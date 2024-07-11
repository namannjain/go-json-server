# JSON Server
I have tried to implement json-server that npm provides but in Golang without the extra features like embed, expand etc. 
It will create end-points to ease-out development for my frontend bros.

## Getting started
I have not published it yet so go get command won't work. 
Instead you can clone it or simply copy the code as it's only in one file main.go as of now

Create a `db.json` file with your desired data


    {
      "posts": [
        { 
          "id": "1", 
          "title": "json-server", 
          "author": "namannjain" 
        }
      ],
      "hobbies": [
        "football", "singing", "games", "gym", "go ofc"
      ],
      "users": [
        {
          "id": "qwert",
          "name": "christ"
        },
        {
          "id": "tyui",
          "name": "jain"
        },
        {
          "id": "dfg",
          "name": "sikh"
        },
        {
          "id": "dfghjm",
          "name": "hindu"
        }
      ]
    }
    

Start JSON Server

`go run main.go`

If you navigate to http://localhost:8080/posts/1, you will get

    { 
      "id": "1", 
      "title": "json-server", 
      "author": "namannjain" 
    }

## Routes

````
GET     /<endpoint>
GET     /<endpoint>/:id
POST    /<endpoint>
PUT     /<endpoint>/:id
PATCH   /<endpoint>/:id
DELETE  /<endpoint>/:id
````

## Just FYI
You'll be able to understand the code if you are familiar with Go.
For routing I have used GorillaMux instead of Gin (just a personal pref this time but Gin is better)
Feel free to use and give feedback
