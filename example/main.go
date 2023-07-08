package main

import (
	"context"
	"fmt"
	"time"

	client "github.com/alex-dna-tech/milvus-client"
	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
)

const (
	serverURL         = "localhost:19530"
	newCollectionName = "test"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	// Create milvus client instance
	client, err := client.New(ctx, serverURL)
	if err != nil {
		fmt.Println("failed to connect to Milvus: ", err.Error())
	}
	defer client.Close()

	//--- Collection Create
	idField := schemapb.FieldSchema{
		Name:         "book_id",
		IsPrimaryKey: true,
		Description:  "ID Description",
		DataType:     schemapb.DataType_Int64,
	}

	introField := schemapb.FieldSchema{
		Name:        "book_intro",
		Description: "Intro Description",
		DataType:    schemapb.DataType_FloatVector,
		TypeParams:  []*commonpb.KeyValuePair{{Key: "dim", Value: "2"}},
	}

	schema := &schemapb.CollectionSchema{
		Name:        newCollectionName,
		Description: "Test collection from simple Milvus client",
		AutoID:      false,
		Fields:      []*schemapb.FieldSchema{&idField, &introField},
	}

	schemaBytes, err := proto.Marshal(schema)
	if err != nil {
		fmt.Println("failed proto marshal schema: ", err.Error())
	}

	sts, err := client.CreateCollection(ctx, &milvuspb.CreateCollectionRequest{
		CollectionName: newCollectionName,
		Schema:         schemaBytes,
	})
	if err != nil {
		fmt.Println("failed create collection: ", err.Error())
	}

	fmt.Printf("Create Status: %#v\n", sts)

	// ShowCollections call procedure with ShowCollectionsRequest
	col, err := client.ShowCollections(ctx, &milvuspb.ShowCollectionsRequest{})
	if err != nil {
		fmt.Printf("ShowCollections err: %v\n", err)
	}

	// Printing ShowCollectionsResponse
	fmt.Printf("Milvus Collection: %#v\n", col)

	dc, err := client.DescribeCollection(ctx, &milvuspb.DescribeCollectionRequest{
		CollectionName: newCollectionName,
	})
	if err != nil {
		fmt.Println("failed describe collection: ", err.Error())
	}

	fmt.Printf("Describe Schema: %#v\n", dc.GetSchema())

}
