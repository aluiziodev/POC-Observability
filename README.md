# Order Service

Este repositório é uma POC de observabilidade em Go.

## Objetivo

A ideia da POC é  estruturar um serviço com stack de observabilidade completa. O foco não está na lógica de negócio, mas na infraestrutura de observabilidade em volta dela.


A feature atual implementada é a API, enquanto a camada de observabilidade está presente de forma inicial.

## Status atual

O projeto conta com:

- API REST para criar, listar, consultar e atualizar status de pedidos
- Arquitetura em camadas: handlers, services e repository
- Logger JSON 
- Endpoints de saúde: /health e /ready
- Configuração via variáveis de ambiente
- Banco de dados PostgreSQL
  
## Tecnologias

- Go
- PostgreSQL
- Docker
- Gorilla Mux
- lib/pq
- godotenv

## Estrutura do projeto

- cmd/api/main.go: ponto de entrada da aplicação
- internal/config: configuração da aplicação
- internal/handler: handlers HTTP
- internal/service: regras de negócio
- internal/repository: acesso ao banco de dados
- internal/observability: logs e métricas (Ainda serão implementadas)
- internal/router: definição das rotas
- sql/postgresql.sql: script de criação da tabela no banco de dados

## Endpoints disponíveis

### Saúde e observabilidade

- GET /health
- GET /ready
- GET /metrics

### Pedidos

- POST /orders
- GET /orders
- GET /orders/{id}
- PATCH /orders/{id}/status

## Exemplo de criação de pedido (POST - /orders)

```json
{
  "item": "Notebook",
  "quantity": 2
}
```

## Como executar

### 1. Subir o banco de dados

```bash
docker compose up -d 
```

### 2. Configurar as variáveis de ambiente

Exemplo:

```bash
export POSTGRES_USER=postgres
export POSTGRES_PASSWORD=postgres
export POSTGRES_DB=orders
export PORT=8080
```

### 3. Rodar a aplicação

```bash
go run ./cmd/api
```

## Banco de dados

O projeto espera um banco PostgreSQL com a tabela criada a partir do script em sql/postgresql.sql

## Próximos passos

Os próximos passos da POC são:

- observabilidade com tracing
- integração com ferramentas como Prometheus e Grafana
- padronização de respostas e tratamento de erros
