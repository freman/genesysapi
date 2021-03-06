// Code generated by go-swagger; DO NOT EDIT.

package locations

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

// NewGetLocationsParams creates a new GetLocationsParams object
// with the default values initialized.
func NewGetLocationsParams() *GetLocationsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLocationsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLocationsParamsWithTimeout creates a new GetLocationsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLocationsParamsWithTimeout(timeout time.Duration) *GetLocationsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLocationsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetLocationsParamsWithContext creates a new GetLocationsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLocationsParamsWithContext(ctx context.Context) *GetLocationsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLocationsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetLocationsParamsWithHTTPClient creates a new GetLocationsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLocationsParamsWithHTTPClient(client *http.Client) *GetLocationsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLocationsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetLocationsParams contains all the parameters to send to the API endpoint
for the get locations operation typically these are written to a http.Request
*/
type GetLocationsParams struct {

	/*ID
	  id

	*/
	ID []string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortOrder
	  Sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get locations params
func (o *GetLocationsParams) WithTimeout(timeout time.Duration) *GetLocationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get locations params
func (o *GetLocationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get locations params
func (o *GetLocationsParams) WithContext(ctx context.Context) *GetLocationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get locations params
func (o *GetLocationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get locations params
func (o *GetLocationsParams) WithHTTPClient(client *http.Client) *GetLocationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get locations params
func (o *GetLocationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get locations params
func (o *GetLocationsParams) WithID(id []string) *GetLocationsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get locations params
func (o *GetLocationsParams) SetID(id []string) {
	o.ID = id
}

// WithPageNumber adds the pageNumber to the get locations params
func (o *GetLocationsParams) WithPageNumber(pageNumber *int32) *GetLocationsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get locations params
func (o *GetLocationsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get locations params
func (o *GetLocationsParams) WithPageSize(pageSize *int32) *GetLocationsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get locations params
func (o *GetLocationsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortOrder adds the sortOrder to the get locations params
func (o *GetLocationsParams) WithSortOrder(sortOrder *string) *GetLocationsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get locations params
func (o *GetLocationsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetLocationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
