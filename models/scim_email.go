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

// ScimEmail Defines a SCIM email address.
//
// swagger:model ScimEmail
type ScimEmail struct {

	// Indicates whether the email address is the primary email address.
	Primary bool `json:"primary"`

	// The type of email address. "value" is immutable if "type" is set to "other".
	// Enum: [work other]
	Type string `json:"type,omitempty"`

	// The email address. Is immutable if "type" is set to "other".
	Value string `json:"value,omitempty"`
}

// Validate validates this scim email
func (m *ScimEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scimEmailTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["work","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimEmailTypeTypePropEnum = append(scimEmailTypeTypePropEnum, v)
	}
}

const (

	// ScimEmailTypeWork captures enum value "work"
	ScimEmailTypeWork string = "work"

	// ScimEmailTypeOther captures enum value "other"
	ScimEmailTypeOther string = "other"
)

// prop value enum
func (m *ScimEmail) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimEmailTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimEmail) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scim email based on context it is used
func (m *ScimEmail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScimEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimEmail) UnmarshalBinary(b []byte) error {
	var res ScimEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
