// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingSkillgroupMembersParams creates a new GetRoutingSkillgroupMembersParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingSkillgroupMembersParams() *GetRoutingSkillgroupMembersParams {
	return &GetRoutingSkillgroupMembersParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingSkillgroupMembersParamsWithTimeout creates a new GetRoutingSkillgroupMembersParams object
// with the ability to set a timeout on a request.
func NewGetRoutingSkillgroupMembersParamsWithTimeout(timeout time.Duration) *GetRoutingSkillgroupMembersParams {
	return &GetRoutingSkillgroupMembersParams{
		timeout: timeout,
	}
}

// NewGetRoutingSkillgroupMembersParamsWithContext creates a new GetRoutingSkillgroupMembersParams object
// with the ability to set a context for a request.
func NewGetRoutingSkillgroupMembersParamsWithContext(ctx context.Context) *GetRoutingSkillgroupMembersParams {
	return &GetRoutingSkillgroupMembersParams{
		Context: ctx,
	}
}

// NewGetRoutingSkillgroupMembersParamsWithHTTPClient creates a new GetRoutingSkillgroupMembersParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingSkillgroupMembersParamsWithHTTPClient(client *http.Client) *GetRoutingSkillgroupMembersParams {
	return &GetRoutingSkillgroupMembersParams{
		HTTPClient: client,
	}
}

/*
GetRoutingSkillgroupMembersParams contains all the parameters to send to the API endpoint

	for the get routing skillgroup members operation.

	Typically these are written to a http.Request.
*/
type GetRoutingSkillgroupMembersParams struct {

	/* After.

	   The cursor that points to the next item
	*/
	After *string

	/* Before.

	   The cursor that points to the previous item
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

	/* SkillGroupID.

	   Skill Group ID
	*/
	SkillGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing skillgroup members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingSkillgroupMembersParams) WithDefaults() *GetRoutingSkillgroupMembersParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing skillgroup members params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingSkillgroupMembersParams) SetDefaults() {
	var (
		pageSizeDefault = int32(25)
	)

	val := GetRoutingSkillgroupMembersParams{
		PageSize: &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithTimeout(timeout time.Duration) *GetRoutingSkillgroupMembersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithContext(ctx context.Context) *GetRoutingSkillgroupMembersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithHTTPClient(client *http.Client) *GetRoutingSkillgroupMembersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithAfter(after *string) *GetRoutingSkillgroupMembersParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithBefore(before *string) *GetRoutingSkillgroupMembersParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetBefore(before *string) {
	o.Before = before
}

// WithExpand adds the expand to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithExpand(expand *string) *GetRoutingSkillgroupMembersParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithPageSize adds the pageSize to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithPageSize(pageSize *int32) *GetRoutingSkillgroupMembersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSkillGroupID adds the skillGroupID to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) WithSkillGroupID(skillGroupID string) *GetRoutingSkillgroupMembersParams {
	o.SetSkillGroupID(skillGroupID)
	return o
}

// SetSkillGroupID adds the skillGroupId to the get routing skillgroup members params
func (o *GetRoutingSkillgroupMembersParams) SetSkillGroupID(skillGroupID string) {
	o.SkillGroupID = skillGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingSkillgroupMembersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param skillGroupId
	if err := r.SetPathParam("skillGroupId", o.SkillGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}