// Code generated by go-swagger; DO NOT EDIT.

package recording

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

// NewDeleteRecordingMediaretentionpoliciesParams creates a new DeleteRecordingMediaretentionpoliciesParams object
// with the default values initialized.
func NewDeleteRecordingMediaretentionpoliciesParams() *DeleteRecordingMediaretentionpoliciesParams {
	var ()
	return &DeleteRecordingMediaretentionpoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteRecordingMediaretentionpoliciesParamsWithTimeout creates a new DeleteRecordingMediaretentionpoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteRecordingMediaretentionpoliciesParamsWithTimeout(timeout time.Duration) *DeleteRecordingMediaretentionpoliciesParams {
	var ()
	return &DeleteRecordingMediaretentionpoliciesParams{

		timeout: timeout,
	}
}

// NewDeleteRecordingMediaretentionpoliciesParamsWithContext creates a new DeleteRecordingMediaretentionpoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteRecordingMediaretentionpoliciesParamsWithContext(ctx context.Context) *DeleteRecordingMediaretentionpoliciesParams {
	var ()
	return &DeleteRecordingMediaretentionpoliciesParams{

		Context: ctx,
	}
}

// NewDeleteRecordingMediaretentionpoliciesParamsWithHTTPClient creates a new DeleteRecordingMediaretentionpoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteRecordingMediaretentionpoliciesParamsWithHTTPClient(client *http.Client) *DeleteRecordingMediaretentionpoliciesParams {
	var ()
	return &DeleteRecordingMediaretentionpoliciesParams{
		HTTPClient: client,
	}
}

/*DeleteRecordingMediaretentionpoliciesParams contains all the parameters to send to the API endpoint
for the delete recording mediaretentionpolicies operation typically these are written to a http.Request
*/
type DeleteRecordingMediaretentionpoliciesParams struct {

	/*Ids*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) WithTimeout(timeout time.Duration) *DeleteRecordingMediaretentionpoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) WithContext(ctx context.Context) *DeleteRecordingMediaretentionpoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) WithHTTPClient(client *http.Client) *DeleteRecordingMediaretentionpoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) WithIds(ids string) *DeleteRecordingMediaretentionpoliciesParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the delete recording mediaretentionpolicies params
func (o *DeleteRecordingMediaretentionpoliciesParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteRecordingMediaretentionpoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
