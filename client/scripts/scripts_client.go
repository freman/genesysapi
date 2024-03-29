// Code generated by go-swagger; DO NOT EDIT.

package scripts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the scripts client
type API interface {
	/*
	   GetScript gets a script
	*/
	GetScript(ctx context.Context, params *GetScriptParams) (*GetScriptOK, error)
	/*
	   GetScriptPage gets a page
	*/
	GetScriptPage(ctx context.Context, params *GetScriptPageParams) (*GetScriptPageOK, error)
	/*
	   GetScriptPages gets the list of pages
	*/
	GetScriptPages(ctx context.Context, params *GetScriptPagesParams) (*GetScriptPagesOK, error)
	/*
	   GetScripts gets the list of scripts
	*/
	GetScripts(ctx context.Context, params *GetScriptsParams) (*GetScriptsOK, error)
	/*
	   GetScriptsDivisionviews gets the metadata for a list of scripts
	*/
	GetScriptsDivisionviews(ctx context.Context, params *GetScriptsDivisionviewsParams) (*GetScriptsDivisionviewsOK, error)
	/*
	   GetScriptsPublished gets the published scripts
	*/
	GetScriptsPublished(ctx context.Context, params *GetScriptsPublishedParams) (*GetScriptsPublishedOK, error)
	/*
	   GetScriptsPublishedDivisionviews gets the published scripts metadata
	*/
	GetScriptsPublishedDivisionviews(ctx context.Context, params *GetScriptsPublishedDivisionviewsParams) (*GetScriptsPublishedDivisionviewsOK, error)
	/*
	   GetScriptsPublishedScriptID gets the published script
	*/
	GetScriptsPublishedScriptID(ctx context.Context, params *GetScriptsPublishedScriptIDParams) (*GetScriptsPublishedScriptIDOK, error)
	/*
	   GetScriptsPublishedScriptIDPage gets the published page
	*/
	GetScriptsPublishedScriptIDPage(ctx context.Context, params *GetScriptsPublishedScriptIDPageParams) (*GetScriptsPublishedScriptIDPageOK, error)
	/*
	   GetScriptsPublishedScriptIDPages gets the list of published pages
	*/
	GetScriptsPublishedScriptIDPages(ctx context.Context, params *GetScriptsPublishedScriptIDPagesParams) (*GetScriptsPublishedScriptIDPagesOK, error)
	/*
	   GetScriptsPublishedScriptIDVariables gets the published variables
	*/
	GetScriptsPublishedScriptIDVariables(ctx context.Context, params *GetScriptsPublishedScriptIDVariablesParams) (*GetScriptsPublishedScriptIDVariablesOK, error)
	/*
	   GetScriptsUploadStatus gets the upload status of an imported script
	*/
	GetScriptsUploadStatus(ctx context.Context, params *GetScriptsUploadStatusParams) (*GetScriptsUploadStatusOK, error)
	/*
	   PostScriptExport exports a script via download service
	*/
	PostScriptExport(ctx context.Context, params *PostScriptExportParams) (*PostScriptExportOK, error)
}

// New creates a new scripts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for scripts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
GetScript gets a script
*/
func (a *Client) GetScript(ctx context.Context, params *GetScriptParams) (*GetScriptOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScript",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/{scriptId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptOK), nil

}

/*
GetScriptPage gets a page
*/
func (a *Client) GetScriptPage(ctx context.Context, params *GetScriptPageParams) (*GetScriptPageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptPage",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/{scriptId}/pages/{pageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptPageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptPageOK), nil

}

/*
GetScriptPages gets the list of pages
*/
func (a *Client) GetScriptPages(ctx context.Context, params *GetScriptPagesParams) (*GetScriptPagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptPages",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/{scriptId}/pages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptPagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptPagesOK), nil

}

/*
GetScripts gets the list of scripts
*/
func (a *Client) GetScripts(ctx context.Context, params *GetScriptsParams) (*GetScriptsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScripts",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsOK), nil

}

/*
GetScriptsDivisionviews gets the metadata for a list of scripts
*/
func (a *Client) GetScriptsDivisionviews(ctx context.Context, params *GetScriptsDivisionviewsParams) (*GetScriptsDivisionviewsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsDivisionviews",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/divisionviews",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsDivisionviewsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsDivisionviewsOK), nil

}

/*
GetScriptsPublished gets the published scripts
*/
func (a *Client) GetScriptsPublished(ctx context.Context, params *GetScriptsPublishedParams) (*GetScriptsPublishedOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsPublished",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/published",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsPublishedReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsPublishedOK), nil

}

/*
GetScriptsPublishedDivisionviews gets the published scripts metadata
*/
func (a *Client) GetScriptsPublishedDivisionviews(ctx context.Context, params *GetScriptsPublishedDivisionviewsParams) (*GetScriptsPublishedDivisionviewsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsPublishedDivisionviews",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/published/divisionviews",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsPublishedDivisionviewsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsPublishedDivisionviewsOK), nil

}

/*
GetScriptsPublishedScriptID gets the published script
*/
func (a *Client) GetScriptsPublishedScriptID(ctx context.Context, params *GetScriptsPublishedScriptIDParams) (*GetScriptsPublishedScriptIDOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsPublishedScriptId",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/published/{scriptId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsPublishedScriptIDReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsPublishedScriptIDOK), nil

}

/*
GetScriptsPublishedScriptIDPage gets the published page
*/
func (a *Client) GetScriptsPublishedScriptIDPage(ctx context.Context, params *GetScriptsPublishedScriptIDPageParams) (*GetScriptsPublishedScriptIDPageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsPublishedScriptIdPage",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/published/{scriptId}/pages/{pageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsPublishedScriptIDPageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsPublishedScriptIDPageOK), nil

}

/*
GetScriptsPublishedScriptIDPages gets the list of published pages
*/
func (a *Client) GetScriptsPublishedScriptIDPages(ctx context.Context, params *GetScriptsPublishedScriptIDPagesParams) (*GetScriptsPublishedScriptIDPagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsPublishedScriptIdPages",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/published/{scriptId}/pages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsPublishedScriptIDPagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsPublishedScriptIDPagesOK), nil

}

/*
GetScriptsPublishedScriptIDVariables gets the published variables
*/
func (a *Client) GetScriptsPublishedScriptIDVariables(ctx context.Context, params *GetScriptsPublishedScriptIDVariablesParams) (*GetScriptsPublishedScriptIDVariablesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsPublishedScriptIdVariables",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/published/{scriptId}/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsPublishedScriptIDVariablesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsPublishedScriptIDVariablesOK), nil

}

/*
GetScriptsUploadStatus gets the upload status of an imported script
*/
func (a *Client) GetScriptsUploadStatus(ctx context.Context, params *GetScriptsUploadStatusParams) (*GetScriptsUploadStatusOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getScriptsUploadStatus",
		Method:             "GET",
		PathPattern:        "/api/v2/scripts/uploads/{uploadId}/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetScriptsUploadStatusReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetScriptsUploadStatusOK), nil

}

/*
PostScriptExport exports a script via download service
*/
func (a *Client) PostScriptExport(ctx context.Context, params *PostScriptExportParams) (*PostScriptExportOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postScriptExport",
		Method:             "POST",
		PathPattern:        "/api/v2/scripts/{scriptId}/export",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostScriptExportReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostScriptExportOK), nil

}
