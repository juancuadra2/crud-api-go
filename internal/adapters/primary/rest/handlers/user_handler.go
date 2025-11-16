package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juancuadra2/crud-api-go/internal/adapters/primary/rest/dtos"
	"github.com/juancuadra2/crud-api-go/internal/core/ports"
)

const (
	IdRequiredError = "ID is required"
)

type UserHandler struct {
	service ports.UserService
}

func NewUserHandler(service ports.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// Create - Add methods to handle HTTP requests
func (handler *UserHandler) Create(c *gin.Context) {
	var createUserDTO dtos.CreateUserDTO
	if err := c.ShouldBindJSON(&createUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponseDTO{Message: err.Error()})
		return
	}

	user, err := handler.service.Create(c.Request.Context(), createUserDTO.ToDomain())
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponseDTO{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dtos.FromDomain(user))
}

// GetAll - Add methods to handle HTTP requests
func (handler *UserHandler) GetAll(c *gin.Context) {
	users, err := handler.service.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dtos.ErrorResponseDTO{Message: err.Error()})
		return
	}

	response := make([]dtos.UserDTO, len(users))
	for i, user := range users {
		response[i] = dtos.FromDomain(user)
	}

	c.JSON(http.StatusOK, response)
}

// GetByID - Add methods to handle HTTP requests
func (handler *UserHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponseDTO{Message: IdRequiredError})
		return
	}

	user, err := handler.service.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, dtos.ErrorResponseDTO{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.FromDomain(user))
}

// Update - Add methods to handle HTTP requests
func (handler *UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponseDTO{Message: IdRequiredError})
		return
	}

	var updateUserDTO dtos.UpdateUserDto
	if err := c.ShouldBindJSON(&updateUserDTO); err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponseDTO{Message: err.Error()})
		return
	}

	user, err := handler.service.Update(c.Request.Context(), updateUserDTO.ToDomain(), id)
	if err != nil {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponseDTO{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dtos.FromDomain(user))
}

// Delete - Add methods to handle HTTP requests
func (handler *UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, dtos.ErrorResponseDTO{Message: IdRequiredError})
		return
	}

	err := handler.service.Delete(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, dtos.ErrorResponseDTO{Message: err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
