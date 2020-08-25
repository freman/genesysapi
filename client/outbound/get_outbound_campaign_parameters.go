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

// NewGetOutboundCampaignParams creates a new GetOutboundCampaignParams object
// with the default values initialized.
func NewGetOutboundCampaignParams() *GetOutboundCampaignParams {
	var ()
	return &GetOutboundCampaignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundCampaignParamsWithTimeout creates a new GetOutboundCampaignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundCampaignParamsWithTimeout(timeout time.Duration) *GetOutboundCampaignParams {
	var ()
	return &GetOutboundCampaignParams{

		timeout: timeout,
	}
}

// NewGetOutboundCampaignParamsWithContext creates a new GetOutboundCampaignParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundCampaignParamsWithContext(ctx context.Context) *GetOutboundCampaignParams {
	var ()
	return &GetOutboundCampaignParams{

		Context: ctx,
	}
}

// NewGetOutboundCampaignParamsWithHTTPClient creates a new GetOutboundCampaignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundCampaignParamsWithHTTPClient(client *http.Client) *GetOutboundCampaignParams {
	var ()
	return &GetOutboundCampaignParams{
		HTTPClient: client,
	}
}

/*GetOutboundCampaignParams contains all the parameters to send to the API endpoint
for the get outbound campaign operation typically these are written to a http.Request
*/
type GetOutboundCampaignParams struct {

	/*CampaignID
	  Campaign ID

	*/
	CampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound campaign params
func (o *GetOutboundCampaignParams) WithTimeout(timeout time.Duration) *GetOutboundCampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound campaign params
func (o *GetOutboundCampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound campaign params
func (o *GetOutboundCampaignParams) WithContext(ctx context.Context) *GetOutboundCampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound campaign params
func (o *GetOutboundCampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound campaign params
func (o *GetOutboundCampaignParams) WithHTTPClient(client *http.Client) *GetOutboundCampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound campaign params
func (o *GetOutboundCampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the get outbound campaign params
func (o *GetOutboundCampaignParams) WithCampaignID(campaignID string) *GetOutboundCampaignParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the get outbound campaign params
func (o *GetOutboundCampaignParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundCampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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