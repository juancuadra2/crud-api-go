package dtos

import "github.com/juancuadra2/crud-api-go/internal/core/domain/models"

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (dto *CreateUserDTO) ToDomain() *models.User {
	return &models.User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}
