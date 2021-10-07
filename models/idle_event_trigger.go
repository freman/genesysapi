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

// IdleEventTrigger Details about an idle event trigger
//
// swagger:model IdleEventTrigger
type IdleEventTrigger struct {

	// Name of event triggered after period of inactivity.
	// Required: true
	EventName *string `json:"eventName"`

	// Number of seconds of inactivity before an event is triggered.
	IdleAfterSeconds int64 `json:"idleAfterSeconds,omitempty"`
}

// Validate validates this idle event trigger
func (m *IdleEventTrigger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdleEventTrigger) validateEventName(formats strfmt.Registry) error {

	if err := validate.Required("eventName", "body", m.EventName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdleEventTrigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdleEventTrigger) UnmarshalBinary(b []byte) error {
	var res IdleEventTrigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
