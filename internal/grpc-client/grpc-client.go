package grpc_client

import (
	"context"
	dto "docs/internal/Dto"
	pd "docs/internal/document"
	"docs/internal/utils"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var (
	address = os.Getenv("GRPC_SERVER")
	timeout = 5 * time.Second
)

func GrpcClient(docData dto.DocPost) chan string {
	docIDChan := make(chan string)

	go func() {
		defer close(docIDChan)

		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Printf("Failed to connect: %v", err)
			return
		}
		defer conn.Close()

		client := pd.NewNewDocumentClient(conn)

		token, err := utils.GetJWTToken(docData.UserUuid.String())
		if err != nil {
			log.Printf("Could not get token: %v", err)
			return
		}

		md := metadata.New(map[string]string{"authorization": token})
		ctx = metadata.NewOutgoingContext(ctx, md)

		request := &pd.DocumentRequest{
			Title: docData.DocName,
		}

		response, err := client.InsertDocument(ctx, request)
		if err != nil {
			log.Printf("Could not insert document: %v", err)
			return
		}
		docID := response.DocumentId
		docIDChan <- docID
	}()
	return docIDChan
}
