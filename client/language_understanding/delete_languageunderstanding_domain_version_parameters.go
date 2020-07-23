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

// NewDeleteLanguageunderstandingDomainVersionParams creates a new DeleteLanguageunderstandingDomainVersionParams object
// with the default values initialized.
func NewDeleteLanguageunderstandingDomainVersionParams() *DeleteLanguageunderstandingDomainVersionParams {
	var ()
	return &DeleteLanguageunderstandingDomainVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLanguageunderstandingDomainVersionParamsWithTimeout creates a new DeleteLanguageunderstandingDomainVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLanguageunderstandingDomainVersionParamsWithTimeout(timeout time.Duration) *DeleteLanguageunderstandingDomainVersionParams {
	var ()
	return &DeleteLanguageunderstandingDomainVersionParams{

		timeout: timeout,
	}
}

// NewDeleteLanguageunderstandingDomainVersionParamsWithContext creates a new DeleteLanguageunderstandingDomainVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLanguageunderstandingDomainVersionParamsWithContext(ctx context.Context) *DeleteLanguageunderstandingDomainVersionParams {
	var ()
	return &DeleteLanguageunderstandingDomainVersionParams{

		Context: ctx,
	}
}

// NewDeleteLanguageunderstandingDomainVersionParamsWithHTTPClient creates a new DeleteLanguageunderstandingDomainVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLanguageunderstandingDomainVersionParamsWithHTTPClient(client *http.Client) *DeleteLanguageunderstandingDomainVersionParams {
	var ()
	return &DeleteLanguageunderstandingDomainVersionParams{
		HTTPClient: client,
	}
}

/*DeleteLanguageunderstandingDomainVersionParams contains all the parameters to send to the API endpoint
for the delete languageunderstanding domain version operation typically these are written to a http.Request
*/
type DeleteLanguageunderstandingDomainVersionParams struct {

	/*DomainID
	  ID of the NLU domain.

	*/
	DomainID string
	/*DomainVersionID
	  ID of the NLU domain version.

	*/
	DomainVersionID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) WithTimeout(timeout time.Duration) *DeleteLanguageunderstandingDomainVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) WithContext(ctx context.Context) *DeleteLanguageunderstandingDomainVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) WithHTTPClient(client *http.Client) *DeleteLanguageunderstandingDomainVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDomainID adds the domainID to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) WithDomainID(domainID string) *DeleteLanguageunderstandingDomainVersionParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WithDomainVersionID adds the domainVersionID to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) WithDomainVersionID(domainVersionID string) *DeleteLanguageunderstandingDomainVersionParams {
	o.SetDomainVersionID(domainVersionID)
	return o
}

// SetDomainVersionID adds the domainVersionId to the delete languageunderstanding domain version params
func (o *DeleteLanguageunderstandingDomainVersionParams) SetDomainVersionID(domainVersionID string) {
	o.DomainVersionID = domainVersionID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLanguageunderstandingDomainVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	// path param domainVersionId
	if err := r.SetPathParam("domainVersionId", o.DomainVersionID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
