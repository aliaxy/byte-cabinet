package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New(fiber.Config{
		AppName: "Byte Cabinet",
	})

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "ok",
			"message": "Byte Cabinet is running",
		})
	})

	// API routes group
	api := app.Group("/api")
	v1 := api.Group("/v1")

	// TODO: Register routes here
	v1.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Byte Cabinet API v1",
		})
	})

	// Get port from environment variable or use default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server
	log.Printf("🚀 Server starting on port %s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
