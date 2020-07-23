// Code generated by go-swagger; DO NOT EDIT.

package utilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the utilities client
type API interface {
	/*
	   GetDate gets the current system date time
	*/
	GetDate(ctx context.Context, params *GetDateParams) (*GetDateOK, error)
	/*
	   GetIpranges gets public ip address ranges for pure cloud
	*/
	GetIpranges(ctx context.Context, params *GetIprangesParams) (*GetIprangesOK, error)
	/*
	   GetTimezones gets time zones list
	*/
	GetTimezones(ctx context.Context, params *GetTimezonesParams) (*GetTimezonesOK, error)
	/*
	   PostCertificateDetails returns the information about an x509 p e m encoded certificate or certificate chain
	*/
	PostCertificateDetails(ctx context.Context, params *PostCertificateDetailsParams) (*PostCertificateDetailsOK, error)
}

// New creates a new utilities API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for utilities API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetDate gets the current system date time
*/
func (a *Client) GetDate(ctx context.Context, params *GetDateParams) (*GetDateOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDate",
		Method:             "GET",
		PathPattern:        "/api/v2/date",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDateReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDateOK), nil

}

/*
GetIpranges gets public ip address ranges for pure cloud
*/
func (a *Client) GetIpranges(ctx context.Context, params *GetIprangesParams) (*GetIprangesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getIpranges",
		Method:             "GET",
		PathPattern:        "/api/v2/ipranges",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetIprangesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetIprangesOK), nil

}

/*
GetTimezones gets time zones list
*/
func (a *Client) GetTimezones(ctx context.Context, params *GetTimezonesParams) (*GetTimezonesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getTimezones",
		Method:             "GET",
		PathPattern:        "/api/v2/timezones",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetTimezonesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetTimezonesOK), nil

}

/*
PostCertificateDetails returns the information about an x509 p e m encoded certificate or certificate chain
*/
func (a *Client) PostCertificateDetails(ctx context.Context, params *PostCertificateDetailsParams) (*PostCertificateDetailsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postCertificateDetails",
		Method:             "POST",
		PathPattern:        "/api/v2/certificate/details",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostCertificateDetailsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostCertificateDetailsOK), nil

}
