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

// NewGetQualityFormsSurveysBulkParams creates a new GetQualityFormsSurveysBulkParams object
// with the default values initialized.
func NewGetQualityFormsSurveysBulkParams() *GetQualityFormsSurveysBulkParams {
	var ()
	return &GetQualityFormsSurveysBulkParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityFormsSurveysBulkParamsWithTimeout creates a new GetQualityFormsSurveysBulkParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQualityFormsSurveysBulkParamsWithTimeout(timeout time.Duration) *GetQualityFormsSurveysBulkParams {
	var ()
	return &GetQualityFormsSurveysBulkParams{

		timeout: timeout,
	}
}

// NewGetQualityFormsSurveysBulkParamsWithContext creates a new GetQualityFormsSurveysBulkParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQualityFormsSurveysBulkParamsWithContext(ctx context.Context) *GetQualityFormsSurveysBulkParams {
	var ()
	return &GetQualityFormsSurveysBulkParams{

		Context: ctx,
	}
}

// NewGetQualityFormsSurveysBulkParamsWithHTTPClient creates a new GetQualityFormsSurveysBulkParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQualityFormsSurveysBulkParamsWithHTTPClient(client *http.Client) *GetQualityFormsSurveysBulkParams {
	var ()
	return &GetQualityFormsSurveysBulkParams{
		HTTPClient: client,
	}
}

/*GetQualityFormsSurveysBulkParams contains all the parameters to send to the API endpoint
for the get quality forms surveys bulk operation typically these are written to a http.Request
*/
type GetQualityFormsSurveysBulkParams struct {

	/*ID
	  A comma-delimited list of valid survey form ids

	*/
	ID []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) WithTimeout(timeout time.Duration) *GetQualityFormsSurveysBulkParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) WithContext(ctx context.Context) *GetQualityFormsSurveysBulkParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) WithHTTPClient(client *http.Client) *GetQualityFormsSurveysBulkParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) WithID(id []string) *GetQualityFormsSurveysBulkParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get quality forms surveys bulk params
func (o *GetQualityFormsSurveysBulkParams) SetID(id []string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityFormsSurveysBulkParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}