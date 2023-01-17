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

// NewPostArchitectPromptHistoryParams creates a new PostArchitectPromptHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostArchitectPromptHistoryParams() *PostArchitectPromptHistoryParams {
	return &PostArchitectPromptHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostArchitectPromptHistoryParamsWithTimeout creates a new PostArchitectPromptHistoryParams object
// with the ability to set a timeout on a request.
func NewPostArchitectPromptHistoryParamsWithTimeout(timeout time.Duration) *PostArchitectPromptHistoryParams {
	return &PostArchitectPromptHistoryParams{
		timeout: timeout,
	}
}

// NewPostArchitectPromptHistoryParamsWithContext creates a new PostArchitectPromptHistoryParams object
// with the ability to set a context for a request.
func NewPostArchitectPromptHistoryParamsWithContext(ctx context.Context) *PostArchitectPromptHistoryParams {
	return &PostArchitectPromptHistoryParams{
		Context: ctx,
	}
}

// NewPostArchitectPromptHistoryParamsWithHTTPClient creates a new PostArchitectPromptHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostArchitectPromptHistoryParamsWithHTTPClient(client *http.Client) *PostArchitectPromptHistoryParams {
	return &PostArchitectPromptHistoryParams{
		HTTPClient: client,
	}
}

/*
PostArchitectPromptHistoryParams contains all the parameters to send to the API endpoint

	for the post architect prompt history operation.

	Typically these are written to a http.Request.
*/
type PostArchitectPromptHistoryParams struct {

	/* PromptID.

	   Prompt ID
	*/
	PromptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post architect prompt history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectPromptHistoryParams) WithDefaults() *PostArchitectPromptHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post architect prompt history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostArchitectPromptHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) WithTimeout(timeout time.Duration) *PostArchitectPromptHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) WithContext(ctx context.Context) *PostArchitectPromptHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) WithHTTPClient(client *http.Client) *PostArchitectPromptHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPromptID adds the promptID to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) WithPromptID(promptID string) *PostArchitectPromptHistoryParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the post architect prompt history params
func (o *PostArchitectPromptHistoryParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WriteToRequest writes these params to a swagger request
func (o *PostArchitectPromptHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param promptId
	if err := r.SetPathParam("promptId", o.PromptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
