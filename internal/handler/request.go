package handler

import "fmt"

func errParamIsRequirered(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

// CreateTodoRequest

type CreateTodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDone      *bool  `json:"is_done"`
}

func (r *CreateTodoRequest) Validate() error {
	if r.Name == "" {
		return fmt.Errorf("name is required")
	}
	if r.Description == "" {
		return fmt.Errorf("description is required")
	}
	if r.IsDone == nil {
		return fmt.Errorf("is_done is required")
	}
	return nil
}

// UpdateTodoRequest

type UpdateTodoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDone      *bool  `json:"is_done"`
}

func (r *UpdateTodoRequest) Validate() error {
	// If any field is provided, validation is successful
	if r.Name != "" || r.Description != "" || r.IsDone != nil {
		return nil
	}
	// If none of the fields were provided, return error
	return fmt.Errorf("at least one field must be provided")
}
