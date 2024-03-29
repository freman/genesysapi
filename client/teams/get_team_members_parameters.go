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
	"github.com/go-openapi/swag"
)

// NewGetTeamMembersParams creates a new GetTeamMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTeamMembersParams() *GetTeamMembersParams {
	return &GetTeamMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTeamMembersParamsWithTimeout creates a new GetTeamMembersParams object
// with the ability to set a timeout on a request.
func NewGetTeamMembersParamsWithTimeout(timeout time.Duration) *GetTeamMembersParams {
	return &GetTeamMembersParams{
		timeout: timeout,
	}
}

// NewGetTeamMembersParamsWithContext creates a new GetTeamMembersParams object
// with the ability to set a context for a request.
func NewGetTeamMembersParamsWithContext(ctx context.Context) *GetTeamMembersParams {
	return &GetTeamMembersParams{
		Context: ctx,
	}
}

// NewGetTeamMembersParamsWithHTTPClient creates a new GetTeamMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTeamMembersParamsWithHTTPClient(client *http.Client) *GetTeamMembersParams {
	return &GetTeamMembersParams{
		HTTPClient: client,
	}
}

/*
GetTeamMembersParams contains all the parameters to send to the API endpoint

	for the get team members operation.

	Typically these are written to a http.Request.
*/
type GetTeamMembersParams struct {

	/* After.

	   The cursor that points to the next item in the complete list of teams
	*/
	After *string

	/* Before.

	   The cursor that points to the previous item in the complete list of teams
	*/
	Before *string

	/* Expand.

	   Expand the name on each user
	*/
	Expand *string

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* TeamID.

	   Team ID
	*/
	TeamID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamMembersParams) WithDefaults() *GetTeamMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get team members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTeamMembersParams) SetDefaults() {
	var (
		pageSizeDefault = int32(25)
	)

	val := GetTeamMembersParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get team members params
func (o *GetTeamMembersParams) WithTimeout(timeout time.Duration) *GetTeamMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get team members params
func (o *GetTeamMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get team members params
func (o *GetTeamMembersParams) WithContext(ctx context.Context) *GetTeamMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get team members params
func (o *GetTeamMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get team members params
func (o *GetTeamMembersParams) WithHTTPClient(client *http.Client) *GetTeamMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get team members params
func (o *GetTeamMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get team members params
func (o *GetTeamMembersParams) WithAfter(after *string) *GetTeamMembersParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get team members params
func (o *GetTeamMembersParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get team members params
func (o *GetTeamMembersParams) WithBefore(before *string) *GetTeamMembersParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get team members params
func (o *GetTeamMembersParams) SetBefore(before *string) {
	o.Before = before
}

// WithExpand adds the expand to the get team members params
func (o *GetTeamMembersParams) WithExpand(expand *string) *GetTeamMembersParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get team members params
func (o *GetTeamMembersParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithPageSize adds the pageSize to the get team members params
func (o *GetTeamMembersParams) WithPageSize(pageSize *int32) *GetTeamMembersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get team members params
func (o *GetTeamMembersParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithTeamID adds the teamID to the get team members params
func (o *GetTeamMembersParams) WithTeamID(teamID string) *GetTeamMembersParams {
	o.SetTeamID(teamID)
	return o
}

// SetTeamID adds the teamId to the get team members params
func (o *GetTeamMembersParams) SetTeamID(teamID string) {
	o.TeamID = teamID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTeamMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	if o.Before != nil {

		// query param before
		var qrBefore string

		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {

			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}
	}

	if o.Expand != nil {

		// query param expand
		var qrExpand string

		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {

			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
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
