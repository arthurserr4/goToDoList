package handler

import (
	"fmt"
	"net/http"

	"github.com/arthurserr4/goToDoList/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Delete todo
// @Description Delete a new job todo
// @Tags Todos
// @Accept json
// @Produce json
// @Param id query string true "Todo identification"
// @Success 200 {object} DeleteTodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /todo [delete]
func DeleteTodoHandler(ctx *gin.Context) { // O handler vai lidar com essa requisição
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequirered("id", "queryParameter").Error())
		return
	}
	todo := schemas.Todo{}
	// Find Todo
	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("todo with id %s not found", id))
		return
	}
	// Delete Todo
	if err := db.Delete(&todo).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting todo with id %s", id))
		return
	}
	sendSuccess(ctx, "delete-todo", todo)
}
