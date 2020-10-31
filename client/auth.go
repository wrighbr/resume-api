package client

import (
	"context"
	"fmt"

	"google.golang.org/api/iterator"
)

func GetUser(username string) (userInfo interface{}) {
	ctx := context.Background()
	client := createClient(ctx)

	iter := client.Collection("users").Where("username", "==", username).Documents(ctx)

	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		d := doc.Data()

		return d

	}
	return nil
}
