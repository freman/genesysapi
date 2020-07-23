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

	"github.com/freman/genesysapi/models"
)

// NewPostFlowsExecutionsParams creates a new PostFlowsExecutionsParams object
// with the default values initialized.
func NewPostFlowsExecutionsParams() *PostFlowsExecutionsParams {
	var ()
	return &PostFlowsExecutionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFlowsExecutionsParamsWithTimeout creates a new PostFlowsExecutionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFlowsExecutionsParamsWithTimeout(timeout time.Duration) *PostFlowsExecutionsParams {
	var ()
	return &PostFlowsExecutionsParams{

		timeout: timeout,
	}
}

// NewPostFlowsExecutionsParamsWithContext creates a new PostFlowsExecutionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFlowsExecutionsParamsWithContext(ctx context.Context) *PostFlowsExecutionsParams {
	var ()
	return &PostFlowsExecutionsParams{

		Context: ctx,
	}
}

// NewPostFlowsExecutionsParamsWithHTTPClient creates a new PostFlowsExecutionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFlowsExecutionsParamsWithHTTPClient(client *http.Client) *PostFlowsExecutionsParams {
	var ()
	return &PostFlowsExecutionsParams{
		HTTPClient: client,
	}
}

/*PostFlowsExecutionsParams contains all the parameters to send to the API endpoint
for the post flows executions operation typically these are written to a http.Request
*/
type PostFlowsExecutionsParams struct {

	/*FlowLaunchRequest*/
	FlowLaunchRequest *models.FlowExecutionLaunchRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post flows executions params
func (o *PostFlowsExecutionsParams) WithTimeout(timeout time.Duration) *PostFlowsExecutionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post flows executions params
func (o *PostFlowsExecutionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post flows executions params
func (o *PostFlowsExecutionsParams) WithContext(ctx context.Context) *PostFlowsExecutionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post flows executions params
func (o *PostFlowsExecutionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post flows executions params
func (o *PostFlowsExecutionsParams) WithHTTPClient(client *http.Client) *PostFlowsExecutionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post flows executions params
func (o *PostFlowsExecutionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFlowLaunchRequest adds the flowLaunchRequest to the post flows executions params
func (o *PostFlowsExecutionsParams) WithFlowLaunchRequest(flowLaunchRequest *models.FlowExecutionLaunchRequest) *PostFlowsExecutionsParams {
	o.SetFlowLaunchRequest(flowLaunchRequest)
	return o
}

// SetFlowLaunchRequest adds the flowLaunchRequest to the post flows executions params
func (o *PostFlowsExecutionsParams) SetFlowLaunchRequest(flowLaunchRequest *models.FlowExecutionLaunchRequest) {
	o.FlowLaunchRequest = flowLaunchRequest
}

// WriteToRequest writes these params to a swagger request
func (o *PostFlowsExecutionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FlowLaunchRequest != nil {
		if err := r.SetBodyParam(o.FlowLaunchRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
