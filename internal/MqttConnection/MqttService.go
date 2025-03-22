package mqtt

import (
    "fmt"
    "log"
    "os"

    MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MQTTService struct {
    client MQTT.Client
}

func NewMQTTService(broker string, clientID string) *MQTTService {
    opts := MQTT.NewClientOptions().AddBroker(broker)
    opts.SetClientID(clientID)

    // Configura usuario y contraseña si es necesario
    mqttUsername := os.Getenv("MQTT_USERNAME")
    mqttPassword := os.Getenv("MQTT_PASSWORD")
    opts.SetUsername(mqttUsername)
    opts.SetPassword(mqttPassword)

    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("Error al conectar al broker MQTT: %s", token.Error())
    }

    return &MQTTService{
        client: client,
    }
}

func (s *MQTTService) Subscribe(topic string, callback MQTT.MessageHandler) {
    if token := s.client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
        log.Fatalf("Error al suscribirse al tópico %s: %s", topic, token.Error())
    }
    fmt.Printf("Suscrito al tópico: %s\n", topic)
}
