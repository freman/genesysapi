// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewPatchTeamParams creates a new PatchTeamParams object
// with the default values initialized.
func NewPatchTeamParams() *PatchTeamParams {
	var ()
	return &PatchTeamParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchTeamParamsWithTimeout creates a new PatchTeamParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchTeamParamsWithTimeout(timeout time.Duration) *PatchTeamParams {
	var ()
	return &PatchTeamParams{

		timeout: timeout,
	}
}

// NewPatchTeamParamsWithContext creates a new PatchTeamParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchTeamParamsWithContext(ctx context.Context) *PatchTeamParams {
	var ()
	return &PatchTeamParams{

		Context: ctx,
	}
}

// NewPatchTeamParamsWithHTTPClient creates a new PatchTeamParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchTeamParamsWithHTTPClient(client *http.Client) *PatchTeamParams {
	var ()
	return &PatchTeamParams{
		HTTPClient: client,
	}
}

/*PatchTeamParams contains all the parameters to send to the API endpoint
for the patch team operation typically these are written to a http.Request
*/
type PatchTeamParams struct {

	/*Body
	  Team

	*/
	Body *models.Team
	/*TeamID
	  Team ID

	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch team params
func (o *PatchTeamParams) WithTimeout(timeout time.Duration) *PatchTeamParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch team params
func (o *PatchTeamParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch team params
func (o *PatchTeamParams) WithContext(ctx context.Context) *PatchTeamParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch team params
func (o *PatchTeamParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch team params
func (o *PatchTeamParams) WithHTTPClient(client *http.Client) *PatchTeamParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch team params
func (o *PatchTeamParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch team params
func (o *PatchTeamParams) WithBody(body *models.Team) *PatchTeamParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch team params
func (o *PatchTeamParams) SetBody(body *models.Team) {
	o.Body = body
}

// WithTeamID adds the teamID to the patch team params
func (o *PatchTeamParams) WithTeamID(teamID string) *PatchTeamParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the patch team params
func (o *PatchTeamParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchTeamParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param teamId
	if err := r.SetPathParam("teamId", o.TeamID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
