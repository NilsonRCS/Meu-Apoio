# 📚 Documentação do Projeto MeuApoio

Bem-vindo à documentação completa do projeto **MeuApoio**! Este diretório contém toda a documentação técnica, arquitetural e operacional do sistema.

## 📑 Índice da Documentação

### 🚀 **Começando**
- **[onboarding.md](./onboarding.md)** - 🎯 **COMECE AQUI** - Guia completo para novos desenvolvedores
  - Setup do ambiente passo a passo
  - Primeira execução e troubleshooting
  - Padrões de código e primeira contribuição
  - Recursos de aprendizado

### 🏗️ **Arquitetura & Design**
- **[architecture.md](./architecture.md)** - Documentação completa da arquitetura
  - Decisões arquiteturais (ADRs)
  - Padrões aplicados
  - Trade-offs e justificativas
  - Diagramas do sistema
  - Roadmap técnico

### 🗄️ **Bancos de Dados**
- **[database.md](./database.md)** - Configuração e administração dos bancos
  - PostgreSQL setup e tabelas
  - MongoDB collections
  - Redis cache strategy
  - Como conectar e administrar

### 📡 **APIs & Serviços** (Em desenvolvimento)
- **api-gateway.md** - Documentação do API Gateway
- **user-service.md** - Serviço de usuários e autenticação
- **content-service.md** - Serviço de conteúdo e CMS
- **audio-service.md** - Serviço de áudio e streaming
- **contact-service.md** - Serviço de contatos de emergência
- **notification-service.md** - Serviço de notificações

### 🔧 **Desenvolvimento & Deploy** (Planejado)
- **development-guide.md** - Guia para desenvolvedores
- **deployment.md** - Deploy e CI/CD
- **testing.md** - Estratégias de teste
- **monitoring.md** - Observabilidade e métricas

### 📋 **Processo & Workflow** (Planejado)
- **contributing.md** - Como contribuir com o projeto
- **code-standards.md** - Padrões de código
- **git-workflow.md** - Workflow de Git e branching

---

## 🚀 Quick Start

### 👨‍💻 **Para Novos Desenvolvedores** 
**👉 [COMECE PELO ONBOARDING](./onboarding.md)**

1. Siga o [guia de onboarding](./onboarding.md) passo a passo
2. Configure o ambiente local
3. Execute: `docker compose up -d` para subir os bancos
4. Faça sua primeira contribuição!

### 🏗️ **Para Arquitetos/Leads**
- **[architecture.md](./architecture.md)** contém todas as decisões arquiteturais
- ADRs (Architecture Decision Records) estão documentados
- Trade-offs e justificativas técnicas estão explicados

### 🔧 **Para DevOps/SRE**
- **[database.md](./database.md)** para configurações de infraestrutura
- Futuramente: deployment.md e monitoring.md

---

## 🎯 Filosofia da Documentação

### Princípios que Seguimos:
- **Documentation as Code** - Versionada junto com o código
- **Living Documentation** - Atualizada a cada mudança significativa
- **ADR Pattern** - Decisões arquiteturais formalmente documentadas
- **Onboarding Friendly** - Fácil para novos desenvolvedores

### Quando Atualizar:
- ✅ Nova feature implementada
- ✅ Mudança arquitetural
- ✅ Novos padrões adotados
- ✅ Lições aprendidas importantes

---

## 🤝 Como Contribuir com a Documentação

### Para Atualizar Documentação:
1. Faça suas mudanças nos arquivos .md
2. Commit com: `docs: descrição da mudança`
3. Pull request para review

### Para Adicionar Nova Documentação:
1. Crie arquivo na pasta `docs/`
2. Adicione link neste README.md
3. Siga o padrão de estrutura existente

---

## 📞 Suporte & Contato

### 🆘 **Preciso de Ajuda!**
1. **Primeiro:** Consulte o [onboarding.md](./onboarding.md) - tem troubleshooting completo
2. **Arquitetura:** Consulte [architecture.md](./architecture.md)
3. **Setup:** Consulte [database.md](./database.md)
4. **Ainda com problema:** Abra issue no projeto

### Dúvidas sobre Desenvolvimento:
- Consulte os guias específicos de cada serviço
- Documentação das APIs (Swagger/OpenAPI em desenvolvimento)

### Problemas de Ambiente:
- **Containers não sobem:** Veja troubleshooting no [onboarding.md](./onboarding.md)
- **Erros de banco:** Verifique logs: `docker compose logs -f`

---

**Última Atualização:** 06/06/2024  
**Próxima Revisão:** A cada milestone  
**Responsável:** Time de Arquitetura

> 💡 **Dica:** Mantenha este README atualizado sempre que adicionar nova documentação! 