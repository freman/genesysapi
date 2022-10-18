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

// SetContentTemplateActionSettings set content template action settings
//
// swagger:model SetContentTemplateActionSettings
type SetContentTemplateActionSettings struct {

	// A string of email contentTemplateId.
	// Required: true
	EmailContentTemplateID *string `json:"emailContentTemplateId"`

	// A string of sms contentTemplateId.
	// Required: true
	SmsContentTemplateID *string `json:"smsContentTemplateId"`
}

// Validate validates this set content template action settings
func (m *SetContentTemplateActionSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEmailContentTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmsContentTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SetContentTemplateActionSettings) validateEmailContentTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("emailContentTemplateId", "body", m.EmailContentTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *SetContentTemplateActionSettings) validateSmsContentTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("smsContentTemplateId", "body", m.SmsContentTemplateID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SetContentTemplateActionSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SetContentTemplateActionSettings) UnmarshalBinary(b []byte) error {
	var res SetContentTemplateActionSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
