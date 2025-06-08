# 🚀 Getting Started - MeuApoio

Guia para executar o **User Service** e **API Gateway** implementados.

## 📋 Pré-requisitos

- **Docker & Docker Compose** instalados
- **Go 1.21+** instalado
- **curl** e **jq** para testes (opcional)

## 🎯 O que foi implementado

### ✅ **User Service** (Porta 8081)
- ✅ Registro de usuários
- ✅ Login com JWT
- ✅ CRUD de perfil
- ✅ Contatos de emergência
- ✅ Autenticação middleware

### ✅ **API Gateway** (Porta 8080)
- ✅ Proxy reverso para microserviços
- ✅ Autenticação centralizada
- ✅ Rate limiting (100 req/min)
- ✅ CORS configurado
- ✅ Health checks

## 🏃‍♂️ Como executar

### 1. **Subir bancos de dados**
```bash
# Na raiz do projeto
docker compose up -d

# Verificar se estão rodando
docker compose ps
```

### 2. **Instalar dependências Go**
```bash
go mod tidy
```

### 3. **Executar User Service**
```bash
# Terminal 1
cd services/user
go run main.go

# Deve exibir: "User Service rodando na porta 8081"
```

### 4. **Executar API Gateway**
```bash
# Terminal 2
cd gateway
go run main.go

# Deve exibir: "🚀 API Gateway rodando na porta 8080"
```

### 5. **Testar APIs**
```bash
# Terminal 3 - Executar script de teste
./test_api.sh

# Ou testar manualmente:
curl http://localhost:8080/health
```

## 📊 Endpoints Disponíveis

### **Públicos (sem autenticação)**
```bash
POST /api/v1/auth/register   # Registrar usuário
POST /api/v1/auth/login      # Fazer login
GET  /health                 # Health check gateway
```

### **Protegidos (com Bearer token)**
```bash
GET    /api/v1/users/profile    # Buscar perfil
PUT    /api/v1/users/profile    # Atualizar perfil
DELETE /api/v1/users/profile    # Deletar conta

GET    /api/v1/contacts         # Listar contatos
POST   /api/v1/contacts         # Criar contato
PUT    /api/v1/contacts/:id     # Atualizar contato
DELETE /api/v1/contacts/:id     # Deletar contato
```

## 🧪 Exemplo de uso

### **1. Registrar usuário**
```bash
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "joao123",
    "email": "joao@email.com",
    "password": "123456",
    "full_name": "João Silva"
  }'
```

### **2. Fazer login**
```bash
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "email": "joao@email.com",
    "password": "123456"
  }'
```

### **3. Buscar perfil (com token)**
```bash
curl -H "Authorization: Bearer SEU_TOKEN_AQUI" \
  http://localhost:8080/api/v1/users/profile
```

## 🔧 Configurações

### **Variáveis de ambiente**
```bash
# Banco de dados
DB_HOST=localhost
DB_PORT=5432
DB_NAME=meuapoio
DB_USER=postgres
DB_PASSWORD=postgres123

# JWT
JWT_SECRET=sua-chave-secreta-super-segura

# Serviços
PORT=8080  # Gateway
PORT=8081  # User Service
```

### **Portas utilizadas**
- **8080**: API Gateway
- **8081**: User Service  
- **5432**: PostgreSQL
- **27017**: MongoDB
- **6379**: Redis

## 🎉 Status

| Componente | Status | Porta |
|------------|--------|-------|
| API Gateway | ✅ Funcionando | 8080 |
| User Service | ✅ Funcionando | 8081 |
| PostgreSQL | ✅ Funcionando | 5432 |
| MongoDB | ✅ Funcionando | 27017 |
| Redis | ✅ Funcionando | 6379 |

## 🚀 Próximos passos

1. **Audio Service** - Meditações e músicas
2. **Content Service** - Artigos e histórias  
3. **Contact Service** - Rede de apoio
4. **Notification Service** - Notificações

---

**🏥 Lembre-se**: Este é um projeto de apoio emocional. Em caso de emergência, procure ajuda profissional imediata!

**CVV**: 188 | **SAMU**: 192 | **Bombeiros**: 193 