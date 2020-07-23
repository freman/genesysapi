// Code generated by go-swagger; DO NOT EDIT.

package language_understanding

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

// NewGetLanguageunderstandingDomainFeedbackParams creates a new GetLanguageunderstandingDomainFeedbackParams object
// with the default values initialized.
func NewGetLanguageunderstandingDomainFeedbackParams() *GetLanguageunderstandingDomainFeedbackParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLanguageunderstandingDomainFeedbackParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguageunderstandingDomainFeedbackParamsWithTimeout creates a new GetLanguageunderstandingDomainFeedbackParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLanguageunderstandingDomainFeedbackParamsWithTimeout(timeout time.Duration) *GetLanguageunderstandingDomainFeedbackParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLanguageunderstandingDomainFeedbackParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		timeout: timeout,
	}
}

// NewGetLanguageunderstandingDomainFeedbackParamsWithContext creates a new GetLanguageunderstandingDomainFeedbackParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLanguageunderstandingDomainFeedbackParamsWithContext(ctx context.Context) *GetLanguageunderstandingDomainFeedbackParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLanguageunderstandingDomainFeedbackParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,

		Context: ctx,
	}
}

// NewGetLanguageunderstandingDomainFeedbackParamsWithHTTPClient creates a new GetLanguageunderstandingDomainFeedbackParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLanguageunderstandingDomainFeedbackParamsWithHTTPClient(client *http.Client) *GetLanguageunderstandingDomainFeedbackParams {
	var (
		pageNumberDefault = int32(1)
		pageSizeDefault   = int32(25)
	)
	return &GetLanguageunderstandingDomainFeedbackParams{
		PageNumber: &pageNumberDefault,
		PageSize:   &pageSizeDefault,
		HTTPClient: client,
	}
}

/*GetLanguageunderstandingDomainFeedbackParams contains all the parameters to send to the API endpoint
for the get languageunderstanding domain feedback operation typically these are written to a http.Request
*/
type GetLanguageunderstandingDomainFeedbackParams struct {

	/*Assessment
	  The top assessment to retrieve feedback for.

	*/
	Assessment *string
	/*DateEnd
	  End of time window as ISO-8601 date.

	*/
	DateEnd *strfmt.Date
	/*DateStart
	  Begin of time window as ISO-8601 date.

	*/
	DateStart *strfmt.Date
	/*DomainID
	  ID of the NLU domain.

	*/
	DomainID string
	/*Fields
	  Fields and properties to get, comma-separated

	*/
	Fields []string
	/*IncludeDeleted
	  Whether to include soft-deleted items in the result.

	*/
	IncludeDeleted *bool
	/*IntentName
	  The top intent name to retrieve feedback for.

	*/
	IntentName *string
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

// WithTimeout adds the timeout to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithTimeout(timeout time.Duration) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithContext(ctx context.Context) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithHTTPClient(client *http.Client) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAssessment adds the assessment to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithAssessment(assessment *string) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetAssessment(assessment)
	return o
}

// SetAssessment adds the assessment to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetAssessment(assessment *string) {
	o.Assessment = assessment
}

// WithDateEnd adds the dateEnd to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithDateEnd(dateEnd *strfmt.Date) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetDateEnd(dateEnd)
	return o
}

// SetDateEnd adds the dateEnd to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetDateEnd(dateEnd *strfmt.Date) {
	o.DateEnd = dateEnd
}

// WithDateStart adds the dateStart to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithDateStart(dateStart *strfmt.Date) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetDateStart(dateStart)
	return o
}

// SetDateStart adds the dateStart to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetDateStart(dateStart *strfmt.Date) {
	o.DateStart = dateStart
}

// WithDomainID adds the domainID to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithDomainID(domainID string) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithFields adds the fields to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithFields(fields []string) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIncludeDeleted adds the includeDeleted to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithIncludeDeleted(includeDeleted *bool) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetIncludeDeleted(includeDeleted)
	return o
}

// SetIncludeDeleted adds the includeDeleted to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetIncludeDeleted(includeDeleted *bool) {
	o.IncludeDeleted = includeDeleted
}

// WithIntentName adds the intentName to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithIntentName(intentName *string) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetIntentName(intentName)
	return o
}

// SetIntentName adds the intentName to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetIntentName(intentName *string) {
	o.IntentName = intentName
}

// WithPageNumber adds the pageNumber to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithPageNumber(pageNumber *int32) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) WithPageSize(pageSize *int32) *GetLanguageunderstandingDomainFeedbackParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get languageunderstanding domain feedback params
func (o *GetLanguageunderstandingDomainFeedbackParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguageunderstandingDomainFeedbackParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Assessment != nil {

		// query param assessment
		var qrAssessment string
		if o.Assessment != nil {
			qrAssessment = *o.Assessment
		}
		qAssessment := qrAssessment
		if qAssessment != "" {
			if err := r.SetQueryParam("assessment", qAssessment); err != nil {
				return err
			}
		}

	}

	if o.DateEnd != nil {

		// query param dateEnd
		var qrDateEnd strfmt.Date
		if o.DateEnd != nil {
			qrDateEnd = *o.DateEnd
		}
		qDateEnd := qrDateEnd.String()
		if qDateEnd != "" {
			if err := r.SetQueryParam("dateEnd", qDateEnd); err != nil {
				return err
			}
		}

	}

	if o.DateStart != nil {

		// query param dateStart
		var qrDateStart strfmt.Date
		if o.DateStart != nil {
			qrDateStart = *o.DateStart
		}
		qDateStart := qrDateStart.String()
		if qDateStart != "" {
			if err := r.SetQueryParam("dateStart", qDateStart); err != nil {
				return err
			}
		}

	}

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	valuesFields := o.Fields

	joinedFields := swag.JoinByFormat(valuesFields, "multi")
	// query array param fields
	if err := r.SetQueryParam("fields", joinedFields...); err != nil {
		return err
	}

	if o.IncludeDeleted != nil {

		// query param includeDeleted
		var qrIncludeDeleted bool
		if o.IncludeDeleted != nil {
			qrIncludeDeleted = *o.IncludeDeleted
		}
		qIncludeDeleted := swag.FormatBool(qrIncludeDeleted)
		if qIncludeDeleted != "" {
			if err := r.SetQueryParam("includeDeleted", qIncludeDeleted); err != nil {
				return err
			}
		}

	}

	if o.IntentName != nil {

		// query param intentName
		var qrIntentName string
		if o.IntentName != nil {
			qrIntentName = *o.IntentName
		}
		qIntentName := qrIntentName
		if qIntentName != "" {
			if err := r.SetQueryParam("intentName", qIntentName); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
