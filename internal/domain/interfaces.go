package domain

type ClienteRepository interface {
	Create(cliente *Cliente) error
	FindAllWithContatos() ([]Cliente, error)
	FindByID(id uint) (*Cliente, error)
}

type ContatoRepository interface {
	Create(contato *Contato) error
	FindByClienteID(clienteId uint) ([]Contato, error)
}