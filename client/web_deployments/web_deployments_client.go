// Code generated by go-swagger; DO NOT EDIT.

package web_deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the web deployments client
type API interface {
	/*
	   DeleteWebdeploymentsConfiguration deletes all versions of a configuration
	*/
	DeleteWebdeploymentsConfiguration(ctx context.Context, params *DeleteWebdeploymentsConfigurationParams) (*DeleteWebdeploymentsConfigurationNoContent, error)
	/*
	   DeleteWebdeploymentsDeployment deletes a deployment
	*/
	DeleteWebdeploymentsDeployment(ctx context.Context, params *DeleteWebdeploymentsDeploymentParams) (*DeleteWebdeploymentsDeploymentNoContent, error)
	/*
	   DeleteWebdeploymentsDeploymentCobrowseSessionID deletes a cobrowse session
	*/
	DeleteWebdeploymentsDeploymentCobrowseSessionID(ctx context.Context, params *DeleteWebdeploymentsDeploymentCobrowseSessionIDParams) (*DeleteWebdeploymentsDeploymentCobrowseSessionIDOK, *DeleteWebdeploymentsDeploymentCobrowseSessionIDNoContent, error)
	/*
	   DeleteWebdeploymentsTokenRevoke invalidates j w t
	*/
	DeleteWebdeploymentsTokenRevoke(ctx context.Context, params *DeleteWebdeploymentsTokenRevokeParams) (*DeleteWebdeploymentsTokenRevokeNoContent, error)
	/*
	   GetWebdeploymentsConfigurationVersion gets a configuration version
	*/
	GetWebdeploymentsConfigurationVersion(ctx context.Context, params *GetWebdeploymentsConfigurationVersionParams) (*GetWebdeploymentsConfigurationVersionOK, error)
	/*
	   GetWebdeploymentsConfigurationVersions gets the versions of a configuration
	   This returns the 50 most recent versions for this configuration
	*/
	GetWebdeploymentsConfigurationVersions(ctx context.Context, params *GetWebdeploymentsConfigurationVersionsParams) (*GetWebdeploymentsConfigurationVersionsOK, error)
	/*
	   GetWebdeploymentsConfigurationVersionsDraft gets the configuration draft
	*/
	GetWebdeploymentsConfigurationVersionsDraft(ctx context.Context, params *GetWebdeploymentsConfigurationVersionsDraftParams) (*GetWebdeploymentsConfigurationVersionsDraftOK, error)
	/*
	   GetWebdeploymentsConfigurations views configuration drafts
	*/
	GetWebdeploymentsConfigurations(ctx context.Context, params *GetWebdeploymentsConfigurationsParams) (*GetWebdeploymentsConfigurationsOK, error)
	/*
	   GetWebdeploymentsDeployment gets a deployment
	*/
	GetWebdeploymentsDeployment(ctx context.Context, params *GetWebdeploymentsDeploymentParams) (*GetWebdeploymentsDeploymentOK, error)
	/*
	   GetWebdeploymentsDeploymentCobrowseSessionID retrieves a cobrowse session
	*/
	GetWebdeploymentsDeploymentCobrowseSessionID(ctx context.Context, params *GetWebdeploymentsDeploymentCobrowseSessionIDParams) (*GetWebdeploymentsDeploymentCobrowseSessionIDOK, error)
	/*
	   GetWebdeploymentsDeploymentConfigurations gets active configuration for a given deployment
	*/
	GetWebdeploymentsDeploymentConfigurations(ctx context.Context, params *GetWebdeploymentsDeploymentConfigurationsParams) (*GetWebdeploymentsDeploymentConfigurationsOK, error)
	/*
	   GetWebdeploymentsDeployments gets deployments
	*/
	GetWebdeploymentsDeployments(ctx context.Context, params *GetWebdeploymentsDeploymentsParams) (*GetWebdeploymentsDeploymentsOK, error)
	/*
	   PostWebdeploymentsConfigurationVersionsDraftPublish publishes the configuration draft and create a new version
	*/
	PostWebdeploymentsConfigurationVersionsDraftPublish(ctx context.Context, params *PostWebdeploymentsConfigurationVersionsDraftPublishParams) (*PostWebdeploymentsConfigurationVersionsDraftPublishOK, error)
	/*
	   PostWebdeploymentsConfigurations creates a configuration draft
	*/
	PostWebdeploymentsConfigurations(ctx context.Context, params *PostWebdeploymentsConfigurationsParams) (*PostWebdeploymentsConfigurationsOK, *PostWebdeploymentsConfigurationsCreated, error)
	/*
	   PostWebdeploymentsDeployments creates a deployment
	*/
	PostWebdeploymentsDeployments(ctx context.Context, params *PostWebdeploymentsDeploymentsParams) (*PostWebdeploymentsDeploymentsOK, *PostWebdeploymentsDeploymentsCreated, error)
	/*
	   PostWebdeploymentsTokenOauthcodegrantjwtexchange exchanges an o auth code obtained using the authorization code flow for a j w t that can be used by webdeployments
	*/
	PostWebdeploymentsTokenOauthcodegrantjwtexchange(ctx context.Context, params *PostWebdeploymentsTokenOauthcodegrantjwtexchangeParams) (*PostWebdeploymentsTokenOauthcodegrantjwtexchangeOK, error)
	/*
	   PostWebdeploymentsTokenRefresh refreshes a j w t
	*/
	PostWebdeploymentsTokenRefresh(ctx context.Context, params *PostWebdeploymentsTokenRefreshParams) (*PostWebdeploymentsTokenRefreshOK, error)
	/*
	   PutWebdeploymentsConfigurationVersionsDraft updates the configuration draft
	*/
	PutWebdeploymentsConfigurationVersionsDraft(ctx context.Context, params *PutWebdeploymentsConfigurationVersionsDraftParams) (*PutWebdeploymentsConfigurationVersionsDraftOK, error)
	/*
	   PutWebdeploymentsDeployment updates a deployment
	*/
	PutWebdeploymentsDeployment(ctx context.Context, params *PutWebdeploymentsDeploymentParams) (*PutWebdeploymentsDeploymentOK, error)
}

// New creates a new web deployments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for web deployments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteWebdeploymentsConfiguration deletes all versions of a configuration
*/
func (a *Client) DeleteWebdeploymentsConfiguration(ctx context.Context, params *DeleteWebdeploymentsConfigurationParams) (*DeleteWebdeploymentsConfigurationNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWebdeploymentsConfiguration",
		Method:             "DELETE",
		PathPattern:        "/api/v2/webdeployments/configurations/{configurationId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWebdeploymentsConfigurationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWebdeploymentsConfigurationNoContent), nil

}

/*
DeleteWebdeploymentsDeployment deletes a deployment
*/
func (a *Client) DeleteWebdeploymentsDeployment(ctx context.Context, params *DeleteWebdeploymentsDeploymentParams) (*DeleteWebdeploymentsDeploymentNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWebdeploymentsDeployment",
		Method:             "DELETE",
		PathPattern:        "/api/v2/webdeployments/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWebdeploymentsDeploymentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWebdeploymentsDeploymentNoContent), nil

}

/*
DeleteWebdeploymentsDeploymentCobrowseSessionID deletes a cobrowse session
*/
func (a *Client) DeleteWebdeploymentsDeploymentCobrowseSessionID(ctx context.Context, params *DeleteWebdeploymentsDeploymentCobrowseSessionIDParams) (*DeleteWebdeploymentsDeploymentCobrowseSessionIDOK, *DeleteWebdeploymentsDeploymentCobrowseSessionIDNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWebdeploymentsDeploymentCobrowseSessionId",
		Method:             "DELETE",
		PathPattern:        "/api/v2/webdeployments/deployments/{deploymentId}/cobrowse/{sessionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWebdeploymentsDeploymentCobrowseSessionIDReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *DeleteWebdeploymentsDeploymentCobrowseSessionIDOK:
		return value, nil, nil
	case *DeleteWebdeploymentsDeploymentCobrowseSessionIDNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
DeleteWebdeploymentsTokenRevoke invalidates j w t
*/
func (a *Client) DeleteWebdeploymentsTokenRevoke(ctx context.Context, params *DeleteWebdeploymentsTokenRevokeParams) (*DeleteWebdeploymentsTokenRevokeNoContent, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteWebdeploymentsTokenRevoke",
		Method:             "DELETE",
		PathPattern:        "/api/v2/webdeployments/token/revoke",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteWebdeploymentsTokenRevokeReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteWebdeploymentsTokenRevokeNoContent), nil

}

/*
GetWebdeploymentsConfigurationVersion gets a configuration version
*/
func (a *Client) GetWebdeploymentsConfigurationVersion(ctx context.Context, params *GetWebdeploymentsConfigurationVersionParams) (*GetWebdeploymentsConfigurationVersionOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsConfigurationVersion",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/configurations/{configurationId}/versions/{versionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsConfigurationVersionReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsConfigurationVersionOK), nil

}

/*
GetWebdeploymentsConfigurationVersions gets the versions of a configuration

This returns the 50 most recent versions for this configuration
*/
func (a *Client) GetWebdeploymentsConfigurationVersions(ctx context.Context, params *GetWebdeploymentsConfigurationVersionsParams) (*GetWebdeploymentsConfigurationVersionsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsConfigurationVersions",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/configurations/{configurationId}/versions",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsConfigurationVersionsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsConfigurationVersionsOK), nil

}

/*
GetWebdeploymentsConfigurationVersionsDraft gets the configuration draft
*/
func (a *Client) GetWebdeploymentsConfigurationVersionsDraft(ctx context.Context, params *GetWebdeploymentsConfigurationVersionsDraftParams) (*GetWebdeploymentsConfigurationVersionsDraftOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsConfigurationVersionsDraft",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/configurations/{configurationId}/versions/draft",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsConfigurationVersionsDraftReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsConfigurationVersionsDraftOK), nil

}

/*
GetWebdeploymentsConfigurations views configuration drafts
*/
func (a *Client) GetWebdeploymentsConfigurations(ctx context.Context, params *GetWebdeploymentsConfigurationsParams) (*GetWebdeploymentsConfigurationsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsConfigurations",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsConfigurationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsConfigurationsOK), nil

}

/*
GetWebdeploymentsDeployment gets a deployment
*/
func (a *Client) GetWebdeploymentsDeployment(ctx context.Context, params *GetWebdeploymentsDeploymentParams) (*GetWebdeploymentsDeploymentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsDeployment",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsDeploymentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsDeploymentOK), nil

}

/*
GetWebdeploymentsDeploymentCobrowseSessionID retrieves a cobrowse session
*/
func (a *Client) GetWebdeploymentsDeploymentCobrowseSessionID(ctx context.Context, params *GetWebdeploymentsDeploymentCobrowseSessionIDParams) (*GetWebdeploymentsDeploymentCobrowseSessionIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsDeploymentCobrowseSessionId",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/deployments/{deploymentId}/cobrowse/{sessionId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsDeploymentCobrowseSessionIDReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsDeploymentCobrowseSessionIDOK), nil

}

/*
GetWebdeploymentsDeploymentConfigurations gets active configuration for a given deployment
*/
func (a *Client) GetWebdeploymentsDeploymentConfigurations(ctx context.Context, params *GetWebdeploymentsDeploymentConfigurationsParams) (*GetWebdeploymentsDeploymentConfigurationsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsDeploymentConfigurations",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/deployments/{deploymentId}/configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsDeploymentConfigurationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsDeploymentConfigurationsOK), nil

}

/*
GetWebdeploymentsDeployments gets deployments
*/
func (a *Client) GetWebdeploymentsDeployments(ctx context.Context, params *GetWebdeploymentsDeploymentsParams) (*GetWebdeploymentsDeploymentsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWebdeploymentsDeployments",
		Method:             "GET",
		PathPattern:        "/api/v2/webdeployments/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetWebdeploymentsDeploymentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWebdeploymentsDeploymentsOK), nil

}

/*
PostWebdeploymentsConfigurationVersionsDraftPublish publishes the configuration draft and create a new version
*/
func (a *Client) PostWebdeploymentsConfigurationVersionsDraftPublish(ctx context.Context, params *PostWebdeploymentsConfigurationVersionsDraftPublishParams) (*PostWebdeploymentsConfigurationVersionsDraftPublishOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWebdeploymentsConfigurationVersionsDraftPublish",
		Method:             "POST",
		PathPattern:        "/api/v2/webdeployments/configurations/{configurationId}/versions/draft/publish",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWebdeploymentsConfigurationVersionsDraftPublishReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWebdeploymentsConfigurationVersionsDraftPublishOK), nil

}

/*
PostWebdeploymentsConfigurations creates a configuration draft
*/
func (a *Client) PostWebdeploymentsConfigurations(ctx context.Context, params *PostWebdeploymentsConfigurationsParams) (*PostWebdeploymentsConfigurationsOK, *PostWebdeploymentsConfigurationsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWebdeploymentsConfigurations",
		Method:             "POST",
		PathPattern:        "/api/v2/webdeployments/configurations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWebdeploymentsConfigurationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostWebdeploymentsConfigurationsOK:
		return value, nil, nil
	case *PostWebdeploymentsConfigurationsCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PostWebdeploymentsDeployments creates a deployment
*/
func (a *Client) PostWebdeploymentsDeployments(ctx context.Context, params *PostWebdeploymentsDeploymentsParams) (*PostWebdeploymentsDeploymentsOK, *PostWebdeploymentsDeploymentsCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWebdeploymentsDeployments",
		Method:             "POST",
		PathPattern:        "/api/v2/webdeployments/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWebdeploymentsDeploymentsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *PostWebdeploymentsDeploymentsOK:
		return value, nil, nil
	case *PostWebdeploymentsDeploymentsCreated:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
PostWebdeploymentsTokenOauthcodegrantjwtexchange exchanges an o auth code obtained using the authorization code flow for a j w t that can be used by webdeployments
*/
func (a *Client) PostWebdeploymentsTokenOauthcodegrantjwtexchange(ctx context.Context, params *PostWebdeploymentsTokenOauthcodegrantjwtexchangeParams) (*PostWebdeploymentsTokenOauthcodegrantjwtexchangeOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWebdeploymentsTokenOauthcodegrantjwtexchange",
		Method:             "POST",
		PathPattern:        "/api/v2/webdeployments/token/oauthcodegrantjwtexchange",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWebdeploymentsTokenOauthcodegrantjwtexchangeReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWebdeploymentsTokenOauthcodegrantjwtexchangeOK), nil

}

/*
PostWebdeploymentsTokenRefresh refreshes a j w t
*/
func (a *Client) PostWebdeploymentsTokenRefresh(ctx context.Context, params *PostWebdeploymentsTokenRefreshParams) (*PostWebdeploymentsTokenRefreshOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postWebdeploymentsTokenRefresh",
		Method:             "POST",
		PathPattern:        "/api/v2/webdeployments/token/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostWebdeploymentsTokenRefreshReader{formats: a.formats},
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostWebdeploymentsTokenRefreshOK), nil

}

/*
PutWebdeploymentsConfigurationVersionsDraft updates the configuration draft
*/
func (a *Client) PutWebdeploymentsConfigurationVersionsDraft(ctx context.Context, params *PutWebdeploymentsConfigurationVersionsDraftParams) (*PutWebdeploymentsConfigurationVersionsDraftOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putWebdeploymentsConfigurationVersionsDraft",
		Method:             "PUT",
		PathPattern:        "/api/v2/webdeployments/configurations/{configurationId}/versions/draft",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutWebdeploymentsConfigurationVersionsDraftReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutWebdeploymentsConfigurationVersionsDraftOK), nil

}

/*
PutWebdeploymentsDeployment updates a deployment
*/
func (a *Client) PutWebdeploymentsDeployment(ctx context.Context, params *PutWebdeploymentsDeploymentParams) (*PutWebdeploymentsDeploymentOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putWebdeploymentsDeployment",
		Method:             "PUT",
		PathPattern:        "/api/v2/webdeployments/deployments/{deploymentId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutWebdeploymentsDeploymentReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutWebdeploymentsDeploymentOK), nil

}
