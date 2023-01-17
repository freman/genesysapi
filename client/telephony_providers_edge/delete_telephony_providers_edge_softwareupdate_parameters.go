// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewDeleteTelephonyProvidersEdgeSoftwareupdateParams creates a new DeleteTelephonyProvidersEdgeSoftwareupdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteTelephonyProvidersEdgeSoftwareupdateParams() *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateParamsWithTimeout creates a new DeleteTelephonyProvidersEdgeSoftwareupdateParams object
// with the ability to set a timeout on a request.
func NewDeleteTelephonyProvidersEdgeSoftwareupdateParamsWithTimeout(timeout time.Duration) *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateParams{
		timeout: timeout,
	}
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateParamsWithContext creates a new DeleteTelephonyProvidersEdgeSoftwareupdateParams object
// with the ability to set a context for a request.
func NewDeleteTelephonyProvidersEdgeSoftwareupdateParamsWithContext(ctx context.Context) *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateParams{
		Context: ctx,
	}
}

// NewDeleteTelephonyProvidersEdgeSoftwareupdateParamsWithHTTPClient creates a new DeleteTelephonyProvidersEdgeSoftwareupdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteTelephonyProvidersEdgeSoftwareupdateParamsWithHTTPClient(client *http.Client) *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	return &DeleteTelephonyProvidersEdgeSoftwareupdateParams{
		HTTPClient: client,
	}
}

/*
DeleteTelephonyProvidersEdgeSoftwareupdateParams contains all the parameters to send to the API endpoint

	for the delete telephony providers edge softwareupdate operation.

	Typically these are written to a http.Request.
*/
type DeleteTelephonyProvidersEdgeSoftwareupdateParams struct {

	/* EdgeID.

	   Edge ID
	*/
	EdgeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete telephony providers edge softwareupdate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) WithDefaults() *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete telephony providers edge softwareupdate params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) WithTimeout(timeout time.Duration) *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) WithContext(ctx context.Context) *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) WithHTTPClient(client *http.Client) *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEdgeID adds the edgeID to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) WithEdgeID(edgeID string) *DeleteTelephonyProvidersEdgeSoftwareupdateParams {
	o.SetEdgeID(edgeID)
	return o
}

// SetEdgeID adds the edgeId to the delete telephony providers edge softwareupdate params
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) SetEdgeID(edgeID string) {
	o.EdgeID = edgeID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteTelephonyProvidersEdgeSoftwareupdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param edgeId
	if err := r.SetPathParam("edgeId", o.EdgeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
