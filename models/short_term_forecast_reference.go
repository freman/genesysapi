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

// ShortTermForecastReference A pointer to a short term forecast
//
// swagger:model ShortTermForecastReference
type ShortTermForecastReference struct {

	// The description of the short term forecast
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The weekDate of the short term forecast in yyyy-MM-dd format
	// Required: true
	WeekDate *string `json:"weekDate"`
}

// Validate validates this short term forecast reference
func (m *ShortTermForecastReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShortTermForecastReference) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShortTermForecastReference) validateWeekDate(formats strfmt.Registry) error {

	if err := validate.Required("weekDate", "body", m.WeekDate); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShortTermForecastReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShortTermForecastReference) UnmarshalBinary(b []byte) error {
	var res ShortTermForecastReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}