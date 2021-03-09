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

// GeneralProgramJobRequest general program job request
//
// swagger:model GeneralProgramJobRequest
type GeneralProgramJobRequest struct {

	// The dialect of the topics to link with the general program, dialect format is {language}-{country} where language follows ISO 639-1 standard and country follows ISO 3166-1 alpha 2 standard
	// Required: true
	// Enum: [en-US es-US en-AU en-GB]
	Dialect *string `json:"dialect"`

	// The mode to use for the general program job, default value is Skip
	// Enum: [Skip Merge]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this general program job request
func (m *GeneralProgramJobRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDialect(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var generalProgramJobRequestTypeDialectPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-US","es-US","en-AU","en-GB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		generalProgramJobRequestTypeDialectPropEnum = append(generalProgramJobRequestTypeDialectPropEnum, v)
	}
}

const (

	// GeneralProgramJobRequestDialectEnUS captures enum value "en-US"
	GeneralProgramJobRequestDialectEnUS string = "en-US"

	// GeneralProgramJobRequestDialectEsUS captures enum value "es-US"
	GeneralProgramJobRequestDialectEsUS string = "es-US"

	// GeneralProgramJobRequestDialectEnAU captures enum value "en-AU"
	GeneralProgramJobRequestDialectEnAU string = "en-AU"

	// GeneralProgramJobRequestDialectEnGB captures enum value "en-GB"
	GeneralProgramJobRequestDialectEnGB string = "en-GB"
)

// prop value enum
func (m *GeneralProgramJobRequest) validateDialectEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, generalProgramJobRequestTypeDialectPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GeneralProgramJobRequest) validateDialect(formats strfmt.Registry) error {

	if err := validate.Required("dialect", "body", m.Dialect); err != nil {
		return err
	}

	// value enum
	if err := m.validateDialectEnum("dialect", "body", *m.Dialect); err != nil {
		return err
	}

	return nil
}

var generalProgramJobRequestTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Skip","Merge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		generalProgramJobRequestTypeModePropEnum = append(generalProgramJobRequestTypeModePropEnum, v)
	}
}

const (

	// GeneralProgramJobRequestModeSkip captures enum value "Skip"
	GeneralProgramJobRequestModeSkip string = "Skip"

	// GeneralProgramJobRequestModeMerge captures enum value "Merge"
	GeneralProgramJobRequestModeMerge string = "Merge"
)

// prop value enum
func (m *GeneralProgramJobRequest) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, generalProgramJobRequestTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GeneralProgramJobRequest) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeneralProgramJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralProgramJobRequest) UnmarshalBinary(b []byte) error {
	var res GeneralProgramJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
