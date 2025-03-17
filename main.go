package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "iot-project/internal/Accesos/application/services"
    "iot-project/internal/Accesos/infraestructure/handler"
    "iot-project/internal/Accesos/infraestructure/repository"
)






func main() {

    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    broker := os.Getenv("MQTT_BROKER")
    clientID := os.Getenv("MQTT_CLIENT_ID")
    topic := os.Getenv("MQTT_TOPIC")

    // Inicializa el handler MQTT
    mqttHandler := handler.NewMQTTHandler(broker, clientID)

    // Inicializa el repositorio MQTT
    mqttRepo := repository.NewMQTTRepository(mqttHandler)

    // Inicializa el repositorio de base de datos
    dbRepo := repository.NewDBRepository()

    // Inicializa el servicio de acceso
    accessService := services.NewAccessService(mqttRepo, dbRepo)

    // Suscríbete al tópico y comienza a escuchar
    accessService.StartListening(topic)

    // Mantén la aplicación en ejecución
    select {}
}