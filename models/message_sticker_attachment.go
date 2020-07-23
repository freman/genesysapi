// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageStickerAttachment message sticker attachment
//
// swagger:model MessageStickerAttachment
type MessageStickerAttachment struct {

	// id
	ID string `json:"id,omitempty"`

	// The location of the media, useful for retrieving it
	URL string `json:"url,omitempty"`
}

// Validate validates this message sticker attachment
func (m *MessageStickerAttachment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MessageStickerAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageStickerAttachment) UnmarshalBinary(b []byte) error {
	var res MessageStickerAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
