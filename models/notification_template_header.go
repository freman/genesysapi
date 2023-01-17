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

// NotificationTemplateHeader Template header object.
//
// swagger:model NotificationTemplateHeader
type NotificationTemplateHeader struct {

	// Media template header image.
	Media *ContentAttachment `json:"media,omitempty"`

	// Template parameters for placeholders in template.
	Parameters []*NotificationTemplateParameter `json:"parameters"`

	// Header text. For WhatsApp, ignored.
	Text string `json:"text,omitempty"`

	// Template header type.
	// Required: true
	// Enum: [Text Media]
	Type *string `json:"type"`
}

// Validate validates this notification template header
func (m *NotificationTemplateHeader) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMedia(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
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

func (m *NotificationTemplateHeader) validateMedia(formats strfmt.Registry) error {
	if swag.IsZero(m.Media) { // not required
		return nil
	}

	if m.Media != nil {
		if err := m.Media.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("media")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("media")
			}
			return err
		}
	}

	return nil
}

func (m *NotificationTemplateHeader) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var notificationTemplateHeaderTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Text","Media"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		notificationTemplateHeaderTypeTypePropEnum = append(notificationTemplateHeaderTypeTypePropEnum, v)
	}
}

const (

	// NotificationTemplateHeaderTypeText captures enum value "Text"
	NotificationTemplateHeaderTypeText string = "Text"

	// NotificationTemplateHeaderTypeMedia captures enum value "Media"
	NotificationTemplateHeaderTypeMedia string = "Media"
)

// prop value enum
func (m *NotificationTemplateHeader) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, notificationTemplateHeaderTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *NotificationTemplateHeader) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this notification template header based on the context it is used
func (m *NotificationTemplateHeader) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMedia(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationTemplateHeader) contextValidateMedia(ctx context.Context, formats strfmt.Registry) error {

	if m.Media != nil {
		if err := m.Media.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("media")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("media")
			}
			return err
		}
	}

	return nil
}

func (m *NotificationTemplateHeader) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Parameters); i++ {

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationTemplateHeader) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationTemplateHeader) UnmarshalBinary(b []byte) error {
	var res NotificationTemplateHeader
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
