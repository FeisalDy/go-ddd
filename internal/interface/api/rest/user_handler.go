package rest

import (
	"net/http"

	appDto "github.com/FeisalDy/go-ddd/internal/application/dto"
	"github.com/FeisalDy/go-ddd/internal/application/interfaces"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	service   interfaces.UserService
	validator *validator.Validate
}

func NewUserHandler(service interfaces.UserService) *UserHandler {
	return &UserHandler{service: service, validator: validator.New()}
}

func (h *UserHandler) Register(c *gin.Context) {
	var request dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.validator.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createUserRequest := &appDto.CreateUserRequest{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	user, err := h.service.CreateUser(createUserRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}
