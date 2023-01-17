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

// SocialHandle social handle
//
// swagger:model SocialHandle
type SocialHandle struct {

	// type
	// Enum: [TWITTER]
	Type string `json:"type,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this social handle
func (m *SocialHandle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var socialHandleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TWITTER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		socialHandleTypeTypePropEnum = append(socialHandleTypeTypePropEnum, v)
	}
}

const (

	// SocialHandleTypeTWITTER captures enum value "TWITTER"
	SocialHandleTypeTWITTER string = "TWITTER"
)

// prop value enum
func (m *SocialHandle) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, socialHandleTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SocialHandle) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this social handle based on context it is used
func (m *SocialHandle) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SocialHandle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SocialHandle) UnmarshalBinary(b []byte) error {
	var res SocialHandle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
