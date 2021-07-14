// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewGetOrganizationsLimitsChangerequestParams creates a new GetOrganizationsLimitsChangerequestParams object
// with the default values initialized.
func NewGetOrganizationsLimitsChangerequestParams() *GetOrganizationsLimitsChangerequestParams {
	var ()
	return &GetOrganizationsLimitsChangerequestParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsLimitsChangerequestParamsWithTimeout creates a new GetOrganizationsLimitsChangerequestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOrganizationsLimitsChangerequestParamsWithTimeout(timeout time.Duration) *GetOrganizationsLimitsChangerequestParams {
	var ()
	return &GetOrganizationsLimitsChangerequestParams{

		timeout: timeout,
	}
}

// NewGetOrganizationsLimitsChangerequestParamsWithContext creates a new GetOrganizationsLimitsChangerequestParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOrganizationsLimitsChangerequestParamsWithContext(ctx context.Context) *GetOrganizationsLimitsChangerequestParams {
	var ()
	return &GetOrganizationsLimitsChangerequestParams{

		Context: ctx,
	}
}

// NewGetOrganizationsLimitsChangerequestParamsWithHTTPClient creates a new GetOrganizationsLimitsChangerequestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOrganizationsLimitsChangerequestParamsWithHTTPClient(client *http.Client) *GetOrganizationsLimitsChangerequestParams {
	var ()
	return &GetOrganizationsLimitsChangerequestParams{
		HTTPClient: client,
	}
}

/*GetOrganizationsLimitsChangerequestParams contains all the parameters to send to the API endpoint
for the get organizations limits changerequest operation typically these are written to a http.Request
*/
type GetOrganizationsLimitsChangerequestParams struct {

	/*RequestID
	  Unique id for the limit change request

	*/
	RequestID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) WithTimeout(timeout time.Duration) *GetOrganizationsLimitsChangerequestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) WithContext(ctx context.Context) *GetOrganizationsLimitsChangerequestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) WithHTTPClient(client *http.Client) *GetOrganizationsLimitsChangerequestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRequestID adds the requestID to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) WithRequestID(requestID string) *GetOrganizationsLimitsChangerequestParams {
	o.SetRequestID(requestID)
	return o
}

// SetRequestID adds the requestId to the get organizations limits changerequest params
func (o *GetOrganizationsLimitsChangerequestParams) SetRequestID(requestID string) {
	o.RequestID = requestID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsLimitsChangerequestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param requestId
	if err := r.SetPathParam("requestId", o.RequestID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}