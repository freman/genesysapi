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

// ClonedUser Represents a cloned user in a trustor organization.
//
// swagger:model ClonedUser
type ClonedUser struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The ID of the trustor organization this clone exists in.
	// Read Only: true
	Trustor *DomainEntityRef `json:"trustor,omitempty"`
}

// Validate validates this cloned user
func (m *ClonedUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrustor(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClonedUser) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ClonedUser) validateTrustor(formats strfmt.Registry) error {
	if swag.IsZero(m.Trustor) { // not required
		return nil
	}

	if m.Trustor != nil {
		if err := m.Trustor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trustor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trustor")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this cloned user based on the context it is used
func (m *ClonedUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrustor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClonedUser) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ClonedUser) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *ClonedUser) contextValidateTrustor(ctx context.Context, formats strfmt.Registry) error {

	if m.Trustor != nil {
		if err := m.Trustor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trustor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trustor")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClonedUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClonedUser) UnmarshalBinary(b []byte) error {
	var res ClonedUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
