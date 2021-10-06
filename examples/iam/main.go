package main

import (
	"context"
	"log"

	"github.com/alphauslabs/blue-sdk-go/iam/v1"
)

func main() {
	ctx := context.Background()
	client, err := iam.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()
	out, err := client.WhoAmI(ctx, &iam.WhoAmIRequest{})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(out)
}
