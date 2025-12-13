package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

// Config holds all configuration for the application
type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
	JWT      JWTConfig      `mapstructure:"jwt"`
	Upload   UploadConfig   `mapstructure:"upload"`
	Blog     BlogConfig     `mapstructure:"blog"`
}

// ServerConfig holds server-related configuration
type ServerConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
	Mode string `mapstructure:"mode"` // development, production
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Path   string `mapstructure:"path"`
}

// JWTConfig holds JWT-related configuration
type JWTConfig struct {
	Secret          string        `mapstructure:"secret"`
	AccessTokenTTL  time.Duration `mapstructure:"access_token_ttl"`
	RefreshTokenTTL time.Duration `mapstructure:"refresh_token_ttl"`
}

// UploadConfig holds file upload configuration
type UploadConfig struct {
	Path         string   `mapstructure:"path"`
	MaxSize      int64    `mapstructure:"max_size"`
	AllowedTypes []string `mapstructure:"allowed_types"`
}

// BlogConfig holds blog-related configuration
type BlogConfig struct {
	Title        string `mapstructure:"title"`
	Description  string `mapstructure:"description"`
	Author       string `mapstructure:"author"`
	URL          string `mapstructure:"url"`
	PostsPerPage int    `mapstructure:"posts_per_page"`
}

// Load reads configuration from file and environment variables
func Load(configPath string) (*Config, error) {
	v := viper.New()

	// Set default values
	setDefaults(v)

	// Set config file
	if configPath != "" {
		v.SetConfigFile(configPath)
	} else {
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".")
		v.AddConfigPath("./config")
	}

	// Read environment variables
	v.AutomaticEnv()

	// Read config file
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("failed to read config file: %w", err)
		}
		// Config file not found, use defaults and env vars
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	// Validate required fields
	if err := validate(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// setDefaults sets default configuration values
func setDefaults(v *viper.Viper) {
	// Server defaults
	v.SetDefault("server.host", "0.0.0.0")
	v.SetDefault("server.port", 3000)
	v.SetDefault("server.mode", "development")

	// Database defaults
	v.SetDefault("database.driver", "sqlite")
	v.SetDefault("database.path", "./data/byte-cabinet.db")

	// JWT defaults
	v.SetDefault("jwt.secret", "")
	v.SetDefault("jwt.access_token_ttl", "15m")
	v.SetDefault("jwt.refresh_token_ttl", "168h") // 7 days

	// Upload defaults
	v.SetDefault("upload.path", "./uploads")
	v.SetDefault("upload.max_size", 10485760) // 10MB
	v.SetDefault("upload.allowed_types", []string{
		"image/jpeg",
		"image/png",
		"image/gif",
		"image/webp",
	})

	// Blog defaults
	v.SetDefault("blog.title", "Byte Cabinet")
	v.SetDefault("blog.description", "Personal technical notes and learning experiences")
	v.SetDefault("blog.author", "")
	v.SetDefault("blog.url", "http://localhost:3000")
	v.SetDefault("blog.posts_per_page", 10)
}

// validate checks required configuration values
func validate(cfg *Config) error {
	if cfg.JWT.Secret == "" {
		return fmt.Errorf("jwt.secret is required")
	}
	return nil
}

// Address returns the server address in host:port format
func (c *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// IsDevelopment returns true if running in development mode
func (c *ServerConfig) IsDevelopment() bool {
	return c.Mode == "development"
}

// IsProduction returns true if running in production mode
func (c *ServerConfig) IsProduction() bool {
	return c.Mode == "production"
}
