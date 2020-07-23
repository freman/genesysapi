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

// NewGetQualityFormsSurveysBulkContextsParams creates a new GetQualityFormsSurveysBulkContextsParams object
// with the default values initialized.
func NewGetQualityFormsSurveysBulkContextsParams() *GetQualityFormsSurveysBulkContextsParams {
	var (
		publishedDefault = bool(true)
	)
	return &GetQualityFormsSurveysBulkContextsParams{
		Published: &publishedDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetQualityFormsSurveysBulkContextsParamsWithTimeout creates a new GetQualityFormsSurveysBulkContextsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetQualityFormsSurveysBulkContextsParamsWithTimeout(timeout time.Duration) *GetQualityFormsSurveysBulkContextsParams {
	var (
		publishedDefault = bool(true)
	)
	return &GetQualityFormsSurveysBulkContextsParams{
		Published: &publishedDefault,

		timeout: timeout,
	}
}

// NewGetQualityFormsSurveysBulkContextsParamsWithContext creates a new GetQualityFormsSurveysBulkContextsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetQualityFormsSurveysBulkContextsParamsWithContext(ctx context.Context) *GetQualityFormsSurveysBulkContextsParams {
	var (
		publishedDefault = bool(true)
	)
	return &GetQualityFormsSurveysBulkContextsParams{
		Published: &publishedDefault,

		Context: ctx,
	}
}

// NewGetQualityFormsSurveysBulkContextsParamsWithHTTPClient creates a new GetQualityFormsSurveysBulkContextsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetQualityFormsSurveysBulkContextsParamsWithHTTPClient(client *http.Client) *GetQualityFormsSurveysBulkContextsParams {
	var (
		publishedDefault = bool(true)
	)
	return &GetQualityFormsSurveysBulkContextsParams{
		Published:  &publishedDefault,
		HTTPClient: client,
	}
}

/*GetQualityFormsSurveysBulkContextsParams contains all the parameters to send to the API endpoint
for the get quality forms surveys bulk contexts operation typically these are written to a http.Request
*/
type GetQualityFormsSurveysBulkContextsParams struct {

	/*ContextID
	  A comma-delimited list of valid survey form context ids

	*/
	ContextID []string
	/*Published
	  If true, the latest published version will be included. If false, only the unpublished version will be included.

	*/
	Published *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) WithTimeout(timeout time.Duration) *GetQualityFormsSurveysBulkContextsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) WithContext(ctx context.Context) *GetQualityFormsSurveysBulkContextsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) WithHTTPClient(client *http.Client) *GetQualityFormsSurveysBulkContextsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContextID adds the contextID to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) WithContextID(contextID []string) *GetQualityFormsSurveysBulkContextsParams {
	o.SetContextID(contextID)
	return o
}

// SetContextID adds the contextId to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) SetContextID(contextID []string) {
	o.ContextID = contextID
}

// WithPublished adds the published to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) WithPublished(published *bool) *GetQualityFormsSurveysBulkContextsParams {
	o.SetPublished(published)
	return o
}

// SetPublished adds the published to the get quality forms surveys bulk contexts params
func (o *GetQualityFormsSurveysBulkContextsParams) SetPublished(published *bool) {
	o.Published = published
}

// WriteToRequest writes these params to a swagger request
func (o *GetQualityFormsSurveysBulkContextsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesContextID := o.ContextID

	joinedContextID := swag.JoinByFormat(valuesContextID, "multi")
	// query array param contextId
	if err := r.SetQueryParam("contextId", joinedContextID...); err != nil {
		return err
	}

	if o.Published != nil {

		// query param published
		var qrPublished bool
		if o.Published != nil {
			qrPublished = *o.Published
		}
		qPublished := swag.FormatBool(qrPublished)
		if qPublished != "" {
			if err := r.SetQueryParam("published", qPublished); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
