// Code generated by go-swagger; DO NOT EDIT.

package languages

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

// NewDeleteLanguageParams creates a new DeleteLanguageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteLanguageParams() *DeleteLanguageParams {
	return &DeleteLanguageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLanguageParamsWithTimeout creates a new DeleteLanguageParams object
// with the ability to set a timeout on a request.
func NewDeleteLanguageParamsWithTimeout(timeout time.Duration) *DeleteLanguageParams {
	return &DeleteLanguageParams{
		timeout: timeout,
	}
}

// NewDeleteLanguageParamsWithContext creates a new DeleteLanguageParams object
// with the ability to set a context for a request.
func NewDeleteLanguageParamsWithContext(ctx context.Context) *DeleteLanguageParams {
	return &DeleteLanguageParams{
		Context: ctx,
	}
}

// NewDeleteLanguageParamsWithHTTPClient creates a new DeleteLanguageParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteLanguageParamsWithHTTPClient(client *http.Client) *DeleteLanguageParams {
	return &DeleteLanguageParams{
		HTTPClient: client,
	}
}

/*
DeleteLanguageParams contains all the parameters to send to the API endpoint

	for the delete language operation.

	Typically these are written to a http.Request.
*/
type DeleteLanguageParams struct {

	/* LanguageID.

	   Language ID
	*/
	LanguageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete language params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLanguageParams) WithDefaults() *DeleteLanguageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete language params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteLanguageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete language params
func (o *DeleteLanguageParams) WithTimeout(timeout time.Duration) *DeleteLanguageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete language params
func (o *DeleteLanguageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete language params
func (o *DeleteLanguageParams) WithContext(ctx context.Context) *DeleteLanguageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete language params
func (o *DeleteLanguageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete language params
func (o *DeleteLanguageParams) WithHTTPClient(client *http.Client) *DeleteLanguageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete language params
func (o *DeleteLanguageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguageID adds the languageID to the delete language params
func (o *DeleteLanguageParams) WithLanguageID(languageID string) *DeleteLanguageParams {
	o.SetLanguageID(languageID)
	return o
}

// SetLanguageID adds the languageId to the delete language params
func (o *DeleteLanguageParams) SetLanguageID(languageID string) {
	o.LanguageID = languageID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLanguageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param languageId
	if err := r.SetPathParam("languageId", o.LanguageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
