package rabbitmq

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQProducer struct {
	QueueName string
	Conn      *amqp.Connection
	Channel   *amqp.Channel
}

func NewRabbitMQProducer() RabbitMQProducer {
	rabbitURL := os.Getenv("RABBITMQ_URL")
	queueName := os.Getenv("RABBITMQ_QUEUE_OUT")

	conn, err := amqp.Dial(rabbitURL)
	if err != nil {
		log.Panic("Error conectando a RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Panic("Error abriendo canal:", err)
	}

	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Panic("Error declarando cola:", err)
	}

	log.Println("Conectado a RabbitMQ, cola:", q.Name)

	return RabbitMQProducer{
		QueueName: queueName,
		Conn:      conn,
		Channel:   ch,
	}
}

func (p *RabbitMQProducer) Publish(message string) {
	log.Printf("📤 Enviando mensaje a RabbitMQ (%s): %q (Bytes: %v)", p.QueueName, message, []byte(message))

	err := p.Channel.Publish(
		"",
		p.QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	if err != nil {
		log.Println("❌ Error enviando mensaje a RabbitMQ:"+p.QueueName, err)
	} else {
		log.Println("✅ Mensaje enviado a RabbitMQ:"+p.QueueName, message)
	}
}
