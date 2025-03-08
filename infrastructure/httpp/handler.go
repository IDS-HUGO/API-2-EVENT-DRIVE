package httpp

import (
	"log"
	"net/http"

	"api2/application"
	"api2/domain"

	"github.com/gin-gonic/gin"
)

type MessageHandler struct {
	Service *application.MessageService
}

func NewMessageHandler(service *application.MessageService) *MessageHandler {
	return &MessageHandler{Service: service}
}

func (h *MessageHandler) HandleMessage(c *gin.Context) {
	var msg domain.Message

	// Usar directamente ShouldBindJSON sin leer el cuerpo manualmente
	if err := c.ShouldBindJSON(&msg); err != nil {
		log.Println("Error en el formato del mensaje:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido"})
		return
	}

	log.Printf("Cuerpo recibido en la API: ID=%d, Name=%s, Price=%.2f", msg.ID, msg.Name, msg.Price)

	if msg.Name == "" {
		log.Println("⚠️ Mensaje vacío recibido en la API")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Mensaje vacío"})
		return
	}

	h.Service.ProcessMessage(msg)
	c.JSON(http.StatusOK, gin.H{"status": "Mensaje procesado y guardado"})
}
