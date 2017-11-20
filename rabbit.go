package hanabi

import (
	"strings"

	"github.com/streadway/amqp"
)

type RabbitTransporter struct {
	rabbit  *amqp.Connection
	channel *amqp.Channel
}

func newRabbit(URL string) (*RabbitTransporter, error) {
	t := &RabbitTransporter{}

	c, err := amqp.Dial(rabbitURL(URL))
	if err != nil {
		return nil, err
	}
	t.rabbit = c

	ch, err := c.Channel()
	if err != nil {
		return nil, err
	}
	t.channel = ch

	err = ch.ExchangeDeclare(
		"logs",   // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"events", // name
		"fanout", // type
		true,     // durable
		false,    // auto-deleted
		false,    // internal
		false,    // no-wait
		nil,      // arguments
	)
	if err != nil {
		return nil, err
	}

	err = ch.ExchangeDeclare(
		"response", // name
		"direct",   // type
		true,       // durable
		false,      // auto-deleted
		false,      // internal
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func rabbitURL(URL string) string {
	s := URL
	if !strings.HasPrefix(s, "amqp") {
		s = "amqp://" + s
	}
	return s
}

func (t *RabbitTransporter) send(key string, value string) error {
	return nil
}

func (t *RabbitTransporter) log(level int, message string) error {
	return nil
}
