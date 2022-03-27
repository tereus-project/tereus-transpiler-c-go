package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/tereus-project/tereus-remixer-c-go/env"
	"github.com/tereus-project/tereus-remixer-c-go/remixer"
	"github.com/tereus-project/tereus-remixer-c-go/services"
)

func main() {
	err := env.LoadEnv()
	if err != nil {
		log.WithError(err).Fatal("Failed to load environment variables")
	}

	level, err := log.ParseLevel(env.LogLevel)
	if err != nil {
		log.Warnf("Invalid log level: '%s'", env.LogLevel)
	} else {
		log.SetLevel(level)
	}

	if len(os.Args) >= 2 {
		path := os.Args[1]

		output, err := remixer.Remix(path)
		if err != nil {
			log.WithError(err).Fatalf("Failed to remix '%s'", path)
		}

		fmt.Println(output)
		return
	}

	initWorker()
}

func initWorker() {
	log.Info("Connecting to RabbitMQ")
	rabbitmqService, err := services.NewRabbitMQService()
	if err != nil {
		log.WithError(err).Fatal()
	}

	log.Info("Connecting to MinIO")
	minioService, err := services.NewMinioService()
	if err != nil {
		log.WithError(err).Fatal()
	}

	log.Info("Initialization done. Waiting for jobs...")

	deliveries, err := rabbitmqService.ConsumeRemixJob()
	if err != nil {
		log.WithError(err).Fatal()
	}

	for job := range deliveries {
		for filepath := range minioService.GetFiles(job.ID) {
			path, err := minioService.GetFile(job.ID, filepath)
			if err != nil {
				log.WithError(err).Fatal()
			}

			log.Debug(path)
		}
	}
}
