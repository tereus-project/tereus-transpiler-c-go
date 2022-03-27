package services

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/tereus-project/tereus-remixer-c-go/env"
)

type RabbitMQService struct {
	connection *amqp.Connection
	channel    *amqp.Channel

	consumersTags []string
}

func NewRabbitMQService() (*RabbitMQService, error) {
	var err error
	service := &RabbitMQService{}

	service.connection, err = amqp.Dial(env.RabbitMQEndpoint)
	if err != nil {
		return nil, err
	}

	service.channel, err = service.connection.Channel()
	if err != nil {
		return nil, err
	}

	// Ensure exchange exists
	err = service.channel.ExchangeDeclare(
		"remix_jobs_ex", // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (s *RabbitMQService) Close() error {
	for _, tag := range s.consumersTags {
		if err := s.channel.Cancel(tag, true); err != nil {
			return fmt.Errorf("Consumer cancel failed: %s", err)
		}
	}

	if err := s.channel.Close(); err != nil {
		return fmt.Errorf("Channel close failed: %s", err)
	}

	if err := s.connection.Close(); err != nil {
		return fmt.Errorf("AMQP connection close error: %s", err)
	}

	return nil
}

type remixJob struct {
	ID             string `json:"id"`
	SourceLanguage string `json:"source_language"`
	TargetLanguage string `json:"target_language"`
}

// Publish a job to the exchange
func (s *RabbitMQService) ConsumeRemixJob() (<-chan remixJob, error) {
	// Ensure queue exists
	queue, err := s.channel.QueueDeclare(
		"remix_jobs_q", // name
		true,           // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		return nil, err
	}

	// Bind queue to exchange
	err = s.channel.QueueBind(
		queue.Name,      // queue name
		"remix_jobs_rk", // routing key
		"remix_jobs_ex", // exchange
		false,           // noWait
		nil,             // arguments
	)
	if err != nil {
		return nil, err
	}

	tag := uuid.New().String()

	deliveries, err := s.channel.Consume(
		queue.Name, // queue name
		tag,        // consumer
		false,      // autoAck
		false,      // exclusive
		false,      // noLocal
		false,      // noWait
		nil,        // arguments
	)
	if err != nil {
		return nil, err
	}

	ch := make(chan remixJob)
	go func() {
		for d := range deliveries {
			var job remixJob

			if err := json.Unmarshal(d.Body, &job); err != nil {
				fmt.Fprintf(os.Stderr, "Error unmarshalling job %s: %s\n", tag, err)
				d.Nack(false, false)
				continue
			}

			if job.SourceLanguage != "c" && job.TargetLanguage != "go" {
				d.Reject(true)
				continue
			}

			d.Ack(false)
			ch <- job
		}

		close(ch)
	}()

	return ch, nil
}
