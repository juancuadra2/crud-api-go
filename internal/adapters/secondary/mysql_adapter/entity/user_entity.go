package entity

import "github.com/juancuadra2/crud-api-go/internal/core/domain/models"

type UserEntity struct {
	ID        string `gorm:"primaryKey;type:char(36)"` // UUID
	Name      string `gorm:"type:varchar(100);not null"`
	Email     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt int64  `gorm:"autoCreateTime"`
	UpdatedAt int64  `gorm:"autoUpdateTime"`
}

func (u *UserEntity) TableName() string {
	return "users"
}

func (u *UserEntity) ToDomain() *models.User {
	return &models.User{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	}
}

func FromDomain(user *models.User) *UserEntity {
	return &UserEntity{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}
}
