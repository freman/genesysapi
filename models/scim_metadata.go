// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScimMetadata Defines the SCIM metadata.
//
// swagger:model ScimMetadata
type ScimMetadata struct {

	// The last time that the resource was modified. Date time is represented as an "ISO-8601 string", for example, yyyy-MM-ddTHH:mm:ss.SSSZ. Not included with "Schema" and "ResourceType" resources.
	// Read Only: true
	// Format: date-time
	LastModified strfmt.DateTime `json:"lastModified,omitempty"`

	// The URI of the resource.
	// Read Only: true
	// Format: uri
	Location strfmt.URI `json:"location,omitempty"`

	// The type of SCIM resource.
	// Read Only: true
	// Enum: [User Group ServiceProviderConfig ResourceType Schema]
	ResourceType string `json:"resourceType,omitempty"`

	// The version of the resource. Matches the ETag HTTP response header. Not included with "Schema" and "ResourceType" resources.
	// Read Only: true
	Version string `json:"version,omitempty"`
}

// Validate validates this scim metadata
func (m *ScimMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimMetadata) validateLastModified(formats strfmt.Registry) error {
	if swag.IsZero(m.LastModified) { // not required
		return nil
	}

	if err := validate.FormatOf("lastModified", "body", "date-time", m.LastModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScimMetadata) validateLocation(formats strfmt.Registry) error {
	if swag.IsZero(m.Location) { // not required
		return nil
	}

	if err := validate.FormatOf("location", "body", "uri", m.Location.String(), formats); err != nil {
		return err
	}

	return nil
}

var scimMetadataTypeResourceTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["User","Group","ServiceProviderConfig","ResourceType","Schema"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimMetadataTypeResourceTypePropEnum = append(scimMetadataTypeResourceTypePropEnum, v)
	}
}

const (

	// ScimMetadataResourceTypeUser captures enum value "User"
	ScimMetadataResourceTypeUser string = "User"

	// ScimMetadataResourceTypeGroup captures enum value "Group"
	ScimMetadataResourceTypeGroup string = "Group"

	// ScimMetadataResourceTypeServiceProviderConfig captures enum value "ServiceProviderConfig"
	ScimMetadataResourceTypeServiceProviderConfig string = "ServiceProviderConfig"

	// ScimMetadataResourceTypeResourceType captures enum value "ResourceType"
	ScimMetadataResourceTypeResourceType string = "ResourceType"

	// ScimMetadataResourceTypeSchema captures enum value "Schema"
	ScimMetadataResourceTypeSchema string = "Schema"
)

// prop value enum
func (m *ScimMetadata) validateResourceTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimMetadataTypeResourceTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimMetadata) validateResourceType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResourceTypeEnum("resourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this scim metadata based on the context it is used
func (m *ScimMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLastModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLocation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResourceType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimMetadata) contextValidateLastModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "lastModified", "body", strfmt.DateTime(m.LastModified)); err != nil {
		return err
	}

	return nil
}

func (m *ScimMetadata) contextValidateLocation(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "location", "body", strfmt.URI(m.Location)); err != nil {
		return err
	}

	return nil
}

func (m *ScimMetadata) contextValidateResourceType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "resourceType", "body", string(m.ResourceType)); err != nil {
		return err
	}

	return nil
}

func (m *ScimMetadata) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", string(m.Version)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimMetadata) UnmarshalBinary(b []byte) error {
	var res ScimMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
