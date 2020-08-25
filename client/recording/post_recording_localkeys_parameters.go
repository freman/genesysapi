// Code generated by go-swagger; DO NOT EDIT.

package recording

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

// NewPostRecordingLocalkeysParams creates a new PostRecordingLocalkeysParams object
// with the default values initialized.
func NewPostRecordingLocalkeysParams() *PostRecordingLocalkeysParams {
	var ()
	return &PostRecordingLocalkeysParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRecordingLocalkeysParamsWithTimeout creates a new PostRecordingLocalkeysParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRecordingLocalkeysParamsWithTimeout(timeout time.Duration) *PostRecordingLocalkeysParams {
	var ()
	return &PostRecordingLocalkeysParams{

		timeout: timeout,
	}
}

// NewPostRecordingLocalkeysParamsWithContext creates a new PostRecordingLocalkeysParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRecordingLocalkeysParamsWithContext(ctx context.Context) *PostRecordingLocalkeysParams {
	var ()
	return &PostRecordingLocalkeysParams{

		Context: ctx,
	}
}

// NewPostRecordingLocalkeysParamsWithHTTPClient creates a new PostRecordingLocalkeysParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostRecordingLocalkeysParamsWithHTTPClient(client *http.Client) *PostRecordingLocalkeysParams {
	var ()
	return &PostRecordingLocalkeysParams{
		HTTPClient: client,
	}
}

/*PostRecordingLocalkeysParams contains all the parameters to send to the API endpoint
for the post recording localkeys operation typically these are written to a http.Request
*/
type PostRecordingLocalkeysParams struct {

	/*Body
	  Local Encryption body

	*/
	Body *models.LocalEncryptionKeyRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) WithTimeout(timeout time.Duration) *PostRecordingLocalkeysParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) WithContext(ctx context.Context) *PostRecordingLocalkeysParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) WithHTTPClient(client *http.Client) *PostRecordingLocalkeysParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) WithBody(body *models.LocalEncryptionKeyRequest) *PostRecordingLocalkeysParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post recording localkeys params
func (o *PostRecordingLocalkeysParams) SetBody(body *models.LocalEncryptionKeyRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRecordingLocalkeysParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}