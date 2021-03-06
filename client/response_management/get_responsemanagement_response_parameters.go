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

// NewGetResponsemanagementResponseParams creates a new GetResponsemanagementResponseParams object
// with the default values initialized.
func NewGetResponsemanagementResponseParams() *GetResponsemanagementResponseParams {
	var ()
	return &GetResponsemanagementResponseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResponsemanagementResponseParamsWithTimeout creates a new GetResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResponsemanagementResponseParamsWithTimeout(timeout time.Duration) *GetResponsemanagementResponseParams {
	var ()
	return &GetResponsemanagementResponseParams{

		timeout: timeout,
	}
}

// NewGetResponsemanagementResponseParamsWithContext creates a new GetResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResponsemanagementResponseParamsWithContext(ctx context.Context) *GetResponsemanagementResponseParams {
	var ()
	return &GetResponsemanagementResponseParams{

		Context: ctx,
	}
}

// NewGetResponsemanagementResponseParamsWithHTTPClient creates a new GetResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResponsemanagementResponseParamsWithHTTPClient(client *http.Client) *GetResponsemanagementResponseParams {
	var ()
	return &GetResponsemanagementResponseParams{
		HTTPClient: client,
	}
}

/*GetResponsemanagementResponseParams contains all the parameters to send to the API endpoint
for the get responsemanagement response operation typically these are written to a http.Request
*/
type GetResponsemanagementResponseParams struct {

	/*Expand
	  Expand instructions for the return value.

	*/
	Expand *string
	/*ResponseID
	  Response ID

	*/
	ResponseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) WithTimeout(timeout time.Duration) *GetResponsemanagementResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) WithContext(ctx context.Context) *GetResponsemanagementResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) WithHTTPClient(client *http.Client) *GetResponsemanagementResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) WithExpand(expand *string) *GetResponsemanagementResponseParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithResponseID adds the responseID to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) WithResponseID(responseID string) *GetResponsemanagementResponseParams {
	o.SetResponseID(responseID)
	return o
}

// SetResponseID adds the responseId to the get responsemanagement response params
func (o *GetResponsemanagementResponseParams) SetResponseID(responseID string) {
	o.ResponseID = responseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResponsemanagementResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// query param expand
		var qrExpand string
		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {
			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}

	}

	// path param responseId
	if err := r.SetPathParam("responseId", o.ResponseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
