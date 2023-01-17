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

// NewPostOutboundCampaignCallbackScheduleParams creates a new PostOutboundCampaignCallbackScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOutboundCampaignCallbackScheduleParams() *PostOutboundCampaignCallbackScheduleParams {
	return &PostOutboundCampaignCallbackScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOutboundCampaignCallbackScheduleParamsWithTimeout creates a new PostOutboundCampaignCallbackScheduleParams object
// with the ability to set a timeout on a request.
func NewPostOutboundCampaignCallbackScheduleParamsWithTimeout(timeout time.Duration) *PostOutboundCampaignCallbackScheduleParams {
	return &PostOutboundCampaignCallbackScheduleParams{
		timeout: timeout,
	}
}

// NewPostOutboundCampaignCallbackScheduleParamsWithContext creates a new PostOutboundCampaignCallbackScheduleParams object
// with the ability to set a context for a request.
func NewPostOutboundCampaignCallbackScheduleParamsWithContext(ctx context.Context) *PostOutboundCampaignCallbackScheduleParams {
	return &PostOutboundCampaignCallbackScheduleParams{
		Context: ctx,
	}
}

// NewPostOutboundCampaignCallbackScheduleParamsWithHTTPClient creates a new PostOutboundCampaignCallbackScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOutboundCampaignCallbackScheduleParamsWithHTTPClient(client *http.Client) *PostOutboundCampaignCallbackScheduleParams {
	return &PostOutboundCampaignCallbackScheduleParams{
		HTTPClient: client,
	}
}

/*
PostOutboundCampaignCallbackScheduleParams contains all the parameters to send to the API endpoint

	for the post outbound campaign callback schedule operation.

	Typically these are written to a http.Request.
*/
type PostOutboundCampaignCallbackScheduleParams struct {

	/* Body.

	   ContactCallbackRequest
	*/
	Body *models.ContactCallbackRequest

	/* CampaignID.

	   Campaign ID
	*/
	CampaignID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post outbound campaign callback schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOutboundCampaignCallbackScheduleParams) WithDefaults() *PostOutboundCampaignCallbackScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post outbound campaign callback schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOutboundCampaignCallbackScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) WithTimeout(timeout time.Duration) *PostOutboundCampaignCallbackScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) WithContext(ctx context.Context) *PostOutboundCampaignCallbackScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) WithHTTPClient(client *http.Client) *PostOutboundCampaignCallbackScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) WithBody(body *models.ContactCallbackRequest) *PostOutboundCampaignCallbackScheduleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) SetBody(body *models.ContactCallbackRequest) {
	o.Body = body
}

// WithCampaignID adds the campaignID to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) WithCampaignID(campaignID string) *PostOutboundCampaignCallbackScheduleParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the post outbound campaign callback schedule params
func (o *PostOutboundCampaignCallbackScheduleParams) SetCampaignID(campaignID string) {
	o.CampaignID = campaignID
}

// WriteToRequest writes these params to a swagger request
func (o *PostOutboundCampaignCallbackScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
