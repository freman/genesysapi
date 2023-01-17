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

// LocationEmergencyNumber location emergency number
//
// swagger:model LocationEmergencyNumber
type LocationEmergencyNumber struct {

	// e164
	E164 string `json:"e164,omitempty"`

	// number
	Number string `json:"number,omitempty"`

	// The type of emergency number.
	// Enum: [default elin]
	Type string `json:"type,omitempty"`
}

// Validate validates this location emergency number
func (m *LocationEmergencyNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var locationEmergencyNumberTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","elin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		locationEmergencyNumberTypeTypePropEnum = append(locationEmergencyNumberTypeTypePropEnum, v)
	}
}

const (

	// LocationEmergencyNumberTypeDefault captures enum value "default"
	LocationEmergencyNumberTypeDefault string = "default"

	// LocationEmergencyNumberTypeElin captures enum value "elin"
	LocationEmergencyNumberTypeElin string = "elin"
)

// prop value enum
func (m *LocationEmergencyNumber) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, locationEmergencyNumberTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LocationEmergencyNumber) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this location emergency number based on context it is used
func (m *LocationEmergencyNumber) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LocationEmergencyNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LocationEmergencyNumber) UnmarshalBinary(b []byte) error {
	var res LocationEmergencyNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
