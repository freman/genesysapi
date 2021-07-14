// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// WebMessagingOfferFields web messaging offer fields
//
// swagger:model WebMessagingOfferFields
type WebMessagingOfferFields struct {

	// Text value to be used when inviting a visitor to engage with a web messaging offer.
	OfferText string `json:"offerText,omitempty"`
}

// Validate validates this web messaging offer fields
func (m *WebMessagingOfferFields) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebMessagingOfferFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingOfferFields) UnmarshalBinary(b []byte) error {
	var res WebMessagingOfferFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
