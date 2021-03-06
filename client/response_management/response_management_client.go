// Code generated by go-swagger; DO NOT EDIT.

package response_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the response management client
type API interface {
	/*
	   DeleteResponsemanagementLibrary deletes an existing response library
	   This will remove any responses associated with the library.
	*/
	DeleteResponsemanagementLibrary(ctx context.Context, params *DeleteResponsemanagementLibraryParams) error
	/*
	   DeleteResponsemanagementResponse deletes an existing response
	   This will remove the response from any libraries associated with it.
	*/
	DeleteResponsemanagementResponse(ctx context.Context, params *DeleteResponsemanagementResponseParams) error
	/*
	   GetResponsemanagementLibraries gets a list of existing response libraries
	*/
	GetResponsemanagementLibraries(ctx context.Context, params *GetResponsemanagementLibrariesParams) (*GetResponsemanagementLibrariesOK, error)
	/*
	   GetResponsemanagementLibrary gets details about an existing response library
	*/
	GetResponsemanagementLibrary(ctx context.Context, params *GetResponsemanagementLibraryParams) (*GetResponsemanagementLibraryOK, error)
	/*
	   GetResponsemanagementResponse gets details about an existing response
	*/
	GetResponsemanagementResponse(ctx context.Context, params *GetResponsemanagementResponseParams) (*GetResponsemanagementResponseOK, error)
	/*
	   GetResponsemanagementResponses gets a list of existing responses
	*/
	GetResponsemanagementResponses(ctx context.Context, params *GetResponsemanagementResponsesParams) (*GetResponsemanagementResponsesOK, error)
	/*
	   PostResponsemanagementLibraries creates a response library
	*/
	PostResponsemanagementLibraries(ctx context.Context, params *PostResponsemanagementLibrariesParams) (*PostResponsemanagementLibrariesOK, error)
	/*
	   PostResponsemanagementResponses creates a response
	*/
	PostResponsemanagementResponses(ctx context.Context, params *PostResponsemanagementResponsesParams) (*PostResponsemanagementResponsesOK, error)
	/*
	   PostResponsemanagementResponsesQuery queries responses
	*/
	PostResponsemanagementResponsesQuery(ctx context.Context, params *PostResponsemanagementResponsesQueryParams) (*PostResponsemanagementResponsesQueryOK, error)
	/*
	   PutResponsemanagementLibrary updates an existing response library
	   Fields that can be updated: name. The most recent version is required for updates.
	*/
	PutResponsemanagementLibrary(ctx context.Context, params *PutResponsemanagementLibraryParams) (*PutResponsemanagementLibraryOK, error)
	/*
	   PutResponsemanagementResponse updates an existing response
	   Fields that can be updated: name, libraries, and texts. The most recent version is required for updates.
	*/
	PutResponsemanagementResponse(ctx context.Context, params *PutResponsemanagementResponseParams) (*PutResponsemanagementResponseOK, error)
}

// New creates a new response management API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for response management API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteResponsemanagementLibrary deletes an existing response library

This will remove any responses associated with the library.
*/
func (a *Client) DeleteResponsemanagementLibrary(ctx context.Context, params *DeleteResponsemanagementLibraryParams) error {

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteResponsemanagementLibrary",
		Method:             "DELETE",
		PathPattern:        "/api/v2/responsemanagement/libraries/{libraryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteResponsemanagementLibraryReader{formats: a.formats},
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
DeleteResponsemanagementResponse deletes an existing response

This will remove the response from any libraries associated with it.
*/
func (a *Client) DeleteResponsemanagementResponse(ctx context.Context, params *DeleteResponsemanagementResponseParams) error {

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteResponsemanagementResponse",
		Method:             "DELETE",
		PathPattern:        "/api/v2/responsemanagement/responses/{responseId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteResponsemanagementResponseReader{formats: a.formats},
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
GetResponsemanagementLibraries gets a list of existing response libraries
*/
func (a *Client) GetResponsemanagementLibraries(ctx context.Context, params *GetResponsemanagementLibrariesParams) (*GetResponsemanagementLibrariesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResponsemanagementLibraries",
		Method:             "GET",
		PathPattern:        "/api/v2/responsemanagement/libraries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResponsemanagementLibrariesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetResponsemanagementLibrariesOK), nil

}

/*
GetResponsemanagementLibrary gets details about an existing response library
*/
func (a *Client) GetResponsemanagementLibrary(ctx context.Context, params *GetResponsemanagementLibraryParams) (*GetResponsemanagementLibraryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResponsemanagementLibrary",
		Method:             "GET",
		PathPattern:        "/api/v2/responsemanagement/libraries/{libraryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResponsemanagementLibraryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetResponsemanagementLibraryOK), nil

}

/*
GetResponsemanagementResponse gets details about an existing response
*/
func (a *Client) GetResponsemanagementResponse(ctx context.Context, params *GetResponsemanagementResponseParams) (*GetResponsemanagementResponseOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResponsemanagementResponse",
		Method:             "GET",
		PathPattern:        "/api/v2/responsemanagement/responses/{responseId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResponsemanagementResponseReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetResponsemanagementResponseOK), nil

}

/*
GetResponsemanagementResponses gets a list of existing responses
*/
func (a *Client) GetResponsemanagementResponses(ctx context.Context, params *GetResponsemanagementResponsesParams) (*GetResponsemanagementResponsesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getResponsemanagementResponses",
		Method:             "GET",
		PathPattern:        "/api/v2/responsemanagement/responses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetResponsemanagementResponsesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetResponsemanagementResponsesOK), nil

}

/*
PostResponsemanagementLibraries creates a response library
*/
func (a *Client) PostResponsemanagementLibraries(ctx context.Context, params *PostResponsemanagementLibrariesParams) (*PostResponsemanagementLibrariesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postResponsemanagementLibraries",
		Method:             "POST",
		PathPattern:        "/api/v2/responsemanagement/libraries",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostResponsemanagementLibrariesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostResponsemanagementLibrariesOK), nil

}

/*
PostResponsemanagementResponses creates a response
*/
func (a *Client) PostResponsemanagementResponses(ctx context.Context, params *PostResponsemanagementResponsesParams) (*PostResponsemanagementResponsesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postResponsemanagementResponses",
		Method:             "POST",
		PathPattern:        "/api/v2/responsemanagement/responses",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostResponsemanagementResponsesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostResponsemanagementResponsesOK), nil

}

/*
PostResponsemanagementResponsesQuery queries responses
*/
func (a *Client) PostResponsemanagementResponsesQuery(ctx context.Context, params *PostResponsemanagementResponsesQueryParams) (*PostResponsemanagementResponsesQueryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postResponsemanagementResponsesQuery",
		Method:             "POST",
		PathPattern:        "/api/v2/responsemanagement/responses/query",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostResponsemanagementResponsesQueryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostResponsemanagementResponsesQueryOK), nil

}

/*
PutResponsemanagementLibrary updates an existing response library

Fields that can be updated: name. The most recent version is required for updates.
*/
func (a *Client) PutResponsemanagementLibrary(ctx context.Context, params *PutResponsemanagementLibraryParams) (*PutResponsemanagementLibraryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putResponsemanagementLibrary",
		Method:             "PUT",
		PathPattern:        "/api/v2/responsemanagement/libraries/{libraryId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutResponsemanagementLibraryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutResponsemanagementLibraryOK), nil

}

/*
PutResponsemanagementResponse updates an existing response

Fields that can be updated: name, libraries, and texts. The most recent version is required for updates.
*/
func (a *Client) PutResponsemanagementResponse(ctx context.Context, params *PutResponsemanagementResponseParams) (*PutResponsemanagementResponseOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putResponsemanagementResponse",
		Method:             "PUT",
		PathPattern:        "/api/v2/responsemanagement/responses/{responseId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutResponsemanagementResponseReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutResponsemanagementResponseOK), nil

}
