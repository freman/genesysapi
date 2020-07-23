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

// UpdateActivityCodeRequest Activity Code
//
// swagger:model UpdateActivityCodeRequest
type UpdateActivityCodeRequest struct {

	// Whether an agent can select this activity code when creating or editing a time off request
	AgentTimeOffSelectable bool `json:"agentTimeOffSelectable,omitempty"`

	// The activity code's category. Attempting to change the category of a default activity code will return an error
	// Enum: [OnQueueWork Break Meal Meeting OffQueueWork TimeOff Training Unavailable Unscheduled]
	Category string `json:"category,omitempty"`

	// Whether an agent is paid while performing this activity
	CountsAsPaidTime bool `json:"countsAsPaidTime,omitempty"`

	// Indicates whether or not the activity should be counted as work time
	CountsAsWorkTime bool `json:"countsAsWorkTime,omitempty"`

	// The default length of the activity in minutes
	LengthInMinutes int32 `json:"lengthInMinutes,omitempty"`

	// Version metadata for the associated management unit's list of activity codes
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`

	// The name of the activity code
	Name string `json:"name,omitempty"`
}

// Validate validates this update activity code request
func (m *UpdateActivityCodeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateActivityCodeRequestTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OnQueueWork","Break","Meal","Meeting","OffQueueWork","TimeOff","Training","Unavailable","Unscheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateActivityCodeRequestTypeCategoryPropEnum = append(updateActivityCodeRequestTypeCategoryPropEnum, v)
	}
}

const (

	// UpdateActivityCodeRequestCategoryOnQueueWork captures enum value "OnQueueWork"
	UpdateActivityCodeRequestCategoryOnQueueWork string = "OnQueueWork"

	// UpdateActivityCodeRequestCategoryBreak captures enum value "Break"
	UpdateActivityCodeRequestCategoryBreak string = "Break"

	// UpdateActivityCodeRequestCategoryMeal captures enum value "Meal"
	UpdateActivityCodeRequestCategoryMeal string = "Meal"

	// UpdateActivityCodeRequestCategoryMeeting captures enum value "Meeting"
	UpdateActivityCodeRequestCategoryMeeting string = "Meeting"

	// UpdateActivityCodeRequestCategoryOffQueueWork captures enum value "OffQueueWork"
	UpdateActivityCodeRequestCategoryOffQueueWork string = "OffQueueWork"

	// UpdateActivityCodeRequestCategoryTimeOff captures enum value "TimeOff"
	UpdateActivityCodeRequestCategoryTimeOff string = "TimeOff"

	// UpdateActivityCodeRequestCategoryTraining captures enum value "Training"
	UpdateActivityCodeRequestCategoryTraining string = "Training"

	// UpdateActivityCodeRequestCategoryUnavailable captures enum value "Unavailable"
	UpdateActivityCodeRequestCategoryUnavailable string = "Unavailable"

	// UpdateActivityCodeRequestCategoryUnscheduled captures enum value "Unscheduled"
	UpdateActivityCodeRequestCategoryUnscheduled string = "Unscheduled"
)

// prop value enum
func (m *UpdateActivityCodeRequest) validateCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateActivityCodeRequestTypeCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateActivityCodeRequest) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *UpdateActivityCodeRequest) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateActivityCodeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateActivityCodeRequest) UnmarshalBinary(b []byte) error {
	var res UpdateActivityCodeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
