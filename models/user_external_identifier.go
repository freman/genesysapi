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

// UserExternalIdentifier Defines a link between an External Identifier and Authority pair to a Entity Type and Entity Identifier pair. Represents the two way, one to one mapping of an External Authority or System of Record's identifier to a PureCloud entity. e.g. (ExternalId='05001',Authority='XyzCRM') to (entityType=user,entityId='8eb03b33-3acb-4bc1-a244-50b9b9f19495')
//
// swagger:model UserExternalIdentifier
type UserExternalIdentifier struct {

	// Authority or System of Record which owns the External Identifier
	// Required: true
	AuthorityName *string `json:"authorityName"`

	// External Key
	// Required: true
	ExternalKey *string `json:"externalKey"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this user external identifier
func (m *UserExternalIdentifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthorityName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalKey(formats); err != nil {
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

func (m *UserExternalIdentifier) validateAuthorityName(formats strfmt.Registry) error {

	if err := validate.Required("authorityName", "body", m.AuthorityName); err != nil {
		return err
	}

	return nil
}

func (m *UserExternalIdentifier) validateExternalKey(formats strfmt.Registry) error {

	if err := validate.Required("externalKey", "body", m.ExternalKey); err != nil {
		return err
	}

	return nil
}

func (m *UserExternalIdentifier) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user external identifier based on the context it is used
func (m *UserExternalIdentifier) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserExternalIdentifier) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserExternalIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserExternalIdentifier) UnmarshalBinary(b []byte) error {
	var res UserExternalIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
