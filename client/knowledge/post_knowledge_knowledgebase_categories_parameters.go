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

// NewPostKnowledgeKnowledgebaseCategoriesParams creates a new PostKnowledgeKnowledgebaseCategoriesParams object
// with the default values initialized.
func NewPostKnowledgeKnowledgebaseCategoriesParams() *PostKnowledgeKnowledgebaseCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseCategoriesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostKnowledgeKnowledgebaseCategoriesParamsWithTimeout creates a new PostKnowledgeKnowledgebaseCategoriesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostKnowledgeKnowledgebaseCategoriesParamsWithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseCategoriesParams{

		timeout: timeout,
	}
}

// NewPostKnowledgeKnowledgebaseCategoriesParamsWithContext creates a new PostKnowledgeKnowledgebaseCategoriesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostKnowledgeKnowledgebaseCategoriesParamsWithContext(ctx context.Context) *PostKnowledgeKnowledgebaseCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseCategoriesParams{

		Context: ctx,
	}
}

// NewPostKnowledgeKnowledgebaseCategoriesParamsWithHTTPClient creates a new PostKnowledgeKnowledgebaseCategoriesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostKnowledgeKnowledgebaseCategoriesParamsWithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseCategoriesParams {
	var ()
	return &PostKnowledgeKnowledgebaseCategoriesParams{
		HTTPClient: client,
	}
}

/*PostKnowledgeKnowledgebaseCategoriesParams contains all the parameters to send to the API endpoint
for the post knowledge knowledgebase categories operation typically these are written to a http.Request
*/
type PostKnowledgeKnowledgebaseCategoriesParams struct {

	/*Body*/
	Body *models.CategoryRequest
	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) WithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseCategoriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) WithContext(ctx context.Context) *PostKnowledgeKnowledgebaseCategoriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) WithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseCategoriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) WithBody(body *models.CategoryRequest) *PostKnowledgeKnowledgebaseCategoriesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) SetBody(body *models.CategoryRequest) {
	o.Body = body
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) WithKnowledgeBaseID(knowledgeBaseID string) *PostKnowledgeKnowledgebaseCategoriesParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the post knowledge knowledgebase categories params
func (o *PostKnowledgeKnowledgebaseCategoriesParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *PostKnowledgeKnowledgebaseCategoriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}