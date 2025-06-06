# MeuApoio - App de Apoio Emocional

## 🎯 Sobre o Projeto

Este projeto implementa uma arquitetura de microserviços em Go Lang para um aplicativo de apoio emocional, oferecendo funcionalidades como meditação guiada, música relaxante, conteúdo de ajuda, histórias inspiradoras e rede de contatos de apoio.

## 📚 Documentação Completa

**👉 [Acesse a Documentação Completa em `docs/`](./docs/README.md)**

### 📖 Links Rápidos:
- **[Arquitetura do Sistema](./docs/architecture.md)** - Decisões técnicas e padrões
- **[Configuração dos Bancos](./docs/database.md)** - Setup de PostgreSQL, MongoDB e Redis

## 🚀 Quick Start

### Pré-requisitos
- Docker & Docker Compose
- Go 1.21+

### Executar o Ambiente
```bash
# Clonar o repositório
git clone <repo-url>
cd meuapoio

# Subir os bancos de dados
docker compose up -d

# Acessar interfaces administrativas
# PostgreSQL: http://localhost:8080 (Adminer)
# MongoDB: http://localhost:8081 (Mongo Express)

# Ver logs
docker compose logs -f

# Parar ambiente
docker compose down
```

## 🏗️ Arquitetura

### Microserviços
1. **API Gateway** (`gateway/`) - Ponto único de entrada, roteamento e autenticação
2. **User Service** (`services/user/`) - Gestão de usuários, autenticação e perfis
3. **Content Service** (`services/content/`) - Informações de ajuda, artigos e histórias
4. **Audio Service** (`services/audio/`) - Meditações guiadas e músicas relaxantes
5. **Contact Service** (`services/contact/`) - Contatos de emergência e suporte
6. **Notification Service** (`services/notification/`) - Notificações e lembretes

### Tecnologias
- **Backend**: Go (Gin framework)
- **Banco Principal**: PostgreSQL para dados estruturados
- **Banco de Conteúdo**: MongoDB para artigos e histórias
- **Cache**: Redis para sessões e performance
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
├── docs/                 # 📚 Documentação completa
├── scripts/              # Scripts de inicialização
└── docker-compose.yml    # Orquestração dos serviços
```

## 🔧 Desenvolvimento

### Para Novos Desenvolvedores
1. **Leia a documentação**: [docs/architecture.md](./docs/architecture.md)
2. **Configure o ambiente**: [docs/database.md](./docs/database.md)
3. **Execute**: `docker compose up -d`
4. **Comece com**: User Service (recomendado)

### Padrões de Commit
```
feat: nova funcionalidade
fix: correção de bug
docs: alteração na documentação
refactor: refatoração de código
test: adição ou modificação de testes
```

## 📊 Status do Projeto

### ✅ Fase 1: Foundations (Concluída)
- [x] Estrutura do projeto
- [x] Configuração dos bancos de dados
- [x] Docker Compose
- [x] Documentação arquitetural

### 🔄 Fase 2: Core Authentication (Em Andamento)
- [ ] User Service implementation
- [ ] JWT authentication flow
- [ ] API Gateway setup
- [ ] Basic CRUD operations

### 📋 Próximas Fases
- **Fase 3**: Business Logic (Audio, Content, Contact Services)
- **Fase 4**: Advanced Features (Notifications, Real-time)
- **Fase 5**: Production Ready (Kubernetes, CI/CD, Monitoring)

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -m 'feat: adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE) para mais detalhes.

---

**🏥 Importante**: Este é um projeto de apoio emocional. Se você ou alguém que conhece está passando por uma crise, procure ajuda profissional imediatamente.

- **CVV**: 188 (24h, gratuito)
- **SAMU**: 192
- **Bombeiros**: 193 