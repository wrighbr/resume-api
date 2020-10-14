package client

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/fatih/structs"
)

func UpdateDocument(collection string, DocID string, body interface{}) {
	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	m := structs.Map(body)
	// fmt.Println(collection)
	// fmt.Println(DocID)
	_, err := client.Collection(collection).Doc(DocID).Set(ctx, m, firestore.MergeAll)
	if err != nil {
		fmt.Println(err)
	}
}
