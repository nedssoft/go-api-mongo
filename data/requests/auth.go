package requests

import "github.com/nedssoft/go-api-mongo/data/responses"

type LoginPayload struct {
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	User responses.UserResponse `json:"user"`
}
