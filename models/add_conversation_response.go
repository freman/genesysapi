// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AddConversationResponse add conversation response
//
// swagger:model AddConversationResponse
type AddConversationResponse struct {

	// The appointment reference
	// Read Only: true
	Appointment *CoachingAppointmentReference `json:"appointment,omitempty"`

	// The conversation reference
	// Read Only: true
	Conversation *ConversationReference `json:"conversation,omitempty"`
}

// Validate validates this add conversation response
func (m *AddConversationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppointment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddConversationResponse) validateAppointment(formats strfmt.Registry) error {

	if swag.IsZero(m.Appointment) { // not required
		return nil
	}

	if m.Appointment != nil {
		if err := m.Appointment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("appointment")
			}
			return err
		}
	}

	return nil
}

func (m *AddConversationResponse) validateConversation(formats strfmt.Registry) error {

	if swag.IsZero(m.Conversation) { // not required
		return nil
	}

	if m.Conversation != nil {
		if err := m.Conversation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddConversationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddConversationResponse) UnmarshalBinary(b []byte) error {
	var res AddConversationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
