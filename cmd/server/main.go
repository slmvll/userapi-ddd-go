package main

import (
	"context"

	user_domain "userapi-ddd-go/pkg/domain/user"

	"userapi-ddd-go/pkg/controllers/user"

	"userapi-ddd-go/pkg/infra/mongodb"

	"userapi-ddd-go/pkg/controllers"
)

// @title 		User API
// @version 	1.0
// @description User API built using Go and Gin with DDD. You can visit the GitHub repository at https://userapi-ddd-go

// @contact.name 	Sergio Lorenzo
// @contact.url 	https://www.linkedin.com/in/sergiolorenzo/
// @contact.email 	slm092013@gmail.com

// @license.name	MIT
// @license.url 	https://opensource.org/licenses/MIT

// @host 		localhost:8080
// @BasePath 	/api/v1
// @query.collection.format multi
func main() {
	userRepository, err := mongodb.NewMongoUserRepository(context.Background(), "mongodb://mongotest:test123@mongo:27017")
	if err != nil {
		panic(err)
	}
	userService := user_domain.NewUserService(userRepository)
	userController := user.NewUserController(userService)

	r := controllers.Router(userController)

	r.Run(":8081")
}
