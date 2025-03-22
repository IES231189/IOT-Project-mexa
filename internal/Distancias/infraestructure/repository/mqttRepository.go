package repository

import (
	"log"
	"time"
	"encoding/json"
	"iot-project/internal/Distancias/domain"
	"iot-project/internal/MqttConnection"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MqttDistanceRepository struct {
	mqttServices *mqtt.MQTTService
}

func NewMqttDistanceRepository(mqttService *mqtt.MQTTService) *MqttDistanceRepository {
	return &MqttDistanceRepository{
		mqttServices: mqttService,
	}
}

func (r *MqttDistanceRepository) Listen(topic string, callback func(distance domain.Distance)) {
	r.mqttServices.Subscribe(topic, func(client MQTT.Client, msg MQTT.Message) {
		var distancia float64
		
		if err := json.Unmarshal(msg.Payload(), &distancia); err != nil {
			log.Printf("Error al decodificar el mensaje: %s", err)
			return
		}

		
		distance := domain.Distance{
			Distancia: distancia,
			Fecha:     time.Now().Unix(), 
		}

		
		distance.IsProximity()

		
		callback(distance)
	})
}
