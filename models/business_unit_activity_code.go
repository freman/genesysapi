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

// BusinessUnitActivityCode business unit activity code
//
// swagger:model BusinessUnitActivityCode
type BusinessUnitActivityCode struct {

	// Whether this activity code is active or has been deleted
	Active bool `json:"active"`

	// Whether an agent can select this activity code when creating or editing a time off request. Null if the activity's category is not time off.
	AgentTimeOffSelectable bool `json:"agentTimeOffSelectable"`

	// The category of the activity code
	// Enum: [OnQueueWork Break Meal Meeting OffQueueWork TimeOff Training Unavailable Unscheduled]
	Category string `json:"category,omitempty"`

	// Whether an agent is paid while performing this activity
	CountsAsPaidTime bool `json:"countsAsPaidTime"`

	// Indicates whether or not the activity should be counted as contiguous work time for calculating daily constraints
	CountsAsWorkTime bool `json:"countsAsWorkTime"`

	// Whether or not this activity code counts toward shrinkage calculations
	CountsTowardShrinkage bool `json:"countsTowardShrinkage"`

	// Whether this is a default activity code
	DefaultCode bool `json:"defaultCode"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Whether this activity code is considered interruptible
	Interruptible bool `json:"interruptible"`

	// The default length of the activity in minutes
	LengthInMinutes int32 `json:"lengthInMinutes,omitempty"`

	// Version metadata of this activity code
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Whether this activity code is considered planned or unplanned shrinkage
	PlannedShrinkage bool `json:"plannedShrinkage"`

	// The secondary presences of this activity code
	SecondaryPresences []*SecondaryPresence `json:"secondaryPresences"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this business unit activity code
func (m *BusinessUnitActivityCode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryPresences(formats); err != nil {
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

var businessUnitActivityCodeTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OnQueueWork","Break","Meal","Meeting","OffQueueWork","TimeOff","Training","Unavailable","Unscheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		businessUnitActivityCodeTypeCategoryPropEnum = append(businessUnitActivityCodeTypeCategoryPropEnum, v)
	}
}

const (

	// BusinessUnitActivityCodeCategoryOnQueueWork captures enum value "OnQueueWork"
	BusinessUnitActivityCodeCategoryOnQueueWork string = "OnQueueWork"

	// BusinessUnitActivityCodeCategoryBreak captures enum value "Break"
	BusinessUnitActivityCodeCategoryBreak string = "Break"

	// BusinessUnitActivityCodeCategoryMeal captures enum value "Meal"
	BusinessUnitActivityCodeCategoryMeal string = "Meal"

	// BusinessUnitActivityCodeCategoryMeeting captures enum value "Meeting"
	BusinessUnitActivityCodeCategoryMeeting string = "Meeting"

	// BusinessUnitActivityCodeCategoryOffQueueWork captures enum value "OffQueueWork"
	BusinessUnitActivityCodeCategoryOffQueueWork string = "OffQueueWork"

	// BusinessUnitActivityCodeCategoryTimeOff captures enum value "TimeOff"
	BusinessUnitActivityCodeCategoryTimeOff string = "TimeOff"

	// BusinessUnitActivityCodeCategoryTraining captures enum value "Training"
	BusinessUnitActivityCodeCategoryTraining string = "Training"

	// BusinessUnitActivityCodeCategoryUnavailable captures enum value "Unavailable"
	BusinessUnitActivityCodeCategoryUnavailable string = "Unavailable"

	// BusinessUnitActivityCodeCategoryUnscheduled captures enum value "Unscheduled"
	BusinessUnitActivityCodeCategoryUnscheduled string = "Unscheduled"
)

// prop value enum
func (m *BusinessUnitActivityCode) validateCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, businessUnitActivityCodeTypeCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BusinessUnitActivityCode) validateCategory(formats strfmt.Registry) error {

	if swag.IsZero(m.Category) { // not required
		return nil
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", m.Category); err != nil {
		return err
	}

	return nil
}

func (m *BusinessUnitActivityCode) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
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

func (m *BusinessUnitActivityCode) validateSecondaryPresences(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryPresences) { // not required
		return nil
	}

	for i := 0; i < len(m.SecondaryPresences); i++ {
		if swag.IsZero(m.SecondaryPresences[i]) { // not required
			continue
		}

		if m.SecondaryPresences[i] != nil {
			if err := m.SecondaryPresences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("secondaryPresences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BusinessUnitActivityCode) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BusinessUnitActivityCode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BusinessUnitActivityCode) UnmarshalBinary(b []byte) error {
	var res BusinessUnitActivityCode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
