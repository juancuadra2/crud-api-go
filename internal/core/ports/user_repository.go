package ports

import (
	"context"

	"github.com/juancuadra2/crud-api-go/internal/core/domain/models"
)

// UserRepository - Puerto de salida - Define lo que el servicio necesita
type UserRepository interface {
	Create(ctx context.Context, user *models.User) (*models.User, error)
	FindByID(ctx context.Context, id string) (*models.User, error)
	FindAll(ctx context.Context) ([]*models.User, error)
	Update(ctx context.Context, user *models.User) (*models.User, error)
	Delete(ctx context.Context, id string) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}
