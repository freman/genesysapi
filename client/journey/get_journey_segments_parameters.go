// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewGetJourneySegmentsParams creates a new GetJourneySegmentsParams object
// with the default values initialized.
func NewGetJourneySegmentsParams() *GetJourneySegmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetJourneySegmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetJourneySegmentsParamsWithTimeout creates a new GetJourneySegmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetJourneySegmentsParamsWithTimeout(timeout time.Duration) *GetJourneySegmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetJourneySegmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetJourneySegmentsParamsWithContext creates a new GetJourneySegmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetJourneySegmentsParamsWithContext(ctx context.Context) *GetJourneySegmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetJourneySegmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetJourneySegmentsParamsWithHTTPClient creates a new GetJourneySegmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetJourneySegmentsParamsWithHTTPClient(client *http.Client) *GetJourneySegmentsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetJourneySegmentsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetJourneySegmentsParams contains all the parameters to send to the API endpoint
for the get journey segments operation typically these are written to a http.Request
*/
type GetJourneySegmentsParams struct {

	/*IsActive
	  Determines whether or not to show only active segments.

	*/
	IsActive *bool
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SortBy
	  Field(s) to sort by. The response can be sorted by any first level property on the Outcome response. Prefix with '-' for descending (e.g. sortBy=displayName,-createdDate).

	*/
	SortBy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get journey segments params
func (o *GetJourneySegmentsParams) WithTimeout(timeout time.Duration) *GetJourneySegmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get journey segments params
func (o *GetJourneySegmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get journey segments params
func (o *GetJourneySegmentsParams) WithContext(ctx context.Context) *GetJourneySegmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get journey segments params
func (o *GetJourneySegmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get journey segments params
func (o *GetJourneySegmentsParams) WithHTTPClient(client *http.Client) *GetJourneySegmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get journey segments params
func (o *GetJourneySegmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIsActive adds the isActive to the get journey segments params
func (o *GetJourneySegmentsParams) WithIsActive(isActive *bool) *GetJourneySegmentsParams {
	o.SetIsActive(isActive)
	return o
}

// SetIsActive adds the isActive to the get journey segments params
func (o *GetJourneySegmentsParams) SetIsActive(isActive *bool) {
	o.IsActive = isActive
}

// WithPageNumber adds the pageNumber to the get journey segments params
func (o *GetJourneySegmentsParams) WithPageNumber(pageNumber *int32) *GetJourneySegmentsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get journey segments params
func (o *GetJourneySegmentsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get journey segments params
func (o *GetJourneySegmentsParams) WithPageSize(pageSize *int32) *GetJourneySegmentsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get journey segments params
func (o *GetJourneySegmentsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get journey segments params
func (o *GetJourneySegmentsParams) WithSortBy(sortBy *string) *GetJourneySegmentsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get journey segments params
func (o *GetJourneySegmentsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WriteToRequest writes these params to a swagger request
func (o *GetJourneySegmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IsActive != nil {

		// query param isActive
		var qrIsActive bool
		if o.IsActive != nil {
			qrIsActive = *o.IsActive
		}
		qIsActive := swag.FormatBool(qrIsActive)
		if qIsActive != "" {
			if err := r.SetQueryParam("isActive", qIsActive); err != nil {
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

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}