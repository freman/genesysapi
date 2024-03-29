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

// ObjectiveZone objective zone
//
// swagger:model ObjectiveZone
type ObjectiveZone struct {

	// direction type
	// Required: true
	// Enum: [Up Down Flat]
	DirectionType *string `json:"directionType"`

	// label
	// Required: true
	Label *string `json:"label"`

	// lower limit points
	// Required: true
	LowerLimitPoints *int32 `json:"lowerLimitPoints"`

	// lower limit value
	LowerLimitValue int32 `json:"lowerLimitValue,omitempty"`

	// upper limit points
	// Required: true
	UpperLimitPoints *int32 `json:"upperLimitPoints"`

	// upper limit value
	UpperLimitValue int32 `json:"upperLimitValue,omitempty"`

	// zone type
	// Required: true
	// Enum: [Good Target Great Out]
	ZoneType *string `json:"zoneType"`
}

// Validate validates this objective zone
func (m *ObjectiveZone) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDirectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLowerLimitPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpperLimitPoints(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZoneType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var objectiveZoneTypeDirectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Up","Down","Flat"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectiveZoneTypeDirectionTypePropEnum = append(objectiveZoneTypeDirectionTypePropEnum, v)
	}
}

const (

	// ObjectiveZoneDirectionTypeUp captures enum value "Up"
	ObjectiveZoneDirectionTypeUp string = "Up"

	// ObjectiveZoneDirectionTypeDown captures enum value "Down"
	ObjectiveZoneDirectionTypeDown string = "Down"

	// ObjectiveZoneDirectionTypeFlat captures enum value "Flat"
	ObjectiveZoneDirectionTypeFlat string = "Flat"
)

// prop value enum
func (m *ObjectiveZone) validateDirectionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectiveZoneTypeDirectionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectiveZone) validateDirectionType(formats strfmt.Registry) error {

	if err := validate.Required("directionType", "body", m.DirectionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateDirectionTypeEnum("directionType", "body", *m.DirectionType); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveZone) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveZone) validateLowerLimitPoints(formats strfmt.Registry) error {

	if err := validate.Required("lowerLimitPoints", "body", m.LowerLimitPoints); err != nil {
		return err
	}

	return nil
}

func (m *ObjectiveZone) validateUpperLimitPoints(formats strfmt.Registry) error {

	if err := validate.Required("upperLimitPoints", "body", m.UpperLimitPoints); err != nil {
		return err
	}

	return nil
}

var objectiveZoneTypeZoneTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Good","Target","Great","Out"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectiveZoneTypeZoneTypePropEnum = append(objectiveZoneTypeZoneTypePropEnum, v)
	}
}

const (

	// ObjectiveZoneZoneTypeGood captures enum value "Good"
	ObjectiveZoneZoneTypeGood string = "Good"

	// ObjectiveZoneZoneTypeTarget captures enum value "Target"
	ObjectiveZoneZoneTypeTarget string = "Target"

	// ObjectiveZoneZoneTypeGreat captures enum value "Great"
	ObjectiveZoneZoneTypeGreat string = "Great"

	// ObjectiveZoneZoneTypeOut captures enum value "Out"
	ObjectiveZoneZoneTypeOut string = "Out"
)

// prop value enum
func (m *ObjectiveZone) validateZoneTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, objectiveZoneTypeZoneTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ObjectiveZone) validateZoneType(formats strfmt.Registry) error {

	if err := validate.Required("zoneType", "body", m.ZoneType); err != nil {
		return err
	}

	// value enum
	if err := m.validateZoneTypeEnum("zoneType", "body", *m.ZoneType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this objective zone based on context it is used
func (m *ObjectiveZone) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectiveZone) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectiveZone) UnmarshalBinary(b []byte) error {
	var res ObjectiveZone
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
