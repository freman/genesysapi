// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserApp Details for a UserApp
//
// swagger:model UserApp
type UserApp struct {

	// config
	Config *UserAppConfigurationInfo `json:"config,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Integration Type for the userApp
	// Read Only: true
	IntegrationType *IntegrationType `json:"integrationType,omitempty"`

	// The name of the userApp, used to distinguish this userApp from others of the same type.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this user app
func (m *UserApp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationType(formats); err != nil {
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

func (m *UserApp) validateConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.Config) { // not required
		return nil
	}

	if m.Config != nil {
		if err := m.Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *UserApp) validateIntegrationType(formats strfmt.Registry) error {
	if swag.IsZero(m.IntegrationType) { // not required
		return nil
	}

	if m.IntegrationType != nil {
		if err := m.IntegrationType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integrationType")
			}
			return err
		}
	}

	return nil
}

func (m *UserApp) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user app based on the context it is used
func (m *UserApp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntegrationType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserApp) contextValidateConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.Config != nil {
		if err := m.Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("config")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("config")
			}
			return err
		}
	}

	return nil
}

func (m *UserApp) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *UserApp) contextValidateIntegrationType(ctx context.Context, formats strfmt.Registry) error {

	if m.IntegrationType != nil {
		if err := m.IntegrationType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrationType")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integrationType")
			}
			return err
		}
	}

	return nil
}

func (m *UserApp) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *UserApp) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserApp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserApp) UnmarshalBinary(b []byte) error {
	var res UserApp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
