package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/juancuadra2/crud-api-go/internal/core/domain/models"
	"github.com/juancuadra2/crud-api-go/internal/core/ports"
	"golang.org/x/crypto/bcrypt"
)

const (
	userAlreadyExistsError = "user with email %s already exists"
	userNotFoundError      = "user not found"
	idIsRequiredError      = "id is required"
)

// userService - Implementa el puerto de entrada UserService
type userService struct {
	repository ports.UserRepository
}

func NewUserService(repo ports.UserRepository) ports.UserService {
	return &userService{
		repository: repo,
	}
}

func (s *userService) Create(ctx context.Context, user *models.User) (*models.User, error) {
	existingUser, _ := s.repository.FindByEmail(ctx, user.Email)
	if existingUser != nil {
		return nil, errors.New(fmt.Sprintf(userAlreadyExistsError, user.Email))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user.ID = uuid.NewString()
	user.Password = string(hashedPassword)

	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := s.repository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (s *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New(idIsRequiredError)
	}

	user, err := s.repository.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New(userNotFoundError)
	}

	return user, nil
}

func (s *userService) GetAll(ctx context.Context) ([]*models.User, error) {
	userList, err := s.repository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return userList, nil
}

func (s *userService) Update(ctx context.Context, user *models.User, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New(idIsRequiredError)
	}

	if existingUser, _ := s.repository.FindByID(ctx, id); existingUser == nil {
		return nil, errors.New(userNotFoundError)
	}

	user.ID = id

	updatedUser, err := s.repository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (s *userService) Delete(ctx context.Context, id string) error {
	if id == "" {
		return errors.New(idIsRequiredError)
	}

	if existingUser, _ := s.repository.FindByID(ctx, id); existingUser == nil {
		return errors.New(userNotFoundError)
	}

	if err := s.repository.Delete(ctx, id); err != nil {
		return err
	}

	return nil
}
