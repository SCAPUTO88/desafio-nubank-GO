package handler

import (
	  "encoding/json"
    "net/http"

    "github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
    "github.com/SCAPUTO88/desafio-nubank-GO/internal/service"
)

type ContatoHandler struct {
	service *service.ContatoService
}

func NewContatoHandler(s *service.ContatoService) *ContatoHandler {
	return &ContatoHandler{service: s}
}

func (h *ContatoHandler) Create(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var dto domain.CreateContatoDTO
    if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
        http.Error(w, "Erro ao ler JSON: "+err.Error(), http.StatusBadRequest)
        return
    }

    if dto.ClienteID == 0 || dto.Tipo == "" || dto.Valor == "" {
        http.Error(w, "ClienteID, Tipo e Valor são obrigatórios", http.StatusBadRequest)
        return
    }

    response, err := h.service.Create(dto)
    if err != nil {
        if err.Error() == "cliente_not_found" {
            http.Error(w, "Cliente não encontrado", http.StatusNotFound)
            return
        }
        http.Error(w, "Erro ao criar contato: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(response)
}