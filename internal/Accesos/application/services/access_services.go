
package services

import (
    "fmt"
    "iot-project/internal/Accesos/domain"
    "iot-project/internal/Accesos/infraestructure/repository"
    "iot-project/internal/websocket"
    "log"
)
type AccessService struct {
    mqttRepo *repository.MQTTRepository
    dbRepo   *repository.DBRepository
    wsHandler *websocket.WebSocketHandler
}

func NewAccessService(mqttRepo *repository.MQTTRepository, dbRepo *repository.DBRepository , wsHandler *websocket.WebSocketHandler)*AccessService {
    return &AccessService{
        mqttRepo: mqttRepo,
        dbRepo:   dbRepo,
        wsHandler:wsHandler,
    }
}

func (s *AccessService) StartListening(topic string) {
    s.mqttRepo.Listen(topic, func(access domain.Access) {
        if access.Estado == "incorrecto"{
            fmt.Printf("no se guardara")
            msg:= websocket.Message{
                Tipo:"Accesos incorrectos",
                Data:access,
            }
            s.wsHandler.Broadcast <- msg
        }else if  err := s.dbRepo.Save(access); err != nil {
            fmt.Printf("Error al guardar  en la database: %s\n", err)
        } else {
            fmt.Printf("Acceso  guardado: %+v\n", access)
            msg:= websocket.Message{
                Tipo:"Accesos correctos",
                Data:access,
            }

            log.Printf("Enviando mensaje al WebSocket: %+v", msg) 
            s.wsHandler.Broadcast <- msg
        }
    })
}