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

// OpenMessageContent Message content element.
//
// swagger:model OpenMessageContent
type OpenMessageContent struct {

	// Attachment content.
	Attachment *ConversationContentAttachment `json:"attachment,omitempty"`

	// Type of this content element. If contentType = "Attachment" only one item is allowed.
	// Required: true
	// Enum: [Attachment]
	ContentType *string `json:"contentType"`
}

// Validate validates this open message content
func (m *OpenMessageContent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttachment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenMessageContent) validateAttachment(formats strfmt.Registry) error {
	if swag.IsZero(m.Attachment) { // not required
		return nil
	}

	if m.Attachment != nil {
		if err := m.Attachment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment")
			}
			return err
		}
	}

	return nil
}

var openMessageContentTypeContentTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attachment"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		openMessageContentTypeContentTypePropEnum = append(openMessageContentTypeContentTypePropEnum, v)
	}
}

const (

	// OpenMessageContentContentTypeAttachment captures enum value "Attachment"
	OpenMessageContentContentTypeAttachment string = "Attachment"
)

// prop value enum
func (m *OpenMessageContent) validateContentTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, openMessageContentTypeContentTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *OpenMessageContent) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("contentType", "body", m.ContentType); err != nil {
		return err
	}

	// value enum
	if err := m.validateContentTypeEnum("contentType", "body", *m.ContentType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this open message content based on the context it is used
func (m *OpenMessageContent) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAttachment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenMessageContent) contextValidateAttachment(ctx context.Context, formats strfmt.Registry) error {

	if m.Attachment != nil {
		if err := m.Attachment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("attachment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("attachment")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenMessageContent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenMessageContent) UnmarshalBinary(b []byte) error {
	var res OpenMessageContent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
