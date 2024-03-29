// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the notifications client
type API interface {
	/*
	   DeleteNotificationsChannelSubscriptions removes all subscriptions
	*/
	DeleteNotificationsChannelSubscriptions(ctx context.Context, params *DeleteNotificationsChannelSubscriptionsParams) error
	/*
	   GetNotificationsAvailabletopics gets available notification topics
	*/
	GetNotificationsAvailabletopics(ctx context.Context, params *GetNotificationsAvailabletopicsParams) (*GetNotificationsAvailabletopicsOK, error)
	/*
	   GetNotificationsChannelSubscriptions thes list of all subscriptions for this channel
	*/
	GetNotificationsChannelSubscriptions(ctx context.Context, params *GetNotificationsChannelSubscriptionsParams) (*GetNotificationsChannelSubscriptionsOK, error)
	/*
	   GetNotificationsChannels thes list of existing channels
	*/
	GetNotificationsChannels(ctx context.Context, params *GetNotificationsChannelsParams) (*GetNotificationsChannelsOK, error)
	/*
	   HeadNotificationsChannel verifies a channel still exists and is valid
	   Returns a 200 OK if channel exists, and a 404 Not Found if it doesn't
	*/
	HeadNotificationsChannel(ctx context.Context, params *HeadNotificationsChannelParams) error
	/*
	   PostNotificationsChannelSubscriptions adds a list of subscriptions to the existing list of subscriptions
	*/
	PostNotificationsChannelSubscriptions(ctx context.Context, params *PostNotificationsChannelSubscriptionsParams) (*PostNotificationsChannelSubscriptionsOK, error)
	/*
	   PostNotificationsChannels creates a new channel
	   There is a limit of 20 channels per user/app combination. Creating a 21st channel will remove the channel with oldest last used date. Channels without an active connection will be removed first.
	*/
	PostNotificationsChannels(ctx context.Context, params *PostNotificationsChannelsParams) (*PostNotificationsChannelsOK, error)
	/*
	   PutNotificationsChannelSubscriptions replaces the current list of subscriptions with a new list
	*/
	PutNotificationsChannelSubscriptions(ctx context.Context, params *PutNotificationsChannelSubscriptionsParams) (*PutNotificationsChannelSubscriptionsOK, error)
}

// New creates a new notifications API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for notifications API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteNotificationsChannelSubscriptions removes all subscriptions
*/
func (a *Client) DeleteNotificationsChannelSubscriptions(ctx context.Context, params *DeleteNotificationsChannelSubscriptionsParams) error {

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNotificationsChannelSubscriptions",
		Method:             "DELETE",
		PathPattern:        "/api/v2/notifications/channels/{channelId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNotificationsChannelSubscriptionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetNotificationsAvailabletopics gets available notification topics
*/
func (a *Client) GetNotificationsAvailabletopics(ctx context.Context, params *GetNotificationsAvailabletopicsParams) (*GetNotificationsAvailabletopicsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNotificationsAvailabletopics",
		Method:             "GET",
		PathPattern:        "/api/v2/notifications/availabletopics",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsAvailabletopicsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNotificationsAvailabletopicsOK), nil

}

/*
GetNotificationsChannelSubscriptions thes list of all subscriptions for this channel
*/
func (a *Client) GetNotificationsChannelSubscriptions(ctx context.Context, params *GetNotificationsChannelSubscriptionsParams) (*GetNotificationsChannelSubscriptionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNotificationsChannelSubscriptions",
		Method:             "GET",
		PathPattern:        "/api/v2/notifications/channels/{channelId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsChannelSubscriptionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNotificationsChannelSubscriptionsOK), nil

}

/*
GetNotificationsChannels thes list of existing channels
*/
func (a *Client) GetNotificationsChannels(ctx context.Context, params *GetNotificationsChannelsParams) (*GetNotificationsChannelsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNotificationsChannels",
		Method:             "GET",
		PathPattern:        "/api/v2/notifications/channels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationsChannelsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetNotificationsChannelsOK), nil

}

/*
HeadNotificationsChannel verifies a channel still exists and is valid

Returns a 200 OK if channel exists, and a 404 Not Found if it doesn't
*/
func (a *Client) HeadNotificationsChannel(ctx context.Context, params *HeadNotificationsChannelParams) error {

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "headNotificationsChannel",
		Method:             "HEAD",
		PathPattern:        "/api/v2/notifications/channels/{channelId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &HeadNotificationsChannelReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
PostNotificationsChannelSubscriptions adds a list of subscriptions to the existing list of subscriptions
*/
func (a *Client) PostNotificationsChannelSubscriptions(ctx context.Context, params *PostNotificationsChannelSubscriptionsParams) (*PostNotificationsChannelSubscriptionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postNotificationsChannelSubscriptions",
		Method:             "POST",
		PathPattern:        "/api/v2/notifications/channels/{channelId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostNotificationsChannelSubscriptionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNotificationsChannelSubscriptionsOK), nil

}

/*
PostNotificationsChannels creates a new channel

There is a limit of 20 channels per user/app combination. Creating a 21st channel will remove the channel with oldest last used date. Channels without an active connection will be removed first.
*/
func (a *Client) PostNotificationsChannels(ctx context.Context, params *PostNotificationsChannelsParams) (*PostNotificationsChannelsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postNotificationsChannels",
		Method:             "POST",
		PathPattern:        "/api/v2/notifications/channels",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostNotificationsChannelsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostNotificationsChannelsOK), nil

}

/*
PutNotificationsChannelSubscriptions replaces the current list of subscriptions with a new list
*/
func (a *Client) PutNotificationsChannelSubscriptions(ctx context.Context, params *PutNotificationsChannelSubscriptionsParams) (*PutNotificationsChannelSubscriptionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putNotificationsChannelSubscriptions",
		Method:             "PUT",
		PathPattern:        "/api/v2/notifications/channels/{channelId}/subscriptions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutNotificationsChannelSubscriptionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutNotificationsChannelSubscriptionsOK), nil

}
