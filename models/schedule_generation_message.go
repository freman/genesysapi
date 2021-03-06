// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScheduleGenerationMessage schedule generation message
//
// swagger:model ScheduleGenerationMessage
type ScheduleGenerationMessage struct {

	// The arguments describing the message
	Arguments []*SchedulerMessageArgument `json:"arguments"`

	// The type of the message
	// Enum: [AgentNotFound AgentNotInSelectedManagementUnit AgentNotLicensed AgentWithoutWorkPlan WorkPlanNotEnabled WorkPlanNotFound AgentWithoutCapability NoNeedDays UnableToProduceAgentSchedule UnableToScheduleMaxConsecutiveWorkingDays UnableToScheduleMaxConsecutiveWorkingWeekends UnableToScheduleMaxWeeklyPaidTime UnableToScheduleMaxWeeklyWorkDays UnableToScheduleMaxWorkDayPaidTime UnableToScheduleMinConsecutiveNonWorkingTimePerWeek UnableToScheduleMinIntershiftTime UnableToScheduleMinShiftStartDistance UnableToScheduleMinWeeklyPaidTime UnableToScheduleMinWeeklyWorkDays UnableToScheduleMinWorkDayPaidTime UnableToSchedulePlanningPeriodMaxDaysOff UnableToSchedulePlanningPeriodMaxPaidTime UnableToSchedulePlanningPeriodMinDaysOff UnableToSchedulePlanningPeriodMinPaidTime UnableToScheduleShiftVariance UnableToScheduleWorkDay]
	Type string `json:"type,omitempty"`
}

// Validate validates this schedule generation message
func (m *ScheduleGenerationMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArguments(formats); err != nil {
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

func (m *ScheduleGenerationMessage) validateArguments(formats strfmt.Registry) error {

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
				}
				return err
			}
		}

	}

	return nil
}

var scheduleGenerationMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AgentNotFound","AgentNotInSelectedManagementUnit","AgentNotLicensed","AgentWithoutWorkPlan","WorkPlanNotEnabled","WorkPlanNotFound","AgentWithoutCapability","NoNeedDays","UnableToProduceAgentSchedule","UnableToScheduleMaxConsecutiveWorkingDays","UnableToScheduleMaxConsecutiveWorkingWeekends","UnableToScheduleMaxWeeklyPaidTime","UnableToScheduleMaxWeeklyWorkDays","UnableToScheduleMaxWorkDayPaidTime","UnableToScheduleMinConsecutiveNonWorkingTimePerWeek","UnableToScheduleMinIntershiftTime","UnableToScheduleMinShiftStartDistance","UnableToScheduleMinWeeklyPaidTime","UnableToScheduleMinWeeklyWorkDays","UnableToScheduleMinWorkDayPaidTime","UnableToSchedulePlanningPeriodMaxDaysOff","UnableToSchedulePlanningPeriodMaxPaidTime","UnableToSchedulePlanningPeriodMinDaysOff","UnableToSchedulePlanningPeriodMinPaidTime","UnableToScheduleShiftVariance","UnableToScheduleWorkDay"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scheduleGenerationMessageTypeTypePropEnum = append(scheduleGenerationMessageTypeTypePropEnum, v)
	}
}

const (

	// ScheduleGenerationMessageTypeAgentNotFound captures enum value "AgentNotFound"
	ScheduleGenerationMessageTypeAgentNotFound string = "AgentNotFound"

	// ScheduleGenerationMessageTypeAgentNotInSelectedManagementUnit captures enum value "AgentNotInSelectedManagementUnit"
	ScheduleGenerationMessageTypeAgentNotInSelectedManagementUnit string = "AgentNotInSelectedManagementUnit"

	// ScheduleGenerationMessageTypeAgentNotLicensed captures enum value "AgentNotLicensed"
	ScheduleGenerationMessageTypeAgentNotLicensed string = "AgentNotLicensed"

	// ScheduleGenerationMessageTypeAgentWithoutWorkPlan captures enum value "AgentWithoutWorkPlan"
	ScheduleGenerationMessageTypeAgentWithoutWorkPlan string = "AgentWithoutWorkPlan"

	// ScheduleGenerationMessageTypeWorkPlanNotEnabled captures enum value "WorkPlanNotEnabled"
	ScheduleGenerationMessageTypeWorkPlanNotEnabled string = "WorkPlanNotEnabled"

	// ScheduleGenerationMessageTypeWorkPlanNotFound captures enum value "WorkPlanNotFound"
	ScheduleGenerationMessageTypeWorkPlanNotFound string = "WorkPlanNotFound"

	// ScheduleGenerationMessageTypeAgentWithoutCapability captures enum value "AgentWithoutCapability"
	ScheduleGenerationMessageTypeAgentWithoutCapability string = "AgentWithoutCapability"

	// ScheduleGenerationMessageTypeNoNeedDays captures enum value "NoNeedDays"
	ScheduleGenerationMessageTypeNoNeedDays string = "NoNeedDays"

	// ScheduleGenerationMessageTypeUnableToProduceAgentSchedule captures enum value "UnableToProduceAgentSchedule"
	ScheduleGenerationMessageTypeUnableToProduceAgentSchedule string = "UnableToProduceAgentSchedule"

	// ScheduleGenerationMessageTypeUnableToScheduleMaxConsecutiveWorkingDays captures enum value "UnableToScheduleMaxConsecutiveWorkingDays"
	ScheduleGenerationMessageTypeUnableToScheduleMaxConsecutiveWorkingDays string = "UnableToScheduleMaxConsecutiveWorkingDays"

	// ScheduleGenerationMessageTypeUnableToScheduleMaxConsecutiveWorkingWeekends captures enum value "UnableToScheduleMaxConsecutiveWorkingWeekends"
	ScheduleGenerationMessageTypeUnableToScheduleMaxConsecutiveWorkingWeekends string = "UnableToScheduleMaxConsecutiveWorkingWeekends"

	// ScheduleGenerationMessageTypeUnableToScheduleMaxWeeklyPaidTime captures enum value "UnableToScheduleMaxWeeklyPaidTime"
	ScheduleGenerationMessageTypeUnableToScheduleMaxWeeklyPaidTime string = "UnableToScheduleMaxWeeklyPaidTime"

	// ScheduleGenerationMessageTypeUnableToScheduleMaxWeeklyWorkDays captures enum value "UnableToScheduleMaxWeeklyWorkDays"
	ScheduleGenerationMessageTypeUnableToScheduleMaxWeeklyWorkDays string = "UnableToScheduleMaxWeeklyWorkDays"

	// ScheduleGenerationMessageTypeUnableToScheduleMaxWorkDayPaidTime captures enum value "UnableToScheduleMaxWorkDayPaidTime"
	ScheduleGenerationMessageTypeUnableToScheduleMaxWorkDayPaidTime string = "UnableToScheduleMaxWorkDayPaidTime"

	// ScheduleGenerationMessageTypeUnableToScheduleMinConsecutiveNonWorkingTimePerWeek captures enum value "UnableToScheduleMinConsecutiveNonWorkingTimePerWeek"
	ScheduleGenerationMessageTypeUnableToScheduleMinConsecutiveNonWorkingTimePerWeek string = "UnableToScheduleMinConsecutiveNonWorkingTimePerWeek"

	// ScheduleGenerationMessageTypeUnableToScheduleMinIntershiftTime captures enum value "UnableToScheduleMinIntershiftTime"
	ScheduleGenerationMessageTypeUnableToScheduleMinIntershiftTime string = "UnableToScheduleMinIntershiftTime"

	// ScheduleGenerationMessageTypeUnableToScheduleMinShiftStartDistance captures enum value "UnableToScheduleMinShiftStartDistance"
	ScheduleGenerationMessageTypeUnableToScheduleMinShiftStartDistance string = "UnableToScheduleMinShiftStartDistance"

	// ScheduleGenerationMessageTypeUnableToScheduleMinWeeklyPaidTime captures enum value "UnableToScheduleMinWeeklyPaidTime"
	ScheduleGenerationMessageTypeUnableToScheduleMinWeeklyPaidTime string = "UnableToScheduleMinWeeklyPaidTime"

	// ScheduleGenerationMessageTypeUnableToScheduleMinWeeklyWorkDays captures enum value "UnableToScheduleMinWeeklyWorkDays"
	ScheduleGenerationMessageTypeUnableToScheduleMinWeeklyWorkDays string = "UnableToScheduleMinWeeklyWorkDays"

	// ScheduleGenerationMessageTypeUnableToScheduleMinWorkDayPaidTime captures enum value "UnableToScheduleMinWorkDayPaidTime"
	ScheduleGenerationMessageTypeUnableToScheduleMinWorkDayPaidTime string = "UnableToScheduleMinWorkDayPaidTime"

	// ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMaxDaysOff captures enum value "UnableToSchedulePlanningPeriodMaxDaysOff"
	ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMaxDaysOff string = "UnableToSchedulePlanningPeriodMaxDaysOff"

	// ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMaxPaidTime captures enum value "UnableToSchedulePlanningPeriodMaxPaidTime"
	ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMaxPaidTime string = "UnableToSchedulePlanningPeriodMaxPaidTime"

	// ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMinDaysOff captures enum value "UnableToSchedulePlanningPeriodMinDaysOff"
	ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMinDaysOff string = "UnableToSchedulePlanningPeriodMinDaysOff"

	// ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMinPaidTime captures enum value "UnableToSchedulePlanningPeriodMinPaidTime"
	ScheduleGenerationMessageTypeUnableToSchedulePlanningPeriodMinPaidTime string = "UnableToSchedulePlanningPeriodMinPaidTime"

	// ScheduleGenerationMessageTypeUnableToScheduleShiftVariance captures enum value "UnableToScheduleShiftVariance"
	ScheduleGenerationMessageTypeUnableToScheduleShiftVariance string = "UnableToScheduleShiftVariance"

	// ScheduleGenerationMessageTypeUnableToScheduleWorkDay captures enum value "UnableToScheduleWorkDay"
	ScheduleGenerationMessageTypeUnableToScheduleWorkDay string = "UnableToScheduleWorkDay"
)

// prop value enum
func (m *ScheduleGenerationMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scheduleGenerationMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScheduleGenerationMessage) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScheduleGenerationMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScheduleGenerationMessage) UnmarshalBinary(b []byte) error {
	var res ScheduleGenerationMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
