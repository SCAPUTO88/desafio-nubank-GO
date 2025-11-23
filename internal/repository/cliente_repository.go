package repository

import (
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"

	"gorm.io/gorm"
)

type ClienteRepository struct {
	db *gorm.DB
}

func NewClienteRepository(db *gorm.DB) *ClienteRepository {
	return &ClienteRepository{db}
}

func (r *ClienteRepository) Create(cliente *domain.Cliente) error {
	return r.db.Create(cliente).Error
}

func (r *ClienteRepository) FindAllWithContatos() ([]domain.Cliente, error) {
	var clientes []domain.Cliente
	err := r.db.Preload("Contatos").Find(&clientes).Error
	return clientes, err
}

func (r *ClienteRepository) FindByID(id uint) (*domain.Cliente, error) {
	var cliente domain.Cliente
	if err := r.db.Preload("Contatos").First(&cliente, id).Error; err != nil {
		return nil, err
	}
	return &cliente, nil
}