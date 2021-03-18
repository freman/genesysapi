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

// CallToAction call to action
//
// swagger:model CallToAction
type CallToAction struct {

	// Where the URL should be opened when the user clicks on the call to action button.
	// Enum: [Blank Self]
	Target string `json:"target,omitempty"`

	// Text displayed on the call to action button.
	Text string `json:"text,omitempty"`

	// URL to open when user clicks on the call to action button.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this call to action
func (m *CallToAction) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var callToActionTypeTargetPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Blank","Self"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callToActionTypeTargetPropEnum = append(callToActionTypeTargetPropEnum, v)
	}
}

const (

	// CallToActionTargetBlank captures enum value "Blank"
	CallToActionTargetBlank string = "Blank"

	// CallToActionTargetSelf captures enum value "Self"
	CallToActionTargetSelf string = "Self"
)

// prop value enum
func (m *CallToAction) validateTargetEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callToActionTypeTargetPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallToAction) validateTarget(formats strfmt.Registry) error {

	if swag.IsZero(m.Target) { // not required
		return nil
	}

	// value enum
	if err := m.validateTargetEnum("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

func (m *CallToAction) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallToAction) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallToAction) UnmarshalBinary(b []byte) error {
	var res CallToAction
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}