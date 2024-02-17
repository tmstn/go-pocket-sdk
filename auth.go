package pocket

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

const (
	AuthorizeUrl             = "https://getpocket.com/auth/authorize?request_token=%s&redirect_uri=%s"
	TokenRequestEndpoint     = "/oauth/request"
	AuthorizeRequestEndpoint = "/oauth/authorize"
)

type RequestAuthParams struct {
	AccessToken string `json:"access_token"`
	ConsumerKey string `json:"consumer_key"`
}

type AuthParams struct {
	*RequestAuthParams
}

type TokenRequestParams struct {
	ConsumerKey string `json:"consumer_key"`
	RedirectURI string `json:"redirect_uri"`
}

type TokenRequestResponse struct {
	Code string `json:"code"`
}

type AuthorizeRequestParams struct {
	ConsumerKey string `json:"consumer_key"`
	Code        string `json:"code"`
}

type AuthorizeRequestResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

type Auth struct {
	client *Client
	*AuthParams
	Username string `json:"username"`
}

func (a *Auth) GetRequestToken(ctx context.Context, redirectUrl string) (*TokenRequestResponse, error) {
	params := &TokenRequestParams{
		ConsumerKey: a.ConsumerKey,
		RedirectURI: redirectUrl,
	}

	b, err := json.Marshal(params)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal input body")
	}

	r, err := a.client.do(ctx, TokenRequestEndpoint, b)
	if err != nil {
		return nil, err
	}

	response := &TokenRequestResponse{}
	err = json.Unmarshal(r, response)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal response body")
	}

	if response.Code == "" {
		return nil, errors.New("empty request token in API response")
	}

	return response, nil
}

func (a *Auth) GetAuthorizationURL(requestToken, redirectUrl string) (string, error) {
	if requestToken == "" || redirectUrl == "" {
		return "", errors.New("empty params")
	}

	return fmt.Sprintf(AuthorizeUrl, requestToken, redirectUrl), nil
}

func (a *Auth) Authorize(ctx context.Context, requestToken string) (*AuthorizeRequestResponse, error) {
	if requestToken == "" {
		return nil, errors.New("empty request token")
	}

	params := &AuthorizeRequestParams{
		Code:        requestToken,
		ConsumerKey: a.ConsumerKey,
	}

	b, err := json.Marshal(params)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to marshal input body")
	}

	r, err := a.client.do(ctx, AuthorizeRequestEndpoint, b)
	if err != nil {
		return nil, err
	}

	response := &AuthorizeRequestResponse{}
	err = json.Unmarshal(r, response)
	if err != nil {
		return nil, errors.WithMessage(err, "failed to unmarshal response body")
	}

	return response, nil
}
