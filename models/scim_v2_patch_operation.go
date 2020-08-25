// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScimV2PatchOperation Defines a SCIM PATCH operation. The path and value follow very specific rules based on operation types. See section 3.5.2 "Modifying with PATCH" in RFC 7644 for details.
//
// swagger:model ScimV2PatchOperation
type ScimV2PatchOperation struct {

	// The PATCH operation to perform.
	// Required: true
	// Enum: [add replace remove]
	Op *string `json:"op"`

	// The attribute path that describes the target of the operation. Required for a "remove" operation.
	Path string `json:"path,omitempty"`

	// The value to set in the path.
	Value *JSONNode `json:"value,omitempty"`
}

// Validate validates this scim v2 patch operation
func (m *ScimV2PatchOperation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scimV2PatchOperationTypeOpPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["add","replace","remove"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimV2PatchOperationTypeOpPropEnum = append(scimV2PatchOperationTypeOpPropEnum, v)
	}
}

const (

	// ScimV2PatchOperationOpAdd captures enum value "add"
	ScimV2PatchOperationOpAdd string = "add"

	// ScimV2PatchOperationOpReplace captures enum value "replace"
	ScimV2PatchOperationOpReplace string = "replace"

	// ScimV2PatchOperationOpRemove captures enum value "remove"
	ScimV2PatchOperationOpRemove string = "remove"
)

// prop value enum
func (m *ScimV2PatchOperation) validateOpEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimV2PatchOperationTypeOpPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimV2PatchOperation) validateOp(formats strfmt.Registry) error {

	if err := validate.Required("op", "body", m.Op); err != nil {
		return err
	}

	// value enum
	if err := m.validateOpEnum("op", "body", *m.Op); err != nil {
		return err
	}

	return nil
}

func (m *ScimV2PatchOperation) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if m.Value != nil {
		if err := m.Value.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("value")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimV2PatchOperation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimV2PatchOperation) UnmarshalBinary(b []byte) error {
	var res ScimV2PatchOperation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}