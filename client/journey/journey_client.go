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
