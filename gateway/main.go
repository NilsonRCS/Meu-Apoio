package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/meuapoio/gateway/middleware"
	"github.com/meuapoio/shared/config"
)

type ServiceRegistry struct {
	UserService string
	// Futuro: AudioService, ContentService, etc.
}

func main() {
	cfg := config.Load()

	// Registry de serviços
	services := &ServiceRegistry{
		UserService: "http://localhost:8081", // User Service
	}

	// Configurar Gin
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	// Middleware global
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Rate limiting
	rateLimiter := middleware.NewRateLimiter(100, time.Minute) // 100 requests por minuto
	r.Use(rateLimiter.Limit())

	// Health check do Gateway
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"service":   "api-gateway",
			"timestamp": time.Now().Unix(),
		})
	})

	// Rotas do User Service
	userGroup := r.Group("/api/v1")
	{
		// Rotas públicas (sem autenticação)
		userGroup.POST("/auth/register", proxyToService(services.UserService))
		userGroup.POST("/auth/login", proxyToService(services.UserService))

		// Rotas protegidas (com autenticação)
		protected := userGroup.Group("")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			// Usuários
			protected.GET("/users/profile", proxyToService(services.UserService))
			protected.PUT("/users/profile", proxyToService(services.UserService))
			protected.DELETE("/users/profile", proxyToService(services.UserService))

			// Contatos
			protected.GET("/contacts", proxyToService(services.UserService))
			protected.POST("/contacts", proxyToService(services.UserService))
			protected.PUT("/contacts/:id", proxyToService(services.UserService))
			protected.DELETE("/contacts/:id", proxyToService(services.UserService))
		}
	}

	// Iniciar servidor
	port := cfg.Port
	if port == "" {
		port = "8080"
	}
	log.Printf("🚀 API Gateway rodando na porta %s", port)
	log.Printf("📋 Serviços registrados:")
	log.Printf("   - User Service: %s", services.UserService)
	log.Fatal(r.Run(":" + port))
}

// proxyToService cria um proxy reverso para um serviço
func proxyToService(serviceURL string) gin.HandlerFunc {
	return gin.WrapH(createReverseProxy(serviceURL))
}

// createReverseProxy cria um proxy reverso HTTP
func createReverseProxy(target string) http.Handler {
	url, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(url)

	// Customizar o director para preservar headers
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)

		// Preservar o path original
		req.URL.Path = req.URL.Path
		req.URL.RawQuery = req.URL.RawQuery

		// Headers importantes
		req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
		req.Header.Set("X-Origin-Service", "api-gateway")
	}

	// Error handling
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("Erro no proxy para %s: %v", target, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"error": "Serviço temporariamente indisponível"}`))
	}

	return proxy
}
