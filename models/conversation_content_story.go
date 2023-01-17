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

// ConversationContentStory An ephemeral story.
//
// swagger:model ConversationContentStory
type ConversationContentStory struct {

	// ID of the ephemeral story being replied to.
	ReplyToID string `json:"replyToId,omitempty"`

	// Type of ephemeral story attachment.
	// Enum: [Mention Reply]
	Type string `json:"type,omitempty"`

	// URL to the ephemeral story.
	URL string `json:"url,omitempty"`
}

// Validate validates this conversation content story
func (m *ConversationContentStory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var conversationContentStoryTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Mention","Reply"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationContentStoryTypeTypePropEnum = append(conversationContentStoryTypeTypePropEnum, v)
	}
}

const (

	// ConversationContentStoryTypeMention captures enum value "Mention"
	ConversationContentStoryTypeMention string = "Mention"

	// ConversationContentStoryTypeReply captures enum value "Reply"
	ConversationContentStoryTypeReply string = "Reply"
)

// prop value enum
func (m *ConversationContentStory) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationContentStoryTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationContentStory) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conversation content story based on context it is used
func (m *ConversationContentStory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationContentStory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationContentStory) UnmarshalBinary(b []byte) error {
	var res ConversationContentStory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
