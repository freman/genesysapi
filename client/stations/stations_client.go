// Code generated by go-swagger; DO NOT EDIT.

package stations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the stations client
type API interface {
	/*
	   DeleteStationAssociateduser unassigns the user assigned to this station
	*/
	DeleteStationAssociateduser(ctx context.Context, params *DeleteStationAssociateduserParams) (*DeleteStationAssociateduserOK, error)
	/*
	   GetStation gets station
	*/
	GetStation(ctx context.Context, params *GetStationParams) (*GetStationOK, error)
	/*
	   GetStations gets the list of available stations
	*/
	GetStations(ctx context.Context, params *GetStationsParams) (*GetStationsOK, error)
	/*
	   GetStationsSettings gets an organization s station settings
	*/
	GetStationsSettings(ctx context.Context, params *GetStationsSettingsParams) (*GetStationsSettingsOK, error)
	/*
	   PatchStationsSettings patches an organization s station settings
	*/
	PatchStationsSettings(ctx context.Context, params *PatchStationsSettingsParams) (*PatchStationsSettingsOK, error)
}

// New creates a new stations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for stations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteStationAssociateduser unassigns the user assigned to this station
*/
func (a *Client) DeleteStationAssociateduser(ctx context.Context, params *DeleteStationAssociateduserParams) (*DeleteStationAssociateduserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteStationAssociateduser",
		Method:             "DELETE",
		PathPattern:        "/api/v2/stations/{stationId}/associateduser",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteStationAssociateduserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteStationAssociateduserOK), nil

}

/*
GetStation gets station
*/
func (a *Client) GetStation(ctx context.Context, params *GetStationParams) (*GetStationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStation",
		Method:             "GET",
		PathPattern:        "/api/v2/stations/{stationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStationOK), nil

}

/*
GetStations gets the list of available stations
*/
func (a *Client) GetStations(ctx context.Context, params *GetStationsParams) (*GetStationsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStations",
		Method:             "GET",
		PathPattern:        "/api/v2/stations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStationsOK), nil

}

/*
GetStationsSettings gets an organization s station settings
*/
func (a *Client) GetStationsSettings(ctx context.Context, params *GetStationsSettingsParams) (*GetStationsSettingsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getStationsSettings",
		Method:             "GET",
		PathPattern:        "/api/v2/stations/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetStationsSettingsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetStationsSettingsOK), nil

}

/*
PatchStationsSettings patches an organization s station settings
*/
func (a *Client) PatchStationsSettings(ctx context.Context, params *PatchStationsSettingsParams) (*PatchStationsSettingsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchStationsSettings",
		Method:             "PATCH",
		PathPattern:        "/api/v2/stations/settings",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchStationsSettingsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchStationsSettingsOK), nil

}
