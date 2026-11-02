package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DBDriver    string
	DBHost      string
	DBPort      string
	DBUser      string
	DBPassword  string
	DBName      string
	DBSSLMode   string
	JWTSecret   string
	UploadDir   string
	ClientURL   string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	port := getEnv("PORT", "8088")
	dbDriver := getEnv("DB_DRIVER", "postgres")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("DB_USER", "postgres")
	dbPass := getEnv("DB_PASSWORD", "postgres")
	dbName := getEnv("DB_NAME", "alhamid_db")
	dbSSL := getEnv("DB_SSLMODE", "disable")
	jwtSecret := getEnv("JWT_SECRET", "alhamid-super-secret-key-2026")
	uploadDir := getEnv("UPLOAD_DIR", "./uploads")
	clientURL := getEnv("CLIENT_URL", "http://localhost:5173")

	return &Config{
		Port:       port,
		DBDriver:   dbDriver,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBUser:     dbUser,
		DBPassword: dbPass,
		DBName:     dbName,
		DBSSLMode:  dbSSL,
		JWTSecret:  jwtSecret,
		UploadDir:  uploadDir,
		ClientURL:  clientURL,
	}
}

func getEnv(key, defaultVal string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultVal
}
