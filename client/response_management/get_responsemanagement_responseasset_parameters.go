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
)

// NewGetResponsemanagementResponseassetParams creates a new GetResponsemanagementResponseassetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetResponsemanagementResponseassetParams() *GetResponsemanagementResponseassetParams {
	return &GetResponsemanagementResponseassetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetResponsemanagementResponseassetParamsWithTimeout creates a new GetResponsemanagementResponseassetParams object
// with the ability to set a timeout on a request.
func NewGetResponsemanagementResponseassetParamsWithTimeout(timeout time.Duration) *GetResponsemanagementResponseassetParams {
	return &GetResponsemanagementResponseassetParams{
		timeout: timeout,
	}
}

// NewGetResponsemanagementResponseassetParamsWithContext creates a new GetResponsemanagementResponseassetParams object
// with the ability to set a context for a request.
func NewGetResponsemanagementResponseassetParamsWithContext(ctx context.Context) *GetResponsemanagementResponseassetParams {
	return &GetResponsemanagementResponseassetParams{
		Context: ctx,
	}
}

// NewGetResponsemanagementResponseassetParamsWithHTTPClient creates a new GetResponsemanagementResponseassetParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetResponsemanagementResponseassetParamsWithHTTPClient(client *http.Client) *GetResponsemanagementResponseassetParams {
	return &GetResponsemanagementResponseassetParams{
		HTTPClient: client,
	}
}

/*
GetResponsemanagementResponseassetParams contains all the parameters to send to the API endpoint

	for the get responsemanagement responseasset operation.

	Typically these are written to a http.Request.
*/
type GetResponsemanagementResponseassetParams struct {

	/* ResponseAssetID.

	   Asset Id
	*/
	ResponseAssetID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get responsemanagement responseasset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResponsemanagementResponseassetParams) WithDefaults() *GetResponsemanagementResponseassetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get responsemanagement responseasset params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetResponsemanagementResponseassetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) WithTimeout(timeout time.Duration) *GetResponsemanagementResponseassetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) WithContext(ctx context.Context) *GetResponsemanagementResponseassetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) WithHTTPClient(client *http.Client) *GetResponsemanagementResponseassetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResponseAssetID adds the responseAssetID to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) WithResponseAssetID(responseAssetID string) *GetResponsemanagementResponseassetParams {
	o.SetResponseAssetID(responseAssetID)
	return o
}

// SetResponseAssetID adds the responseAssetId to the get responsemanagement responseasset params
func (o *GetResponsemanagementResponseassetParams) SetResponseAssetID(responseAssetID string) {
	o.ResponseAssetID = responseAssetID
}

// WriteToRequest writes these params to a swagger request
func (o *GetResponsemanagementResponseassetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param responseAssetId
	if err := r.SetPathParam("responseAssetId", o.ResponseAssetID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
