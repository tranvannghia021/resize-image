package configs

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var (
	port     int
	env      string
	appUrl   string
	mongoDns string
	mongoDB  string
)

func GetPort() int {
	if port == 0 {
		loadEnv()
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	return port
}

func GetAppUrl() string {
	if appUrl == "" {
		loadEnv()
		appUrl = os.Getenv("APP_URL")
	}
	return appUrl
}

func GetAppEnv() string {
	if env == "" {
		loadEnv()
		env = os.Getenv("APP_ENV")
	}
	return env
}

func GetMongoDns() string {
	if mongoDns == "" {
		loadEnv()
		mongoDns = os.Getenv("MONGO_DNS")
	}
	return mongoDns
}

func loadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
}

func GetMongoDB() string {
	if mongoDB == "" {
		loadEnv()
		mongoDB = os.Getenv("MONGO_DB")
	}
	return mongoDB
}
