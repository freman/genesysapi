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

// WorkPlanValidationMessageArgument work plan validation message argument
//
// swagger:model WorkPlanValidationMessageArgument
type WorkPlanValidationMessageArgument struct {

	// The type of the argument associated with violation messages
	// Enum: [ActivityId ActivityId2 ActivityPaidTimeMinutes ActivityStartTimeMinutes ActivityValidationId ActivityValidationId2 ApplicableDays Count DailyPaidTimeMinutes MaximumDays MaxShiftCount Minutes PaidTimeGranularityMinutes RequiredDays ShiftId ShiftPaidTimeMinutes ShiftStartTimeMinutes ShiftStopTimeMinutes ShiftValidationId WeeklyPaidTimeMinutes Weeks WorkTimeMinutes]
	Type string `json:"type,omitempty"`

	// The value of the argument
	Value string `json:"value,omitempty"`
}

// Validate validates this work plan validation message argument
func (m *WorkPlanValidationMessageArgument) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workPlanValidationMessageArgumentTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ActivityId","ActivityId2","ActivityPaidTimeMinutes","ActivityStartTimeMinutes","ActivityValidationId","ActivityValidationId2","ApplicableDays","Count","DailyPaidTimeMinutes","MaximumDays","MaxShiftCount","Minutes","PaidTimeGranularityMinutes","RequiredDays","ShiftId","ShiftPaidTimeMinutes","ShiftStartTimeMinutes","ShiftStopTimeMinutes","ShiftValidationId","WeeklyPaidTimeMinutes","Weeks","WorkTimeMinutes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workPlanValidationMessageArgumentTypeTypePropEnum = append(workPlanValidationMessageArgumentTypeTypePropEnum, v)
	}
}

const (

	// WorkPlanValidationMessageArgumentTypeActivityID captures enum value "ActivityId"
	WorkPlanValidationMessageArgumentTypeActivityID string = "ActivityId"

	// WorkPlanValidationMessageArgumentTypeActivityId2 captures enum value "ActivityId2"
	WorkPlanValidationMessageArgumentTypeActivityId2 string = "ActivityId2"

	// WorkPlanValidationMessageArgumentTypeActivityPaidTimeMinutes captures enum value "ActivityPaidTimeMinutes"
	WorkPlanValidationMessageArgumentTypeActivityPaidTimeMinutes string = "ActivityPaidTimeMinutes"

	// WorkPlanValidationMessageArgumentTypeActivityStartTimeMinutes captures enum value "ActivityStartTimeMinutes"
	WorkPlanValidationMessageArgumentTypeActivityStartTimeMinutes string = "ActivityStartTimeMinutes"

	// WorkPlanValidationMessageArgumentTypeActivityValidationID captures enum value "ActivityValidationId"
	WorkPlanValidationMessageArgumentTypeActivityValidationID string = "ActivityValidationId"

	// WorkPlanValidationMessageArgumentTypeActivityValidationId2 captures enum value "ActivityValidationId2"
	WorkPlanValidationMessageArgumentTypeActivityValidationId2 string = "ActivityValidationId2"

	// WorkPlanValidationMessageArgumentTypeApplicableDays captures enum value "ApplicableDays"
	WorkPlanValidationMessageArgumentTypeApplicableDays string = "ApplicableDays"

	// WorkPlanValidationMessageArgumentTypeCount captures enum value "Count"
	WorkPlanValidationMessageArgumentTypeCount string = "Count"

	// WorkPlanValidationMessageArgumentTypeDailyPaidTimeMinutes captures enum value "DailyPaidTimeMinutes"
	WorkPlanValidationMessageArgumentTypeDailyPaidTimeMinutes string = "DailyPaidTimeMinutes"

	// WorkPlanValidationMessageArgumentTypeMaximumDays captures enum value "MaximumDays"
	WorkPlanValidationMessageArgumentTypeMaximumDays string = "MaximumDays"

	// WorkPlanValidationMessageArgumentTypeMaxShiftCount captures enum value "MaxShiftCount"
	WorkPlanValidationMessageArgumentTypeMaxShiftCount string = "MaxShiftCount"

	// WorkPlanValidationMessageArgumentTypeMinutes captures enum value "Minutes"
	WorkPlanValidationMessageArgumentTypeMinutes string = "Minutes"

	// WorkPlanValidationMessageArgumentTypePaidTimeGranularityMinutes captures enum value "PaidTimeGranularityMinutes"
	WorkPlanValidationMessageArgumentTypePaidTimeGranularityMinutes string = "PaidTimeGranularityMinutes"

	// WorkPlanValidationMessageArgumentTypeRequiredDays captures enum value "RequiredDays"
	WorkPlanValidationMessageArgumentTypeRequiredDays string = "RequiredDays"

	// WorkPlanValidationMessageArgumentTypeShiftID captures enum value "ShiftId"
	WorkPlanValidationMessageArgumentTypeShiftID string = "ShiftId"

	// WorkPlanValidationMessageArgumentTypeShiftPaidTimeMinutes captures enum value "ShiftPaidTimeMinutes"
	WorkPlanValidationMessageArgumentTypeShiftPaidTimeMinutes string = "ShiftPaidTimeMinutes"

	// WorkPlanValidationMessageArgumentTypeShiftStartTimeMinutes captures enum value "ShiftStartTimeMinutes"
	WorkPlanValidationMessageArgumentTypeShiftStartTimeMinutes string = "ShiftStartTimeMinutes"

	// WorkPlanValidationMessageArgumentTypeShiftStopTimeMinutes captures enum value "ShiftStopTimeMinutes"
	WorkPlanValidationMessageArgumentTypeShiftStopTimeMinutes string = "ShiftStopTimeMinutes"

	// WorkPlanValidationMessageArgumentTypeShiftValidationID captures enum value "ShiftValidationId"
	WorkPlanValidationMessageArgumentTypeShiftValidationID string = "ShiftValidationId"

	// WorkPlanValidationMessageArgumentTypeWeeklyPaidTimeMinutes captures enum value "WeeklyPaidTimeMinutes"
	WorkPlanValidationMessageArgumentTypeWeeklyPaidTimeMinutes string = "WeeklyPaidTimeMinutes"

	// WorkPlanValidationMessageArgumentTypeWeeks captures enum value "Weeks"
	WorkPlanValidationMessageArgumentTypeWeeks string = "Weeks"

	// WorkPlanValidationMessageArgumentTypeWorkTimeMinutes captures enum value "WorkTimeMinutes"
	WorkPlanValidationMessageArgumentTypeWorkTimeMinutes string = "WorkTimeMinutes"
)

// prop value enum
func (m *WorkPlanValidationMessageArgument) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workPlanValidationMessageArgumentTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WorkPlanValidationMessageArgument) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this work plan validation message argument based on context it is used
func (m *WorkPlanValidationMessageArgument) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkPlanValidationMessageArgument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkPlanValidationMessageArgument) UnmarshalBinary(b []byte) error {
	var res WorkPlanValidationMessageArgument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
