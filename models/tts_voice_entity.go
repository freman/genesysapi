// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TtsVoiceEntity tts voice entity
//
// swagger:model TtsVoiceEntity
type TtsVoiceEntity struct {

	// Ths TTS engine this voice belongs to
	// Required: true
	Engine *TtsEngineEntity `json:"engine"`

	// The gender of the TTS voice
	// Required: true
	Gender *string `json:"gender"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The voice is the default voice for its language
	IsDefault bool `json:"isDefault"`

	// The language supported by the TTS voice
	// Required: true
	Language *string `json:"language"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this tts voice entity
func (m *TtsVoiceEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEngine(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGender(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TtsVoiceEntity) validateEngine(formats strfmt.Registry) error {

	if err := validate.Required("engine", "body", m.Engine); err != nil {
		return err
	}

	if m.Engine != nil {
		if err := m.Engine.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("engine")
			}
			return err
		}
	}

	return nil
}

func (m *TtsVoiceEntity) validateGender(formats strfmt.Registry) error {

	if err := validate.Required("gender", "body", m.Gender); err != nil {
		return err
	}

	return nil
}

func (m *TtsVoiceEntity) validateLanguage(formats strfmt.Registry) error {

	if err := validate.Required("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *TtsVoiceEntity) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TtsVoiceEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TtsVoiceEntity) UnmarshalBinary(b []byte) error {
	var res TtsVoiceEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
