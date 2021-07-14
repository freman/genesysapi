// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConversationNormalizedMessage General rich media message structure with normalized feature support across many messaging channels.
//
// swagger:model ConversationNormalizedMessage
type ConversationNormalizedMessage struct {

	// Channel-specific information that describes the message and the message channel/provider.
	// Read Only: true
	Channel *ConversationMessagingChannel `json:"channel,omitempty"`

	// List of content elements
	Content []*ConversationMessageContent `json:"content"`

	// The direction of the message.
	// Read Only: true
	// Enum: [Inbound Outbound]
	Direction string `json:"direction,omitempty"`

	// Unique ID of the message. Message receipts will have the same ID as the message they reference.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Indicates if this is the last message receipt for this message, or if another message receipt can be expected.
	// Read Only: true
	IsFinalReceipt *bool `json:"isFinalReceipt"`

	// Additional metadata about this message.
	Metadata map[string]string `json:"metadata,omitempty"`

	// Specifies if this message was sent by a human agent or bot. The platform may use this to apply appropriate provider policies.
	// Enum: [Human Bot]
	OriginatingEntity string `json:"originatingEntity,omitempty"`

	// List of reasons for a message receipt that indicates the message has failed. Only used with Failed status.
	// Read Only: true
	Reasons []*ConversationReason `json:"reasons"`

	// Message receipt status, only used with type Receipt.
	// Read Only: true
	// Enum: [Sent Delivered Read Failed Published Removed]
	Status string `json:"status,omitempty"`

	// Message text.
	Text string `json:"text,omitempty"`

	// Message type.
	// Required: true
	// Enum: [Text Structured Receipt]
	Type *string `json:"type"`
}

// Validate validates this conversation normalized message
func (m *ConversationNormalizedMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDirection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOriginatingEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReasons(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationNormalizedMessage) validateChannel(formats strfmt.Registry) error {

	if swag.IsZero(m.Channel) { // not required
		return nil
	}

	if m.Channel != nil {
		if err := m.Channel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("channel")
			}
			return err
		}
	}

	return nil
}

func (m *ConversationNormalizedMessage) validateContent(formats strfmt.Registry) error {

	if swag.IsZero(m.Content) { // not required
		return nil
	}

	for i := 0; i < len(m.Content); i++ {
		if swag.IsZero(m.Content[i]) { // not required
			continue
		}

		if m.Content[i] != nil {
			if err := m.Content[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var conversationNormalizedMessageTypeDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Inbound","Outbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationNormalizedMessageTypeDirectionPropEnum = append(conversationNormalizedMessageTypeDirectionPropEnum, v)
	}
}

const (

	// ConversationNormalizedMessageDirectionInbound captures enum value "Inbound"
	ConversationNormalizedMessageDirectionInbound string = "Inbound"

	// ConversationNormalizedMessageDirectionOutbound captures enum value "Outbound"
	ConversationNormalizedMessageDirectionOutbound string = "Outbound"
)

// prop value enum
func (m *ConversationNormalizedMessage) validateDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationNormalizedMessageTypeDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationNormalizedMessage) validateDirection(formats strfmt.Registry) error {

	if swag.IsZero(m.Direction) { // not required
		return nil
	}

	// value enum
	if err := m.validateDirectionEnum("direction", "body", m.Direction); err != nil {
		return err
	}

	return nil
}

var conversationNormalizedMessageTypeOriginatingEntityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Human","Bot"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationNormalizedMessageTypeOriginatingEntityPropEnum = append(conversationNormalizedMessageTypeOriginatingEntityPropEnum, v)
	}
}

const (

	// ConversationNormalizedMessageOriginatingEntityHuman captures enum value "Human"
	ConversationNormalizedMessageOriginatingEntityHuman string = "Human"

	// ConversationNormalizedMessageOriginatingEntityBot captures enum value "Bot"
	ConversationNormalizedMessageOriginatingEntityBot string = "Bot"
)

// prop value enum
func (m *ConversationNormalizedMessage) validateOriginatingEntityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationNormalizedMessageTypeOriginatingEntityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationNormalizedMessage) validateOriginatingEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.OriginatingEntity) { // not required
		return nil
	}

	// value enum
	if err := m.validateOriginatingEntityEnum("originatingEntity", "body", m.OriginatingEntity); err != nil {
		return err
	}

	return nil
}

func (m *ConversationNormalizedMessage) validateReasons(formats strfmt.Registry) error {

	if swag.IsZero(m.Reasons) { // not required
		return nil
	}

	for i := 0; i < len(m.Reasons); i++ {
		if swag.IsZero(m.Reasons[i]) { // not required
			continue
		}

		if m.Reasons[i] != nil {
			if err := m.Reasons[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reasons" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var conversationNormalizedMessageTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sent","Delivered","Read","Failed","Published","Removed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationNormalizedMessageTypeStatusPropEnum = append(conversationNormalizedMessageTypeStatusPropEnum, v)
	}
}

const (

	// ConversationNormalizedMessageStatusSent captures enum value "Sent"
	ConversationNormalizedMessageStatusSent string = "Sent"

	// ConversationNormalizedMessageStatusDelivered captures enum value "Delivered"
	ConversationNormalizedMessageStatusDelivered string = "Delivered"

	// ConversationNormalizedMessageStatusRead captures enum value "Read"
	ConversationNormalizedMessageStatusRead string = "Read"

	// ConversationNormalizedMessageStatusFailed captures enum value "Failed"
	ConversationNormalizedMessageStatusFailed string = "Failed"

	// ConversationNormalizedMessageStatusPublished captures enum value "Published"
	ConversationNormalizedMessageStatusPublished string = "Published"

	// ConversationNormalizedMessageStatusRemoved captures enum value "Removed"
	ConversationNormalizedMessageStatusRemoved string = "Removed"
)

// prop value enum
func (m *ConversationNormalizedMessage) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationNormalizedMessageTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationNormalizedMessage) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var conversationNormalizedMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text","Structured","Receipt"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationNormalizedMessageTypeTypePropEnum = append(conversationNormalizedMessageTypeTypePropEnum, v)
	}
}

const (

	// ConversationNormalizedMessageTypeText captures enum value "Text"
	ConversationNormalizedMessageTypeText string = "Text"

	// ConversationNormalizedMessageTypeStructured captures enum value "Structured"
	ConversationNormalizedMessageTypeStructured string = "Structured"

	// ConversationNormalizedMessageTypeReceipt captures enum value "Receipt"
	ConversationNormalizedMessageTypeReceipt string = "Receipt"
)

// prop value enum
func (m *ConversationNormalizedMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationNormalizedMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationNormalizedMessage) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationNormalizedMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationNormalizedMessage) UnmarshalBinary(b []byte) error {
	var res ConversationNormalizedMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
