// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StorySetting story setting
//
// swagger:model StorySetting
type StorySetting struct {

	// Setting relating to Story Mentions
	Mention *InboundOnlySetting `json:"mention,omitempty"`

	// Setting relating to Story Replies
	Reply *InboundOnlySetting `json:"reply,omitempty"`
}

// Validate validates this story setting
func (m *StorySetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMention(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReply(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorySetting) validateMention(formats strfmt.Registry) error {
	if swag.IsZero(m.Mention) { // not required
		return nil
	}

	if m.Mention != nil {
		if err := m.Mention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mention")
			}
			return err
		}
	}

	return nil
}

func (m *StorySetting) validateReply(formats strfmt.Registry) error {
	if swag.IsZero(m.Reply) { // not required
		return nil
	}

	if m.Reply != nil {
		if err := m.Reply.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reply")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this story setting based on the context it is used
func (m *StorySetting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReply(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StorySetting) contextValidateMention(ctx context.Context, formats strfmt.Registry) error {

	if m.Mention != nil {
		if err := m.Mention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mention")
			}
			return err
		}
	}

	return nil
}

func (m *StorySetting) contextValidateReply(ctx context.Context, formats strfmt.Registry) error {

	if m.Reply != nil {
		if err := m.Reply.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reply")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reply")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StorySetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StorySetting) UnmarshalBinary(b []byte) error {
	var res StorySetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
