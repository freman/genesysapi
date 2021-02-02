// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SupportedContent Supported content for inbound and outbound messages
//
// swagger:model SupportedContent
type SupportedContent struct {

	// Defines the allowable media that may be accepted for an inbound message or to be sent in an outbound message. The following is an example of allowing all inbound media, and for outbound all images and only mpeg video: {
	//   "mediaTypes": {
	//     "allow": {
	//       "inbound": [{"type": "*/*"}],
	//       "outbound": [{"type": "image/*"}, {"type": "video/mpeg"}]
	//     }
	//   }
	// }
	MediaTypes *MediaTypes `json:"mediaTypes,omitempty"`
}

// Validate validates this supported content
func (m *SupportedContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMediaTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportedContent) validateMediaTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaTypes) { // not required
		return nil
	}

	if m.MediaTypes != nil {
		if err := m.MediaTypes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaTypes")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportedContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedContent) UnmarshalBinary(b []byte) error {
	var res SupportedContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}