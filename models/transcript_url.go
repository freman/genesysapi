// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TranscriptURL transcript Url
//
// swagger:model TranscriptUrl
type TranscriptURL struct {

	// The pre-signed S3 URL of the transcript
	URL string `json:"url,omitempty"`
}

// Validate validates this transcript Url
func (m *TranscriptURL) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TranscriptURL) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptURL) UnmarshalBinary(b []byte) error {
	var res TranscriptURL
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
