package handler

import (
	"github.com/arthurserr4/goToDoList/internal/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	db = config.GetSQLite()
}
