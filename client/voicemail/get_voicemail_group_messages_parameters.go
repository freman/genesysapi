// Code generated by go-swagger; DO NOT EDIT.

package voicemail

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

// NewGetVoicemailGroupMessagesParams creates a new GetVoicemailGroupMessagesParams object
// with the default values initialized.
func NewGetVoicemailGroupMessagesParams() *GetVoicemailGroupMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetVoicemailGroupMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetVoicemailGroupMessagesParamsWithTimeout creates a new GetVoicemailGroupMessagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetVoicemailGroupMessagesParamsWithTimeout(timeout time.Duration) *GetVoicemailGroupMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetVoicemailGroupMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetVoicemailGroupMessagesParamsWithContext creates a new GetVoicemailGroupMessagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetVoicemailGroupMessagesParamsWithContext(ctx context.Context) *GetVoicemailGroupMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetVoicemailGroupMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetVoicemailGroupMessagesParamsWithHTTPClient creates a new GetVoicemailGroupMessagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetVoicemailGroupMessagesParamsWithHTTPClient(client *http.Client) *GetVoicemailGroupMessagesParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetVoicemailGroupMessagesParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetVoicemailGroupMessagesParams contains all the parameters to send to the API endpoint
for the get voicemail group messages operation typically these are written to a http.Request
*/
type GetVoicemailGroupMessagesParams struct {

	/*GroupID
	  Group ID

	*/
	GroupID string
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

// WithTimeout adds the timeout to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) WithTimeout(timeout time.Duration) *GetVoicemailGroupMessagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) WithContext(ctx context.Context) *GetVoicemailGroupMessagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) WithHTTPClient(client *http.Client) *GetVoicemailGroupMessagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGroupID adds the groupID to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) WithGroupID(groupID string) *GetVoicemailGroupMessagesParams {
	o.SetGroupID(groupID)
	return o
}

// SetGroupID adds the groupId to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) SetGroupID(groupID string) {
	o.GroupID = groupID
}

// WithPageNumber adds the pageNumber to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) WithPageNumber(pageNumber *int32) *GetVoicemailGroupMessagesParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) WithPageSize(pageSize *int32) *GetVoicemailGroupMessagesParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get voicemail group messages params
func (o *GetVoicemailGroupMessagesParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetVoicemailGroupMessagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param groupId
	if err := r.SetPathParam("groupId", o.GroupID); err != nil {
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
