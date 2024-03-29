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

// RoutingTransferEvent routing transfer event
//
// swagger:model RoutingTransferEvent
type RoutingTransferEvent struct {

	// The id (V4 UUID) used by the external platform to refer to the transfer in subsequent *Transfer events.
	// Required: true
	CommandID *string `json:"commandId"`

	// A unique Id (V4 UUID) identifying this conversation
	// Required: true
	ConversationID *string `json:"conversationId"`

	// The id (V4 UUID) of the desired destination queue that the object communication should be transferred to.
	// Required: true
	DestinationQueueID *string `json:"destinationQueueId"`

	// A Date Time representing the time this event occurred. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	EventDateTime *strfmt.DateTime `json:"eventDateTime"`

	// A unique (V4 UUID) eventId for this event
	// Required: true
	EventID *string `json:"eventId"`

	// Indicates the desired type of transfer.
	// Required: true
	InitiatingCommunicationID *string `json:"initiatingCommunicationId"`

	// The unique identifier (V4 UUID) for the language that should be used to determine the destination for the conversation.
	LanguageID string `json:"languageId,omitempty"`

	// The id (V4 UUID) of the communication that is being transferred.
	// Required: true
	ObjectCommunicationID *string `json:"objectCommunicationId"`

	// The unique identifiers (V4 UUID) for the skills that should be used to determine the destination for the conversation.
	SkillIds []string `json:"skillIds"`

	// The id (V4 UUID) of the communication that is being transferred away from. In many cases this will be the same as the `initiatingCommunicationId`.
	// Required: true
	TargetCommunicationID *string `json:"targetCommunicationId"`

	// Indicates the desired type of transfer.
	// Required: true
	// Enum: [Attended Unattended]
	TransferType *string `json:"transferType"`
}

// Validate validates this routing transfer event
func (m *RoutingTransferEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommandID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestinationQueueID(formats); err != nil {
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

func (m *RoutingTransferEvent) validateCommandID(formats strfmt.Registry) error {

	if err := validate.Required("commandId", "body", m.CommandID); err != nil {
		return err
	}

	return nil
}

func (m *RoutingTransferEvent) validateConversationID(formats strfmt.Registry) error {

	if err := validate.Required("conversationId", "body", m.ConversationID); err != nil {
		return err
	}

	return nil
}

func (m *RoutingTransferEvent) validateDestinationQueueID(formats strfmt.Registry) error {

	if err := validate.Required("destinationQueueId", "body", m.DestinationQueueID); err != nil {
		return err
	}

	return nil
}

func (m *RoutingTransferEvent) validateEventDateTime(formats strfmt.Registry) error {

	if err := validate.Required("eventDateTime", "body", m.EventDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventDateTime", "body", "date-time", m.EventDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *RoutingTransferEvent) validateEventID(formats strfmt.Registry) error {

	if err := validate.Required("eventId", "body", m.EventID); err != nil {
		return err
	}

	return nil
}

func (m *RoutingTransferEvent) validateInitiatingCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("initiatingCommunicationId", "body", m.InitiatingCommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *RoutingTransferEvent) validateObjectCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("objectCommunicationId", "body", m.ObjectCommunicationID); err != nil {
		return err
	}

	return nil
}

func (m *RoutingTransferEvent) validateTargetCommunicationID(formats strfmt.Registry) error {

	if err := validate.Required("targetCommunicationId", "body", m.TargetCommunicationID); err != nil {
		return err
	}

	return nil
}

var routingTransferEventTypeTransferTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attended","Unattended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		routingTransferEventTypeTransferTypePropEnum = append(routingTransferEventTypeTransferTypePropEnum, v)
	}
}

const (

	// RoutingTransferEventTransferTypeAttended captures enum value "Attended"
	RoutingTransferEventTransferTypeAttended string = "Attended"

	// RoutingTransferEventTransferTypeUnattended captures enum value "Unattended"
	RoutingTransferEventTransferTypeUnattended string = "Unattended"
)

// prop value enum
func (m *RoutingTransferEvent) validateTransferTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, routingTransferEventTypeTransferTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RoutingTransferEvent) validateTransferType(formats strfmt.Registry) error {

	if err := validate.Required("transferType", "body", m.TransferType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTransferTypeEnum("transferType", "body", *m.TransferType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this routing transfer event based on context it is used
func (m *RoutingTransferEvent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RoutingTransferEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RoutingTransferEvent) UnmarshalBinary(b []byte) error {
	var res RoutingTransferEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
