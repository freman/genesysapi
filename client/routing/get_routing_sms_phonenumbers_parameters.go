// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewGetRoutingSmsPhonenumbersParams creates a new GetRoutingSmsPhonenumbersParams object
// with the default values initialized.
func NewGetRoutingSmsPhonenumbersParams() *GetRoutingSmsPhonenumbersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRoutingSmsPhonenumbersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoutingSmsPhonenumbersParamsWithTimeout creates a new GetRoutingSmsPhonenumbersParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoutingSmsPhonenumbersParamsWithTimeout(timeout time.Duration) *GetRoutingSmsPhonenumbersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRoutingSmsPhonenumbersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetRoutingSmsPhonenumbersParamsWithContext creates a new GetRoutingSmsPhonenumbersParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoutingSmsPhonenumbersParamsWithContext(ctx context.Context) *GetRoutingSmsPhonenumbersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRoutingSmsPhonenumbersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetRoutingSmsPhonenumbersParamsWithHTTPClient creates a new GetRoutingSmsPhonenumbersParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoutingSmsPhonenumbersParamsWithHTTPClient(client *http.Client) *GetRoutingSmsPhonenumbersParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetRoutingSmsPhonenumbersParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetRoutingSmsPhonenumbersParams contains all the parameters to send to the API endpoint
for the get routing sms phonenumbers operation typically these are written to a http.Request
*/
type GetRoutingSmsPhonenumbersParams struct {

	/*PageNumber
	  Page number

	*/
	PageNumber *int32
	/*PageSize
	  Page size

	*/
	PageSize *int32
	/*PhoneNumber
	  Filter on phone number address. Allowable characters are the digits '0-9' and the wild card character '\*'. If just digits are present, a contains search is done on the address pattern. For example, '317' could be matched anywhere in the address. An '\*' will match multiple digits. For example, to match a specific area code within the US a pattern like '1317*' could be used.

	*/
	PhoneNumber *string
	/*PhoneNumberStatus
	  Filter on phone number status

	*/
	PhoneNumberStatus *string
	/*PhoneNumberType
	  Filter on phone number type

	*/
	PhoneNumberType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithTimeout(timeout time.Duration) *GetRoutingSmsPhonenumbersParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithContext(ctx context.Context) *GetRoutingSmsPhonenumbersParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithHTTPClient(client *http.Client) *GetRoutingSmsPhonenumbersParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithPageNumber(pageNumber *int32) *GetRoutingSmsPhonenumbersParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithPageSize(pageSize *int32) *GetRoutingSmsPhonenumbersParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithPhoneNumber adds the phoneNumber to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithPhoneNumber(phoneNumber *string) *GetRoutingSmsPhonenumbersParams {
	o.SetPhoneNumber(phoneNumber)
	return o
}

// SetPhoneNumber adds the phoneNumber to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetPhoneNumber(phoneNumber *string) {
	o.PhoneNumber = phoneNumber
}

// WithPhoneNumberStatus adds the phoneNumberStatus to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithPhoneNumberStatus(phoneNumberStatus *string) *GetRoutingSmsPhonenumbersParams {
	o.SetPhoneNumberStatus(phoneNumberStatus)
	return o
}

// SetPhoneNumberStatus adds the phoneNumberStatus to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetPhoneNumberStatus(phoneNumberStatus *string) {
	o.PhoneNumberStatus = phoneNumberStatus
}

// WithPhoneNumberType adds the phoneNumberType to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) WithPhoneNumberType(phoneNumberType *string) *GetRoutingSmsPhonenumbersParams {
	o.SetPhoneNumberType(phoneNumberType)
	return o
}

// SetPhoneNumberType adds the phoneNumberType to the get routing sms phonenumbers params
func (o *GetRoutingSmsPhonenumbersParams) SetPhoneNumberType(phoneNumberType *string) {
	o.PhoneNumberType = phoneNumberType
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoutingSmsPhonenumbersParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.PhoneNumber != nil {

		// query param phoneNumber
		var qrPhoneNumber string
		if o.PhoneNumber != nil {
			qrPhoneNumber = *o.PhoneNumber
		}
		qPhoneNumber := qrPhoneNumber
		if qPhoneNumber != "" {
			if err := r.SetQueryParam("phoneNumber", qPhoneNumber); err != nil {
				return err
			}
		}

	}

	if o.PhoneNumberStatus != nil {

		// query param phoneNumberStatus
		var qrPhoneNumberStatus string
		if o.PhoneNumberStatus != nil {
			qrPhoneNumberStatus = *o.PhoneNumberStatus
		}
		qPhoneNumberStatus := qrPhoneNumberStatus
		if qPhoneNumberStatus != "" {
			if err := r.SetQueryParam("phoneNumberStatus", qPhoneNumberStatus); err != nil {
				return err
			}
		}

	}

	if o.PhoneNumberType != nil {

		// query param phoneNumberType
		var qrPhoneNumberType string
		if o.PhoneNumberType != nil {
			qrPhoneNumberType = *o.PhoneNumberType
		}
		qPhoneNumberType := qrPhoneNumberType
		if qPhoneNumberType != "" {
			if err := r.SetQueryParam("phoneNumberType", qPhoneNumberType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
