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

// MessageMediaData message media data
//
// swagger:model MessageMediaData
type MessageMediaData struct {

	// The optional content length of the the media object, in bytes.
	ContentLengthBytes int32 `json:"contentLengthBytes,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The detected internet media type of the the media object.  If null then the media type should be dictated by the url.
	MediaType string `json:"mediaType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of the media, indicates if the media is in the process of uploading. If the upload fails, the media becomes invalid
	// Read Only: true
	// Enum: [uploading valid invalid]
	Status string `json:"status,omitempty"`

	// The URL returned to upload an attachment
	UploadURL string `json:"uploadUrl,omitempty"`

	// The location of the media, useful for retrieving it
	URL string `json:"url,omitempty"`
}

// Validate validates this message media data
func (m *MessageMediaData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessageMediaData) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var messageMediaDataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["uploading","valid","invalid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		messageMediaDataTypeStatusPropEnum = append(messageMediaDataTypeStatusPropEnum, v)
	}
}

const (

	// MessageMediaDataStatusUploading captures enum value "uploading"
	MessageMediaDataStatusUploading string = "uploading"

	// MessageMediaDataStatusValid captures enum value "valid"
	MessageMediaDataStatusValid string = "valid"

	// MessageMediaDataStatusInvalid captures enum value "invalid"
	MessageMediaDataStatusInvalid string = "invalid"
)

// prop value enum
func (m *MessageMediaData) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, messageMediaDataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MessageMediaData) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessageMediaData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessageMediaData) UnmarshalBinary(b []byte) error {
	var res MessageMediaData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
