package controllers

import (
	"userapi-ddd-go/pkg/controllers/user"

	"github.com/gin-gonic/gin"

	_ "userapi-ddd-go/docs"

	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func Router(ctrl *user.UserController) *gin.Engine {

	r := gin.Default()

	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.POST("user", ctrl.AddUser)
	}

	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
