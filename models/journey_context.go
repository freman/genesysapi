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

// JourneyContext journey context
//
// swagger:model JourneyContext
type JourneyContext struct {

	// A subset of the Journey System's customer data at a point-in-time (for external linkage and internal usage/context)
	// Required: true
	Customer *JourneyCustomer `json:"customer"`

	// A subset of the Journey System's tracked customer session data at a point-in-time (for external linkage and internal usage/context)
	CustomerSession *JourneyCustomerSession `json:"customerSession,omitempty"`

	// A subset of the Journey System's action data relevant to a part of a conversation (for external linkage and internal usage/context)
	TriggeringAction *JourneyAction `json:"triggeringAction,omitempty"`
}

// Validate validates this journey context
func (m *JourneyContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerSession(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTriggeringAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JourneyContext) validateCustomer(formats strfmt.Registry) error {

	if err := validate.Required("customer", "body", m.Customer); err != nil {
		return err
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *JourneyContext) validateCustomerSession(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomerSession) { // not required
		return nil
	}

	if m.CustomerSession != nil {
		if err := m.CustomerSession.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerSession")
			}
			return err
		}
	}

	return nil
}

func (m *JourneyContext) validateTriggeringAction(formats strfmt.Registry) error {

	if swag.IsZero(m.TriggeringAction) { // not required
		return nil
	}

	if m.TriggeringAction != nil {
		if err := m.TriggeringAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("triggeringAction")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JourneyContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JourneyContext) UnmarshalBinary(b []byte) error {
	var res JourneyContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
