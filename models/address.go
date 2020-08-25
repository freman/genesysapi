// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Address address
//
// swagger:model Address
type Address struct {

	// The displayable address. This field is acquired from the Address Normalization Table.  The addressRaw could have gone through some transformations, such as only using the numeric portion, before being run through the Address Normalization Table.
	AddressDisplayable string `json:"addressDisplayable,omitempty"`

	// The normalized address. This field is acquired from the Address Normalization Table.  The addressRaw could have gone through some transformations, such as only using the numeric portion, before being run through the Address Normalization Table.
	AddressNormalized string `json:"addressNormalized,omitempty"`

	// The address as close to the bits on the wire as possible.
	AddressRaw string `json:"addressRaw,omitempty"`

	// This will be nameRaw if present, or a locality lookup of the address field otherwise.
	Name string `json:"name,omitempty"`

	// The name as close to the bits on the wire as possible.
	NameRaw string `json:"nameRaw,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}