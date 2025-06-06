# Arquitetura do Sistema MeuApoio

## 📋 Índice
- [Visão Geral](#visão-geral)
- [Decisões Arquiteturais](#decisões-arquiteturais)
- [Estrutura de Microserviços](#estrutura-de-microserviços)
- [Estratégia de Dados](#estratégia-de-dados)
- [Padrões Aplicados](#padrões-aplicados)
- [Trade-offs e Justificativas](#trade-offs-e-justificativas)
- [Roadmap Técnico](#roadmap-técnico)
- [Diagramas](#diagramas)

---

## 🎯 Visão Geral

O **MeuApoio** é um aplicativo de apoio emocional construído com arquitetura de microserviços em Go. O sistema oferece funcionalidades de meditação, música relaxante, conteúdo de ajuda, histórias inspiradoras e contatos de emergência.

### Domínios Principais
- **Autenticação & Usuários** - Gestão de contas e perfis
- **Conteúdo Editorial** - Artigos, histórias e material de apoio  
- **Áudio & Mídia** - Meditações guiadas e música relaxante
- **Contatos de Apoio** - Rede de suporte e emergência
- **Notificações** - Lembretes e comunicação com usuários

---

## 🏗️ Decisões Arquiteturais

### ADR-001: Arquitetura de Microserviços

**Status:** ✅ Aprovado  
**Data:** 06/06/2024  

**Contexto:**
Sistema com múltiplos domínios distintos, necessidade de escalabilidade independente e equipes trabalhando em paralelo.

**Decisão:**
Adotar arquitetura de microserviços com separação por domínio de negócio.

**Justificativa:**
- **Independência de Deploy:** Cada serviço pode ser atualizado sem afetar outros
- **Escalabilidade Granular:** Serviço de áudio pode escalar independentemente
- **Technology Diversity:** Diferentes tecnologias para diferentes problemas
- **Team Autonomy:** Equipes podem trabalhar independentemente
- **Fault Isolation:** Falha em um serviço não derruba o sistema todo

**Consequências:**
- ✅ Maior flexibilidade e escalabilidade
- ✅ Deploy independente e continuous delivery
- ❌ Maior complexidade operacional
- ❌ Overhead de comunicação entre serviços

### ADR-002: Database per Service Pattern

**Status:** ✅ Aprovado  
**Data:** 06/06/2024  

**Contexto:**
Necessidade de autonomia de dados e diferentes características de acesso.

**Decisão:**
Cada microserviço possui seu próprio banco de dados otimizado para seu domínio.

**Justificativa:**
- **PostgreSQL** para dados estruturados (users, contacts, audio metadata)
- **MongoDB** para conteúdo flexível (articles, stories, help content)
- **Redis** para caching e sessões temporárias

**Consequências:**
- ✅ Otimização específica por domínio
- ✅ Isolamento de falhas
- ❌ Eventual consistency entre serviços
- ❌ Joins cross-service não são possíveis

### ADR-003: Shared Kernel Pattern

**Status:** ✅ Aprovado  
**Data:** 06/06/2024  

**Contexto:**
Evitar duplicação de código comum entre microserviços.

**Decisão:**
Criar pasta `shared/` com código reutilizável (config, database, middleware, utils).

**Justificativa:**
- DRY (Don't Repeat Yourself)
- Consistência de implementação
- Facilita manutenção e updates

---

## 🔧 Estrutura de Microserviços

```
📦 MeuApoio Ecosystem
├── 🌐 API Gateway (Port 8080)
│   ├── Routing & Load Balancing
│   ├── Authentication & Authorization  
│   ├── Rate Limiting
│   └── Request/Response Transformation
│
├── 👤 User Service (Port 8081)
│   ├── User Registration & Authentication
│   ├── Profile Management
│   ├── JWT Token Management
│   └── User Preferences
│
├── 📝 Content Service (Port 8082)
│   ├── Articles & Stories Management
│   ├── Help Content CMS
│   ├── Content Categories
│   └── Search & Filtering
│
├── 🎵 Audio Service (Port 8083)
│   ├── Audio Upload & Processing
│   ├── Streaming & Download
│   ├── Playlist Management
│   └── Favorites & History
│
├── 📞 Contact Service (Port 8084)
│   ├── Emergency Contacts
│   ├── Support Network
│   ├── Professional Contacts
│   └── Contact Validation
│
└── 🔔 Notification Service (Port 8085)
    ├── Push Notifications
    ├── Email Notifications
    ├── SMS Alerts
    └── Scheduled Reminders
```

---

## 🗄️ Estratégia de Dados

### PostgreSQL - Banco Relacional Principal
**Porta:** 5432 | **Admin UI:** Adminer (http://localhost:8080)

```sql
-- Principais Entidades
Users            // Usuários do sistema
Emergency_Contacts // Contatos de emergência  
Audios           // Metadados dos áudios
User_Favorites   // Favoritos dos usuários
User_Play_History // Histórico de reprodução
Notifications    // Sistema de notificações
```

**Características:**
- ACID Transactions
- Strong Consistency  
- UUID Primary Keys (distributed-friendly)
- Timestamps automáticos para auditoria
- Soft delete pattern (`is_active`)
- Índices otimizados para consultas frequentes

### MongoDB - Banco NoSQL para Conteúdo
**Porta:** 27017 | **Admin UI:** Mongo Express (http://localhost:8081)

```javascript
// Collections Planejadas
articles: {        // Artigos de apoio emocional
  title, content, author, tags, category, publishedAt
}
stories: {         // Histórias de superação
  title, story, author, emotions, inspiration
}  
help_content: {    // Conteúdo de ajuda
  title, content, category, difficulty, steps
}
categories: {      // Categorias de conteúdo
  name, description, icon, color
}
```

**Características:**
- Schema Flexível
- Document-oriented
- Rich queries & aggregations
- Horizontal scaling ready

### Redis - Cache & Sessions
**Porta:** 6379 | **Sem Auth (Development)**

**Uso Cases:**
- Session storage (JWT tokens)
- Rate limiting counters
- Cached queries (frequent reads)
- Temporary data (OTP codes)
- Real-time data (active users)

---

## 🎨 Padrões Aplicados

### 1. **Domain-Driven Design (DDD)**
- Estrutura organizada por domínios de negócio
- Bounded contexts bem definidos
- Ubiquitous language entre dev e business

### 2. **API Gateway Pattern**
```
Client Apps → API Gateway → Microservices
```
- Single entry point
- Cross-cutting concerns (auth, logging, rate limiting)
- Service discovery & load balancing

### 3. **Database per Service**
- Data autonomy
- Technology choice per domain
- Fault isolation

### 4. **Event-Driven Architecture** (Futuro)
- Async communication via message queues
- Event sourcing for audit trails
- CQRS for read/write optimization

### 5. **Circuit Breaker Pattern** (Futuro)
- Fault tolerance between services
- Graceful degradation
- Auto-recovery mechanisms

---

## ⚖️ Trade-offs e Justificativas

### Por que Microserviços? (vs Monolito)

| Aspecto | Monolito | Microserviços | Nossa Escolha |
|---------|----------|---------------|---------------|
| **Complexidade** | Baixa | Alta | ✅ Aceitável (team expertise) |
| **Deploy** | Tudo junto | Independente | ✅ Critical para CI/CD |
| **Escalabilidade** | Vertical | Horizontal granular | ✅ Audio service needs more resources |
| **Technology Stack** | Homogêneo | Heterogêneo | ✅ Right tool for right job |
| **Team Structure** | Central | Distributed | ✅ Multiple teams working parallel |
| **Data Consistency** | ACID | Eventual | ⚠️ Acceptable for our use case |

### Por que Go? (vs outras linguagens)

- **Performance:** Concurrency nativa, baixa latência
- **Ecosystem:** Rico em bibliotecas para microserviços  
- **Deploy:** Single binary, containers lightweight
- **Learning Curve:** Sintaxe simples, produtividade alta
- **Community:** Grande adoção em cloud-native apps

### Por que Docker Compose? (vs Kubernetes)

**Current Stage:** Development & Prototyping
- ✅ Simplicidade para desenvolvimento local
- ✅ Easy setup para novos developers
- ✅ Resource efficient (development machines)

**Future Migration:** Production → Kubernetes
- Quando escala passa de ~10 containers
- Quando precisar de auto-scaling
- Quando deployment frequency aumentar

---

## 🛣️ Roadmap Técnico

### 🏗️ **Fase 1: Foundations** ✅
- [x] Project structure & documentation
- [x] Database design & setup  
- [x] Docker orchestration
- [x] Shared libraries (config, middleware)

### 🔐 **Fase 2: Core Authentication**
- [ ] User Service implementation
- [ ] JWT authentication flow
- [ ] API Gateway setup
- [ ] Basic CRUD operations

### 🎵 **Fase 3: Business Logic**  
- [ ] Audio Service (upload, streaming)
- [ ] Content Service (CMS functionality)
- [ ] Contact Service (emergency contacts)

### 🔔 **Fase 4: Advanced Features**
- [ ] Notification Service (push, email, SMS)
- [ ] Real-time features (WebSocket)
- [ ] Analytics & monitoring

### 🚀 **Fase 5: Production Ready**
- [ ] Kubernetes migration
- [ ] CI/CD pipelines  
- [ ] Monitoring & observability (Prometheus/Grafana)
- [ ] Security hardening
- [ ] Performance optimization

---

## 📊 Diagramas

### System Architecture Overview
```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Mobile App    │    │    Web App       │    │   Admin Panel   │
└─────────┬───────┘    └────────┬─────────┘    └─────────┬───────┘
          │                     │                        │
          └─────────────────────┼────────────────────────┘
                                │
                    ┌───────────▼────────────┐
                    │     API Gateway        │
                    │   (Port 8080)          │
                    │ • Auth & Rate Limiting │
                    │ • Request Routing      │
                    └───────────┬────────────┘
                                │
            ┌───────────────────┼───────────────────┐
            │                   │                   │
    ┌───────▼────────┐ ┌────────▼────────┐ ┌───────▼────────┐
    │  User Service  │ │ Content Service │ │  Audio Service │
    │  (Port 8081)   │ │  (Port 8082)    │ │  (Port 8083)   │
    │                │ │                 │ │                │
    │ ┌─PostgreSQL─┐ │ │ ┌─MongoDB────┐  │ │ ┌─PostgreSQL─┐ │
    │ │• Users     │ │ │ │• Articles  │  │ │ │• Audios    │ │
    │ │• Contacts  │ │ │ │• Stories   │  │ │ │• Favorites │ │
    │ └────────────┘ │ │ │• Help      │  │ │ │• History   │ │
    └────────────────┘ │ └────────────┘  │ │ └────────────┘ │
                       └─────────────────┘ └────────────────┘
                                │
                    ┌───────────▼────────────┐
                    │       Redis            │
                    │   (Cache & Sessions)   │
                    │     Port 6379          │
                    └────────────────────────┘
```

### Data Flow Example: User Authentication
```
1. Mobile App → API Gateway: POST /login {email, password}
2. API Gateway → User Service: Forward request
3. User Service → PostgreSQL: Validate credentials
4. User Service → Redis: Store session
5. User Service → API Gateway: Return JWT token
6. API Gateway → Mobile App: Return authenticated response
```

---

## 📚 Referências & Recursos

### Livros & Artigos
- **"Building Microservices"** - Sam Newman
- **"Microservices Patterns"** - Chris Richardson  
- **"Clean Architecture"** - Robert C. Martin
- **Martin Fowler on Microservices** - martinfowler.com

### Tools & Technologies  
- **Go:** https://golang.org/
- **Gin Framework:** https://gin-gonic.com/
- **PostgreSQL:** https://postgresql.org/
- **MongoDB:** https://mongodb.com/
- **Redis:** https://redis.io/
- **Docker:** https://docker.com/

### Monitoring & Observability (Future)
- **Prometheus:** Metrics collection
- **Grafana:** Metrics visualization  
- **Jaeger:** Distributed tracing
- **ELK Stack:** Centralized logging

---

## 👥 Team Guidelines

### Code Standards
- **Go:** Follow effective Go guidelines
- **Git:** Conventional commits pattern
- **Documentation:** Code comments + architectural decisions
- **Testing:** Unit tests + integration tests

### Development Workflow
1. Feature branch from `main`
2. Implement with tests
3. Update documentation if needed
4. Pull request with peer review
5. Merge after CI/CD passes

---

**Última Atualização:** 06/06/2024  
**Próxima Revisão:** A cada milestone completado  
**Maintainers:** Time de Arquitetura 