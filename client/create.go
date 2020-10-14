package client

import (
	"context"
	"fmt"
)

func CreateDocument(colletion string, body interface{}) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	_, _, err := client.Collection(colletion).Add(ctx, body)
	if err != nil {
		fmt.Println(err)
	}

}
