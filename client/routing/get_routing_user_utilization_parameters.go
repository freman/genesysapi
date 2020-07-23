// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingUserUtilizationParams creates a new GetRoutingUserUtilizationParams object
// with the default values initialized.
func NewGetRoutingUserUtilizationParams() *GetRoutingUserUtilizationParams {
	var ()
	return &GetRoutingUserUtilizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingUserUtilizationParamsWithTimeout creates a new GetRoutingUserUtilizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingUserUtilizationParamsWithTimeout(timeout time.Duration) *GetRoutingUserUtilizationParams {
	var ()
	return &GetRoutingUserUtilizationParams{

		timeout: timeout,
	}
}

// NewGetRoutingUserUtilizationParamsWithContext creates a new GetRoutingUserUtilizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingUserUtilizationParamsWithContext(ctx context.Context) *GetRoutingUserUtilizationParams {
	var ()
	return &GetRoutingUserUtilizationParams{

		Context: ctx,
	}
}

// NewGetRoutingUserUtilizationParamsWithHTTPClient creates a new GetRoutingUserUtilizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingUserUtilizationParamsWithHTTPClient(client *http.Client) *GetRoutingUserUtilizationParams {
	var ()
	return &GetRoutingUserUtilizationParams{
		HTTPClient: client,
	}
}

/*GetRoutingUserUtilizationParams contains all the parameters to send to the API endpoint
for the get routing user utilization operation typically these are written to a http.Request
*/
type GetRoutingUserUtilizationParams struct {

	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) WithTimeout(timeout time.Duration) *GetRoutingUserUtilizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) WithContext(ctx context.Context) *GetRoutingUserUtilizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) WithHTTPClient(client *http.Client) *GetRoutingUserUtilizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserID adds the userID to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) WithUserID(userID string) *GetRoutingUserUtilizationParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get routing user utilization params
func (o *GetRoutingUserUtilizationParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingUserUtilizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
