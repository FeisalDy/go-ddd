package rest

import (
	"net/http"

	appDto "github.com/FeisalDy/go-ddd/internal/application/dto"
	"github.com/FeisalDy/go-ddd/internal/application/interfaces"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest/dto"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	service   interfaces.AuthService
	validator *validator.Validate
}

func NewAuthHandler(service interfaces.AuthService) *AuthHandler {
	return &AuthHandler{
		service:   service,
		validator: validator.New(),
	}
}

func (h *AuthHandler) Login(c *gin.Context) {
	var request dto.LoginRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := h.validator.Struct(request); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	loginRequest := &appDto.LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}

	res, err := h.service.Login(loginRequest)
	if err != nil {
		response.Error(c, http.StatusUnauthorized, err.Error())
		return
	}

	response.Success(c, res)
}
