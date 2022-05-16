package services

import (
	"context"
	"encoding/json"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type KafkaService struct {
	endpoint        string
	consumerGroupID string
	writers         map[string]*kafka.Writer
}

func NewKafkaService(endpoint string) (*KafkaService, error) {
	return &KafkaService{
		endpoint:        endpoint,
		consumerGroupID: "remixer-c-to-go",
		writers:         make(map[string]*kafka.Writer),
	}, nil
}

type SubmissionStatus string

const (
	StatusPending    SubmissionStatus = "pending"
	StatusProcessing SubmissionStatus = "processing"
	StatusDone       SubmissionStatus = "done"
	StatusFailed     SubmissionStatus = "failed"
)

type SubmissionStatusMessage struct {
	ID     string           `json:"id"`
	Status SubmissionStatus `json:"status"`
	Reason string           `json:"reason"`
}

func (k *KafkaService) PublishSubmissionStatus(status SubmissionStatusMessage) error {
	msgBytes, err := json.Marshal(&status)
	if err != nil {
		return err
	}

	writer, err := k.getWriterForTopic("submission_status")
	if err != nil {
		return err
	}

	return writer.WriteMessages(context.Background(),
		kafka.Message{
			Key:   []byte(status.ID),
			Value: msgBytes,
		},
	)
}

type SubmissionMessage struct {
	ID string `json:"id"`
}

func (k *KafkaService) ConsumeSubmissions(ctx context.Context) <-chan SubmissionMessage {
	ch := make(chan SubmissionMessage)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{k.endpoint},
		GroupID:     k.consumerGroupID,
		Topic:       "remix_jobs_c_to_go",
		MaxWait:     1 * time.Second,
		StartOffset: kafka.LastOffset,
	})

	go func() {
		defer r.Close()
		defer close(ch)

		for {
			msg, err := r.ReadMessage(ctx)
			if err != nil {
				logrus.WithError(err).Error("Error reading message")
			}

			var job SubmissionMessage
			err = json.Unmarshal(msg.Value, &job)
			if err != nil {
				logrus.WithError(err).Error("Failed to unmarshal submission status message")
				continue
			}

			ch <- job
		}
	}()

	return ch
}

func (k *KafkaService) CloseAllWriters() {
	for _, w := range k.writers {
		err := w.Close()
		if err != nil {
			logrus.WithError(err).Error("Failed to close kafka writer")
		}
	}
}

func (k *KafkaService) getWriterForTopic(topicName string) (*kafka.Writer, error) {
	w, ok := k.writers[topicName]
	if !ok {
		w := &kafka.Writer{
			Addr:                   kafka.TCP(k.endpoint),
			Topic:                  topicName,
			Balancer:               &kafka.LeastBytes{},
			AllowAutoTopicCreation: true,
		}
		k.writers[topicName] = w

		return w, nil
	}

	return w, nil
}
