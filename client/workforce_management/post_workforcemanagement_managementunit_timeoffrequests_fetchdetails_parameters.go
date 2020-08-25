// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

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

// NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams creates a new PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams object
// with the default values initialized.
func NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams() *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParamsWithTimeout creates a new PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParamsWithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams{

		timeout: timeout,
	}
}

// NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParamsWithContext creates a new PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParamsWithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams{

		Context: ctx,
	}
}

// NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParamsWithHTTPClient creates a new PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParamsWithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	var ()
	return &PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams{
		HTTPClient: client,
	}
}

/*PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams contains all the parameters to send to the API endpoint
for the post workforcemanagement managementunit timeoffrequests fetchdetails operation typically these are written to a http.Request
*/
type PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams struct {

	/*Body
	  body

	*/
	Body *models.TimeOffRequestListing
	/*MuID
	  The muId of the management unit, or 'mine' for the management unit of the logged-in user.

	*/
	MuID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) WithTimeout(timeout time.Duration) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) WithContext(ctx context.Context) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) WithHTTPClient(client *http.Client) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) WithBody(body *models.TimeOffRequestListing) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) SetBody(body *models.TimeOffRequestListing) {
	o.Body = body
}

// WithMuID adds the muID to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) WithMuID(muID string) *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams {
	o.SetMuID(muID)
	return o
}

// SetMuID adds the muId to the post workforcemanagement managementunit timeoffrequests fetchdetails params
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) SetMuID(muID string) {
	o.MuID = muID
}

// WriteToRequest writes these params to a swagger request
func (o *PostWorkforcemanagementManagementunitTimeoffrequestsFetchdetailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param muId
	if err := r.SetPathParam("muId", o.MuID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}