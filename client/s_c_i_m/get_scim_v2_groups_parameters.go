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

// NewGetScimV2GroupsParams creates a new GetScimV2GroupsParams object
// with the default values initialized.
func NewGetScimV2GroupsParams() *GetScimV2GroupsParams {
	var (
		countDefault      = int32(25)
		startIndexDefault = int32(1)
	)
	return &GetScimV2GroupsParams{
		Count:      &countDefault,
		StartIndex: &startIndexDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetScimV2GroupsParamsWithTimeout creates a new GetScimV2GroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetScimV2GroupsParamsWithTimeout(timeout time.Duration) *GetScimV2GroupsParams {
	var (
		countDefault      = int32(25)
		startIndexDefault = int32(1)
	)
	return &GetScimV2GroupsParams{
		Count:      &countDefault,
		StartIndex: &startIndexDefault,

		timeout: timeout,
	}
}

// NewGetScimV2GroupsParamsWithContext creates a new GetScimV2GroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetScimV2GroupsParamsWithContext(ctx context.Context) *GetScimV2GroupsParams {
	var (
		countDefault      = int32(25)
		startIndexDefault = int32(1)
	)
	return &GetScimV2GroupsParams{
		Count:      &countDefault,
		StartIndex: &startIndexDefault,

		Context: ctx,
	}
}

// NewGetScimV2GroupsParamsWithHTTPClient creates a new GetScimV2GroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetScimV2GroupsParamsWithHTTPClient(client *http.Client) *GetScimV2GroupsParams {
	var (
		countDefault      = int32(25)
		startIndexDefault = int32(1)
	)
	return &GetScimV2GroupsParams{
		Count:      &countDefault,
		StartIndex: &startIndexDefault,
		HTTPClient: client,
	}
}

/*GetScimV2GroupsParams contains all the parameters to send to the API endpoint
for the get scim v2 groups operation typically these are written to a http.Request
*/
type GetScimV2GroupsParams struct {

	/*Attributes
	  Indicates which attributes to include. Returns these attributes and the 'id', 'active', and 'meta attributes . Use "attributes" to avoid expensive secondary calls for the default attributes.

	*/
	Attributes []string
	/*Count
	  The requested number of items per page. A value of 0 returns "totalResults". Note that a page size over 25 will likely cause a 429 error by exceeding internal resource limits. Page sizes over 25 will require using excludedAttributes and includeAttributes query parameters to exclude secondary lookup values -- (i.e. externalId, roles, urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:routingLanguages, urn:ietf:params:scim:schemas:extension:genesys:purecloud:2.0:User:routingSkills)

	*/
	Count *int32
	/*ExcludedAttributes
	  Indicates which attributes to exclude. Returns the default attributes minus "excludedAttributes". Use "excludedAttributes" to avoid expensive secondary calls for the default attributes. The'id', 'active', and 'meta' attributes will always be present in the output.

	*/
	ExcludedAttributes []string
	/*Filter
	  Filters results. If nothing is specified, returns all groups. Examples of valid values: "id eq 5f4bc742-a019-4e38-8e2a-d39d5bc0b0f3", "displayname eq Sales".

	*/
	Filter string
	/*StartIndex
	  The 1-based index of the first query result.

	*/
	StartIndex *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithTimeout(timeout time.Duration) *GetScimV2GroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithContext(ctx context.Context) *GetScimV2GroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithHTTPClient(client *http.Client) *GetScimV2GroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithAttributes(attributes []string) *GetScimV2GroupsParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithCount adds the count to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithCount(count *int32) *GetScimV2GroupsParams {
	o.SetCount(count)
	return o
}

// SetCount adds the count to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetCount(count *int32) {
	o.Count = count
}

// WithExcludedAttributes adds the excludedAttributes to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithExcludedAttributes(excludedAttributes []string) *GetScimV2GroupsParams {
	o.SetExcludedAttributes(excludedAttributes)
	return o
}

// SetExcludedAttributes adds the excludedAttributes to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetExcludedAttributes(excludedAttributes []string) {
	o.ExcludedAttributes = excludedAttributes
}

// WithFilter adds the filter to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithFilter(filter string) *GetScimV2GroupsParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetFilter(filter string) {
	o.Filter = filter
}

// WithStartIndex adds the startIndex to the get scim v2 groups params
func (o *GetScimV2GroupsParams) WithStartIndex(startIndex *int32) *GetScimV2GroupsParams {
	o.SetStartIndex(startIndex)
	return o
}

// SetStartIndex adds the startIndex to the get scim v2 groups params
func (o *GetScimV2GroupsParams) SetStartIndex(startIndex *int32) {
	o.StartIndex = startIndex
}

// WriteToRequest writes these params to a swagger request
func (o *GetScimV2GroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAttributes := o.Attributes

	joinedAttributes := swag.JoinByFormat(valuesAttributes, "multi")
	// query array param attributes
	if err := r.SetQueryParam("attributes", joinedAttributes...); err != nil {
		return err
	}

	if o.Count != nil {

		// query param count
		var qrCount int32
		if o.Count != nil {
			qrCount = *o.Count
		}
		qCount := swag.FormatInt32(qrCount)
		if qCount != "" {
			if err := r.SetQueryParam("count", qCount); err != nil {
				return err
			}
		}

	}

	valuesExcludedAttributes := o.ExcludedAttributes

	joinedExcludedAttributes := swag.JoinByFormat(valuesExcludedAttributes, "multi")
	// query array param excludedAttributes
	if err := r.SetQueryParam("excludedAttributes", joinedExcludedAttributes...); err != nil {
		return err
	}

	// query param filter
	qrFilter := o.Filter
	qFilter := qrFilter
	if qFilter != "" {
		if err := r.SetQueryParam("filter", qFilter); err != nil {
			return err
		}
	}

	if o.StartIndex != nil {

		// query param startIndex
		var qrStartIndex int32
		if o.StartIndex != nil {
			qrStartIndex = *o.StartIndex
		}
		qStartIndex := swag.FormatInt32(qrStartIndex)
		if qStartIndex != "" {
			if err := r.SetQueryParam("startIndex", qStartIndex); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}