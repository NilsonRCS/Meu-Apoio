package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/meuapoio/services/user/handlers"
	"github.com/meuapoio/services/user/repository"
	"github.com/meuapoio/shared/config"
	"github.com/meuapoio/shared/database"
)

func main() {
	// Carregar configurações
	cfg := config.Load()

	// Conectar ao banco de dados
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("Falha ao conectar ao banco de dados:", err)
	}
	defer db.Close()

	// Inicializar repositórios
	userRepo := repository.NewUserRepository(db)
	contactRepo := repository.NewContactRepository(db)

	// Inicializar handlers
	userHandler := handlers.NewUserHandler(userRepo)
	contactHandler := handlers.NewContactHandler(contactRepo)
	authHandler := handlers.NewAuthHandler(userRepo, cfg.JWTSecret)

	// Configurar Gin
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Middlewares globais
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Rotas públicas (sem autenticação)
	public := r.Group("/api/v1")
	{
		public.POST("/auth/register", authHandler.Register)
		public.POST("/auth/login", authHandler.Login)
		public.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{"status": "ok", "service": "user-service"})
		})
	}

	// Rotas protegidas (com autenticação)
	protected := r.Group("/api/v1")
	protected.Use(authHandler.AuthMiddleware())
	{
		// Usuários
		protected.GET("/users/profile", userHandler.GetProfile)
		protected.PUT("/users/profile", userHandler.UpdateProfile)
		protected.DELETE("/users/profile", userHandler.DeleteAccount)

		// Contatos de emergência
		protected.GET("/contacts", contactHandler.GetContacts)
		protected.POST("/contacts", contactHandler.CreateContact)
		protected.PUT("/contacts/:id", contactHandler.UpdateContact)
		protected.DELETE("/contacts/:id", contactHandler.DeleteContact)
	}

	// Iniciar servidor
	port := cfg.Port
	if port == "" {
		port = "8081"
	}
	log.Printf("User Service rodando na porta %s", port)
	log.Fatal(r.Run(":" + port))
}
