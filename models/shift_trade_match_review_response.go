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

// ShiftTradeMatchReviewResponse shift trade match review response
//
// swagger:model ShiftTradeMatchReviewResponse
type ShiftTradeMatchReviewResponse struct {

	// Constraint violations associated with this shift trade which require shift trade administrator review
	AdminReviewViolations []*ShiftTradeMatchViolation `json:"adminReviewViolations"`

	// Details for the initiatingUser side of the shift trade
	InitiatingUser *ShiftTradeMatchReviewUserResponse `json:"initiatingUser,omitempty"`

	// Details for the receivingUser side of the shift trade
	ReceivingUser *ShiftTradeMatchReviewUserResponse `json:"receivingUser,omitempty"`

	// Constraint violations introduced after being matched that would normally disallow a trade, but which can still be overridden by the shift trade administrator
	Violations []*ShiftTradeMatchViolation `json:"violations"`
}

// Validate validates this shift trade match review response
func (m *ShiftTradeMatchReviewResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminReviewViolations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatingUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivingUser(formats); err != nil {
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

func (m *ShiftTradeMatchReviewResponse) validateAdminReviewViolations(formats strfmt.Registry) error {

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

func (m *ShiftTradeMatchReviewResponse) validateInitiatingUser(formats strfmt.Registry) error {

	if swag.IsZero(m.InitiatingUser) { // not required
		return nil
	}

	if m.InitiatingUser != nil {
		if err := m.InitiatingUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiatingUser")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeMatchReviewResponse) validateReceivingUser(formats strfmt.Registry) error {

	if swag.IsZero(m.ReceivingUser) { // not required
		return nil
	}

	if m.ReceivingUser != nil {
		if err := m.ReceivingUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receivingUser")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeMatchReviewResponse) validateViolations(formats strfmt.Registry) error {

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
func (m *ShiftTradeMatchReviewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShiftTradeMatchReviewResponse) UnmarshalBinary(b []byte) error {
	var res ShiftTradeMatchReviewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}