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

// NewGetConversationsMessagingStickerParams creates a new GetConversationsMessagingStickerParams object
// with the default values initialized.
func NewGetConversationsMessagingStickerParams() *GetConversationsMessagingStickerParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingStickerParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetConversationsMessagingStickerParamsWithTimeout creates a new GetConversationsMessagingStickerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetConversationsMessagingStickerParamsWithTimeout(timeout time.Duration) *GetConversationsMessagingStickerParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingStickerParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetConversationsMessagingStickerParamsWithContext creates a new GetConversationsMessagingStickerParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetConversationsMessagingStickerParamsWithContext(ctx context.Context) *GetConversationsMessagingStickerParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingStickerParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetConversationsMessagingStickerParamsWithHTTPClient creates a new GetConversationsMessagingStickerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetConversationsMessagingStickerParamsWithHTTPClient(client *http.Client) *GetConversationsMessagingStickerParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetConversationsMessagingStickerParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetConversationsMessagingStickerParams contains all the parameters to send to the API endpoint
for the get conversations messaging sticker operation typically these are written to a http.Request
*/
type GetConversationsMessagingStickerParams struct {

	/*MessengerType
	  Messenger Type

	*/
	MessengerType string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) WithTimeout(timeout time.Duration) *GetConversationsMessagingStickerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) WithContext(ctx context.Context) *GetConversationsMessagingStickerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) WithHTTPClient(client *http.Client) *GetConversationsMessagingStickerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMessengerType adds the messengerType to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) WithMessengerType(messengerType string) *GetConversationsMessagingStickerParams {
	o.SetMessengerType(messengerType)
	return o
}

// SetMessengerType adds the messengerType to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) SetMessengerType(messengerType string) {
	o.MessengerType = messengerType
}

// WithPageNumber adds the pageNumber to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) WithPageNumber(pageNumber *int32) *GetConversationsMessagingStickerParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) WithPageSize(pageSize *int32) *GetConversationsMessagingStickerParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get conversations messaging sticker params
func (o *GetConversationsMessagingStickerParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetConversationsMessagingStickerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param messengerType
	if err := r.SetPathParam("messengerType", o.MessengerType); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
