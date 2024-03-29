// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DocumentThumbnail document thumbnail
//
// swagger:model DocumentThumbnail
type DocumentThumbnail struct {

	// height
	Height int32 `json:"height,omitempty"`

	// image Uri
	ImageURI string `json:"imageUri,omitempty"`

	// resolution
	Resolution string `json:"resolution,omitempty"`

	// width
	Width int32 `json:"width,omitempty"`
}

// Validate validates this document thumbnail
func (m *DocumentThumbnail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this document thumbnail based on context it is used
func (m *DocumentThumbnail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DocumentThumbnail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentThumbnail) UnmarshalBinary(b []byte) error {
	var res DocumentThumbnail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
