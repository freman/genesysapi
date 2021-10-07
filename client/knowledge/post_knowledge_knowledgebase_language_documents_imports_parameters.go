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

// NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParams creates a new PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams object
// with the default values initialized.
func NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParams() *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParamsWithTimeout creates a new PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParamsWithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams{

		timeout: timeout,
	}
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParamsWithContext creates a new PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParamsWithContext(ctx context.Context) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams{

		Context: ctx,
	}
}

// NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParamsWithHTTPClient creates a new PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostKnowledgeKnowledgebaseLanguageDocumentsImportsParamsWithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	var ()
	return &PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams{
		HTTPClient: client,
	}
}

/*PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams contains all the parameters to send to the API endpoint
for the post knowledge knowledgebase language documents imports operation typically these are written to a http.Request
*/
type PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams struct {

	/*Body*/
	Body *models.KnowledgeImport
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

// WithTimeout adds the timeout to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) WithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) WithContext(ctx context.Context) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) WithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) WithBody(body *models.KnowledgeImport) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) SetBody(body *models.KnowledgeImport) {
	o.Body = body
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) WithKnowledgeBaseID(knowledgeBaseID string) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLanguageCode adds the languageCode to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) WithLanguageCode(languageCode string) *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the post knowledge knowledgebase language documents imports params
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WriteToRequest writes these params to a swagger request
func (o *PostKnowledgeKnowledgebaseLanguageDocumentsImportsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
