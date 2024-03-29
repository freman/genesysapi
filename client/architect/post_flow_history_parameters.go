// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewPostFlowHistoryParams creates a new PostFlowHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostFlowHistoryParams() *PostFlowHistoryParams {
	return &PostFlowHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostFlowHistoryParamsWithTimeout creates a new PostFlowHistoryParams object
// with the ability to set a timeout on a request.
func NewPostFlowHistoryParamsWithTimeout(timeout time.Duration) *PostFlowHistoryParams {
	return &PostFlowHistoryParams{
		timeout: timeout,
	}
}

// NewPostFlowHistoryParamsWithContext creates a new PostFlowHistoryParams object
// with the ability to set a context for a request.
func NewPostFlowHistoryParamsWithContext(ctx context.Context) *PostFlowHistoryParams {
	return &PostFlowHistoryParams{
		Context: ctx,
	}
}

// NewPostFlowHistoryParamsWithHTTPClient creates a new PostFlowHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostFlowHistoryParamsWithHTTPClient(client *http.Client) *PostFlowHistoryParams {
	return &PostFlowHistoryParams{
		HTTPClient: client,
	}
}

/*
PostFlowHistoryParams contains all the parameters to send to the API endpoint

	for the post flow history operation.

	Typically these are written to a http.Request.
*/
type PostFlowHistoryParams struct {

	/* FlowID.

	   Flow ID
	*/
	FlowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post flow history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFlowHistoryParams) WithDefaults() *PostFlowHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post flow history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostFlowHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post flow history params
func (o *PostFlowHistoryParams) WithTimeout(timeout time.Duration) *PostFlowHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post flow history params
func (o *PostFlowHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post flow history params
func (o *PostFlowHistoryParams) WithContext(ctx context.Context) *PostFlowHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post flow history params
func (o *PostFlowHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post flow history params
func (o *PostFlowHistoryParams) WithHTTPClient(client *http.Client) *PostFlowHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post flow history params
func (o *PostFlowHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowID adds the flowID to the post flow history params
func (o *PostFlowHistoryParams) WithFlowID(flowID string) *PostFlowHistoryParams {
	o.SetFlowID(flowID)
	return o
}

// SetFlowID adds the flowId to the post flow history params
func (o *PostFlowHistoryParams) SetFlowID(flowID string) {
	o.FlowID = flowID
}

// WriteToRequest writes these params to a swagger request
func (o *PostFlowHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param flowId
	if err := r.SetPathParam("flowId", o.FlowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
