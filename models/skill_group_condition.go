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

// SkillGroupCondition skill group condition
//
// swagger:model SkillGroupCondition
type SkillGroupCondition struct {

	// Routing skill conditions that will be used for building the query
	// Required: true
	LanguageSkillConditions []*SkillGroupLanguageCondition `json:"languageSkillConditions"`

	// Operator that will be applied to the conditions
	// Required: true
	// Enum: [And Not Or]
	Operation *string `json:"operation"`

	// Routing skill conditions that will be used for building the query
	// Required: true
	RoutingSkillConditions []*SkillGroupRoutingCondition `json:"routingSkillConditions"`
}

// Validate validates this skill group condition
func (m *SkillGroupCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLanguageSkillConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingSkillConditions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SkillGroupCondition) validateLanguageSkillConditions(formats strfmt.Registry) error {

	if err := validate.Required("languageSkillConditions", "body", m.LanguageSkillConditions); err != nil {
		return err
	}

	for i := 0; i < len(m.LanguageSkillConditions); i++ {
		if swag.IsZero(m.LanguageSkillConditions[i]) { // not required
			continue
		}

		if m.LanguageSkillConditions[i] != nil {
			if err := m.LanguageSkillConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languageSkillConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languageSkillConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var skillGroupConditionTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["And","Not","Or"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		skillGroupConditionTypeOperationPropEnum = append(skillGroupConditionTypeOperationPropEnum, v)
	}
}

const (

	// SkillGroupConditionOperationAnd captures enum value "And"
	SkillGroupConditionOperationAnd string = "And"

	// SkillGroupConditionOperationNot captures enum value "Not"
	SkillGroupConditionOperationNot string = "Not"

	// SkillGroupConditionOperationOr captures enum value "Or"
	SkillGroupConditionOperationOr string = "Or"
)

// prop value enum
func (m *SkillGroupCondition) validateOperationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, skillGroupConditionTypeOperationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SkillGroupCondition) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", *m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *SkillGroupCondition) validateRoutingSkillConditions(formats strfmt.Registry) error {

	if err := validate.Required("routingSkillConditions", "body", m.RoutingSkillConditions); err != nil {
		return err
	}

	for i := 0; i < len(m.RoutingSkillConditions); i++ {
		if swag.IsZero(m.RoutingSkillConditions[i]) { // not required
			continue
		}

		if m.RoutingSkillConditions[i] != nil {
			if err := m.RoutingSkillConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routingSkillConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("routingSkillConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this skill group condition based on the context it is used
func (m *SkillGroupCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLanguageSkillConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRoutingSkillConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SkillGroupCondition) contextValidateLanguageSkillConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.LanguageSkillConditions); i++ {

		if m.LanguageSkillConditions[i] != nil {
			if err := m.LanguageSkillConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languageSkillConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languageSkillConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SkillGroupCondition) contextValidateRoutingSkillConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RoutingSkillConditions); i++ {

		if m.RoutingSkillConditions[i] != nil {
			if err := m.RoutingSkillConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routingSkillConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("routingSkillConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SkillGroupCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SkillGroupCondition) UnmarshalBinary(b []byte) error {
	var res SkillGroupCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}