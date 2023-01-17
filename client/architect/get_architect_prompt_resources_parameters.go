// Code generated by go-swagger; DO NOT EDIT.

package architect

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

// NewGetArchitectPromptResourcesParams creates a new GetArchitectPromptResourcesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetArchitectPromptResourcesParams() *GetArchitectPromptResourcesParams {
	return &GetArchitectPromptResourcesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetArchitectPromptResourcesParamsWithTimeout creates a new GetArchitectPromptResourcesParams object
// with the ability to set a timeout on a request.
func NewGetArchitectPromptResourcesParamsWithTimeout(timeout time.Duration) *GetArchitectPromptResourcesParams {
	return &GetArchitectPromptResourcesParams{
		timeout: timeout,
	}
}

// NewGetArchitectPromptResourcesParamsWithContext creates a new GetArchitectPromptResourcesParams object
// with the ability to set a context for a request.
func NewGetArchitectPromptResourcesParamsWithContext(ctx context.Context) *GetArchitectPromptResourcesParams {
	return &GetArchitectPromptResourcesParams{
		Context: ctx,
	}
}

// NewGetArchitectPromptResourcesParamsWithHTTPClient creates a new GetArchitectPromptResourcesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetArchitectPromptResourcesParamsWithHTTPClient(client *http.Client) *GetArchitectPromptResourcesParams {
	return &GetArchitectPromptResourcesParams{
		HTTPClient: client,
	}
}

/*
GetArchitectPromptResourcesParams contains all the parameters to send to the API endpoint

	for the get architect prompt resources operation.

	Typically these are written to a http.Request.
*/
type GetArchitectPromptResourcesParams struct {

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

	/* PromptID.

	   Prompt ID
	*/
	PromptID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get architect prompt resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectPromptResourcesParams) WithDefaults() *GetArchitectPromptResourcesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get architect prompt resources params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetArchitectPromptResourcesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetArchitectPromptResourcesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) WithTimeout(timeout time.Duration) *GetArchitectPromptResourcesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) WithContext(ctx context.Context) *GetArchitectPromptResourcesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) WithHTTPClient(client *http.Client) *GetArchitectPromptResourcesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) WithPageNumber(pageNumber *int32) *GetArchitectPromptResourcesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) WithPageSize(pageSize *int32) *GetArchitectPromptResourcesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPromptID adds the promptID to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) WithPromptID(promptID string) *GetArchitectPromptResourcesParams {
	o.SetPromptID(promptID)
	return o
}

// SetPromptID adds the promptId to the get architect prompt resources params
func (o *GetArchitectPromptResourcesParams) SetPromptID(promptID string) {
	o.PromptID = promptID
}

// WriteToRequest writes these params to a swagger request
func (o *GetArchitectPromptResourcesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param promptId
	if err := r.SetPathParam("promptId", o.PromptID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
