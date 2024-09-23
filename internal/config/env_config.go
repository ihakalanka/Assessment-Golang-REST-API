package config

import (
	"github.com/joho/godotenv"
	"github.com/ihakalanka/Assessment-Golang-REST-API/pkg"
)

func InitEnv() {
    if err := godotenv.Load(); err != nil {
        pkg.Log.Warn("No .env file found, using environment variables")
    }
}