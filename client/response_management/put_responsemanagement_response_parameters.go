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

	"github.com/freman/genesysapi/models"
)

// NewPutResponsemanagementResponseParams creates a new PutResponsemanagementResponseParams object
// with the default values initialized.
func NewPutResponsemanagementResponseParams() *PutResponsemanagementResponseParams {
	var ()
	return &PutResponsemanagementResponseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutResponsemanagementResponseParamsWithTimeout creates a new PutResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutResponsemanagementResponseParamsWithTimeout(timeout time.Duration) *PutResponsemanagementResponseParams {
	var ()
	return &PutResponsemanagementResponseParams{

		timeout: timeout,
	}
}

// NewPutResponsemanagementResponseParamsWithContext creates a new PutResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutResponsemanagementResponseParamsWithContext(ctx context.Context) *PutResponsemanagementResponseParams {
	var ()
	return &PutResponsemanagementResponseParams{

		Context: ctx,
	}
}

// NewPutResponsemanagementResponseParamsWithHTTPClient creates a new PutResponsemanagementResponseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutResponsemanagementResponseParamsWithHTTPClient(client *http.Client) *PutResponsemanagementResponseParams {
	var ()
	return &PutResponsemanagementResponseParams{
		HTTPClient: client,
	}
}

/*PutResponsemanagementResponseParams contains all the parameters to send to the API endpoint
for the put responsemanagement response operation typically these are written to a http.Request
*/
type PutResponsemanagementResponseParams struct {

	/*Body
	  Response

	*/
	Body *models.Response
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

// WithTimeout adds the timeout to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) WithTimeout(timeout time.Duration) *PutResponsemanagementResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) WithContext(ctx context.Context) *PutResponsemanagementResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) WithHTTPClient(client *http.Client) *PutResponsemanagementResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) WithBody(body *models.Response) *PutResponsemanagementResponseParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) SetBody(body *models.Response) {
	o.Body = body
}

// WithExpand adds the expand to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) WithExpand(expand *string) *PutResponsemanagementResponseParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithResponseID adds the responseID to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) WithResponseID(responseID string) *PutResponsemanagementResponseParams {
	o.SetResponseID(responseID)
	return o
}

// SetResponseID adds the responseId to the put responsemanagement response params
func (o *PutResponsemanagementResponseParams) SetResponseID(responseID string) {
	o.ResponseID = responseID
}

// WriteToRequest writes these params to a swagger request
func (o *PutResponsemanagementResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
