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

// Objective objective
//
// swagger:model Objective
type Objective struct {

	// start date of the objective. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	DateStart strfmt.Date `json:"dateStart,omitempty"`

	// A flag for whether this objective is enabled for the related metric
	Enabled bool `json:"enabled"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The id of this objective's base template
	TemplateID string `json:"templateId,omitempty"`

	// A filter type for topic Ids. It's only used for objectives with topicIds. Default filter behavior is "or".
	// Enum: [and or]
	TopicIdsFilterType string `json:"topicIdsFilterType,omitempty"`

	// A list of topic ids for detected topic metrics
	Topics []*AddressableEntityRef `json:"topics"`

	// Objective zone specifies min,max points and values for the associated metric
	Zones []*ObjectiveZone `json:"zones"`
}

// Validate validates this objective
func (m *Objective) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopicIdsFilterType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZones(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Objective) validateDateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStart) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStart", "body", "date", m.DateStart.String(), formats); err != nil {
		return err
	}

	return nil
}

var objectiveTypeTopicIdsFilterTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["and","or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectiveTypeTopicIdsFilterTypePropEnum = append(objectiveTypeTopicIdsFilterTypePropEnum, v)
	}
}

const (

	// ObjectiveTopicIdsFilterTypeAnd captures enum value "and"
	ObjectiveTopicIdsFilterTypeAnd string = "and"

	// ObjectiveTopicIdsFilterTypeOr captures enum value "or"
	ObjectiveTopicIdsFilterTypeOr string = "or"
)

// prop value enum
func (m *Objective) validateTopicIdsFilterTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectiveTypeTopicIdsFilterTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Objective) validateTopicIdsFilterType(formats strfmt.Registry) error {

	if swag.IsZero(m.TopicIdsFilterType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTopicIdsFilterTypeEnum("topicIdsFilterType", "body", m.TopicIdsFilterType); err != nil {
		return err
	}

	return nil
}

func (m *Objective) validateTopics(formats strfmt.Registry) error {

	if swag.IsZero(m.Topics) { // not required
		return nil
	}

	for i := 0; i < len(m.Topics); i++ {
		if swag.IsZero(m.Topics[i]) { // not required
			continue
		}

		if m.Topics[i] != nil {
			if err := m.Topics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("topics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Objective) validateZones(formats strfmt.Registry) error {

	if swag.IsZero(m.Zones) { // not required
		return nil
	}

	for i := 0; i < len(m.Zones); i++ {
		if swag.IsZero(m.Zones[i]) { // not required
			continue
		}

		if m.Zones[i] != nil {
			if err := m.Zones[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("zones" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Objective) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Objective) UnmarshalBinary(b []byte) error {
	var res Objective
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
