package schemas

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string
	Description string
	IsDone      bool
}

type TodoResponse struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
}
