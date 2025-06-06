# 🚀 Guia de Onboarding - MeuApoio

Bem-vindo ao time de desenvolvimento do **MeuApoio**! 👋 Este guia vai te levar do zero até estar produtivo no projeto.

## 📋 Checklist de Onboarding

### 🏁 **Dia 1: Setup Inicial**
- [ ] Configurar ambiente de desenvolvimento
- [ ] Clonar repositório e executar primeira build
- [ ] Subir ambiente local com bancos de dados
- [ ] Explorar estrutura do projeto
- [ ] Ler documentação arquitetural básica

### 📚 **Semana 1: Entendimento**
- [ ] Compreender domínios de negócio
- [ ] Entender padrões de código utilizados
- [ ] Fazer primeira contribuição (bug fix simples)
- [ ] Participar de code review
- [ ] Configurar ferramentas de desenvolvimento

### 🛠️ **Semana 2-3: Desenvolvimento**
- [ ] Implementar primeira feature
- [ ] Escrever testes unitários
- [ ] Entender processo de deploy
- [ ] Colaborar em pair programming

---

## 🔧 Setup do Ambiente

### Pré-requisitos Obrigatórios

```bash
# Verificar instalações
go version        # Go 1.21+
docker --version  # Docker 20.0+
git --version     # Git 2.0+
```

**Não tem algo instalado?**
- **Go**: https://golang.org/doc/install
- **Docker**: https://docs.docker.com/get-docker/
- **Git**: https://git-scm.com/downloads

### Ferramentas Recomendadas

**Editor/IDE:**
- **VS Code** + Go Extension (recomendado para iniciantes)
- **GoLand** (JetBrains, pago mas excelente)
- **Neovim** + gopls (para ninjas 🥷)

**Extensões VS Code Essenciais:**
```bash
# Instalar via VS Code ou comando:
code --install-extension golang.go
code --install-extension ms-vscode.docker
code --install-extension bradlc.vscode-tailwindcss
code --install-extension humao.rest-client
```

**CLI Tools Úteis:**
```bash
# Ferramentas Go
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install golang.org/x/tools/cmd/goimports@latest

# Database clients (opcional)
brew install postgresql  # psql client
brew install mongosh     # MongoDB shell
brew install redis       # redis-cli
```

---

## 🏗️ Primeira Execução

### 1. Clone e Setup
```bash
# Clone do repositório
git clone <repo-url>
cd meuapoio

# Baixar dependências Go
go mod download
go mod tidy

# Verificar se tudo compila
go build ./...
```

### 2. Subir Ambiente Local
```bash
# Subir bancos de dados
docker compose up -d

# Verificar se containers estão rodando
docker compose ps

# Ver logs se algo der errado
docker compose logs -f
```

### 3. Acessar Interfaces Admin
- **PostgreSQL**: http://localhost:8080 (Adminer)
  - Sistema: PostgreSQL
  - Servidor: localhost:5432
  - Usuário: postgres
  - Senha: postgres123
  - Base: meuapoio

- **MongoDB**: http://localhost:8081 (Mongo Express)
  - Usuário: mongo
  - Senha: mongo123

### 4. Testar Conexões
```bash
# PostgreSQL
psql -h localhost -p 5432 -U postgres -d meuapoio
# \dt para listar tabelas
# \q para sair

# Redis
redis-cli -h localhost -p 6379
# ping (deve retornar PONG)
# exit

# MongoDB
mongosh "mongodb://mongo:mongo123@localhost:27017/meuapoio"
# show collections
# exit
```

---

## 🏛️ Entendendo a Arquitetura

### Conceitos Fundamentais

**1. Domain-Driven Design (DDD)**
```
📦 MeuApoio é dividido em domínios:
├── 👤 Users - Autenticação e perfis
├── 🎵 Audio - Meditação e música  
├── 📝 Content - Artigos e histórias
├── 📞 Contact - Rede de apoio
└── 🔔 Notification - Comunicação
```

**2. Database per Service**
- Cada microserviço tem seu próprio banco
- PostgreSQL para dados estruturados
- MongoDB para conteúdo flexível
- Redis para cache e sessões

**3. Shared Kernel**
- Pasta `shared/` tem código comum
- Config, middleware, utils reutilizáveis
- Evita duplicação entre serviços

### Fluxo de Dados Típico
```
1. 📱 App faz request → API Gateway
2. 🌐 Gateway autentica → User Service  
3. 🔐 User Service valida → PostgreSQL
4. ✅ Retorna dados → App
```

---

## 📁 Navegando no Código

### Estrutura de Pastas - O que é o quê?

```
meuapoio/
├── 📂 docs/              # 📚 COMECE AQUI - Documentação
│   ├── architecture.md   # 🏗️ Decisões arquiteturais  
│   ├── database.md      # 🗄️ Setup dos bancos
│   └── onboarding.md    # 🚀 Este arquivo!
│
├── 📂 shared/            # 🔗 Código compartilhado
│   ├── config/          # ⚙️ Configurações centralizadas
│   ├── database/        # 🔌 Conexões com bancos
│   ├── middleware/      # 🛡️ Auth, CORS, Rate limiting
│   └── utils/           # 🛠️ Funções auxiliares
│
├── 📂 gateway/           # 🌐 API Gateway (porta 8080)
│   ├── main.go          # 🚪 Ponto de entrada
│   ├── routes/          # 🛣️ Roteamento
│   └── handlers/        # 🎯 Lógica das rotas
│
├── 📂 services/          # 🎛️ Microserviços
│   ├── user/            # 👤 Usuários (porta 8081)
│   ├── content/         # 📝 Conteúdo (porta 8082)
│   ├── audio/           # 🎵 Áudio (porta 8083)
│   ├── contact/         # 📞 Contatos (porta 8084)
│   └── notification/    # 🔔 Notificações (porta 8085)
│
├── 📂 scripts/           # 📜 Scripts de inicialização
│   └── init.sql         # 🗃️ Schema do PostgreSQL
│
├── 🐳 docker-compose.yml # 📦 Orquestração dos containers
├── 📄 go.mod            # 📋 Dependências Go
└── 📖 README.md         # 📌 Visão geral do projeto
```

### Padrão de Organização por Serviço
```
services/user/
├── 📄 main.go           # Entry point do serviço
├── 📂 handlers/         # HTTP handlers (controllers)
├── 📂 models/           # Structs de dados
├── 📂 repository/       # Acesso a dados (DB layer)
├── 📂 service/          # Lógica de negócio  
├── 📂 routes/           # Definição das rotas
└── 🐳 Dockerfile        # Container do serviço
```

---

## 🛠️ Padrões de Desenvolvimento

### Convenções de Nomenclatura

**Arquivos e Pastas:**
```go
// ✅ Bom
user_service.go
audio_handler.go
database_config.go

// ❌ Evite
UserService.go
audioHandler.go
databaseConfig.go
```

**Variáveis e Funções:**
```go
// ✅ Bom - camelCase
var userRepository UserRepository
func createUser(userData User) error

// ✅ Bom - Exported (PascalCase)
func CreateUserHandler(c *gin.Context)
type UserService interface{}

// ❌ Evite
var user_repository UserRepository
func Create_User(user_data User) error
```

**Constantes:**
```go
// ✅ Bom
const (
    DefaultPort = 8080
    MaxRetries  = 3
    APIVersion  = "v1"
)
```

### Estrutura de Handler Padrão
```go
package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/meuapoio/shared/middleware"
)

// CreateUserHandler cria um novo usuário
func CreateUserHandler(userService *service.UserService) gin.HandlerFunc {
    return func(c *gin.Context) {
        // 1. Validar entrada
        var req CreateUserRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        // 2. Chamar serviço
        user, err := userService.CreateUser(req)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // 3. Retornar resposta
        c.JSON(http.StatusCreated, user)
    }
}
```

### Padrão de Testes
```go
package handlers_test

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCreateUserHandler(t *testing.T) {
    // Arrange
    userService := &mockUserService{}
    handler := CreateUserHandler(userService)
    
    // Act
    // ... código de teste
    
    // Assert
    assert.Equal(t, http.StatusCreated, response.Code)
}
```

---

## 🎯 Primeira Contribuição

### Começando com Bug Fix Simples

**1. Escolha uma issue marcada como `good-first-issue`**
```bash
# Criar branch para sua mudança
git checkout -b fix/issue-123-user-validation

# Fazer suas mudanças
# ...

# Commit seguindo padrão
git commit -m "fix: validação de email em CreateUser

- Adiciona validação de formato de email
- Adiciona testes para casos de erro
- Fixes #123"
```

**2. Padrões de Commit**
```
feat: nova funcionalidade
fix: correção de bug  
docs: mudança na documentação
refactor: refatoração sem mudança de comportamento
test: adição ou modificação de testes
style: mudanças de formatação
perf: melhoria de performance
chore: mudanças no build, dependências, etc
```

**3. Checklist antes do PR**
- [ ] Código compila sem erros: `go build ./...`
- [ ] Testes passam: `go test ./...`
- [ ] Linter limpo: `golangci-lint run`
- [ ] Documentação atualizada se necessário

---

## 🐛 Debugging e Troubleshooting

### Problemas Comuns

**🚨 "Containers não sobem"**
```bash
# Verificar se portas estão ocupadas
lsof -i :5432 -i :27017 -i :6379

# Limpar containers antigos
docker compose down -v
docker system prune -f

# Subir novamente
docker compose up -d
```

**🚨 "Erro de conexão com banco"**
```bash
# Verificar se containers estão healthy
docker compose ps

# Ver logs dos containers
docker compose logs postgres
docker compose logs mongodb
```

**🚨 "Import não encontrado"**
```bash
# Limpar cache de módulos
go clean -modcache
go mod download
go mod tidy
```

### Logs e Debugging

**Ver logs dos serviços:**
```bash
# Todos os containers
docker compose logs -f

# Container específico  
docker compose logs -f postgres

# Logs do Go (quando desenvolver)
go run main.go 2>&1 | tee app.log
```

**Debug no VS Code:**
1. Criar `.vscode/launch.json`:
```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Debug User Service",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "./services/user/main.go",
            "env": {
                "PORT": "8081"
            }
        }
    ]
}
```

---

## 📚 Recursos de Aprendizado

### Go Lang - Se você é novo em Go

**📖 Leitura Essencial:**
- [Tour of Go](https://tour.golang.org/) - Tutorial interativo oficial
- [Effective Go](https://golang.org/doc/effective_go.html) - Boas práticas
- [Go by Example](https://gobyexample.com/) - Exemplos práticos

**🎥 Vídeos Recomendados:**
- "Learn Go Programming" - TechWorld with Nana
- "Go Course" - FreeCodeCamp
- "Golang Tutorial" - Derek Banas

### Microserviços & Arquitetura

**📚 Livros:**
- "Building Microservices" - Sam Newman
- "Clean Architecture" - Robert C. Martin
- "Microservices Patterns" - Chris Richardson

**🔗 Artigos:**
- [Martin Fowler on Microservices](https://martinfowler.com/articles/microservices.html)
- [Microservices.io](https://microservices.io/) - Patterns catalog

### Ferramentas Específicas

**🐳 Docker:**
- [Docker Getting Started](https://docs.docker.com/get-started/)
- [Docker Compose Tutorial](https://docs.docker.com/compose/gettingstarted/)

**🗄️ Bancos de Dados:**
- [PostgreSQL Tutorial](https://www.postgresqltutorial.com/)
- [MongoDB University](https://university.mongodb.com/)
- [Redis Tutorial](https://redis.io/docs/manual/)

---

## 🤝 Cultura e Processo do Time

### Code Review Guidelines

**👀 O que procurar ao revisar:**
- Código funciona e atende requisitos?
- Testes adequados estão incluídos?
- Código é legível e bem documentado?
- Segue padrões estabelecidos?
- Performance é adequada?

**✍️ Como dar feedback construtivo:**
```
✅ Bom: "Consider using a map here for O(1) lookup instead of array iteration"
❌ Evite: "This is wrong"

✅ Bom: "Could you add a test for the error case?"  
❌ Evite: "No tests"
```

### Comunicação

**💬 Canais:**
- Slack `#meuapoio-dev` - Discussões gerais
- Slack `#meuapoio-tech` - Discussões técnicas
- Issues GitHub - Bugs e features
- PRs GitHub - Code review

**🕐 Ceremonias:**
- Daily Standup - 9h (15min)
- Sprint Planning - Segunda (1h)
- Retrospective - Sexta (30min)
- Tech Talk - Quinta (30min opcional)

---

## 🎯 Metas dos Primeiros 30 Dias

### Semana 1: Orientação
- [ ] Setup completo do ambiente
- [ ] Entendimento da arquitetura
- [ ] Primeira contribuição (bug fix)
- [ ] Participação ativa em code reviews

### Semana 2: Desenvolvimento
- [ ] Implementar primeira feature simples
- [ ] Escrever testes unitários
- [ ] Melhorar algo existente (refactor)

### Semana 3: Autonomia
- [ ] Propor melhoria técnica
- [ ] Contribuir para documentação
- [ ] Mentorear próximo onboarding

### Semana 4: Impacto
- [ ] Liderar implementação de feature
- [ ] Apresentar resultado para o time
- [ ] Contribuir para planning

---

## 🆘 Quando Precisar de Ajuda

### Hierarquia de Suporte

**1. 📖 Self-Service (tente primeiro):**
- Esta documentação
- README dos serviços específicos
- Issues antigas no GitHub
- Stack Overflow

**2. 🤝 Peer Support:**
- Pergunte no Slack `#meuapoio-dev`
- Pair programming com colega
- Code review discussions

**3. 👨‍💼 Lead/Senior:**
- Decisões arquiteturais
- Bloqueios técnicos complexos
- Direcionamento de carreira

**4. 🚨 Emergência:**
- Production down
- Security issues
- Data loss

### Contatos Importantes

| Área | Pessoa | Slack | Responsabilidade |
|------|--------|-------|------------------|
| **Tech Lead** | @lead | `@tech-lead` | Decisões técnicas, arquitetura |
| **Backend Senior** | @senior | `@backend-senior` | Code review, mentoria |
| **DevOps** | @devops | `@devops` | Infraestrutura, deploy |
| **Product Owner** | @po | `@product` | Requirements, priorização |

---

## 🎉 Parabéns!

Se chegou até aqui, você está pronto para contribuir com o **MeuApoio**! 🚀

**Próximos passos:**
1. Configure seu ambiente seguindo este guia
2. Explore o código por algumas horas
3. Pegue sua primeira issue `good-first-issue`
4. Apareça no Slack para se apresentar! 👋

**Lembre-se:**
- ❓ Perguntar é sempre bem-vindo
- 🐛 Cometer erros faz parte do aprendizado  
- 🤝 Colaboração > competição
- 💡 Suas ideias são valiosas!

---

**Welcome to the team! 🎉**

*Última atualização: 06/06/2024*  
*Próxima revisão: Mensal ou quando processo mudar* 