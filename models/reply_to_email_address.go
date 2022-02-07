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

// ReplyToEmailAddress reply to email address
//
// swagger:model ReplyToEmailAddress
type ReplyToEmailAddress struct {

	// The InboundDomain used for the email address.
	// Required: true
	Domain *DomainEntityRef `json:"domain"`

	// The InboundRoute used for the email address.
	// Required: true
	Route *DomainEntityRef `json:"route"`
}

// Validate validates this reply to email address
func (m *ReplyToEmailAddress) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoute(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplyToEmailAddress) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	if m.Domain != nil {
		if err := m.Domain.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("domain")
			}
			return err
		}
	}

	return nil
}

func (m *ReplyToEmailAddress) validateRoute(formats strfmt.Registry) error {

	if err := validate.Required("route", "body", m.Route); err != nil {
		return err
	}

	if m.Route != nil {
		if err := m.Route.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("route")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplyToEmailAddress) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplyToEmailAddress) UnmarshalBinary(b []byte) error {
	var res ReplyToEmailAddress
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}