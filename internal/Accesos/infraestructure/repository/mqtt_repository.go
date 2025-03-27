package repository

import (
    "encoding/json"
    "log"
    "iot-project/internal/Accesos/domain"
    "iot-project/internal/MqttConnection"
    "iot-project/internal/Accesos/infraestructure/model"
    "time"
    MQTT "github.com/eclipse/paho.mqtt.golang" 
)

type MQTTRepository struct {
    mqttService *mqtt.MQTTService
}

func NewMQTTRepository(mqttService *mqtt.MQTTService) *MQTTRepository {
    return &MQTTRepository{
        mqttService: mqttService,
    }
}

func (r *MQTTRepository) Listen(topic string, callback func(access domain.Access)) {
    r.mqttService.Subscribe(topic, func(client MQTT.Client, msg MQTT.Message) {
        var mqtt model.MQTTResult
        if err := json.Unmarshal(msg.Payload(), &mqtt); err != nil {
            log.Printf("Error decoding MQTT message: %s", err)
            return
        }
        
        if mqtt.Estado == "incorrecto" {
            log.Printf("No se puede guardar el mensaje: %s", mqtt.Pin)
            access := domain.Access{
                CodigoIngresado: mqtt.Pin,
                Estado:          mqtt.Estado,
                Fecha:           time.Now(),
            }

            callback(access)
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