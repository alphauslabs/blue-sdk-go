package session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

type Option interface {
	Apply(*Session)
}

type withClientId string

func (w withClientId) Apply(s *Session) { s.clientId = string(w) }
func WithClientId(v string) Option      { return withClientId(v) }

type withClientSecret string

func (w withClientSecret) Apply(s *Session) { s.clientSecret = string(w) }
func WithClientSecret(v string) Option      { return withClientSecret(v) }

type withGrantType string

func (w withGrantType) Apply(s *Session) { s.grantType = string(w) }
func WithGrantType(v string) Option      { return withGrantType(v) }

type withScope string

func (w withScope) Apply(s *Session) { s.scope = string(w) }
func WithScope(v string) Option      { return withScope(v) }

type withUsername string

func (w withUsername) Apply(s *Session) { s.username = string(w) }
func WithUsername(v string) Option      { return withUsername(v) }

type withPassword string

func (w withPassword) Apply(s *Session) { s.password = string(w) }
func WithPassword(v string) Option      { return withPassword(v) }

type withLoginUrl string

func (w withLoginUrl) Apply(s *Session) { s.loginUrl = string(w) }
func WithLoginUrl(v string) Option      { return withLoginUrl(v) }

type withHttpClient struct{ client *http.Client }

func (w withHttpClient) Apply(s *Session)  { s.httpClient = w.client }
func WithHttpClient(v *http.Client) Option { return withHttpClient{v} }

type Session struct {
	clientId     string
	clientSecret string
	grantType    string
	scope        string
	username     string
	password     string
	loginUrl     string
	httpClient   *http.Client
}

func (s *Session) AccessToken() (string, error) {
	var err error
	var token string
	var body []byte
	var resp *http.Response

	form := url.Values{}
	form.Add("client_id", s.clientId)
	form.Add("client_secret", s.clientSecret)
	form.Add("grant_type", s.grantType)
	form.Add("scope", s.scope)
	if s.grantType == "password" {
		form.Add("username", s.username)
		form.Add("password", s.password)
	}

	resp, err = http.PostForm(s.loginUrl, form)
	if err != nil {
		return token, err
	}

	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if (resp.StatusCode / 100) != 2 {
		return token, fmt.Errorf(resp.Status)
	}

	var m map[string]interface{}
	if err = json.Unmarshal(body, &m); err != nil {
		return token, err
	}

	t, found := m["access_token"]
	if !found {
		return token, fmt.Errorf("cannot find access token")
	}

	token = fmt.Sprintf("%s", t)
	return token, nil
}

func New(o ...Option) *Session {
	s := &Session{
		loginUrl:     "https://login.alphaus.cloud/ripple/access_token",
		clientId:     os.Getenv("ALPHAUS_CLIENT_ID"),
		clientSecret: os.Getenv("ALPHAUS_CLIENT_SECRET"),
		grantType:    "client_credentials",
		scope:        "openid",
		username:     os.Getenv("ALPHAUS_USERNAME"),
		password:     os.Getenv("ALPHAUS_PASSWORD"),
	}

	for _, opt := range o {
		opt.Apply(s)
	}

	return s
}
