## Golang Rest Api Starter pack
The project can be used by developers to quickly start a new rest api written in golang.
___
##### get the project
```bach
git pull https://github.com/Negat1v9/golang-api-start
```
##### quick start on port **8088**
```bach 
go run app.go
```
### application layers
1. model
- layer with data models for the database and json responses
2. internal/config
- config model for service
3. internal/web
- web layer with server, all routers and handlers
4. services
- layer for working with layers for a database or external services
