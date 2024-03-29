// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConversationContentNotificationTemplate Template notification object.
//
// swagger:model ConversationContentNotificationTemplate
type ConversationContentNotificationTemplate struct {

	// The template body.
	// Required: true
	Body *ConversationNotificationTemplateBody `json:"body"`

	// The template footer.
	Footer *ConversationNotificationTemplateFooter `json:"footer,omitempty"`

	// The template header.
	Header *ConversationNotificationTemplateHeader `json:"header,omitempty"`

	// The messaging provider template ID. For WhatsApp, 'namespace@name'.
	ID string `json:"id,omitempty"`

	// Template language.
	Language string `json:"language,omitempty"`
}

// Validate validates this conversation content notification template
func (m *ConversationContentNotificationTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBody(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFooter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationContentNotificationTemplate) validateBody(formats strfmt.Registry) error {

	if err := validate.Required("body", "body", m.Body); err != nil {
		return err
	}

	if m.Body != nil {
		if err := m.Body.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

func (m *ConversationContentNotificationTemplate) validateFooter(formats strfmt.Registry) error {
	if swag.IsZero(m.Footer) { // not required
		return nil
	}

	if m.Footer != nil {
		if err := m.Footer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("footer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("footer")
			}
			return err
		}
	}

	return nil
}

func (m *ConversationContentNotificationTemplate) validateHeader(formats strfmt.Registry) error {
	if swag.IsZero(m.Header) { // not required
		return nil
	}

	if m.Header != nil {
		if err := m.Header.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this conversation content notification template based on the context it is used
func (m *ConversationContentNotificationTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBody(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFooter(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHeader(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConversationContentNotificationTemplate) contextValidateBody(ctx context.Context, formats strfmt.Registry) error {

	if m.Body != nil {
		if err := m.Body.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("body")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("body")
			}
			return err
		}
	}

	return nil
}

func (m *ConversationContentNotificationTemplate) contextValidateFooter(ctx context.Context, formats strfmt.Registry) error {

	if m.Footer != nil {
		if err := m.Footer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("footer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("footer")
			}
			return err
		}
	}

	return nil
}

func (m *ConversationContentNotificationTemplate) contextValidateHeader(ctx context.Context, formats strfmt.Registry) error {

	if m.Header != nil {
		if err := m.Header.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("header")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("header")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConversationContentNotificationTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConversationContentNotificationTemplate) UnmarshalBinary(b []byte) error {
	var res ConversationContentNotificationTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
