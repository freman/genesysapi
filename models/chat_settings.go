// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ChatSettings chat settings
//
// swagger:model ChatSettings
type ChatSettings struct {

	// Retention time for messages in days
	MessageRetentionPeriodDays int32 `json:"messageRetentionPeriodDays,omitempty"`
}

// Validate validates this chat settings
func (m *ChatSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this chat settings based on context it is used
func (m *ChatSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ChatSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChatSettings) UnmarshalBinary(b []byte) error {
	var res ChatSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
