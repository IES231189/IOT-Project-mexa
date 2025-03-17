// application/access_service.go
package services

import (
    "fmt"
    "iot-project/internal/Accesos/domain"
    "iot-project/internal/Accesos/infraestructure/repository"

)
type AccessService struct {
    mqttRepo *repository.MQTTRepository
    dbRepo   *repository.DBRepository
}

func NewAccessService(mqttRepo *repository.MQTTRepository, dbRepo *repository.DBRepository) *AccessService {
    return &AccessService{
        mqttRepo: mqttRepo,
        dbRepo:   dbRepo,
    }
}

func (s *AccessService) StartListening(topic string) {
    s.mqttRepo.Listen(topic, func(access domain.Access) {
       
        if err := s.dbRepo.Save(access); err != nil {
            fmt.Printf("Error al guardar  en la database: %s\n", err)
        } else {
            fmt.Printf("Acceso  guardado: %+v\n", access)
        }
    })
}