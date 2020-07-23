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

// NewGetRoutingLanguagesParams creates a new GetRoutingLanguagesParams object
// with the default values initialized.
func NewGetRoutingLanguagesParams() *GetRoutingLanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetRoutingLanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingLanguagesParamsWithTimeout creates a new GetRoutingLanguagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingLanguagesParamsWithTimeout(timeout time.Duration) *GetRoutingLanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetRoutingLanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetRoutingLanguagesParamsWithContext creates a new GetRoutingLanguagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingLanguagesParamsWithContext(ctx context.Context) *GetRoutingLanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetRoutingLanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetRoutingLanguagesParamsWithHTTPClient creates a new GetRoutingLanguagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingLanguagesParamsWithHTTPClient(client *http.Client) *GetRoutingLanguagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("ASC")
	)
	return &GetRoutingLanguagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetRoutingLanguagesParams contains all the parameters to send to the API endpoint
for the get routing languages operation typically these are written to a http.Request
*/
type GetRoutingLanguagesParams struct {

	/*ID
	  id

	*/
	ID []string
	/*Name
	  Name

	*/
	Name *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortOrder
	  Ascending or descending sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing languages params
func (o *GetRoutingLanguagesParams) WithTimeout(timeout time.Duration) *GetRoutingLanguagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing languages params
func (o *GetRoutingLanguagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing languages params
func (o *GetRoutingLanguagesParams) WithContext(ctx context.Context) *GetRoutingLanguagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing languages params
func (o *GetRoutingLanguagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing languages params
func (o *GetRoutingLanguagesParams) WithHTTPClient(client *http.Client) *GetRoutingLanguagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing languages params
func (o *GetRoutingLanguagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get routing languages params
func (o *GetRoutingLanguagesParams) WithID(id []string) *GetRoutingLanguagesParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get routing languages params
func (o *GetRoutingLanguagesParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get routing languages params
func (o *GetRoutingLanguagesParams) WithName(name *string) *GetRoutingLanguagesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get routing languages params
func (o *GetRoutingLanguagesParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get routing languages params
func (o *GetRoutingLanguagesParams) WithPageNumber(pageNumber *int32) *GetRoutingLanguagesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing languages params
func (o *GetRoutingLanguagesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing languages params
func (o *GetRoutingLanguagesParams) WithPageSize(pageSize *int32) *GetRoutingLanguagesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing languages params
func (o *GetRoutingLanguagesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortOrder adds the sortOrder to the get routing languages params
func (o *GetRoutingLanguagesParams) WithSortOrder(sortOrder *string) *GetRoutingLanguagesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get routing languages params
func (o *GetRoutingLanguagesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingLanguagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
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

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string
		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {
			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
