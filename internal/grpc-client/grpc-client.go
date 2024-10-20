package grpc_client

import (
	"context"
	pd "docs/internal/document" // Adjust the import path as necessary
	"docs/internal/utils"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

const (
	address = "localhost:50051"
	timeout = 5 * time.Second
)

func GrpcClient(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to gRPC server"})
		return
	}
	defer conn.Close()

	client := pd.NewNewDocumentClient(conn)

	// Get JWT token
	token, err := utils.GetJWTToken("123s") // Pass user ID or any relevant identifier
	if err != nil {
		log.Printf("Could not get token: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		return
	}

	// Set the token in the metadata
	md := metadata.New(map[string]string{"authorization": token})
	ctx = metadata.NewOutgoingContext(ctx, md)

	request := &pd.DocumentRequest{
		Title: "Hello",
	}

	response, err := client.InsertDocument(ctx, request)
	if err != nil {
		log.Printf("Could not insert document: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document"})
		return
	}

	c.JSON(http.StatusOK, response)
}
