package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"rsc.io/quote"
)

type Person struct {
	Name string `json:"name"`
}

func handleRequest(ctx context.Context, event json.RawMessage) error {
	var person Person
	if err := json.Unmarshal(event, &person); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return err
	}

	fmt.Printf("Hello, %s!\n", person.Name)
	fmt.Println(quote.Go())

	return nil
}

func main() {
	lambda.Start(handleRequest)
}
