
package domain

type Distance struct {
    ID        string  `gorm:"primaryKey"`
    Value     float64 // Valor de la distancia en cm
    Timestamp int64   // Fecha y hora del evento
}