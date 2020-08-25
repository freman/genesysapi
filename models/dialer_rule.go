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

// DialerRule dialer rule
//
// swagger:model DialerRule
type DialerRule struct {

	// The list of actions to be taken if the conditions are true.
	Actions []*DialerAction `json:"actions"`

	// The category of the rule.
	// Required: true
	// Enum: [DIALER_PRECALL DIALER_WRAPUP]
	Category *string `json:"category"`

	// A list of Conditions. All of the Conditions must evaluate to true to trigger the actions.
	// Required: true
	Conditions []*Condition `json:"conditions"`

	// The identifier of the rule.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the rule.
	// Required: true
	Name *string `json:"name"`

	// The ranked order of the rule. Rules are processed from lowest number to highest.
	Order int32 `json:"order,omitempty"`
}

// Validate validates this dialer rule
func (m *DialerRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DialerRule) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {
		if swag.IsZero(m.Actions[i]) { // not required
			continue
		}

		if m.Actions[i] != nil {
			if err := m.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var dialerRuleTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIALER_PRECALL","DIALER_WRAPUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dialerRuleTypeCategoryPropEnum = append(dialerRuleTypeCategoryPropEnum, v)
	}
}

const (

	// DialerRuleCategoryDIALERPRECALL captures enum value "DIALER_PRECALL"
	DialerRuleCategoryDIALERPRECALL string = "DIALER_PRECALL"

	// DialerRuleCategoryDIALERWRAPUP captures enum value "DIALER_WRAPUP"
	DialerRuleCategoryDIALERWRAPUP string = "DIALER_WRAPUP"
)

// prop value enum
func (m *DialerRule) validateCategoryEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dialerRuleTypeCategoryPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DialerRule) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", *m.Category); err != nil {
		return err
	}

	return nil
}

func (m *DialerRule) validateConditions(formats strfmt.Registry) error {

	if err := validate.Required("conditions", "body", m.Conditions); err != nil {
		return err
	}

	for i := 0; i < len(m.Conditions); i++ {
		if swag.IsZero(m.Conditions[i]) { // not required
			continue
		}

		if m.Conditions[i] != nil {
			if err := m.Conditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *DialerRule) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DialerRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DialerRule) UnmarshalBinary(b []byte) error {
	var res DialerRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}