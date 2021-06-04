package awscost

import (
	"context"

	"github.com/alphauslabs/blue-sdk-go/conn"
)

type GrpcClient struct {
	AwsCostClient
	conn *conn.GrpcClientConn
}

func (c *GrpcClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// NewClientOptions represents the optional options to NewClient.
type NewClientOptions struct {
	Conn *conn.GrpcClientConn
}

func NewClient(ctx context.Context, opts ...*NewClientOptions) (*GrpcClient, error) {
	var fconn *conn.GrpcClientConn
	var err error
	switch {
	case len(opts) > 0:
		if opts[0].Conn != nil {
			fconn = opts[0].Conn
		}
	default:
		fconn, err = conn.New(ctx)
		if err != nil {
			return nil, err
		}
	}

	cc := NewAwsCostClient(fconn)
	return &GrpcClient{cc, fconn}, nil
}
