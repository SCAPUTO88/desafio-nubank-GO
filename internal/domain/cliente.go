package domain

import (
	"time"
)

type Cliente struct {
	ID uint `gorm:"primaryKey" json: "id"`
	Nome string `gorm:"size:120;not null" json:"nome"`
	Email string `gorm:"size:200;uniqueIndex;not null" json:"email"`
	Contatos []Contato `json:"contatos, omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}