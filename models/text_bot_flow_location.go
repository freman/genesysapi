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

// TextBotFlowLocation Describes a flow location.
//
// swagger:model TextBotFlowLocation
type TextBotFlowLocation struct {

	// The name of the action that was active when the event of interest happened.
	// Required: true
	ActionName *string `json:"actionName"`

	// The number of the action that was active when the event of interest happened.
	// Required: true
	ActionNumber *int32 `json:"actionNumber"`

	// The name of the state or task which was active when the event of interest happened.
	// Required: true
	SequenceName *string `json:"sequenceName"`
}

// Validate validates this text bot flow location
func (m *TextBotFlowLocation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActionNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSequenceName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextBotFlowLocation) validateActionName(formats strfmt.Registry) error {

	if err := validate.Required("actionName", "body", m.ActionName); err != nil {
		return err
	}

	return nil
}

func (m *TextBotFlowLocation) validateActionNumber(formats strfmt.Registry) error {

	if err := validate.Required("actionNumber", "body", m.ActionNumber); err != nil {
		return err
	}

	return nil
}

func (m *TextBotFlowLocation) validateSequenceName(formats strfmt.Registry) error {

	if err := validate.Required("sequenceName", "body", m.SequenceName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextBotFlowLocation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextBotFlowLocation) UnmarshalBinary(b []byte) error {
	var res TextBotFlowLocation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
