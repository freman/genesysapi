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

// NewGetSpeechandtextanalyticsProgramsMappingsParams creates a new GetSpeechandtextanalyticsProgramsMappingsParams object
// with the default values initialized.
func NewGetSpeechandtextanalyticsProgramsMappingsParams() *GetSpeechandtextanalyticsProgramsMappingsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsMappingsParams{
		PageSize: &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsProgramsMappingsParamsWithTimeout creates a new GetSpeechandtextanalyticsProgramsMappingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpeechandtextanalyticsProgramsMappingsParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramsMappingsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsMappingsParams{
		PageSize: &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsProgramsMappingsParamsWithContext creates a new GetSpeechandtextanalyticsProgramsMappingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpeechandtextanalyticsProgramsMappingsParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramsMappingsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsMappingsParams{
		PageSize: &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsProgramsMappingsParamsWithHTTPClient creates a new GetSpeechandtextanalyticsProgramsMappingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpeechandtextanalyticsProgramsMappingsParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramsMappingsParams {
	var (
		pageSizeDefault = int32(20)
	)
	return &GetSpeechandtextanalyticsProgramsMappingsParams{
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetSpeechandtextanalyticsProgramsMappingsParams contains all the parameters to send to the API endpoint
for the get speechandtextanalytics programs mappings operation typically these are written to a http.Request
*/
type GetSpeechandtextanalyticsProgramsMappingsParams struct {

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

// WithTimeout adds the timeout to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramsMappingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramsMappingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramsMappingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNextPage adds the nextPage to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) WithNextPage(nextPage *string) *GetSpeechandtextanalyticsProgramsMappingsParams {
	o.SetNextPage(nextPage)
	return o
}

// SetNextPage adds the nextPage to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) SetNextPage(nextPage *string) {
	o.NextPage = nextPage
}

// WithPageSize adds the pageSize to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) WithPageSize(pageSize *int32) *GetSpeechandtextanalyticsProgramsMappingsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get speechandtextanalytics programs mappings params
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsProgramsMappingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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