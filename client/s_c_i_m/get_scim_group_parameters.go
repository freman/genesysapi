// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

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

// NewGetScimGroupParams creates a new GetScimGroupParams object
// with the default values initialized.
func NewGetScimGroupParams() *GetScimGroupParams {
	var ()
	return &GetScimGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScimGroupParamsWithTimeout creates a new GetScimGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScimGroupParamsWithTimeout(timeout time.Duration) *GetScimGroupParams {
	var ()
	return &GetScimGroupParams{

		timeout: timeout,
	}
}

// NewGetScimGroupParamsWithContext creates a new GetScimGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScimGroupParamsWithContext(ctx context.Context) *GetScimGroupParams {
	var ()
	return &GetScimGroupParams{

		Context: ctx,
	}
}

// NewGetScimGroupParamsWithHTTPClient creates a new GetScimGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScimGroupParamsWithHTTPClient(client *http.Client) *GetScimGroupParams {
	var ()
	return &GetScimGroupParams{
		HTTPClient: client,
	}
}

/*GetScimGroupParams contains all the parameters to send to the API endpoint
for the get scim group operation typically these are written to a http.Request
*/
type GetScimGroupParams struct {

	/*IfNoneMatch
	  The ETag of a resource in double quotes. Returned as header and meta.version with initial call to GET /api/v2/scim/groups/{groupId}. Example: "42". If the ETag is different from the version on the server, returns the current configuration of the resource. If the ETag is current, returns 304 Not Modified.

	*/
	IfNoneMatch *string
	/*Attributes
	  Indicates which attributes to include. Returns these attributes and the 'id', 'active', and 'meta attributes . Use "attributes" to avoid expensive secondary calls for the default attributes.

	*/
	Attributes []string
	/*ExcludedAttributes
	  Indicates which attributes to exclude. Returns the default attributes minus "excludedAttributes". Use "excludedAttributes" to avoid expensive secondary calls for the default attributes. The'id', 'active', and 'meta' attributes will always be present in the output.

	*/
	ExcludedAttributes []string
	/*GroupID
	  The ID of a group. Returned with GET /api/v2/scim/groups.

	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scim group params
func (o *GetScimGroupParams) WithTimeout(timeout time.Duration) *GetScimGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scim group params
func (o *GetScimGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scim group params
func (o *GetScimGroupParams) WithContext(ctx context.Context) *GetScimGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scim group params
func (o *GetScimGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scim group params
func (o *GetScimGroupParams) WithHTTPClient(client *http.Client) *GetScimGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scim group params
func (o *GetScimGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get scim group params
func (o *GetScimGroupParams) WithIfNoneMatch(ifNoneMatch *string) *GetScimGroupParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get scim group params
func (o *GetScimGroupParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithAttributes adds the attributes to the get scim group params
func (o *GetScimGroupParams) WithAttributes(attributes []string) *GetScimGroupParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get scim group params
func (o *GetScimGroupParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithExcludedAttributes adds the excludedAttributes to the get scim group params
func (o *GetScimGroupParams) WithExcludedAttributes(excludedAttributes []string) *GetScimGroupParams {
	o.SetExcludedAttributes(excludedAttributes)
	return o
}

// SetExcludedAttributes adds the excludedAttributes to the get scim group params
func (o *GetScimGroupParams) SetExcludedAttributes(excludedAttributes []string) {
	o.ExcludedAttributes = excludedAttributes
}

// WithGroupID adds the groupID to the get scim group params
func (o *GetScimGroupParams) WithGroupID(groupID string) *GetScimGroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get scim group params
func (o *GetScimGroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScimGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}

	}

	valuesAttributes := o.Attributes

	joinedAttributes := swag.JoinByFormat(valuesAttributes, "multi")
	// query array param attributes
	if err := r.SetQueryParam("attributes", joinedAttributes...); err != nil {
		return err
	}

	valuesExcludedAttributes := o.ExcludedAttributes

	joinedExcludedAttributes := swag.JoinByFormat(valuesExcludedAttributes, "multi")
	// query array param excludedAttributes
	if err := r.SetQueryParam("excludedAttributes", joinedExcludedAttributes...); err != nil {
		return err
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