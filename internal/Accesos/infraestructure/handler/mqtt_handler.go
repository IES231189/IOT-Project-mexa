// internal/Accesos/infraestructure/handler/mqtt_handler.go
package handler

import (
    "fmt"
    "log"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "os"
)

type MQTTHandler struct {
    client MQTT.Client
}

func NewMQTTHandler(broker string, clientID string) *MQTTHandler {
    opts := MQTT.NewClientOptions().AddBroker(broker)
    opts.SetClientID(clientID)

    // Agrega las credenciales de MQTT (usuario y contraseña)
    mqttUsername := os.Getenv("MQTT_USERNAME")
    mqttPassword := os.Getenv("MQTT_PASSWORD")
    opts.SetUsername(mqttUsername)
    opts.SetPassword(mqttPassword)

    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        log.Fatalf("Error al conectarce a MQTT broker: %s", token.Error())
    }

    return &MQTTHandler{
        client: client,
    }
}

func (h *MQTTHandler) Subscribe(topic string, callback MQTT.MessageHandler) {
    if token := h.client.Subscribe(topic, 0, callback); token.Wait() && token.Error() != nil {
        log.Fatalf("Error al suscribirce al  topic %s: %s", topic, token.Error())
    }
    fmt.Printf("Subscrito al topico: %s\n", topic)
}