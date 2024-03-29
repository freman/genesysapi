// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewPatchConversationsCallbacksParams creates a new PatchConversationsCallbacksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchConversationsCallbacksParams() *PatchConversationsCallbacksParams {
	return &PatchConversationsCallbacksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchConversationsCallbacksParamsWithTimeout creates a new PatchConversationsCallbacksParams object
// with the ability to set a timeout on a request.
func NewPatchConversationsCallbacksParamsWithTimeout(timeout time.Duration) *PatchConversationsCallbacksParams {
	return &PatchConversationsCallbacksParams{
		timeout: timeout,
	}
}

// NewPatchConversationsCallbacksParamsWithContext creates a new PatchConversationsCallbacksParams object
// with the ability to set a context for a request.
func NewPatchConversationsCallbacksParamsWithContext(ctx context.Context) *PatchConversationsCallbacksParams {
	return &PatchConversationsCallbacksParams{
		Context: ctx,
	}
}

// NewPatchConversationsCallbacksParamsWithHTTPClient creates a new PatchConversationsCallbacksParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchConversationsCallbacksParamsWithHTTPClient(client *http.Client) *PatchConversationsCallbacksParams {
	return &PatchConversationsCallbacksParams{
		HTTPClient: client,
	}
}

/*
PatchConversationsCallbacksParams contains all the parameters to send to the API endpoint

	for the patch conversations callbacks operation.

	Typically these are written to a http.Request.
*/
type PatchConversationsCallbacksParams struct {

	/* Body.

	   PatchCallbackRequest
	*/
	Body *models.PatchCallbackRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch conversations callbacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsCallbacksParams) WithDefaults() *PatchConversationsCallbacksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch conversations callbacks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchConversationsCallbacksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) WithTimeout(timeout time.Duration) *PatchConversationsCallbacksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) WithContext(ctx context.Context) *PatchConversationsCallbacksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) WithHTTPClient(client *http.Client) *PatchConversationsCallbacksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) WithBody(body *models.PatchCallbackRequest) *PatchConversationsCallbacksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch conversations callbacks params
func (o *PatchConversationsCallbacksParams) SetBody(body *models.PatchCallbackRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchConversationsCallbacksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
