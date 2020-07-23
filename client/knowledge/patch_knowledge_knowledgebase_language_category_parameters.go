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

// NewPatchKnowledgeKnowledgebaseLanguageCategoryParams creates a new PatchKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized.
func NewPatchKnowledgeKnowledgebaseLanguageCategoryParams() *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &PatchKnowledgeKnowledgebaseLanguageCategoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryParamsWithTimeout creates a new PatchKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchKnowledgeKnowledgebaseLanguageCategoryParamsWithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &PatchKnowledgeKnowledgebaseLanguageCategoryParams{

		timeout: timeout,
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryParamsWithContext creates a new PatchKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchKnowledgeKnowledgebaseLanguageCategoryParamsWithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &PatchKnowledgeKnowledgebaseLanguageCategoryParams{

		Context: ctx,
	}
}

// NewPatchKnowledgeKnowledgebaseLanguageCategoryParamsWithHTTPClient creates a new PatchKnowledgeKnowledgebaseLanguageCategoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchKnowledgeKnowledgebaseLanguageCategoryParamsWithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	var ()
	return &PatchKnowledgeKnowledgebaseLanguageCategoryParams{
		HTTPClient: client,
	}
}

/*PatchKnowledgeKnowledgebaseLanguageCategoryParams contains all the parameters to send to the API endpoint
for the patch knowledge knowledgebase language category operation typically these are written to a http.Request
*/
type PatchKnowledgeKnowledgebaseLanguageCategoryParams struct {

	/*Body*/
	Body *models.KnowledgeCategoryRequest
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

// WithTimeout adds the timeout to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WithTimeout(timeout time.Duration) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WithContext(ctx context.Context) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WithHTTPClient(client *http.Client) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WithBody(body *models.KnowledgeCategoryRequest) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) SetBody(body *models.KnowledgeCategoryRequest) {
	o.Body = body
}

// WithCategoryID adds the categoryID to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WithCategoryID(categoryID string) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetCategoryID(categoryID)
	return o
}

// SetCategoryID adds the categoryId to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) SetCategoryID(categoryID string) {
	o.CategoryID = categoryID
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WithKnowledgeBaseID(knowledgeBaseID string) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLanguageCode adds the languageCode to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WithLanguageCode(languageCode string) *PatchKnowledgeKnowledgebaseLanguageCategoryParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the patch knowledge knowledgebase language category params
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WriteToRequest writes these params to a swagger request
func (o *PatchKnowledgeKnowledgebaseLanguageCategoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
