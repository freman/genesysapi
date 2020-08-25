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

// NewPutOutboundSchedulesCampaignParams creates a new PutOutboundSchedulesCampaignParams object
// with the default values initialized.
func NewPutOutboundSchedulesCampaignParams() *PutOutboundSchedulesCampaignParams {
	var ()
	return &PutOutboundSchedulesCampaignParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutOutboundSchedulesCampaignParamsWithTimeout creates a new PutOutboundSchedulesCampaignParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutOutboundSchedulesCampaignParamsWithTimeout(timeout time.Duration) *PutOutboundSchedulesCampaignParams {
	var ()
	return &PutOutboundSchedulesCampaignParams{

		timeout: timeout,
	}
}

// NewPutOutboundSchedulesCampaignParamsWithContext creates a new PutOutboundSchedulesCampaignParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutOutboundSchedulesCampaignParamsWithContext(ctx context.Context) *PutOutboundSchedulesCampaignParams {
	var ()
	return &PutOutboundSchedulesCampaignParams{

		Context: ctx,
	}
}

// NewPutOutboundSchedulesCampaignParamsWithHTTPClient creates a new PutOutboundSchedulesCampaignParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutOutboundSchedulesCampaignParamsWithHTTPClient(client *http.Client) *PutOutboundSchedulesCampaignParams {
	var ()
	return &PutOutboundSchedulesCampaignParams{
		HTTPClient: client,
	}
}

/*PutOutboundSchedulesCampaignParams contains all the parameters to send to the API endpoint
for the put outbound schedules campaign operation typically these are written to a http.Request
*/
type PutOutboundSchedulesCampaignParams struct {

	/*Body
	  CampaignSchedule

	*/
	Body *models.CampaignSchedule
	/*CampaignID
	  Campaign ID

	*/
	CampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) WithTimeout(timeout time.Duration) *PutOutboundSchedulesCampaignParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) WithContext(ctx context.Context) *PutOutboundSchedulesCampaignParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) WithHTTPClient(client *http.Client) *PutOutboundSchedulesCampaignParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) WithBody(body *models.CampaignSchedule) *PutOutboundSchedulesCampaignParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) SetBody(body *models.CampaignSchedule) {
	o.Body = body
}

// WithCampaignID adds the campaignID to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) WithCampaignID(campaignID string) *PutOutboundSchedulesCampaignParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the put outbound schedules campaign params
func (o *PutOutboundSchedulesCampaignParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WriteToRequest writes these params to a swagger request
func (o *PutOutboundSchedulesCampaignParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param campaignId
	if err := r.SetPathParam("campaignId", o.CampaignID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}