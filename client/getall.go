package client

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"
)

func GetAllDocuments(collection string) {
	ctx := context.Background()
	client := createClient(ctx)

	iter := client.Collection(collection).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(doc.Data())
	}
}
