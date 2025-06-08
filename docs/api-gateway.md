# 🌐 API Gateway - Documentação Técnica

## 📋 Índice

- [Visão Geral](#visão-geral)
- [Por que API Gateway?](#por-que-api-gateway)
- [Arquitetura](#arquitetura)
- [Funcionalidades Implementadas](#funcionalidades-implementadas)
- [Configuração e Execução](#configuração-e-execução)
- [Middleware](#middleware)
- [Proxy Reverso](#proxy-reverso)
- [Roteamento](#roteamento)
- [Autenticação Centralizada](#autenticação-centralizada)
- [Rate Limiting](#rate-limiting)
- [CORS](#cors)
- [Health Checks](#health-checks)
- [Logs e Monitoramento](#logs-e-monitoramento)
- [Exemplos de Uso](#exemplos-de-uso)
- [Troubleshooting](#troubleshooting)
- [Próximos Passos](#próximos-passos)

---

## 🎯 Visão Geral

O **API Gateway** é o ponto de entrada único para todos os clientes do sistema MeuApoio. Ele atua como um **proxy reverso inteligente** que roteia requisições para os microserviços apropriados, gerencia autenticação, implementa rate limiting e fornece uma camada de abstração entre clientes e serviços.

### **Características Principais:**
- ✅ **Ponto único de entrada** para todas as APIs
- ✅ **Proxy reverso** para microserviços
- ✅ **Autenticação centralizada** via JWT
- ✅ **Rate limiting** por IP
- ✅ **CORS** configurado
- ✅ **Health checks** integrados
- ✅ **Logs detalhados** de requisições

### **Tecnologias Utilizadas:**
- **Go 1.21+** com framework **Gin**
- **JWT** para autenticação
- **Token Bucket** para rate limiting
- **Docker** para containerização

---

## 🤔 Por que API Gateway?

### **Problemas Resolvidos:**

#### **1. Complexidade do Cliente**
**❌ Sem Gateway:**
```javascript
// Cliente precisa conhecer todos os serviços
const userAPI = 'http://user-service:8081'
const audioAPI = 'http://audio-service:8082'
const contentAPI = 'http://content-service:8083'
```

**✅ Com Gateway:**
```javascript
// Cliente conhece apenas o Gateway
const API_BASE = 'http://api-gateway:8080'
```

#### **2. Autenticação Duplicada**
**❌ Sem Gateway:**
- Cada serviço implementa validação JWT
- Código duplicado e inconsistente
- Múltiplos pontos de falha

**✅ Com Gateway:**
- Autenticação centralizada
- Token validado uma vez
- Propagação de contexto do usuário

#### **3. CORS e Políticas de Segurança**
**❌ Sem Gateway:**
- Cada serviço configura CORS
- Políticas inconsistentes
- Headers duplicados

**✅ Com Gateway:**
- CORS configurado uma vez
- Políticas uniformes
- Headers padronizados

#### **4. Rate Limiting Distribuído**
**❌ Sem Gateway:**
- Difícil implementar rate limiting global
- Cada serviço tem seu próprio limite
- Possível abuso por distribuição

**✅ Com Gateway:**
- Rate limiting global por IP
- Proteção contra DDoS
- Controle centralizado

---

## 🏗️ Arquitetura

### **Diagrama de Fluxo:**

```
┌─────────────┐    ┌─────────────────┐    ┌──────────────────┐
│   Cliente   │────│   API Gateway   │────│  User Service    │
│  (Frontend) │    │   (Port 8080)   │    │   (Port 8081)    │
└─────────────┘    └─────────────────┘    └──────────────────┘
                          │
                          ├─────────────────────┬──────────────────────┐
                          │                     │                      │
                   ┌──────▼──────┐      ┌──────▼──────┐      ┌──────▼──────┐
                   │Audio Service│      │Content Serv.│      │Contact Serv.│
                   │(Port 8082)  │      │(Port 8083)  │      │(Port 8084)  │
                   └─────────────┘      └─────────────┘      └─────────────┘
```

### **Camadas do Gateway:**

```
┌─────────────────────────────────────────────────────────────┐
│                     CLIENT REQUEST                          │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                   MIDDLEWARE STACK                          │
│  1. Logger          │ 2. Recovery     │ 3. CORS            │
│  4. Rate Limiter    │ 5. Auth (JWT)   │ 6. Request Transform│
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                    ROUTING LAYER                            │
│  • Path matching    • Service discovery • Load balancing   │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                   PROXY LAYER                               │
│  • Reverse proxy    • Header management  • Error handling  │
└─────────────────────────┬───────────────────────────────────┘
                          │
┌─────────────────────────▼───────────────────────────────────┐
│                 TARGET MICROSERVICE                         │
└─────────────────────────────────────────────────────────────┘
```

---

## ⚙️ Funcionalidades Implementadas

### **1. Service Registry**
Registro centralizado dos microserviços:

```go
type ServiceRegistry struct {
    UserService    string // http://localhost:8081
    AudioService   string // http://localhost:8082 (futuro)
    ContentService string // http://localhost:8083 (futuro)
    ContactService string // http://localhost:8084 (futuro)
}
```

### **2. Roteamento Inteligente**
Mapeamento de rotas para serviços:

```go
// Rotas públicas (sem autenticação)
userGroup.POST("/auth/register", proxyToService(services.UserService))
userGroup.POST("/auth/login", proxyToService(services.UserService))

// Rotas protegidas (com autenticação JWT)
protected.GET("/users/profile", proxyToService(services.UserService))
protected.GET("/contacts", proxyToService(services.UserService))
```

### **3. Middleware Pipeline**
Processamento sequencial de middlewares:

1. **Logger**: Log de todas as requisições
2. **Recovery**: Captura panics e retorna erro 500
3. **CORS**: Headers para cross-origin requests
4. **Rate Limiter**: Controle de taxa por IP
5. **Auth**: Validação JWT (apenas rotas protegidas)

---

## 🚀 Configuração e Execução

### **Variáveis de Ambiente:**

```bash
# Porta do Gateway
PORT=8080

# Ambiente (development/production)
ENVIRONMENT=development

# JWT Secret para validação de tokens
JWT_SECRET=sua-chave-secreta-super-segura

# URLs dos microserviços
USER_SERVICE_URL=http://localhost:8081
AUDIO_SERVICE_URL=http://localhost:8082
CONTENT_SERVICE_URL=http://localhost:8083
```

### **Como Executar:**

```bash
# Desenvolvimento
cd gateway
PORT=8080 go run main.go

# Produção (com Docker)
docker build -t meuapoio-gateway .
docker run -p 8080:8080 -e JWT_SECRET=secret meuapoio-gateway
```

---

## 🛡️ Middleware

### **1. Rate Limiter**

**Implementação:** Token Bucket Algorithm

```go
type RateLimiter struct {
    visitors map[string]*Visitor // Cache de visitantes por IP
    mutex    sync.RWMutex       // Proteção concorrente
    rate     int                // Tokens por período
    duration time.Duration      // Período de refill
}
```

**Características:**
- ✅ **100 requests/minuto** por IP
- ✅ **Token bucket** com refill automático
- ✅ **Cleanup automático** de visitantes antigos
- ✅ **Thread-safe** com mutex

**Funcionamento:**
1. Cliente faz requisição
2. Gateway identifica IP do cliente
3. Verifica tokens disponíveis no bucket
4. Se disponível: consome token e prossegue
5. Se não: retorna `429 Too Many Requests`

### **2. Auth Middleware**

**Fluxo de Autenticação:**

```
┌─────────────┐    ┌──────────────┐    ┌─────────────────┐
│   Request   │───▶│Extract Token │───▶│  Validate JWT   │
└─────────────┘    └──────────────┘    └─────────────────┘
                           │                      │
                           ▼                      ▼
                   ┌──────────────┐    ┌─────────────────┐
                   │ Return 401   │    │ Set User Context│
                   └──────────────┘    └─────────────────┘
                                              │
                                              ▼
                                    ┌─────────────────┐
                                    │  Continue to    │
                                    │   Next Handler  │
                                    └─────────────────┘
```

**Implementação:**
- ✅ Extrai token do header `Authorization: Bearer <token>`
- ✅ Valida JWT usando chave secreta
- ✅ Adiciona `user_id` e `user_email` ao contexto
- ✅ Retorna `401` se token inválido ou ausente

### **3. CORS Middleware**

**Configuração:**

```go
cors.Config{
    AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}
```

---

## 🔄 Proxy Reverso

### **Implementação:**

O Gateway utiliza `httputil.ReverseProxy` do Go para redirecionamento:

```go
func createReverseProxy(target string) http.Handler {
    url, _ := url.Parse(target)
    proxy := httputil.NewSingleHostReverseProxy(url)
    
    // Customizar headers
    originalDirector := proxy.Director
    proxy.Director = func(req *http.Request) {
        originalDirector(req)
        req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
        req.Header.Set("X-Origin-Service", "api-gateway")
    }
    
    // Error handling personalizado
    proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
        log.Printf("Erro no proxy para %s: %v", target, err)
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusBadGateway)
        w.Write([]byte(`{"error": "Serviço temporariamente indisponível"}`))
    }
    
    return proxy
}
```

### **Headers Adicionados:**

| **Header** | **Valor** | **Propósito** |
|------------|-----------|---------------|
| `X-Forwarded-Host` | Host original | Preservar host do cliente |
| `X-Origin-Service` | `api-gateway` | Identificar origem da requisição |
| `User-ID` | ID do usuário JWT | Contexto do usuário autenticado |
| `User-Email` | Email do usuário | Contexto do usuário autenticado |

### **Error Handling:**

- ✅ **502 Bad Gateway**: Quando serviço está indisponível
- ✅ **503 Service Unavailable**: Quando serviço está sobrecarregado
- ✅ **Logs detalhados**: Para debugging e monitoramento

---

## 🛣️ Roteamento

### **Mapeamento de Rotas:**

| **Padrão** | **Serviço de Destino** | **Descrição** |
|------------|-------------------------|---------------|
| `/health` | Gateway interno | Health check do gateway |
| `/api/v1/auth/*` | User Service | Autenticação e registro |
| `/api/v1/users/*` | User Service | Gestão de usuários |
| `/api/v1/contacts/*` | User Service | Contatos de emergência |
| `/api/v1/audio/*` | Audio Service | Meditações e músicas (futuro) |
| `/api/v1/content/*` | Content Service | Artigos e histórias (futuro) |

### **Tipos de Rotas:**

#### **Públicas (sem autenticação):**
```bash
POST /api/v1/auth/register   # Registro de usuário
POST /api/v1/auth/login      # Login
GET  /health                 # Health check
```

#### **Protegidas (requer JWT):**
```bash
GET    /api/v1/users/profile    # Buscar perfil
PUT    /api/v1/users/profile    # Atualizar perfil
DELETE /api/v1/users/profile    # Deletar conta
GET    /api/v1/contacts         # Listar contatos
POST   /api/v1/contacts         # Criar contato
PUT    /api/v1/contacts/:id     # Atualizar contato
DELETE /api/v1/contacts/:id     # Deletar contato
```

---

## 🔐 Autenticação Centralizada

### **Fluxo de Autenticação:**

```mermaid
sequenceDiagram
    participant C as Cliente
    participant G as Gateway
    participant U as User Service
    
    C->>G: POST /api/v1/auth/login
    G->>U: Proxy para User Service
    U->>U: Valida credenciais
    U->>U: Gera JWT
    U->>G: Retorna JWT + dados do usuário
    G->>C: Retorna JWT + dados do usuário
    
    C->>G: GET /api/v1/users/profile (com JWT)
    G->>G: Valida JWT
    G->>G: Extrai user_id
    G->>U: Proxy com contexto do usuário
    U->>G: Retorna dados do perfil
    G->>C: Retorna dados do perfil
```

### **Validação de Token:**

1. **Extração**: Token extraído do header `Authorization`
2. **Formato**: Valida formato `Bearer <token>`
3. **Assinatura**: Verifica assinatura JWT com chave secreta
4. **Expiração**: Valida se token não expirou
5. **Contexto**: Adiciona `user_id` e `user_email` à requisição

### **Propagação de Contexto:**

O Gateway adiciona headers à requisição para o microserviço:

```http
GET /api/v1/users/profile HTTP/1.1
Host: user-service:8081
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
X-User-ID: d81a2c11-b489-4b1c-9d2a-c3d02ba2afa7
X-User-Email: usuario@example.com
X-Origin-Service: api-gateway
```

---

## ⏱️ Rate Limiting

### **Algoritmo Token Bucket:**

```go
type TokenBucket struct {
    tokens     int           // Tokens disponíveis
    capacity   int           // Capacidade máxima
    refillRate time.Duration // Taxa de refill
    lastRefill time.Time     // Último refill
    mutex      sync.Mutex    // Thread safety
}
```

### **Configuração Atual:**
- ✅ **100 requests por minuto** por IP
- ✅ **Refill automático** a cada 600ms
- ✅ **Cleanup automático** de visitantes inativos (>1 hora)

### **Funcionamento:**

1. **Primeira requisição**: Cria bucket com 100 tokens
2. **Requisições subsequentes**: Consome 1 token por requisição
3. **Refill**: Adiciona tokens baseado no tempo decorrido
4. **Esgotamento**: Retorna `429 Too Many Requests`
5. **Cleanup**: Remove visitantes antigos automaticamente

### **Resposta de Rate Limit:**

```json
{
  "error": "Rate limit exceeded",
  "retry_after": "60s"
}
```

---

## 🌐 CORS

### **Configuração Implementada:**

```go
cors.Config{
    AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
    AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
    AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
    ExposeHeaders:    []string{"Content-Length"},
    AllowCredentials: true,
    MaxAge:           12 * time.Hour,
}
```

### **Explicação dos Parâmetros:**

- **AllowOrigins**: Domínios permitidos (desenvolvimento local)
- **AllowMethods**: HTTP methods permitidos
- **AllowHeaders**: Headers permitidos nas requisições
- **ExposeHeaders**: Headers expostos ao cliente
- **AllowCredentials**: Permite cookies/credenciais
- **MaxAge**: Cache de preflight requests (12 horas)

---

## 🏥 Health Checks

### **Endpoint Interno:**

```bash
GET /health
```

**Resposta:**
```json
{
  "status": "ok",
  "service": "api-gateway",
  "timestamp": 1749407301
}
```

### **Monitoramento de Serviços:**

O Gateway pode ser estendido para incluir health checks dos microserviços:

```go
// Futuro: Health check agregado
func (g *Gateway) AggregatedHealthCheck() HealthStatus {
    checks := map[string]bool{
        "user-service":    g.checkService(g.registry.UserService),
        "audio-service":   g.checkService(g.registry.AudioService),
        "content-service": g.checkService(g.registry.ContentService),
    }
    
    return HealthStatus{
        Status:   calculateOverallStatus(checks),
        Services: checks,
        Timestamp: time.Now().Unix(),
    }
}
```

---

## 📊 Logs e Monitoramento

### **Logs Atuais:**

O Gateway registra automaticamente:

```bash
[GIN] 2025/06/08 - 15:28:43 | 201 | 69.660555ms | ::1 | POST "/api/v1/auth/register"
[GIN] 2025/06/08 - 15:29:10 | 200 |  1.232486ms | ::1 | GET  "/api/v1/users/profile"
[GIN] 2025/06/08 - 15:29:39 | 401 |    35.511µs | ::1 | GET  "/api/v1/users/profile"
```

**Informações incluídas:**
- ✅ **Timestamp** da requisição
- ✅ **Status Code** da resposta
- ✅ **Latência** da requisição
- ✅ **IP do cliente**
- ✅ **Método HTTP** e **Path**

### **Logs de Erro:**

```bash
2025/06/08 15:30:45 Erro no proxy para http://localhost:8081: dial tcp: connection refused
```

### **Métricas Recomendadas (Futuro):**

```go
type Metrics struct {
    RequestCount    int64         // Total de requests
    RequestLatency  time.Duration // Latência média
    ErrorRate       float64       // Taxa de erro
    RateLimitHits   int64         // Hits de rate limit
    ActiveUsers     int64         // Usuários ativos
}
```

---

## 🧪 Exemplos de Uso

### **1. Registro de Usuário:**

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

**Fluxo interno:**
1. Request chega no Gateway `:8080`
2. Passa pelos middlewares (CORS, Rate Limit, Logger)
3. Rota `/auth/register` → Proxy para User Service `:8081`
4. User Service processa e retorna JWT
5. Gateway retorna resposta ao cliente

### **2. Buscar Perfil (Autenticado):**

```bash
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..." \
     http://localhost:8080/api/v1/users/profile
```

**Fluxo interno:**
1. Request chega no Gateway
2. Passa por middlewares básicos
3. **Auth Middleware** valida JWT
4. Adiciona `user_id` e `user_email` ao contexto
5. Proxy para User Service com headers de contexto
6. User Service usa contexto para buscar dados
7. Gateway retorna resposta

### **3. Rate Limiting:**

```bash
# Fazer 101 requests rapidamente
for i in {1..101}; do
  curl http://localhost:8080/health
done

# A partir da 101ª requisição:
# HTTP 429 Too Many Requests
# {"error": "Rate limit exceeded", "retry_after": "60s"}
```

---

## 🔧 Troubleshooting

### **Problemas Comuns:**

#### **1. Gateway não inicia (porta ocupada):**

```bash
Error: listen tcp :8080: bind: address already in use
```

**Solução:**
```bash
# Verificar processo usando a porta
lsof -i :8080

# Parar processo ou usar porta diferente
PORT=8081 go run main.go
```

#### **2. Serviço downstream indisponível:**

```bash
Erro no proxy para http://localhost:8081: dial tcp: connection refused
```

**Solução:**
1. Verificar se User Service está rodando
2. Verificar conectividade de rede
3. Verificar configuração de Service Registry

#### **3. JWT inválido:**

```json
{"error": "Token inválido"}
```

**Causas possíveis:**
- Token expirado
- Chave JWT diferente entre Gateway e User Service
- Token malformado
- Header `Authorization` ausente

#### **4. CORS errors no browser:**

```
Access to fetch at 'http://localhost:8080/api/v1/users/profile' from origin 'http://localhost:3000' has been blocked by CORS policy
```

**Solução:**
1. Verificar se origem está em `AllowOrigins`
2. Verificar se headers estão em `AllowHeaders`
3. Verificar se método está em `AllowMethods`

### **Debug Commands:**

```bash
# Verificar health do Gateway
curl http://localhost:8080/health

# Verificar conectividade com User Service
curl http://localhost:8081/api/v1/health

# Testar autenticação
curl -H "Authorization: Bearer INVALID_TOKEN" \
     http://localhost:8080/api/v1/users/profile

# Verificar rate limiting
for i in {1..5}; do curl -w "%{http_code}\n" http://localhost:8080/health; done
```

---

## 🚀 Próximos Passos

### **Melhorias Planejadas:**

#### **1. Service Discovery:**
```go
type ServiceRegistry interface {
    RegisterService(name, url string) error
    DiscoverServices() map[string]string
    HealthCheck(service string) bool
}
```

#### **2. Load Balancing:**
```go
type LoadBalancer interface {
    SelectInstance(service string) string
    UpdateInstances(service string, instances []string)
}
```

#### **3. Circuit Breaker:**
```go
type CircuitBreaker struct {
    maxFailures   int
    timeout       time.Duration
    state         State // Open, Closed, HalfOpen
}
```

#### **4. Métricas e Observabilidade:**
```go
// Prometheus metrics
var (
    requestsTotal = prometheus.NewCounterVec(...)
    requestDuration = prometheus.NewHistogramVec(...)
    activeConnections = prometheus.NewGauge(...)
)
```

#### **5. Cache de Responses:**
```go
type ResponseCache struct {
    store map[string]CachedResponse
    ttl   time.Duration
}
```

#### **6. API Versioning:**
```go
// Suporte a múltiplas versões
/api/v1/users/profile → UserService v1
/api/v2/users/profile → UserService v2
```

#### **7. Request/Response Transformation:**
```go
type Transformer interface {
    TransformRequest(req *http.Request) error
    TransformResponse(resp *http.Response) error
}
```

### **Funcionalidades para Próximos Microserviços:**

- **Audio Service**: Streaming, upload, transcoding
- **Content Service**: CMS, busca, recomendações  
- **Contact Service**: Rede social, matching
- **Notification Service**: Push, email, SMS
- **Analytics Service**: Tracking, relatórios

---

## 📚 Referências

### **Documentação Técnica:**
- [Go HTTP ReverseProxy](https://pkg.go.dev/net/http/httputil#ReverseProxy)
- [Gin Framework](https://gin-gonic.com/docs/)
- [JWT Go Library](https://github.com/golang-jwt/jwt)

### **Padrões de Arquitetura:**
- [Microservices Patterns](https://microservices.io/patterns/)
- [API Gateway Pattern](https://microservices.io/patterns/apigateway.html)
- [Circuit Breaker Pattern](https://martinfowler.com/bliki/CircuitBreaker.html)

### **Melhores Práticas:**
- [12 Factor App](https://12factor.net/)
- [REST API Design](https://restfulapi.net/)
- [HTTP Status Codes](https://httpstatuses.com/)

---

**📧 Contato:** Para dúvidas técnicas sobre o API Gateway, consulte a equipe de arquitetura.

**🔄 Última atualização:** Junho 2025 