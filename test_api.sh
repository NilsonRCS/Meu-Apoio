#!/bin/bash

# Script de teste das APIs do MeuApoio
# Testa User Service e API Gateway

echo "🚀 Testando APIs do MeuApoio"
echo "=================================="

API_BASE="http://localhost:8080/api/v1"

echo ""
echo "1. 🏥 Health Check do Gateway"
echo "------------------------------"
curl -s "$API_BASE/../health" | jq '.'

echo ""
echo "2. 👤 Registrar novo usuário"
echo "------------------------------"
REGISTER_RESPONSE=$(curl -s -X POST "$API_BASE/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "teste_user",
    "email": "teste@meuapoio.com",
    "password": "123456",
    "full_name": "Usuário de Teste"
  }')

echo "$REGISTER_RESPONSE" | jq '.'

# Extrair token do registro
TOKEN=$(echo "$REGISTER_RESPONSE" | jq -r '.token // empty')

if [ -z "$TOKEN" ]; then
  echo ""
  echo "🔑 Fazendo login para obter token..."
  LOGIN_RESPONSE=$(curl -s -X POST "$API_BASE/auth/login" \
    -H "Content-Type: application/json" \
    -d '{
      "email": "teste@meuapoio.com",
      "password": "123456"
    }')
  
  echo "$LOGIN_RESPONSE" | jq '.'
  TOKEN=$(echo "$LOGIN_RESPONSE" | jq -r '.token // empty')
fi

if [ -n "$TOKEN" ]; then
  echo ""
  echo "✅ Token obtido: ${TOKEN:0:20}..."
  
  echo ""
  echo "3. 👤 Buscar perfil do usuário"
  echo "------------------------------"
  curl -s -H "Authorization: Bearer $TOKEN" "$API_BASE/users/profile" | jq '.'

  echo ""
  echo "4. 📞 Criar contato de emergência"
  echo "--------------------------------"
  curl -s -X POST "$API_BASE/contacts" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
      "name": "João Silva",
      "phone": "(11) 99999-9999",
      "relationship": "Irmão",
      "is_primary": true
    }' | jq '.'

  echo ""
  echo "5. 📞 Listar contatos"
  echo "--------------------"
  curl -s -H "Authorization: Bearer $TOKEN" "$API_BASE/contacts" | jq '.'

  echo ""
  echo "6. 👤 Atualizar perfil"
  echo "---------------------"
  curl -s -X PUT "$API_BASE/users/profile" \
    -H "Authorization: Bearer $TOKEN" \
    -H "Content-Type: application/json" \
    -d '{
      "full_name": "Usuário de Teste Atualizado",
      "phone": "(11) 88888-8888"
    }' | jq '.'

else
  echo "❌ Erro: Não foi possível obter o token de autenticação"
fi

echo ""
echo "🎉 Testes concluídos!" 