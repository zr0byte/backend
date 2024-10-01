package config

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

func LoadConfigs() {
	_, b, _, _ := runtime.Caller(0)
	rootPath := filepath.Join(filepath.Dir(b), "../")

	err := godotenv.Load(rootPath + "/.env")
	if err != nil {
		log.Fatal("error loading .env file")
	}

	loadAppConfig()
}
