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

// CustomEventAttribute custom event attribute
//
// swagger:model CustomEventAttribute
type CustomEventAttribute struct {

	// The data type of the custom attribute.
	// Required: true
	DataType *string `json:"dataType"`

	// The value of the custom attribute.
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this custom event attribute
func (m *CustomEventAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataType(formats); err != nil {
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

func (m *CustomEventAttribute) validateDataType(formats strfmt.Registry) error {

	if err := validate.Required("dataType", "body", m.DataType); err != nil {
		return err
	}

	return nil
}

func (m *CustomEventAttribute) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this custom event attribute based on context it is used
func (m *CustomEventAttribute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CustomEventAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomEventAttribute) UnmarshalBinary(b []byte) error {
	var res CustomEventAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
