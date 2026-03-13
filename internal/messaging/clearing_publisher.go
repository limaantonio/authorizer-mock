package messaging

import (
	"encoding/json"

	amqp091 "github.com/rabbitmq/amqp091-go"
)

type ClearingPublisher struct {
	conn *amqp091.Connection
	exchangeName  *amqp091.Channel
	queue amqp091.Queue
}

func NewClearingPublisher(url string) (*ClearingPublisher, error) {
	conn, err := amqp091.Dial(url)

	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()

	if err != nil {
		return nil, err
	}

	queue, err := ch.QueueDeclare(
		"clearing_queue",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	return &ClearingPublisher{
		conn: conn,
		exchangeName: ch,
		queue: queue,
	}, nil
}

func (p *ClearingPublisher) Publish(event interface{}) error {

	body, err := json.Marshal(event)

	if err != nil {
		return err
	}

	err = p.exchangeName.Publish(
		"",
		p.queue.Name,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	return err
}