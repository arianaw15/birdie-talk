package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string

	DBUser        string
	DBPassword    string
	DBName        string
	DBAddress     string
	JWTExpiration int
	JWTSecret     string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost:    getEnv("PUBLIC_HOST", "http://localhost"),
		Port:          getEnv("PORT", "8080"),
		DBUser:        getEnv("DB_USER", "root"),
		DBPassword:    getEnv("DB_PASSWORD", "Coggy254!"),
		DBAddress:     fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:        getEnv("DB_NAME", "birdie_talk"),
		JWTExpiration: getEnvInt("JWT_EXPIRATION", 3600*24*3),
		JWTSecret:     getEnv("JWT_SECRET", "secret-information"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	if value, ok := os.LookupEnv(key); ok {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}

	return fallback
}
