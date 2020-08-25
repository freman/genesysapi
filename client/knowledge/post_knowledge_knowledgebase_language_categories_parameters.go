// Code generated by go-swagger; DO NOT EDIT.

package knowledge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// NewPostKnowledgeKnowledgebaseLanguageCategoriesParams creates a new PostKnowledgeKnowledgebaseLanguageCategoriesParams object
// with the default values initialized.
func NewPostKnowledgeKnowledgebaseLanguageCategoriesParams() *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesParamsWithTimeout creates a new PostKnowledgeKnowledgebaseLanguageCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostKnowledgeKnowledgebaseLanguageCategoriesParamsWithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageCategoriesParams{

		timeout: timeout,
	}
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesParamsWithContext creates a new PostKnowledgeKnowledgebaseLanguageCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostKnowledgeKnowledgebaseLanguageCategoriesParamsWithContext(ctx context.Context) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageCategoriesParams{

		Context: ctx,
	}
}

// NewPostKnowledgeKnowledgebaseLanguageCategoriesParamsWithHTTPClient creates a new PostKnowledgeKnowledgebaseLanguageCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostKnowledgeKnowledgebaseLanguageCategoriesParamsWithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageCategoriesParams{
		HTTPClient: client,
	}
}

/*PostKnowledgeKnowledgebaseLanguageCategoriesParams contains all the parameters to send to the API endpoint
for the post knowledge knowledgebase language categories operation typically these are written to a http.Request
*/
type PostKnowledgeKnowledgebaseLanguageCategoriesParams struct {

	/*Body*/
	Body *models.KnowledgeCategoryRequest
	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string
	/*LanguageCode
	  Language code, format: iso2-LOCALE

	*/
	LanguageCode string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) WithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) WithContext(ctx context.Context) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) WithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) WithBody(body *models.KnowledgeCategoryRequest) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) SetBody(body *models.KnowledgeCategoryRequest) {
	o.Body = body
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) WithKnowledgeBaseID(knowledgeBaseID string) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLanguageCode adds the languageCode to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) WithLanguageCode(languageCode string) *PostKnowledgeKnowledgebaseLanguageCategoriesParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the post knowledge knowledgebase language categories params
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WriteToRequest writes these params to a swagger request
func (o *PostKnowledgeKnowledgebaseLanguageCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	// path param languageCode
	if err := r.SetPathParam("languageCode", o.LanguageCode); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}