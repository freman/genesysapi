// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MediaResult media result
//
// swagger:model MediaResult
type MediaResult struct {

	// media Uri
	MediaURI string `json:"mediaUri,omitempty"`

	// waveform data
	WaveformData []float32 `json:"waveformData"`
}

// Validate validates this media result
func (m *MediaResult) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MediaResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaResult) UnmarshalBinary(b []byte) error {
	var res MediaResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
