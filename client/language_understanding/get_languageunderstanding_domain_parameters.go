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
)

// NewGetLanguageunderstandingDomainParams creates a new GetLanguageunderstandingDomainParams object
// with the default values initialized.
func NewGetLanguageunderstandingDomainParams() *GetLanguageunderstandingDomainParams {
	var ()
	return &GetLanguageunderstandingDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetLanguageunderstandingDomainParamsWithTimeout creates a new GetLanguageunderstandingDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetLanguageunderstandingDomainParamsWithTimeout(timeout time.Duration) *GetLanguageunderstandingDomainParams {
	var ()
	return &GetLanguageunderstandingDomainParams{

		timeout: timeout,
	}
}

// NewGetLanguageunderstandingDomainParamsWithContext creates a new GetLanguageunderstandingDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetLanguageunderstandingDomainParamsWithContext(ctx context.Context) *GetLanguageunderstandingDomainParams {
	var ()
	return &GetLanguageunderstandingDomainParams{

		Context: ctx,
	}
}

// NewGetLanguageunderstandingDomainParamsWithHTTPClient creates a new GetLanguageunderstandingDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetLanguageunderstandingDomainParamsWithHTTPClient(client *http.Client) *GetLanguageunderstandingDomainParams {
	var ()
	return &GetLanguageunderstandingDomainParams{
		HTTPClient: client,
	}
}

/*GetLanguageunderstandingDomainParams contains all the parameters to send to the API endpoint
for the get languageunderstanding domain operation typically these are written to a http.Request
*/
type GetLanguageunderstandingDomainParams struct {

	/*DomainID
	  ID of the NLU domain.

	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) WithTimeout(timeout time.Duration) *GetLanguageunderstandingDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) WithContext(ctx context.Context) *GetLanguageunderstandingDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) WithHTTPClient(client *http.Client) *GetLanguageunderstandingDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) WithDomainID(domainID string) *GetLanguageunderstandingDomainParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the get languageunderstanding domain params
func (o *GetLanguageunderstandingDomainParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *GetLanguageunderstandingDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
