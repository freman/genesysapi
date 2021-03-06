// Code generated by go-swagger; DO NOT EDIT.

package response_management

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

// NewGetResponsemanagementLibrariesParams creates a new GetResponsemanagementLibrariesParams object
// with the default values initialized.
func NewGetResponsemanagementLibrariesParams() *GetResponsemanagementLibrariesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetResponsemanagementLibrariesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetResponsemanagementLibrariesParamsWithTimeout creates a new GetResponsemanagementLibrariesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetResponsemanagementLibrariesParamsWithTimeout(timeout time.Duration) *GetResponsemanagementLibrariesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetResponsemanagementLibrariesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetResponsemanagementLibrariesParamsWithContext creates a new GetResponsemanagementLibrariesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetResponsemanagementLibrariesParamsWithContext(ctx context.Context) *GetResponsemanagementLibrariesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetResponsemanagementLibrariesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetResponsemanagementLibrariesParamsWithHTTPClient creates a new GetResponsemanagementLibrariesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetResponsemanagementLibrariesParamsWithHTTPClient(client *http.Client) *GetResponsemanagementLibrariesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetResponsemanagementLibrariesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetResponsemanagementLibrariesParams contains all the parameters to send to the API endpoint
for the get responsemanagement libraries operation typically these are written to a http.Request
*/
type GetResponsemanagementLibrariesParams struct {

	/*MessagingTemplateFilter
	  Returns a list of libraries that contain responses with at least one messaging template defined for a specific message channel

	*/
	MessagingTemplateFilter *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) WithTimeout(timeout time.Duration) *GetResponsemanagementLibrariesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) WithContext(ctx context.Context) *GetResponsemanagementLibrariesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) WithHTTPClient(client *http.Client) *GetResponsemanagementLibrariesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessagingTemplateFilter adds the messagingTemplateFilter to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) WithMessagingTemplateFilter(messagingTemplateFilter *string) *GetResponsemanagementLibrariesParams {
	o.SetMessagingTemplateFilter(messagingTemplateFilter)
	return o
}

// SetMessagingTemplateFilter adds the messagingTemplateFilter to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) SetMessagingTemplateFilter(messagingTemplateFilter *string) {
	o.MessagingTemplateFilter = messagingTemplateFilter
}

// WithPageNumber adds the pageNumber to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) WithPageNumber(pageNumber *int32) *GetResponsemanagementLibrariesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) WithPageSize(pageSize *int32) *GetResponsemanagementLibrariesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get responsemanagement libraries params
func (o *GetResponsemanagementLibrariesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetResponsemanagementLibrariesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MessagingTemplateFilter != nil {

		// query param messagingTemplateFilter
		var qrMessagingTemplateFilter string
		if o.MessagingTemplateFilter != nil {
			qrMessagingTemplateFilter = *o.MessagingTemplateFilter
		}
		qMessagingTemplateFilter := qrMessagingTemplateFilter
		if qMessagingTemplateFilter != "" {
			if err := r.SetQueryParam("messagingTemplateFilter", qMessagingTemplateFilter); err != nil {
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
