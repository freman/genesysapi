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

// WebDeploymentsJourneyContext web deployments journey context
//
// swagger:model WebDeploymentsJourneyContext
type WebDeploymentsJourneyContext struct {

	// Journey customer information. Used for linking the authenticated customer with the journey.
	Customer *JourneyCustomer `json:"customer,omitempty"`

	// Contains the Journey System's customer session details.
	CustomerSession *JourneyCustomerSession `json:"customerSession,omitempty"`
}

// Validate validates this web deployments journey context
func (m *WebDeploymentsJourneyContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomerSession(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebDeploymentsJourneyContext) validateCustomer(formats strfmt.Registry) error {
	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentsJourneyContext) validateCustomerSession(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerSession) { // not required
		return nil
	}

	if m.CustomerSession != nil {
		if err := m.CustomerSession.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerSession")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerSession")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this web deployments journey context based on the context it is used
func (m *WebDeploymentsJourneyContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCustomerSession(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebDeploymentsJourneyContext) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

	if m.Customer != nil {
		if err := m.Customer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *WebDeploymentsJourneyContext) contextValidateCustomerSession(ctx context.Context, formats strfmt.Registry) error {

	if m.CustomerSession != nil {
		if err := m.CustomerSession.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customerSession")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customerSession")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebDeploymentsJourneyContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebDeploymentsJourneyContext) UnmarshalBinary(b []byte) error {
	var res WebDeploymentsJourneyContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
