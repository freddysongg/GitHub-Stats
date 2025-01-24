package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"

	"github.com/freddysongg/GitHub-Stats/internal/auth"
	"github.com/freddysongg/GitHub-Stats/internal/cache"
	"github.com/freddysongg/GitHub-Stats/internal/github"
	"github.com/freddysongg/GitHub-Stats/internal/graphql"
	"github.com/freddysongg/GitHub-Stats/internal/health"
	"github.com/freddysongg/GitHub-Stats/internal/kafka"
	"github.com/freddysongg/GitHub-Stats/internal/websocket"
)

func main() {
	// Initialize logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	// Load configuration
	config := loadConfig()

	// Initialize services
	oauthConfig := &oauth2.Config{
		ClientID:     config.GitHubClientID,
		ClientSecret: config.GitHubClientSecret,
		Scopes:       []string{"read:user", "repo"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}

	cache := cache.NewCache()
	githubService := github.NewService(config.GitHubToken, cache, logger)
	wsHub := websocket.NewHub()
	kafkaProducer := kafka.NewProducer(config.KafkaBrokers, logger)
	authService := auth.NewService(oauthConfig, logger)

	// Initialize GraphQL server
	srv := handler.NewDefaultServer(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: &graphql.Resolver{
			GitHubService: githubService,
			Cache:         cache,
			Logger:        logger,
		},
	}))

	// Set up HTTP server
	router := http.NewServeMux()
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.Handle("/health", health.NewHandler())
	router.Handle("/ws", websocket.NewHandler(wsHub))

	server := &http.Server{
		Addr:    ":" + config.Port,
		Handler: router,
	}

	// Start services
	go wsHub.Run()
	go kafkaProducer.Run()

	// Graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	logger.Info("Starting server on port " + config.Port)
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatalf("listen: %s\n", err)
		}
	}()

	<-done
	logger.Info("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("Server Shutdown Failed:%+v", err)
	}
	logger.Info("Server exited properly")
}

type Config struct {
	Port            string
	GitHubClientID  string
	GitHubClientSecret string
	GitHubToken     string
	KafkaBrokers    []string
}

func loadConfig() Config {
	return Config{
		Port:            getEnv("PORT", "8080"),
		GitHubClientID:  getEnv("GITHUB_CLIENT_ID", ""),
		GitHubClientSecret: getEnv("GITHUB_CLIENT_SECRET", ""),
		GitHubToken:     getEnv("GITHUB_TOKEN", ""),
		KafkaBrokers:    []string{getEnv("KAFKA_BROKER", "localhost:9092")},
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
