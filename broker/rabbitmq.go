package broker

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(user, password, host, port string) (*amqp.Channel, func() error) {
	address := fmt.Sprintf("amqp://%s:%s@%s:%s", user, password, host, port)

	conn, err := amqp.Dial(address)
	if err != nil {
		log.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare(MangaAnimeSamaInitAll, "direct", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.ExchangeDeclare(MangaAnimeSamaUpdate, "fanout", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	return ch, conn.Close

}
