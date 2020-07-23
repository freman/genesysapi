// Code generated by go-swagger; DO NOT EDIT.

package response_management

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

// NewDeleteResponsemanagementResponseParams creates a new DeleteResponsemanagementResponseParams object
// with the default values initialized.
func NewDeleteResponsemanagementResponseParams() *DeleteResponsemanagementResponseParams {
	var ()
	return &DeleteResponsemanagementResponseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteResponsemanagementResponseParamsWithTimeout creates a new DeleteResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteResponsemanagementResponseParamsWithTimeout(timeout time.Duration) *DeleteResponsemanagementResponseParams {
	var ()
	return &DeleteResponsemanagementResponseParams{

		timeout: timeout,
	}
}

// NewDeleteResponsemanagementResponseParamsWithContext creates a new DeleteResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteResponsemanagementResponseParamsWithContext(ctx context.Context) *DeleteResponsemanagementResponseParams {
	var ()
	return &DeleteResponsemanagementResponseParams{

		Context: ctx,
	}
}

// NewDeleteResponsemanagementResponseParamsWithHTTPClient creates a new DeleteResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteResponsemanagementResponseParamsWithHTTPClient(client *http.Client) *DeleteResponsemanagementResponseParams {
	var ()
	return &DeleteResponsemanagementResponseParams{
		HTTPClient: client,
	}
}

/*DeleteResponsemanagementResponseParams contains all the parameters to send to the API endpoint
for the delete responsemanagement response operation typically these are written to a http.Request
*/
type DeleteResponsemanagementResponseParams struct {

	/*ResponseID
	  Response ID

	*/
	ResponseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) WithTimeout(timeout time.Duration) *DeleteResponsemanagementResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) WithContext(ctx context.Context) *DeleteResponsemanagementResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) WithHTTPClient(client *http.Client) *DeleteResponsemanagementResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResponseID adds the responseID to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) WithResponseID(responseID string) *DeleteResponsemanagementResponseParams {
	o.SetResponseID(responseID)
	return o
}

// SetResponseID adds the responseId to the delete responsemanagement response params
func (o *DeleteResponsemanagementResponseParams) SetResponseID(responseID string) {
	o.ResponseID = responseID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteResponsemanagementResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param responseId
	if err := r.SetPathParam("responseId", o.ResponseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
