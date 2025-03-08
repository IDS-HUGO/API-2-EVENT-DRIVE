package application

import (
	"api2/domain"
	"api2/infrastructure/database"
	"api2/infrastructure/rabbitmq"
	"log"
)

type MessageService struct {
	Producer rabbitmq.RabbitMQProducer
}

func NewMessageService(producer rabbitmq.RabbitMQProducer) *MessageService {
	return &MessageService{Producer: producer}
}

func (s *MessageService) ProcessMessage(msg domain.Message) {
	log.Printf("Procesando mensaje: ID=%d, Name=%s, Price=%.2f", msg.ID, msg.Name, msg.Price)

	// Guardar en MySQL
	_, err := database.DB.Exec("INSERT INTO mensajes (id, name, price) VALUES (?, ?, ?)", msg.ID, msg.Name, msg.Price)
	if err != nil {
		log.Println("Error guardando mensaje en MySQL:", err)
		return
	}
	log.Println("Mensaje guardado en la base de datos")

	// Enviar a RabbitMQ
	s.Producer.Publish(msg.Name) // Enviar solo el 'Name' a RabbitMQ
}
