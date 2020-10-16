// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the journey client
type API interface {
	/*
	   GetJourneyActiontarget retrieves a single action target
	*/
	GetJourneyActiontarget(ctx context.Context, params *GetJourneyActiontargetParams) (*GetJourneyActiontargetOK, error)
	/*
	   GetJourneyActiontargets retrieves all action targets
	*/
	GetJourneyActiontargets(ctx context.Context, params *GetJourneyActiontargetsParams) (*GetJourneyActiontargetsOK, error)
	/*
	   PatchJourneyActiontarget updates a single action target
	*/
	PatchJourneyActiontarget(ctx context.Context, params *PatchJourneyActiontargetParams) (*PatchJourneyActiontargetOK, error)
	/*
	   PostAnalyticsJourneysAggregatesQuery queries for journey aggregates
	*/
	PostAnalyticsJourneysAggregatesQuery(ctx context.Context, params *PostAnalyticsJourneysAggregatesQueryParams) (*PostAnalyticsJourneysAggregatesQueryOK, error)
}

// New creates a new journey API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for journey API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetJourneyActiontarget retrieves a single action target
*/
func (a *Client) GetJourneyActiontarget(ctx context.Context, params *GetJourneyActiontargetParams) (*GetJourneyActiontargetOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActiontarget",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actiontargets/{actionTargetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActiontargetReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActiontargetOK), nil

}

/*
GetJourneyActiontargets retrieves all action targets
*/
func (a *Client) GetJourneyActiontargets(ctx context.Context, params *GetJourneyActiontargetsParams) (*GetJourneyActiontargetsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJourneyActiontargets",
		Method:             "GET",
		PathPattern:        "/api/v2/journey/actiontargets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetJourneyActiontargetsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJourneyActiontargetsOK), nil

}

/*
PatchJourneyActiontarget updates a single action target
*/
func (a *Client) PatchJourneyActiontarget(ctx context.Context, params *PatchJourneyActiontargetParams) (*PatchJourneyActiontargetOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "patchJourneyActiontarget",
		Method:             "PATCH",
		PathPattern:        "/api/v2/journey/actiontargets/{actionTargetId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchJourneyActiontargetReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PatchJourneyActiontargetOK), nil

}

/*
PostAnalyticsJourneysAggregatesQuery queries for journey aggregates
*/
func (a *Client) PostAnalyticsJourneysAggregatesQuery(ctx context.Context, params *PostAnalyticsJourneysAggregatesQueryParams) (*PostAnalyticsJourneysAggregatesQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAnalyticsJourneysAggregatesQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/analytics/journeys/aggregates/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAnalyticsJourneysAggregatesQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAnalyticsJourneysAggregatesQueryOK), nil

}
