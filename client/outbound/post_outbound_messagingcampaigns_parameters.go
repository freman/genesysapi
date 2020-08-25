// Code generated by go-swagger; DO NOT EDIT.

package outbound

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

// NewPostOutboundMessagingcampaignsParams creates a new PostOutboundMessagingcampaignsParams object
// with the default values initialized.
func NewPostOutboundMessagingcampaignsParams() *PostOutboundMessagingcampaignsParams {
	var ()
	return &PostOutboundMessagingcampaignsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostOutboundMessagingcampaignsParamsWithTimeout creates a new PostOutboundMessagingcampaignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostOutboundMessagingcampaignsParamsWithTimeout(timeout time.Duration) *PostOutboundMessagingcampaignsParams {
	var ()
	return &PostOutboundMessagingcampaignsParams{

		timeout: timeout,
	}
}

// NewPostOutboundMessagingcampaignsParamsWithContext creates a new PostOutboundMessagingcampaignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostOutboundMessagingcampaignsParamsWithContext(ctx context.Context) *PostOutboundMessagingcampaignsParams {
	var ()
	return &PostOutboundMessagingcampaignsParams{

		Context: ctx,
	}
}

// NewPostOutboundMessagingcampaignsParamsWithHTTPClient creates a new PostOutboundMessagingcampaignsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostOutboundMessagingcampaignsParamsWithHTTPClient(client *http.Client) *PostOutboundMessagingcampaignsParams {
	var ()
	return &PostOutboundMessagingcampaignsParams{
		HTTPClient: client,
	}
}

/*PostOutboundMessagingcampaignsParams contains all the parameters to send to the API endpoint
for the post outbound messagingcampaigns operation typically these are written to a http.Request
*/
type PostOutboundMessagingcampaignsParams struct {

	/*Body
	  Messaging Campaign

	*/
	Body *models.MessagingCampaign

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) WithTimeout(timeout time.Duration) *PostOutboundMessagingcampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) WithContext(ctx context.Context) *PostOutboundMessagingcampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) WithHTTPClient(client *http.Client) *PostOutboundMessagingcampaignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) WithBody(body *models.MessagingCampaign) *PostOutboundMessagingcampaignsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post outbound messagingcampaigns params
func (o *PostOutboundMessagingcampaignsParams) SetBody(body *models.MessagingCampaign) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostOutboundMessagingcampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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