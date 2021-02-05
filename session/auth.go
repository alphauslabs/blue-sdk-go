package session

import (
	"context"
	"os"
)

type tokenAuth struct {
	loginUrl     string
	clientId     string
	clientSecret string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	s := New(
		WithLoginUrl(t.loginUrl),
		WithClientId(t.clientId),
		WithClientSecret(t.clientSecret),
	)

	token, err := s.AccessToken()
	if err != nil {
		return map[string]string{}, err
	}

	return map[string]string{"authorization": "Bearer " + token}, nil
}

func (tokenAuth) RequireTransportSecurity() bool { return false }

type RpcCredentialsInput struct {
	LoginUrl     string // default: Ripple
	ClientId     string // default: $ALPHAUS_CLIENT_ID
	ClientSecret string // default: $ALPHAUS_CLIENT_SECRET
}

func NewRpcCredentials(in ...RpcCredentialsInput) tokenAuth {
	loginUrl := LoginUrlRipple
	clientId := os.Getenv("ALPHAUS_CLIENT_ID")
	clientSecret := os.Getenv("ALPHAUS_CLIENT_SECRET")
	if len(in) > 0 {
		if in[0].LoginUrl != "" {
			loginUrl = in[0].LoginUrl
		}

		if in[0].ClientId != "" {
			clientId = in[0].ClientId
		}

		if in[0].ClientSecret != "" {
			clientSecret = in[0].ClientSecret
		}
	}

	return tokenAuth{
		loginUrl:     loginUrl,
		clientId:     clientId,
		clientSecret: clientSecret,
	}
}
