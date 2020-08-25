// Code generated by go-swagger; DO NOT EDIT.

package voicemail

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

// NewPatchVoicemailGroupPolicyParams creates a new PatchVoicemailGroupPolicyParams object
// with the default values initialized.
func NewPatchVoicemailGroupPolicyParams() *PatchVoicemailGroupPolicyParams {
	var ()
	return &PatchVoicemailGroupPolicyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVoicemailGroupPolicyParamsWithTimeout creates a new PatchVoicemailGroupPolicyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchVoicemailGroupPolicyParamsWithTimeout(timeout time.Duration) *PatchVoicemailGroupPolicyParams {
	var ()
	return &PatchVoicemailGroupPolicyParams{

		timeout: timeout,
	}
}

// NewPatchVoicemailGroupPolicyParamsWithContext creates a new PatchVoicemailGroupPolicyParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchVoicemailGroupPolicyParamsWithContext(ctx context.Context) *PatchVoicemailGroupPolicyParams {
	var ()
	return &PatchVoicemailGroupPolicyParams{

		Context: ctx,
	}
}

// NewPatchVoicemailGroupPolicyParamsWithHTTPClient creates a new PatchVoicemailGroupPolicyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchVoicemailGroupPolicyParamsWithHTTPClient(client *http.Client) *PatchVoicemailGroupPolicyParams {
	var ()
	return &PatchVoicemailGroupPolicyParams{
		HTTPClient: client,
	}
}

/*PatchVoicemailGroupPolicyParams contains all the parameters to send to the API endpoint
for the patch voicemail group policy operation typically these are written to a http.Request
*/
type PatchVoicemailGroupPolicyParams struct {

	/*Body
	  The group's voicemail policy

	*/
	Body *models.VoicemailGroupPolicy
	/*GroupID
	  Group ID

	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) WithTimeout(timeout time.Duration) *PatchVoicemailGroupPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) WithContext(ctx context.Context) *PatchVoicemailGroupPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) WithHTTPClient(client *http.Client) *PatchVoicemailGroupPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) WithBody(body *models.VoicemailGroupPolicy) *PatchVoicemailGroupPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) SetBody(body *models.VoicemailGroupPolicy) {
	o.Body = body
}

// WithGroupID adds the groupID to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) WithGroupID(groupID string) *PatchVoicemailGroupPolicyParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the patch voicemail group policy params
func (o *PatchVoicemailGroupPolicyParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVoicemailGroupPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}