// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FieldConfig field config
//
// swagger:model FieldConfig
type FieldConfig struct {

	// entity type
	// Enum: [person group org externalContact]
	EntityType string `json:"entityType,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// schema version
	SchemaVersion string `json:"schemaVersion,omitempty"`

	// sections
	Sections []*Section `json:"sections"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this field config
func (m *FieldConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSections(formats); err != nil {
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

var fieldConfigTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["person","group","org","externalContact"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fieldConfigTypeEntityTypePropEnum = append(fieldConfigTypeEntityTypePropEnum, v)
	}
}

const (

	// FieldConfigEntityTypePerson captures enum value "person"
	FieldConfigEntityTypePerson string = "person"

	// FieldConfigEntityTypeGroup captures enum value "group"
	FieldConfigEntityTypeGroup string = "group"

	// FieldConfigEntityTypeOrg captures enum value "org"
	FieldConfigEntityTypeOrg string = "org"

	// FieldConfigEntityTypeExternalContact captures enum value "externalContact"
	FieldConfigEntityTypeExternalContact string = "externalContact"
)

// prop value enum
func (m *FieldConfig) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fieldConfigTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FieldConfig) validateEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) validateSections(formats strfmt.Registry) error {
	if swag.IsZero(m.Sections) { // not required
		return nil
	}

	for i := 0; i < len(m.Sections); i++ {
		if swag.IsZero(m.Sections[i]) { // not required
			continue
		}

		if m.Sections[i] != nil {
			if err := m.Sections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FieldConfig) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this field config based on the context it is used
func (m *FieldConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSections(ctx, formats); err != nil {
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

func (m *FieldConfig) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *FieldConfig) contextValidateSections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sections); i++ {

		if m.Sections[i] != nil {
			if err := m.Sections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FieldConfig) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FieldConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FieldConfig) UnmarshalBinary(b []byte) error {
	var res FieldConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
