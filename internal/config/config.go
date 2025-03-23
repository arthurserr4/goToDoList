package config

import (
	"fmt"
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("error initializing SQLite: %v", err)
	}

	return nil
}
