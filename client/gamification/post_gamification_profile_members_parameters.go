// Code generated by go-swagger; DO NOT EDIT.

package gamification

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

// NewPostGamificationProfileMembersParams creates a new PostGamificationProfileMembersParams object
// with the default values initialized.
func NewPostGamificationProfileMembersParams() *PostGamificationProfileMembersParams {
	var ()
	return &PostGamificationProfileMembersParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostGamificationProfileMembersParamsWithTimeout creates a new PostGamificationProfileMembersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostGamificationProfileMembersParamsWithTimeout(timeout time.Duration) *PostGamificationProfileMembersParams {
	var ()
	return &PostGamificationProfileMembersParams{

		timeout: timeout,
	}
}

// NewPostGamificationProfileMembersParamsWithContext creates a new PostGamificationProfileMembersParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostGamificationProfileMembersParamsWithContext(ctx context.Context) *PostGamificationProfileMembersParams {
	var ()
	return &PostGamificationProfileMembersParams{

		Context: ctx,
	}
}

// NewPostGamificationProfileMembersParamsWithHTTPClient creates a new PostGamificationProfileMembersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostGamificationProfileMembersParamsWithHTTPClient(client *http.Client) *PostGamificationProfileMembersParams {
	var ()
	return &PostGamificationProfileMembersParams{
		HTTPClient: client,
	}
}

/*PostGamificationProfileMembersParams contains all the parameters to send to the API endpoint
for the post gamification profile members operation typically these are written to a http.Request
*/
type PostGamificationProfileMembersParams struct {

	/*Body
	  assignUsers

	*/
	Body *models.AssignUsers
	/*ProfileID
	  Profile Id

	*/
	ProfileID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) WithTimeout(timeout time.Duration) *PostGamificationProfileMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) WithContext(ctx context.Context) *PostGamificationProfileMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) WithHTTPClient(client *http.Client) *PostGamificationProfileMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) WithBody(body *models.AssignUsers) *PostGamificationProfileMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) SetBody(body *models.AssignUsers) {
	o.Body = body
}

// WithProfileID adds the profileID to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) WithProfileID(profileID string) *PostGamificationProfileMembersParams {
	o.SetProfileID(profileID)
	return o
}

// SetProfileID adds the profileId to the post gamification profile members params
func (o *PostGamificationProfileMembersParams) SetProfileID(profileID string) {
	o.ProfileID = profileID
}

// WriteToRequest writes these params to a swagger request
func (o *PostGamificationProfileMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param profileId
	if err := r.SetPathParam("profileId", o.ProfileID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
