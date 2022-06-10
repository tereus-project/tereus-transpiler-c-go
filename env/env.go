package env

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Env struct {
	S3Bucket    string `env:"S3_BUCKET" env-required:"true"`
	S3AccessKey string `env:"S3_ACCESS_KEY" env-required:"true"`
	S3SecretKey string `env:"S3_SECRET_KEY" env-required:"true"`
	S3Endpoint  string `env:"S3_ENDPOINT" env-required:"true"`

	NSQEndpoint        string `env:"NSQ_ENDPOINT" env-required:"true"`
	NSQLookupdEndpoint string `env:"NSQ_LOOKUPD" env-required:"true"`

	LogFormat string `env:"LOG_FORMAT" env-default:"json"`
	LogLevel  string `env:"LOG_LEVEL" env-default:"info"`
	SentryDSN string `env:"SENTRY_DSN"`
	Env       string `env:"ENV" env-required:"true"`

	MetricsPort string `env:"METRICS_PORT" env-default:"8080"`
}

var env Env

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		logrus.WithError(err).Warn("Failed to load env variables from file")
	}

	return cleanenv.ReadEnv(&env)
}

func Get() *Env {
	return &env
}
