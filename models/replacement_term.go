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

// ReplacementTerm replacement term
//
// swagger:model ReplacementTerm
type ReplacementTerm struct {

	// existing value
	ExistingValue string `json:"existingValue,omitempty"`

	// type
	// Enum: [NAME ADDRESS PHONE EMAIL TWITTER]
	Type string `json:"type,omitempty"`

	// updated value
	UpdatedValue string `json:"updatedValue,omitempty"`
}

// Validate validates this replacement term
func (m *ReplacementTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var replacementTermTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NAME","ADDRESS","PHONE","EMAIL","TWITTER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		replacementTermTypeTypePropEnum = append(replacementTermTypeTypePropEnum, v)
	}
}

const (

	// ReplacementTermTypeNAME captures enum value "NAME"
	ReplacementTermTypeNAME string = "NAME"

	// ReplacementTermTypeADDRESS captures enum value "ADDRESS"
	ReplacementTermTypeADDRESS string = "ADDRESS"

	// ReplacementTermTypePHONE captures enum value "PHONE"
	ReplacementTermTypePHONE string = "PHONE"

	// ReplacementTermTypeEMAIL captures enum value "EMAIL"
	ReplacementTermTypeEMAIL string = "EMAIL"

	// ReplacementTermTypeTWITTER captures enum value "TWITTER"
	ReplacementTermTypeTWITTER string = "TWITTER"
)

// prop value enum
func (m *ReplacementTerm) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, replacementTermTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReplacementTerm) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplacementTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplacementTerm) UnmarshalBinary(b []byte) error {
	var res ReplacementTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}