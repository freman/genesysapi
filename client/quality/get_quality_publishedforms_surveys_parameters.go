// Code generated by go-swagger; DO NOT EDIT.

package quality

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

// NewGetQualityPublishedformsSurveysParams creates a new GetQualityPublishedformsSurveysParams object
// with the default values initialized.
func NewGetQualityPublishedformsSurveysParams() *GetQualityPublishedformsSurveysParams {
	var (
		onlyLatestEnabledPerContextDefault = bool(false)
		pageNumberDefault                  = int32(1)
		pageSizeDefault                    = int32(25)
	)
	return &GetQualityPublishedformsSurveysParams{
		OnlyLatestEnabledPerContext: &onlyLatestEnabledPerContextDefault,
		PageNumber:                  &pageNumberDefault,
		PageSize:                    &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityPublishedformsSurveysParamsWithTimeout creates a new GetQualityPublishedformsSurveysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQualityPublishedformsSurveysParamsWithTimeout(timeout time.Duration) *GetQualityPublishedformsSurveysParams {
	var (
		onlyLatestEnabledPerContextDefault = bool(false)
		pageNumberDefault                  = int32(1)
		pageSizeDefault                    = int32(25)
	)
	return &GetQualityPublishedformsSurveysParams{
		OnlyLatestEnabledPerContext: &onlyLatestEnabledPerContextDefault,
		PageNumber:                  &pageNumberDefault,
		PageSize:                    &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetQualityPublishedformsSurveysParamsWithContext creates a new GetQualityPublishedformsSurveysParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQualityPublishedformsSurveysParamsWithContext(ctx context.Context) *GetQualityPublishedformsSurveysParams {
	var (
		onlyLatestEnabledPerContextDefault = bool(false)
		pageNumberDefault                  = int32(1)
		pageSizeDefault                    = int32(25)
	)
	return &GetQualityPublishedformsSurveysParams{
		OnlyLatestEnabledPerContext: &onlyLatestEnabledPerContextDefault,
		PageNumber:                  &pageNumberDefault,
		PageSize:                    &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetQualityPublishedformsSurveysParamsWithHTTPClient creates a new GetQualityPublishedformsSurveysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQualityPublishedformsSurveysParamsWithHTTPClient(client *http.Client) *GetQualityPublishedformsSurveysParams {
	var (
		onlyLatestEnabledPerContextDefault = bool(false)
		pageNumberDefault                  = int32(1)
		pageSizeDefault                    = int32(25)
	)
	return &GetQualityPublishedformsSurveysParams{
		OnlyLatestEnabledPerContext: &onlyLatestEnabledPerContextDefault,
		PageNumber:                  &pageNumberDefault,
		PageSize:                    &pageSizeDefault,
		HTTPClient:                  client,
	}
}

/*GetQualityPublishedformsSurveysParams contains all the parameters to send to the API endpoint
for the get quality publishedforms surveys operation typically these are written to a http.Request
*/
type GetQualityPublishedformsSurveysParams struct {

	/*Name
	  Name

	*/
	Name *string
	/*OnlyLatestEnabledPerContext
	  onlyLatestEnabledPerContext

	*/
	OnlyLatestEnabledPerContext *bool
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

// WithTimeout adds the timeout to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) WithTimeout(timeout time.Duration) *GetQualityPublishedformsSurveysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) WithContext(ctx context.Context) *GetQualityPublishedformsSurveysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) WithHTTPClient(client *http.Client) *GetQualityPublishedformsSurveysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) WithName(name *string) *GetQualityPublishedformsSurveysParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) SetName(name *string) {
	o.Name = name
}

// WithOnlyLatestEnabledPerContext adds the onlyLatestEnabledPerContext to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) WithOnlyLatestEnabledPerContext(onlyLatestEnabledPerContext *bool) *GetQualityPublishedformsSurveysParams {
	o.SetOnlyLatestEnabledPerContext(onlyLatestEnabledPerContext)
	return o
}

// SetOnlyLatestEnabledPerContext adds the onlyLatestEnabledPerContext to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) SetOnlyLatestEnabledPerContext(onlyLatestEnabledPerContext *bool) {
	o.OnlyLatestEnabledPerContext = onlyLatestEnabledPerContext
}

// WithPageNumber adds the pageNumber to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) WithPageNumber(pageNumber *int32) *GetQualityPublishedformsSurveysParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) WithPageSize(pageSize *int32) *GetQualityPublishedformsSurveysParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get quality publishedforms surveys params
func (o *GetQualityPublishedformsSurveysParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityPublishedformsSurveysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.OnlyLatestEnabledPerContext != nil {

		// query param onlyLatestEnabledPerContext
		var qrOnlyLatestEnabledPerContext bool
		if o.OnlyLatestEnabledPerContext != nil {
			qrOnlyLatestEnabledPerContext = *o.OnlyLatestEnabledPerContext
		}
		qOnlyLatestEnabledPerContext := swag.FormatBool(qrOnlyLatestEnabledPerContext)
		if qOnlyLatestEnabledPerContext != "" {
			if err := r.SetQueryParam("onlyLatestEnabledPerContext", qOnlyLatestEnabledPerContext); err != nil {
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