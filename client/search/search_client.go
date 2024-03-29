// Code generated by go-swagger; DO NOT EDIT.

package search

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the search client
type API interface {
	/*
	   GetDocumentationGknSearch searches gkn documentation using the q64 value returned from a previous search
	*/
	GetDocumentationGknSearch(ctx context.Context, params *GetDocumentationGknSearchParams) (*GetDocumentationGknSearchOK, error)
	/*
	   GetDocumentationSearch searches documentation using the q64 value returned from a previous search
	*/
	GetDocumentationSearch(ctx context.Context, params *GetDocumentationSearchParams) (*GetDocumentationSearchOK, error)
	/*
	   GetSearch searches using the q64 value returned from a previous search
	*/
	GetSearch(ctx context.Context, params *GetSearchParams) (*GetSearchOK, error)
	/*
	   GetSearchSuggest suggests resources using the q64 value returned from a previous suggest query
	*/
	GetSearchSuggest(ctx context.Context, params *GetSearchSuggestParams) (*GetSearchSuggestOK, error)
	/*
	   PostAnalyticsConversationsTranscriptsQuery searches resources
	*/
	PostAnalyticsConversationsTranscriptsQuery(ctx context.Context, params *PostAnalyticsConversationsTranscriptsQueryParams) (*PostAnalyticsConversationsTranscriptsQueryOK, error)
	/*
	   PostDocumentationGknSearch searches gkn documentation
	*/
	PostDocumentationGknSearch(ctx context.Context, params *PostDocumentationGknSearchParams) (*PostDocumentationGknSearchOK, error)
	/*
	   PostDocumentationSearch searches documentation
	*/
	PostDocumentationSearch(ctx context.Context, params *PostDocumentationSearchParams) (*PostDocumentationSearchOK, error)
	/*
	   PostSearch searches resources
	*/
	PostSearch(ctx context.Context, params *PostSearchParams) (*PostSearchOK, error)
	/*
	   PostSearchSuggest suggests resources
	*/
	PostSearchSuggest(ctx context.Context, params *PostSearchSuggestParams) (*PostSearchSuggestOK, error)
	/*
	   PostSpeechandtextanalyticsTranscriptsSearch searches resources
	*/
	PostSpeechandtextanalyticsTranscriptsSearch(ctx context.Context, params *PostSpeechandtextanalyticsTranscriptsSearchParams) (*PostSpeechandtextanalyticsTranscriptsSearchOK, error)
}

// New creates a new search API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for search API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetDocumentationGknSearch searches gkn documentation using the q64 value returned from a previous search
*/
func (a *Client) GetDocumentationGknSearch(ctx context.Context, params *GetDocumentationGknSearchParams) (*GetDocumentationGknSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDocumentationGknSearch",
		Method:             "GET",
		PathPattern:        "/api/v2/documentation/gkn/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDocumentationGknSearchReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDocumentationGknSearchOK), nil

}

/*
GetDocumentationSearch searches documentation using the q64 value returned from a previous search
*/
func (a *Client) GetDocumentationSearch(ctx context.Context, params *GetDocumentationSearchParams) (*GetDocumentationSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDocumentationSearch",
		Method:             "GET",
		PathPattern:        "/api/v2/documentation/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDocumentationSearchReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetDocumentationSearchOK), nil

}

/*
GetSearch searches using the q64 value returned from a previous search
*/
func (a *Client) GetSearch(ctx context.Context, params *GetSearchParams) (*GetSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSearch",
		Method:             "GET",
		PathPattern:        "/api/v2/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchOK), nil

}

/*
GetSearchSuggest suggests resources using the q64 value returned from a previous suggest query
*/
func (a *Client) GetSearchSuggest(ctx context.Context, params *GetSearchSuggestParams) (*GetSearchSuggestOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getSearchSuggest",
		Method:             "GET",
		PathPattern:        "/api/v2/search/suggest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetSearchSuggestReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetSearchSuggestOK), nil

}

/*
PostAnalyticsConversationsTranscriptsQuery searches resources
*/
func (a *Client) PostAnalyticsConversationsTranscriptsQuery(ctx context.Context, params *PostAnalyticsConversationsTranscriptsQueryParams) (*PostAnalyticsConversationsTranscriptsQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postAnalyticsConversationsTranscriptsQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/analytics/conversations/transcripts/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAnalyticsConversationsTranscriptsQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAnalyticsConversationsTranscriptsQueryOK), nil

}

/*
PostDocumentationGknSearch searches gkn documentation
*/
func (a *Client) PostDocumentationGknSearch(ctx context.Context, params *PostDocumentationGknSearchParams) (*PostDocumentationGknSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDocumentationGknSearch",
		Method:             "POST",
		PathPattern:        "/api/v2/documentation/gkn/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDocumentationGknSearchReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDocumentationGknSearchOK), nil

}

/*
PostDocumentationSearch searches documentation
*/
func (a *Client) PostDocumentationSearch(ctx context.Context, params *PostDocumentationSearchParams) (*PostDocumentationSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postDocumentationSearch",
		Method:             "POST",
		PathPattern:        "/api/v2/documentation/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostDocumentationSearchReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostDocumentationSearchOK), nil

}

/*
PostSearch searches resources
*/
func (a *Client) PostSearch(ctx context.Context, params *PostSearchParams) (*PostSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSearch",
		Method:             "POST",
		PathPattern:        "/api/v2/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSearchReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSearchOK), nil

}

/*
PostSearchSuggest suggests resources
*/
func (a *Client) PostSearchSuggest(ctx context.Context, params *PostSearchSuggestParams) (*PostSearchSuggestOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSearchSuggest",
		Method:             "POST",
		PathPattern:        "/api/v2/search/suggest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSearchSuggestReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSearchSuggestOK), nil

}

/*
PostSpeechandtextanalyticsTranscriptsSearch searches resources
*/
func (a *Client) PostSpeechandtextanalyticsTranscriptsSearch(ctx context.Context, params *PostSpeechandtextanalyticsTranscriptsSearchParams) (*PostSpeechandtextanalyticsTranscriptsSearchOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postSpeechandtextanalyticsTranscriptsSearch",
		Method:             "POST",
		PathPattern:        "/api/v2/speechandtextanalytics/transcripts/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostSpeechandtextanalyticsTranscriptsSearchReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostSpeechandtextanalyticsTranscriptsSearchOK), nil

}
