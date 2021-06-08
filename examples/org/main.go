package main

import (
	"context"
	"log"

	"github.com/alphauslabs/blue-sdk-go/org/v1"
)

func main() {
	ctx := context.Background()
	client, err := org.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	out, err := client.GetInfo(ctx, &org.GetInfoRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}
