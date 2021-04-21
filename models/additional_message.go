// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AdditionalMessage additional message
//
// swagger:model AdditionalMessage
type AdditionalMessage struct {

	// The media ids associated with the text message. See https://developer.genesys.cloud/api/rest/v2/conversations/messaging-media-upload for example usage.
	MediaIds []string `json:"mediaIds"`

	// The messaging template use to send a predefined canned response with the message
	MessagingTemplate *MessagingTemplateRequest `json:"messagingTemplate,omitempty"`

	// The sticker ids associated with the text message.
	StickerIds []string `json:"stickerIds"`

	// The body of the text message.
	// Required: true
	TextBody *string `json:"textBody"`
}

// Validate validates this additional message
func (m *AdditionalMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessagingTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTextBody(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdditionalMessage) validateMessagingTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.MessagingTemplate) { // not required
		return nil
	}

	if m.MessagingTemplate != nil {
		if err := m.MessagingTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingTemplate")
			}
			return err
		}
	}

	return nil
}

func (m *AdditionalMessage) validateTextBody(formats strfmt.Registry) error {

	if err := validate.Required("textBody", "body", m.TextBody); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdditionalMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdditionalMessage) UnmarshalBinary(b []byte) error {
	var res AdditionalMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
