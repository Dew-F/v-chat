package config

import (
	"fmt"
	"os"
)

type Config struct {
	Port              string
	PostgresURL       string
	RedisURL          string
	JWTSecret         string
	SnowflakeWorkerID int64
}

func Load() Config {
	return Config{
		Port: ":8080",

		PostgresURL: fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			getenv("DB_USER", "postgres"),
			getenv("DB_PASSWORD", "postgres"),
			getenv("DB_HOST", "localhost"),
			getenv("DB_PORT", "5432"),
			getenv("DB_NAME", "postgres"),
		),

		RedisURL: fmt.Sprintf("redis://%s:%s",
			getenv("REDIS_HOST", "localhost"),
			getenv("REDIS_PORT", "6379"),
		),

		JWTSecret: getenv("JWT_SECRET", "secret"),

		SnowflakeWorkerID: 1,
	}
}

func getenv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return fallback
}
