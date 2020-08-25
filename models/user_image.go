// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserImage user image
//
// swagger:model UserImage
type UserImage struct {

	// image Uri
	ImageURI string `json:"imageUri,omitempty"`

	// Height and/or width of image. ex: 640x480 or x128
	Resolution string `json:"resolution,omitempty"`
}

// Validate validates this user image
func (m *UserImage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserImage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserImage) UnmarshalBinary(b []byte) error {
	var res UserImage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}