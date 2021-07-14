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

// NewGetConversationsMessagingIntegrationsOpenParams creates a new GetConversationsMessagingIntegrationsOpenParams object
// with the default values initialized.
func NewGetConversationsMessagingIntegrationsOpenParams() *GetConversationsMessagingIntegrationsOpenParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingIntegrationsOpenParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessagingIntegrationsOpenParamsWithTimeout creates a new GetConversationsMessagingIntegrationsOpenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsMessagingIntegrationsOpenParamsWithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsOpenParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingIntegrationsOpenParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetConversationsMessagingIntegrationsOpenParamsWithContext creates a new GetConversationsMessagingIntegrationsOpenParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsMessagingIntegrationsOpenParamsWithContext(ctx context.Context) *GetConversationsMessagingIntegrationsOpenParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingIntegrationsOpenParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetConversationsMessagingIntegrationsOpenParamsWithHTTPClient creates a new GetConversationsMessagingIntegrationsOpenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsMessagingIntegrationsOpenParamsWithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsOpenParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingIntegrationsOpenParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetConversationsMessagingIntegrationsOpenParams contains all the parameters to send to the API endpoint
for the get conversations messaging integrations open operation typically these are written to a http.Request
*/
type GetConversationsMessagingIntegrationsOpenParams struct {

	/*Expand
	  Expand instructions for the return value.

	*/
	Expand *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*SupportedContentID
	  Filter integrations returned based on the supported content ID

	*/
	SupportedContentID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) WithTimeout(timeout time.Duration) *GetConversationsMessagingIntegrationsOpenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) WithContext(ctx context.Context) *GetConversationsMessagingIntegrationsOpenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) WithHTTPClient(client *http.Client) *GetConversationsMessagingIntegrationsOpenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) WithExpand(expand *string) *GetConversationsMessagingIntegrationsOpenParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) SetExpand(expand *string) {
	o.Expand = expand
}

// WithPageNumber adds the pageNumber to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) WithPageNumber(pageNumber *int32) *GetConversationsMessagingIntegrationsOpenParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) WithPageSize(pageSize *int32) *GetConversationsMessagingIntegrationsOpenParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSupportedContentID adds the supportedContentID to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) WithSupportedContentID(supportedContentID *string) *GetConversationsMessagingIntegrationsOpenParams {
	o.SetSupportedContentID(supportedContentID)
	return o
}

// SetSupportedContentID adds the supportedContentId to the get conversations messaging integrations open params
func (o *GetConversationsMessagingIntegrationsOpenParams) SetSupportedContentID(supportedContentID *string) {
	o.SupportedContentID = supportedContentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessagingIntegrationsOpenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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