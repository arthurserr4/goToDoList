package handler

import (
	"net/http"

	"github.com/arthurserr4/goToDoList/internal/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Create an todo
// @Description Create a new job todo
// @Tags Todos
// @Accept json
// @Produce json
// @Param request body CreateTodoRequest true "Request body"
// @Success 200 {object} CreateTodoResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /todo [post]
func CreateTodoHandler(ctx *gin.Context) {
	request := CreateTodoRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error()) // Diferenciando a API: gerenciando os erros e dando respostas
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	todo := schemas.Todo{
		Name:        request.Name,
		Description: request.Description,
		IsDone:      *request.IsDone,
	}

	if err := db.Create(&todo).Error; err != nil {
		logger.Errorf("error creating todo: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating todo on database")
		return
	}

	sendSuccess(ctx, "create-todo", todo)
}
