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

// JourneyCustomerSession journey customer session
//
// swagger:model JourneyCustomerSession
type JourneyCustomerSession struct {

	// An ID of a Customer/User's session within the Journey System at a point-in-time
	// Required: true
	ID *string `json:"id"`

	// The type of the Customer/User's session within the Journey System (e.g. web, app)
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this journey customer session
func (m *JourneyCustomerSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JourneyCustomerSession) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *JourneyCustomerSession) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JourneyCustomerSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JourneyCustomerSession) UnmarshalBinary(b []byte) error {
	var res JourneyCustomerSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}