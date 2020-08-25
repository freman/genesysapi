// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MatchShiftTradeResponse match shift trade response
//
// swagger:model MatchShiftTradeResponse
type MatchShiftTradeResponse struct {

	// Constraint violations for this shift trade which require shift trade administrator review
	AdminReviewViolations []*ShiftTradeMatchViolation `json:"adminReviewViolations"`

	// The associated shift trade
	Trade *ShiftTradeResponse `json:"trade,omitempty"`

	// Constraint violations which disallow this shift trade
	Violations []*ShiftTradeMatchViolation `json:"violations"`
}

// Validate validates this match shift trade response
func (m *MatchShiftTradeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminReviewViolations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViolations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MatchShiftTradeResponse) validateAdminReviewViolations(formats strfmt.Registry) error {

	if swag.IsZero(m.AdminReviewViolations) { // not required
		return nil
	}

	for i := 0; i < len(m.AdminReviewViolations); i++ {
		if swag.IsZero(m.AdminReviewViolations[i]) { // not required
			continue
		}

		if m.AdminReviewViolations[i] != nil {
			if err := m.AdminReviewViolations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adminReviewViolations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MatchShiftTradeResponse) validateTrade(formats strfmt.Registry) error {

	if swag.IsZero(m.Trade) { // not required
		return nil
	}

	if m.Trade != nil {
		if err := m.Trade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trade")
			}
			return err
		}
	}

	return nil
}

func (m *MatchShiftTradeResponse) validateViolations(formats strfmt.Registry) error {

	if swag.IsZero(m.Violations) { // not required
		return nil
	}

	for i := 0; i < len(m.Violations); i++ {
		if swag.IsZero(m.Violations[i]) { // not required
			continue
		}

		if m.Violations[i] != nil {
			if err := m.Violations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("violations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MatchShiftTradeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MatchShiftTradeResponse) UnmarshalBinary(b []byte) error {
	var res MatchShiftTradeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}