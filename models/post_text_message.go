// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostTextMessage post text message
//
// swagger:model PostTextMessage
type PostTextMessage struct {

	// A list of content elements in message
	Content []*ConversationMessageContent `json:"content"`

	// Message text. If type is structured, used as fallback for clients that do not support particular structured content
	Text string `json:"text,omitempty"`

	// Message type
	// Required: true
	// Enum: [Text Structured Receipt Event Message Unknown]
	Type *string `json:"type"`
}

// Validate validates this post text message
func (m *PostTextMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContent(formats); err != nil {
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

func (m *PostTextMessage) validateContent(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var postTextMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text","Structured","Receipt","Event","Message","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postTextMessageTypeTypePropEnum = append(postTextMessageTypeTypePropEnum, v)
	}
}

const (

	// PostTextMessageTypeText captures enum value "Text"
	PostTextMessageTypeText string = "Text"

	// PostTextMessageTypeStructured captures enum value "Structured"
	PostTextMessageTypeStructured string = "Structured"

	// PostTextMessageTypeReceipt captures enum value "Receipt"
	PostTextMessageTypeReceipt string = "Receipt"

	// PostTextMessageTypeEvent captures enum value "Event"
	PostTextMessageTypeEvent string = "Event"

	// PostTextMessageTypeMessage captures enum value "Message"
	PostTextMessageTypeMessage string = "Message"

	// PostTextMessageTypeUnknown captures enum value "Unknown"
	PostTextMessageTypeUnknown string = "Unknown"
)

// prop value enum
func (m *PostTextMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postTextMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostTextMessage) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post text message based on the context it is used
func (m *PostTextMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostTextMessage) contextValidateContent(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Content); i++ {

		if m.Content[i] != nil {
			if err := m.Content[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("content" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("content" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostTextMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTextMessage) UnmarshalBinary(b []byte) error {
	var res PostTextMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
