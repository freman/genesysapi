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

// SourceConfiguration source configuration
//
// swagger:model SourceConfiguration
type SourceConfiguration struct {

	// The customer's unique external identifier associated with the conversation that comes from the external platform.
	// Required: true
	InteractionID *string `json:"interactionId"`

	// Identifies the external platform that is the source of the conversation.
	// Required: true
	SourceID *string `json:"sourceId"`

	// The customer's external identifier or tag associated with the conversation. If set, it will be used to tag the conversation.
	TagID string `json:"tagId,omitempty"`
}

// Validate validates this source configuration
func (m *SourceConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInteractionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSourceID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SourceConfiguration) validateInteractionID(formats strfmt.Registry) error {

	if err := validate.Required("interactionId", "body", m.InteractionID); err != nil {
		return err
	}

	return nil
}

func (m *SourceConfiguration) validateSourceID(formats strfmt.Registry) error {

	if err := validate.Required("sourceId", "body", m.SourceID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this source configuration based on context it is used
func (m *SourceConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SourceConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SourceConfiguration) UnmarshalBinary(b []byte) error {
	var res SourceConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
