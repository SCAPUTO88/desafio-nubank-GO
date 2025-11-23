# Arquitetura do Sistema - Desafio Backend Nubank

Este diagrama ilustra a **Clean Architecture** e o **Fluxo de Dados** para a API Go leve e segura.

```mermaid
graph TD
    subgraph "Externo"
        Client[Cliente HTTP]
    end

    subgraph "Ponto de Entrada (cmd/server)"
        Main[main.go]
        Router[Go 1.23 ServeMux]
    end

    subgraph "Camada de Segurança (internal/middleware)"
        SecMW[Middleware de Segurança]
        LogMW[Middleware de Logger]
        LimitMW[Limitador de Tamanho de Body]
    end

    subgraph "Camada de Interface (internal/handler)"
        CH[Handler de Cliente]
        CtH[Handler de Contato]
    end

    subgraph "Camada de Negócio (internal/service)"
        CS[Service de Cliente]
        CtS[Service de Contato]
    end

    subgraph "Camada de Acesso a Dados (internal/repository)"
        CR[Repository de Cliente]
        CtR[Repository de Contato]
    end

    subgraph "Camada de Domínio (internal/domain)"
        Ent[Entidades & DTOs]
    end

    subgraph "Infraestrutura"
        DB[(PostgreSQL)]
    end

    %% Conexões de Fluxo
    Client -->|"Requisição HTTP"| Main
    Main --> Router
    Router -->|"1. Requisição"| SecMW
    SecMW --> LimitMW
    LimitMW --> LogMW
    LogMW -->|"2. Requisição Validada"| CH
    LogMW -->|"2. Requisição Validada"| CtH

    CH -->|"3. Chama Lógica"| CS
    CtH -->|"3. Chama Lógica"| CtS

    CS -->|"4. Op. Dados"| CR
    CtS -->|"4. Op. Dados"| CtR

    CR -->|"5. Consulta SQL"| DB
    CtR -->|"5. Consulta SQL"| DB

    %% Dependências
    CH -.-> Ent
    CtH -.-> Ent
    CS -.-> Ent
    CtS -.-> Ent
    CR -.-> Ent
    CtR -.-> Ent
```

## Responsabilidades dos Componentes

1.  **Middleware de Segurança**:

    - **Strict-Transport-Security**: Força HTTPS.
    - **CSP & No-Sniff**: Mitiga XSS e MIME-sniffing.
    - **Body Limiter**: Previne DoS via payloads grandes.

2.  **Handlers**:

    - Parseiam entrada JSON.
    - Validam DTOs.
    - Mapeiam erros HTTP.

3.  **Services**:

    - Regras de negócio (ex: "Cliente deve existir para adicionar Contato").
    - Gerenciamento de transação (se necessário).

4.  **Repositories**:
    - Interações puras com banco de dados usando GORM.
    - Sem lógica de negócio.
