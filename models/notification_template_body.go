// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NotificationTemplateBody Template body object.
//
// swagger:model NotificationTemplateBody
type NotificationTemplateBody struct {

	// Template parameters for placeholders in template.
	// Required: true
	Parameters []*NotificationTemplateParameter `json:"parameters"`

	// Body text. For WhatsApp, ignored.
	Text string `json:"text,omitempty"`
}

// Validate validates this notification template body
func (m *NotificationTemplateBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationTemplateBody) validateParameters(formats strfmt.Registry) error {

	if err := validate.Required("parameters", "body", m.Parameters); err != nil {
		return err
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationTemplateBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationTemplateBody) UnmarshalBinary(b []byte) error {
	var res NotificationTemplateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
