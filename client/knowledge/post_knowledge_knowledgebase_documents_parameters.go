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

// NewPostKnowledgeKnowledgebaseDocumentsParams creates a new PostKnowledgeKnowledgebaseDocumentsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostKnowledgeKnowledgebaseDocumentsParams() *PostKnowledgeKnowledgebaseDocumentsParams {
	return &PostKnowledgeKnowledgebaseDocumentsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostKnowledgeKnowledgebaseDocumentsParamsWithTimeout creates a new PostKnowledgeKnowledgebaseDocumentsParams object
// with the ability to set a timeout on a request.
func NewPostKnowledgeKnowledgebaseDocumentsParamsWithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseDocumentsParams {
	return &PostKnowledgeKnowledgebaseDocumentsParams{
		timeout: timeout,
	}
}

// NewPostKnowledgeKnowledgebaseDocumentsParamsWithContext creates a new PostKnowledgeKnowledgebaseDocumentsParams object
// with the ability to set a context for a request.
func NewPostKnowledgeKnowledgebaseDocumentsParamsWithContext(ctx context.Context) *PostKnowledgeKnowledgebaseDocumentsParams {
	return &PostKnowledgeKnowledgebaseDocumentsParams{
		Context: ctx,
	}
}

// NewPostKnowledgeKnowledgebaseDocumentsParamsWithHTTPClient creates a new PostKnowledgeKnowledgebaseDocumentsParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostKnowledgeKnowledgebaseDocumentsParamsWithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseDocumentsParams {
	return &PostKnowledgeKnowledgebaseDocumentsParams{
		HTTPClient: client,
	}
}

/*
PostKnowledgeKnowledgebaseDocumentsParams contains all the parameters to send to the API endpoint

	for the post knowledge knowledgebase documents operation.

	Typically these are written to a http.Request.
*/
type PostKnowledgeKnowledgebaseDocumentsParams struct {

	// Body.
	Body *models.KnowledgeDocumentReq

	/* KnowledgeBaseID.

	   Knowledge base ID
	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post knowledge knowledgebase documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostKnowledgeKnowledgebaseDocumentsParams) WithDefaults() *PostKnowledgeKnowledgebaseDocumentsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post knowledge knowledgebase documents params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostKnowledgeKnowledgebaseDocumentsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) WithTimeout(timeout time.Duration) *PostKnowledgeKnowledgebaseDocumentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) WithContext(ctx context.Context) *PostKnowledgeKnowledgebaseDocumentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) WithHTTPClient(client *http.Client) *PostKnowledgeKnowledgebaseDocumentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) WithBody(body *models.KnowledgeDocumentReq) *PostKnowledgeKnowledgebaseDocumentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) SetBody(body *models.KnowledgeDocumentReq) {
	o.Body = body
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) WithKnowledgeBaseID(knowledgeBaseID string) *PostKnowledgeKnowledgebaseDocumentsParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the post knowledge knowledgebase documents params
func (o *PostKnowledgeKnowledgebaseDocumentsParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *PostKnowledgeKnowledgebaseDocumentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
