package common

import (
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
  prod := os.Getenv("PROD")

  if prod != "true" {
    err := godotenv.Load()

    if err != nil {
      return err
    }
  }

  return nil
}
