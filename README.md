# Card Authorization System

Projeto de estudo que simula a arquitetura de um sistema de autorização de cartões, similar ao utilizado por bandeiras e processadoras de pagamento.

## Visão Geral

O sistema recebe mensagens no formato **ISO 8583** (padrão utilizado em redes de pagamento), executa um pipeline de validações e decide aprovar ou negar a transação. Caso aprovada, publica um evento de clearing via **RabbitMQ**.

## Funcionalidades

- Parse de mensagens ISO 8583
- Consulta de dados de conta e cartão (Account Service)
- Validação de risco de fraude (Anti-Fraud Service)
- Validação de saldo (Ledger Service)
- Persistência da autorização (PostgreSQL)
- Publicação de evento de clearing (RabbitMQ)

## Stack

- **Go** — linguagem principal
- **RabbitMQ** (`amqp091-go`) — mensageria para clearing
- **PostgreSQL** — base de autorizações
- **Docker Compose** — infraestrutura local
- **PlantUML / C4 Model** — documentação de arquitetura

## Estrutura do Projeto

```
cmd/server/         → entrypoint da aplicação
internal/
  authorization/    → engine de autorização (service, processor, rules)
  clients/          → interfaces e mocks dos serviços externos
  domain/           → entidades de domínio
  messaging/        → publisher RabbitMQ (clearing)
  parser/           → parse de mensagem ISO 8583
  repository/       → persistência de autorizações
architecture/       → diagramas C4 em PlantUML
```

## Arquitetura (C4 Model)

Os diagramas estão em `/architecture` no formato PlantUML. Para visualizá-los, use a extensão [PlantUML](https://marketplace.visualstudio.com/items?itemName=jebbs.plantuml) no VS Code ou o site [plantuml.com](https://www.plantuml.com/plantuml).

---

### Nível 1 — Contexto

> Mostra os atores externos e sistemas com os quais o autorizador se relaciona.

![C4 Context](architecture/context.puml)

```plantuml
' architecture/context.puml
Person_Ext(network, "Card Network")          → envia ISO 8583
System(authorizer, "Card Authorization System")
System_Ext(account, "Account Service")
System_Ext(ledger, "Ledger Service")
System_Ext(antifraud, "Anti-Fraud Service")
System_Ext(clearing, "Clearing Service")
System_Ext(authdb, "Authorization Database")
```

---

### Nível 2 — Container

> Detalha os containers (processos/serviços) dentro do sistema de autorização.

![C4 Container](architecture/container.puml)

```
Authorization System
├── ISO Listener       (Go) — recebe mensagens ISO 8583 da bandeira
├── Authorization Engine (Go) — executa regras de autorização
└── Authorization DB   (PostgreSQL) — armazena autorizações

Sistemas externos: Account Service, Anti-Fraud Service, Ledger Service, Clearing Service
```

---

### Nível 3 — Componente

> Detalha os componentes internos do Authorization Engine.

![C4 Component](architecture/component.puml)

```
Authorization Engine
├── ISO Parser              — parse da mensagem ISO 8583
├── Pre-Auth Validator      — validações iniciais
├── Rule Engine             — executa regras de negócio
├── Account Client          — HTTP client → Account Service
├── Fraud Client            — HTTP client → Anti-Fraud Service
├── Ledger Client           — HTTP client → Ledger Service
├── Authorization Processor — decide approve/deny
├── Authorization Repository — persiste no PostgreSQL
└── Clearing Publisher      — publica evento no RabbitMQ
```

---

### Nível 4 — Código (Classes)

> Diagrama de classes das principais entidades do domínio e seus relacionamentos.

![C4 Code](architecture/codigo.puml)

```
authorization
├── AuthorizationService    → Authorize(request) AuthorizationResponse
├── AuthorizationProcessor  → Process(request) AuthorizationResponse
├── RuleEngine              → Validate(request) bool
├── AuthorizationRepository → Save(auth)
└── ClearingPublisher       → Publish(auth)

clients
├── AccountClient  → GetAccount(cardNumber)
├── FraudClient    → CheckFraud(request)
└── LedgerClient   → CheckBalance(accountId, amount)

domain
├── AuthorizationRequest  { cardNumber, amount, mcc, merchantId }
└── AuthorizationResponse { approved, responseCode }
```

## Como Rodar

```bash
# Subir infraestrutura (RabbitMQ, PostgreSQL)
docker compose up -d

# Rodar a aplicação
go run cmd/server/main.go
```
