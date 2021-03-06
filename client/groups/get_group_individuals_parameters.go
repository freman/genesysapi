// Code generated by go-swagger; DO NOT EDIT.

package groups

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

// NewGetGroupIndividualsParams creates a new GetGroupIndividualsParams object
// with the default values initialized.
func NewGetGroupIndividualsParams() *GetGroupIndividualsParams {
	var ()
	return &GetGroupIndividualsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupIndividualsParamsWithTimeout creates a new GetGroupIndividualsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupIndividualsParamsWithTimeout(timeout time.Duration) *GetGroupIndividualsParams {
	var ()
	return &GetGroupIndividualsParams{

		timeout: timeout,
	}
}

// NewGetGroupIndividualsParamsWithContext creates a new GetGroupIndividualsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupIndividualsParamsWithContext(ctx context.Context) *GetGroupIndividualsParams {
	var ()
	return &GetGroupIndividualsParams{

		Context: ctx,
	}
}

// NewGetGroupIndividualsParamsWithHTTPClient creates a new GetGroupIndividualsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupIndividualsParamsWithHTTPClient(client *http.Client) *GetGroupIndividualsParams {
	var ()
	return &GetGroupIndividualsParams{
		HTTPClient: client,
	}
}

/*GetGroupIndividualsParams contains all the parameters to send to the API endpoint
for the get group individuals operation typically these are written to a http.Request
*/
type GetGroupIndividualsParams struct {

	/*GroupID
	  Group ID

	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get group individuals params
func (o *GetGroupIndividualsParams) WithTimeout(timeout time.Duration) *GetGroupIndividualsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get group individuals params
func (o *GetGroupIndividualsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get group individuals params
func (o *GetGroupIndividualsParams) WithContext(ctx context.Context) *GetGroupIndividualsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get group individuals params
func (o *GetGroupIndividualsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get group individuals params
func (o *GetGroupIndividualsParams) WithHTTPClient(client *http.Client) *GetGroupIndividualsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get group individuals params
func (o *GetGroupIndividualsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get group individuals params
func (o *GetGroupIndividualsParams) WithGroupID(groupID string) *GetGroupIndividualsParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get group individuals params
func (o *GetGroupIndividualsParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupIndividualsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
