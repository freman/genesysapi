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

// NewGetOutboundMessagingcampaignsParams creates a new GetOutboundMessagingcampaignsParams object
// with the default values initialized.
func NewGetOutboundMessagingcampaignsParams() *GetOutboundMessagingcampaignsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ascending")
	)
	return &GetOutboundMessagingcampaignsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundMessagingcampaignsParamsWithTimeout creates a new GetOutboundMessagingcampaignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundMessagingcampaignsParamsWithTimeout(timeout time.Duration) *GetOutboundMessagingcampaignsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ascending")
	)
	return &GetOutboundMessagingcampaignsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetOutboundMessagingcampaignsParamsWithContext creates a new GetOutboundMessagingcampaignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundMessagingcampaignsParamsWithContext(ctx context.Context) *GetOutboundMessagingcampaignsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ascending")
	)
	return &GetOutboundMessagingcampaignsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetOutboundMessagingcampaignsParamsWithHTTPClient creates a new GetOutboundMessagingcampaignsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundMessagingcampaignsParamsWithHTTPClient(client *http.Client) *GetOutboundMessagingcampaignsParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortByDefault     = string("name")
		sortOrderDefault  = string("ascending")
	)
	return &GetOutboundMessagingcampaignsParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetOutboundMessagingcampaignsParams contains all the parameters to send to the API endpoint
for the get outbound messagingcampaigns operation typically these are written to a http.Request
*/
type GetOutboundMessagingcampaignsParams struct {

	/*ContactListID
	  Contact List ID

	*/
	ContactListID *string
	/*DivisionID
	  Division ID(s)

	*/
	DivisionID []string
	/*ID
	  A list of messaging campaign ids to bulk fetch

	*/
	ID []string
	/*Name
	  Name

	*/
	Name *string
	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size. The max that will be returned is 100.

	*/
	PageSize *int32
	/*SenderSmsPhoneNumber
	  Sender SMS Phone Number

	*/
	SenderSmsPhoneNumber *string
	/*SortBy
	  The field to sort by

	*/
	SortBy *string
	/*SortOrder
	  The direction to sort

	*/
	SortOrder *string
	/*Type
	  Campaign Type

	*/
	Type *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithTimeout(timeout time.Duration) *GetOutboundMessagingcampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithContext(ctx context.Context) *GetOutboundMessagingcampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithHTTPClient(client *http.Client) *GetOutboundMessagingcampaignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContactListID adds the contactListID to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithContactListID(contactListID *string) *GetOutboundMessagingcampaignsParams {
	o.SetContactListID(contactListID)
	return o
}

// SetContactListID adds the contactListId to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetContactListID(contactListID *string) {
	o.ContactListID = contactListID
}

// WithDivisionID adds the divisionID to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithDivisionID(divisionID []string) *GetOutboundMessagingcampaignsParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithID adds the id to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithID(id []string) *GetOutboundMessagingcampaignsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithName(name *string) *GetOutboundMessagingcampaignsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithPageNumber(pageNumber *int32) *GetOutboundMessagingcampaignsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithPageSize(pageSize *int32) *GetOutboundMessagingcampaignsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSenderSmsPhoneNumber adds the senderSmsPhoneNumber to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithSenderSmsPhoneNumber(senderSmsPhoneNumber *string) *GetOutboundMessagingcampaignsParams {
	o.SetSenderSmsPhoneNumber(senderSmsPhoneNumber)
	return o
}

// SetSenderSmsPhoneNumber adds the senderSmsPhoneNumber to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetSenderSmsPhoneNumber(senderSmsPhoneNumber *string) {
	o.SenderSmsPhoneNumber = senderSmsPhoneNumber
}

// WithSortBy adds the sortBy to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithSortBy(sortBy *string) *GetOutboundMessagingcampaignsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithSortOrder(sortOrder *string) *GetOutboundMessagingcampaignsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithType adds the typeVar to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) WithType(typeVar *string) *GetOutboundMessagingcampaignsParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the get outbound messagingcampaigns params
func (o *GetOutboundMessagingcampaignsParams) SetType(typeVar *string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundMessagingcampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContactListID != nil {

		// query param contactListId
		var qrContactListID string
		if o.ContactListID != nil {
			qrContactListID = *o.ContactListID
		}
		qContactListID := qrContactListID
		if qContactListID != "" {
			if err := r.SetQueryParam("contactListId", qContactListID); err != nil {
				return err
			}
		}

	}

	valuesDivisionID := o.DivisionID

	joinedDivisionID := swag.JoinByFormat(valuesDivisionID, "multi")
	// query array param divisionId
	if err := r.SetQueryParam("divisionId", joinedDivisionID...); err != nil {
		return err
	}

	valuesID := o.ID

	joinedID := swag.JoinByFormat(valuesID, "multi")
	// query array param id
	if err := r.SetQueryParam("id", joinedID...); err != nil {
		return err
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

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
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