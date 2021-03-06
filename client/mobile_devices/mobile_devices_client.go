// Code generated by go-swagger; DO NOT EDIT.

package mobile_devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the mobile devices client
type API interface {
	/*
	   DeleteMobiledevice deletes device
	*/
	DeleteMobiledevice(ctx context.Context, params *DeleteMobiledeviceParams) (*DeleteMobiledeviceNoContent, error)
	/*
	   GetMobiledevice gets device
	*/
	GetMobiledevice(ctx context.Context, params *GetMobiledeviceParams) (*GetMobiledeviceOK, error)
	/*
	   GetMobiledevices gets a list of all devices
	*/
	GetMobiledevices(ctx context.Context, params *GetMobiledevicesParams) (*GetMobiledevicesOK, error)
	/*
	   PostMobiledevices creates user device
	*/
	PostMobiledevices(ctx context.Context, params *PostMobiledevicesParams) (*PostMobiledevicesOK, error)
	/*
	   PutMobiledevice updates device
	*/
	PutMobiledevice(ctx context.Context, params *PutMobiledeviceParams) (*PutMobiledeviceOK, error)
}

// New creates a new mobile devices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for mobile devices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteMobiledevice deletes device
*/
func (a *Client) DeleteMobiledevice(ctx context.Context, params *DeleteMobiledeviceParams) (*DeleteMobiledeviceNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteMobiledevice",
		Method:             "DELETE",
		PathPattern:        "/api/v2/mobiledevices/{deviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteMobiledeviceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteMobiledeviceNoContent), nil

}

/*
GetMobiledevice gets device
*/
func (a *Client) GetMobiledevice(ctx context.Context, params *GetMobiledeviceParams) (*GetMobiledeviceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMobiledevice",
		Method:             "GET",
		PathPattern:        "/api/v2/mobiledevices/{deviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMobiledeviceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMobiledeviceOK), nil

}

/*
GetMobiledevices gets a list of all devices
*/
func (a *Client) GetMobiledevices(ctx context.Context, params *GetMobiledevicesParams) (*GetMobiledevicesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMobiledevices",
		Method:             "GET",
		PathPattern:        "/api/v2/mobiledevices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMobiledevicesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMobiledevicesOK), nil

}

/*
PostMobiledevices creates user device
*/
func (a *Client) PostMobiledevices(ctx context.Context, params *PostMobiledevicesParams) (*PostMobiledevicesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postMobiledevices",
		Method:             "POST",
		PathPattern:        "/api/v2/mobiledevices",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostMobiledevicesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostMobiledevicesOK), nil

}

/*
PutMobiledevice updates device
*/
func (a *Client) PutMobiledevice(ctx context.Context, params *PutMobiledeviceParams) (*PutMobiledeviceOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putMobiledevice",
		Method:             "PUT",
		PathPattern:        "/api/v2/mobiledevices/{deviceId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutMobiledeviceReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutMobiledeviceOK), nil

}
