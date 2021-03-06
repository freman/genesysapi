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
)

// NewDeleteOutboundMessagingcampaignParams creates a new DeleteOutboundMessagingcampaignParams object
// with the default values initialized.
func NewDeleteOutboundMessagingcampaignParams() *DeleteOutboundMessagingcampaignParams {
	var ()
	return &DeleteOutboundMessagingcampaignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOutboundMessagingcampaignParamsWithTimeout creates a new DeleteOutboundMessagingcampaignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOutboundMessagingcampaignParamsWithTimeout(timeout time.Duration) *DeleteOutboundMessagingcampaignParams {
	var ()
	return &DeleteOutboundMessagingcampaignParams{

		timeout: timeout,
	}
}

// NewDeleteOutboundMessagingcampaignParamsWithContext creates a new DeleteOutboundMessagingcampaignParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOutboundMessagingcampaignParamsWithContext(ctx context.Context) *DeleteOutboundMessagingcampaignParams {
	var ()
	return &DeleteOutboundMessagingcampaignParams{

		Context: ctx,
	}
}

// NewDeleteOutboundMessagingcampaignParamsWithHTTPClient creates a new DeleteOutboundMessagingcampaignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOutboundMessagingcampaignParamsWithHTTPClient(client *http.Client) *DeleteOutboundMessagingcampaignParams {
	var ()
	return &DeleteOutboundMessagingcampaignParams{
		HTTPClient: client,
	}
}

/*DeleteOutboundMessagingcampaignParams contains all the parameters to send to the API endpoint
for the delete outbound messagingcampaign operation typically these are written to a http.Request
*/
type DeleteOutboundMessagingcampaignParams struct {

	/*MessagingCampaignID
	  The Messaging Campaign ID

	*/
	MessagingCampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) WithTimeout(timeout time.Duration) *DeleteOutboundMessagingcampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) WithContext(ctx context.Context) *DeleteOutboundMessagingcampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) WithHTTPClient(client *http.Client) *DeleteOutboundMessagingcampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessagingCampaignID adds the messagingCampaignID to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) WithMessagingCampaignID(messagingCampaignID string) *DeleteOutboundMessagingcampaignParams {
	o.SetMessagingCampaignID(messagingCampaignID)
	return o
}

// SetMessagingCampaignID adds the messagingCampaignId to the delete outbound messagingcampaign params
func (o *DeleteOutboundMessagingcampaignParams) SetMessagingCampaignID(messagingCampaignID string) {
	o.MessagingCampaignID = messagingCampaignID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOutboundMessagingcampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param messagingCampaignId
	if err := r.SetPathParam("messagingCampaignId", o.MessagingCampaignID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
