package dtos

import "github.com/juancuadra2/crud-api-go/internal/core/domain/models"

type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func FromDomain(user *models.User) UserDTO {
	return UserDTO{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
