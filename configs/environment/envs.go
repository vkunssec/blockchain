package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(".env file not exists")
	}
	return os.Getenv(key)
}
