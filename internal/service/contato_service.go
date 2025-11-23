package service

import (
	"errors"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/repository"
)

var ErrClienteNotFound = errors.New("cliente_not_found")

type ContatoService struct {
    repoContato repository.ContatoRepository
    repoCliente repository.ClienteRepository
}

func NewContatoService(cr repository.ContatoRepository, cl repository.ClienteRepository) *ContatoService {
    return &ContatoService{repoContato: cr, repoCliente: cl}
}

func (s *ContatoService) Create(dto domain.CreateContatoDTO) (domain.ContatoDTO, error) {
    // valida se cliente existe
    cli, err := s.repoCliente.FindByID(dto.ClienteID)
    if err != nil {
        return domain.ContatoDTO{}, err
    }
    if cli == nil {
        return domain.ContatoDTO{}, ErrClienteNotFound
    }

    contato := domain.Contato{
        ClienteID: dto.ClienteID,
        Tipo:      dto.Tipo,
        Valor:     dto.Valor,
    }
    if err := s.repoContato.Create(&contato); err != nil {
        return domain.ContatoDTO{}, err
    }

    return domain.ContatoDTO{
        ID:        contato.ID,
        Tipo:      contato.Tipo,
        Valor:     contato.Valor,
        ClienteId: contato.ClienteID,
    }, nil
}

func (s *ContatoService) ListByClienteID(id uint) ([]domain.ContatoDTO, error) {
    contatos, err := s.repoContato.FindByClienteID(id)
    if err != nil {
        return nil, err
    }

    var result []domain.ContatoDTO
    for _, ct := range contatos {
        result = append(result, domain.ContatoDTO{
            ID:        ct.ID,
            Tipo:      ct.Tipo,
            Valor:     ct.Valor,
            ClienteId: ct.ClienteID,
        })
    }
    return result, nil
	}