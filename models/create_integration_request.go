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

// CreateIntegrationRequest Details for an Integration
//
// swagger:model CreateIntegrationRequest
type CreateIntegrationRequest struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Type of the integration to create.
	// Required: true
	IntegrationType *IntegrationType `json:"integrationType"`

	// The name of the integration, used to distinguish this integration from others of the same type.
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this create integration request
func (m *CreateIntegrationRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationType(formats); err != nil {
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

func (m *CreateIntegrationRequest) validateIntegrationType(formats strfmt.Registry) error {

	if err := validate.Required("integrationType", "body", m.IntegrationType); err != nil {
		return err
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

func (m *CreateIntegrationRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CreateIntegrationRequest) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create integration request based on the context it is used
func (m *CreateIntegrationRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntegrationType(ctx, formats); err != nil {
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

func (m *CreateIntegrationRequest) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CreateIntegrationRequest) contextValidateIntegrationType(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateIntegrationRequest) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateIntegrationRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateIntegrationRequest) UnmarshalBinary(b []byte) error {
	var res CreateIntegrationRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
