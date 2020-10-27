// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ScimConfigResourceType Defines a SCIM resource.
//
// swagger:model ScimConfigResourceType
type ScimConfigResourceType struct {

	// The description of the resource type.
	// Read Only: true
	Description string `json:"description,omitempty"`

	// The HTTP-addressable endpoint of the resource type. Appears after the base URL.
	// Read Only: true
	Endpoint string `json:"endpoint,omitempty"`

	// The ID of the SCIM resource. Set by the service provider. "caseExact" is set to "true". "mutability" is set to "readOnly". "returned" is set to "always".
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The metadata of the SCIM resource. Only "location" and "resourceType" are set for "ResourceType" resources.
	// Read Only: true
	Meta *ScimMetadata `json:"meta,omitempty"`

	// The name of the resource type.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The URI of the primary or base schema for the resource type.
	// Read Only: true
	Schema string `json:"schema,omitempty"`

	// The list of schema extensions for the resource type.
	// Read Only: true
	SchemaExtensions []*ScimConfigResourceTypeSchemaExtension `json:"schemaExtensions"`

	// The list of supported schemas.
	// Read Only: true
	Schemas []string `json:"schemas"`
}

// Validate validates this scim config resource type
func (m *ScimConfigResourceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemaExtensions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimConfigResourceType) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *ScimConfigResourceType) validateSchemaExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.SchemaExtensions) { // not required
		return nil
	}

	for i := 0; i < len(m.SchemaExtensions); i++ {
		if swag.IsZero(m.SchemaExtensions[i]) { // not required
			continue
		}

		if m.SchemaExtensions[i] != nil {
			if err := m.SchemaExtensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("schemaExtensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimConfigResourceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimConfigResourceType) UnmarshalBinary(b []byte) error {
	var res ScimConfigResourceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
