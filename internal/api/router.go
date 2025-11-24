package api

import (
	"net/http"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/handler"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/middleware"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/service"
)

func NewRouter(
    clienteHandler *handler.ClienteHandler,
    contatoHandler *handler.ContatoHandler,
    authHandler *handler.AuthHandler,
    authService *service.AuthService,
) http.Handler {
    mux := http.NewServeMux()

    mux.HandleFunc("POST /login", authHandler.Login)

    authMiddleware := middleware.AuthMiddleware(authService)

    protected := func(h http.HandlerFunc) http.HandlerFunc {
        return authMiddleware(http.HandlerFunc(h)).ServeHTTP
    }

    mux.HandleFunc("POST /clientes", protected(clienteHandler.Create))
    mux.HandleFunc("GET /clientes", protected(clienteHandler.List))
    mux.HandleFunc("GET /clientes/{id}/contatos", protected(clienteHandler.ListContatos))
    mux.HandleFunc("POST /contatos", protected(contatoHandler.Create))

    var h http.Handler = mux
    h = middleware.SecurityHeaders(h)
    h = middleware.Logger(h)
    h = middleware.BodySizeLimiter(1024 * 1024)(h)
    h = middleware.RateLimitMiddleware(h)

    return h
}