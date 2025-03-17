package repository

import (
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"iot-project/internal/Accesos/domain"
	"iot-project/internal/Accesos/infraestructure/handler"
	"iot-project/internal/Accesos/infraestructure/model"
	"log"
	"time"
)

type MQTTRepository struct {
	handler *handler.MQTTHandler
}

func NewMQTTRepository(handler *handler.MQTTHandler) *MQTTRepository {
	return &MQTTRepository{
		handler: handler,
	}
}

func (r *MQTTRepository) Listen(topic string, callback func(access domain.Access)) {
	r.handler.Subscribe(topic, func(client MQTT.Client, msg MQTT.Message) {
		var mqtt model.MQTTResult
		if err := json.Unmarshal(msg.Payload(), &mqtt); err != nil {
			log.Printf("Error decoding MQTT message: %s", err)
			return
		}

        

		if mqtt.Estado == "incorrecto" {
			log.Printf("No se puede guardar el mensaje: %s", mqtt.Pin)
		} else if mqtt.Estado == "correcto" {
			access := domain.Access{
				CodigoIngresado: mqtt.Pin,
				Estado:          mqtt.Estado,
				Fecha:           time.Now(),
			}

			callback(access)
		}

	})
}
