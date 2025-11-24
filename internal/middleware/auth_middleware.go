package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/service"
)

type contextKey string

const UserIDKey contextKey = "user_id"

func AuthMiddleware(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Token ausente", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
				return
			}

			tokenStr := parts[1]

			claims, err := authService.ValidateToken(tokenStr)
			if err != nil {
				http.Error(w, "Token inválido: "+err.Error(), http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
			next.ServeHTTP(w, r.WithContext(ctx))
			})
		}
	}
