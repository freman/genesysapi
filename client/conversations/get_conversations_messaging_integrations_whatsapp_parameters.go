// Code generated by go-swagger; DO NOT EDIT.

package conversations

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

// NewGetConversationsMessagingIntegrationsWhatsappParams creates a new GetConversationsMessagingIntegrationsWhatsappParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetConversationsMessagingIntegrationsWhatsappParams() *GetConversationsMessagingIntegrationsWhatsappParams {
	return &GetConversationsMessagingIntegrationsWhatsappParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessagingIntegrationsWhatsappParamsWithTimeout creates a new GetConversationsMessagingIntegrationsWhatsappParams object
// with the ability to set a timeout on a request.
func NewGetConversationsMessagingIntegrationsWhatsappParamsWithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsWhatsappParams {
	return &GetConversationsMessagingIntegrationsWhatsappParams{
		timeout: timeout,
	}
}

// NewGetConversationsMessagingIntegrationsWhatsappParamsWithContext creates a new GetConversationsMessagingIntegrationsWhatsappParams object
// with the ability to set a context for a request.
func NewGetConversationsMessagingIntegrationsWhatsappParamsWithContext(ctx context.Context) *GetConversationsMessagingIntegrationsWhatsappParams {
	return &GetConversationsMessagingIntegrationsWhatsappParams{
		Context: ctx,
	}
}

// NewGetConversationsMessagingIntegrationsWhatsappParamsWithHTTPClient creates a new GetConversationsMessagingIntegrationsWhatsappParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetConversationsMessagingIntegrationsWhatsappParamsWithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsWhatsappParams {
	return &GetConversationsMessagingIntegrationsWhatsappParams{
		HTTPClient: client,
	}
}

/*
GetConversationsMessagingIntegrationsWhatsappParams contains all the parameters to send to the API endpoint

	for the get conversations messaging integrations whatsapp operation.

	Typically these are written to a http.Request.
*/
type GetConversationsMessagingIntegrationsWhatsappParams struct {

	/* Expand.

	   Expand instructions for the return value.
	*/
	Expand *string

	/* MessagingSettingID.

	   Filter integrations returned based on the setting ID
	*/
	MessagingSettingID *string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* SupportedContentID.

	   Filter integrations returned based on the supported content ID
	*/
	SupportedContentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get conversations messaging integrations whatsapp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithDefaults() *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get conversations messaging integrations whatsapp params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)
	)

	val := GetConversationsMessagingIntegrationsWhatsappParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithContext(ctx context.Context) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithExpand(expand *string) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithMessagingSettingID adds the messagingSettingID to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithMessagingSettingID(messagingSettingID *string) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetMessagingSettingID(messagingSettingID)
	return o
}

// SetMessagingSettingID adds the messagingSettingId to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetMessagingSettingID(messagingSettingID *string) {
	o.MessagingSettingID = messagingSettingID
}

// WithPageNumber adds the pageNumber to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithPageNumber(pageNumber *int32) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithPageSize(pageSize *int32) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSupportedContentID adds the supportedContentID to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WithSupportedContentID(supportedContentID *string) *GetConversationsMessagingIntegrationsWhatsappParams {
	o.SetSupportedContentID(supportedContentID)
	return o
}

// SetSupportedContentID adds the supportedContentId to the get conversations messaging integrations whatsapp params
func (o *GetConversationsMessagingIntegrationsWhatsappParams) SetSupportedContentID(supportedContentID *string) {
	o.SupportedContentID = supportedContentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessagingIntegrationsWhatsappParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// query param expand
		var qrExpand string

		if o.Expand != nil {
			qrExpand = *o.Expand
		}
		qExpand := qrExpand
		if qExpand != "" {

			if err := r.SetQueryParam("expand", qExpand); err != nil {
				return err
			}
		}
	}

	if o.MessagingSettingID != nil {

		// query param messagingSetting.id
		var qrMessagingSettingID string

		if o.MessagingSettingID != nil {
			qrMessagingSettingID = *o.MessagingSettingID
		}
		qMessagingSettingID := qrMessagingSettingID
		if qMessagingSettingID != "" {

			if err := r.SetQueryParam("messagingSetting.id", qMessagingSettingID); err != nil {
				return err
			}
		}
	}

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
	}

	if o.SupportedContentID != nil {

		// query param supportedContent.id
		var qrSupportedContentID string

		if o.SupportedContentID != nil {
			qrSupportedContentID = *o.SupportedContentID
		}
		qSupportedContentID := qrSupportedContentID
		if qSupportedContentID != "" {

			if err := r.SetQueryParam("supportedContent.id", qSupportedContentID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
