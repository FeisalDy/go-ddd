package rest

import (
	"net/http"
	"strconv"

	appif "github.com/FeisalDy/go-ddd/internal/application/interfaces"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest/dto"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserRoleHandler struct {
	service   appif.UserRoleService
	validator *validator.Validate
}

func NewUserRoleHandler(service appif.UserRoleService) *UserRoleHandler {
	return &UserRoleHandler{service: service, validator: validator.New()}
}

func (h *UserRoleHandler) AssignRoles(c *gin.Context) {
	userIDParam := c.Param("id")
	uid64, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}
	var req dto.AssignRolesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.validator.Struct(req); err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	roles, err := h.service.AssignRoles(uint(uid64), req.RoleIDs)
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, roles)
}

func (h *UserRoleHandler) ListUserRoles(c *gin.Context) {
	userIDParam := c.Param("id")
	uid64, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}
	roles, err := h.service.ListUserRoles(uint(uid64))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, roles)
}

func (h *UserRoleHandler) RemoveRole(c *gin.Context) {
	userIDParam := c.Param("id")
	uid64, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "invalid user id")
		return
	}
	roleIDParam := c.Param("roleId")
	rid64, err := strconv.ParseUint(roleIDParam, 10, 64)
	if err != nil {
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	roles, err := h.service.RemoveRole(uint(uid64), uint(rid64))
	if err != nil {
		response.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
	response.Success(c, roles)
}
