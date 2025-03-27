package services

import(
	"fmt"
	"iot-project/internal/Distancias/domain"
	"iot-project/internal/Distancias/infraestructure/repository"
	"iot-project/internal/websocket"
	"time"
)



type DistanceSensorServices struct{
	mqttRepo *repository.MqttDistanceRepository
	dbRepo *repository.DBRepository
	wsHandler *websocket.WebSocketHandler
	lastState *domain.Distance 
}

func NewDistanceSensorService(mqttRepo *repository.MqttDistanceRepository , dbRepo *repository.DBRepository , wsHandler *websocket.WebSocketHandler)*DistanceSensorServices{
	return &DistanceSensorServices{
		mqttRepo: mqttRepo,
		dbRepo:dbRepo,
		wsHandler:wsHandler,
		lastState : nil,
	}
}


func (s *DistanceSensorServices) shouldSave(distance domain.Distance) bool {
	
	if s.lastState == nil {
		return true
	}

	
	if distance.Proximidad != s.lastState.Proximidad {
		return true
	}

	
	if distance.Proximidad == "muy cerca" {
		elapsed := distance.Fecha.Sub(s.lastState.Fecha)
		return elapsed > 10*time.Minute
	}

	
	if distance.Proximidad == "muy cerca" && s.lastState.Proximidad != "muy cerca" {
		return true // Entrando a zona crítica
	}
	if distance.Proximidad != "muy cerca" && s.lastState.Proximidad == "muy cerca" {
		return true // Saliendo de zona crítica
	}

	
	return false
}


func (s *DistanceSensorServices) StartListening(topic string) {
	s.mqttRepo.Listen(topic, func(distance domain.Distance) {
		if distance.ValidDistance() {
			distance.IsProximity()

			
			if s.shouldSave(distance) {
				if err := s.dbRepo.Save(distance); err != nil {
					fmt.Printf("Error al guardar la distancia: %s\n", err)
				} else {
					fmt.Printf("Datos de distancia guardados: %+v\n", distance)
					s.lastState = &distance 
					
					msg := websocket.Message{
						Tipo: "distancia",
						Data: distance,
					}
					s.wsHandler.Broadcast <- msg
				}
			} else {
				return
			}
		} else {
			fmt.Printf("Distancia no válida: %+v\n", distance)
		}
	})
}



