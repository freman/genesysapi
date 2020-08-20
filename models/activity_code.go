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

// ActivityCode Activity code data
//
// swagger:model ActivityCode
type ActivityCode struct {

	// Whether an agent can select this activity code when creating or editing a time off request. Null if the activity's category is not time off.
	AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`

	// The activity code's category.
	// Enum: [OnQueueWork Break Meal Meeting OffQueueWork TimeOff Training Unavailable Unscheduled]
	Category string `json:"category,omitempty"`

	// Whether an agent is paid while performing this activity
	CountsAsPaidTime bool `json:"countsAsPaidTime"`

	// Indicates whether or not the activity should be counted as contiguous work time for calculating daily constraints
	CountsAsWorkTime bool `json:"countsAsWorkTime"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Whether this activity code is active or has been deleted
	IsActive bool `json:"isActive"`

	// Whether this is a default activity code
	IsDefault bool `json:"isDefault"`

	// The default length of the activity in minutes
	LengthInMinutes int32 `json:"lengthInMinutes,omitempty"`

	// Version metadata for the associated management unit's list of activity codes
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`

	// The name of the activity code. Default activity codes will be created with an empty name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this activity code
func (m *ActivityCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var activityCodeTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OnQueueWork","Break","Meal","Meeting","OffQueueWork","TimeOff","Training","Unavailable","Unscheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		activityCodeTypeCategoryPropEnum = append(activityCodeTypeCategoryPropEnum, v)
	}
}

const (

	// ActivityCodeCategoryOnQueueWork captures enum value "OnQueueWork"
	ActivityCodeCategoryOnQueueWork string = "OnQueueWork"

	// ActivityCodeCategoryBreak captures enum value "Break"
	ActivityCodeCategoryBreak string = "Break"

	// ActivityCodeCategoryMeal captures enum value "Meal"
	ActivityCodeCategoryMeal string = "Meal"

	// ActivityCodeCategoryMeeting captures enum value "Meeting"
	ActivityCodeCategoryMeeting string = "Meeting"

	// ActivityCodeCategoryOffQueueWork captures enum value "OffQueueWork"
	ActivityCodeCategoryOffQueueWork string = "OffQueueWork"

	// ActivityCodeCategoryTimeOff captures enum value "TimeOff"
	ActivityCodeCategoryTimeOff string = "TimeOff"

	// ActivityCodeCategoryTraining captures enum value "Training"
	ActivityCodeCategoryTraining string = "Training"

	// ActivityCodeCategoryUnavailable captures enum value "Unavailable"
	ActivityCodeCategoryUnavailable string = "Unavailable"

	// ActivityCodeCategoryUnscheduled captures enum value "Unscheduled"
	ActivityCodeCategoryUnscheduled string = "Unscheduled"
)

// prop value enum
func (m *ActivityCode) validateCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, activityCodeTypeCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActivityCode) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *ActivityCode) validateMetadata(formats strfmt.Registry) error {

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

func (m *ActivityCode) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActivityCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActivityCode) UnmarshalBinary(b []byte) error {
	var res ActivityCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
