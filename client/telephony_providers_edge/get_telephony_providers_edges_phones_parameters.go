// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

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

// NewGetTelephonyProvidersEdgesPhonesParams creates a new GetTelephonyProvidersEdgesPhonesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetTelephonyProvidersEdgesPhonesParams() *GetTelephonyProvidersEdgesPhonesParams {
	return &GetTelephonyProvidersEdgesPhonesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonesParamsWithTimeout creates a new GetTelephonyProvidersEdgesPhonesParams object
// with the ability to set a timeout on a request.
func NewGetTelephonyProvidersEdgesPhonesParamsWithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonesParams {
	return &GetTelephonyProvidersEdgesPhonesParams{
		timeout: timeout,
	}
}

// NewGetTelephonyProvidersEdgesPhonesParamsWithContext creates a new GetTelephonyProvidersEdgesPhonesParams object
// with the ability to set a context for a request.
func NewGetTelephonyProvidersEdgesPhonesParamsWithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonesParams {
	return &GetTelephonyProvidersEdgesPhonesParams{
		Context: ctx,
	}
}

// NewGetTelephonyProvidersEdgesPhonesParamsWithHTTPClient creates a new GetTelephonyProvidersEdgesPhonesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetTelephonyProvidersEdgesPhonesParamsWithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonesParams {
	return &GetTelephonyProvidersEdgesPhonesParams{
		HTTPClient: client,
	}
}

/*
GetTelephonyProvidersEdgesPhonesParams contains all the parameters to send to the API endpoint

	for the get telephony providers edges phones operation.

	Typically these are written to a http.Request.
*/
type GetTelephonyProvidersEdgesPhonesParams struct {

	/* Expand.

	   Fields to expand in the response, comma-separated
	*/
	Expand []string

	/* Fields.

	   Fields and properties to get, comma-separated
	*/
	Fields []string

	/* LinesDefaultForUserID.

	   Filter by lines.defaultForUser.id
	*/
	LinesDefaultForUserID *string

	/* LinesID.

	   Filter by lines.id
	*/
	LinesID *string

	/* LinesLoggedInUserID.

	   Filter by lines.loggedInUser.id
	*/
	LinesLoggedInUserID *string

	/* LinesName.

	   Filter by lines.name
	*/
	LinesName *string

	/* Name.

	   Name of the Phone to filter by
	*/
	Name *string

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

	/* PhoneBaseSettingsID.

	   Filter by phoneBaseSettings.id
	*/
	PhoneBaseSettingsID *string

	/* PhoneHardwareID.

	   Filter by phone_hardwareId
	*/
	PhoneHardwareID *string

	/* SecondaryStatusOperationalStatus.

	   The secondary status to filter by
	*/
	SecondaryStatusOperationalStatus *string

	/* SiteID.

	   Filter by site.id
	*/
	SiteID *string

	/* SortBy.

	   The field to sort by

	   Default: "name"
	*/
	SortBy *string

	/* SortOrder.

	   Sort order

	   Default: "ASC"
	*/
	SortOrder *string

	/* StatusOperationalStatus.

	   The primary status to filter by
	*/
	StatusOperationalStatus *string

	/* WebRtcUserID.

	   Filter by webRtcUser.id
	*/
	WebRtcUserID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get telephony providers edges phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesPhonesParams) WithDefaults() *GetTelephonyProvidersEdgesPhonesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get telephony providers edges phones params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetTelephonyProvidersEdgesPhonesParams) SetDefaults() {
	var (
		pageNumberDefault = int32(1)

		pageSizeDefault = int32(25)

		sortByDefault = string("name")

		sortOrderDefault = string("ASC")
	)

	val := GetTelephonyProvidersEdgesPhonesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		SortBy:     &sortByDefault,
		SortOrder:  &sortOrderDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithTimeout(timeout time.Duration) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithContext(ctx context.Context) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithHTTPClient(client *http.Client) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExpand adds the expand to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithExpand(expand []string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetExpand(expand)
	return o
}

// SetExpand adds the expand to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetExpand(expand []string) {
	o.Expand = expand
}

// WithFields adds the fields to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithFields(fields []string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLinesDefaultForUserID adds the linesDefaultForUserID to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithLinesDefaultForUserID(linesDefaultForUserID *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetLinesDefaultForUserID(linesDefaultForUserID)
	return o
}

// SetLinesDefaultForUserID adds the linesDefaultForUserId to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetLinesDefaultForUserID(linesDefaultForUserID *string) {
	o.LinesDefaultForUserID = linesDefaultForUserID
}

// WithLinesID adds the linesID to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithLinesID(linesID *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetLinesID(linesID)
	return o
}

// SetLinesID adds the linesId to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetLinesID(linesID *string) {
	o.LinesID = linesID
}

// WithLinesLoggedInUserID adds the linesLoggedInUserID to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithLinesLoggedInUserID(linesLoggedInUserID *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetLinesLoggedInUserID(linesLoggedInUserID)
	return o
}

// SetLinesLoggedInUserID adds the linesLoggedInUserId to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetLinesLoggedInUserID(linesLoggedInUserID *string) {
	o.LinesLoggedInUserID = linesLoggedInUserID
}

// WithLinesName adds the linesName to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithLinesName(linesName *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetLinesName(linesName)
	return o
}

// SetLinesName adds the linesName to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetLinesName(linesName *string) {
	o.LinesName = linesName
}

// WithName adds the name to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithName(name *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetName(name *string) {
	o.Name = name
}

// WithPageNumber adds the pageNumber to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithPageNumber(pageNumber *int32) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithPageSize(pageSize *int32) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPhoneBaseSettingsID adds the phoneBaseSettingsID to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithPhoneBaseSettingsID(phoneBaseSettingsID *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetPhoneBaseSettingsID(phoneBaseSettingsID)
	return o
}

// SetPhoneBaseSettingsID adds the phoneBaseSettingsId to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetPhoneBaseSettingsID(phoneBaseSettingsID *string) {
	o.PhoneBaseSettingsID = phoneBaseSettingsID
}

// WithPhoneHardwareID adds the phoneHardwareID to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithPhoneHardwareID(phoneHardwareID *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetPhoneHardwareID(phoneHardwareID)
	return o
}

// SetPhoneHardwareID adds the phoneHardwareId to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetPhoneHardwareID(phoneHardwareID *string) {
	o.PhoneHardwareID = phoneHardwareID
}

// WithSecondaryStatusOperationalStatus adds the secondaryStatusOperationalStatus to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithSecondaryStatusOperationalStatus(secondaryStatusOperationalStatus *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetSecondaryStatusOperationalStatus(secondaryStatusOperationalStatus)
	return o
}

// SetSecondaryStatusOperationalStatus adds the secondaryStatusOperationalStatus to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetSecondaryStatusOperationalStatus(secondaryStatusOperationalStatus *string) {
	o.SecondaryStatusOperationalStatus = secondaryStatusOperationalStatus
}

// WithSiteID adds the siteID to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithSiteID(siteID *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetSiteID(siteID *string) {
	o.SiteID = siteID
}

// WithSortBy adds the sortBy to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithSortBy(sortBy *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithSortOrder(sortOrder *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WithStatusOperationalStatus adds the statusOperationalStatus to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithStatusOperationalStatus(statusOperationalStatus *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetStatusOperationalStatus(statusOperationalStatus)
	return o
}

// SetStatusOperationalStatus adds the statusOperationalStatus to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetStatusOperationalStatus(statusOperationalStatus *string) {
	o.StatusOperationalStatus = statusOperationalStatus
}

// WithWebRtcUserID adds the webRtcUserID to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) WithWebRtcUserID(webRtcUserID *string) *GetTelephonyProvidersEdgesPhonesParams {
	o.SetWebRtcUserID(webRtcUserID)
	return o
}

// SetWebRtcUserID adds the webRtcUserId to the get telephony providers edges phones params
func (o *GetTelephonyProvidersEdgesPhonesParams) SetWebRtcUserID(webRtcUserID *string) {
	o.WebRtcUserID = webRtcUserID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTelephonyProvidersEdgesPhonesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Expand != nil {

		// binding items for expand
		joinedExpand := o.bindParamExpand(reg)

		// query array param expand
		if err := r.SetQueryParam("expand", joinedExpand...); err != nil {
			return err
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.LinesDefaultForUserID != nil {

		// query param lines.defaultForUser.id
		var qrLinesDefaultForUserID string

		if o.LinesDefaultForUserID != nil {
			qrLinesDefaultForUserID = *o.LinesDefaultForUserID
		}
		qLinesDefaultForUserID := qrLinesDefaultForUserID
		if qLinesDefaultForUserID != "" {

			if err := r.SetQueryParam("lines.defaultForUser.id", qLinesDefaultForUserID); err != nil {
				return err
			}
		}
	}

	if o.LinesID != nil {

		// query param lines.id
		var qrLinesID string

		if o.LinesID != nil {
			qrLinesID = *o.LinesID
		}
		qLinesID := qrLinesID
		if qLinesID != "" {

			if err := r.SetQueryParam("lines.id", qLinesID); err != nil {
				return err
			}
		}
	}

	if o.LinesLoggedInUserID != nil {

		// query param lines.loggedInUser.id
		var qrLinesLoggedInUserID string

		if o.LinesLoggedInUserID != nil {
			qrLinesLoggedInUserID = *o.LinesLoggedInUserID
		}
		qLinesLoggedInUserID := qrLinesLoggedInUserID
		if qLinesLoggedInUserID != "" {

			if err := r.SetQueryParam("lines.loggedInUser.id", qLinesLoggedInUserID); err != nil {
				return err
			}
		}
	}

	if o.LinesName != nil {

		// query param lines.name
		var qrLinesName string

		if o.LinesName != nil {
			qrLinesName = *o.LinesName
		}
		qLinesName := qrLinesName
		if qLinesName != "" {

			if err := r.SetQueryParam("lines.name", qLinesName); err != nil {
				return err
			}
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

	if o.PhoneBaseSettingsID != nil {

		// query param phoneBaseSettings.id
		var qrPhoneBaseSettingsID string

		if o.PhoneBaseSettingsID != nil {
			qrPhoneBaseSettingsID = *o.PhoneBaseSettingsID
		}
		qPhoneBaseSettingsID := qrPhoneBaseSettingsID
		if qPhoneBaseSettingsID != "" {

			if err := r.SetQueryParam("phoneBaseSettings.id", qPhoneBaseSettingsID); err != nil {
				return err
			}
		}
	}

	if o.PhoneHardwareID != nil {

		// query param phone_hardwareId
		var qrPhoneHardwareID string

		if o.PhoneHardwareID != nil {
			qrPhoneHardwareID = *o.PhoneHardwareID
		}
		qPhoneHardwareID := qrPhoneHardwareID
		if qPhoneHardwareID != "" {

			if err := r.SetQueryParam("phone_hardwareId", qPhoneHardwareID); err != nil {
				return err
			}
		}
	}

	if o.SecondaryStatusOperationalStatus != nil {

		// query param secondaryStatus.operationalStatus
		var qrSecondaryStatusOperationalStatus string

		if o.SecondaryStatusOperationalStatus != nil {
			qrSecondaryStatusOperationalStatus = *o.SecondaryStatusOperationalStatus
		}
		qSecondaryStatusOperationalStatus := qrSecondaryStatusOperationalStatus
		if qSecondaryStatusOperationalStatus != "" {

			if err := r.SetQueryParam("secondaryStatus.operationalStatus", qSecondaryStatusOperationalStatus); err != nil {
				return err
			}
		}
	}

	if o.SiteID != nil {

		// query param site.id
		var qrSiteID string

		if o.SiteID != nil {
			qrSiteID = *o.SiteID
		}
		qSiteID := qrSiteID
		if qSiteID != "" {

			if err := r.SetQueryParam("site.id", qSiteID); err != nil {
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

	if o.StatusOperationalStatus != nil {

		// query param status.operationalStatus
		var qrStatusOperationalStatus string

		if o.StatusOperationalStatus != nil {
			qrStatusOperationalStatus = *o.StatusOperationalStatus
		}
		qStatusOperationalStatus := qrStatusOperationalStatus
		if qStatusOperationalStatus != "" {

			if err := r.SetQueryParam("status.operationalStatus", qStatusOperationalStatus); err != nil {
				return err
			}
		}
	}

	if o.WebRtcUserID != nil {

		// query param webRtcUser.id
		var qrWebRtcUserID string

		if o.WebRtcUserID != nil {
			qrWebRtcUserID = *o.WebRtcUserID
		}
		qWebRtcUserID := qrWebRtcUserID
		if qWebRtcUserID != "" {

			if err := r.SetQueryParam("webRtcUser.id", qWebRtcUserID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamGetTelephonyProvidersEdgesPhones binds the parameter expand
func (o *GetTelephonyProvidersEdgesPhonesParams) bindParamExpand(formats strfmt.Registry) []string {
	expandIR := o.Expand

	var expandIC []string
	for _, expandIIR := range expandIR { // explode []string

		expandIIV := expandIIR // string as string
		expandIC = append(expandIC, expandIIV)
	}

	// items.CollectionFormat: "multi"
	expandIS := swag.JoinByFormat(expandIC, "multi")

	return expandIS
}

// bindParamGetTelephonyProvidersEdgesPhones binds the parameter fields
func (o *GetTelephonyProvidersEdgesPhonesParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "multi"
	fieldsIS := swag.JoinByFormat(fieldsIC, "multi")

	return fieldsIS
}
