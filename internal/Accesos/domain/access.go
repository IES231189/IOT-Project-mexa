

package domain

import "time"

type Access struct {
    ID             int       `gorm:"primaryKey;autoIncrement"` 
    CodigoIngresado string    `gorm:"size:10;not null"`        
    Estado         string    `gorm:"type:enum('correcto', 'incorrecto');not null"`
    Fecha          time.Time `gorm:"default:CURRENT_TIMESTAMP"` 
}

func (Access) TableName() string {
    return "accesos"
}