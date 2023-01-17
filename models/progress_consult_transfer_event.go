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

// ProgressConsultTransferEvent progress consult transfer event
//
// swagger:model ProgressConsultTransferEvent
type ProgressConsultTransferEvent struct {

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

	// The id (V4 UUID) of the communication representing the participant that is initiating the transfer.
	// Required: true
	InitiatingCommunicationID *string `json:"initiatingCommunicationId"`

	// The id (V4 UUID) of the communication that is being transferred.
	// Required: true
	ObjectCommunicationID *string `json:"objectCommunicationId"`
}

// Validate validates this progress consult transfer event
func (m *ProgressConsultTransferEvent) Validate(formats strfmt.Registry) error {
	var res []error

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

	if err := m.validateInitiatingCommunicationID(formats); err != nil {
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

func (m *ProgressConsultTransferEvent) validateConversationID(formats strfmt.Registry) error {

	if err := validate.Required("conversationId", "body", m.ConversationID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressConsultTransferEvent) validateDestinationCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("destinationCommunicationId", "body", m.DestinationCommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressConsultTransferEvent) validateEventDateTime(formats strfmt.Registry) error {

	if err := validate.Required("eventDateTime", "body", m.EventDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventDateTime", "body", "date-time", m.EventDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProgressConsultTransferEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventId", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressConsultTransferEvent) validateInitiatingCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("initiatingCommunicationId", "body", m.InitiatingCommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *ProgressConsultTransferEvent) validateObjectCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("objectCommunicationId", "body", m.ObjectCommunicationID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this progress consult transfer event based on context it is used
func (m *ProgressConsultTransferEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProgressConsultTransferEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProgressConsultTransferEvent) UnmarshalBinary(b []byte) error {
	var res ProgressConsultTransferEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
