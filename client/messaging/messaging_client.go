// Code generated by go-swagger; DO NOT EDIT.

package messaging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the messaging client
type API interface {
	/*
	   DeleteMessagingSupportedcontentSupportedContentID deletes a supported content profile
	*/
	DeleteMessagingSupportedcontentSupportedContentID(ctx context.Context, params *DeleteMessagingSupportedcontentSupportedContentIDParams) (*DeleteMessagingSupportedcontentSupportedContentIDNoContent, error)
	/*
	   GetMessagingSupportedcontent gets a list of supported content profiles
	*/
	GetMessagingSupportedcontent(ctx context.Context, params *GetMessagingSupportedcontentParams) (*GetMessagingSupportedcontentOK, error)
	/*
	   GetMessagingSupportedcontentSupportedContentID gets a supported content profile
	*/
	GetMessagingSupportedcontentSupportedContentID(ctx context.Context, params *GetMessagingSupportedcontentSupportedContentIDParams) (*GetMessagingSupportedcontentSupportedContentIDOK, error)
	/*
	   PatchMessagingSupportedcontentSupportedContentID updates a supported content profile
	*/
	PatchMessagingSupportedcontentSupportedContentID(ctx context.Context, params *PatchMessagingSupportedcontentSupportedContentIDParams) (*PatchMessagingSupportedcontentSupportedContentIDOK, error)
	/*
	   PostMessagingSupportedcontent creates a supported content profile
	*/
	PostMessagingSupportedcontent(ctx context.Context, params *PostMessagingSupportedcontentParams) (*PostMessagingSupportedcontentOK, *PostMessagingSupportedcontentCreated, error)
}

// New creates a new messaging API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for messaging API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteMessagingSupportedcontentSupportedContentID deletes a supported content profile
*/
func (a *Client) DeleteMessagingSupportedcontentSupportedContentID(ctx context.Context, params *DeleteMessagingSupportedcontentSupportedContentIDParams) (*DeleteMessagingSupportedcontentSupportedContentIDNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMessagingSupportedcontentSupportedContentId",
		Method:             "DELETE",
		PathPattern:        "/api/v2/messaging/supportedcontent/{supportedContentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMessagingSupportedcontentSupportedContentIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMessagingSupportedcontentSupportedContentIDNoContent), nil

}

/*
GetMessagingSupportedcontent gets a list of supported content profiles
*/
func (a *Client) GetMessagingSupportedcontent(ctx context.Context, params *GetMessagingSupportedcontentParams) (*GetMessagingSupportedcontentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMessagingSupportedcontent",
		Method:             "GET",
		PathPattern:        "/api/v2/messaging/supportedcontent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMessagingSupportedcontentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMessagingSupportedcontentOK), nil

}

/*
GetMessagingSupportedcontentSupportedContentID gets a supported content profile
*/
func (a *Client) GetMessagingSupportedcontentSupportedContentID(ctx context.Context, params *GetMessagingSupportedcontentSupportedContentIDParams) (*GetMessagingSupportedcontentSupportedContentIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMessagingSupportedcontentSupportedContentId",
		Method:             "GET",
		PathPattern:        "/api/v2/messaging/supportedcontent/{supportedContentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMessagingSupportedcontentSupportedContentIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMessagingSupportedcontentSupportedContentIDOK), nil

}

/*
PatchMessagingSupportedcontentSupportedContentID updates a supported content profile
*/
func (a *Client) PatchMessagingSupportedcontentSupportedContentID(ctx context.Context, params *PatchMessagingSupportedcontentSupportedContentIDParams) (*PatchMessagingSupportedcontentSupportedContentIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchMessagingSupportedcontentSupportedContentId",
		Method:             "PATCH",
		PathPattern:        "/api/v2/messaging/supportedcontent/{supportedContentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchMessagingSupportedcontentSupportedContentIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchMessagingSupportedcontentSupportedContentIDOK), nil

}

/*
PostMessagingSupportedcontent creates a supported content profile
*/
func (a *Client) PostMessagingSupportedcontent(ctx context.Context, params *PostMessagingSupportedcontentParams) (*PostMessagingSupportedcontentOK, *PostMessagingSupportedcontentCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postMessagingSupportedcontent",
		Method:             "POST",
		PathPattern:        "/api/v2/messaging/supportedcontent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostMessagingSupportedcontentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostMessagingSupportedcontentOK:
		return value, nil, nil
	case *PostMessagingSupportedcontentCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}