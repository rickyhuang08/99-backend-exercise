package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	SQLite struct {
		Route string `yaml:"route"` // SQLite database file path
	} `yaml:"sqlite"`
	Jwt struct {
		PublicKey  string `yaml:"publickey_path"`
		PrivateKey string `yaml:"privatekey_path"`
	} `yaml:"jwt"`
}

func NewConfig() (*Config, error) {

	// 1. Get the environment (default: development)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development" // Default to development
	}

	// 2. Set the corresponding config file
	configFile := fmt.Sprintf("config/config.%s.yaml", env)

	// Load YAML config
	file, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read %s: %w", configFile, err)
	}

	// 4. Parse YAML
	var cfg Config
	if err := yaml.Unmarshal(file, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse YAML: %w", err)
	}

	cfg, err = LoadEnv(cfg, env)
	if err != nil {
		return nil, fmt.Errorf("failed to load env: %w", err)
	}

	return &cfg, nil
}

func LoadEnv(cfg Config, env string) (Config, error) {
	if env == "" {
		env = "development"
	}

	// Load the corresponding .env file
	envFile := ".env." + env
	if err := godotenv.Load(envFile); err != nil {
		return Config{}, fmt.Errorf("Error loading %s file", envFile)
	}

	log.Printf("Loaded environment variables from %s", envFile)

	if user := os.Getenv("DB_USER"); user != "" {
		cfg.Database.User = user
	}
	if pass := os.Getenv("DB_PASS"); pass != "" {
		cfg.Database.Password = pass
	}

	return cfg, nil
}