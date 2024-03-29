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

// KnowledgeConversationContextResponse knowledge conversation context response
//
// swagger:model KnowledgeConversationContextResponse
type KnowledgeConversationContextResponse struct {

	// The conversation.
	// Required: true
	Conversation *AddressableEntityRef `json:"conversation"`

	// The end-user participant of the conversation.
	ExternalContact *AddressableEntityRef `json:"externalContact,omitempty"`

	// The media type of the conversation.
	// Enum: [Unknown Callback Chat Cobrowse Email Message Screenshare Video Voice]
	MediaType string `json:"mediaType,omitempty"`

	// The queue used to assign the interaction to the user.
	Queue *AddressableEntityRef `json:"queue,omitempty"`
}

// Validate validates this knowledge conversation context response
func (m *KnowledgeConversationContextResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeConversationContextResponse) validateConversation(formats strfmt.Registry) error {

	if err := validate.Required("conversation", "body", m.Conversation); err != nil {
		return err
	}

	if m.Conversation != nil {
		if err := m.Conversation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeConversationContextResponse) validateExternalContact(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalContact) { // not required
		return nil
	}

	if m.ExternalContact != nil {
		if err := m.ExternalContact.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalContact")
			}
			return err
		}
	}

	return nil
}

var knowledgeConversationContextResponseTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Callback","Chat","Cobrowse","Email","Message","Screenshare","Video","Voice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		knowledgeConversationContextResponseTypeMediaTypePropEnum = append(knowledgeConversationContextResponseTypeMediaTypePropEnum, v)
	}
}

const (

	// KnowledgeConversationContextResponseMediaTypeUnknown captures enum value "Unknown"
	KnowledgeConversationContextResponseMediaTypeUnknown string = "Unknown"

	// KnowledgeConversationContextResponseMediaTypeCallback captures enum value "Callback"
	KnowledgeConversationContextResponseMediaTypeCallback string = "Callback"

	// KnowledgeConversationContextResponseMediaTypeChat captures enum value "Chat"
	KnowledgeConversationContextResponseMediaTypeChat string = "Chat"

	// KnowledgeConversationContextResponseMediaTypeCobrowse captures enum value "Cobrowse"
	KnowledgeConversationContextResponseMediaTypeCobrowse string = "Cobrowse"

	// KnowledgeConversationContextResponseMediaTypeEmail captures enum value "Email"
	KnowledgeConversationContextResponseMediaTypeEmail string = "Email"

	// KnowledgeConversationContextResponseMediaTypeMessage captures enum value "Message"
	KnowledgeConversationContextResponseMediaTypeMessage string = "Message"

	// KnowledgeConversationContextResponseMediaTypeScreenshare captures enum value "Screenshare"
	KnowledgeConversationContextResponseMediaTypeScreenshare string = "Screenshare"

	// KnowledgeConversationContextResponseMediaTypeVideo captures enum value "Video"
	KnowledgeConversationContextResponseMediaTypeVideo string = "Video"

	// KnowledgeConversationContextResponseMediaTypeVoice captures enum value "Voice"
	KnowledgeConversationContextResponseMediaTypeVoice string = "Voice"
)

// prop value enum
func (m *KnowledgeConversationContextResponse) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, knowledgeConversationContextResponseTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *KnowledgeConversationContextResponse) validateMediaType(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

func (m *KnowledgeConversationContextResponse) validateQueue(formats strfmt.Registry) error {
	if swag.IsZero(m.Queue) { // not required
		return nil
	}

	if m.Queue != nil {
		if err := m.Queue.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this knowledge conversation context response based on the context it is used
func (m *KnowledgeConversationContextResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConversation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalContact(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeConversationContextResponse) contextValidateConversation(ctx context.Context, formats strfmt.Registry) error {

	if m.Conversation != nil {
		if err := m.Conversation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conversation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conversation")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeConversationContextResponse) contextValidateExternalContact(ctx context.Context, formats strfmt.Registry) error {

	if m.ExternalContact != nil {
		if err := m.ExternalContact.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("externalContact")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("externalContact")
			}
			return err
		}
	}

	return nil
}

func (m *KnowledgeConversationContextResponse) contextValidateQueue(ctx context.Context, formats strfmt.Registry) error {

	if m.Queue != nil {
		if err := m.Queue.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queue")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("queue")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeConversationContextResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeConversationContextResponse) UnmarshalBinary(b []byte) error {
	var res KnowledgeConversationContextResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
