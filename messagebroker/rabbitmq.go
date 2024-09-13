package messagebroker

import (
	"api_gateway/genproto/auth"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	regQueue *amqp.Queue
}

func NewRabbitMQ(conn *amqp.Connection, ch *amqp.Channel) *RabbitMQ {
	return &RabbitMQ{
		conn:    conn,
		channel: ch,
	}
}

func (r *RabbitMQ) PublishRegister(message *auth.RegisterRequest) error {

	byteData, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		"",              // exchange
		r.regQueue.Name, // routing key (queue name)
		false,           // mandatory
		false,           // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(byteData),
		},
	)
}

func (r *RabbitMQ) Consume(queueName string, handler func(string)) error {
	msgs, err := r.channel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			handler(string(d.Body))
		}
	}()
	return nil
}

func DeclareQueue(name string, ch *amqp.Channel) (*amqp.Queue, error) {
	q, err := ch.QueueDeclare(
		name,  // queue name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return nil, err
	}

	return &q, nil
}
