I have tried to implement json-server that npm provides but in Golang without the extra features like embed, expand etc. 
It will only work with a db.json filename just like json-server.
It will create end-points to ease-out development for my frontend bros.

It supports
/GET/endpoint
/GET/endpoint/:id
/POST/endpoint
/PUT/endpoint/:id
/PATCH/endpoint/:id
/DELETE/endpoint/:id

For using it clone it or simply copy the code as it is only in one file as of now. You'll be able to understand the code if you are familiar with Go.
For routing I have used GorillaMux instead of Gin (just a personal pref this time but Gin is better)
