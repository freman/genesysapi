// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the events client
type API interface {
	/*
	   PostEventsConversations publishes conversation batch events
	*/
	PostEventsConversations(ctx context.Context, params *PostEventsConversationsParams) (*PostEventsConversationsOK, error)
	/*
	   PostEventsUsersPresence publishes user presence status batch events
	*/
	PostEventsUsersPresence(ctx context.Context, params *PostEventsUsersPresenceParams) (*PostEventsUsersPresenceOK, error)
	/*
	   PostEventsUsersRoutingstatus publishes agent routing status batch events
	*/
	PostEventsUsersRoutingstatus(ctx context.Context, params *PostEventsUsersRoutingstatusParams) (*PostEventsUsersRoutingstatusOK, error)
}

// New creates a new events API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for events API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
PostEventsConversations publishes conversation batch events
*/
func (a *Client) PostEventsConversations(ctx context.Context, params *PostEventsConversationsParams) (*PostEventsConversationsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postEventsConversations",
		Method:             "POST",
		PathPattern:        "/api/v2/events/conversations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEventsConversationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEventsConversationsOK), nil

}

/*
PostEventsUsersPresence publishes user presence status batch events
*/
func (a *Client) PostEventsUsersPresence(ctx context.Context, params *PostEventsUsersPresenceParams) (*PostEventsUsersPresenceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postEventsUsersPresence",
		Method:             "POST",
		PathPattern:        "/api/v2/events/users/presence",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEventsUsersPresenceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEventsUsersPresenceOK), nil

}

/*
PostEventsUsersRoutingstatus publishes agent routing status batch events
*/
func (a *Client) PostEventsUsersRoutingstatus(ctx context.Context, params *PostEventsUsersRoutingstatusParams) (*PostEventsUsersRoutingstatusOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postEventsUsersRoutingstatus",
		Method:             "POST",
		PathPattern:        "/api/v2/events/users/routingstatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostEventsUsersRoutingstatusReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostEventsUsersRoutingstatusOK), nil

}
