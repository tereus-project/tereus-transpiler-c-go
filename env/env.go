package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	LogLevel         string
	S3Bucket         string
	S3AccessKey      string
	S3SecretKey      string
	S3Endpoint       string
	RabbitMQEndpoint string
	KafkaEndpoint    string
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	LogLevel = os.Getenv("LOG_LEVEL")
	if LogLevel == "" {
		LogLevel = "info"
	}

	S3Bucket = os.Getenv("S3_BUCKET")
	if S3Bucket == "" {
		return fmt.Errorf("S3_BUCKET is not set")
	}

	S3AccessKey = os.Getenv("S3_ACCESS_KEY")
	if S3AccessKey == "" {
		return fmt.Errorf("S3_ACCESS_KEY is not set")
	}

	S3SecretKey = os.Getenv("S3_SECRET_KEY")
	if S3SecretKey == "" {
		return fmt.Errorf("S3_SECRET_KEY is not set")
	}

	S3Endpoint = os.Getenv("S3_ENDPOINT")
	if S3Endpoint == "" {
		return fmt.Errorf("S3_ENDPOINT is not set")
	}

	// RabbitMQEndpoint = os.Getenv("RABBITMQ_ENDPOINT")
	// if RabbitMQEndpoint == "" {
	// 	return fmt.Errorf("RABBITMQ_ENDPOINT is not set")
	// }

	KafkaEndpoint = os.Getenv("KAFKA_ENDPOINT")
	if KafkaEndpoint == "" {
		return fmt.Errorf("KAFKA_ENDPOINT is not set")
	}

	return nil
}
