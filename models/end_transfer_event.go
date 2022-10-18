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

// EndTransferEvent end transfer event
//
// swagger:model EndTransferEvent
type EndTransferEvent struct {

	// The id (V4 UUID) used to identify the transfer already started by the external platform.
	// Required: true
	CommandID *string `json:"commandId"`

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

	// Indicates whether the transfer completed successfully, was cancelled, or failed for some reason.
	// Required: true
	// Enum: [Completed Canceled Failed]
	FinalState *string `json:"finalState"`

	// The id (V4 UUID) of the communication that was being transferred.
	// Required: true
	ObjectCommunicationID *string `json:"objectCommunicationId"`
}

// Validate validates this end transfer event
func (m *EndTransferEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandID(formats); err != nil {
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

	if err := m.validateFinalState(formats); err != nil {
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

func (m *EndTransferEvent) validateCommandID(formats strfmt.Registry) error {

	if err := validate.Required("commandId", "body", m.CommandID); err != nil {
		return err
	}

	return nil
}

func (m *EndTransferEvent) validateConversationID(formats strfmt.Registry) error {

	if err := validate.Required("conversationId", "body", m.ConversationID); err != nil {
		return err
	}

	return nil
}

func (m *EndTransferEvent) validateEventDateTime(formats strfmt.Registry) error {

	if err := validate.Required("eventDateTime", "body", m.EventDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventDateTime", "body", "date-time", m.EventDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EndTransferEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventId", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

var endTransferEventTypeFinalStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Completed","Canceled","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endTransferEventTypeFinalStatePropEnum = append(endTransferEventTypeFinalStatePropEnum, v)
	}
}

const (

	// EndTransferEventFinalStateCompleted captures enum value "Completed"
	EndTransferEventFinalStateCompleted string = "Completed"

	// EndTransferEventFinalStateCanceled captures enum value "Canceled"
	EndTransferEventFinalStateCanceled string = "Canceled"

	// EndTransferEventFinalStateFailed captures enum value "Failed"
	EndTransferEventFinalStateFailed string = "Failed"
)

// prop value enum
func (m *EndTransferEvent) validateFinalStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, endTransferEventTypeFinalStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EndTransferEvent) validateFinalState(formats strfmt.Registry) error {

	if err := validate.Required("finalState", "body", m.FinalState); err != nil {
		return err
	}

	// value enum
	if err := m.validateFinalStateEnum("finalState", "body", *m.FinalState); err != nil {
		return err
	}

	return nil
}

func (m *EndTransferEvent) validateObjectCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("objectCommunicationId", "body", m.ObjectCommunicationID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndTransferEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndTransferEvent) UnmarshalBinary(b []byte) error {
	var res EndTransferEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
