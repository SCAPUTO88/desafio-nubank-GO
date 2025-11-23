package api

import (
    "net/http"

    "github.com/SCAPUTO88/desafio-nubank-GO/internal/handler"
    "github.com/SCAPUTO88/desafio-nubank-GO/internal/middleware"
)

func NewRouter(clienteHandler *handler.ClienteHandler, contatoHandler *handler.ContatoHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /clientes", clienteHandler.Create)
	mux.HandleFunc("GET /clientes", clienteHandler.List)
	mux.HandleFunc("GET /clientes/{id}/contatos", clienteHandler.ListContatos)

	mux.HandleFunc("POST /contatos", contatoHandler.Create)

	var h http.Handler = mux
	h = middleware.SecurityHeaders(h)
	h = middleware.Logger(h)
	h = middleware.BodySizeLimiter(1024 * 1024)(h)

	return h

}