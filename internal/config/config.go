package config

import (
	"log"
	"os"
)

type Config struct {
	DBURL string
	Port  string
	Env   string
}

func Load() Config {
	cfg := Config{
		DBURL: getenv("DATABASE_URL", "postgres://admin:admin@localhost:5432/nubank?sslmode=disable"),
		Port:  getenv("PORT", "8080"),
		Env:   getenv("ENV", "development"),
	}
	return cfg 
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func MustHave(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("Missing required environment variable %s", key)
	}
	return v
}
