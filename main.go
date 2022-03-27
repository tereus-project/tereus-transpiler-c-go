package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/tereus-project/tereus-remixer-c-go/env"
	"github.com/tereus-project/tereus-remixer-c-go/modules"
	"github.com/tereus-project/tereus-remixer-c-go/remixer"
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
	rabbitmqService, err := modules.NewRabbitMQService()
	if err != nil {
		fmt.Println(err)
	}

	ch, err := rabbitmqService.ConsumeRemixJob()
	if err != nil {
		fmt.Println(err)
	}

	for job := range ch {
		fmt.Println(job)
	}
}
