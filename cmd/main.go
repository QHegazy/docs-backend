package main

import (
	"context"
	"docs/internal/server"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func gracefulShutdown(apiServer *http.Server) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")
}

// @title   Docs APIs
// @version  1.0
// @description Testing Swagger APIs.
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email mohamed__hegazy@outlook.com
// @license.name MIT License
// @license.url https://github.com/aws/mit-0
// @host   localhost:8080
// @BasePath  /v1
// @schemes  http https
func main() {

	server := server.NewServer()
	go gracefulShutdown(server)

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
