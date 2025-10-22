package rest

import (
	"github.com/FeisalDy/go-ddd/internal/domain/services"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(userHandler *UserHandler, authHandler *AuthHandler, roleHandler *RoleHandler, userRoleHandler *UserRoleHandler, jwtService services.JWTService) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
				"status":  "healthy",
			})
		})

		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		users := v1.Group("/users")
		{
			users.POST("/register", userHandler.Register)
			users.GET("/:id/roles", userRoleHandler.ListUserRoles)
			users.POST("/:id/roles", userRoleHandler.AssignRoles)
			users.DELETE("/:id/roles/:roleId", userRoleHandler.RemoveRole)
		}

		roles := v1.Group("/roles")
		{
			roles.GET("/", roleHandler.GetAll)
			roles.GET("/:id", roleHandler.GetById)
		}

		protected := v1.Group("")
		protected.Use(middleware.JWTAuth(jwtService))
		{
			protected.GET("/profile", func(c *gin.Context) {
				userID := c.GetUint("user_id")
				c.JSON(200, gin.H{
					"message": "This is a protected route",
					"user_id": userID,
				})
			})
		}
	}

	return r
}
