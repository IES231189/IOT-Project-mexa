package main

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "iot-project/internal/Accesos/application/services"
    "iot-project/internal/Accesos/infraestructure/repository"
    MqttConnection "iot-project/internal/MqttConnection"

    //Distance in sensor ultrasonico 
    DistanceServices"iot-project/internal/Distancias/application/services"
    DistanceRepo "iot-project/internal/Distancias/infraestructure/repository"
    //DistanceMQTT "iot-project/internal/MqttConnection"

)

func main() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    broker := os.Getenv("MQTT_BROKER")
    clientID := os.Getenv("MQTT_CLIENT_ID")
    topicAcceso := os.Getenv("MQTT_TOPIC1") 
    topicDistancia := os.Getenv("MQTT_TOPIC2") // Asegúrate de definir esta variable en tu .env

    // Conexión con MQTT
    mqttService := MqttConnection.NewMQTTService(broker, clientID)

    // Accesos
    dbRepo := repository.NewDBRepository()
    accessRepo := repository.NewMQTTRepository(mqttService)
    accessService := services.NewAccessService(accessRepo, dbRepo)

    // Distancia
    distanceMqttRepo := DistanceRepo.NewMqttDistanceRepository(mqttService)
    distanceDbRepo := DistanceRepo.NewDBRepository()
    distanceService := DistanceServices.NewDistanceSensorService(distanceMqttRepo, distanceDbRepo)

    // Iniciar escucha de tópicos
    accessService.StartListening(topicAcceso)
    distanceService.StartListening(topicDistancia)

    select {}
}