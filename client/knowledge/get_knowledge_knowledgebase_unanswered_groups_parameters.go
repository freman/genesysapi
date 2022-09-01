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

// NewGetKnowledgeKnowledgebaseUnansweredGroupsParams creates a new GetKnowledgeKnowledgebaseUnansweredGroupsParams object
// with the default values initialized.
func NewGetKnowledgeKnowledgebaseUnansweredGroupsParams() *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	var ()
	return &GetKnowledgeKnowledgebaseUnansweredGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupsParamsWithTimeout creates a new GetKnowledgeKnowledgebaseUnansweredGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetKnowledgeKnowledgebaseUnansweredGroupsParamsWithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	var ()
	return &GetKnowledgeKnowledgebaseUnansweredGroupsParams{

		timeout: timeout,
	}
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupsParamsWithContext creates a new GetKnowledgeKnowledgebaseUnansweredGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetKnowledgeKnowledgebaseUnansweredGroupsParamsWithContext(ctx context.Context) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	var ()
	return &GetKnowledgeKnowledgebaseUnansweredGroupsParams{

		Context: ctx,
	}
}

// NewGetKnowledgeKnowledgebaseUnansweredGroupsParamsWithHTTPClient creates a new GetKnowledgeKnowledgebaseUnansweredGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetKnowledgeKnowledgebaseUnansweredGroupsParamsWithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	var ()
	return &GetKnowledgeKnowledgebaseUnansweredGroupsParams{
		HTTPClient: client,
	}
}

/*GetKnowledgeKnowledgebaseUnansweredGroupsParams contains all the parameters to send to the API endpoint
for the get knowledge knowledgebase unanswered groups operation typically these are written to a http.Request
*/
type GetKnowledgeKnowledgebaseUnansweredGroupsParams struct {

	/*App
	  The app value to be used for filtering phrases.

	*/
	App *string
	/*KnowledgeBaseID
	  Knowledge base ID

	*/
	KnowledgeBaseID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) WithTimeout(timeout time.Duration) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) WithContext(ctx context.Context) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) WithHTTPClient(client *http.Client) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApp adds the app to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) WithApp(app *string) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	o.SetApp(app)
	return o
}

// SetApp adds the app to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) SetApp(app *string) {
	o.App = app
}

// WithKnowledgeBaseID adds the knowledgeBaseID to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) WithKnowledgeBaseID(knowledgeBaseID string) *GetKnowledgeKnowledgebaseUnansweredGroupsParams {
	o.SetKnowledgeBaseID(knowledgeBaseID)
	return o
}

// SetKnowledgeBaseID adds the knowledgeBaseId to the get knowledge knowledgebase unanswered groups params
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) SetKnowledgeBaseID(knowledgeBaseID string) {
	o.KnowledgeBaseID = knowledgeBaseID
}

// WriteToRequest writes these params to a swagger request
func (o *GetKnowledgeKnowledgebaseUnansweredGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.App != nil {

		// query param app
		var qrApp string
		if o.App != nil {
			qrApp = *o.App
		}
		qApp := qrApp
		if qApp != "" {
			if err := r.SetQueryParam("app", qApp); err != nil {
				return err
			}
		}

	}

	// path param knowledgeBaseId
	if err := r.SetPathParam("knowledgeBaseId", o.KnowledgeBaseID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}