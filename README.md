# Simple go Milvus DB client

## Usage example
```go
package main

import (
	"context"
	"fmt"
	"time"

	"github.com/alex-dna-tech/milvus-client"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
)

const (
    serverURL = "localhost:19530"
    collection = "test"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Create milvus client instance
	client, err := client.New(ctx, serverURL)
	if err != nil {
		fmt.Printf("failed to connect to Milvus: ", err.Error())
	}
	defer client.Close()

    // ShowCollections call procedure with ShowCollectionsRequest
    col, err := client.ShowCollections(ctx, &milvuspb.ShowCollectionsRequest{})
    if err != nil {
        fmt.Printf("ShowCollections err: %v\n", err)
    }

    // Printing ShowCollectionsResponse
    fmt.Printf("Milvus Collection: %v\n", col)

}
```
