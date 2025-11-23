package main

import (
	"log"

	"github.com/SCAPUTO88/desafio-nubank-GO/internal/config"
	"github.com/SCAPUTO88/desafio-nubank-GO/internal/domain"
)

func main() {
	cfg := config.Load()
	db := config.ConnectDB(cfg.DBURL)

	if err := db.AutoMigrate(&domain.Cliente{}, &domain.Contato{}); err != nil {
		log.Fatalf("Erro ao migrar tabelas: %v", err)
	}

	log.Println("✅ Migração concluída, tabelas prontas!")
}