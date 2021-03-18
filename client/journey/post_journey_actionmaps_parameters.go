// Code generated by go-swagger; DO NOT EDIT.

package journey

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

// NewPostJourneyActionmapsParams creates a new PostJourneyActionmapsParams object
// with the default values initialized.
func NewPostJourneyActionmapsParams() *PostJourneyActionmapsParams {
	var ()
	return &PostJourneyActionmapsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostJourneyActionmapsParamsWithTimeout creates a new PostJourneyActionmapsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostJourneyActionmapsParamsWithTimeout(timeout time.Duration) *PostJourneyActionmapsParams {
	var ()
	return &PostJourneyActionmapsParams{

		timeout: timeout,
	}
}

// NewPostJourneyActionmapsParamsWithContext creates a new PostJourneyActionmapsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostJourneyActionmapsParamsWithContext(ctx context.Context) *PostJourneyActionmapsParams {
	var ()
	return &PostJourneyActionmapsParams{

		Context: ctx,
	}
}

// NewPostJourneyActionmapsParamsWithHTTPClient creates a new PostJourneyActionmapsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostJourneyActionmapsParamsWithHTTPClient(client *http.Client) *PostJourneyActionmapsParams {
	var ()
	return &PostJourneyActionmapsParams{
		HTTPClient: client,
	}
}

/*PostJourneyActionmapsParams contains all the parameters to send to the API endpoint
for the post journey actionmaps operation typically these are written to a http.Request
*/
type PostJourneyActionmapsParams struct {

	/*Body*/
	Body *models.ActionMap

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) WithTimeout(timeout time.Duration) *PostJourneyActionmapsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) WithContext(ctx context.Context) *PostJourneyActionmapsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) WithHTTPClient(client *http.Client) *PostJourneyActionmapsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) WithBody(body *models.ActionMap) *PostJourneyActionmapsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post journey actionmaps params
func (o *PostJourneyActionmapsParams) SetBody(body *models.ActionMap) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostJourneyActionmapsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}