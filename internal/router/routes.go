package router

import (
	docs "github.com/arthurserr4/goToDoList/docs"
	"github.com/arthurserr4/goToDoList/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.POST("/todo", handler.CreateTodoHandler)
		v1.GET("/todo", handler.ShowTodoHandler)
		v1.PUT("/todo", handler.UpdateTodoHandler)
		v1.DELETE("/todo", handler.DeleteTodoHandler)
		v1.GET("/todos", handler.ListTodosHandler)
	}
	// Initialize Swagger
	router.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
