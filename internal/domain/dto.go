package domain

type CreateClienteDTO struct {
	Nome  string `json:"nome" binding:"required, min=2"`
	Email string `json:"email" binding:"required,email"`
}

type ClienteResponseDTO struct {
	ID       uint         `json:"id"`
	Nome     string       `json:"nome"`
	Email    string       `json:"email"`
	Contatos []ContatoDTO `json:"contatos"`
}

type CreateContatoDTO struct {
	ClienteID uint   `json:"cliente_id" binding:"required"`
	Tipo      string `json:"tipo" binding:"required, oneof=telefone email whatsapp"`
	Valor     string `json:"valor" binding:"required"`
}

type ContatoDTO struct {
	ID        uint   `json:"id"`
	Tipo      string `json:"tipo"`
	Valor     string `json:"valor"`
	ClienteId uint   `json:"cliente_id"`
}