// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UserTransferEvent user transfer event
//
// swagger:model UserTransferEvent
type UserTransferEvent struct {

	// The id (V4 UUID) used by the external platform to refer to the transfer in subsequent Transfer events.
	// Required: true
	CommandID *string `json:"commandId"`

	// A unique Id (V4 UUID) identifying this conversation
	// Required: true
	ConversationID *string `json:"conversationId"`

	// The id (V4 UUID) of the desired destination user that the object communication should be transferred to.
	// Required: true
	DestinationUserID *string `json:"destinationUserId"`

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

	// The id (V4 UUID) of the communication that is being transferred away from. In many cases this will be the same as the `initiatingCommunicationId`.
	// Required: true
	TargetCommunicationID *string `json:"targetCommunicationId"`

	// Indicates the desired type of transfer.
	// Required: true
	// Enum: [Attended Unattended]
	TransferType *string `json:"transferType"`
}

// Validate validates this user transfer event
func (m *UserTransferEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationUserID(formats); err != nil {
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

	if err := m.validateTargetCommunicationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTransferType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserTransferEvent) validateCommandID(formats strfmt.Registry) error {

	if err := validate.Required("commandId", "body", m.CommandID); err != nil {
		return err
	}

	return nil
}

func (m *UserTransferEvent) validateConversationID(formats strfmt.Registry) error {

	if err := validate.Required("conversationId", "body", m.ConversationID); err != nil {
		return err
	}

	return nil
}

func (m *UserTransferEvent) validateDestinationUserID(formats strfmt.Registry) error {

	if err := validate.Required("destinationUserId", "body", m.DestinationUserID); err != nil {
		return err
	}

	return nil
}

func (m *UserTransferEvent) validateEventDateTime(formats strfmt.Registry) error {

	if err := validate.Required("eventDateTime", "body", m.EventDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventDateTime", "body", "date-time", m.EventDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserTransferEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventId", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

func (m *UserTransferEvent) validateInitiatingCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("initiatingCommunicationId", "body", m.InitiatingCommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *UserTransferEvent) validateObjectCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("objectCommunicationId", "body", m.ObjectCommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *UserTransferEvent) validateTargetCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("targetCommunicationId", "body", m.TargetCommunicationID); err != nil {
		return err
	}

	return nil
}

var userTransferEventTypeTransferTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attended","Unattended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userTransferEventTypeTransferTypePropEnum = append(userTransferEventTypeTransferTypePropEnum, v)
	}
}

const (

	// UserTransferEventTransferTypeAttended captures enum value "Attended"
	UserTransferEventTransferTypeAttended string = "Attended"

	// UserTransferEventTransferTypeUnattended captures enum value "Unattended"
	UserTransferEventTransferTypeUnattended string = "Unattended"
)

// prop value enum
func (m *UserTransferEvent) validateTransferTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userTransferEventTypeTransferTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserTransferEvent) validateTransferType(formats strfmt.Registry) error {

	if err := validate.Required("transferType", "body", m.TransferType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTransferTypeEnum("transferType", "body", *m.TransferType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this user transfer event based on context it is used
func (m *UserTransferEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserTransferEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserTransferEvent) UnmarshalBinary(b []byte) error {
	var res UserTransferEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}