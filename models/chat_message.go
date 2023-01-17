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

// ChatMessage chat message
//
// swagger:model ChatMessage
type ChatMessage struct {

	// The message body
	Body string `json:"body,omitempty"`

	// Type of the message body (v2 chats only)
	// Enum: [STANDARD ACTIVITY TYPING NOTICE MEMBERJOIN MEMBERLEAVE MEDIAREQUEST]
	BodyType string `json:"bodyType,omitempty"`

	// The interaction id (if available)
	Chat string `json:"chat,omitempty"`

	// The message sender
	From string `json:"from,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// The message id
	Message string `json:"message,omitempty"`

	// Participant purpose of sender (v2 chats only)
	ParticipantPurpose string `json:"participantPurpose,omitempty"`

	// Communication of sender (v2 chats only)
	SenderCommunicationID string `json:"senderCommunicationId,omitempty"`

	// The message recipient
	To string `json:"to,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// The user information for the sender (if available)
	User *ChatMessageUser `json:"user,omitempty"`

	// utc
	Utc string `json:"utc,omitempty"`
}

// Validate validates this chat message
func (m *ChatMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBodyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var chatMessageTypeBodyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STANDARD","ACTIVITY","TYPING","NOTICE","MEMBERJOIN","MEMBERLEAVE","MEDIAREQUEST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		chatMessageTypeBodyTypePropEnum = append(chatMessageTypeBodyTypePropEnum, v)
	}
}

const (

	// ChatMessageBodyTypeSTANDARD captures enum value "STANDARD"
	ChatMessageBodyTypeSTANDARD string = "STANDARD"

	// ChatMessageBodyTypeACTIVITY captures enum value "ACTIVITY"
	ChatMessageBodyTypeACTIVITY string = "ACTIVITY"

	// ChatMessageBodyTypeTYPING captures enum value "TYPING"
	ChatMessageBodyTypeTYPING string = "TYPING"

	// ChatMessageBodyTypeNOTICE captures enum value "NOTICE"
	ChatMessageBodyTypeNOTICE string = "NOTICE"

	// ChatMessageBodyTypeMEMBERJOIN captures enum value "MEMBERJOIN"
	ChatMessageBodyTypeMEMBERJOIN string = "MEMBERJOIN"

	// ChatMessageBodyTypeMEMBERLEAVE captures enum value "MEMBERLEAVE"
	ChatMessageBodyTypeMEMBERLEAVE string = "MEMBERLEAVE"

	// ChatMessageBodyTypeMEDIAREQUEST captures enum value "MEDIAREQUEST"
	ChatMessageBodyTypeMEDIAREQUEST string = "MEDIAREQUEST"
)

// prop value enum
func (m *ChatMessage) validateBodyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, chatMessageTypeBodyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChatMessage) validateBodyType(formats strfmt.Registry) error {
	if swag.IsZero(m.BodyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBodyTypeEnum("bodyType", "body", m.BodyType); err != nil {
		return err
	}

	return nil
}

func (m *ChatMessage) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this chat message based on the context it is used
func (m *ChatMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChatMessage) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChatMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChatMessage) UnmarshalBinary(b []byte) error {
	var res ChatMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
