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

	"github.com/freman/genesysapi/models"
)

// NewPutSpeechandtextanalyticsProgramParams creates a new PutSpeechandtextanalyticsProgramParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutSpeechandtextanalyticsProgramParams() *PutSpeechandtextanalyticsProgramParams {
	return &PutSpeechandtextanalyticsProgramParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutSpeechandtextanalyticsProgramParamsWithTimeout creates a new PutSpeechandtextanalyticsProgramParams object
// with the ability to set a timeout on a request.
func NewPutSpeechandtextanalyticsProgramParamsWithTimeout(timeout time.Duration) *PutSpeechandtextanalyticsProgramParams {
	return &PutSpeechandtextanalyticsProgramParams{
		timeout: timeout,
	}
}

// NewPutSpeechandtextanalyticsProgramParamsWithContext creates a new PutSpeechandtextanalyticsProgramParams object
// with the ability to set a context for a request.
func NewPutSpeechandtextanalyticsProgramParamsWithContext(ctx context.Context) *PutSpeechandtextanalyticsProgramParams {
	return &PutSpeechandtextanalyticsProgramParams{
		Context: ctx,
	}
}

// NewPutSpeechandtextanalyticsProgramParamsWithHTTPClient creates a new PutSpeechandtextanalyticsProgramParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutSpeechandtextanalyticsProgramParamsWithHTTPClient(client *http.Client) *PutSpeechandtextanalyticsProgramParams {
	return &PutSpeechandtextanalyticsProgramParams{
		HTTPClient: client,
	}
}

/*
PutSpeechandtextanalyticsProgramParams contains all the parameters to send to the API endpoint

	for the put speechandtextanalytics program operation.

	Typically these are written to a http.Request.
*/
type PutSpeechandtextanalyticsProgramParams struct {

	/* Body.

	   The program to update
	*/
	Body *models.ProgramRequest

	/* ProgramID.

	   The id of the program
	*/
	ProgramID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put speechandtextanalytics program params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSpeechandtextanalyticsProgramParams) WithDefaults() *PutSpeechandtextanalyticsProgramParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put speechandtextanalytics program params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutSpeechandtextanalyticsProgramParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) WithTimeout(timeout time.Duration) *PutSpeechandtextanalyticsProgramParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) WithContext(ctx context.Context) *PutSpeechandtextanalyticsProgramParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) WithHTTPClient(client *http.Client) *PutSpeechandtextanalyticsProgramParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) WithBody(body *models.ProgramRequest) *PutSpeechandtextanalyticsProgramParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) SetBody(body *models.ProgramRequest) {
	o.Body = body
}

// WithProgramID adds the programID to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) WithProgramID(programID string) *PutSpeechandtextanalyticsProgramParams {
	o.SetProgramID(programID)
	return o
}

// SetProgramID adds the programId to the put speechandtextanalytics program params
func (o *PutSpeechandtextanalyticsProgramParams) SetProgramID(programID string) {
	o.ProgramID = programID
}

// WriteToRequest writes these params to a swagger request
func (o *PutSpeechandtextanalyticsProgramParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param programId
	if err := r.SetPathParam("programId", o.ProgramID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
