package handler

import (
	"net/http"

	"github.com/arthurserr4/goToDoList/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Update todo
// @Description Update a job todo
// @Tags Todos
// @Accept json
// @Produce json
// @Param id query string true "Todo identification"
// @Param todo body UpdateTodoRequest true "Todo data to update"
// @Success 200 {object} UpdateTodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /todo [put]
func UpdateTodoHandler(ctx *gin.Context) {
	request := UpdateTodoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

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
	// Update Todo
	if request.Name != "" {
		todo.Name = request.Name
	}
	if request.Description != "" {
		todo.Description = request.Description
	}
	if request.IsDone != nil {
		todo.IsDone = *request.IsDone
	}
	// Save todo
	if err := db.Save(&todo).Error; err != nil {
		logger.Errorf("error updating todo: %v", err.Error()) // Não esquecer de logar os erros
		sendError(ctx, http.StatusInternalServerError, "error updating todo")
		return
	}
	sendSuccess(ctx, "update-todo", todo)
}
