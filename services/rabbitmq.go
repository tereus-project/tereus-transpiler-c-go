package services

import (
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQService struct {
	conn *amqp.Connection

	queues []*RabbitMQQueue
}

type RabbitMQQueue struct {
	name       string
	exchange   string
	routingKey string

	channel *amqp.Channel
}

func NewRabbitMQService(endpoint string) (*RabbitMQService, error) {
	conn, err := amqp.Dial(endpoint)
	if err != nil {
		return nil, err
	}

	return &RabbitMQService{
		conn: conn,
	}, nil
}

func (s *RabbitMQService) Close() error {
	return s.conn.Close()
}

func (s *RabbitMQService) NewQueue(name string, exchange string, routingKey string) (*RabbitMQQueue, error) {
	channel, err := s.conn.Channel()
	if err != nil {
		return nil, err
	}

	// Ensure exchange exists
	err = channel.ExchangeDeclare(
		exchange, // name
		"direct", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	// Ensure queue exists
	queue, err := channel.QueueDeclare(
		name,  // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	// Bind queue to exchange
	err = channel.QueueBind(
		queue.Name, // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return &RabbitMQQueue{
		name:       queue.Name,
		exchange:   exchange,
		routingKey: routingKey,
		channel:    channel,
	}, nil
}

func (q *RabbitMQQueue) Close() error {
	return q.channel.Close()
}

func (q *RabbitMQQueue) Publish(data interface{}) error {
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return q.channel.Publish(
		q.exchange,   // exchange
		q.routingKey, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (q *RabbitMQQueue) Consume() (<-chan amqp.Delivery, error) {
	return q.channel.Consume(
		q.name, // queue name
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
}
