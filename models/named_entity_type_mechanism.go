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

// NamedEntityTypeMechanism named entity type mechanism
//
// swagger:model NamedEntityTypeMechanism
type NamedEntityTypeMechanism struct {

	// The items that define the named entity type.
	// Required: true
	Items []*NamedEntityTypeItem `json:"items"`

	// Whether the named entity type is restricted to the items provided. Default: false
	Restricted bool `json:"restricted"`

	// The type of the mechanism.
	// Required: true
	// Enum: [DynamicList List Regex Unknown]
	Type *string `json:"type"`
}

// Validate validates this named entity type mechanism
func (m *NamedEntityTypeMechanism) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedEntityTypeMechanism) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("items", "body", m.Items); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var namedEntityTypeMechanismTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DynamicList","List","Regex","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		namedEntityTypeMechanismTypeTypePropEnum = append(namedEntityTypeMechanismTypeTypePropEnum, v)
	}
}

const (

	// NamedEntityTypeMechanismTypeDynamicList captures enum value "DynamicList"
	NamedEntityTypeMechanismTypeDynamicList string = "DynamicList"

	// NamedEntityTypeMechanismTypeList captures enum value "List"
	NamedEntityTypeMechanismTypeList string = "List"

	// NamedEntityTypeMechanismTypeRegex captures enum value "Regex"
	NamedEntityTypeMechanismTypeRegex string = "Regex"

	// NamedEntityTypeMechanismTypeUnknown captures enum value "Unknown"
	NamedEntityTypeMechanismTypeUnknown string = "Unknown"
)

// prop value enum
func (m *NamedEntityTypeMechanism) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, namedEntityTypeMechanismTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NamedEntityTypeMechanism) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this named entity type mechanism based on the context it is used
func (m *NamedEntityTypeMechanism) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NamedEntityTypeMechanism) contextValidateItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Items); i++ {

		if m.Items[i] != nil {
			if err := m.Items[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NamedEntityTypeMechanism) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NamedEntityTypeMechanism) UnmarshalBinary(b []byte) error {
	var res NamedEntityTypeMechanism
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
