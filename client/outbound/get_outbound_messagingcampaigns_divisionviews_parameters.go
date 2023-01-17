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
	"github.com/go-openapi/swag"
)

// NewGetOutboundMessagingcampaignsDivisionviewsParams creates a new GetOutboundMessagingcampaignsDivisionviewsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOutboundMessagingcampaignsDivisionviewsParams() *GetOutboundMessagingcampaignsDivisionviewsParams {
	return &GetOutboundMessagingcampaignsDivisionviewsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewsParamsWithTimeout creates a new GetOutboundMessagingcampaignsDivisionviewsParams object
// with the ability to set a timeout on a request.
func NewGetOutboundMessagingcampaignsDivisionviewsParamsWithTimeout(timeout time.Duration) *GetOutboundMessagingcampaignsDivisionviewsParams {
	return &GetOutboundMessagingcampaignsDivisionviewsParams{
		timeout: timeout,
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewsParamsWithContext creates a new GetOutboundMessagingcampaignsDivisionviewsParams object
// with the ability to set a context for a request.
func NewGetOutboundMessagingcampaignsDivisionviewsParamsWithContext(ctx context.Context) *GetOutboundMessagingcampaignsDivisionviewsParams {
	return &GetOutboundMessagingcampaignsDivisionviewsParams{
		Context: ctx,
	}
}

// NewGetOutboundMessagingcampaignsDivisionviewsParamsWithHTTPClient creates a new GetOutboundMessagingcampaignsDivisionviewsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOutboundMessagingcampaignsDivisionviewsParamsWithHTTPClient(client *http.Client) *GetOutboundMessagingcampaignsDivisionviewsParams {
	return &GetOutboundMessagingcampaignsDivisionviewsParams{
		HTTPClient: client,
	}
}

/*
GetOutboundMessagingcampaignsDivisionviewsParams contains all the parameters to send to the API endpoint

	for the get outbound messagingcampaigns divisionviews operation.

	Typically these are written to a http.Request.
*/
type GetOutboundMessagingcampaignsDivisionviewsParams struct {

	/* ID.

	   id
	*/
	ID []string

	/* Name.

	   Name
	*/
	Name *string

	/* PageNumber.

	   Page number

	   Format: int32
	   Default: 1
	*/
	PageNumber *int32

	/* PageSize.

	   Page size. The max that will be returned is 100.

	   Format: int32
	   Default: 25
	*/
	PageSize *int32

	/* SenderSmsPhoneNumber.

	   Sender SMS Phone Number
	*/
	SenderSmsPhoneNumber *string

	/* SortOrder.

	   The direction to sort

	   Default: "a"
	*/
	SortOrder *string

	/* Type.

	   Campaign Type
	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get outbound messagingcampaigns divisionviews params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithDefaults() *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get outbound messagingcampaigns divisionviews params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortOrderDefault = string("a")
	)

	val := GetOutboundMessagingcampaignsDivisionviewsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithTimeout(timeout time.Duration) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithContext(ctx context.Context) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithHTTPClient(client *http.Client) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithID(id []string) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithName(name *string) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithPageNumber(pageNumber *int32) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithPageSize(pageSize *int32) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSenderSmsPhoneNumber adds the senderSmsPhoneNumber to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithSenderSmsPhoneNumber(senderSmsPhoneNumber *string) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetSenderSmsPhoneNumber(senderSmsPhoneNumber)
	return o
}

// SetSenderSmsPhoneNumber adds the senderSmsPhoneNumber to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetSenderSmsPhoneNumber(senderSmsPhoneNumber *string) {
	o.SenderSmsPhoneNumber = senderSmsPhoneNumber
}

// WithSortOrder adds the sortOrder to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithSortOrder(sortOrder *string) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithType adds the typeVar to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WithType(typeVar *string) *GetOutboundMessagingcampaignsDivisionviewsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get outbound messagingcampaigns divisionviews params
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// binding items for id
		joinedID := o.bindParamID(reg)

		// query array param id
		if err := r.SetQueryParam("id", joinedID...); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
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

	if o.SenderSmsPhoneNumber != nil {

		// query param senderSmsPhoneNumber
		var qrSenderSmsPhoneNumber string

		if o.SenderSmsPhoneNumber != nil {
			qrSenderSmsPhoneNumber = *o.SenderSmsPhoneNumber
		}
		qSenderSmsPhoneNumber := qrSenderSmsPhoneNumber
		if qSenderSmsPhoneNumber != "" {

			if err := r.SetQueryParam("senderSmsPhoneNumber", qSenderSmsPhoneNumber); err != nil {
				return err
			}
		}
	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string

		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {

			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}
	}

	if o.Type != nil {

		// query param type
		var qrType string

		if o.Type != nil {
			qrType = *o.Type
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetOutboundMessagingcampaignsDivisionviews binds the parameter id
func (o *GetOutboundMessagingcampaignsDivisionviewsParams) bindParamID(formats strfmt.Registry) []string {
	iDIR := o.ID

	var iDIC []string
	for _, iDIIR := range iDIR { // explode []string

		iDIIV := iDIIR // string as string
		iDIC = append(iDIC, iDIIV)
	}

	// items.CollectionFormat: "multi"
	iDIS := swag.JoinByFormat(iDIC, "multi")

	return iDIS
}
