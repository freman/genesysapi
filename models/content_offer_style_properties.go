// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContentOfferStyleProperties content offer style properties
//
// swagger:model ContentOfferStyleProperties
type ContentOfferStyleProperties struct {

	// Background color of the offer. (eg. #000000)
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Text color of the offer. (eg. #FF0000)
	Color string `json:"color,omitempty"`

	// Padding of the offer. (eg. 10px)
	Padding string `json:"padding,omitempty"`
}

// Validate validates this content offer style properties
func (m *ContentOfferStyleProperties) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this content offer style properties based on context it is used
func (m *ContentOfferStyleProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContentOfferStyleProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentOfferStyleProperties) UnmarshalBinary(b []byte) error {
	var res ContentOfferStyleProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
