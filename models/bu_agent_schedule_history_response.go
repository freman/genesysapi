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

// BuAgentScheduleHistoryResponse bu agent schedule history response
//
// swagger:model BuAgentScheduleHistoryResponse
type BuAgentScheduleHistoryResponse struct {

	// The originally published agent schedules
	BasePublishedSchedule *BuAgentScheduleHistoryChange `json:"basePublishedSchedule,omitempty"`

	// The list of changes for the schedule history
	Changes []*BuAgentScheduleHistoryChange `json:"changes"`

	// The changes dropped from the schedule history. This will happen if the schedule history is too large
	DroppedChanges []*BuAgentScheduleHistoryDroppedChange `json:"droppedChanges"`

	// The list of previously published schedules
	PriorPublishedSchedules []*BuScheduleReference `json:"priorPublishedSchedules"`
}

// Validate validates this bu agent schedule history response
func (m *BuAgentScheduleHistoryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBasePublishedSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDroppedChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriorPublishedSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAgentScheduleHistoryResponse) validateBasePublishedSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.BasePublishedSchedule) { // not required
		return nil
	}

	if m.BasePublishedSchedule != nil {
		if err := m.BasePublishedSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basePublishedSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basePublishedSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *BuAgentScheduleHistoryResponse) validateChanges(formats strfmt.Registry) error {
	if swag.IsZero(m.Changes) { // not required
		return nil
	}

	for i := 0; i < len(m.Changes); i++ {
		if swag.IsZero(m.Changes[i]) { // not required
			continue
		}

		if m.Changes[i] != nil {
			if err := m.Changes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuAgentScheduleHistoryResponse) validateDroppedChanges(formats strfmt.Registry) error {
	if swag.IsZero(m.DroppedChanges) { // not required
		return nil
	}

	for i := 0; i < len(m.DroppedChanges); i++ {
		if swag.IsZero(m.DroppedChanges[i]) { // not required
			continue
		}

		if m.DroppedChanges[i] != nil {
			if err := m.DroppedChanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("droppedChanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("droppedChanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuAgentScheduleHistoryResponse) validatePriorPublishedSchedules(formats strfmt.Registry) error {
	if swag.IsZero(m.PriorPublishedSchedules) { // not required
		return nil
	}

	for i := 0; i < len(m.PriorPublishedSchedules); i++ {
		if swag.IsZero(m.PriorPublishedSchedules[i]) { // not required
			continue
		}

		if m.PriorPublishedSchedules[i] != nil {
			if err := m.PriorPublishedSchedules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priorPublishedSchedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("priorPublishedSchedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bu agent schedule history response based on the context it is used
func (m *BuAgentScheduleHistoryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBasePublishedSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateChanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDroppedChanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriorPublishedSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAgentScheduleHistoryResponse) contextValidateBasePublishedSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.BasePublishedSchedule != nil {
		if err := m.BasePublishedSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("basePublishedSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("basePublishedSchedule")
			}
			return err
		}
	}

	return nil
}

func (m *BuAgentScheduleHistoryResponse) contextValidateChanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Changes); i++ {

		if m.Changes[i] != nil {
			if err := m.Changes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("changes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("changes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuAgentScheduleHistoryResponse) contextValidateDroppedChanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DroppedChanges); i++ {

		if m.DroppedChanges[i] != nil {
			if err := m.DroppedChanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("droppedChanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("droppedChanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuAgentScheduleHistoryResponse) contextValidatePriorPublishedSchedules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PriorPublishedSchedules); i++ {

		if m.PriorPublishedSchedules[i] != nil {
			if err := m.PriorPublishedSchedules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priorPublishedSchedules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("priorPublishedSchedules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAgentScheduleHistoryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAgentScheduleHistoryResponse) UnmarshalBinary(b []byte) error {
	var res BuAgentScheduleHistoryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
