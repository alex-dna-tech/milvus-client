// Package client implement Milvus Client
package client

import (
	"context"
	"log"

	pb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// milvusClient struct must imlpement milvuspb.MilvusServiceClient interface
var _ pb.MilvusServiceClient = (*milvusClient)(nil)

// NewMilvusServiceClient function creating MilvusClient
func New(ctx context.Context, serverAddr string, opts ...grpc.DialOption) (*milvusClient, error) {
	// Add default options if no option is given
	if opts == nil {
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	}

	conn, err := grpc.DialContext(ctx, serverAddr, opts...)
	if err != nil {
		return nil, err
	}

	return &milvusClient{conn, pb.NewMilvusServiceClient(conn), nil}, nil
}

// WithCallOptions method
func (mc *milvusClient) WithCallOptions(opts ...grpc.CallOption) *milvusClient {
	mc.opts = append(mc.opts, opts...)
	return mc
}

type milvusClient struct {
	conn *grpc.ClientConn
	pb.MilvusServiceClient
	opts []grpc.CallOption
}

// Close close the remaining connection resources
func (mc *milvusClient) Close() {
	err := mc.conn.Close()
	if err != nil {
		log.Fatalf("connection close error: %v", err)
	}
}
