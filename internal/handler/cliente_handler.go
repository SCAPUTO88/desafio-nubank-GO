package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

  "github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
  "github.com/SCAPUTO88/desafio-nubank-GO/internal/service"
)

type ClienteHandler struct {
	service *service.ClienteService
}

func NewClienteHandler(s *service.ClienteService) *ClienteHandler {
	return &ClienteHandler{service: s}
}

func (h *ClienteHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var dto domain.CreateClienteDTO
	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, "Erro ao ler JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if dto.Nome == "" || dto.Email == "" {
		http.Error(w, "Nome e Email são obrigatórios", http.StatusBadRequest)
		return
	}

	response, err := h.service.Create(dto)
	if err != nil {
		http.Error(w, "Erro ao criar clientes: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *ClienteHandler) List(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	clientes, err := h.service.ListAll()
	if err != nil {
		http.Error(w, "Erro ao listar clientes", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func (h *ClienteHandler) ListContatos(w http.ResponseWriter, r *http.Request) {
	
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "ID não fornecido", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	cliente, err := h.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente.Contatos)

}