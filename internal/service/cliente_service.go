package service

import (
	"errors"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/repository"
)

type ClienteService struct {
	repo repository.ClienteRepository
}

func NewClienteService(r repository.ClienteRepository) *ClienteService {
	return &ClienteService{repo: r}
}

func (s *ClienteService) Create(dto domain.CreateClienteDTO) (*domain.ClienteResponseDTO, error) {
	cliente := domain.Cliente{
		Nome: dto.Nome,
		Email: dto.Email,
	}
	if err := s.repo.Create(&cliente); err != nil {
		return nil, err
	}

	return &domain.ClienteResponseDTO{
		ID: cliente.ID,
		Nome: cliente.Nome,
		Email: cliente.Email,
	}, nil
}

func (s *ClienteService) ListAll() ([]domain.ClienteResponseDTO, error) {
	clientes, err := s.repo.FindAllWithContatos()
	if err != nil {
		return nil, err
	}

	var result []domain.ClienteResponseDTO
	for _, c := range clientes {
		dto := domain.ClienteResponseDTO{
			ID: c.ID,
			Nome: c.Nome,
			Email: c.Email,
		}
		for _, ct := range c.Contatos {
			dto.Contatos = append(dto.Contatos, domain.ContatoDTO{
				ID: ct.ID,
				Tipo: ct.Tipo,
				Valor: ct.Valor,
				ClienteId: ct.ClienteID,
			})
		}
		result = append(result, dto)
	}
	return result, nil
	}

	func (s *ClienteService) GetByID(id uint) (*domain.Cliente,error) {
		cliente, err := s.repo.FindByID(id)
		if err != nil {
			return nil, err
		}
		if cliente == nil {
			return nil, errors.New("cliente_not_found")
		}
		return cliente, nil
	}