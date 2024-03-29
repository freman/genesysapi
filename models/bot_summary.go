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

// BotSummary A summary description for a botConnector bot
//
// swagger:model BotSummary
type BotSummary struct {

	// A system-generated string that contains metadata about this bot.
	// Read Only: true
	BotCompositeTag string `json:"botCompositeTag,omitempty"`

	// An optional description of the bot.
	Description string `json:"description,omitempty"`

	// The id of the bot.
	// Required: true
	ID *string `json:"id"`

	// The name of the bot.
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this bot summary
func (m *BotSummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BotSummary) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *BotSummary) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bot summary based on the context it is used
func (m *BotSummary) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBotCompositeTag(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BotSummary) contextValidateBotCompositeTag(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "botCompositeTag", "body", string(m.BotCompositeTag)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BotSummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BotSummary) UnmarshalBinary(b []byte) error {
	var res BotSummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
