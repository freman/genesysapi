// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

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

// NewGetSpeechandtextanalyticsProgramsUnpublishedParams creates a new GetSpeechandtextanalyticsProgramsUnpublishedParams object
// with the default values initialized.
func NewGetSpeechandtextanalyticsProgramsUnpublishedParams() *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsUnpublishedParams{
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsProgramsUnpublishedParamsWithTimeout creates a new GetSpeechandtextanalyticsProgramsUnpublishedParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpeechandtextanalyticsProgramsUnpublishedParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsUnpublishedParams{
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsProgramsUnpublishedParamsWithContext creates a new GetSpeechandtextanalyticsProgramsUnpublishedParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpeechandtextanalyticsProgramsUnpublishedParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsUnpublishedParams{
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsProgramsUnpublishedParamsWithHTTPClient creates a new GetSpeechandtextanalyticsProgramsUnpublishedParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpeechandtextanalyticsProgramsUnpublishedParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsUnpublishedParams{
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetSpeechandtextanalyticsProgramsUnpublishedParams contains all the parameters to send to the API endpoint
for the get speechandtextanalytics programs unpublished operation typically these are written to a http.Request
*/
type GetSpeechandtextanalyticsProgramsUnpublishedParams struct {

	/*NextPage
	  The key for listing the next page

	*/
	NextPage *string
	/*PageSize
	  The page size for the listing

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNextPage adds the nextPage to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) WithNextPage(nextPage *string) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageSize adds the pageSize to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) WithPageSize(pageSize *int32) *GetSpeechandtextanalyticsProgramsUnpublishedParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get speechandtextanalytics programs unpublished params
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsProgramsUnpublishedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.NextPage != nil {

		// query param nextPage
		var qrNextPage string
		if o.NextPage != nil {
			qrNextPage = *o.NextPage
		}
		qNextPage := qrNextPage
		if qNextPage != "" {
			if err := r.SetQueryParam("nextPage", qNextPage); err != nil {
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
