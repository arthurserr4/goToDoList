package handler

import (
	"net/http"

	"github.com/arthurserr4/goToDoList/internal/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Show todo
// @Description Show a new job todo
// @Tags Todos
// @Accept json
// @Produce json
// @Param id query string true "Todo identification"
// @Success 200 {object} ShowTodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /todo [get]
func ShowTodoHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequirered("id", "queryParameter").Error())
		return
	}
	todo := schemas.Todo{}
	if err := db.First(&todo, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "todo not found")
		return
	}

	sendSuccess(ctx, "show-todo", todo)
}
