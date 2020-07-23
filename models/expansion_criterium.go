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

// ExpansionCriterium expansion criterium
//
// swagger:model ExpansionCriterium
type ExpansionCriterium struct {

	// threshold
	Threshold float64 `json:"threshold,omitempty"`

	// type
	// Enum: [TIMEOUT_SECONDS]
	Type string `json:"type,omitempty"`
}

// Validate validates this expansion criterium
func (m *ExpansionCriterium) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var expansionCriteriumTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TIMEOUT_SECONDS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		expansionCriteriumTypeTypePropEnum = append(expansionCriteriumTypeTypePropEnum, v)
	}
}

const (

	// ExpansionCriteriumTypeTIMEOUTSECONDS captures enum value "TIMEOUT_SECONDS"
	ExpansionCriteriumTypeTIMEOUTSECONDS string = "TIMEOUT_SECONDS"
)

// prop value enum
func (m *ExpansionCriterium) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, expansionCriteriumTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExpansionCriterium) validateType(formats strfmt.Registry) error {

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
func (m *ExpansionCriterium) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExpansionCriterium) UnmarshalBinary(b []byte) error {
	var res ExpansionCriterium
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
