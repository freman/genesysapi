// Code generated by go-swagger; DO NOT EDIT.

package widgets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the widgets client
type API interface {
	/*
	   DeleteWidgetsDeployment deletes a widget deployment
	*/
	DeleteWidgetsDeployment(ctx context.Context, params *DeleteWidgetsDeploymentParams) (*DeleteWidgetsDeploymentNoContent, error)
	/*
	   GetWidgetsDeployment gets a widget deployment
	*/
	GetWidgetsDeployment(ctx context.Context, params *GetWidgetsDeploymentParams) (*GetWidgetsDeploymentOK, error)
	/*
	   GetWidgetsDeployments lists widget deployments
	*/
	GetWidgetsDeployments(ctx context.Context, params *GetWidgetsDeploymentsParams) (*GetWidgetsDeploymentsOK, error)
	/*
	   PostWidgetsDeployments creates widget deployment
	*/
	PostWidgetsDeployments(ctx context.Context, params *PostWidgetsDeploymentsParams) (*PostWidgetsDeploymentsOK, error)
	/*
	   PutWidgetsDeployment updates a widget deployment
	*/
	PutWidgetsDeployment(ctx context.Context, params *PutWidgetsDeploymentParams) (*PutWidgetsDeploymentOK, error)
}

// New creates a new widgets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for widgets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteWidgetsDeployment deletes a widget deployment
*/
func (a *Client) DeleteWidgetsDeployment(ctx context.Context, params *DeleteWidgetsDeploymentParams) (*DeleteWidgetsDeploymentNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWidgetsDeployment",
		Method:             "DELETE",
		PathPattern:        "/api/v2/widgets/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWidgetsDeploymentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWidgetsDeploymentNoContent), nil

}

/*
GetWidgetsDeployment gets a widget deployment
*/
func (a *Client) GetWidgetsDeployment(ctx context.Context, params *GetWidgetsDeploymentParams) (*GetWidgetsDeploymentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWidgetsDeployment",
		Method:             "GET",
		PathPattern:        "/api/v2/widgets/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWidgetsDeploymentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWidgetsDeploymentOK), nil

}

/*
GetWidgetsDeployments lists widget deployments
*/
func (a *Client) GetWidgetsDeployments(ctx context.Context, params *GetWidgetsDeploymentsParams) (*GetWidgetsDeploymentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWidgetsDeployments",
		Method:             "GET",
		PathPattern:        "/api/v2/widgets/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWidgetsDeploymentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWidgetsDeploymentsOK), nil

}

/*
PostWidgetsDeployments creates widget deployment
*/
func (a *Client) PostWidgetsDeployments(ctx context.Context, params *PostWidgetsDeploymentsParams) (*PostWidgetsDeploymentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWidgetsDeployments",
		Method:             "POST",
		PathPattern:        "/api/v2/widgets/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWidgetsDeploymentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWidgetsDeploymentsOK), nil

}

/*
PutWidgetsDeployment updates a widget deployment
*/
func (a *Client) PutWidgetsDeployment(ctx context.Context, params *PutWidgetsDeploymentParams) (*PutWidgetsDeploymentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putWidgetsDeployment",
		Method:             "PUT",
		PathPattern:        "/api/v2/widgets/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutWidgetsDeploymentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutWidgetsDeploymentOK), nil

}
