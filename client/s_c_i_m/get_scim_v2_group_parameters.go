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

// NewGetScimV2GroupParams creates a new GetScimV2GroupParams object
// with the default values initialized.
func NewGetScimV2GroupParams() *GetScimV2GroupParams {
	var ()
	return &GetScimV2GroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScimV2GroupParamsWithTimeout creates a new GetScimV2GroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScimV2GroupParamsWithTimeout(timeout time.Duration) *GetScimV2GroupParams {
	var ()
	return &GetScimV2GroupParams{

		timeout: timeout,
	}
}

// NewGetScimV2GroupParamsWithContext creates a new GetScimV2GroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScimV2GroupParamsWithContext(ctx context.Context) *GetScimV2GroupParams {
	var ()
	return &GetScimV2GroupParams{

		Context: ctx,
	}
}

// NewGetScimV2GroupParamsWithHTTPClient creates a new GetScimV2GroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScimV2GroupParamsWithHTTPClient(client *http.Client) *GetScimV2GroupParams {
	var ()
	return &GetScimV2GroupParams{
		HTTPClient: client,
	}
}

/*GetScimV2GroupParams contains all the parameters to send to the API endpoint
for the get scim v2 group operation typically these are written to a http.Request
*/
type GetScimV2GroupParams struct {

	/*IfNoneMatch
	  TThe ETag of a resource in double quotes. Returned as header and meta.version with initial call to GET /api/v2/scim/v2/groups/{groupId}. Example: "42". If the ETag is different from the version on the server, returns the current configuration of the resource. If the ETag is current, returns 304 Not Modified.

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
	  The ID of a group. Returned with GET /api/v2/scim/v2/groups.

	*/
	GroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scim v2 group params
func (o *GetScimV2GroupParams) WithTimeout(timeout time.Duration) *GetScimV2GroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scim v2 group params
func (o *GetScimV2GroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scim v2 group params
func (o *GetScimV2GroupParams) WithContext(ctx context.Context) *GetScimV2GroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scim v2 group params
func (o *GetScimV2GroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scim v2 group params
func (o *GetScimV2GroupParams) WithHTTPClient(client *http.Client) *GetScimV2GroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scim v2 group params
func (o *GetScimV2GroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get scim v2 group params
func (o *GetScimV2GroupParams) WithIfNoneMatch(ifNoneMatch *string) *GetScimV2GroupParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get scim v2 group params
func (o *GetScimV2GroupParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithAttributes adds the attributes to the get scim v2 group params
func (o *GetScimV2GroupParams) WithAttributes(attributes []string) *GetScimV2GroupParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get scim v2 group params
func (o *GetScimV2GroupParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithExcludedAttributes adds the excludedAttributes to the get scim v2 group params
func (o *GetScimV2GroupParams) WithExcludedAttributes(excludedAttributes []string) *GetScimV2GroupParams {
	o.SetExcludedAttributes(excludedAttributes)
	return o
}

// SetExcludedAttributes adds the excludedAttributes to the get scim v2 group params
func (o *GetScimV2GroupParams) SetExcludedAttributes(excludedAttributes []string) {
	o.ExcludedAttributes = excludedAttributes
}

// WithGroupID adds the groupID to the get scim v2 group params
func (o *GetScimV2GroupParams) WithGroupID(groupID string) *GetScimV2GroupParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get scim v2 group params
func (o *GetScimV2GroupParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScimV2GroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
