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

// FlowOutcome flow outcome
//
// swagger:model FlowOutcome
type FlowOutcome struct {

	// current operation
	CurrentOperation *Operation `json:"currentOperation,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// The flow outcome identifier
	ID string `json:"id,omitempty"`

	// The flow outcome name.
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this flow outcome
func (m *FlowOutcome) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrentOperation(formats); err != nil {
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

func (m *FlowOutcome) validateCurrentOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentOperation) { // not required
		return nil
	}

	if m.CurrentOperation != nil {
		if err := m.CurrentOperation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentOperation")
			}
			return err
		}
	}

	return nil
}

func (m *FlowOutcome) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FlowOutcome) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FlowOutcome) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FlowOutcome) UnmarshalBinary(b []byte) error {
	var res FlowOutcome
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
