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

// ScimConfigResourceTypeSchemaExtension Defines a SCIM resource type's schema extension.
//
// swagger:model ScimConfigResourceTypeSchemaExtension
type ScimConfigResourceTypeSchemaExtension struct {

	// Indicates whether a schema extension is required.
	// Read Only: true
	Required *bool `json:"required"`

	// The URI of an extended schema, for example, "urn:edu:2.0:Staff". Must be equal to the "id" attribute of a schema.
	// Read Only: true
	Schema string `json:"schema,omitempty"`
}

// Validate validates this scim config resource type schema extension
func (m *ScimConfigResourceTypeSchemaExtension) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this scim config resource type schema extension based on the context it is used
func (m *ScimConfigResourceTypeSchemaExtension) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRequired(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimConfigResourceTypeSchemaExtension) contextValidateRequired(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "required", "body", m.Required); err != nil {
		return err
	}

	return nil
}

func (m *ScimConfigResourceTypeSchemaExtension) contextValidateSchema(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "schema", "body", string(m.Schema)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimConfigResourceTypeSchemaExtension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimConfigResourceTypeSchemaExtension) UnmarshalBinary(b []byte) error {
	var res ScimConfigResourceTypeSchemaExtension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
