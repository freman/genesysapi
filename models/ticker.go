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

// Ticker ticker
//
// swagger:model Ticker
type Ticker struct {

	// The exchange for this ticker symbol. Examples: NYSE, FTSE, NASDAQ, etc.
	// Required: true
	Exchange *string `json:"exchange"`

	// The ticker symbol for this organization. Example: ININ, AAPL, MSFT, etc.
	// Required: true
	Symbol *string `json:"symbol"`
}

// Validate validates this ticker
func (m *Ticker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExchange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSymbol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Ticker) validateExchange(formats strfmt.Registry) error {

	if err := validate.Required("exchange", "body", m.Exchange); err != nil {
		return err
	}

	return nil
}

func (m *Ticker) validateSymbol(formats strfmt.Registry) error {

	if err := validate.Required("symbol", "body", m.Symbol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this ticker based on context it is used
func (m *Ticker) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Ticker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Ticker) UnmarshalBinary(b []byte) error {
	var res Ticker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
