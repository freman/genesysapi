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

// ScimError Defines a SCIM error.
//
// swagger:model ScimError
type ScimError struct {

	// The detailed description of the SCIM error.
	// Read Only: true
	Detail string `json:"detail,omitempty"`

	// The list of schemas for the SCIM error.
	// Read Only: true
	Schemas []string `json:"schemas"`

	// The type of SCIM error when httpStatus is a "400" error.
	// Read Only: true
	// Enum: [invalidFilter tooMany uniqueness mutability invalidSyntax invalidPath noTarget invalidValue invalidVers sensitive]
	ScimType string `json:"scimType,omitempty"`

	// The HTTP status code returned for the SCIM error.
	// Read Only: true
	Status string `json:"status,omitempty"`
}

// Validate validates this scim error
func (m *ScimError) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateScimType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scimErrorTypeScimTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["invalidFilter","tooMany","uniqueness","mutability","invalidSyntax","invalidPath","noTarget","invalidValue","invalidVers","sensitive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimErrorTypeScimTypePropEnum = append(scimErrorTypeScimTypePropEnum, v)
	}
}

const (

	// ScimErrorScimTypeInvalidFilter captures enum value "invalidFilter"
	ScimErrorScimTypeInvalidFilter string = "invalidFilter"

	// ScimErrorScimTypeTooMany captures enum value "tooMany"
	ScimErrorScimTypeTooMany string = "tooMany"

	// ScimErrorScimTypeUniqueness captures enum value "uniqueness"
	ScimErrorScimTypeUniqueness string = "uniqueness"

	// ScimErrorScimTypeMutability captures enum value "mutability"
	ScimErrorScimTypeMutability string = "mutability"

	// ScimErrorScimTypeInvalidSyntax captures enum value "invalidSyntax"
	ScimErrorScimTypeInvalidSyntax string = "invalidSyntax"

	// ScimErrorScimTypeInvalidPath captures enum value "invalidPath"
	ScimErrorScimTypeInvalidPath string = "invalidPath"

	// ScimErrorScimTypeNoTarget captures enum value "noTarget"
	ScimErrorScimTypeNoTarget string = "noTarget"

	// ScimErrorScimTypeInvalidValue captures enum value "invalidValue"
	ScimErrorScimTypeInvalidValue string = "invalidValue"

	// ScimErrorScimTypeInvalidVers captures enum value "invalidVers"
	ScimErrorScimTypeInvalidVers string = "invalidVers"

	// ScimErrorScimTypeSensitive captures enum value "sensitive"
	ScimErrorScimTypeSensitive string = "sensitive"
)

// prop value enum
func (m *ScimError) validateScimTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimErrorTypeScimTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimError) validateScimType(formats strfmt.Registry) error {
	if swag.IsZero(m.ScimType) { // not required
		return nil
	}

	// value enum
	if err := m.validateScimTypeEnum("scimType", "body", m.ScimType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this scim error based on the context it is used
func (m *ScimError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDetail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchemas(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScimType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimError) contextValidateDetail(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "detail", "body", string(m.Detail)); err != nil {
		return err
	}

	return nil
}

func (m *ScimError) contextValidateSchemas(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "schemas", "body", []string(m.Schemas)); err != nil {
		return err
	}

	return nil
}

func (m *ScimError) contextValidateScimType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "scimType", "body", string(m.ScimType)); err != nil {
		return err
	}

	return nil
}

func (m *ScimError) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimError) UnmarshalBinary(b []byte) error {
	var res ScimError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
