package grpc

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCClient
type GRPCClient struct {
	*grpc.ClientConn
}

// NewGRPCClient
func NewGRPCClient(host, port string, tlsConfig *tls.Config) (*GRPCClient, error) {
	address := fmt.Sprintf("%s:%s", host, port)
	tlsCredentials := insecure.NewCredentials()
	if tlsConfig != nil {
		tlsCredentials = credentials.NewTLS(tlsConfig)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(tlsCredentials),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to dial gRPC")
	}

	return &GRPCClient{
		ClientConn: conn,
	}, nil
}
