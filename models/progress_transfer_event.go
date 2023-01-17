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

// ProgressTransferEvent progress transfer event
//
// swagger:model ProgressTransferEvent
type ProgressTransferEvent struct {

	// The id (V4 UUID) used to identify the transfer already started by the external platform.
	// Required: true
	CommandID *string `json:"commandId"`

	// A unique Id (V4 UUID) identifying this conversation
	// Required: true
	ConversationID *string `json:"conversationId"`

	// The id (V4 UUID) of the communication that is being transferred to.
	// Required: true
	DestinationCommunicationID *string `json:"destinationCommunicationId"`

	// A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	EventDateTime *strfmt.DateTime `json:"eventDateTime"`

	// A unique (V4 UUID) eventId for this event
	// Required: true
	EventID *string `json:"eventId"`

	// The id (V4 UUID) of the communication that is being transferred.
	// Required: true
	ObjectCommunicationID *string `json:"objectCommunicationId"`
}

// Validate validates this progress transfer event
func (m *ProgressTransferEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationCommunicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectCommunicationID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProgressTransferEvent) validateCommandID(formats strfmt.Registry) error {

	if err := validate.Required("commandId", "body", m.CommandID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressTransferEvent) validateConversationID(formats strfmt.Registry) error {

	if err := validate.Required("conversationId", "body", m.ConversationID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressTransferEvent) validateDestinationCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("destinationCommunicationId", "body", m.DestinationCommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressTransferEvent) validateEventDateTime(formats strfmt.Registry) error {

	if err := validate.Required("eventDateTime", "body", m.EventDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventDateTime", "body", "date-time", m.EventDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProgressTransferEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventId", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressTransferEvent) validateObjectCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("objectCommunicationId", "body", m.ObjectCommunicationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this progress transfer event based on context it is used
func (m *ProgressTransferEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProgressTransferEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressTransferEvent) UnmarshalBinary(b []byte) error {
	var res ProgressTransferEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
