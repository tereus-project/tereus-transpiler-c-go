package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/tereus-project/tereus-go-std/logging"
	"github.com/tereus-project/tereus-remixer-c-go/env"
	"github.com/tereus-project/tereus-remixer-c-go/remixer"
	"github.com/tereus-project/tereus-remixer-c-go/services"
)

var (
	prom_namespace       = "remixer_c_go"
	remixingDurationHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: prom_namespace,
		Name:      "remixing_duration_seconds",
		Help:      "Histogram of remixing duration",
	}, []string{"status"})
)

func main() {
	err := env.LoadEnv()
	if err != nil {
		log.WithError(err).Fatal("Failed to load environment variables")
	}

	config := env.Get()

	sentryHook, err := logging.SetupLog(logging.LogConfig{
		Format:       config.LogFormat,
		LogLevel:     config.LogLevel,
		ShowFilename: true,
		ReportCaller: true,
		SentryDSN:    config.SentryDSN,
	})
	if err != nil {
		logrus.WithError(err).Fatal("Failed to set log configuration")
	}
	defer sentryHook.Flush()
	defer logging.RecoverAndLogPanic()

	if len(os.Args) >= 2 {
		path := os.Args[1]

		output, err := remixer.Remix(path)
		if err != nil {
			log.WithError(err).Fatalf("Failed to remix '%s'", path)
		}

		fmt.Println(output)
		return
	}

	go exposeMetrics()

	initWorker(config)
}

func exposeMetrics() {
	config := env.Get()

	prometheus.MustRegister(remixingDurationHist)

	http.Handle("/metrics", promhttp.Handler())
	listenAddr := fmt.Sprintf("0.0.0.0:%s", config.MetricsPort)
	log.Infof("Beginning to serve metrics on %s", listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}

func initWorker(config *env.Env) {
	log.Info("Connecting to Kafka...")
	kafkaService, err := services.NewKafkaService(config.KafkaEndpoint)
	if err != nil {
		log.WithError(err).Fatal()
	}
	defer kafkaService.CloseAllWriters()

	log.Info("Connecting to MinIO...")
	minioService, err := services.NewMinioService(config.S3Endpoint, config.S3AccessKey, config.S3SecretKey, config.S3Bucket)
	if err != nil {
		log.WithError(err).Fatal()
	}

	log.Info("Starting remix job listener...")
	startRemixJobListener(kafkaService, minioService)
	log.Warn("Listener has stopped")
}

func startRemixJobListener(k *services.KafkaService, minioService *services.MinioService) {
	submissions := k.ConsumeSubmissions(context.Background())

	for job := range submissions {
		startTime := time.Now()

		err := k.PublishSubmissionStatus(services.SubmissionStatusMessage{
			ID:     job.ID,
			Status: services.StatusProcessing,
		})
		if err != nil {
			log.WithError(err).WithField("job_id", job.ID).Errorf("Error publishing status message for job")
		}

		log.Debugf("Job '%s' started", job.ID)

		err = remix(job.ID, minioService)
		if err != nil {
			log.WithError(err).WithField("job_id", job.ID).Errorf("Failed to remix and upload job")
			remixingDurationHist.WithLabelValues(string(services.StatusFailed)).Observe(time.Since(startTime).Seconds())

			err := k.PublishSubmissionStatus(services.SubmissionStatusMessage{
				ID:     job.ID,
				Status: services.StatusFailed,
				Reason: err.Error(),
			})
			if err != nil {
				log.WithError(err).WithField("job_id", job.ID).Errorf("Error publishing status message for job")
			}
		} else {
			remixingDurationHist.WithLabelValues(string(services.StatusDone)).Observe(time.Since(startTime).Seconds())
			err := k.PublishSubmissionStatus(services.SubmissionStatusMessage{
				ID:     job.ID,
				Status: services.StatusDone,
			})
			if err != nil {
				log.WithError(err).WithField("job_id", job.ID).Errorf("Error publishing status message for job")
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
