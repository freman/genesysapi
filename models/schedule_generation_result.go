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

// ScheduleGenerationResult schedule generation result
//
// swagger:model ScheduleGenerationResult
type ScheduleGenerationResult struct {

	// Whether the schedule generation run failed
	Failed bool `json:"failed"`

	// The number of schedule generation messages for this schedule generation run
	MessageCount int32 `json:"messageCount,omitempty"`

	// User facing messages related to the schedule generation run
	Messages []*ScheduleGenerationMessage `json:"messages"`

	// The run ID for the schedule generation. Reference this when requesting support
	RunID string `json:"runId,omitempty"`
}

// Validate validates this schedule generation result
func (m *ScheduleGenerationResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScheduleGenerationResult) validateMessages(formats strfmt.Registry) error {

	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleGenerationResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleGenerationResult) UnmarshalBinary(b []byte) error {
	var res ScheduleGenerationResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
