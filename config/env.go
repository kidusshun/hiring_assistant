package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	Port string
	DBUser string
	DBPassword string
	DBAddress string
	DBName string
}

type JWTConfig struct {
	JWTSecret string
}

type GoogleClientConfig struct {
	GoogleClientID string
	GoogleClientSecret string
}

type GeminiConfig struct {
	GeminiAPIKey string
}

func initDatabaseConfig() DBConfig {
	godotenv.Load()

	return DBConfig{
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", ""), getEnv("DB_PORT", "")),
		DBName:     getEnv("DB_NAME", "planLog"),
	}
}

func initJWTConfig() JWTConfig {
	godotenv.Load()

	return JWTConfig{
		JWTSecret: getEnv("JWT_SECRET", ""),
	}
}

func initGoogleClientConfig() GoogleClientConfig {
	godotenv.Load()

	return GoogleClientConfig{
		GoogleClientID:     getEnv("GOOGLE_CLIENT_ID", ""),
		GoogleClientSecret: getEnv("GOOGLE_CLIENT_SECRET", ""),
	}
}


func initGeminiAPIConfig() GeminiConfig {
	godotenv.Load()
	return GeminiConfig{
		GeminiAPIKey: getEnv("GEMINI_API_KEY", ""),
	}
}

func getEnv(key, fallback string) string {

	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

var DB = initDatabaseConfig()

var JWTEnvs = initJWTConfig()

var GoogleClient = initGoogleClientConfig()

var GeminiAPI = initGeminiAPIConfig()