// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MessageMedia message media
//
// swagger:model MessageMedia
type MessageMedia struct {

	// The optional content length of the the media object, in bytes.
	ContentLengthBytes int32 `json:"contentLengthBytes,omitempty"`

	// The optional id of the the media object.
	ID string `json:"id,omitempty"`

	// The optional internet media type of the the media object.  If null then the media type should be dictated by the url
	MediaType string `json:"mediaType,omitempty"`

	// The optional name of the the media object.
	Name string `json:"name,omitempty"`

	// The location of the media, useful for retrieving it
	URL string `json:"url,omitempty"`
}

// Validate validates this message media
func (m *MessageMedia) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this message media based on context it is used
func (m *MessageMedia) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MessageMedia) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageMedia) UnmarshalBinary(b []byte) error {
	var res MessageMedia
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
