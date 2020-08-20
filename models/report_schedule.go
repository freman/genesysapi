// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportSchedule report schedule
//
// swagger:model ReportSchedule
type ReportSchedule struct {

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	// Required: true
	Interval *string `json:"interval"`

	// last run
	LastRun *ReportRunEntry `json:"lastRun,omitempty"`

	// locale
	Locale string `json:"locale,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	NextFireTime strfmt.DateTime `json:"nextFireTime,omitempty"`

	// parameters
	Parameters map[string]interface{} `json:"parameters,omitempty"`

	// Quartz Cron Expression
	// Required: true
	QuartzCronExpression *string `json:"quartzCronExpression"`

	// report format
	ReportFormat string `json:"reportFormat,omitempty"`

	// Report ID
	// Required: true
	ReportID *string `json:"reportId"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// time period
	TimePeriod string `json:"timePeriod,omitempty"`

	// time zone
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this report schedule
func (m *ReportSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastRun(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNextFireTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuartzCronExpression(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportSchedule) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateLastRun(formats strfmt.Registry) error {

	if swag.IsZero(m.LastRun) { // not required
		return nil
	}

	if m.LastRun != nil {
		if err := m.LastRun.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lastRun")
			}
			return err
		}
	}

	return nil
}

func (m *ReportSchedule) validateNextFireTime(formats strfmt.Registry) error {

	if swag.IsZero(m.NextFireTime) { // not required
		return nil
	}

	if err := validate.FormatOf("nextFireTime", "body", "date-time", m.NextFireTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateQuartzCronExpression(formats strfmt.Registry) error {

	if err := validate.Required("quartzCronExpression", "body", m.QuartzCronExpression); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateReportID(formats strfmt.Registry) error {

	if err := validate.Required("reportId", "body", m.ReportID); err != nil {
		return err
	}

	return nil
}

func (m *ReportSchedule) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportSchedule) UnmarshalBinary(b []byte) error {
	var res ReportSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
