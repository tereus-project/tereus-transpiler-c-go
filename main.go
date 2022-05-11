package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
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
	kafkaService, err := services.NewKafkaService(env.KafkaEndpoint)
	if err != nil {
		log.WithError(err).Fatal()
	}

	log.Info("Connecting to MinIO")
	minioService, err := services.NewMinioService()
	if err != nil {
		log.WithError(err).Fatal()
	}

	startRemixJobListener(kafkaService, minioService)
}

type remixJob struct {
	ID string `json:"id"`
}

func startRemixJobListener(k *services.KafkaService, minioService *services.MinioService) {
	for {
		msg, err := k.RemixSubmissionsConsumer.ReadMessage(-1)
		if err != nil {
			logrus.WithError(err).Error("Failed to read message")
			continue
		}
		var job remixJob

		if err := json.Unmarshal(msg.Value, &job); err != nil {
			log.WithError(err).Errorf("Error unmarshalling job %s: %s\n", err)
			continue
		}

		err = k.PublishSubmissionStatus(services.SubmissionStatusMessage{
			ID:     job.ID,
			Status: services.StatusProcessing,
			Reason: err.Error(),
		})
		if err != nil {
			log.WithError(err).Errorf("Error publishing status message for job '%s'", job.ID)
		}

		log.Debugf("Job '%s' started", job.ID)

		err = remix(job.ID, minioService)
		if err != nil {
			log.WithError(err).Errorf("Failed to remix and upload job '%s'", job.ID)

			err := k.PublishSubmissionStatus(services.SubmissionStatusMessage{
				ID:     job.ID,
				Status: services.StatusFailed,
				Reason: err.Error(),
			})
			if err != nil {
				log.WithError(err).Errorf("Error publishing status message for job '%s'", job.ID)
			}
		} else {
			err := k.PublishSubmissionStatus(services.SubmissionStatusMessage{
				ID:     job.ID,
				Status: services.StatusDone,
			})
			if err != nil {
				log.WithError(err).Errorf("Error publishing status message for job '%s'", job.ID)
			}

			log.Debugf("Job '%s' completed", job.ID)
		}
	}
}

type jobFile struct {
	objectPath string
	localPath  string
}

func remix(jobId string, minioService *services.MinioService) error {
	files, err := downloadObjectsToTempFiles(minioService, jobId)
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.HasSuffix(file.objectPath, ".c") {
			data, err := os.ReadFile(file.localPath)
			if err != nil {
				return err
			}

			err = minioService.PutFile(jobId, file.objectPath, string(data))
			if err != nil {
				return err
			}

			continue
		}

		log.Debugf("Remixing file '%s'", file.objectPath)

		output, err := remixer.Remix(file.localPath)
		if err != nil {
			return err
		}

		objectPath := fmt.Sprintf("%s.go", strings.TrimSuffix(file.objectPath, ".c"))

		log.Debugf("Uploading file '%s'", objectPath)
		err = minioService.PutFile(jobId, objectPath, output)
		if err != nil {
			return err
		}
	}

	return nil
}

func downloadObjectsToTempFiles(minioService *services.MinioService, jobId string) ([]*jobFile, error) {
	var files []*jobFile
	log.Debug("Downloading job files...")

	for objectPath := range minioService.GetFiles(jobId) {
		log.Debugf("Downloading file '%s'", objectPath)

		localPath, err := minioService.GetFile(jobId, objectPath)
		if err != nil {
			return nil, err
		}

		files = append(files, &jobFile{
			objectPath: objectPath,
			localPath:  localPath,
		})
	}

	return files, nil
}
