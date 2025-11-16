package dtos

import "github.com/juancuadra2/crud-api-go/internal/core/domain/models"

type UpdateUserDto struct {
	Name     *string `json:"name,omitempty"`
	Email    *string `json:"email,omitempty"`
	Password *string `json:"password,omitempty"`
}

func (dto *UpdateUserDto) ToDomain() *models.User {
	user := &models.User{}

	if dto.Name != nil {
		user.Name = *dto.Name
	}

	if dto.Email != nil {
		user.Email = *dto.Email
	}

	return user
}
