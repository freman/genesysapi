// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateCallbackResponse create callback response
//
// swagger:model CreateCallbackResponse
type CreateCallbackResponse struct {

	// The list of communication identifiers for the callback participants
	// Required: true
	CallbackIdentifiers []*CallbackIdentifier `json:"callbackIdentifiers"`

	// The conversation associated with the callback
	// Required: true
	Conversation *DomainEntityRef `json:"conversation"`
}

// Validate validates this create callback response
func (m *CreateCallbackResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallbackIdentifiers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversation(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCallbackResponse) validateCallbackIdentifiers(formats strfmt.Registry) error {

	if err := validate.Required("callbackIdentifiers", "body", m.CallbackIdentifiers); err != nil {
		return err
	}

	for i := 0; i < len(m.CallbackIdentifiers); i++ {
		if swag.IsZero(m.CallbackIdentifiers[i]) { // not required
			continue
		}

		if m.CallbackIdentifiers[i] != nil {
			if err := m.CallbackIdentifiers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("callbackIdentifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("callbackIdentifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateCallbackResponse) validateConversation(formats strfmt.Registry) error {

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

// ContextValidate validate this create callback response based on the context it is used
func (m *CreateCallbackResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCallbackIdentifiers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateCallbackResponse) contextValidateCallbackIdentifiers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CallbackIdentifiers); i++ {

		if m.CallbackIdentifiers[i] != nil {
			if err := m.CallbackIdentifiers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("callbackIdentifiers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("callbackIdentifiers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateCallbackResponse) contextValidateConversation(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CreateCallbackResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateCallbackResponse) UnmarshalBinary(b []byte) error {
	var res CreateCallbackResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
