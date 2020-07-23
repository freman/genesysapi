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
	"github.com/go-openapi/swag"
)

// NewDeleteArchitectPromptsParams creates a new DeleteArchitectPromptsParams object
// with the default values initialized.
func NewDeleteArchitectPromptsParams() *DeleteArchitectPromptsParams {
	var ()
	return &DeleteArchitectPromptsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteArchitectPromptsParamsWithTimeout creates a new DeleteArchitectPromptsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteArchitectPromptsParamsWithTimeout(timeout time.Duration) *DeleteArchitectPromptsParams {
	var ()
	return &DeleteArchitectPromptsParams{

		timeout: timeout,
	}
}

// NewDeleteArchitectPromptsParamsWithContext creates a new DeleteArchitectPromptsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteArchitectPromptsParamsWithContext(ctx context.Context) *DeleteArchitectPromptsParams {
	var ()
	return &DeleteArchitectPromptsParams{

		Context: ctx,
	}
}

// NewDeleteArchitectPromptsParamsWithHTTPClient creates a new DeleteArchitectPromptsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteArchitectPromptsParamsWithHTTPClient(client *http.Client) *DeleteArchitectPromptsParams {
	var ()
	return &DeleteArchitectPromptsParams{
		HTTPClient: client,
	}
}

/*DeleteArchitectPromptsParams contains all the parameters to send to the API endpoint
for the delete architect prompts operation typically these are written to a http.Request
*/
type DeleteArchitectPromptsParams struct {

	/*ID
	  List of Prompt IDs

	*/
	ID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithTimeout(timeout time.Duration) *DeleteArchitectPromptsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithContext(ctx context.Context) *DeleteArchitectPromptsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithHTTPClient(client *http.Client) *DeleteArchitectPromptsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) WithID(id []string) *DeleteArchitectPromptsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete architect prompts params
func (o *DeleteArchitectPromptsParams) SetID(id []string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteArchitectPromptsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
