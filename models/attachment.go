// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Attachment attachment
//
// swagger:model Attachment
type Attachment struct {

	// The unique identifier for the attachment.
	AttachmentID string `json:"attachmentId,omitempty"`

	// The length of the attachment file.
	ContentLength int32 `json:"contentLength,omitempty"`

	// The type of file the attachment is.
	ContentType string `json:"contentType,omitempty"`

	// The content uri of the attachment. If set, this is commonly a public api download location.
	ContentURI string `json:"contentUri,omitempty"`

	// Whether or not the attachment was attached inline.,
	InlineImage bool `json:"inlineImage,omitempty"`

	// The name of the attachment.
	Name string `json:"name,omitempty"`
}

// Validate validates this attachment
func (m *Attachment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Attachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Attachment) UnmarshalBinary(b []byte) error {
	var res Attachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
