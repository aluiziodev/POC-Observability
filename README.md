# Order Service


Este repositório é uma POC de observabilidade em Go.


## Objetivo


A ideia da POC é estruturar um serviço com stack de observabilidade completa. O foco não está na lógica de negócio, mas na infraestrutura de observabilidade em volta dela.




## Status atual


O projeto conta com:


- API REST para criar, listar, consultar e atualizar status de pedidos
- Arquitetura em camadas: handler, service e repository
- Métricas (Prometheus)**: métricas técnicas e de negócio, utilizando o padrao RED method, expostas em `/metrics`
- Logs estruturados (JSON via `log/slog`): correlacionados por `request_id` e `trace_id`
- Tracing distribuído (OpenTelemetry): span raiz por requisição, propagação de contexto via header `traceparent`, e spans filhos em cada consulta ao banco
- Endpoints de saúde: `/health` e `/ready`
- Configuração via variáveis de ambiente
- Banco de dados PostgreSQL


## Tecnologias


- Go
- PostgreSQL
- Docker / Docker Compose
- `prometheus/client_golang`
- `go.opentelemetry.io/otel`
- `lib/pq`
- `godotenv`
- `google/uuid`


## Estrutura do projeto


- `cmd/api/main.go`: ponto de entrada da aplicação sobe o servidor HTTP, inicializa o banco e o Tracer e trata encerramento gracioso
- `internal/config`: configuração da aplicação
- `internal/handler`: handlers HTTP
- `internal/service`: regras de negócio
- `internal/repository`: acesso ao banco de dados
- `internal/observability`:
 - `technicalMetrics.go`, `bussinesMetrics.go`, `handler.go`: métricas utilizando Prometheus e o handler do `/metrics`
 - `logger.go`: logger JSON e helpers `WarnContext`/`ErrorContext`
 - `middleware.go`: cria o span raiz, propaga `trace_id`, loga a requisição e alimenta as métricas
 - `tracing.go`: inicialização do `TracerProvider` e exportação via OTLP/HTTP
 - `requestId.go`: geração/leitura do `request_id`
- `internal/router`: definição das rotas
- `internal/models`: modelos de request/response
- `internal/domain`: entidades de domínio
- `internal/response`: helpers de resposta HTTP
- `sql/postgresql.sql`: script de criação da tabela no banco de dados


## Endpoints disponíveis


### Saúde e observabilidade


- `GET /health`
- `GET /ready`
- `GET /metrics` — métricas no formato do Prometheus


### Pedidos


- `POST /orders`
- `GET /orders`
- `GET /orders/{id}`
- `PATCH /orders/{id}/status`


#### Exemplo de criação de pedido (`POST /orders`)


```json
{
 "item": "Notebook",
 "quantity": 2
}
```


## Observabilidade


### Métricas


Utilizando Prometheus, expostas em `GET /metrics`:


| Métrica | Tipo | Labels | O que mede |
|---|---|---|---|
| `http_requests_total` | counter | `method`, `path`, `status` | total de requisições HTTP |
| `http_request_duration` | histograma | `method`, `path` | latência das requisições HTTP |
| `database_query_duration` | histograma | `operation` | latência das queries ao banco (`orders.create`, `orders.get_all`, `orders.get_by_id`, `orders.update_status`) |
| `orders_created_total` | counter | (não possui) | total de pedidos criados com sucesso |


O label `path` sempre usa o padrão da rota (ex: `/orders/{id}`), nunca a URL crua com o ID real, para evitar explosão de cardinalidade.


### Logs estruturados


Formato JSON, em `stdout`, com `request_id` e `trace_id` para correlação com os traces:


```json
{"time":"2026-07-17T14:00:00Z","level":"INFO","msg":"http_request","request_id":"a1b2c3d4","trace_id":"4bf92f3577b34da6a3ce929d0e0e4736","method":"GET","path":"/orders/xyz","status":200,"duration":"12ms"}
```


Erros de negócio nos handlers também são logados (`WARN` para erro de cliente/4xx, `ERROR` para erro de negócio ou infra/5xx), carregando o mesmo `request_id` da linha de requisição.


### Tracing distribuído (OpenTelemetry)


- Span raiz criado no middleware HTTP para cada requisição, nomeado com o padrão da rota (ex: `GET /orders/{id}`).
- Propagação de contexto via header `traceparent` (padrão W3C Trace Context), se a requisição já chegar com um trace em andamento, ele é continuado em vez de criar um novo.
- Um span filho é criado por operação de banco no repository (ex: `db.orders.create`, `db.orders.get_by_id`), aninhado sob o span da requisição.
- Exportação via OTLP/HTTP para um coletor (Jaeger, Tempo, ou o OpenTelemetry Collector), endpoint configurável via `OTEL_EXPORTER_OTLP_ENDPOINT`.


## Como executar


### 1. Subir a infraestrutura (Banco de dados + Jaeger)


```bash
docker compose up -d
```


Esse comando sobe:
- **Postgres** em `localhost:5432`
- **Jaeger** (UI em `http://localhost:16686`, recebendo traces via OTLP/HTTP em `localhost:4318`)


### 2. Configurar as variáveis de ambiente


Exemplo (`.env`):


```bash
export POSTGRES_USER=admin
export POSTGRES_PASSWORD=admin123
export POSTGRES_DB=observability
```


Variáveis opcionais:


```bash
export API_PORT=8080 # padrão: 8080
export OTEL_EXPORTER_OTLP_ENDPOINT=localhost:4318 # padrão: localhost:4318
```


### 3. Rodar a aplicação


```bash
go run ./cmd/api
```


### 4. Verificar a observabilidade


- Métricas: `curl http://localhost:8080/metrics`
- Traces: `http://localhost:16686` (procure pelo serviço `order-service`)
- Logs: stdout do processo


## Banco de dados


O projeto espera um banco PostgreSQL com a tabela criada a partir do script em `sql/postgresql.sql`.


## Próximos passos


- Integração com Prometheus + Grafana (scrape do `/metrics` e dashboards)
- Padronização de respostas e tratamento de erros
- worker assíncrono, validando a propagação de trace através de uma fila (trace atravessando de um processo síncrono para um assíncrono)
