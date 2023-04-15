// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// LearningSlot learning slot
//
// swagger:model LearningSlot
type LearningSlot struct {

	// Start date and time of scheduled Learning activity slot. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateStart strfmt.DateTime `json:"dateStart,omitempty"`

	// Rating based on the staffing difference for scheduled slot
	// Enum: [Poor Neutral Good]
	DifferenceRating string `json:"differenceRating,omitempty"`

	// Length of Learning activity slot in minutes
	LengthInMinutes int32 `json:"lengthInMinutes,omitempty"`

	// Difference between scheduled and forecast headcount for this slot after scheduling the Learning activity
	StaffingDifference float64 `json:"staffingDifference,omitempty"`
}

// Validate validates this learning slot
func (m *LearningSlot) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDifferenceRating(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningSlot) validateDateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date-time", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

var learningSlotTypeDifferenceRatingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Poor","Neutral","Good"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningSlotTypeDifferenceRatingPropEnum = append(learningSlotTypeDifferenceRatingPropEnum, v)
	}
}

const (

	// LearningSlotDifferenceRatingPoor captures enum value "Poor"
	LearningSlotDifferenceRatingPoor string = "Poor"

	// LearningSlotDifferenceRatingNeutral captures enum value "Neutral"
	LearningSlotDifferenceRatingNeutral string = "Neutral"

	// LearningSlotDifferenceRatingGood captures enum value "Good"
	LearningSlotDifferenceRatingGood string = "Good"
)

// prop value enum
func (m *LearningSlot) validateDifferenceRatingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningSlotTypeDifferenceRatingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningSlot) validateDifferenceRating(formats strfmt.Registry) error {
	if swag.IsZero(m.DifferenceRating) { // not required
		return nil
	}

	// value enum
	if err := m.validateDifferenceRatingEnum("differenceRating", "body", m.DifferenceRating); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this learning slot based on context it is used
func (m *LearningSlot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *LearningSlot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningSlot) UnmarshalBinary(b []byte) error {
	var res LearningSlot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}