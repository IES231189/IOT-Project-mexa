package services

import(
	"fmt"
	"iot-project/internal/Distancias/domain"
	"iot-project/internal/Distancias/infraestructure/repository"

)



type DistanceSensorServices struct{
	mqttRepo *repository.MqttDistanceRepository
	dbRepo *repository.DBRepository

}

func NewDistanceSensorService(mqttRepo *repository.MqttDistanceRepository , dbRepo *repository.DBRepository)*DistanceSensorServices{
	return &DistanceSensorServices{
		mqttRepo: mqttRepo,
		dbRepo:dbRepo,
	}
}

func (s * DistanceSensorServices) StartListening(topic string){
	s.mqttRepo.Listen(topic ,func(distance domain.Distance){
		if distance.ValidDistance(){
				distance.IsProximity()
			if err :=s.dbRepo.Save(distance); err != nil{
				fmt.Printf("Error al guardar la distancia de los datos:%s\n" ,err)
			}else{
				fmt.Printf("Datos de la distancia guardada: %+v\n",distance)
			}
		}else{
			fmt.Printf("Distancia no valida: %+v\n",distance)
		}
	} )
}