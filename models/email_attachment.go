// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EmailAttachment email attachment
//
// swagger:model EmailAttachment
type EmailAttachment struct {

	// attachment Id
	AttachmentID string `json:"attachmentId,omitempty"`

	// content length
	ContentLength int32 `json:"contentLength,omitempty"`

	// content path
	ContentPath string `json:"contentPath,omitempty"`

	// content type
	ContentType string `json:"contentType,omitempty"`

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this email attachment
func (m *EmailAttachment) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailAttachment) UnmarshalBinary(b []byte) error {
	var res EmailAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}