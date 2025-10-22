package rest

import (
	"net/http"
	"strconv"

	"github.com/FeisalDy/go-ddd/internal/application/interfaces"
	"github.com/gin-gonic/gin"
)

type RoleHandler struct {
	service interfaces.RoleService
}

func NewRoleHandler(service interfaces.RoleService) *RoleHandler {
	return &RoleHandler{
		service: service,
	}
}

func (h *RoleHandler) GetAll(c *gin.Context) {
	roles, err := h.service.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, roles)

}

func (h *RoleHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	parsedUint64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format: " + err.Error()})
		return
	}

	roleID := uint(parsedUint64)
	role, err := h.service.FindByID(roleID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}
