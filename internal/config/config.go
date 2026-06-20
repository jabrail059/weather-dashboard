package config

import (
	"os"
	"strconv"
)

type Config struct {
	AppAddr       string
	SqlitePath    string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
}

func New() *Config {

	return &Config{
		AppAddr:       getEnv("APP_ADDR", ":8081"),
		SqlitePath:    getEnv("SQLITE_PATH", ""),
		RedisAddr:     getEnv("REDIS_ADDR", ""),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       getEnvAsInt("REDIS_DB", 0),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getEnvAsInt(key string, defaultVal int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}
