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

// FlowOutcomeDivisionView flow outcome division view
//
// swagger:model FlowOutcomeDivisionView
type FlowOutcomeDivisionView struct {

	// The division to which this entity belongs.
	Division *WritableDivision `json:"division,omitempty"`

	// The flow outcome identifier
	ID string `json:"id,omitempty"`

	// The flow outcome name
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this flow outcome division view
func (m *FlowOutcomeDivisionView) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
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

func (m *FlowOutcomeDivisionView) validateDivision(formats strfmt.Registry) error {

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

func (m *FlowOutcomeDivisionView) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FlowOutcomeDivisionView) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowOutcomeDivisionView) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowOutcomeDivisionView) UnmarshalBinary(b []byte) error {
	var res FlowOutcomeDivisionView
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}