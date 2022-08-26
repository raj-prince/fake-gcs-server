package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func main() {
	client, err := storage.NewClient(context.TODO(), option.WithEndpoint("http://localhost:8080/storage/v1/"))

	if err != nil {
		log.Fatalf("Server creation failed")
	}
	fmt.Println("Prince")
	fmt.Println(client)
}
