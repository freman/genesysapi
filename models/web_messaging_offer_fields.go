// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebMessagingOfferFields web messaging offer fields
//
// swagger:model WebMessagingOfferFields
type WebMessagingOfferFields struct {

	// Flow to be invoked, overrides default flow when specified.
	ArchitectFlow *AddressableEntityRef `json:"architectFlow,omitempty"`

	// Text value to be used when inviting a visitor to engage with a web messaging offer.
	OfferText string `json:"offerText,omitempty"`
}

// Validate validates this web messaging offer fields
func (m *WebMessagingOfferFields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchitectFlow(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebMessagingOfferFields) validateArchitectFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchitectFlow) { // not required
		return nil
	}

	if m.ArchitectFlow != nil {
		if err := m.ArchitectFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("architectFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("architectFlow")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this web messaging offer fields based on the context it is used
func (m *WebMessagingOfferFields) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchitectFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebMessagingOfferFields) contextValidateArchitectFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.ArchitectFlow != nil {
		if err := m.ArchitectFlow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("architectFlow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("architectFlow")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebMessagingOfferFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingOfferFields) UnmarshalBinary(b []byte) error {
	var res WebMessagingOfferFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
