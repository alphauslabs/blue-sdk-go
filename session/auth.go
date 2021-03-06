package session

import (
	"context"
)

type tokenAuth struct {
	loginUrl     string
	clientId     string
	clientSecret string
	accessToken  string
}

func (t tokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	s := New(
		WithLoginUrl(t.loginUrl),
		WithClientId(t.clientId),
		WithClientSecret(t.clientSecret),
	)

	token := t.accessToken
	if token == "" {
		var err error
		token, err = s.AccessToken()
		if err != nil {
			return map[string]string{}, err
		}
	}

	return map[string]string{"authorization": "Bearer " + token}, nil
}

func (tokenAuth) RequireTransportSecurity() bool { return false }

type RpcCredentialsInput struct {
	LoginUrl     string
	ClientId     string
	ClientSecret string
	AccessToken  string // use directly if non-empty; disregard others
}

func NewRpcCredentials(in ...RpcCredentialsInput) tokenAuth {
	var accessToken string
	sess := New()
	loginUrl := sess.LoginUrl()
	clientId := sess.ClientId()
	clientSecret := sess.ClientSecret()
	if len(in) > 0 {
		accessToken = in[0].AccessToken
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
		accessToken:  accessToken,
	}
}
