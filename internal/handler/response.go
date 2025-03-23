package handler

import (
	"fmt"
	"net/http"

	"github.com/arthurserr4/goToDoList/schemas"
	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json") // Já está indicando que o conteúdo é JSON
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successful", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode int    `json:"errorCode"`
}

type CreateTodoResponse struct {
	Message string               `json:"message"`
	Data    schemas.TodoResponse `json:"data"`
}

type DeleteTodoResponse struct {
	Message string               `json:"message"`
	Data    schemas.TodoResponse `json:"data"`
}

type ShowTodoResponse struct {
	Message string               `json:"message"`
	Data    schemas.TodoResponse `json:"data"`
}

type ListTodosResponse struct {
	Message string                 `json:"message"`
	Data    []schemas.TodoResponse `json:"data"`
}

type UpdateTodoResponse struct {
	Message string               `json:"message"`
	Data    schemas.TodoResponse `json:"data"`
}
