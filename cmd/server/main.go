package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/aliaxy/byte-cabinet/internal/config"
	"github.com/aliaxy/byte-cabinet/internal/database"
	"github.com/aliaxy/byte-cabinet/internal/handler"
	"github.com/aliaxy/byte-cabinet/internal/middleware"
	"github.com/aliaxy/byte-cabinet/internal/repository"
	"github.com/aliaxy/byte-cabinet/internal/service"
	"github.com/aliaxy/byte-cabinet/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Load configuration
	cfg, err := config.Load("")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Initialize database
	db, err := database.New(cfg.Database.Path)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Run migrations
	log.Println("📦 Running database migrations...")
	if err := db.Migrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	log.Println("✅ Database migrations completed")

	// Initialize JWT manager
	jwtManager := utils.NewJWTManager(
		cfg.JWT.Secret,
		cfg.JWT.AccessTokenTTL,
		cfg.JWT.RefreshTokenTTL,
	)

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.DB)

	// Initialize services
	authService := service.NewAuthService(userRepo, jwtManager)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		AppName:      cfg.Blog.Title,
		ErrorHandler: customErrorHandler,
	})

	// Global middleware
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: false,
	}))

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"data": fiber.Map{
				"status":  "healthy",
				"service": cfg.Blog.Title,
			},
		})
	})

	// API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// Create auth middleware
	authMiddleware := middleware.AuthMiddleware(jwtManager)

	// Register routes
	authHandler.RegisterRoutes(v1, authMiddleware)

	// API welcome route
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"success": true,
			"data": fiber.Map{
				"message": "Welcome to " + cfg.Blog.Title + " API v1",
				"version": "1.0.0",
			},
		})
	})

	// Graceful shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		log.Println("🛑 Shutting down server...")
		if err := app.Shutdown(); err != nil {
			log.Printf("Error during shutdown: %v", err)
		}
	}()

	// Start server
	addr := cfg.Server.Address()
	log.Printf("🚀 Server starting on %s", addr)
	log.Printf("📝 Blog: %s", cfg.Blog.Title)
	log.Printf("🔧 Mode: %s", cfg.Server.Mode)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// customErrorHandler handles errors globally
func customErrorHandler(c *fiber.Ctx, err error) error {
	// Default status code
	code := fiber.StatusInternalServerError
	message := "Internal Server Error"

	// Check if it's a Fiber error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
	}

	return c.Status(code).JSON(fiber.Map{
		"success": false,
		"error": fiber.Map{
			"code":    code,
			"message": message,
		},
	})
}
