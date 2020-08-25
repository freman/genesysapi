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
)

// NewGetKnowledgeKnowledgebaseLanguageCategoryParams creates a new GetKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized.
func NewGetKnowledgeKnowledgebaseLanguageCategoryParams() *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageCategoryParamsWithTimeout creates a new GetKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeKnowledgebaseLanguageCategoryParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageCategoryParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageCategoryParamsWithContext creates a new GetKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeKnowledgebaseLanguageCategoryParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageCategoryParams{

		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageCategoryParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeKnowledgebaseLanguageCategoryParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageCategoryParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeKnowledgebaseLanguageCategoryParams contains all the parameters to send to the API endpoint
for the get knowledge knowledgebase language category operation typically these are written to a http.Request
*/
type GetKnowledgeKnowledgebaseLanguageCategoryParams struct {

	/*CategoryID
	  Category ID

	*/
	CategoryID string
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

// WithTimeout adds the timeout to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCategoryID adds the categoryID to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) WithCategoryID(categoryID string) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLanguageCode adds the languageCode to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) WithLanguageCode(languageCode string) *GetKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the get knowledge knowledgebase language category params
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseLanguageCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param categoryId
	if err := r.SetPathParam("categoryId", o.CategoryID); err != nil {
		return err
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