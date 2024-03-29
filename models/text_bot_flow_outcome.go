// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TextBotFlowOutcome Flow Outcome data related to a bot flow which is exiting gracefully.
//
// swagger:model TextBotFlowOutcome
type TextBotFlowOutcome struct {

	// The timestamp for when the Flow Outcome finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateEnd strfmt.DateTime `json:"dateEnd,omitempty"`

	// The timestamp for when the Flow Outcome began. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateStart strfmt.DateTime `json:"dateStart,omitempty"`

	// The Flow Milestones for the Flow Outcome.
	Milestones []*TextBotFlowMilestone `json:"milestones"`

	// The Flow Outcome ID.
	OutcomeID string `json:"outcomeId,omitempty"`

	// The value of the FlowOutcome.
	// Enum: [SUCCESS FAILURE]
	OutcomeValue string `json:"outcomeValue,omitempty"`
}

// Validate validates this text bot flow outcome
func (m *TextBotFlowOutcome) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMilestones(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutcomeValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextBotFlowOutcome) validateDateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.DateEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("dateEnd", "body", "date-time", m.DateEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TextBotFlowOutcome) validateDateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date-time", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TextBotFlowOutcome) validateMilestones(formats strfmt.Registry) error {
	if swag.IsZero(m.Milestones) { // not required
		return nil
	}

	for i := 0; i < len(m.Milestones); i++ {
		if swag.IsZero(m.Milestones[i]) { // not required
			continue
		}

		if m.Milestones[i] != nil {
			if err := m.Milestones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("milestones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("milestones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var textBotFlowOutcomeTypeOutcomeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","FAILURE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		textBotFlowOutcomeTypeOutcomeValuePropEnum = append(textBotFlowOutcomeTypeOutcomeValuePropEnum, v)
	}
}

const (

	// TextBotFlowOutcomeOutcomeValueSUCCESS captures enum value "SUCCESS"
	TextBotFlowOutcomeOutcomeValueSUCCESS string = "SUCCESS"

	// TextBotFlowOutcomeOutcomeValueFAILURE captures enum value "FAILURE"
	TextBotFlowOutcomeOutcomeValueFAILURE string = "FAILURE"
)

// prop value enum
func (m *TextBotFlowOutcome) validateOutcomeValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, textBotFlowOutcomeTypeOutcomeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TextBotFlowOutcome) validateOutcomeValue(formats strfmt.Registry) error {
	if swag.IsZero(m.OutcomeValue) { // not required
		return nil
	}

	// value enum
	if err := m.validateOutcomeValueEnum("outcomeValue", "body", m.OutcomeValue); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this text bot flow outcome based on the context it is used
func (m *TextBotFlowOutcome) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateMilestones(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TextBotFlowOutcome) contextValidateMilestones(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Milestones); i++ {

		if m.Milestones[i] != nil {
			if err := m.Milestones[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("milestones" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("milestones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TextBotFlowOutcome) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TextBotFlowOutcome) UnmarshalBinary(b []byte) error {
	var res TextBotFlowOutcome
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
