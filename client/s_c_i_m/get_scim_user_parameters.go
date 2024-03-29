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

// NewGetScimUserParams creates a new GetScimUserParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetScimUserParams() *GetScimUserParams {
	return &GetScimUserParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetScimUserParamsWithTimeout creates a new GetScimUserParams object
// with the ability to set a timeout on a request.
func NewGetScimUserParamsWithTimeout(timeout time.Duration) *GetScimUserParams {
	return &GetScimUserParams{
		timeout: timeout,
	}
}

// NewGetScimUserParamsWithContext creates a new GetScimUserParams object
// with the ability to set a context for a request.
func NewGetScimUserParamsWithContext(ctx context.Context) *GetScimUserParams {
	return &GetScimUserParams{
		Context: ctx,
	}
}

// NewGetScimUserParamsWithHTTPClient creates a new GetScimUserParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetScimUserParamsWithHTTPClient(client *http.Client) *GetScimUserParams {
	return &GetScimUserParams{
		HTTPClient: client,
	}
}

/*
GetScimUserParams contains all the parameters to send to the API endpoint

	for the get scim user operation.

	Typically these are written to a http.Request.
*/
type GetScimUserParams struct {

	/* IfNoneMatch.

	   The ETag of a resource in double quotes. Returned as header and meta.version with initial call to GET /api/v2/scim/users/{userId}. Example: "42". If the ETag is different from the version on the server, returns the current configuration of the resource. If the ETag is current, returns 304 Not Modified.
	*/
	IfNoneMatch *string

	/* Attributes.

	   Indicates which attributes to include. Returns these attributes and the "id", "userName", "active", and "meta" attributes. Use "attributes" to avoid expensive secondary calls for the default attributes.
	*/
	Attributes []string

	/* ExcludedAttributes.

	   Indicates which attributes to exclude. Returns the default attributes minus "excludedAttributes". Always returns the "id", "userName", "active", and "meta" attributes. Use "excludedAttributes" to avoid expensive secondary calls for the default attributes.
	*/
	ExcludedAttributes []string

	/* UserID.

	   The ID of a user. Returned with GET /api/v2/scim/users.
	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get scim user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScimUserParams) WithDefaults() *GetScimUserParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get scim user params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetScimUserParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get scim user params
func (o *GetScimUserParams) WithTimeout(timeout time.Duration) *GetScimUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scim user params
func (o *GetScimUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scim user params
func (o *GetScimUserParams) WithContext(ctx context.Context) *GetScimUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scim user params
func (o *GetScimUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scim user params
func (o *GetScimUserParams) WithHTTPClient(client *http.Client) *GetScimUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scim user params
func (o *GetScimUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get scim user params
func (o *GetScimUserParams) WithIfNoneMatch(ifNoneMatch *string) *GetScimUserParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get scim user params
func (o *GetScimUserParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithAttributes adds the attributes to the get scim user params
func (o *GetScimUserParams) WithAttributes(attributes []string) *GetScimUserParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get scim user params
func (o *GetScimUserParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithExcludedAttributes adds the excludedAttributes to the get scim user params
func (o *GetScimUserParams) WithExcludedAttributes(excludedAttributes []string) *GetScimUserParams {
	o.SetExcludedAttributes(excludedAttributes)
	return o
}

// SetExcludedAttributes adds the excludedAttributes to the get scim user params
func (o *GetScimUserParams) SetExcludedAttributes(excludedAttributes []string) {
	o.ExcludedAttributes = excludedAttributes
}

// WithUserID adds the userID to the get scim user params
func (o *GetScimUserParams) WithUserID(userID string) *GetScimUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the get scim user params
func (o *GetScimUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *GetScimUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Attributes != nil {

		// binding items for attributes
		joinedAttributes := o.bindParamAttributes(reg)

		// query array param attributes
		if err := r.SetQueryParam("attributes", joinedAttributes...); err != nil {
			return err
		}
	}

	if o.ExcludedAttributes != nil {

		// binding items for excludedAttributes
		joinedExcludedAttributes := o.bindParamExcludedAttributes(reg)

		// query array param excludedAttributes
		if err := r.SetQueryParam("excludedAttributes", joinedExcludedAttributes...); err != nil {
			return err
		}
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

// bindParamGetScimUser binds the parameter attributes
func (o *GetScimUserParams) bindParamAttributes(formats strfmt.Registry) []string {
	attributesIR := o.Attributes

	var attributesIC []string
	for _, attributesIIR := range attributesIR { // explode []string

		attributesIIV := attributesIIR // string as string
		attributesIC = append(attributesIC, attributesIIV)
	}

	// items.CollectionFormat: "multi"
	attributesIS := swag.JoinByFormat(attributesIC, "multi")

	return attributesIS
}

// bindParamGetScimUser binds the parameter excludedAttributes
func (o *GetScimUserParams) bindParamExcludedAttributes(formats strfmt.Registry) []string {
	excludedAttributesIR := o.ExcludedAttributes

	var excludedAttributesIC []string
	for _, excludedAttributesIIR := range excludedAttributesIR { // explode []string

		excludedAttributesIIV := excludedAttributesIIR // string as string
		excludedAttributesIC = append(excludedAttributesIC, excludedAttributesIIV)
	}

	// items.CollectionFormat: "multi"
	excludedAttributesIS := swag.JoinByFormat(excludedAttributesIC, "multi")

	return excludedAttributesIS
}
