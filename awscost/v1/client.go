package awscost

import (
	"context"
	"crypto/tls"
	"os"

	"github.com/alphauslabs/blue-sdk-go/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	Address = ":443" // to be updated after prod deployment
)

type clientOptions struct {
	target       string // gRPC server address
	loginUrl     string // url to get access token
	clientId     string // Alphaus client id
	clientSecret string // Alphaus client secret
	conn         *grpc.ClientConn
}

type ClientOption interface {
	apply(*clientOptions)
}

// fnClientOption wraps a function that modifies clientOptions into an
// implementation of the ClientOption interface.
type fnClientOption struct {
	f func(*clientOptions)
}

func (o *fnClientOption) apply(do *clientOptions) { o.f(do) }

func newFnClientOption(f func(*clientOptions)) *fnClientOption {
	return &fnClientOption{f: f}
}

func WithTarget(s string) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.target = s
	})
}

func WithLoginUrl(s string) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.loginUrl = s
	})
}

func WithClientId(s string) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.clientId = s
	})
}

func WithClientSecret(s string) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.clientSecret = s
	})
}

func WithGrpcConnection(v *grpc.ClientConn) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.conn = v
	})
}

type grpcClient struct {
	AwsCostClient
	opts clientOptions
}

func (c *grpcClient) Close() { c.opts.conn.Close() }

func NewClient(ctx context.Context, opts ...ClientOption) (*grpcClient, error) {
	co := clientOptions{
		target:       Address,
		loginUrl:     session.LoginUrlRipple,
		clientId:     os.Getenv("ALPHAUS_CLIENT_ID"),
		clientSecret: os.Getenv("ALPHAUS_CLIENT_SECRET"),
	}

	for _, opt := range opts {
		opt.apply(&co)
	}

	if co.conn == nil {
		var err error
		var gopts []grpc.DialOption
		creds := credentials.NewTLS(&tls.Config{})
		gopts = append(gopts, grpc.WithTransportCredentials(creds))
		gopts = append(gopts, grpc.WithBlock())
		gopts = append(gopts, grpc.WithPerRPCCredentials(
			session.NewRpcCredentials(session.RpcCredentialsInput{
				LoginUrl:     co.loginUrl,
				ClientId:     co.clientId,
				ClientSecret: co.clientSecret,
			}),
		))

		co.conn, err = grpc.DialContext(ctx, co.target, gopts...)
		if err != nil {
			return nil, err
		}
	}

	cc := NewAwsCostClient(co.conn)
	return &grpcClient{cc, co}, nil
}
