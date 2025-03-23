package handler

import (
	"net/http"

	"github.com/arthurserr4/goToDoList/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary List todos
// @Description List all job todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {object} ListTodosResponse
// @Failure 500 {object} ErrorResponse
// @Router /todos [get]
func ListTodosHandler(ctx *gin.Context) {
	todos := []schemas.Todo{}

	if err := db.Find(&todos).Error; err != nil { // O db.find() é um método do GORM que faz uma query no banco de dados e atribui as alterações no objeto que foi passado como parâmetro
		sendError(ctx, http.StatusInternalServerError, "error listing todos")
		return
	}

	sendSuccess(ctx, "list-todos", todos)
}
