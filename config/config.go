package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config estrutura as configurações globais da aplicação
type Config struct {
	ServerPort         string
	DefaultURL         string
	DefaultRequests    int
	DefaultConcurrency int
	HTTPTimeout        int
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando valores padrão")
	}

	return &Config{
		ServerPort:         getEnv("SERVER_PORT", "3000"),
		DefaultURL:         getEnv("DEFAULT_URL", "http://example.com"),
		DefaultRequests:    getEnvInt("DEFAULT_REQUESTS", 100),
		DefaultConcurrency: getEnvInt("DEFAULT_CONCURRENCY", 10),
		HTTPTimeout:        getEnvInt("HTTP_TIMEOUT", 10),
	}
}

// Funções auxiliares para buscar variáveis de ambiente
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		intVal, err := strconv.Atoi(value)
		if err == nil {
			return intVal
		}
	}
	return defaultValue
}
