package conn

import (
	"context"
	"crypto/tls"

	"github.com/alphauslabs/blue-sdk-go/session"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type clientOptions struct {
	target string // gRPC server address
	sess   *session.Session
	conn   *grpc.ClientConn
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

func WithSession(v *session.Session) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.sess = v
	})
}

func WithGrpcConnection(v *grpc.ClientConn) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.conn = v
	})
}

type GrpcClientConn struct {
	*grpc.ClientConn
	opts clientOptions
}

func (c *GrpcClientConn) Close() {
	if c.opts.conn != nil {
		c.opts.conn.Close()
	}
}

func NewClientConn(ctx context.Context, opts ...ClientOption) (*GrpcClientConn, error) {
	sess := session.New()
	co := clientOptions{
		target: session.BlueEndpoint,
		sess:   sess,
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
				LoginUrl:     co.sess.LoginUrl(),
				ClientId:     co.sess.ClientId(),
				ClientSecret: co.sess.ClientSecret(),
			}),
		))

		co.conn, err = grpc.DialContext(ctx, co.target, gopts...)
		if err != nil {
			return nil, err
		}
	}

	return &GrpcClientConn{co.conn, co}, nil
}
