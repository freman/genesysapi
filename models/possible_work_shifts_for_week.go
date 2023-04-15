// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PossibleWorkShiftsForWeek possible work shifts for week
//
// swagger:model PossibleWorkShiftsForWeek
type PossibleWorkShiftsForWeek struct {

	// Daily shifts in this possible weekly shift
	DailyPossibleShifts []*DailyPossibleShift `json:"dailyPossibleShifts"`

	// ID of this possible weekly shift
	ID int32 `json:"id,omitempty"`
}

// Validate validates this possible work shifts for week
func (m *PossibleWorkShiftsForWeek) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDailyPossibleShifts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PossibleWorkShiftsForWeek) validateDailyPossibleShifts(formats strfmt.Registry) error {
	if swag.IsZero(m.DailyPossibleShifts) { // not required
		return nil
	}

	for i := 0; i < len(m.DailyPossibleShifts); i++ {
		if swag.IsZero(m.DailyPossibleShifts[i]) { // not required
			continue
		}

		if m.DailyPossibleShifts[i] != nil {
			if err := m.DailyPossibleShifts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dailyPossibleShifts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dailyPossibleShifts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this possible work shifts for week based on the context it is used
func (m *PossibleWorkShiftsForWeek) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDailyPossibleShifts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PossibleWorkShiftsForWeek) contextValidateDailyPossibleShifts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DailyPossibleShifts); i++ {

		if m.DailyPossibleShifts[i] != nil {
			if err := m.DailyPossibleShifts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dailyPossibleShifts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dailyPossibleShifts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PossibleWorkShiftsForWeek) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PossibleWorkShiftsForWeek) UnmarshalBinary(b []byte) error {
	var res PossibleWorkShiftsForWeek
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
