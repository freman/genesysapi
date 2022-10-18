// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WebMessagingAttachment Attachment object.
//
// swagger:model WebMessagingAttachment
type WebMessagingAttachment struct {

	// The file size associated with the file
	// Read Only: true
	FileSize int32 `json:"fileSize,omitempty"`

	// Suggested file name for attachment.
	// Read Only: true
	Filename string `json:"filename,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The type of attachment this instance represents.
	// Read Only: true
	// Enum: [Image Video Audio File Link]
	MediaType string `json:"mediaType,omitempty"`

	// Attachment mime type (https://www.iana.org/assignments/media-types/media-types.xhtml).
	// Read Only: true
	Mime string `json:"mime,omitempty"`

	// Secure hash of the attachment content.
	// Read Only: true
	Sha256 string `json:"sha256,omitempty"`

	// Text associated with attachment such as an image caption.
	// Read Only: true
	Text string `json:"text,omitempty"`

	// URL of the attachment.
	// Read Only: true
	URL string `json:"url,omitempty"`
}

// Validate validates this web messaging attachment
func (m *WebMessagingAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webMessagingAttachmentTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Image","Video","Audio","File","Link"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webMessagingAttachmentTypeMediaTypePropEnum = append(webMessagingAttachmentTypeMediaTypePropEnum, v)
	}
}

const (

	// WebMessagingAttachmentMediaTypeImage captures enum value "Image"
	WebMessagingAttachmentMediaTypeImage string = "Image"

	// WebMessagingAttachmentMediaTypeVideo captures enum value "Video"
	WebMessagingAttachmentMediaTypeVideo string = "Video"

	// WebMessagingAttachmentMediaTypeAudio captures enum value "Audio"
	WebMessagingAttachmentMediaTypeAudio string = "Audio"

	// WebMessagingAttachmentMediaTypeFile captures enum value "File"
	WebMessagingAttachmentMediaTypeFile string = "File"

	// WebMessagingAttachmentMediaTypeLink captures enum value "Link"
	WebMessagingAttachmentMediaTypeLink string = "Link"
)

// prop value enum
func (m *WebMessagingAttachment) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webMessagingAttachmentTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebMessagingAttachment) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebMessagingAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebMessagingAttachment) UnmarshalBinary(b []byte) error {
	var res WebMessagingAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
