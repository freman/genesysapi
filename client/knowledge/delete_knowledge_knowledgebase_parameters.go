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

// NewDeleteKnowledgeKnowledgebaseParams creates a new DeleteKnowledgeKnowledgebaseParams object
// with the default values initialized.
func NewDeleteKnowledgeKnowledgebaseParams() *DeleteKnowledgeKnowledgebaseParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseParamsWithTimeout creates a new DeleteKnowledgeKnowledgebaseParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteKnowledgeKnowledgebaseParamsWithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseParams{

		timeout: timeout,
	}
}

// NewDeleteKnowledgeKnowledgebaseParamsWithContext creates a new DeleteKnowledgeKnowledgebaseParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteKnowledgeKnowledgebaseParamsWithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseParams{

		Context: ctx,
	}
}

// NewDeleteKnowledgeKnowledgebaseParamsWithHTTPClient creates a new DeleteKnowledgeKnowledgebaseParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteKnowledgeKnowledgebaseParamsWithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseParams {
	var ()
	return &DeleteKnowledgeKnowledgebaseParams{
		HTTPClient: client,
	}
}

/*DeleteKnowledgeKnowledgebaseParams contains all the parameters to send to the API endpoint
for the delete knowledge knowledgebase operation typically these are written to a http.Request
*/
type DeleteKnowledgeKnowledgebaseParams struct {

	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) WithTimeout(timeout time.Duration) *DeleteKnowledgeKnowledgebaseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) WithContext(ctx context.Context) *DeleteKnowledgeKnowledgebaseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) WithHTTPClient(client *http.Client) *DeleteKnowledgeKnowledgebaseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) WithKnowledgeBaseID(knowledgeBaseID string) *DeleteKnowledgeKnowledgebaseParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the delete knowledge knowledgebase params
func (o *DeleteKnowledgeKnowledgebaseParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteKnowledgeKnowledgebaseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
