package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meuapoio/services/user/models"
	"github.com/meuapoio/services/user/repository"
)

type ContactHandler struct {
	contactRepo *repository.ContactRepository
}

func NewContactHandler(contactRepo *repository.ContactRepository) *ContactHandler {
	return &ContactHandler{contactRepo: contactRepo}
}

func (h *ContactHandler) GetContacts(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	contacts, err := h.contactRepo.GetByUserID(userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contatos"})
		return
	}

	if contacts == nil {
		contacts = []*models.EmergencyContact{}
	}

	c.JSON(http.StatusOK, gin.H{"contacts": contacts})
}

func (h *ContactHandler) CreateContact(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	var req models.CreateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	contact, err := h.contactRepo.Create(userID.(string), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar contato"})
		return
	}

	c.JSON(http.StatusCreated, contact)
}

func (h *ContactHandler) UpdateContact(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	contactID := c.Param("id")
	if contactID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do contato é obrigatório"})
		return
	}

	var req models.UpdateContactRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar se contato existe e pertence ao usuário
	_, err := h.contactRepo.GetByID(contactID, userID.(string))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contato não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	// Atualizar contato
	if err := h.contactRepo.Update(contactID, userID.(string), &req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar contato"})
		return
	}

	// Buscar contato atualizado
	updatedContact, err := h.contactRepo.GetByID(contactID, userID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar contato atualizado"})
		return
	}

	c.JSON(http.StatusOK, updatedContact)
}

func (h *ContactHandler) DeleteContact(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuário não autenticado"})
		return
	}

	contactID := c.Param("id")
	if contactID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID do contato é obrigatório"})
		return
	}

	// Verificar se contato existe e pertence ao usuário
	_, err := h.contactRepo.GetByID(contactID, userID.(string))
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Contato não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno do servidor"})
		return
	}

	// Deletar contato
	if err := h.contactRepo.Delete(contactID, userID.(string)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar contato"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Contato deletado com sucesso"})
}
