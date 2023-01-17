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

// WebChatRoutingTarget web chat routing target
//
// swagger:model WebChatRoutingTarget
type WebChatRoutingTarget struct {

	// The language name to use for routing.
	Language string `json:"language,omitempty"`

	// The priority to assign to the conversation for routing.
	Priority int64 `json:"priority,omitempty"`

	// The list of skill names to use for routing.
	Skills []string `json:"skills"`

	// The target of the route, in the format appropriate given the 'targetType'.
	// Required: true
	TargetAddress *string `json:"targetAddress"`

	// The target type of the routing target, such as 'QUEUE'.
	// Required: true
	// Enum: [QUEUE]
	TargetType *string `json:"targetType"`
}

// Validate validates this web chat routing target
func (m *WebChatRoutingTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTargetAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebChatRoutingTarget) validateTargetAddress(formats strfmt.Registry) error {

	if err := validate.Required("targetAddress", "body", m.TargetAddress); err != nil {
		return err
	}

	return nil
}

var webChatRoutingTargetTypeTargetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QUEUE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webChatRoutingTargetTypeTargetTypePropEnum = append(webChatRoutingTargetTypeTargetTypePropEnum, v)
	}
}

const (

	// WebChatRoutingTargetTargetTypeQUEUE captures enum value "QUEUE"
	WebChatRoutingTargetTargetTypeQUEUE string = "QUEUE"
)

// prop value enum
func (m *WebChatRoutingTarget) validateTargetTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webChatRoutingTargetTypeTargetTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebChatRoutingTarget) validateTargetType(formats strfmt.Registry) error {

	if err := validate.Required("targetType", "body", m.TargetType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTargetTypeEnum("targetType", "body", *m.TargetType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web chat routing target based on context it is used
func (m *WebChatRoutingTarget) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebChatRoutingTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebChatRoutingTarget) UnmarshalBinary(b []byte) error {
	var res WebChatRoutingTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
