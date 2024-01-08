package configs

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var (
	port   int
	env    string
	appUrl string
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
		env = os.Getenv("APP_ENV")
	}
	return env
}

func loadEnv() {
	err := godotenv.Load("./.env")
	if err != nil {
		panic(err)
	}
}

func GetConfigResize() (int, int) {
	return 800, 600
}
