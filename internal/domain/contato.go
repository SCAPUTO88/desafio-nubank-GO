package domain

import "time"

type Contato struct {
	ID uint `gorm:"primary_key" json:"id"`
	Tipo string `gorm:"size:50;not null" json:"tipo"`
	Valor string `gorm:size:200;not null" json:"valor"`
	ClienteID uint `gorm:"not null" json:"cliente_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}