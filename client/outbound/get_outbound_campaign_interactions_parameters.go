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

// NewGetOutboundCampaignInteractionsParams creates a new GetOutboundCampaignInteractionsParams object
// with the default values initialized.
func NewGetOutboundCampaignInteractionsParams() *GetOutboundCampaignInteractionsParams {
	var ()
	return &GetOutboundCampaignInteractionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundCampaignInteractionsParamsWithTimeout creates a new GetOutboundCampaignInteractionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundCampaignInteractionsParamsWithTimeout(timeout time.Duration) *GetOutboundCampaignInteractionsParams {
	var ()
	return &GetOutboundCampaignInteractionsParams{

		timeout: timeout,
	}
}

// NewGetOutboundCampaignInteractionsParamsWithContext creates a new GetOutboundCampaignInteractionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundCampaignInteractionsParamsWithContext(ctx context.Context) *GetOutboundCampaignInteractionsParams {
	var ()
	return &GetOutboundCampaignInteractionsParams{

		Context: ctx,
	}
}

// NewGetOutboundCampaignInteractionsParamsWithHTTPClient creates a new GetOutboundCampaignInteractionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundCampaignInteractionsParamsWithHTTPClient(client *http.Client) *GetOutboundCampaignInteractionsParams {
	var ()
	return &GetOutboundCampaignInteractionsParams{
		HTTPClient: client,
	}
}

/*GetOutboundCampaignInteractionsParams contains all the parameters to send to the API endpoint
for the get outbound campaign interactions operation typically these are written to a http.Request
*/
type GetOutboundCampaignInteractionsParams struct {

	/*CampaignID
	  Campaign ID

	*/
	CampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) WithTimeout(timeout time.Duration) *GetOutboundCampaignInteractionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) WithContext(ctx context.Context) *GetOutboundCampaignInteractionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) WithHTTPClient(client *http.Client) *GetOutboundCampaignInteractionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) WithCampaignID(campaignID string) *GetOutboundCampaignInteractionsParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the get outbound campaign interactions params
func (o *GetOutboundCampaignInteractionsParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundCampaignInteractionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", o.CampaignID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
