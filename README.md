# User API in GO (DDD)

User API allows the client to create a user or get user by its id.

This project is sample project that shows a REST API 
with domain driven approach (DDD) with hexagonal architecture in Golang and Gin.

## Persistence

1. MongoDB

## Requirements

1. Go 1.22
1. Gin
1. Testify
1. Mockery
1. Swagger

## Usage

`make run`

## Build

`make build`

## Using Docker

To build the app with docker, run

`docker build -t userapi .`

`docker run -d -p 8080:80 --name userapiapp userapi`

## Mocks

Generate mocks with mockery: 

```
$ mockery --dir=pkg/domain/user --name=UserRepository --filename=service_mock.go --output=pkg/domain/user/mocks --outpkg=mocks
```

## Swagger

Generate swagger documentation: 

```
$ go get github.com/swaggo/swag/cmd/swag@latest
$ swag init -g cmd/server/main.go
```

URL: http://localhost:8081/docs/index.html