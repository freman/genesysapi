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

// NewDeleteLanguageunderstandingDomainParams creates a new DeleteLanguageunderstandingDomainParams object
// with the default values initialized.
func NewDeleteLanguageunderstandingDomainParams() *DeleteLanguageunderstandingDomainParams {
	var ()
	return &DeleteLanguageunderstandingDomainParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLanguageunderstandingDomainParamsWithTimeout creates a new DeleteLanguageunderstandingDomainParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLanguageunderstandingDomainParamsWithTimeout(timeout time.Duration) *DeleteLanguageunderstandingDomainParams {
	var ()
	return &DeleteLanguageunderstandingDomainParams{

		timeout: timeout,
	}
}

// NewDeleteLanguageunderstandingDomainParamsWithContext creates a new DeleteLanguageunderstandingDomainParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLanguageunderstandingDomainParamsWithContext(ctx context.Context) *DeleteLanguageunderstandingDomainParams {
	var ()
	return &DeleteLanguageunderstandingDomainParams{

		Context: ctx,
	}
}

// NewDeleteLanguageunderstandingDomainParamsWithHTTPClient creates a new DeleteLanguageunderstandingDomainParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLanguageunderstandingDomainParamsWithHTTPClient(client *http.Client) *DeleteLanguageunderstandingDomainParams {
	var ()
	return &DeleteLanguageunderstandingDomainParams{
		HTTPClient: client,
	}
}

/*DeleteLanguageunderstandingDomainParams contains all the parameters to send to the API endpoint
for the delete languageunderstanding domain operation typically these are written to a http.Request
*/
type DeleteLanguageunderstandingDomainParams struct {

	/*DomainID
	  ID of the NLU domain.

	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) WithTimeout(timeout time.Duration) *DeleteLanguageunderstandingDomainParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) WithContext(ctx context.Context) *DeleteLanguageunderstandingDomainParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) WithHTTPClient(client *http.Client) *DeleteLanguageunderstandingDomainParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) WithDomainID(domainID string) *DeleteLanguageunderstandingDomainParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete languageunderstanding domain params
func (o *DeleteLanguageunderstandingDomainParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLanguageunderstandingDomainParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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