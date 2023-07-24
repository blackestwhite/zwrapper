package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/blackestwhite/zwrapper/config"
	"github.com/blackestwhite/zwrapper/db"
	"github.com/blackestwhite/zwrapper/gateway"
	"github.com/blackestwhite/zwrapper/router"
	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Load()
	db.Connect()
	gateway.Initiate()

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.LoadHTMLGlob("./templates/*")
	router.Setup(engine)

	// Create a channel to receive the OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Start the server in a separate goroutine
	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		log.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for an OS signal to shut down the server gracefully
	<-quit
	log.Println("Received OS signal. Shutting down gracefully...")

	// Set a timeout for the graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown the server with the given context
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	log.Println("Server shut down.")

	// Disconnect the database
	db.Disconnect()
}
