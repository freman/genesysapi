// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConversationContentLocation Location object.
//
// swagger:model ConversationContentLocation
type ConversationContentLocation struct {

	// Location postal address.
	Address string `json:"address,omitempty"`

	// Latitude of the location.
	Latitude float64 `json:"latitude,omitempty"`

	// Longitude of the location.
	Longitude float64 `json:"longitude,omitempty"`

	// Location name.
	Text string `json:"text,omitempty"`

	// URL of the Location.
	URL string `json:"url,omitempty"`
}

// Validate validates this conversation content location
func (m *ConversationContentLocation) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationContentLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationContentLocation) UnmarshalBinary(b []byte) error {
	var res ConversationContentLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
