package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

var (
	githubClient *githubv4.Client
	wsServer    *WebSocketServer
)

type GitHubStats struct {
	Profile      interface{} `json:"profile"`
	Repositories []RepoStats `json:"repositories"`
	Languages    []Language  `json:"languages"`
	Contributions interface{} `json:"contributions"`
}

type RepoStats struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Stars       int    `json:"stars"`
	Forks       int    `json:"forks"`
	OpenIssues  int    `json:"open_issues"`
	PullRequests int    `json:"pull_requests"`
}

type Language struct {
	Name  string `json:"name"`
	Bytes int    `json:"bytes"`
}

func main() {
	// Initialize GitHub client
	initGitHubClient()

	// Create router
	r := mux.NewRouter()

	// API endpoints
	r.HandleFunc("/profile", getProfile).Methods("GET")
	r.HandleFunc("/repositories", getRepositories).Methods("GET")
	r.HandleFunc("/languages", getLanguages).Methods("GET")
	r.HandleFunc("/contributions", getContributions).Methods("GET")
	r.HandleFunc("/ws", handleWebSocket)

	// Health check
	r.HandleFunc("/health", healthCheck).Methods("GET")

	// Start server
	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}

	// Create WebSocket server
	wsServer = NewWebSocketServer()

	// Start HTTP server (includes WebSocket endpoint)
	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("Shutting down")
	os.Exit(0)
}

func initGitHubClient() {
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
	)
	httpClient := oauth2.NewClient(context.Background(), src)
	githubClient = githubv4.NewClient(httpClient)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	// Implementation to fetch profile data
}

func getRepositories(w http.ResponseWriter, r *http.Request) {
	// Implementation to fetch repository data
}

func getLanguages(w http.ResponseWriter, r *http.Request) {
	// Implementation to fetch language data
}

func getContributions(w http.ResponseWriter, r *http.Request) {
	// Implementation to fetch contribution data
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	wsServer.HandleConnections(w, r)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"status":   "ok",
		"services": []string{
			"http_server:8080",
			"websocket_server:8081",
		},
		"timestamp": time.Now().UTC(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
