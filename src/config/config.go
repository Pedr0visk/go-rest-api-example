package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var SecretKey []byte

func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
