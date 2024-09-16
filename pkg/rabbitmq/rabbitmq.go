package rabbitmq

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectRabbitMQ(url string) (*amqp.Connection, error) {

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return conn, nil
}

func CreateChannel(conn *amqp.Connection) (*amqp.Channel, error) {
	ch, err := conn.Channel()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return ch, nil
}
