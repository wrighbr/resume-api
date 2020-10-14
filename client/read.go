package client

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"
)

func ReadDocument(col string, id int) (docID string, jsonstring interface{}) {
	ctx := context.Background()
	client := createClient(ctx)

	iter := client.Collection(col).Where("ID", "==", id).Documents(ctx)

	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		id := doc.Ref.ID
		d := doc.Data()

		return id, d

	}
	return "", ""
}
