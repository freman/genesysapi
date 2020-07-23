// Code generated by go-swagger; DO NOT EDIT.

package tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the tokens client
type API interface {
	/*
	   DeleteToken deletes all auth tokens for the specified user
	*/
	DeleteToken(ctx context.Context, params *DeleteTokenParams) (*DeleteTokenNoContent, error)
	/*
	   DeleteTokensMe deletes auth token used to make the request
	*/
	DeleteTokensMe(ctx context.Context, params *DeleteTokensMeParams) (*DeleteTokensMeOK, error)
	/*
	   GetTokensMe fetches information about the current token
	*/
	GetTokensMe(ctx context.Context, params *GetTokensMeParams) (*GetTokensMeOK, error)
}

// New creates a new tokens API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for tokens API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteToken deletes all auth tokens for the specified user
*/
func (a *Client) DeleteToken(ctx context.Context, params *DeleteTokenParams) (*DeleteTokenNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteToken",
		Method:             "DELETE",
		PathPattern:        "/api/v2/tokens/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTokenReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTokenNoContent), nil

}

/*
DeleteTokensMe deletes auth token used to make the request
*/
func (a *Client) DeleteTokensMe(ctx context.Context, params *DeleteTokensMeParams) (*DeleteTokensMeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteTokensMe",
		Method:             "DELETE",
		PathPattern:        "/api/v2/tokens/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteTokensMeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteTokensMeOK), nil

}

/*
GetTokensMe fetches information about the current token
*/
func (a *Client) GetTokensMe(ctx context.Context, params *GetTokensMeParams) (*GetTokensMeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTokensMe",
		Method:             "GET",
		PathPattern:        "/api/v2/tokens/me",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTokensMeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTokensMeOK), nil

}
