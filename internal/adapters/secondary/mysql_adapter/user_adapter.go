package mysql_adapter

import (
	"context"

	"github.com/juancuadra2/crud-api-go/internal/adapters/secondary/mysql_adapter/entity"
	"github.com/juancuadra2/crud-api-go/internal/core/domain/models"
	"gorm.io/gorm"
)

type UserAdapter struct {
	db *gorm.DB
}

func NewUserAdapter(db *gorm.DB) *UserAdapter {
	return &UserAdapter{db: db}
}

func (adapter *UserAdapter) Create(ctx context.Context, user *models.User) (*models.User, error) {
	userEntity := entity.FromDomain(user)
	err := adapter.db.WithContext(ctx).Create(userEntity).Error
	if err != nil {
		return nil, err
	}
	// Retorna una copia del usuario creado para evitar problemas con punteros
	userCreated := *userEntity.ToDomain()
	return &userCreated, nil
}

func (adapter *UserAdapter) FindByID(ctx context.Context, id string) (*models.User, error) {
	var user entity.UserEntity
	err := adapter.db.WithContext(ctx).First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return user.ToDomain(), nil
}

func (adapter *UserAdapter) FindAll(ctx context.Context) ([]*models.User, error) {
	var users []*entity.UserEntity
	err := adapter.db.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}

	userList := make([]*models.User, len(users))
	for i, user := range users {
		userList[i] = user.ToDomain()
	}

	return userList, nil
}

func (adapter *UserAdapter) Update(ctx context.Context, user *models.User) (*models.User, error) {
	model := entity.FromDomain(user)
	err := adapter.db.WithContext(ctx).Save(model).Error
	if err != nil {
		return nil, err
	}
	updatedUser := *model.ToDomain()
	return &updatedUser, nil
}

func (adapter *UserAdapter) Delete(ctx context.Context, id string) error {
	return adapter.db.WithContext(ctx).Delete(&entity.UserEntity{}, "id = ?", id).Error
}

func (adapter *UserAdapter) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var model entity.UserEntity
	err := adapter.db.WithContext(ctx).First(&model, "email = ?", email).Error
	if err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}
