package ports

import (
	"context"

	"github.com/juancuadra2/crud-api-go/internal/core/domain/models"
)

// UserService - Puerto de entrada - Define las operaciones del caso de uso que el servicio ofrece
type UserService interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	GetByID(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User, id string) (*models.User, error)
	Delete(ctx context.Context, id string) error
}
