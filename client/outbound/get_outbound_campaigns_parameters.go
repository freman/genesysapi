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

// NewGetOutboundCampaignsParams creates a new GetOutboundCampaignsParams object
// with the default values initialized.
func NewGetOutboundCampaignsParams() *GetOutboundCampaignsParams {
	var (
		filterTypeDefault = string("Prefix")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("a")
	)
	return &GetOutboundCampaignsParams{
		FilterType: &filterTypeDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetOutboundCampaignsParamsWithTimeout creates a new GetOutboundCampaignsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetOutboundCampaignsParamsWithTimeout(timeout time.Duration) *GetOutboundCampaignsParams {
	var (
		filterTypeDefault = string("Prefix")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("a")
	)
	return &GetOutboundCampaignsParams{
		FilterType: &filterTypeDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		timeout: timeout,
	}
}

// NewGetOutboundCampaignsParamsWithContext creates a new GetOutboundCampaignsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetOutboundCampaignsParamsWithContext(ctx context.Context) *GetOutboundCampaignsParams {
	var (
		filterTypeDefault = string("Prefix")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("a")
	)
	return &GetOutboundCampaignsParams{
		FilterType: &filterTypeDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,

		Context: ctx,
	}
}

// NewGetOutboundCampaignsParamsWithHTTPClient creates a new GetOutboundCampaignsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetOutboundCampaignsParamsWithHTTPClient(client *http.Client) *GetOutboundCampaignsParams {
	var (
		filterTypeDefault = string("Prefix")
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
		sortOrderDefault  = string("a")
	)
	return &GetOutboundCampaignsParams{
		FilterType: &filterTypeDefault,
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortOrder:  &sortOrderDefault,
		HTTPClient: client,
	}
}

/*GetOutboundCampaignsParams contains all the parameters to send to the API endpoint
for the get outbound campaigns operation typically these are written to a http.Request
*/
type GetOutboundCampaignsParams struct {

	/*CallAnalysisResponseSetID
	  Call analysis response set ID

	*/
	CallAnalysisResponseSetID *string
	/*ContactListID
	  Contact List ID

	*/
	ContactListID *string
	/*DistributionQueueID
	  Distribution queue ID

	*/
	DistributionQueueID *string
	/*DivisionID
	  Division ID(s)

	*/
	DivisionID []string
	/*DncListIds
	  DNC list ID

	*/
	DncListIds *string
	/*EdgeGroupID
	  Edge group ID

	*/
	EdgeGroupID *string
	/*FilterType
	  Filter type

	*/
	FilterType *string
	/*ID
	  id

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
	/*SortBy
	  Sort by

	*/
	SortBy *string
	/*SortOrder
	  Sort order

	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithTimeout(timeout time.Duration) *GetOutboundCampaignsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithContext(ctx context.Context) *GetOutboundCampaignsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithHTTPClient(client *http.Client) *GetOutboundCampaignsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCallAnalysisResponseSetID adds the callAnalysisResponseSetID to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithCallAnalysisResponseSetID(callAnalysisResponseSetID *string) *GetOutboundCampaignsParams {
	o.SetCallAnalysisResponseSetID(callAnalysisResponseSetID)
	return o
}

// SetCallAnalysisResponseSetID adds the callAnalysisResponseSetId to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetCallAnalysisResponseSetID(callAnalysisResponseSetID *string) {
	o.CallAnalysisResponseSetID = callAnalysisResponseSetID
}

// WithContactListID adds the contactListID to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithContactListID(contactListID *string) *GetOutboundCampaignsParams {
	o.SetContactListID(contactListID)
	return o
}

// SetContactListID adds the contactListId to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetContactListID(contactListID *string) {
	o.ContactListID = contactListID
}

// WithDistributionQueueID adds the distributionQueueID to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithDistributionQueueID(distributionQueueID *string) *GetOutboundCampaignsParams {
	o.SetDistributionQueueID(distributionQueueID)
	return o
}

// SetDistributionQueueID adds the distributionQueueId to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetDistributionQueueID(distributionQueueID *string) {
	o.DistributionQueueID = distributionQueueID
}

// WithDivisionID adds the divisionID to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithDivisionID(divisionID []string) *GetOutboundCampaignsParams {
	o.SetDivisionID(divisionID)
	return o
}

// SetDivisionID adds the divisionId to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetDivisionID(divisionID []string) {
	o.DivisionID = divisionID
}

// WithDncListIds adds the dncListIds to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithDncListIds(dncListIds *string) *GetOutboundCampaignsParams {
	o.SetDncListIds(dncListIds)
	return o
}

// SetDncListIds adds the dncListIds to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetDncListIds(dncListIds *string) {
	o.DncListIds = dncListIds
}

// WithEdgeGroupID adds the edgeGroupID to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithEdgeGroupID(edgeGroupID *string) *GetOutboundCampaignsParams {
	o.SetEdgeGroupID(edgeGroupID)
	return o
}

// SetEdgeGroupID adds the edgeGroupId to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetEdgeGroupID(edgeGroupID *string) {
	o.EdgeGroupID = edgeGroupID
}

// WithFilterType adds the filterType to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithFilterType(filterType *string) *GetOutboundCampaignsParams {
	o.SetFilterType(filterType)
	return o
}

// SetFilterType adds the filterType to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetFilterType(filterType *string) {
	o.FilterType = filterType
}

// WithID adds the id to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithID(id []string) *GetOutboundCampaignsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetID(id []string) {
	o.ID = id
}

// WithName adds the name to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithName(name *string) *GetOutboundCampaignsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithPageNumber(pageNumber *int32) *GetOutboundCampaignsParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithPageSize(pageSize *int32) *GetOutboundCampaignsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithSortBy adds the sortBy to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithSortBy(sortBy *string) *GetOutboundCampaignsParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) WithSortOrder(sortOrder *string) *GetOutboundCampaignsParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get outbound campaigns params
func (o *GetOutboundCampaignsParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *GetOutboundCampaignsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CallAnalysisResponseSetID != nil {

		// query param callAnalysisResponseSetId
		var qrCallAnalysisResponseSetID string
		if o.CallAnalysisResponseSetID != nil {
			qrCallAnalysisResponseSetID = *o.CallAnalysisResponseSetID
		}
		qCallAnalysisResponseSetID := qrCallAnalysisResponseSetID
		if qCallAnalysisResponseSetID != "" {
			if err := r.SetQueryParam("callAnalysisResponseSetId", qCallAnalysisResponseSetID); err != nil {
				return err
			}
		}

	}

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

	if o.DistributionQueueID != nil {

		// query param distributionQueueId
		var qrDistributionQueueID string
		if o.DistributionQueueID != nil {
			qrDistributionQueueID = *o.DistributionQueueID
		}
		qDistributionQueueID := qrDistributionQueueID
		if qDistributionQueueID != "" {
			if err := r.SetQueryParam("distributionQueueId", qDistributionQueueID); err != nil {
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

	if o.DncListIds != nil {

		// query param dncListIds
		var qrDncListIds string
		if o.DncListIds != nil {
			qrDncListIds = *o.DncListIds
		}
		qDncListIds := qrDncListIds
		if qDncListIds != "" {
			if err := r.SetQueryParam("dncListIds", qDncListIds); err != nil {
				return err
			}
		}

	}

	if o.EdgeGroupID != nil {

		// query param edgeGroupId
		var qrEdgeGroupID string
		if o.EdgeGroupID != nil {
			qrEdgeGroupID = *o.EdgeGroupID
		}
		qEdgeGroupID := qrEdgeGroupID
		if qEdgeGroupID != "" {
			if err := r.SetQueryParam("edgeGroupId", qEdgeGroupID); err != nil {
				return err
			}
		}

	}

	if o.FilterType != nil {

		// query param filterType
		var qrFilterType string
		if o.FilterType != nil {
			qrFilterType = *o.FilterType
		}
		qFilterType := qrFilterType
		if qFilterType != "" {
			if err := r.SetQueryParam("filterType", qFilterType); err != nil {
				return err
			}
		}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}