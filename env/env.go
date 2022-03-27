package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	RabbitMQEndpoint string
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	RabbitMQEndpoint = os.Getenv("RABBITMQ_ENDPOINT")
	if RabbitMQEndpoint == "" {
		return fmt.Errorf("RABBITMQ_ENDPOINT is not set")
	}

	return nil
}
