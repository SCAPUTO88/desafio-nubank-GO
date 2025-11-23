package repository

import (
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
	"gorm.io/gorm"
)

type ContatoRepository struct {
	db *gorm.DB
}

func NewContatoRepository(db *gorm.DB) *ContatoRepository {
	return &ContatoRepository{db}
}

func (r *ContatoRepository) Create(contato *domain.Contato) error {
	return r.db.Create(contato).Error
}

func (r *ContatoRepository) FindByClienteID(clienteID uint) ([]domain.Contato, error) {
	var contatos []domain.Contato
	err := r.db.Where("cliente_id = ?", clienteID).Find(&contatos).Error
	return contatos, err
}