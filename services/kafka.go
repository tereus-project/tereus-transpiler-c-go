package services

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type KafkaService struct {
	RemixSubmissionsConsumer *kafka.Consumer
	Producer                 *kafka.Producer
}

func NewKafkaService(endpoint string) (*KafkaService, error) {
	remixSubmissionsConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": endpoint,
		"group.id":          "remixer-c-to-go",
		"auto.offset.reset": "latest",
	})
	if err != nil {
		return nil, err
	}
	err = remixSubmissionsConsumer.SubscribeTopics([]string{"remix_jobs_c_to_go"}, nil)
	if err != nil {
		return nil, err
	}

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": endpoint,
	})
	if err != nil {
		return nil, err
	}

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &KafkaService{
		RemixSubmissionsConsumer: remixSubmissionsConsumer,
		Producer:                 producer,
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

	topic := "submission_status"
	err = k.Producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          msgBytes,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
