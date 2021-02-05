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
	addr        = ":443" // to be updated after prod deployment
	loginRipple = "https://login.alphaus.cloud/ripple/access_token"
	loginWave   = "https://login.alphaus.cloud/access_token"
)

type clientOptions struct {
	target   string // gRPC server address
	loginUrl string // url to get access token
	key      string // Alphaus client id
	secret   string // Alphaus client secret
	conn     *grpc.ClientConn
}

type ClientOption interface {
	apply(*clientOptions)
}

type client struct {
	AwsCostClient
	opts clientOptions
}

func (c *client) Close() { c.opts.conn.Close() }

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

func WithKey(s string) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.key = s
	})
}

func WithSecret(s string) ClientOption {
	return newFnClientOption(func(o *clientOptions) {
		o.secret = s
	})
}

func NewClient(opts ...ClientOption) (*client, error) {
	co := clientOptions{
		target:   addr,
		loginUrl: loginRipple,
		key:      os.Getenv("ALPHAUS_CLIENT_ID"),
		secret:   os.Getenv("ALPHAUS_CLIENT_SECRET"),
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
		gopts = append(gopts, grpc.WithPerRPCCredentials(tokenAuth{
			loginUrl: co.loginUrl,
			key:      co.key,
			secret:   co.secret,
		}))

		co.conn, err = grpc.Dial(co.target, gopts...)
		if err != nil {
			return nil, err
		}
	}

	cc := NewAwsCostClient(co.conn)
	return &client{cc, co}, nil
}

type tokenAuth struct {
	loginUrl string
	key      string
	secret   string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	s := session.New(
		session.WithLoginUrl(t.loginUrl),
		session.WithClientId(t.key),
		session.WithClientSecret(t.secret),
	)

	token, err := s.AccessToken()
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"authorization": "Bearer " + token}, nil
}

func (tokenAuth) RequireTransportSecurity() bool { return false }
