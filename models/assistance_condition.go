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

// AssistanceCondition assistance condition
//
// swagger:model AssistanceCondition
type AssistanceCondition struct {

	// The operator for the assistance condition. The operator defines whether the listed topicIds should EXIST or NOTEXIST for the condition to be evaluated as true.
	// Enum: [EXISTS NOTEXISTS]
	Operator string `json:"operator,omitempty"`

	// List of topicIds within the assistance condition which would be combined together using logical OR operator. Eg ( topicId_1 || topicId_2 ) .
	TopicIds []string `json:"topicIds"`
}

// Validate validates this assistance condition
func (m *AssistanceCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperator(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var assistanceConditionTypeOperatorPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXISTS","NOTEXISTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		assistanceConditionTypeOperatorPropEnum = append(assistanceConditionTypeOperatorPropEnum, v)
	}
}

const (

	// AssistanceConditionOperatorEXISTS captures enum value "EXISTS"
	AssistanceConditionOperatorEXISTS string = "EXISTS"

	// AssistanceConditionOperatorNOTEXISTS captures enum value "NOTEXISTS"
	AssistanceConditionOperatorNOTEXISTS string = "NOTEXISTS"
)

// prop value enum
func (m *AssistanceCondition) validateOperatorEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, assistanceConditionTypeOperatorPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AssistanceCondition) validateOperator(formats strfmt.Registry) error {
	if swag.IsZero(m.Operator) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperatorEnum("operator", "body", m.Operator); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this assistance condition based on context it is used
func (m *AssistanceCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AssistanceCondition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssistanceCondition) UnmarshalBinary(b []byte) error {
	var res AssistanceCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
