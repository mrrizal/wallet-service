package configs

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DBURI string
	Port  string
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var config *Config
var once sync.Once

func GetConfig() *Config {
	godotenv.Load(".env")
	pgUser := getEnv("POSTGRES_USER", "")
	pgPassword := getEnv("POSTGRES_PASSWORD", "")
	pgDB := getEnv("POSTGRES_DB", "")
	pgHost := getEnv("POSTGRES_HOST", "")
	pgPort := getEnv("POSTGRES_PORT", "")

	dbURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", pgUser, pgPassword, pgHost, pgPort, pgDB)
	once.Do(func() {
		config = &Config{
			DBURI: dbURI,
			Port:  getEnv("PORT", "3000"),
		}
	})
	return config
}
