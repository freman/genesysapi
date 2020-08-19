// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// HistoricalAdherenceActuals historical adherence actuals
//
// swagger:model HistoricalAdherenceActuals
type HistoricalAdherenceActuals struct {

	// Activity in which the user is actually engaged
	// Enum: [OnQueueWork Break Meal Meeting OffQueueWork TimeOff Training Unavailable Unscheduled]
	ActualActivityCategory string `json:"actualActivityCategory,omitempty"`

	// Actual end offset in seconds relative to query start time
	EndOffsetSeconds int32 `json:"endOffsetSeconds,omitempty"`

	// Actual start offset in seconds relative to query start time
	StartOffsetSeconds int32 `json:"startOffsetSeconds,omitempty"`
}

// Validate validates this historical adherence actuals
func (m *HistoricalAdherenceActuals) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActualActivityCategory(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var historicalAdherenceActualsTypeActualActivityCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OnQueueWork","Break","Meal","Meeting","OffQueueWork","TimeOff","Training","Unavailable","Unscheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		historicalAdherenceActualsTypeActualActivityCategoryPropEnum = append(historicalAdherenceActualsTypeActualActivityCategoryPropEnum, v)
	}
}

const (

	// HistoricalAdherenceActualsActualActivityCategoryOnQueueWork captures enum value "OnQueueWork"
	HistoricalAdherenceActualsActualActivityCategoryOnQueueWork string = "OnQueueWork"

	// HistoricalAdherenceActualsActualActivityCategoryBreak captures enum value "Break"
	HistoricalAdherenceActualsActualActivityCategoryBreak string = "Break"

	// HistoricalAdherenceActualsActualActivityCategoryMeal captures enum value "Meal"
	HistoricalAdherenceActualsActualActivityCategoryMeal string = "Meal"

	// HistoricalAdherenceActualsActualActivityCategoryMeeting captures enum value "Meeting"
	HistoricalAdherenceActualsActualActivityCategoryMeeting string = "Meeting"

	// HistoricalAdherenceActualsActualActivityCategoryOffQueueWork captures enum value "OffQueueWork"
	HistoricalAdherenceActualsActualActivityCategoryOffQueueWork string = "OffQueueWork"

	// HistoricalAdherenceActualsActualActivityCategoryTimeOff captures enum value "TimeOff"
	HistoricalAdherenceActualsActualActivityCategoryTimeOff string = "TimeOff"

	// HistoricalAdherenceActualsActualActivityCategoryTraining captures enum value "Training"
	HistoricalAdherenceActualsActualActivityCategoryTraining string = "Training"

	// HistoricalAdherenceActualsActualActivityCategoryUnavailable captures enum value "Unavailable"
	HistoricalAdherenceActualsActualActivityCategoryUnavailable string = "Unavailable"

	// HistoricalAdherenceActualsActualActivityCategoryUnscheduled captures enum value "Unscheduled"
	HistoricalAdherenceActualsActualActivityCategoryUnscheduled string = "Unscheduled"
)

// prop value enum
func (m *HistoricalAdherenceActuals) validateActualActivityCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, historicalAdherenceActualsTypeActualActivityCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *HistoricalAdherenceActuals) validateActualActivityCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.ActualActivityCategory) { // not required
		return nil
	}

	// value enum
	if err := m.validateActualActivityCategoryEnum("actualActivityCategory", "body", m.ActualActivityCategory); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HistoricalAdherenceActuals) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HistoricalAdherenceActuals) UnmarshalBinary(b []byte) error {
	var res HistoricalAdherenceActuals
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
