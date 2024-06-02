package handlers

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New(validator.WithRequiredStructEnabled())

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=8,containsany=!@#$%^&*()_+"`
	RoleIds  []string `json:"role_ids" validate:"required,dive,uuid"`
}
