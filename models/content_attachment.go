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

// ContentAttachment Attachment object.
//
// swagger:model ContentAttachment
type ContentAttachment struct {

	// Suggested file name for attachment.
	Filename string `json:"filename,omitempty"`

	// Provider specific ID for attachment. For example, a LINE sticker ID.
	ID string `json:"id,omitempty"`

	// The type of attachment this instance represents.
	// Required: true
	// Enum: [Image Video Audio File Link]
	MediaType *string `json:"mediaType"`

	// Attachment mime type (https://www.iana.org/assignments/media-types/media-types.xhtml).
	Mime string `json:"mime,omitempty"`

	// Secure hash of the attachment content.
	Sha256 string `json:"sha256,omitempty"`

	// Text associated with attachment such as an image caption.
	Text string `json:"text,omitempty"`

	// URL of the attachment.
	URL string `json:"url,omitempty"`
}

// Validate validates this content attachment
func (m *ContentAttachment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contentAttachmentTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Image","Video","Audio","File","Link"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentAttachmentTypeMediaTypePropEnum = append(contentAttachmentTypeMediaTypePropEnum, v)
	}
}

const (

	// ContentAttachmentMediaTypeImage captures enum value "Image"
	ContentAttachmentMediaTypeImage string = "Image"

	// ContentAttachmentMediaTypeVideo captures enum value "Video"
	ContentAttachmentMediaTypeVideo string = "Video"

	// ContentAttachmentMediaTypeAudio captures enum value "Audio"
	ContentAttachmentMediaTypeAudio string = "Audio"

	// ContentAttachmentMediaTypeFile captures enum value "File"
	ContentAttachmentMediaTypeFile string = "File"

	// ContentAttachmentMediaTypeLink captures enum value "Link"
	ContentAttachmentMediaTypeLink string = "Link"
)

// prop value enum
func (m *ContentAttachment) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contentAttachmentTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContentAttachment) validateMediaType(formats strfmt.Registry) error {

	if err := validate.Required("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", *m.MediaType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentAttachment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentAttachment) UnmarshalBinary(b []byte) error {
	var res ContentAttachment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
