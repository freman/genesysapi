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

// AssociatedValueField associated value field
//
// swagger:model AssociatedValueField
type AssociatedValueField struct {

	// The data type of the value field.
	// Required: true
	// Enum: [Number Integer]
	DataType *string `json:"dataType"`

	// The field name for extracting value from event.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this associated value field
func (m *AssociatedValueField) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var associatedValueFieldTypeDataTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Number","Integer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		associatedValueFieldTypeDataTypePropEnum = append(associatedValueFieldTypeDataTypePropEnum, v)
	}
}

const (

	// AssociatedValueFieldDataTypeNumber captures enum value "Number"
	AssociatedValueFieldDataTypeNumber string = "Number"

	// AssociatedValueFieldDataTypeInteger captures enum value "Integer"
	AssociatedValueFieldDataTypeInteger string = "Integer"
)

// prop value enum
func (m *AssociatedValueField) validateDataTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, associatedValueFieldTypeDataTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssociatedValueField) validateDataType(formats strfmt.Registry) error {

	if err := validate.Required("dataType", "body", m.DataType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDataTypeEnum("dataType", "body", *m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *AssociatedValueField) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this associated value field based on context it is used
func (m *AssociatedValueField) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssociatedValueField) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssociatedValueField) UnmarshalBinary(b []byte) error {
	var res AssociatedValueField
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
