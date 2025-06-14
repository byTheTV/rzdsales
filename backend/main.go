package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"rzd-sales/backend/internal/config"
	"rzd-sales/backend/internal/handlers"
	"rzd-sales/backend/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load configuration
	cfg := config.New()

	// Initialize router with custom configuration
	r := gin.Default()
	r.Use(gin.Recovery())

	// Initialize handlers
	h := handlers.NewHandler(cfg)

	// Setup routes
	setupRoutes(r, h)

	// Create HTTP server
	srv := server.New(cfg.Server, r)

	// Start server in a goroutine
	go func() {
		log.Printf("Server starting on port %s", cfg.Server.Port)
		if err := srv.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create shutdown context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}

func setupRoutes(r *gin.Engine, h *handlers.Handler) {
	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		})

		// Station routes
		stations := v1.Group("/stations")
		{
			stations.GET("", h.SearchStations)
		}

		// Train routes
		trains := v1.Group("/trains")
		{
			trains.GET("", h.SearchTrains)
		}
	}
}
