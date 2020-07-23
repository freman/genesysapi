// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GDPRJourneyCustomer g d p r journey customer
//
// swagger:model GDPRJourneyCustomer
type GDPRJourneyCustomer struct {

	// An ID of a customer within the Journey System at a point-in-time. Required if `type` is defined.
	ID string `json:"id,omitempty"`

	// The type of the customerId within the Journey System (e.g. cookie). Required if `id` is defined.
	Type string `json:"type,omitempty"`
}

// Validate validates this g d p r journey customer
func (m *GDPRJourneyCustomer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GDPRJourneyCustomer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GDPRJourneyCustomer) UnmarshalBinary(b []byte) error {
	var res GDPRJourneyCustomer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
