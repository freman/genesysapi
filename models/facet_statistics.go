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

// FacetStatistics facet statistics
//
// swagger:model FacetStatistics
type FacetStatistics struct {

	// count
	Count int64 `json:"count,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateMax strfmt.DateTime `json:"dateMax,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateMin strfmt.DateTime `json:"dateMin,omitempty"`

	// max
	Max float64 `json:"max,omitempty"`

	// mean
	Mean float64 `json:"mean,omitempty"`

	// min
	Min float64 `json:"min,omitempty"`

	// std deviation
	StdDeviation float64 `json:"stdDeviation,omitempty"`
}

// Validate validates this facet statistics
func (m *FacetStatistics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateMax(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateMin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacetStatistics) validateDateMax(formats strfmt.Registry) error {

	if swag.IsZero(m.DateMax) { // not required
		return nil
	}

	if err := validate.FormatOf("dateMax", "body", "date-time", m.DateMax.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FacetStatistics) validateDateMin(formats strfmt.Registry) error {

	if swag.IsZero(m.DateMin) { // not required
		return nil
	}

	if err := validate.FormatOf("dateMin", "body", "date-time", m.DateMin.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacetStatistics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacetStatistics) UnmarshalBinary(b []byte) error {
	var res FacetStatistics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
