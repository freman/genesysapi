// Code generated by go-swagger; DO NOT EDIT.

package routing

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

// NewDeleteUserRoutinglanguageParams creates a new DeleteUserRoutinglanguageParams object
// with the default values initialized.
func NewDeleteUserRoutinglanguageParams() *DeleteUserRoutinglanguageParams {
	var ()
	return &DeleteUserRoutinglanguageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteUserRoutinglanguageParamsWithTimeout creates a new DeleteUserRoutinglanguageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteUserRoutinglanguageParamsWithTimeout(timeout time.Duration) *DeleteUserRoutinglanguageParams {
	var ()
	return &DeleteUserRoutinglanguageParams{

		timeout: timeout,
	}
}

// NewDeleteUserRoutinglanguageParamsWithContext creates a new DeleteUserRoutinglanguageParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteUserRoutinglanguageParamsWithContext(ctx context.Context) *DeleteUserRoutinglanguageParams {
	var ()
	return &DeleteUserRoutinglanguageParams{

		Context: ctx,
	}
}

// NewDeleteUserRoutinglanguageParamsWithHTTPClient creates a new DeleteUserRoutinglanguageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteUserRoutinglanguageParamsWithHTTPClient(client *http.Client) *DeleteUserRoutinglanguageParams {
	var ()
	return &DeleteUserRoutinglanguageParams{
		HTTPClient: client,
	}
}

/*DeleteUserRoutinglanguageParams contains all the parameters to send to the API endpoint
for the delete user routinglanguage operation typically these are written to a http.Request
*/
type DeleteUserRoutinglanguageParams struct {

	/*LanguageID
	  languageId

	*/
	LanguageID string
	/*UserID
	  User ID

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) WithTimeout(timeout time.Duration) *DeleteUserRoutinglanguageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) WithContext(ctx context.Context) *DeleteUserRoutinglanguageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) WithHTTPClient(client *http.Client) *DeleteUserRoutinglanguageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLanguageID adds the languageID to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) WithLanguageID(languageID string) *DeleteUserRoutinglanguageParams {
	o.SetLanguageID(languageID)
	return o
}

// SetLanguageID adds the languageId to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) SetLanguageID(languageID string) {
	o.LanguageID = languageID
}

// WithUserID adds the userID to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) WithUserID(userID string) *DeleteUserRoutinglanguageParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the delete user routinglanguage params
func (o *DeleteUserRoutinglanguageParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteUserRoutinglanguageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param languageId
	if err := r.SetPathParam("languageId", o.LanguageID); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
