package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(s *service.AuthService) *AuthHandler {
	return &AuthHandler{authService: s}
}

type LoginRequest struct {
	Email    string `json:"email"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    var req LoginRequest
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Erro ao ler JSON", http.StatusBadRequest)
        return
    }

		if req.Email != "admin@desafio.com.br" {
			http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
			return
		}

		token, err := h.authService.GenerateToken(1)
		if err != nil {
			http.Error(w, "Erro ao gerar token", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(LoginResponse{Token: token})
	}