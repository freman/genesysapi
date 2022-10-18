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

// ConversationMessageMetadataEvent Metadata information about a message event.
//
// swagger:model ConversationMessageMetadataEvent
type ConversationMessageMetadataEvent struct {

	// Type of this event element
	// Enum: [CoBrowse Typing Presence Unknown]
	EventType string `json:"eventType,omitempty"`

	// Event subtype
	// Enum: [On Join Offering OfferingExpired OfferingAccepted OfferingRejected Disconnect Unknown]
	SubType string `json:"subType,omitempty"`
}

// Validate validates this conversation message metadata event
func (m *ConversationMessageMetadataEvent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationMessageMetadataEventTypeEventTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CoBrowse","Typing","Presence","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationMessageMetadataEventTypeEventTypePropEnum = append(conversationMessageMetadataEventTypeEventTypePropEnum, v)
	}
}

const (

	// ConversationMessageMetadataEventEventTypeCoBrowse captures enum value "CoBrowse"
	ConversationMessageMetadataEventEventTypeCoBrowse string = "CoBrowse"

	// ConversationMessageMetadataEventEventTypeTyping captures enum value "Typing"
	ConversationMessageMetadataEventEventTypeTyping string = "Typing"

	// ConversationMessageMetadataEventEventTypePresence captures enum value "Presence"
	ConversationMessageMetadataEventEventTypePresence string = "Presence"

	// ConversationMessageMetadataEventEventTypeUnknown captures enum value "Unknown"
	ConversationMessageMetadataEventEventTypeUnknown string = "Unknown"
)

// prop value enum
func (m *ConversationMessageMetadataEvent) validateEventTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationMessageMetadataEventTypeEventTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationMessageMetadataEvent) validateEventType(formats strfmt.Registry) error {

	if swag.IsZero(m.EventType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventTypeEnum("eventType", "body", m.EventType); err != nil {
		return err
	}

	return nil
}

var conversationMessageMetadataEventTypeSubTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["On","Join","Offering","OfferingExpired","OfferingAccepted","OfferingRejected","Disconnect","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationMessageMetadataEventTypeSubTypePropEnum = append(conversationMessageMetadataEventTypeSubTypePropEnum, v)
	}
}

const (

	// ConversationMessageMetadataEventSubTypeOn captures enum value "On"
	ConversationMessageMetadataEventSubTypeOn string = "On"

	// ConversationMessageMetadataEventSubTypeJoin captures enum value "Join"
	ConversationMessageMetadataEventSubTypeJoin string = "Join"

	// ConversationMessageMetadataEventSubTypeOffering captures enum value "Offering"
	ConversationMessageMetadataEventSubTypeOffering string = "Offering"

	// ConversationMessageMetadataEventSubTypeOfferingExpired captures enum value "OfferingExpired"
	ConversationMessageMetadataEventSubTypeOfferingExpired string = "OfferingExpired"

	// ConversationMessageMetadataEventSubTypeOfferingAccepted captures enum value "OfferingAccepted"
	ConversationMessageMetadataEventSubTypeOfferingAccepted string = "OfferingAccepted"

	// ConversationMessageMetadataEventSubTypeOfferingRejected captures enum value "OfferingRejected"
	ConversationMessageMetadataEventSubTypeOfferingRejected string = "OfferingRejected"

	// ConversationMessageMetadataEventSubTypeDisconnect captures enum value "Disconnect"
	ConversationMessageMetadataEventSubTypeDisconnect string = "Disconnect"

	// ConversationMessageMetadataEventSubTypeUnknown captures enum value "Unknown"
	ConversationMessageMetadataEventSubTypeUnknown string = "Unknown"
)

// prop value enum
func (m *ConversationMessageMetadataEvent) validateSubTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationMessageMetadataEventTypeSubTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationMessageMetadataEvent) validateSubType(formats strfmt.Registry) error {

	if swag.IsZero(m.SubType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubTypeEnum("subType", "body", m.SubType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationMessageMetadataEvent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationMessageMetadataEvent) UnmarshalBinary(b []byte) error {
	var res ConversationMessageMetadataEvent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
