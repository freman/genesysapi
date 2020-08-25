// Code generated by go-swagger; DO NOT EDIT.

package knowledge

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
)

// NewGetKnowledgeKnowledgebaseLanguageTrainingsParams creates a new GetKnowledgeKnowledgebaseLanguageTrainingsParams object
// with the default values initialized.
func NewGetKnowledgeKnowledgebaseLanguageTrainingsParams() *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageTrainingsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsParamsWithTimeout creates a new GetKnowledgeKnowledgebaseLanguageTrainingsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeKnowledgebaseLanguageTrainingsParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageTrainingsParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsParamsWithContext creates a new GetKnowledgeKnowledgebaseLanguageTrainingsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeKnowledgebaseLanguageTrainingsParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageTrainingsParams{

		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseLanguageTrainingsParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseLanguageTrainingsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeKnowledgebaseLanguageTrainingsParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	var ()
	return &GetKnowledgeKnowledgebaseLanguageTrainingsParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeKnowledgebaseLanguageTrainingsParams contains all the parameters to send to the API endpoint
for the get knowledge knowledgebase language trainings operation typically these are written to a http.Request
*/
type GetKnowledgeKnowledgebaseLanguageTrainingsParams struct {

	/*After
	  The cursor that points to the end of the set of entities that has been returned.

	*/
	After *string
	/*Before
	  The cursor that points to the start of the set of entities that has been returned.

	*/
	Before *string
	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string
	/*LanguageCode
	  Language code, format: iso2-LOCALE

	*/
	LanguageCode string
	/*Limit
	  Number of entities to return. Maximum of 200.

	*/
	Limit *string
	/*PageSize
	  Number of entities to return. Maximum of 200.

	*/
	PageSize *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAfter adds the after to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithAfter(after *string) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetAfter(after)
	return o
}

// SetAfter adds the after to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetAfter(after *string) {
	o.After = after
}

// WithBefore adds the before to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithBefore(before *string) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetBefore(before)
	return o
}

// SetBefore adds the before to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetBefore(before *string) {
	o.Before = before
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WithLanguageCode adds the languageCode to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithLanguageCode(languageCode string) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetLanguageCode(languageCode)
	return o
}

// SetLanguageCode adds the languageCode to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetLanguageCode(languageCode string) {
	o.LanguageCode = languageCode
}

// WithLimit adds the limit to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithLimit(limit *string) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetLimit(limit *string) {
	o.Limit = limit
}

// WithPageSize adds the pageSize to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WithPageSize(pageSize *string) *GetKnowledgeKnowledgebaseLanguageTrainingsParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the get knowledge knowledgebase language trainings params
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) SetPageSize(pageSize *string) {
	o.PageSize = pageSize
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseLanguageTrainingsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Before != nil {

		// query param before
		var qrBefore string
		if o.Before != nil {
			qrBefore = *o.Before
		}
		qBefore := qrBefore
		if qBefore != "" {
			if err := r.SetQueryParam("before", qBefore); err != nil {
				return err
			}
		}

	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	// path param languageCode
	if err := r.SetPathParam("languageCode", o.LanguageCode); err != nil {
		return err
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit string
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := qrLimit
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
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