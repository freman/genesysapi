// Code generated by go-swagger; DO NOT EDIT.

package quality

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

// NewDeleteQualityCalibrationParams creates a new DeleteQualityCalibrationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteQualityCalibrationParams() *DeleteQualityCalibrationParams {
	return &DeleteQualityCalibrationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteQualityCalibrationParamsWithTimeout creates a new DeleteQualityCalibrationParams object
// with the ability to set a timeout on a request.
func NewDeleteQualityCalibrationParamsWithTimeout(timeout time.Duration) *DeleteQualityCalibrationParams {
	return &DeleteQualityCalibrationParams{
		timeout: timeout,
	}
}

// NewDeleteQualityCalibrationParamsWithContext creates a new DeleteQualityCalibrationParams object
// with the ability to set a context for a request.
func NewDeleteQualityCalibrationParamsWithContext(ctx context.Context) *DeleteQualityCalibrationParams {
	return &DeleteQualityCalibrationParams{
		Context: ctx,
	}
}

// NewDeleteQualityCalibrationParamsWithHTTPClient creates a new DeleteQualityCalibrationParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteQualityCalibrationParamsWithHTTPClient(client *http.Client) *DeleteQualityCalibrationParams {
	return &DeleteQualityCalibrationParams{
		HTTPClient: client,
	}
}

/*
DeleteQualityCalibrationParams contains all the parameters to send to the API endpoint

	for the delete quality calibration operation.

	Typically these are written to a http.Request.
*/
type DeleteQualityCalibrationParams struct {

	/* CalibrationID.

	   Calibration ID
	*/
	CalibrationID string

	/* CalibratorID.

	   calibratorId
	*/
	CalibratorID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete quality calibration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQualityCalibrationParams) WithDefaults() *DeleteQualityCalibrationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete quality calibration params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteQualityCalibrationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) WithTimeout(timeout time.Duration) *DeleteQualityCalibrationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) WithContext(ctx context.Context) *DeleteQualityCalibrationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) WithHTTPClient(client *http.Client) *DeleteQualityCalibrationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCalibrationID adds the calibrationID to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) WithCalibrationID(calibrationID string) *DeleteQualityCalibrationParams {
	o.SetCalibrationID(calibrationID)
	return o
}

// SetCalibrationID adds the calibrationId to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) SetCalibrationID(calibrationID string) {
	o.CalibrationID = calibrationID
}

// WithCalibratorID adds the calibratorID to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) WithCalibratorID(calibratorID string) *DeleteQualityCalibrationParams {
	o.SetCalibratorID(calibratorID)
	return o
}

// SetCalibratorID adds the calibratorId to the delete quality calibration params
func (o *DeleteQualityCalibrationParams) SetCalibratorID(calibratorID string) {
	o.CalibratorID = calibratorID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteQualityCalibrationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param calibrationId
	if err := r.SetPathParam("calibrationId", o.CalibrationID); err != nil {
		return err
	}

	// query param calibratorId
	qrCalibratorID := o.CalibratorID
	qCalibratorID := qrCalibratorID
	if qCalibratorID != "" {

		if err := r.SetQueryParam("calibratorId", qCalibratorID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
