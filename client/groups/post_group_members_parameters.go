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

	"github.com/freman/genesysapi/models"
)

// NewPostGroupMembersParams creates a new PostGroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostGroupMembersParams() *PostGroupMembersParams {
	return &PostGroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostGroupMembersParamsWithTimeout creates a new PostGroupMembersParams object
// with the ability to set a timeout on a request.
func NewPostGroupMembersParamsWithTimeout(timeout time.Duration) *PostGroupMembersParams {
	return &PostGroupMembersParams{
		timeout: timeout,
	}
}

// NewPostGroupMembersParamsWithContext creates a new PostGroupMembersParams object
// with the ability to set a context for a request.
func NewPostGroupMembersParamsWithContext(ctx context.Context) *PostGroupMembersParams {
	return &PostGroupMembersParams{
		Context: ctx,
	}
}

// NewPostGroupMembersParamsWithHTTPClient creates a new PostGroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostGroupMembersParamsWithHTTPClient(client *http.Client) *PostGroupMembersParams {
	return &PostGroupMembersParams{
		HTTPClient: client,
	}
}

/*
PostGroupMembersParams contains all the parameters to send to the API endpoint

	for the post group members operation.

	Typically these are written to a http.Request.
*/
type PostGroupMembersParams struct {

	/* Body.

	   Add members
	*/
	Body *models.GroupMembersUpdate

	/* GroupID.

	   Group ID
	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGroupMembersParams) WithDefaults() *PostGroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post group members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostGroupMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post group members params
func (o *PostGroupMembersParams) WithTimeout(timeout time.Duration) *PostGroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post group members params
func (o *PostGroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post group members params
func (o *PostGroupMembersParams) WithContext(ctx context.Context) *PostGroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post group members params
func (o *PostGroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post group members params
func (o *PostGroupMembersParams) WithHTTPClient(client *http.Client) *PostGroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post group members params
func (o *PostGroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post group members params
func (o *PostGroupMembersParams) WithBody(body *models.GroupMembersUpdate) *PostGroupMembersParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post group members params
func (o *PostGroupMembersParams) SetBody(body *models.GroupMembersUpdate) {
	o.Body = body
}

// WithGroupID adds the groupID to the post group members params
func (o *PostGroupMembersParams) WithGroupID(groupID string) *PostGroupMembersParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the post group members params
func (o *PostGroupMembersParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *PostGroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
