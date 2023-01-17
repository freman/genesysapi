// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CommunicationDispositionAppliedEvent communication disposition applied event
//
// swagger:model CommunicationDispositionAppliedEvent
type CommunicationDispositionAppliedEvent struct {

	// The wrapup-code (V4 UUID) used to disposition this interaction. If this value is not provided the disposition is considered skipped.
	Code string `json:"code,omitempty"`

	// A unique Id (V4 UUID) identifying this communication
	// Required: true
	CommunicationID *string `json:"communicationId"`

	// A unique Id (V4 UUID) identifying this conversation
	// Required: true
	ConversationID *string `json:"conversationId"`

	// A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	EventDateTime *strfmt.DateTime `json:"eventDateTime"`

	// A unique (V4 UUID) eventId for this event
	// Required: true
	EventID *string `json:"eventId"`

	// Text entered by the agent to describe the interaction or disposition. Ignored if the disposition is considered skipped.
	Notes string `json:"notes,omitempty"`

	// The list of tags selected by the agent to describe the interaction or disposition. Ignored if the disposition is considered skipped.
	Tags []string `json:"tags"`
}

// Validate validates this communication disposition applied event
func (m *CommunicationDispositionAppliedEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommunicationDispositionAppliedEvent) validateCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("communicationId", "body", m.CommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *CommunicationDispositionAppliedEvent) validateConversationID(formats strfmt.Registry) error {

	if err := validate.Required("conversationId", "body", m.ConversationID); err != nil {
		return err
	}

	return nil
}

func (m *CommunicationDispositionAppliedEvent) validateEventDateTime(formats strfmt.Registry) error {

	if err := validate.Required("eventDateTime", "body", m.EventDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventDateTime", "body", "date-time", m.EventDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CommunicationDispositionAppliedEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventId", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this communication disposition applied event based on context it is used
func (m *CommunicationDispositionAppliedEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CommunicationDispositionAppliedEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommunicationDispositionAppliedEvent) UnmarshalBinary(b []byte) error {
	var res CommunicationDispositionAppliedEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}