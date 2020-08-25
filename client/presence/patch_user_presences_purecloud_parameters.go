// Code generated by go-swagger; DO NOT EDIT.

package presence

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

// NewPatchUserPresencesPurecloudParams creates a new PatchUserPresencesPurecloudParams object
// with the default values initialized.
func NewPatchUserPresencesPurecloudParams() *PatchUserPresencesPurecloudParams {
	var ()
	return &PatchUserPresencesPurecloudParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUserPresencesPurecloudParamsWithTimeout creates a new PatchUserPresencesPurecloudParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchUserPresencesPurecloudParamsWithTimeout(timeout time.Duration) *PatchUserPresencesPurecloudParams {
	var ()
	return &PatchUserPresencesPurecloudParams{

		timeout: timeout,
	}
}

// NewPatchUserPresencesPurecloudParamsWithContext creates a new PatchUserPresencesPurecloudParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchUserPresencesPurecloudParamsWithContext(ctx context.Context) *PatchUserPresencesPurecloudParams {
	var ()
	return &PatchUserPresencesPurecloudParams{

		Context: ctx,
	}
}

// NewPatchUserPresencesPurecloudParamsWithHTTPClient creates a new PatchUserPresencesPurecloudParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchUserPresencesPurecloudParamsWithHTTPClient(client *http.Client) *PatchUserPresencesPurecloudParams {
	var ()
	return &PatchUserPresencesPurecloudParams{
		HTTPClient: client,
	}
}

/*PatchUserPresencesPurecloudParams contains all the parameters to send to the API endpoint
for the patch user presences purecloud operation typically these are written to a http.Request
*/
type PatchUserPresencesPurecloudParams struct {

	/*Body
	  User presence

	*/
	Body *models.UserPresence
	/*UserID
	  user Id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) WithTimeout(timeout time.Duration) *PatchUserPresencesPurecloudParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) WithContext(ctx context.Context) *PatchUserPresencesPurecloudParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) WithHTTPClient(client *http.Client) *PatchUserPresencesPurecloudParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) WithBody(body *models.UserPresence) *PatchUserPresencesPurecloudParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) SetBody(body *models.UserPresence) {
	o.Body = body
}

// WithUserID adds the userID to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) WithUserID(userID string) *PatchUserPresencesPurecloudParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch user presences purecloud params
func (o *PatchUserPresencesPurecloudParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUserPresencesPurecloudParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}