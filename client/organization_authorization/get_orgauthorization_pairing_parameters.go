// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

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

// NewGetOrgauthorizationPairingParams creates a new GetOrgauthorizationPairingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrgauthorizationPairingParams() *GetOrgauthorizationPairingParams {
	return &GetOrgauthorizationPairingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrgauthorizationPairingParamsWithTimeout creates a new GetOrgauthorizationPairingParams object
// with the ability to set a timeout on a request.
func NewGetOrgauthorizationPairingParamsWithTimeout(timeout time.Duration) *GetOrgauthorizationPairingParams {
	return &GetOrgauthorizationPairingParams{
		timeout: timeout,
	}
}

// NewGetOrgauthorizationPairingParamsWithContext creates a new GetOrgauthorizationPairingParams object
// with the ability to set a context for a request.
func NewGetOrgauthorizationPairingParamsWithContext(ctx context.Context) *GetOrgauthorizationPairingParams {
	return &GetOrgauthorizationPairingParams{
		Context: ctx,
	}
}

// NewGetOrgauthorizationPairingParamsWithHTTPClient creates a new GetOrgauthorizationPairingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrgauthorizationPairingParamsWithHTTPClient(client *http.Client) *GetOrgauthorizationPairingParams {
	return &GetOrgauthorizationPairingParams{
		HTTPClient: client,
	}
}

/*
GetOrgauthorizationPairingParams contains all the parameters to send to the API endpoint

	for the get orgauthorization pairing operation.

	Typically these are written to a http.Request.
*/
type GetOrgauthorizationPairingParams struct {

	/* PairingID.

	   Pairing Id
	*/
	PairingID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get orgauthorization pairing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationPairingParams) WithDefaults() *GetOrgauthorizationPairingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get orgauthorization pairing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrgauthorizationPairingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) WithTimeout(timeout time.Duration) *GetOrgauthorizationPairingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) WithContext(ctx context.Context) *GetOrgauthorizationPairingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) WithHTTPClient(client *http.Client) *GetOrgauthorizationPairingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPairingID adds the pairingID to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) WithPairingID(pairingID string) *GetOrgauthorizationPairingParams {
	o.SetPairingID(pairingID)
	return o
}

// SetPairingID adds the pairingId to the get orgauthorization pairing params
func (o *GetOrgauthorizationPairingParams) SetPairingID(pairingID string) {
	o.PairingID = pairingID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrgauthorizationPairingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param pairingId
	if err := r.SetPathParam("pairingId", o.PairingID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
