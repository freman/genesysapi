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
)

// NewDeleteTeamMembersParams creates a new DeleteTeamMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTeamMembersParams() *DeleteTeamMembersParams {
	return &DeleteTeamMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTeamMembersParamsWithTimeout creates a new DeleteTeamMembersParams object
// with the ability to set a timeout on a request.
func NewDeleteTeamMembersParamsWithTimeout(timeout time.Duration) *DeleteTeamMembersParams {
	return &DeleteTeamMembersParams{
		timeout: timeout,
	}
}

// NewDeleteTeamMembersParamsWithContext creates a new DeleteTeamMembersParams object
// with the ability to set a context for a request.
func NewDeleteTeamMembersParamsWithContext(ctx context.Context) *DeleteTeamMembersParams {
	return &DeleteTeamMembersParams{
		Context: ctx,
	}
}

// NewDeleteTeamMembersParamsWithHTTPClient creates a new DeleteTeamMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTeamMembersParamsWithHTTPClient(client *http.Client) *DeleteTeamMembersParams {
	return &DeleteTeamMembersParams{
		HTTPClient: client,
	}
}

/*
DeleteTeamMembersParams contains all the parameters to send to the API endpoint

	for the delete team members operation.

	Typically these are written to a http.Request.
*/
type DeleteTeamMembersParams struct {

	/* ID.

	   Comma separated list of member ids to remove
	*/
	ID string

	/* TeamID.

	   Team ID
	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamMembersParams) WithDefaults() *DeleteTeamMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTeamMembersParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete team members params
func (o *DeleteTeamMembersParams) WithTimeout(timeout time.Duration) *DeleteTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete team members params
func (o *DeleteTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete team members params
func (o *DeleteTeamMembersParams) WithContext(ctx context.Context) *DeleteTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete team members params
func (o *DeleteTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete team members params
func (o *DeleteTeamMembersParams) WithHTTPClient(client *http.Client) *DeleteTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete team members params
func (o *DeleteTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete team members params
func (o *DeleteTeamMembersParams) WithID(id string) *DeleteTeamMembersParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete team members params
func (o *DeleteTeamMembersParams) SetID(id string) {
	o.ID = id
}

// WithTeamID adds the teamID to the delete team members params
func (o *DeleteTeamMembersParams) WithTeamID(teamID string) *DeleteTeamMembersParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the delete team members params
func (o *DeleteTeamMembersParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
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