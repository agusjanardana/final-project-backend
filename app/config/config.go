package config

import (
	"github.com/joho/godotenv"
	"os"
	"vaccine-app-be/exceptions"
)

type Config interface {
	Get(key string) string
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exceptions.PanicIfError(err)
	return &config{}
}

type config struct {
}

func (c *config) Get(key string) string {
	return os.Getenv(key)
}