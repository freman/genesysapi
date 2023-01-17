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

// WorkPlanConfigurationViolationMessage work plan configuration violation message
//
// swagger:model WorkPlanConfigurationViolationMessage
type WorkPlanConfigurationViolationMessage struct {

	// Arguments of the message that provide information about the misconfigured value or the threshold that is exceeded by the misconfigured value
	Arguments []*WorkPlanValidationMessageArgument `json:"arguments"`

	// Severity of the message. A message with Error severity indicates the scheduler won't be able to produce schedules and thus the work plan is invalid.
	// Enum: [Information Warning Error]
	Severity string `json:"severity,omitempty"`

	// Type of configuration violation message for this work plan
	// Enum: [ActivitiesOverlap ActivityEndGreaterThanShiftStop ActivityPaidTimeGreaterThanShiftPaidTime ActivityStartBeforeShiftStart ActivityStartGreaterThanEqualToShiftStop ActivityStartIncrementMinutesNotDivisibleByScheduleIntervalMinutes DailyExactPaidMinutes DailyMaxTotalLessThanWeeklyMin DailyMaxTotalLessThanWeeklyMinWithOptional DailyMaxTotalLessThanWeeklyMinWithoutOptional DailyMinTotalGreaterThanWeeklyMax DailyMinTotalGreaterThanWeeklyMaxWithOptional DailyMinTotalGreaterThanWeeklyMaxWithoutOptional DailyRequiredDaysGreaterThanWeeklyMaxDays DailyShiftHasNoDaysSelected DailyShiftMaxPossibilitiesViolated EarliestShiftStopIsTooLate ExactPaidTimeNotDivisibleByGranularity MaxConsecutiveWorkingDaysNoMoreThanDoubleMaxWorkingDaysPerWeek MaxDaysOffPerPlanningPeriodNotCorrect MaxPaidTimeIsMoreThanShiftLength MaxPaidTimeNotDivisibleByGranularity MaxPaidTimePerPlanningPeriod MaxShifts MinPaidTimeNotDivisibleByGranularity MinPaidTimePerPlanningPeriod NoShifts PaidTimeGreaterThanMaxWorkTime PaidTimeLessThanMinWorkTime PaidTimeNotMetByShiftStartStop ShiftDaysSelectMoreThanMinWorkingDays ShiftStopEarlierThanStart ShiftVarianceCannotBeMet WeeklyExactPaidMinutes]
	Type string `json:"type,omitempty"`
}

// Validate validates this work plan configuration violation message
func (m *WorkPlanConfigurationViolationMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkPlanConfigurationViolationMessage) validateArguments(formats strfmt.Registry) error {
	if swag.IsZero(m.Arguments) { // not required
		return nil
	}

	for i := 0; i < len(m.Arguments); i++ {
		if swag.IsZero(m.Arguments[i]) { // not required
			continue
		}

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var workPlanConfigurationViolationMessageTypeSeverityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Information","Warning","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workPlanConfigurationViolationMessageTypeSeverityPropEnum = append(workPlanConfigurationViolationMessageTypeSeverityPropEnum, v)
	}
}

const (

	// WorkPlanConfigurationViolationMessageSeverityInformation captures enum value "Information"
	WorkPlanConfigurationViolationMessageSeverityInformation string = "Information"

	// WorkPlanConfigurationViolationMessageSeverityWarning captures enum value "Warning"
	WorkPlanConfigurationViolationMessageSeverityWarning string = "Warning"

	// WorkPlanConfigurationViolationMessageSeverityError captures enum value "Error"
	WorkPlanConfigurationViolationMessageSeverityError string = "Error"
)

// prop value enum
func (m *WorkPlanConfigurationViolationMessage) validateSeverityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workPlanConfigurationViolationMessageTypeSeverityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WorkPlanConfigurationViolationMessage) validateSeverity(formats strfmt.Registry) error {
	if swag.IsZero(m.Severity) { // not required
		return nil
	}

	// value enum
	if err := m.validateSeverityEnum("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

var workPlanConfigurationViolationMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ActivitiesOverlap","ActivityEndGreaterThanShiftStop","ActivityPaidTimeGreaterThanShiftPaidTime","ActivityStartBeforeShiftStart","ActivityStartGreaterThanEqualToShiftStop","ActivityStartIncrementMinutesNotDivisibleByScheduleIntervalMinutes","DailyExactPaidMinutes","DailyMaxTotalLessThanWeeklyMin","DailyMaxTotalLessThanWeeklyMinWithOptional","DailyMaxTotalLessThanWeeklyMinWithoutOptional","DailyMinTotalGreaterThanWeeklyMax","DailyMinTotalGreaterThanWeeklyMaxWithOptional","DailyMinTotalGreaterThanWeeklyMaxWithoutOptional","DailyRequiredDaysGreaterThanWeeklyMaxDays","DailyShiftHasNoDaysSelected","DailyShiftMaxPossibilitiesViolated","EarliestShiftStopIsTooLate","ExactPaidTimeNotDivisibleByGranularity","MaxConsecutiveWorkingDaysNoMoreThanDoubleMaxWorkingDaysPerWeek","MaxDaysOffPerPlanningPeriodNotCorrect","MaxPaidTimeIsMoreThanShiftLength","MaxPaidTimeNotDivisibleByGranularity","MaxPaidTimePerPlanningPeriod","MaxShifts","MinPaidTimeNotDivisibleByGranularity","MinPaidTimePerPlanningPeriod","NoShifts","PaidTimeGreaterThanMaxWorkTime","PaidTimeLessThanMinWorkTime","PaidTimeNotMetByShiftStartStop","ShiftDaysSelectMoreThanMinWorkingDays","ShiftStopEarlierThanStart","ShiftVarianceCannotBeMet","WeeklyExactPaidMinutes"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workPlanConfigurationViolationMessageTypeTypePropEnum = append(workPlanConfigurationViolationMessageTypeTypePropEnum, v)
	}
}

const (

	// WorkPlanConfigurationViolationMessageTypeActivitiesOverlap captures enum value "ActivitiesOverlap"
	WorkPlanConfigurationViolationMessageTypeActivitiesOverlap string = "ActivitiesOverlap"

	// WorkPlanConfigurationViolationMessageTypeActivityEndGreaterThanShiftStop captures enum value "ActivityEndGreaterThanShiftStop"
	WorkPlanConfigurationViolationMessageTypeActivityEndGreaterThanShiftStop string = "ActivityEndGreaterThanShiftStop"

	// WorkPlanConfigurationViolationMessageTypeActivityPaidTimeGreaterThanShiftPaidTime captures enum value "ActivityPaidTimeGreaterThanShiftPaidTime"
	WorkPlanConfigurationViolationMessageTypeActivityPaidTimeGreaterThanShiftPaidTime string = "ActivityPaidTimeGreaterThanShiftPaidTime"

	// WorkPlanConfigurationViolationMessageTypeActivityStartBeforeShiftStart captures enum value "ActivityStartBeforeShiftStart"
	WorkPlanConfigurationViolationMessageTypeActivityStartBeforeShiftStart string = "ActivityStartBeforeShiftStart"

	// WorkPlanConfigurationViolationMessageTypeActivityStartGreaterThanEqualToShiftStop captures enum value "ActivityStartGreaterThanEqualToShiftStop"
	WorkPlanConfigurationViolationMessageTypeActivityStartGreaterThanEqualToShiftStop string = "ActivityStartGreaterThanEqualToShiftStop"

	// WorkPlanConfigurationViolationMessageTypeActivityStartIncrementMinutesNotDivisibleByScheduleIntervalMinutes captures enum value "ActivityStartIncrementMinutesNotDivisibleByScheduleIntervalMinutes"
	WorkPlanConfigurationViolationMessageTypeActivityStartIncrementMinutesNotDivisibleByScheduleIntervalMinutes string = "ActivityStartIncrementMinutesNotDivisibleByScheduleIntervalMinutes"

	// WorkPlanConfigurationViolationMessageTypeDailyExactPaidMinutes captures enum value "DailyExactPaidMinutes"
	WorkPlanConfigurationViolationMessageTypeDailyExactPaidMinutes string = "DailyExactPaidMinutes"

	// WorkPlanConfigurationViolationMessageTypeDailyMaxTotalLessThanWeeklyMin captures enum value "DailyMaxTotalLessThanWeeklyMin"
	WorkPlanConfigurationViolationMessageTypeDailyMaxTotalLessThanWeeklyMin string = "DailyMaxTotalLessThanWeeklyMin"

	// WorkPlanConfigurationViolationMessageTypeDailyMaxTotalLessThanWeeklyMinWithOptional captures enum value "DailyMaxTotalLessThanWeeklyMinWithOptional"
	WorkPlanConfigurationViolationMessageTypeDailyMaxTotalLessThanWeeklyMinWithOptional string = "DailyMaxTotalLessThanWeeklyMinWithOptional"

	// WorkPlanConfigurationViolationMessageTypeDailyMaxTotalLessThanWeeklyMinWithoutOptional captures enum value "DailyMaxTotalLessThanWeeklyMinWithoutOptional"
	WorkPlanConfigurationViolationMessageTypeDailyMaxTotalLessThanWeeklyMinWithoutOptional string = "DailyMaxTotalLessThanWeeklyMinWithoutOptional"

	// WorkPlanConfigurationViolationMessageTypeDailyMinTotalGreaterThanWeeklyMax captures enum value "DailyMinTotalGreaterThanWeeklyMax"
	WorkPlanConfigurationViolationMessageTypeDailyMinTotalGreaterThanWeeklyMax string = "DailyMinTotalGreaterThanWeeklyMax"

	// WorkPlanConfigurationViolationMessageTypeDailyMinTotalGreaterThanWeeklyMaxWithOptional captures enum value "DailyMinTotalGreaterThanWeeklyMaxWithOptional"
	WorkPlanConfigurationViolationMessageTypeDailyMinTotalGreaterThanWeeklyMaxWithOptional string = "DailyMinTotalGreaterThanWeeklyMaxWithOptional"

	// WorkPlanConfigurationViolationMessageTypeDailyMinTotalGreaterThanWeeklyMaxWithoutOptional captures enum value "DailyMinTotalGreaterThanWeeklyMaxWithoutOptional"
	WorkPlanConfigurationViolationMessageTypeDailyMinTotalGreaterThanWeeklyMaxWithoutOptional string = "DailyMinTotalGreaterThanWeeklyMaxWithoutOptional"

	// WorkPlanConfigurationViolationMessageTypeDailyRequiredDaysGreaterThanWeeklyMaxDays captures enum value "DailyRequiredDaysGreaterThanWeeklyMaxDays"
	WorkPlanConfigurationViolationMessageTypeDailyRequiredDaysGreaterThanWeeklyMaxDays string = "DailyRequiredDaysGreaterThanWeeklyMaxDays"

	// WorkPlanConfigurationViolationMessageTypeDailyShiftHasNoDaysSelected captures enum value "DailyShiftHasNoDaysSelected"
	WorkPlanConfigurationViolationMessageTypeDailyShiftHasNoDaysSelected string = "DailyShiftHasNoDaysSelected"

	// WorkPlanConfigurationViolationMessageTypeDailyShiftMaxPossibilitiesViolated captures enum value "DailyShiftMaxPossibilitiesViolated"
	WorkPlanConfigurationViolationMessageTypeDailyShiftMaxPossibilitiesViolated string = "DailyShiftMaxPossibilitiesViolated"

	// WorkPlanConfigurationViolationMessageTypeEarliestShiftStopIsTooLate captures enum value "EarliestShiftStopIsTooLate"
	WorkPlanConfigurationViolationMessageTypeEarliestShiftStopIsTooLate string = "EarliestShiftStopIsTooLate"

	// WorkPlanConfigurationViolationMessageTypeExactPaidTimeNotDivisibleByGranularity captures enum value "ExactPaidTimeNotDivisibleByGranularity"
	WorkPlanConfigurationViolationMessageTypeExactPaidTimeNotDivisibleByGranularity string = "ExactPaidTimeNotDivisibleByGranularity"

	// WorkPlanConfigurationViolationMessageTypeMaxConsecutiveWorkingDaysNoMoreThanDoubleMaxWorkingDaysPerWeek captures enum value "MaxConsecutiveWorkingDaysNoMoreThanDoubleMaxWorkingDaysPerWeek"
	WorkPlanConfigurationViolationMessageTypeMaxConsecutiveWorkingDaysNoMoreThanDoubleMaxWorkingDaysPerWeek string = "MaxConsecutiveWorkingDaysNoMoreThanDoubleMaxWorkingDaysPerWeek"

	// WorkPlanConfigurationViolationMessageTypeMaxDaysOffPerPlanningPeriodNotCorrect captures enum value "MaxDaysOffPerPlanningPeriodNotCorrect"
	WorkPlanConfigurationViolationMessageTypeMaxDaysOffPerPlanningPeriodNotCorrect string = "MaxDaysOffPerPlanningPeriodNotCorrect"

	// WorkPlanConfigurationViolationMessageTypeMaxPaidTimeIsMoreThanShiftLength captures enum value "MaxPaidTimeIsMoreThanShiftLength"
	WorkPlanConfigurationViolationMessageTypeMaxPaidTimeIsMoreThanShiftLength string = "MaxPaidTimeIsMoreThanShiftLength"

	// WorkPlanConfigurationViolationMessageTypeMaxPaidTimeNotDivisibleByGranularity captures enum value "MaxPaidTimeNotDivisibleByGranularity"
	WorkPlanConfigurationViolationMessageTypeMaxPaidTimeNotDivisibleByGranularity string = "MaxPaidTimeNotDivisibleByGranularity"

	// WorkPlanConfigurationViolationMessageTypeMaxPaidTimePerPlanningPeriod captures enum value "MaxPaidTimePerPlanningPeriod"
	WorkPlanConfigurationViolationMessageTypeMaxPaidTimePerPlanningPeriod string = "MaxPaidTimePerPlanningPeriod"

	// WorkPlanConfigurationViolationMessageTypeMaxShifts captures enum value "MaxShifts"
	WorkPlanConfigurationViolationMessageTypeMaxShifts string = "MaxShifts"

	// WorkPlanConfigurationViolationMessageTypeMinPaidTimeNotDivisibleByGranularity captures enum value "MinPaidTimeNotDivisibleByGranularity"
	WorkPlanConfigurationViolationMessageTypeMinPaidTimeNotDivisibleByGranularity string = "MinPaidTimeNotDivisibleByGranularity"

	// WorkPlanConfigurationViolationMessageTypeMinPaidTimePerPlanningPeriod captures enum value "MinPaidTimePerPlanningPeriod"
	WorkPlanConfigurationViolationMessageTypeMinPaidTimePerPlanningPeriod string = "MinPaidTimePerPlanningPeriod"

	// WorkPlanConfigurationViolationMessageTypeNoShifts captures enum value "NoShifts"
	WorkPlanConfigurationViolationMessageTypeNoShifts string = "NoShifts"

	// WorkPlanConfigurationViolationMessageTypePaidTimeGreaterThanMaxWorkTime captures enum value "PaidTimeGreaterThanMaxWorkTime"
	WorkPlanConfigurationViolationMessageTypePaidTimeGreaterThanMaxWorkTime string = "PaidTimeGreaterThanMaxWorkTime"

	// WorkPlanConfigurationViolationMessageTypePaidTimeLessThanMinWorkTime captures enum value "PaidTimeLessThanMinWorkTime"
	WorkPlanConfigurationViolationMessageTypePaidTimeLessThanMinWorkTime string = "PaidTimeLessThanMinWorkTime"

	// WorkPlanConfigurationViolationMessageTypePaidTimeNotMetByShiftStartStop captures enum value "PaidTimeNotMetByShiftStartStop"
	WorkPlanConfigurationViolationMessageTypePaidTimeNotMetByShiftStartStop string = "PaidTimeNotMetByShiftStartStop"

	// WorkPlanConfigurationViolationMessageTypeShiftDaysSelectMoreThanMinWorkingDays captures enum value "ShiftDaysSelectMoreThanMinWorkingDays"
	WorkPlanConfigurationViolationMessageTypeShiftDaysSelectMoreThanMinWorkingDays string = "ShiftDaysSelectMoreThanMinWorkingDays"

	// WorkPlanConfigurationViolationMessageTypeShiftStopEarlierThanStart captures enum value "ShiftStopEarlierThanStart"
	WorkPlanConfigurationViolationMessageTypeShiftStopEarlierThanStart string = "ShiftStopEarlierThanStart"

	// WorkPlanConfigurationViolationMessageTypeShiftVarianceCannotBeMet captures enum value "ShiftVarianceCannotBeMet"
	WorkPlanConfigurationViolationMessageTypeShiftVarianceCannotBeMet string = "ShiftVarianceCannotBeMet"

	// WorkPlanConfigurationViolationMessageTypeWeeklyExactPaidMinutes captures enum value "WeeklyExactPaidMinutes"
	WorkPlanConfigurationViolationMessageTypeWeeklyExactPaidMinutes string = "WeeklyExactPaidMinutes"
)

// prop value enum
func (m *WorkPlanConfigurationViolationMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workPlanConfigurationViolationMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WorkPlanConfigurationViolationMessage) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this work plan configuration violation message based on the context it is used
func (m *WorkPlanConfigurationViolationMessage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArguments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkPlanConfigurationViolationMessage) contextValidateArguments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Arguments); i++ {

		if m.Arguments[i] != nil {
			if err := m.Arguments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("arguments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("arguments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkPlanConfigurationViolationMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkPlanConfigurationViolationMessage) UnmarshalBinary(b []byte) error {
	var res WorkPlanConfigurationViolationMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
