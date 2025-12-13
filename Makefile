.PHONY: all build run clean test dev help install-deps

# Variables
APP_NAME := byte-cabinet
SERVER_DIR := cmd/server
BUILD_DIR := bin
WEB_DIR := web

# Go related variables
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOTEST := $(GOCMD) test
GOMOD := $(GOCMD) mod

# Default target
all: build

# Install Go dependencies
install-deps:
	@echo "📦 Installing Go dependencies..."
	$(GOMOD) download
	$(GOMOD) tidy

# Build the server binary
build:
	@echo "🔨 Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/server ./$(SERVER_DIR)
	@echo "✅ Build complete: $(BUILD_DIR)/server"

# Run the server
run:
	@echo "🚀 Starting server..."
	$(GORUN) ./$(SERVER_DIR)/main.go

# Run tests
test:
	@echo "✅ Running tests..."
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	@echo "📊 Running tests with coverage..."
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "📄 Coverage report: coverage.html"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "✅ Clean complete"

# Development mode with hot reload (requires air)
dev:
	@echo "🔥 Starting development server with hot reload..."
	@which air > /dev/null || (echo "❌ 'air' not found. Install with: go install github.com/air-verse/air@latest" && exit 1)
	air

# Install development tools
install-tools:
	@echo "🛠️ Installing development tools..."
	go install github.com/air-verse/air@latest
	@echo "✅ Tools installed"

# Frontend commands
web-install:
	@echo "📦 Installing frontend dependencies..."
	cd $(WEB_DIR) && pnpm install

web-dev:
	@echo "🌐 Starting frontend dev server..."
	cd $(WEB_DIR) && pnpm dev

web-build:
	@echo "🔨 Building frontend..."
	cd $(WEB_DIR) && pnpm build

# Database migrations (placeholder)
migrate-up:
	@echo "⬆️ Running migrations..."
	@echo "TODO: Implement migrations"

migrate-down:
	@echo "⬇️ Rolling back migrations..."
	@echo "TODO: Implement migrations"

# Help
help:
	@echo "Byte Cabinet - Available Commands"
	@echo ""
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all            Build the application (default)"
	@echo "  build          Build the server binary"
	@echo "  run            Run the server"
	@echo "  dev            Run with hot reload (requires air)"
	@echo "  test           Run tests"
	@echo "  test-coverage  Run tests with coverage report"
	@echo "  clean          Remove build artifacts"
	@echo "  install-deps   Install Go dependencies"
	@echo "  install-tools  Install development tools (air)"
	@echo ""
	@echo "Frontend:"
	@echo "  web-install    Install frontend dependencies"
	@echo "  web-dev        Start frontend dev server"
	@echo "  web-build      Build frontend for production"
	@echo ""
	@echo "Database:"
	@echo "  migrate-up     Run database migrations"
	@echo "  migrate-down   Rollback database migrations"
