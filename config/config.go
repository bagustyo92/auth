package config

import (
	"os"

	"github.com/bagustyo92/auth/middleware/logger"
	"github.com/joho/godotenv"
)

var (
	AppPort    string
	DBUsername string
	DBPassword string
	DBName     string
	DBURL      string
	DBPort     string
)

func getEnv(keyEnv string, fileEnv map[string]string) string {
	envVal, ok := os.LookupEnv(keyEnv)
	if !ok {
		return fileEnv[keyEnv]
	}
	return envVal
}

// InitApp will get all the important env from env file
// Do not include env file if you want get the env from host
func InitApp(envPath string) {
	envFile, err := godotenv.Read(envPath)
	if err != nil {
		logger.MakeLogEntry(nil).Panic(err)
	}

	AppPort = getEnv("APP_PORT", envFile)

	// DB Env
	DBName = getEnv("DB_NAME", envFile)
	DBUsername = getEnv("DB_USER", envFile)
	DBPassword = getEnv("DB_PASS", envFile)
	DBPort = getEnv("DB_PORT", envFile)
	DBURL = getEnv("DB_HOST", envFile)
}
