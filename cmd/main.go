package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rickyhuang08/99-backend-exercise/config"
	"github.com/rickyhuang08/99-backend-exercise/delivery/http"
	"github.com/rickyhuang08/99-backend-exercise/helpers"
	"github.com/rickyhuang08/99-backend-exercise/internal/repository/sql"
	"github.com/rickyhuang08/99-backend-exercise/internal/usecase"
	"github.com/rickyhuang08/99-backend-exercise/middleware"
	"github.com/rickyhuang08/99-backend-exercise/pkg/auth"
	"github.com/rickyhuang08/99-backend-exercise/pkg/database/sqldb"
)

func startServer() error {
	// Load configuration
	cfg, err := config.NewConfig()
	log.Printf("config : %+v", cfg)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	db, err := sqldb.Init(&sqldb.SQLiteAdapter{}, cfg.SQLite.Route)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}

	// Initialize Gin router
	router := gin.Default()

	// Initialize helper for time abstraction
	timeProvider := helpers.NewRealTimeProvider()

	// Initialize middleware module
	mw := middleware.NewMiddlewareModule(timeProvider, cfg.Jwt.PublicKey)

	// Initialize repository
	userRepo := sql.NewUserRepository(db, timeProvider)
	listingRepo := sql.NewListingRepository(db, timeProvider)

	jwtHelper := auth.NewJWTHelper(timeProvider, cfg.Jwt.PrivateKey)
	authUC := usecase.NewAuthUsecase(userRepo, jwtHelper)
	userUC := usecase.NewUserUsecase(userRepo)
	listingUC := usecase.NewListingUsecase(listingRepo)

	// Initialize handler
	handler := http.NewHandler(authUC, userUC, listingUC)

	// Register routes
	http.RegisterRoutes(router, handler, mw)

	// Start server
	log.Printf("Server running on port %s", cfg.Server.Port)
	return router.Run(":" + cfg.Server.Port)
}

func main() {
	// Handle errors properly
	if err := startServer(); err != nil {
		log.Fatal(err) // Only main.go decides to stop the app
	}
}
