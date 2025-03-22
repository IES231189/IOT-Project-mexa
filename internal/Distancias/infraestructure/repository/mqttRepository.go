package repository

import (
	"log"
	"encoding/json"
	"iot-project/internal/Distancias/domain"
	"iot-project/internal/MqttConnection"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"time"
)

type MqttDistanceRepository struct {
	mqttServices *mqtt.MQTTService
}

func NewMqttDistanceRepository(mqttService *mqtt.MQTTService) *MqttDistanceRepository {
	return &MqttDistanceRepository{
		mqttServices: mqttService,
	}
}

// Listen se suscribe al tópico MQTT y recibe la distancia
func (r *MqttDistanceRepository) Listen(topic string, callback func(distance domain.Distance)) {
	r.mqttServices.Subscribe(topic, func(client MQTT.Client, msg MQTT.Message) {
		var distancia float64

		// Intentamos decodificar el mensaje
		if err := json.Unmarshal(msg.Payload(), &distancia); err != nil {
			log.Printf("Error al decodificar el mensaje: %s", err)
			return
		}


		distance := domain.Distance{
			Distancia: distancia,
			Fecha:     time.Now(),  // Usamos time.Now() directamente
		}

		
		if !distance.ValidDistance() {
			log.Println("Distancia no válida, no se procesará el mensaje.")
			return
		}

		
		distance.IsProximity()

		// Llamar al callback para procesar la distancia
		callback(distance)
	})
}
