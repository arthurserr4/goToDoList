package main

import (
	"github.com/arthurserr4/goToDoList/internal/config"
	"github.com/arthurserr4/goToDoList/internal/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	router.Initialize()
}
