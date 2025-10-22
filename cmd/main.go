package main

import (
	"log"

	"github.com/FeisalDy/go-ddd/config"
	"github.com/FeisalDy/go-ddd/internal/application/services"
	"github.com/FeisalDy/go-ddd/internal/infrastructure/db/postgres"
	"github.com/FeisalDy/go-ddd/internal/infrastructure/security"
	"github.com/FeisalDy/go-ddd/internal/interface/api/rest"
)

func main() {
	cfg := config.LoadConfig()
	config.InitializeApp(cfg.App)
	postgres.Init(cfg.DB)

	hasher := security.BcryptHasher{}
	jwtService := security.NewJWTService(cfg.App.JWTSecret)

	userRepo := postgres.NewUserRepository(postgres.DB)
	roleRepo := postgres.NewRoleRepository(postgres.DB)
	userService := services.NewUserService(userRepo, hasher)
	roleService := services.NewRoleService(roleRepo)
	authService := services.NewAuthService(userRepo, hasher, jwtService)

	userHandler := rest.NewUserHandler(userService)
	roleHandler := rest.NewRoleHandler(roleService)
	authHandler := rest.NewAuthHandler(authService)

	r := rest.SetupRoutes(userHandler, authHandler, roleHandler, jwtService)
	serverAddr := ":" + cfg.App.Port

	if err := r.Run(serverAddr); err != nil {
		panic("failed to start server: " + err.Error())
	} else {
		log.Printf("Server started on %s", serverAddr)
		log.Printf("Environment: %s", cfg.App.Environment)
	}
}
