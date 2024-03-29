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

// NewGetRoutingSkillsParams creates a new GetRoutingSkillsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRoutingSkillsParams() *GetRoutingSkillsParams {
	return &GetRoutingSkillsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingSkillsParamsWithTimeout creates a new GetRoutingSkillsParams object
// with the ability to set a timeout on a request.
func NewGetRoutingSkillsParamsWithTimeout(timeout time.Duration) *GetRoutingSkillsParams {
	return &GetRoutingSkillsParams{
		timeout: timeout,
	}
}

// NewGetRoutingSkillsParamsWithContext creates a new GetRoutingSkillsParams object
// with the ability to set a context for a request.
func NewGetRoutingSkillsParamsWithContext(ctx context.Context) *GetRoutingSkillsParams {
	return &GetRoutingSkillsParams{
		Context: ctx,
	}
}

// NewGetRoutingSkillsParamsWithHTTPClient creates a new GetRoutingSkillsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRoutingSkillsParamsWithHTTPClient(client *http.Client) *GetRoutingSkillsParams {
	return &GetRoutingSkillsParams{
		HTTPClient: client,
	}
}

/*
GetRoutingSkillsParams contains all the parameters to send to the API endpoint

	for the get routing skills operation.

	Typically these are written to a http.Request.
*/
type GetRoutingSkillsParams struct {

	/* ID.

	   id
	*/
	ID []string

	/* Name.

	   Filter for results that start with this value
	*/
	Name *string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get routing skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingSkillsParams) WithDefaults() *GetRoutingSkillsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get routing skills params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRoutingSkillsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetRoutingSkillsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get routing skills params
func (o *GetRoutingSkillsParams) WithTimeout(timeout time.Duration) *GetRoutingSkillsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing skills params
func (o *GetRoutingSkillsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing skills params
func (o *GetRoutingSkillsParams) WithContext(ctx context.Context) *GetRoutingSkillsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing skills params
func (o *GetRoutingSkillsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing skills params
func (o *GetRoutingSkillsParams) WithHTTPClient(client *http.Client) *GetRoutingSkillsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing skills params
func (o *GetRoutingSkillsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get routing skills params
func (o *GetRoutingSkillsParams) WithID(id []string) *GetRoutingSkillsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get routing skills params
func (o *GetRoutingSkillsParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get routing skills params
func (o *GetRoutingSkillsParams) WithName(name *string) *GetRoutingSkillsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get routing skills params
func (o *GetRoutingSkillsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get routing skills params
func (o *GetRoutingSkillsParams) WithPageNumber(pageNumber *int32) *GetRoutingSkillsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing skills params
func (o *GetRoutingSkillsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing skills params
func (o *GetRoutingSkillsParams) WithPageSize(pageSize *int32) *GetRoutingSkillsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing skills params
func (o *GetRoutingSkillsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingSkillsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetRoutingSkills binds the parameter id
func (o *GetRoutingSkillsParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
