# Configuração dos Bancos de Dados

## PostgreSQL - Banco Principal

**Porta**: 5432  
**Usuário**: postgres  
**Senha**: postgres123  
**Database**: meuapoio  

### Tabelas Criadas:
- `users` - Usuários do sistema
- `emergency_contacts` - Contatos de emergência
- `audios` - Catálogo de áudios (meditação/música)
- `user_favorites` - Favoritos dos usuários
- `user_play_history` - Histórico de reprodução
- `notifications` - Sistema de notificações

### Acessar via Adminer:
- URL: http://localhost:8080
- Sistema: PostgreSQL
- Servidor: postgres
- Usuário: postgres
- Senha: postgres123
- Base de dados: meuapoio

## MongoDB - Conteúdo e Artigos

**Porta**: 27017  
**Usuário**: mongo  
**Senha**: mongo123  
**Database**: meuapoio  

### Collections Planejadas:
- `articles` - Artigos de apoio emocional
- `stories` - Histórias de superação
- `help_content` - Conteúdo de ajuda
- `categories` - Categorias de conteúdo

### Acessar via Mongo Express:
- URL: http://localhost:8081
- Usuário: mongo
- Senha: mongo123

## Redis - Cache

**Porta**: 6379  
**Sem autenticação para desenvolvimento**

### Uso Previsto:
- Cache de sessões
- Cache de consultas frequentes
- Rate limiting
- Dados temporários

## Como executar

```bash
# Subir apenas os bancos de dados
docker-compose up -d

# Ver logs
docker-compose logs -f

# Parar tudo
docker-compose down

# Parar e remover volumes (CUIDADO: apaga os dados)
docker-compose down -v
```

## Conectar aos bancos externamente

### PostgreSQL
```bash
psql -h localhost -p 5432 -U postgres -d meuapoio
```

### MongoDB
```bash
mongosh "mongodb://mongo:mongo123@localhost:27017/meuapoio"
```

### Redis
```bash
redis-cli -h localhost -p 6379
``` 