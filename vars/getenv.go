package vars

import (
	"os"

	"github.com/joho/godotenv"
)

func Getenv(key string) string {

	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	return os.Getenv(key)
}