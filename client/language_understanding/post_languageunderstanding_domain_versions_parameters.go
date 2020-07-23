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

	"github.com/freman/genesysapi/models"
)

// NewPostLanguageunderstandingDomainVersionsParams creates a new PostLanguageunderstandingDomainVersionsParams object
// with the default values initialized.
func NewPostLanguageunderstandingDomainVersionsParams() *PostLanguageunderstandingDomainVersionsParams {
	var ()
	return &PostLanguageunderstandingDomainVersionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostLanguageunderstandingDomainVersionsParamsWithTimeout creates a new PostLanguageunderstandingDomainVersionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostLanguageunderstandingDomainVersionsParamsWithTimeout(timeout time.Duration) *PostLanguageunderstandingDomainVersionsParams {
	var ()
	return &PostLanguageunderstandingDomainVersionsParams{

		timeout: timeout,
	}
}

// NewPostLanguageunderstandingDomainVersionsParamsWithContext creates a new PostLanguageunderstandingDomainVersionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostLanguageunderstandingDomainVersionsParamsWithContext(ctx context.Context) *PostLanguageunderstandingDomainVersionsParams {
	var ()
	return &PostLanguageunderstandingDomainVersionsParams{

		Context: ctx,
	}
}

// NewPostLanguageunderstandingDomainVersionsParamsWithHTTPClient creates a new PostLanguageunderstandingDomainVersionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostLanguageunderstandingDomainVersionsParamsWithHTTPClient(client *http.Client) *PostLanguageunderstandingDomainVersionsParams {
	var ()
	return &PostLanguageunderstandingDomainVersionsParams{
		HTTPClient: client,
	}
}

/*PostLanguageunderstandingDomainVersionsParams contains all the parameters to send to the API endpoint
for the post languageunderstanding domain versions operation typically these are written to a http.Request
*/
type PostLanguageunderstandingDomainVersionsParams struct {

	/*Body
	  The NLU Domain Version to create.

	*/
	Body *models.NluDomainVersion
	/*DomainID
	  ID of the NLU domain.

	*/
	DomainID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) WithTimeout(timeout time.Duration) *PostLanguageunderstandingDomainVersionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) WithContext(ctx context.Context) *PostLanguageunderstandingDomainVersionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) WithHTTPClient(client *http.Client) *PostLanguageunderstandingDomainVersionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) WithBody(body *models.NluDomainVersion) *PostLanguageunderstandingDomainVersionsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) SetBody(body *models.NluDomainVersion) {
	o.Body = body
}

// WithDomainID adds the domainID to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) WithDomainID(domainID string) *PostLanguageunderstandingDomainVersionsParams {
	o.SetDomainID(domainID)
	return o
}

// SetDomainID adds the domainId to the post languageunderstanding domain versions params
func (o *PostLanguageunderstandingDomainVersionsParams) SetDomainID(domainID string) {
	o.DomainID = domainID
}

// WriteToRequest writes these params to a swagger request
func (o *PostLanguageunderstandingDomainVersionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param domainId
	if err := r.SetPathParam("domainId", o.DomainID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
