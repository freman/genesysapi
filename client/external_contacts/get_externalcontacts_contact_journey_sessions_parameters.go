// Code generated by go-swagger; DO NOT EDIT.

package external_contacts

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

// NewGetExternalcontactsContactJourneySessionsParams creates a new GetExternalcontactsContactJourneySessionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetExternalcontactsContactJourneySessionsParams() *GetExternalcontactsContactJourneySessionsParams {
	return &GetExternalcontactsContactJourneySessionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetExternalcontactsContactJourneySessionsParamsWithTimeout creates a new GetExternalcontactsContactJourneySessionsParams object
// with the ability to set a timeout on a request.
func NewGetExternalcontactsContactJourneySessionsParamsWithTimeout(timeout time.Duration) *GetExternalcontactsContactJourneySessionsParams {
	return &GetExternalcontactsContactJourneySessionsParams{
		timeout: timeout,
	}
}

// NewGetExternalcontactsContactJourneySessionsParamsWithContext creates a new GetExternalcontactsContactJourneySessionsParams object
// with the ability to set a context for a request.
func NewGetExternalcontactsContactJourneySessionsParamsWithContext(ctx context.Context) *GetExternalcontactsContactJourneySessionsParams {
	return &GetExternalcontactsContactJourneySessionsParams{
		Context: ctx,
	}
}

// NewGetExternalcontactsContactJourneySessionsParamsWithHTTPClient creates a new GetExternalcontactsContactJourneySessionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetExternalcontactsContactJourneySessionsParamsWithHTTPClient(client *http.Client) *GetExternalcontactsContactJourneySessionsParams {
	return &GetExternalcontactsContactJourneySessionsParams{
		HTTPClient: client,
	}
}

/*
GetExternalcontactsContactJourneySessionsParams contains all the parameters to send to the API endpoint

	for the get externalcontacts contact journey sessions operation.

	Typically these are written to a http.Request.
*/
type GetExternalcontactsContactJourneySessionsParams struct {

	/* After.

	   The cursor that points to the end of the set of entities that has been returned.
	*/
	After *string

	/* ContactID.

	   ExternalContact ID
	*/
	ContactID string

	/* IncludeMerged.

	   Indicates whether to return sessions from all external contacts in the merge-set of the given one.
	*/
	IncludeMerged *bool

	/* PageSize.

	   Number of entities to return. Maximum of 200.
	*/
	PageSize *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get externalcontacts contact journey sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsContactJourneySessionsParams) WithDefaults() *GetExternalcontactsContactJourneySessionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get externalcontacts contact journey sessions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetExternalcontactsContactJourneySessionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) WithTimeout(timeout time.Duration) *GetExternalcontactsContactJourneySessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) WithContext(ctx context.Context) *GetExternalcontactsContactJourneySessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) WithHTTPClient(client *http.Client) *GetExternalcontactsContactJourneySessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) WithAfter(after *string) *GetExternalcontactsContactJourneySessionsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) SetAfter(after *string) {
	o.After = after
}

// WithContactID adds the contactID to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) WithContactID(contactID string) *GetExternalcontactsContactJourneySessionsParams {
	o.SetContactID(contactID)
	return o
}

// SetContactID adds the contactId to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) SetContactID(contactID string) {
	o.ContactID = contactID
}

// WithIncludeMerged adds the includeMerged to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) WithIncludeMerged(includeMerged *bool) *GetExternalcontactsContactJourneySessionsParams {
	o.SetIncludeMerged(includeMerged)
	return o
}

// SetIncludeMerged adds the includeMerged to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) SetIncludeMerged(includeMerged *bool) {
	o.IncludeMerged = includeMerged
}

// WithPageSize adds the pageSize to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) WithPageSize(pageSize *string) *GetExternalcontactsContactJourneySessionsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get externalcontacts contact journey sessions params
func (o *GetExternalcontactsContactJourneySessionsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetExternalcontactsContactJourneySessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.After != nil {

		// query param after
		var qrAfter string

		if o.After != nil {
			qrAfter = *o.After
		}
		qAfter := qrAfter
		if qAfter != "" {

			if err := r.SetQueryParam("after", qAfter); err != nil {
				return err
			}
		}
	}

	// path param contactId
	if err := r.SetPathParam("contactId", o.ContactID); err != nil {
		return err
	}

	if o.IncludeMerged != nil {

		// query param includeMerged
		var qrIncludeMerged bool

		if o.IncludeMerged != nil {
			qrIncludeMerged = *o.IncludeMerged
		}
		qIncludeMerged := swag.FormatBool(qrIncludeMerged)
		if qIncludeMerged != "" {

			if err := r.SetQueryParam("includeMerged", qIncludeMerged); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize string

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := qrPageSize
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