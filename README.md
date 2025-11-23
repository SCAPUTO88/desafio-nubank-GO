# Desafio Backend Nubank - Implementação em Go 🚀

Esta é uma implementação robusta, segura e performática do desafio técnico de backend do Nubank, escrita inteiramente em **Go (Golang)**.

O projeto foca em **Clean Architecture**, **Segurança Extrema** e **Baixo Acoplamento**, evitando frameworks web pesados e priorizando a biblioteca padrão do Go.

Descrição do desafio:

Desafio Técnico - Vaga Backend - Nubank

Olá candidato(a), Este desafio faz parte do processo seletivo para a vaga de Desenvolvedor Backend no Nubank. Nosso objetivo é avaliar sua capacidade de estruturar uma API RESTful utilizando boas práticas, relacionamentos de entidades e persistência de dados com Spring Boot e PostgreSQL.

Desafio Construa uma API REST para gerenciamento de clientes e seus contatos. Cada cliente pode ter um ou mais contatos associados.

Requisitos Técnicos Sua aplicação deve conter: • Cadastro de Cliente (POST /clientes) • Cadastro de Contato (POST /contatos) associado a um cliente existente • Listagem de todos os clientes com seus contatos (GET /clientes) • Listagem de contatos de um cliente específico (GET /clientes/{id}/contatos) • Uso de Spring Boot + Spring Data JPA • Banco de dados PostgreSQL • Entidades Cliente e Contato com relacionamento OneToMany / ManyToOne

Requisitos de Código Esperamos que seu código siga boas práticas de desenvolvimento, incluindo: • Separação de responsabilidades (camadas: controller, service, repository) • Uso de DTOs para entrada e saída de dados • Tratamento adequado de erros • Usar Lombok

Diferenciais (Não obrigatórios) • Uso de Docker para subir o PostgreSQL

• Testes automatizados • Documentação com Swagger

## 🛠 Tech Stack

- **Linguagem**: Go 1.23+
- **Banco de Dados**: PostgreSQL
- **ORM**: GORM
- **Autenticação**: JWT (JSON Web Tokens)
- **Testes**: Testify + Mocks
- **Segurança**: Rate Limiting, Security Headers, Body Limiter

## 🏛 Arquitetura

O projeto segue os princípios da **Clean Architecture**, garantindo separação de responsabilidades e testabilidade:

- `cmd/server`: Ponto de entrada (Main).
- `internal/domain`: Entidades e Interfaces (Core do negócio).
- `internal/repository`: Implementação do acesso a dados (GORM).
- `internal/service`: Regras de negócio puras.
- `internal/handler`: Controladores HTTP (Entrada/Saída).
- `internal/middleware`: Camada de segurança e interceptadores.

## 🔒 Funcionalidades de Segurança (Destaque)

Camadas de segurança foram implementadas:

1.  **Autenticação JWT**: Proteção contra acesso não autorizado.
2.  **Proteção contra IDOR**: Validação de contexto de usuário via Token.
3.  **Rate Limiting**: Proteção contra ataques de força bruta e DoS (Token Bucket Algorithm).
4.  **Security Headers**:
    - `Strict-Transport-Security` (HSTS)
    - `Content-Security-Policy` (CSP)
    - `X-Frame-Options` (Clickjacking protection)
    - `X-XSS-Protection`
5.  **Body Size Limiter**: Previne ataques de exaustão de memória limitando payloads a 1MB.
6.  **Sanitização**: Uso de GORM e encoding/json previne SQL Injection e XSS.

## 🚀 Como Rodar

### Pré-requisitos

- Go 1.23+
- Docker (para o banco de dados)

### Passos

1.  **Subir o Banco de Dados**:

    ```bash
    docker-compose -f build/docker-compose.yml up -d
    ```

2.  **Rodar a Aplicação**:
    ```bash
    go run cmd/server/main.go
    ```
    O servidor iniciará em `http://localhost:8080`.

## 🧪 Testes

O projeto possui cobertura de testes unitários na camada de serviço, utilizando Mocks para isolar o banco de dados.

```bash
go test ./internal/service/... -v
```

## 📡 API Endpoints

### Autenticação

- `POST /login`: Recebe email, retorna JWT.
  - _Email Admin_: `admin@desafio.com.br`

### Clientes (Requer Token Bearer)

- `POST /clientes`: Cria um novo cliente.
- `GET /clientes`: Lista todos os clientes.
- `GET /clientes/{id}/contatos`: Lista contatos de um cliente.

### Contatos (Requer Token Bearer)

- `POST /contatos`: Adiciona um contato a um cliente.

---

**Desenvolvido com foco em Excelência Técnica.**
