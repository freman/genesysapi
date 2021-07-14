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

// NotificationTemplateParameter Template parameters for placeholders in template.
//
// swagger:model NotificationTemplateParameter
type NotificationTemplateParameter struct {

	// Parameter name.
	Name string `json:"name,omitempty"`

	// Parameter text value.
	// Required: true
	Text *string `json:"text"`
}

// Validate validates this notification template parameter
func (m *NotificationTemplateParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateText(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationTemplateParameter) validateText(formats strfmt.Registry) error {

	if err := validate.Required("text", "body", m.Text); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationTemplateParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationTemplateParameter) UnmarshalBinary(b []byte) error {
	var res NotificationTemplateParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
