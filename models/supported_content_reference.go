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

// SupportedContentReference Reference to supported content profile associated with the integration
//
// swagger:model SupportedContentReference
type SupportedContentReference struct {

	// The SupportedContent unique identifier associated with this integration
	// Required: true
	ID *string `json:"id"`

	// Media types definition for the supported content
	// Read Only: true
	MediaTypes *MediaTypes `json:"mediaTypes,omitempty"`

	// The SupportedContent profile name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The SupportedContent profile URI
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this supported content reference
func (m *SupportedContentReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaTypes(formats); err != nil {
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

func (m *SupportedContentReference) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *SupportedContentReference) validateMediaTypes(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaTypes) { // not required
		return nil
	}

	if m.MediaTypes != nil {
		if err := m.MediaTypes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaTypes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaTypes")
			}
			return err
		}
	}

	return nil
}

func (m *SupportedContentReference) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this supported content reference based on the context it is used
func (m *SupportedContentReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMediaTypes(ctx, formats); err != nil {
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

func (m *SupportedContentReference) contextValidateMediaTypes(ctx context.Context, formats strfmt.Registry) error {

	if m.MediaTypes != nil {
		if err := m.MediaTypes.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaTypes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaTypes")
			}
			return err
		}
	}

	return nil
}

func (m *SupportedContentReference) contextValidateName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *SupportedContentReference) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportedContentReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedContentReference) UnmarshalBinary(b []byte) error {
	var res SupportedContentReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
