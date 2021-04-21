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

// AnalyticsSessionMetric analytics session metric
//
// swagger:model AnalyticsSessionMetric
type AnalyticsSessionMetric struct {

	// Metric emission date. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EmitDate strfmt.DateTime `json:"emitDate,omitempty"`

	// Unique name of this metric
	Name string `json:"name,omitempty"`

	// The metric value
	Value int64 `json:"value,omitempty"`
}

// Validate validates this analytics session metric
func (m *AnalyticsSessionMetric) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmitDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsSessionMetric) validateEmitDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EmitDate) { // not required
		return nil
	}

	if err := validate.FormatOf("emitDate", "body", "date-time", m.EmitDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsSessionMetric) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsSessionMetric) UnmarshalBinary(b []byte) error {
	var res AnalyticsSessionMetric
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
