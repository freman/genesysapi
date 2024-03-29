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

// MediaPolicies media policies
//
// swagger:model MediaPolicies
type MediaPolicies struct {

	// Conditions and actions for calls
	CallPolicy *CallMediaPolicy `json:"callPolicy,omitempty"`

	// Conditions and actions for chats
	ChatPolicy *ChatMediaPolicy `json:"chatPolicy,omitempty"`

	// Conditions and actions for emails
	EmailPolicy *EmailMediaPolicy `json:"emailPolicy,omitempty"`

	// Conditions and actions for messages
	MessagePolicy *MessageMediaPolicy `json:"messagePolicy,omitempty"`
}

// Validate validates this media policies
func (m *MediaPolicies) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChatPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailPolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessagePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaPolicies) validateCallPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.CallPolicy) { // not required
		return nil
	}

	if m.CallPolicy != nil {
		if err := m.CallPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *MediaPolicies) validateChatPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.ChatPolicy) { // not required
		return nil
	}

	if m.ChatPolicy != nil {
		if err := m.ChatPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chatPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chatPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *MediaPolicies) validateEmailPolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.EmailPolicy) { // not required
		return nil
	}

	if m.EmailPolicy != nil {
		if err := m.EmailPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *MediaPolicies) validateMessagePolicy(formats strfmt.Registry) error {
	if swag.IsZero(m.MessagePolicy) { // not required
		return nil
	}

	if m.MessagePolicy != nil {
		if err := m.MessagePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("messagePolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this media policies based on the context it is used
func (m *MediaPolicies) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCallPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChatPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEmailPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessagePolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MediaPolicies) contextValidateCallPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.CallPolicy != nil {
		if err := m.CallPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *MediaPolicies) contextValidateChatPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.ChatPolicy != nil {
		if err := m.ChatPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("chatPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("chatPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *MediaPolicies) contextValidateEmailPolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.EmailPolicy != nil {
		if err := m.EmailPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emailPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("emailPolicy")
			}
			return err
		}
	}

	return nil
}

func (m *MediaPolicies) contextValidateMessagePolicy(ctx context.Context, formats strfmt.Registry) error {

	if m.MessagePolicy != nil {
		if err := m.MessagePolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagePolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("messagePolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MediaPolicies) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MediaPolicies) UnmarshalBinary(b []byte) error {
	var res MediaPolicies
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
