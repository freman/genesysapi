// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailsSettings emails settings
//
// swagger:model EmailsSettings
type EmailsSettings struct {

	// sending size limit
	SendingSizeLimit int32 `json:"sendingSizeLimit,omitempty"`
}

// Validate validates this emails settings
func (m *EmailsSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailsSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailsSettings) UnmarshalBinary(b []byte) error {
	var res EmailsSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}