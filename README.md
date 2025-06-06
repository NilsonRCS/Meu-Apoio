# MeuApoio - App de Apoio Emocional

## Arquitetura de Microserviços

Este projeto implementa uma arquitetura de microserviços em Go Lang para um aplicativo de apoio emocional.

### Serviços

1. **API Gateway** (`gateway/`) - Ponto único de entrada, roteamento e autenticação
2. **User Service** (`services/user/`) - Gestão de usuários, autenticação e perfis
3. **Content Service** (`services/content/`) - Informações de ajuda, artigos e histórias
4. **Audio Service** (`services/audio/`) - Meditações guiadas e músicas relaxantes
5. **Contact Service** (`services/contact/`) - Contatos de emergência e suporte
6. **Notification Service** (`services/notification/`) - Notificações e lembretes

### Tecnologias

- **Backend**: Go (Gin/Echo framework)
- **Banco de Dados**: PostgreSQL para dados estruturados, MongoDB para conteúdo
- **Cache**: Redis
- **Message Queue**: RabbitMQ/NATS
- **Storage**: MinIO/S3 para arquivos de áudio
- **Monitoring**: Prometheus + Grafana
- **Containerização**: Docker + Docker Compose

### Estrutura do Projeto

```
meuapoio/
├── gateway/                 # API Gateway
├── services/               # Microserviços
│   ├── user/              # Serviço de usuários
│   ├── content/           # Serviço de conteúdo
│   ├── audio/             # Serviço de áudio
│   ├── contact/           # Serviço de contatos
│   └── notification/      # Serviço de notificações
├── shared/                # Código compartilhado
│   ├── config/           # Configurações
│   ├── database/         # Conexões DB
│   ├── middleware/       # Middlewares comuns
│   └── utils/            # Utilitários
├── docker-compose.yml    # Orquestração dos serviços
└── docs/                # Documentação da API
```

## Como executar

```bash
# Subir todos os serviços
docker-compose up -d

# Ou executar individualmente
cd gateway && go run main.go
cd services/user && go run main.go
``` 