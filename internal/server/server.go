package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"docs/internal/handlers"
)

type Server struct {
	port int
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	newServer := &Server{
		port: port,
	}
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", newServer.port),
		Handler:      handlers.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 25 * time.Second,
	}

	return server
}
