package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/api"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/config"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/event" // Importe o pacote event
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/handler"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/repository"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/service"
)

func main() {
	os.Setenv("PUBSUB_EMULATOR_HOST", "localhost:8085")

	cfg := config.Load()
	db := config.ConnectDB(cfg.DBURL)

	if err := db.AutoMigrate(&domain.Cliente{}, &domain.Contato{}); err != nil {
		log.Fatalf("Erro ao migrar tabelas: %v", err)
	}

	log.Println("✅ Migração concluída, tabelas prontas!")

	publisher, err := event.NewGCPPublisher("SCAPUT88_DESAFIO_GO")
	if err != nil {
		log.Fatalf("Erro ao criar publisher: %v", err)
	}
	defer publisher.Close()

	clienteRepo := repository.NewClienteRepository(db)
	contatoRepo := repository.NewContatoRepository(db)
	
	authService := service.NewAuthService()
	clienteService := service.NewClienteService(clienteRepo, publisher)
	contatoService := service.NewContatoService(contatoRepo, clienteRepo)

	authHandler := handler.NewAuthHandler(authService)
	clienteHandler := handler.NewClienteHandler(clienteService)
	contatoHandler := handler.NewContatoHandler(contatoService)

	router := api.NewRouter(clienteHandler, contatoHandler, authHandler, authService)

	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
		ErrorLog: log.New(os.Stderr, "server: ", log.Lshortfile),
		ReadTimeout:  10 * time.Second,
		WriteTimeout:	10 * time.Second,
		IdleTimeout: 	120 * time.Second,
	}

	log.Println("🚀 Servidor rodando na porta 8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}

}