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

// LearningSlotWfmScheduleActivity learning slot wfm schedule activity
//
// swagger:model LearningSlotWfmScheduleActivity
type LearningSlotWfmScheduleActivity struct {

	// List of user's scheduled activities
	Activities []*LearningSlotScheduleActivity `json:"activities"`

	// List of user's days off
	FullDayTimeOffMarkers []*LearningSlotFullDayTimeOffMarker `json:"fullDayTimeOffMarkers"`

	// User that the schedule is for
	User *UserReference `json:"user,omitempty"`
}

// Validate validates this learning slot wfm schedule activity
func (m *LearningSlotWfmScheduleActivity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullDayTimeOffMarkers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningSlotWfmScheduleActivity) validateActivities(formats strfmt.Registry) error {
	if swag.IsZero(m.Activities) { // not required
		return nil
	}

	for i := 0; i < len(m.Activities); i++ {
		if swag.IsZero(m.Activities[i]) { // not required
			continue
		}

		if m.Activities[i] != nil {
			if err := m.Activities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningSlotWfmScheduleActivity) validateFullDayTimeOffMarkers(formats strfmt.Registry) error {
	if swag.IsZero(m.FullDayTimeOffMarkers) { // not required
		return nil
	}

	for i := 0; i < len(m.FullDayTimeOffMarkers); i++ {
		if swag.IsZero(m.FullDayTimeOffMarkers[i]) { // not required
			continue
		}

		if m.FullDayTimeOffMarkers[i] != nil {
			if err := m.FullDayTimeOffMarkers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fullDayTimeOffMarkers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fullDayTimeOffMarkers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningSlotWfmScheduleActivity) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this learning slot wfm schedule activity based on the context it is used
func (m *LearningSlotWfmScheduleActivity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActivities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateFullDayTimeOffMarkers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningSlotWfmScheduleActivity) contextValidateActivities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Activities); i++ {

		if m.Activities[i] != nil {
			if err := m.Activities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("activities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("activities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningSlotWfmScheduleActivity) contextValidateFullDayTimeOffMarkers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.FullDayTimeOffMarkers); i++ {

		if m.FullDayTimeOffMarkers[i] != nil {
			if err := m.FullDayTimeOffMarkers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("fullDayTimeOffMarkers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("fullDayTimeOffMarkers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningSlotWfmScheduleActivity) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningSlotWfmScheduleActivity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningSlotWfmScheduleActivity) UnmarshalBinary(b []byte) error {
	var res LearningSlotWfmScheduleActivity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}