package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

type AppConfig struct {
	Environment    string
	Port           string
	BaseURL        string
	Timezone       string
	LogLevel       string
	Debug          bool
	RequestTimeout time.Duration
	JWTSecret      string
}

type Config struct {
	DB  DBConfig
	App AppConfig
}

func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	return Config{
		App: LoadAppConfig(),
		DB:  LoadDBConfig(),
	}
}

func LoadAppConfig() AppConfig {
	timeout, err := strconv.Atoi(getEnv("REQUEST_TIMEOUT_SECONDS", "30"))
	if err != nil {
		timeout = 30
	}

	debug, err := strconv.ParseBool(getEnv("DEBUG", "false"))
	if err != nil {
		debug = false
	}

	return AppConfig{
		Environment:    getEnv("ENVIRONMENT", "development"),
		Port:           getEnv("PORT", "8080"),
		BaseURL:        getEnv("BASE_URL", "http://localhost:8080"),
		Timezone:       getEnv("TIMEZONE", "UTC"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		Debug:          debug,
		RequestTimeout: time.Duration(timeout) * time.Second,
		JWTSecret:      getEnv("JWT_SECRET_KEY", "your-secret-key-change-in-production"),
	}
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "postgres"),
	}
}

func InitializeApp(config AppConfig) {
	if err := setTimezone(config.Timezone); err != nil {
		panic(fmt.Sprintf("failed to set timezone: %v", err))
	} else {
		log.Printf("Timezone set to: %s", config.Timezone)
	}

	setGinMode(config.Environment)

	log.Printf("Application initialized:")
	log.Printf("  Environment: %s", config.Environment)
	log.Printf("  Port: %s", config.Port)
	log.Printf("  BaseURL: %s", config.BaseURL)
	log.Printf("  Debug: %t", config.Debug)
	log.Printf("  LogLevel: %s", config.LogLevel)
}

func setTimezone(timezone string) error {
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		return err
	}

	time.Local = loc

	os.Setenv("TZ", timezone)

	return nil
}

func setGinMode(environment string) {
	switch environment {
	case "production":
		os.Setenv("GIN_MODE", "release")
		log.Println("Gin mode set to: release")
	case "staging":
		os.Setenv("GIN_MODE", "test")
		log.Println("Gin mode set to: test")
	default:
		os.Setenv("GIN_MODE", "debug")
		log.Println("Gin mode set to: debug")
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
