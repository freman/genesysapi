// Code generated by go-swagger; DO NOT EDIT.

package languages

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the languages client
type API interface {
	/*
	   DeleteLanguage deletes language deprecated
	   This endpoint is deprecated. It has been moved to /routing/languages/{languageId}
	*/
	DeleteLanguage(ctx context.Context, params *DeleteLanguageParams) error
	/*
	   DeleteRoutingLanguage deletes language
	*/
	DeleteRoutingLanguage(ctx context.Context, params *DeleteRoutingLanguageParams) error
	/*
	   GetLanguage gets language deprecated
	   This endpoint is deprecated. It has been moved to /routing/languages/{languageId}
	*/
	GetLanguage(ctx context.Context, params *GetLanguageParams) (*GetLanguageOK, error)
	/*
	   GetLanguages gets the list of supported languages deprecated
	   This endpoint is deprecated. It has been moved to /routing/languages
	*/
	GetLanguages(ctx context.Context, params *GetLanguagesParams) (*GetLanguagesOK, error)
	/*
	   GetLanguagesTranslations gets all available languages for translation
	*/
	GetLanguagesTranslations(ctx context.Context, params *GetLanguagesTranslationsParams) (*GetLanguagesTranslationsOK, error)
	/*
	   GetLanguagesTranslationsBuiltin gets the builtin translation for a language
	*/
	GetLanguagesTranslationsBuiltin(ctx context.Context, params *GetLanguagesTranslationsBuiltinParams) (*GetLanguagesTranslationsBuiltinOK, error)
	/*
	   GetLanguagesTranslationsOrganization gets effective translation for an organization by language
	*/
	GetLanguagesTranslationsOrganization(ctx context.Context, params *GetLanguagesTranslationsOrganizationParams) (*GetLanguagesTranslationsOrganizationOK, error)
	/*
	   GetLanguagesTranslationsUser gets effective language translation for a user
	*/
	GetLanguagesTranslationsUser(ctx context.Context, params *GetLanguagesTranslationsUserParams) (*GetLanguagesTranslationsUserOK, error)
	/*
	   GetRoutingLanguage gets language
	*/
	GetRoutingLanguage(ctx context.Context, params *GetRoutingLanguageParams) (*GetRoutingLanguageOK, error)
	/*
	   PostLanguages creates language deprecated
	   This endpoint is deprecated. It has been moved to /routing/languages
	*/
	PostLanguages(ctx context.Context, params *PostLanguagesParams) (*PostLanguagesOK, error)
}

// New creates a new languages API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for languages API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
DeleteLanguage deletes language deprecated

This endpoint is deprecated. It has been moved to /routing/languages/{languageId}
*/
func (a *Client) DeleteLanguage(ctx context.Context, params *DeleteLanguageParams) error {

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLanguage",
		Method:             "DELETE",
		PathPattern:        "/api/v2/languages/{languageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLanguageReader{formats: a.formats},
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
DeleteRoutingLanguage deletes language
*/
func (a *Client) DeleteRoutingLanguage(ctx context.Context, params *DeleteRoutingLanguageParams) error {

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRoutingLanguage",
		Method:             "DELETE",
		PathPattern:        "/api/v2/routing/languages/{languageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteRoutingLanguageReader{formats: a.formats},
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
GetLanguage gets language deprecated

This endpoint is deprecated. It has been moved to /routing/languages/{languageId}
*/
func (a *Client) GetLanguage(ctx context.Context, params *GetLanguageParams) (*GetLanguageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLanguage",
		Method:             "GET",
		PathPattern:        "/api/v2/languages/{languageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLanguageOK), nil

}

/*
GetLanguages gets the list of supported languages deprecated

This endpoint is deprecated. It has been moved to /routing/languages
*/
func (a *Client) GetLanguages(ctx context.Context, params *GetLanguagesParams) (*GetLanguagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLanguages",
		Method:             "GET",
		PathPattern:        "/api/v2/languages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLanguagesOK), nil

}

/*
GetLanguagesTranslations gets all available languages for translation
*/
func (a *Client) GetLanguagesTranslations(ctx context.Context, params *GetLanguagesTranslationsParams) (*GetLanguagesTranslationsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLanguagesTranslations",
		Method:             "GET",
		PathPattern:        "/api/v2/languages/translations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguagesTranslationsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLanguagesTranslationsOK), nil

}

/*
GetLanguagesTranslationsBuiltin gets the builtin translation for a language
*/
func (a *Client) GetLanguagesTranslationsBuiltin(ctx context.Context, params *GetLanguagesTranslationsBuiltinParams) (*GetLanguagesTranslationsBuiltinOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLanguagesTranslationsBuiltin",
		Method:             "GET",
		PathPattern:        "/api/v2/languages/translations/builtin",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguagesTranslationsBuiltinReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLanguagesTranslationsBuiltinOK), nil

}

/*
GetLanguagesTranslationsOrganization gets effective translation for an organization by language
*/
func (a *Client) GetLanguagesTranslationsOrganization(ctx context.Context, params *GetLanguagesTranslationsOrganizationParams) (*GetLanguagesTranslationsOrganizationOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLanguagesTranslationsOrganization",
		Method:             "GET",
		PathPattern:        "/api/v2/languages/translations/organization",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguagesTranslationsOrganizationReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLanguagesTranslationsOrganizationOK), nil

}

/*
GetLanguagesTranslationsUser gets effective language translation for a user
*/
func (a *Client) GetLanguagesTranslationsUser(ctx context.Context, params *GetLanguagesTranslationsUserParams) (*GetLanguagesTranslationsUserOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getLanguagesTranslationsUser",
		Method:             "GET",
		PathPattern:        "/api/v2/languages/translations/users/{userId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetLanguagesTranslationsUserReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetLanguagesTranslationsUserOK), nil

}

/*
GetRoutingLanguage gets language
*/
func (a *Client) GetRoutingLanguage(ctx context.Context, params *GetRoutingLanguageParams) (*GetRoutingLanguageOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRoutingLanguage",
		Method:             "GET",
		PathPattern:        "/api/v2/routing/languages/{languageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRoutingLanguageReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRoutingLanguageOK), nil

}

/*
PostLanguages creates language deprecated

This endpoint is deprecated. It has been moved to /routing/languages
*/
func (a *Client) PostLanguages(ctx context.Context, params *PostLanguagesParams) (*PostLanguagesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postLanguages",
		Method:             "POST",
		PathPattern:        "/api/v2/languages",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostLanguagesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostLanguagesOK), nil

}
