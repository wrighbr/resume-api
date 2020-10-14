package client

import "context"

func DeleteDocument(Collection string, ID string) {

	ctx := context.Background()
	client := createClient(ctx)
	defer client.Close()

	client.Collection(Collection).Doc(ID).Delete(ctx)

}
