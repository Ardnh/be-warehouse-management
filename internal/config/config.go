package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type DatabaseConfig struct {
	Host            string
	Port            string
	User            string
	Password        string
	DBName          string
	SSLMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
}

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
	PoolSize int
}

type Config struct {
	Database DatabaseConfig
	Redis    RedisConfig
	App      AppConfig
}

type AppConfig struct {
	Port      string
	JWTSecret string
}

func LoadConfig() *Config {

	return &Config{
		Database: DatabaseConfig{
			Host:            getRequiredEnv("DB_HOST"),
			Port:            getRequiredEnv("DB_HOST"),
			User:            getRequiredEnv("DB_HOST"),
			Password:        getRequiredEnv("DB_HOST"),
			DBName:          getRequiredEnv("DB_HOST"),
			SSLMode:         getRequiredEnv("DB_HOST"),
			MaxIdleConns:    getEnvAsInt("DB_MAX_IDLE_CONNS"),
			MaxOpenConns:    getRequiredEnv("DB_MAX_OPEN_CONNS"),
			ConnMaxLifetime: getRequiredEnv("DB_HOST"),
		},
		Redis: RedisConfig{},
		App: AppConfig{
			Port:      getRequiredEnv("APP_SERVER_PORT"),
			JWTSecret: getRequiredEnv("APP_JWT_SECRET"),
		},
	}
}

func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}

func (c *RedisConfig) GetAddr() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func getRequiredEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		log.Fatalf("required environment variable %s is not set", key)
	}

	return value
}

func getEnvAsInt(key string) int {
	valueStr := os.Getenv(key)

	if valueStr == "" {
		log.Fatalf("required environment variable %s is not set", key)
	}

	value, err := strconv.Atoi(valueStr)

	if err != nil {
		log.Fatalf("required environment variable %s is not set", key)
	}
	return value
}

func getEnvAsDuration(key string) time.Duration {
	valueStr := os.Getenv(key)

	if valueStr == "" {
		log.Fatalf("required environment variable %s is not set", key)
	}

	value, err := time.ParseDuration(valueStr)

	if err == nil {
		log.Fatalf("required environment variable %s is not set", key)
	}
	return value
}
