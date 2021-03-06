// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

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

// NewGetSpeechandtextanalyticsProgramParams creates a new GetSpeechandtextanalyticsProgramParams object
// with the default values initialized.
func NewGetSpeechandtextanalyticsProgramParams() *GetSpeechandtextanalyticsProgramParams {
	var ()
	return &GetSpeechandtextanalyticsProgramParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSpeechandtextanalyticsProgramParamsWithTimeout creates a new GetSpeechandtextanalyticsProgramParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSpeechandtextanalyticsProgramParamsWithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramParams {
	var ()
	return &GetSpeechandtextanalyticsProgramParams{

		timeout: timeout,
	}
}

// NewGetSpeechandtextanalyticsProgramParamsWithContext creates a new GetSpeechandtextanalyticsProgramParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSpeechandtextanalyticsProgramParamsWithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramParams {
	var ()
	return &GetSpeechandtextanalyticsProgramParams{

		Context: ctx,
	}
}

// NewGetSpeechandtextanalyticsProgramParamsWithHTTPClient creates a new GetSpeechandtextanalyticsProgramParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSpeechandtextanalyticsProgramParamsWithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramParams {
	var ()
	return &GetSpeechandtextanalyticsProgramParams{
		HTTPClient: client,
	}
}

/*GetSpeechandtextanalyticsProgramParams contains all the parameters to send to the API endpoint
for the get speechandtextanalytics program operation typically these are written to a http.Request
*/
type GetSpeechandtextanalyticsProgramParams struct {

	/*ProgramID
	  The id of the program

	*/
	ProgramID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) WithTimeout(timeout time.Duration) *GetSpeechandtextanalyticsProgramParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) WithContext(ctx context.Context) *GetSpeechandtextanalyticsProgramParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) WithHTTPClient(client *http.Client) *GetSpeechandtextanalyticsProgramParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProgramID adds the programID to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) WithProgramID(programID string) *GetSpeechandtextanalyticsProgramParams {
	o.SetProgramID(programID)
	return o
}

// SetProgramID adds the programId to the get speechandtextanalytics program params
func (o *GetSpeechandtextanalyticsProgramParams) SetProgramID(programID string) {
	o.ProgramID = programID
}

// WriteToRequest writes these params to a swagger request
func (o *GetSpeechandtextanalyticsProgramParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param programId
	if err := r.SetPathParam("programId", o.ProgramID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
