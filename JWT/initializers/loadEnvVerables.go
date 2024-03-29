package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVerables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading env")
	}
}
