package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port uint16
}

func NewConfiguration() *Configuration {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("unable to load config")
	}

	port, _ := strconv.Atoi(os.Getenv("PORT"))
	return &Configuration{
		Port: uint16(port),
	}
}
