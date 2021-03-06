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

// NewPatchUserPresenceParams creates a new PatchUserPresenceParams object
// with the default values initialized.
func NewPatchUserPresenceParams() *PatchUserPresenceParams {
	var ()
	return &PatchUserPresenceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchUserPresenceParamsWithTimeout creates a new PatchUserPresenceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchUserPresenceParamsWithTimeout(timeout time.Duration) *PatchUserPresenceParams {
	var ()
	return &PatchUserPresenceParams{

		timeout: timeout,
	}
}

// NewPatchUserPresenceParamsWithContext creates a new PatchUserPresenceParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchUserPresenceParamsWithContext(ctx context.Context) *PatchUserPresenceParams {
	var ()
	return &PatchUserPresenceParams{

		Context: ctx,
	}
}

// NewPatchUserPresenceParamsWithHTTPClient creates a new PatchUserPresenceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchUserPresenceParamsWithHTTPClient(client *http.Client) *PatchUserPresenceParams {
	var ()
	return &PatchUserPresenceParams{
		HTTPClient: client,
	}
}

/*PatchUserPresenceParams contains all the parameters to send to the API endpoint
for the patch user presence operation typically these are written to a http.Request
*/
type PatchUserPresenceParams struct {

	/*Body
	  User presence

	*/
	Body *models.UserPresence
	/*SourceID
	  Presence source ID

	*/
	SourceID string
	/*UserID
	  user Id

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch user presence params
func (o *PatchUserPresenceParams) WithTimeout(timeout time.Duration) *PatchUserPresenceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch user presence params
func (o *PatchUserPresenceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch user presence params
func (o *PatchUserPresenceParams) WithContext(ctx context.Context) *PatchUserPresenceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch user presence params
func (o *PatchUserPresenceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch user presence params
func (o *PatchUserPresenceParams) WithHTTPClient(client *http.Client) *PatchUserPresenceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch user presence params
func (o *PatchUserPresenceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch user presence params
func (o *PatchUserPresenceParams) WithBody(body *models.UserPresence) *PatchUserPresenceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch user presence params
func (o *PatchUserPresenceParams) SetBody(body *models.UserPresence) {
	o.Body = body
}

// WithSourceID adds the sourceID to the patch user presence params
func (o *PatchUserPresenceParams) WithSourceID(sourceID string) *PatchUserPresenceParams {
	o.SetSourceID(sourceID)
	return o
}

// SetSourceID adds the sourceId to the patch user presence params
func (o *PatchUserPresenceParams) SetSourceID(sourceID string) {
	o.SourceID = sourceID
}

// WithUserID adds the userID to the patch user presence params
func (o *PatchUserPresenceParams) WithUserID(userID string) *PatchUserPresenceParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the patch user presence params
func (o *PatchUserPresenceParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchUserPresenceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param sourceId
	if err := r.SetPathParam("sourceId", o.SourceID); err != nil {
		return err
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
