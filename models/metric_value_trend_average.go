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

// MetricValueTrendAverage metric value trend average
//
// swagger:model MetricValueTrendAverage
type MetricValueTrendAverage struct {

	// The targeted end workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	DateEndWorkday strfmt.Date `json:"dateEndWorkday,omitempty"`

	// The targeted reference workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	DateReferenceWorkday strfmt.Date `json:"dateReferenceWorkday,omitempty"`

	// The targeted start workday. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	DateStartWorkday strfmt.Date `json:"dateStartWorkday,omitempty"`

	// The targeted division for the metrics
	// Read Only: true
	Division *Division `json:"division,omitempty"`

	// The targeted performance profile for the average points
	// Read Only: true
	Metric *AddressableEntityRef `json:"metric,omitempty"`

	// The targeted performance profile for the average points
	// Read Only: true
	PerformanceProfile *AddressableEntityRef `json:"performanceProfile,omitempty"`

	// The metric value trend and average
	// Read Only: true
	Result *WorkdayValuesMetricItem `json:"result,omitempty"`

	// The time zone used for aggregating metric values
	// Read Only: true
	Timezone string `json:"timezone,omitempty"`

	// The targeted user for the metrics
	// Read Only: true
	User *UserReference `json:"user,omitempty"`
}

// Validate validates this metric value trend average
func (m *MetricValueTrendAverage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateEndWorkday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateReferenceWorkday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStartWorkday(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePerformanceProfile(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetricValueTrendAverage) validateDateEndWorkday(formats strfmt.Registry) error {

	if swag.IsZero(m.DateEndWorkday) { // not required
		return nil
	}

	if err := validate.FormatOf("dateEndWorkday", "body", "date", m.DateEndWorkday.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MetricValueTrendAverage) validateDateReferenceWorkday(formats strfmt.Registry) error {

	if swag.IsZero(m.DateReferenceWorkday) { // not required
		return nil
	}

	if err := validate.FormatOf("dateReferenceWorkday", "body", "date", m.DateReferenceWorkday.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MetricValueTrendAverage) validateDateStartWorkday(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStartWorkday) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStartWorkday", "body", "date", m.DateStartWorkday.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MetricValueTrendAverage) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *MetricValueTrendAverage) validateMetric(formats strfmt.Registry) error {

	if swag.IsZero(m.Metric) { // not required
		return nil
	}

	if m.Metric != nil {
		if err := m.Metric.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metric")
			}
			return err
		}
	}

	return nil
}

func (m *MetricValueTrendAverage) validatePerformanceProfile(formats strfmt.Registry) error {

	if swag.IsZero(m.PerformanceProfile) { // not required
		return nil
	}

	if m.PerformanceProfile != nil {
		if err := m.PerformanceProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("performanceProfile")
			}
			return err
		}
	}

	return nil
}

func (m *MetricValueTrendAverage) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

func (m *MetricValueTrendAverage) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetricValueTrendAverage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricValueTrendAverage) UnmarshalBinary(b []byte) error {
	var res MetricValueTrendAverage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
