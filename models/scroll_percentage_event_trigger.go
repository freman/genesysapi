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

// ScrollPercentageEventTrigger Details about a scroll percentage event trigger
//
// swagger:model ScrollPercentageEventTrigger
type ScrollPercentageEventTrigger struct {

	// Name of event triggered after scrolling to the specified percentage.
	// Required: true
	EventName *string `json:"eventName"`

	// Percentage of a webpage at which an event is triggered.
	// Required: true
	Percentage *int32 `json:"percentage"`
}

// Validate validates this scroll percentage event trigger
func (m *ScrollPercentageEventTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScrollPercentageEventTrigger) validateEventName(formats strfmt.Registry) error {

	if err := validate.Required("eventName", "body", m.EventName); err != nil {
		return err
	}

	return nil
}

func (m *ScrollPercentageEventTrigger) validatePercentage(formats strfmt.Registry) error {

	if err := validate.Required("percentage", "body", m.Percentage); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScrollPercentageEventTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScrollPercentageEventTrigger) UnmarshalBinary(b []byte) error {
	var res ScrollPercentageEventTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}