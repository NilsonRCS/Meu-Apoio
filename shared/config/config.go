package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port        string
	Environment string

	// Database
	DBHost     string
	DBPort     string
	DBName     string
	DBUser     string
	DBPassword string

	// MongoDB
	MongoURI string

	// Redis
	RedisURL string

	// RabbitMQ
	RabbitMQURL string

	// MinIO
	MinIOEndpoint  string
	MinIOAccessKey string
	MinIOSecretKey string

	// JWT
	JWTSecret string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8080"),
		Environment: getEnv("ENVIRONMENT", "development"),

		// Database
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBName:     getEnv("DB_NAME", "meuapoio"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "postgres123"),

		// MongoDB
		MongoURI: getEnv("MONGO_URI", "mongodb://mongo:mongo123@localhost:27017/meuapoio"),

		// Redis
		RedisURL: getEnv("REDIS_URL", "localhost:6379"),

		// RabbitMQ
		RabbitMQURL: getEnv("RABBITMQ_URL", "amqp://admin:admin123@localhost:5672/"),

		// MinIO
		MinIOEndpoint:  getEnv("MINIO_ENDPOINT", "localhost:9000"),
		MinIOAccessKey: getEnv("MINIO_ACCESS_KEY", "minio"),
		MinIOSecretKey: getEnv("MINIO_SECRET_KEY", "minio123"),

		// JWT
		JWTSecret: getEnv("JWT_SECRET", "sua-chave-secreta-super-segura"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}
