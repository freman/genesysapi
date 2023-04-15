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

// LearningScheduleSlotsQueryResponse learning schedule slots query response
//
// swagger:model LearningScheduleSlotsQueryResponse
type LearningScheduleSlotsQueryResponse struct {

	// List of slots where Learning activity can be scheduled
	SuggestedSlots []*LearningSlot `json:"suggestedSlots"`

	// Detailed data for WFM scheduled activities
	WfmScheduleActivities []*LearningSlotWfmScheduleActivity `json:"wfmScheduleActivities"`
}

// Validate validates this learning schedule slots query response
func (m *LearningScheduleSlotsQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSuggestedSlots(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWfmScheduleActivities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningScheduleSlotsQueryResponse) validateSuggestedSlots(formats strfmt.Registry) error {
	if swag.IsZero(m.SuggestedSlots) { // not required
		return nil
	}

	for i := 0; i < len(m.SuggestedSlots); i++ {
		if swag.IsZero(m.SuggestedSlots[i]) { // not required
			continue
		}

		if m.SuggestedSlots[i] != nil {
			if err := m.SuggestedSlots[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suggestedSlots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suggestedSlots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningScheduleSlotsQueryResponse) validateWfmScheduleActivities(formats strfmt.Registry) error {
	if swag.IsZero(m.WfmScheduleActivities) { // not required
		return nil
	}

	for i := 0; i < len(m.WfmScheduleActivities); i++ {
		if swag.IsZero(m.WfmScheduleActivities[i]) { // not required
			continue
		}

		if m.WfmScheduleActivities[i] != nil {
			if err := m.WfmScheduleActivities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wfmScheduleActivities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wfmScheduleActivities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this learning schedule slots query response based on the context it is used
func (m *LearningScheduleSlotsQueryResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSuggestedSlots(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWfmScheduleActivities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningScheduleSlotsQueryResponse) contextValidateSuggestedSlots(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.SuggestedSlots); i++ {

		if m.SuggestedSlots[i] != nil {
			if err := m.SuggestedSlots[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("suggestedSlots" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("suggestedSlots" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningScheduleSlotsQueryResponse) contextValidateWfmScheduleActivities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WfmScheduleActivities); i++ {

		if m.WfmScheduleActivities[i] != nil {
			if err := m.WfmScheduleActivities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wfmScheduleActivities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wfmScheduleActivities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningScheduleSlotsQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningScheduleSlotsQueryResponse) UnmarshalBinary(b []byte) error {
	var res LearningScheduleSlotsQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
