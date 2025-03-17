package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando as configuraçoes Default do gin
	router := gin.Default()
	// Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a nossa API
	router.Run("8080") // listen and serve on 0.0.0.0:8080
}
