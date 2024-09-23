package config

import (
	"github.com/joho/godotenv"
	"github.com/ihakalanka/Assessment-Golang-REST-API/utils"
)

func InitEnv() {
    if err := godotenv.Load(); err != nil {
        utils.Log.Warn("No .env file found, using environment variables")
    }
}