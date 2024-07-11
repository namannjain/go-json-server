package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var db map[string]interface{}

func main() {
	loadDatabase()

	r := mux.NewRouter()
	r.Use(loggingMiddleware)

	r.HandleFunc("/{resource}", handleResource).Methods("GET", "POST")
	r.HandleFunc("/{resource}/{id}", handleResourceItem).Methods("GET", "PUT", "PATCH", "DELETE")

	fmt.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func loadDatabase() {
	file, err := ioutil.ReadFile("db.json")
	if err != nil {
		log.Fatal("Error reading db.json file:", err)
	}
	err = json.Unmarshal(file, &db)
	if err != nil {
		log.Fatal("Error parsing db.json:", err)
	}
}

func saveDatabase() {
	file, _ := json.MarshalIndent(db, "", "  ")
	ioutil.WriteFile("db.json", file, 0644)
}

func handleResource(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resource := vars["resource"]

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(db[resource])
	case "POST":
		var newData interface{}
		json.NewDecoder(r.Body).Decode(&newData)

		// Handle different types of resources
		switch resourceData := db[resource].(type) {
		case []interface{}:
			db[resource] = append(resourceData, newData)
		case map[string]interface{}:
			if newMap, ok := newData.(map[string]interface{}); ok {
				for k, v := range newMap {
					resourceData[k] = v
				}
			}
		default:
			db[resource] = newData
		}

		saveDatabase()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newData)
	}
}

func handleResourceItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	resource := vars["resource"]
	id := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	resourceData, ok := db[resource].([]interface{})
	if !ok {
		http.Error(w, "Resource not found or not an array", http.StatusNotFound)
		return
	}

	index, err := findItemIndex(resourceData, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(resourceData[index])
	case "PUT", "PATCH":
		var updatedData map[string]interface{}
		json.NewDecoder(r.Body).Decode(&updatedData)

		if r.Method == "PATCH" {
			// For PATCH, we merge the updated data with the existing data
			existingData := resourceData[index].(map[string]interface{})
			for k, v := range updatedData {
				existingData[k] = v
			}
			updatedData = existingData
		}

		resourceData[index] = updatedData
		db[resource] = resourceData
		saveDatabase()
		json.NewEncoder(w).Encode(updatedData)
	case "DELETE":
		db[resource] = append(resourceData[:index], resourceData[index+1:]...)
		saveDatabase()
		w.WriteHeader(http.StatusNoContent)
	}
}

func findItemIndex(resourceData []interface{}, id string) (int, error) {
	for i, item := range resourceData {
		if m, ok := item.(map[string]interface{}); ok {
			if fmt.Sprintf("%v", m["id"]) == id {
				return i, nil
			}
		}
	}
	return -1, fmt.Errorf("item not found")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
