package middleware

import (
	"net/http"
	"strings"

	"github.com/FeisalDy/go-ddd/internal/domain/services"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest/response"

	"github.com/gin-gonic/gin"
)

func JWTAuth(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, http.StatusUnauthorized, "authorization header is required")
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, http.StatusUnauthorized, "invalid authorization header format")
			return
		}

		token := parts[1]
		userID, err := jwtService.ValidateToken(token)
		if err != nil {
			response.Error(c, http.StatusUnauthorized, "invalid or expired token")
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
