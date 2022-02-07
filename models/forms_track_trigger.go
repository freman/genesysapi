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

// FormsTrackTrigger Details about a forms tracking event trigger
//
// swagger:model FormsTrackTrigger
type FormsTrackTrigger struct {

	// Whether to capture the form data in the form abandoned event.
	// Required: true
	CaptureDataOnFormAbandon *bool `json:"captureDataOnFormAbandon"`

	// Whether to capture the form data in the form submitted event.
	// Required: true
	CaptureDataOnFormSubmit *bool `json:"captureDataOnFormSubmit"`

	// Prefix for the form submitted or abandoned event name.
	// Required: true
	FormName *string `json:"formName"`

	// Form element that triggers the form submitted or abandoned event.
	// Required: true
	Selector *string `json:"selector"`
}

// Validate validates this forms track trigger
func (m *FormsTrackTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCaptureDataOnFormAbandon(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCaptureDataOnFormSubmit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFormName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FormsTrackTrigger) validateCaptureDataOnFormAbandon(formats strfmt.Registry) error {

	if err := validate.Required("captureDataOnFormAbandon", "body", m.CaptureDataOnFormAbandon); err != nil {
		return err
	}

	return nil
}

func (m *FormsTrackTrigger) validateCaptureDataOnFormSubmit(formats strfmt.Registry) error {

	if err := validate.Required("captureDataOnFormSubmit", "body", m.CaptureDataOnFormSubmit); err != nil {
		return err
	}

	return nil
}

func (m *FormsTrackTrigger) validateFormName(formats strfmt.Registry) error {

	if err := validate.Required("formName", "body", m.FormName); err != nil {
		return err
	}

	return nil
}

func (m *FormsTrackTrigger) validateSelector(formats strfmt.Registry) error {

	if err := validate.Required("selector", "body", m.Selector); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FormsTrackTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FormsTrackTrigger) UnmarshalBinary(b []byte) error {
	var res FormsTrackTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}