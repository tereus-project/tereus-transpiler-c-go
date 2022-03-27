package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tereus-project/tereus-remixer-c-go/env"
	"github.com/tereus-project/tereus-remixer-c-go/remixer"
	"github.com/tereus-project/tereus-remixer-c-go/services"
)

func main() {
	if len(os.Args) >= 2 {
		path := os.Args[1]

		output, err := remixer.Remix(path)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to remix '%s'", path)
		}

		fmt.Println(output)
		return
	}

	err := env.LoadEnv()
	if err != nil {
		logrus.WithError(err).Fatal("Failed to connect to database")
	}

	initWorker()
}

func initWorker() {
	rabbitmqService, err := services.NewRabbitMQService()
	if err != nil {
		fmt.Println(err)
	}

	minioService, err := services.NewMinioService()
	if err != nil {
		fmt.Println(err)
	}

	deliveries, err := rabbitmqService.ConsumeRemixJob()
	if err != nil {
		fmt.Println(err)
	}

	for job := range deliveries {
		for filepath := range minioService.GetFiles(job.ID) {
			path, err := minioService.GetFile(job.ID, filepath)
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(path)
		}
	}
}
