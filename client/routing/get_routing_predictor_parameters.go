// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingPredictorParams creates a new GetRoutingPredictorParams object
// with the default values initialized.
func NewGetRoutingPredictorParams() *GetRoutingPredictorParams {
	var ()
	return &GetRoutingPredictorParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingPredictorParamsWithTimeout creates a new GetRoutingPredictorParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingPredictorParamsWithTimeout(timeout time.Duration) *GetRoutingPredictorParams {
	var ()
	return &GetRoutingPredictorParams{

		timeout: timeout,
	}
}

// NewGetRoutingPredictorParamsWithContext creates a new GetRoutingPredictorParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingPredictorParamsWithContext(ctx context.Context) *GetRoutingPredictorParams {
	var ()
	return &GetRoutingPredictorParams{

		Context: ctx,
	}
}

// NewGetRoutingPredictorParamsWithHTTPClient creates a new GetRoutingPredictorParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingPredictorParamsWithHTTPClient(client *http.Client) *GetRoutingPredictorParams {
	var ()
	return &GetRoutingPredictorParams{
		HTTPClient: client,
	}
}

/*GetRoutingPredictorParams contains all the parameters to send to the API endpoint
for the get routing predictor operation typically these are written to a http.Request
*/
type GetRoutingPredictorParams struct {

	/*PredictorID
	  Predictor ID

	*/
	PredictorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing predictor params
func (o *GetRoutingPredictorParams) WithTimeout(timeout time.Duration) *GetRoutingPredictorParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing predictor params
func (o *GetRoutingPredictorParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing predictor params
func (o *GetRoutingPredictorParams) WithContext(ctx context.Context) *GetRoutingPredictorParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing predictor params
func (o *GetRoutingPredictorParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing predictor params
func (o *GetRoutingPredictorParams) WithHTTPClient(client *http.Client) *GetRoutingPredictorParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing predictor params
func (o *GetRoutingPredictorParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPredictorID adds the predictorID to the get routing predictor params
func (o *GetRoutingPredictorParams) WithPredictorID(predictorID string) *GetRoutingPredictorParams {
	o.SetPredictorID(predictorID)
	return o
}

// SetPredictorID adds the predictorId to the get routing predictor params
func (o *GetRoutingPredictorParams) SetPredictorID(predictorID string) {
	o.PredictorID = predictorID
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingPredictorParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param predictorId
	if err := r.SetPathParam("predictorId", o.PredictorID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
