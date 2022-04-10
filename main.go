package main

import (
	"fmt"
	"os"
	"strings"

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
		log.Debugf("Job '%s' started", job.ID)

		type jobFiles struct {
			objectPath string
			localPath  string
		}

		var files []*jobFiles

		log.Debug("Downloading job files...")
		for objectPath := range minioService.GetFiles(job.ID) {
			log.Debugf("Downloading file '%s'", objectPath)
			localPath, err := minioService.GetFile(job.ID, objectPath)
			if err != nil {
				log.WithError(err).Fatal()
			}

			files = append(files, &jobFiles{
				objectPath: objectPath,
				localPath:  localPath,
			})
		}

		log.Debug("Remixing job files...")
		for _, file := range files {
			if !strings.HasSuffix(file.objectPath, ".c") {
				data, err := os.ReadFile(file.localPath)
				if err != nil {
					log.WithError(err).Fatal()
				}

				err = minioService.PutFile(job.ID, file.objectPath, string(data))
				if err != nil {
					log.WithError(err).Fatal()
				}

				continue
			}

			log.Debugf("Remixing file '%s'", file.objectPath)

			output, err := remixer.Remix(file.localPath)
			if err != nil {
				log.WithError(err).Fatal()
			}

			objectPath := fmt.Sprintf("%s.go", strings.TrimSuffix(file.objectPath, ".c"))

			log.Debugf("Uploading file '%s'", objectPath)
			err = minioService.PutFile(job.ID, objectPath, output)
			if err != nil {
				log.WithError(err).Fatal()
			}
		}

		log.Debugf("Job '%s' completed", job.ID)
	}
}
