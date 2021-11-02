package main

import (
	"context"
	"io"
	"log"

	"github.com/alphauslabs/blue-sdk-go/cost/v1"
)

func main() {
	ctx := context.Background()
	client, err := cost.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	request := cost.ListAccountsRequest{
		Vendor: "aws",
	}

	stream, err := client.ListAccounts(ctx, &request)
	if err != nil {
		log.Fatal(err)
	}

	for {

		acct, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		log.Println(acct)
	}
}
