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

// SearchShiftTradeResponse search shift trade response
//
// swagger:model SearchShiftTradeResponse
type SearchShiftTradeResponse struct {

	// IDs of shifts which match the search criteria
	MatchingReceivingShiftIds []string `json:"matchingReceivingShiftIds"`

	// A preview of what the shift trade would look like if matched
	Preview *ShiftTradePreviewResponse `json:"preview,omitempty"`

	// A trade which matches search criteria
	Trade *ShiftTradeResponse `json:"trade,omitempty"`
}

// Validate validates this search shift trade response
func (m *SearchShiftTradeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePreview(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrade(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchShiftTradeResponse) validatePreview(formats strfmt.Registry) error {
	if swag.IsZero(m.Preview) { // not required
		return nil
	}

	if m.Preview != nil {
		if err := m.Preview.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preview")
			}
			return err
		}
	}

	return nil
}

func (m *SearchShiftTradeResponse) validateTrade(formats strfmt.Registry) error {
	if swag.IsZero(m.Trade) { // not required
		return nil
	}

	if m.Trade != nil {
		if err := m.Trade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trade")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trade")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this search shift trade response based on the context it is used
func (m *SearchShiftTradeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePreview(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrade(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SearchShiftTradeResponse) contextValidatePreview(ctx context.Context, formats strfmt.Registry) error {

	if m.Preview != nil {
		if err := m.Preview.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("preview")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("preview")
			}
			return err
		}
	}

	return nil
}

func (m *SearchShiftTradeResponse) contextValidateTrade(ctx context.Context, formats strfmt.Registry) error {

	if m.Trade != nil {
		if err := m.Trade.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trade")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trade")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SearchShiftTradeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchShiftTradeResponse) UnmarshalBinary(b []byte) error {
	var res SearchShiftTradeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
