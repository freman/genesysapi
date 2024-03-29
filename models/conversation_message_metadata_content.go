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

// ConversationMessageMetadataContent Metadata information about a message content.
//
// swagger:model ConversationMessageMetadataContent
type ConversationMessageMetadataContent struct {

	// Type of this content element.
	// Enum: [Attachment Location QuickReply Notification ButtonResponse Story Mention Card Carousel Text QuickReplyV2 Unknown]
	ContentType string `json:"contentType,omitempty"`

	// Content subtype
	// Enum: [Image Video Audio File Link Mention Reply Button QuickReply Postback Unknown]
	SubType string `json:"subType,omitempty"`
}

// Validate validates this conversation message metadata content
func (m *ConversationMessageMetadataContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentType(formats); err != nil {
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

var conversationMessageMetadataContentTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attachment","Location","QuickReply","Notification","ButtonResponse","Story","Mention","Card","Carousel","Text","QuickReplyV2","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationMessageMetadataContentTypeContentTypePropEnum = append(conversationMessageMetadataContentTypeContentTypePropEnum, v)
	}
}

const (

	// ConversationMessageMetadataContentContentTypeAttachment captures enum value "Attachment"
	ConversationMessageMetadataContentContentTypeAttachment string = "Attachment"

	// ConversationMessageMetadataContentContentTypeLocation captures enum value "Location"
	ConversationMessageMetadataContentContentTypeLocation string = "Location"

	// ConversationMessageMetadataContentContentTypeQuickReply captures enum value "QuickReply"
	ConversationMessageMetadataContentContentTypeQuickReply string = "QuickReply"

	// ConversationMessageMetadataContentContentTypeNotification captures enum value "Notification"
	ConversationMessageMetadataContentContentTypeNotification string = "Notification"

	// ConversationMessageMetadataContentContentTypeButtonResponse captures enum value "ButtonResponse"
	ConversationMessageMetadataContentContentTypeButtonResponse string = "ButtonResponse"

	// ConversationMessageMetadataContentContentTypeStory captures enum value "Story"
	ConversationMessageMetadataContentContentTypeStory string = "Story"

	// ConversationMessageMetadataContentContentTypeMention captures enum value "Mention"
	ConversationMessageMetadataContentContentTypeMention string = "Mention"

	// ConversationMessageMetadataContentContentTypeCard captures enum value "Card"
	ConversationMessageMetadataContentContentTypeCard string = "Card"

	// ConversationMessageMetadataContentContentTypeCarousel captures enum value "Carousel"
	ConversationMessageMetadataContentContentTypeCarousel string = "Carousel"

	// ConversationMessageMetadataContentContentTypeText captures enum value "Text"
	ConversationMessageMetadataContentContentTypeText string = "Text"

	// ConversationMessageMetadataContentContentTypeQuickReplyV2 captures enum value "QuickReplyV2"
	ConversationMessageMetadataContentContentTypeQuickReplyV2 string = "QuickReplyV2"

	// ConversationMessageMetadataContentContentTypeUnknown captures enum value "Unknown"
	ConversationMessageMetadataContentContentTypeUnknown string = "Unknown"
)

// prop value enum
func (m *ConversationMessageMetadataContent) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationMessageMetadataContentTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationMessageMetadataContent) validateContentType(formats strfmt.Registry) error {
	if swag.IsZero(m.ContentType) { // not required
		return nil
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

var conversationMessageMetadataContentTypeSubTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Image","Video","Audio","File","Link","Mention","Reply","Button","QuickReply","Postback","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationMessageMetadataContentTypeSubTypePropEnum = append(conversationMessageMetadataContentTypeSubTypePropEnum, v)
	}
}

const (

	// ConversationMessageMetadataContentSubTypeImage captures enum value "Image"
	ConversationMessageMetadataContentSubTypeImage string = "Image"

	// ConversationMessageMetadataContentSubTypeVideo captures enum value "Video"
	ConversationMessageMetadataContentSubTypeVideo string = "Video"

	// ConversationMessageMetadataContentSubTypeAudio captures enum value "Audio"
	ConversationMessageMetadataContentSubTypeAudio string = "Audio"

	// ConversationMessageMetadataContentSubTypeFile captures enum value "File"
	ConversationMessageMetadataContentSubTypeFile string = "File"

	// ConversationMessageMetadataContentSubTypeLink captures enum value "Link"
	ConversationMessageMetadataContentSubTypeLink string = "Link"

	// ConversationMessageMetadataContentSubTypeMention captures enum value "Mention"
	ConversationMessageMetadataContentSubTypeMention string = "Mention"

	// ConversationMessageMetadataContentSubTypeReply captures enum value "Reply"
	ConversationMessageMetadataContentSubTypeReply string = "Reply"

	// ConversationMessageMetadataContentSubTypeButton captures enum value "Button"
	ConversationMessageMetadataContentSubTypeButton string = "Button"

	// ConversationMessageMetadataContentSubTypeQuickReply captures enum value "QuickReply"
	ConversationMessageMetadataContentSubTypeQuickReply string = "QuickReply"

	// ConversationMessageMetadataContentSubTypePostback captures enum value "Postback"
	ConversationMessageMetadataContentSubTypePostback string = "Postback"

	// ConversationMessageMetadataContentSubTypeUnknown captures enum value "Unknown"
	ConversationMessageMetadataContentSubTypeUnknown string = "Unknown"
)

// prop value enum
func (m *ConversationMessageMetadataContent) validateSubTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationMessageMetadataContentTypeSubTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConversationMessageMetadataContent) validateSubType(formats strfmt.Registry) error {
	if swag.IsZero(m.SubType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSubTypeEnum("subType", "body", m.SubType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this conversation message metadata content based on context it is used
func (m *ConversationMessageMetadataContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConversationMessageMetadataContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationMessageMetadataContent) UnmarshalBinary(b []byte) error {
	var res ConversationMessageMetadataContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
