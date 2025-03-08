package main

import (
	"api2/application"
	"api2/core"
	"api2/infrastructure/database"
	"api2/infrastructure/httpp"
	"api2/infrastructure/rabbitmq"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	core.LoadEnv()

	// Conectar a la base de datos
	database.ConnectDB()
	defer database.DB.Close()

	// Inicializar RabbitMQ
	producer := rabbitmq.NewRabbitMQProducer()

	// Inicializar servicio
	messageService := application.NewMessageService(producer)

	// Inicializar controlador
	messageHandler := httpp.NewMessageHandler(messageService)

	// Configurar servidor HTTP con Gin
	r := gin.Default()
	r.POST("/procesar-mensaje", messageHandler.HandleMessage)

	// Iniciar servidor HTTP
	port := core.GetEnv("API_PORT", "8000")
	log.Println("API REST corriendo en puerto", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
