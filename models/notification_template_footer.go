// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NotificationTemplateFooter Template footer object
//
// swagger:model NotificationTemplateFooter
type NotificationTemplateFooter struct {

	// Footer text. For WhatsApp, ignored
	Text string `json:"text,omitempty"`
}

// Validate validates this notification template footer
func (m *NotificationTemplateFooter) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationTemplateFooter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationTemplateFooter) UnmarshalBinary(b []byte) error {
	var res NotificationTemplateFooter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
