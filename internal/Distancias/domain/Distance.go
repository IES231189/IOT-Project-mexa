package domain
import "time"

type Distance struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	Distancia  float64
	Proximidad string
	Fecha      time.Time
}

func (D *Distance) ValidDistance() bool {
	return D.Distancia <= 100 && D.Distancia > 0
}

func (D *Distance) IsProximity() {

	if D.Distancia <= 100 && D.Distancia > 50 {
		D.Proximidad = "lejos"
	} else if D.Distancia > 100 {
		D.Proximidad = "muy lejos"
	} else if D.Distancia <= 50 && D.Distancia >=20 {
		D.Proximidad = "cerca"
	} else if D.Distancia <20 && D.Distancia>0 {
		D.Proximidad = "muy cerca"
	}
}

