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

// TranscriptionEngines transcription engines
//
// swagger:model TranscriptionEngines
type TranscriptionEngines struct {

	// dialects
	Dialects []string `json:"dialects"`

	// engine
	// Enum: [Genesys GenesysExtended]
	Engine string `json:"engine,omitempty"`
}

// Validate validates this transcription engines
func (m *TranscriptionEngines) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var transcriptionEnginesTypeEnginePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Genesys","GenesysExtended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transcriptionEnginesTypeEnginePropEnum = append(transcriptionEnginesTypeEnginePropEnum, v)
	}
}

const (

	// TranscriptionEnginesEngineGenesys captures enum value "Genesys"
	TranscriptionEnginesEngineGenesys string = "Genesys"

	// TranscriptionEnginesEngineGenesysExtended captures enum value "GenesysExtended"
	TranscriptionEnginesEngineGenesysExtended string = "GenesysExtended"
)

// prop value enum
func (m *TranscriptionEngines) validateEngineEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transcriptionEnginesTypeEnginePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TranscriptionEngines) validateEngine(formats strfmt.Registry) error {
	if swag.IsZero(m.Engine) { // not required
		return nil
	}

	// value enum
	if err := m.validateEngineEnum("engine", "body", m.Engine); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transcription engines based on context it is used
func (m *TranscriptionEngines) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TranscriptionEngines) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptionEngines) UnmarshalBinary(b []byte) error {
	var res TranscriptionEngines
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
