// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourcePermissionPolicy resource permission policy
//
// swagger:model ResourcePermissionPolicy
type ResourcePermissionPolicy struct {

	// action set
	// Unique: true
	ActionSet []string `json:"actionSet"`

	// action set key
	ActionSetKey string `json:"actionSetKey,omitempty"`

	// allow conditions
	AllowConditions bool `json:"allowConditions"`

	// domain
	Domain string `json:"domain,omitempty"`

	// entity name
	EntityName string `json:"entityName,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// named resources
	NamedResources []string `json:"namedResources"`

	// policy description
	PolicyDescription string `json:"policyDescription,omitempty"`

	// policy name
	PolicyName string `json:"policyName,omitempty"`

	// resource condition
	ResourceCondition string `json:"resourceCondition,omitempty"`

	// resource condition node
	ResourceConditionNode *ResourceConditionNode `json:"resourceConditionNode,omitempty"`
}

// Validate validates this resource permission policy
func (m *ResourcePermissionPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceConditionNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcePermissionPolicy) validateActionSet(formats strfmt.Registry) error {
	if swag.IsZero(m.ActionSet) { // not required
		return nil
	}

	if err := validate.UniqueItems("actionSet", "body", m.ActionSet); err != nil {
		return err
	}

	return nil
}

func (m *ResourcePermissionPolicy) validateResourceConditionNode(formats strfmt.Registry) error {
	if swag.IsZero(m.ResourceConditionNode) { // not required
		return nil
	}

	if m.ResourceConditionNode != nil {
		if err := m.ResourceConditionNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceConditionNode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceConditionNode")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this resource permission policy based on the context it is used
func (m *ResourcePermissionPolicy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateResourceConditionNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ResourcePermissionPolicy) contextValidateResourceConditionNode(ctx context.Context, formats strfmt.Registry) error {

	if m.ResourceConditionNode != nil {
		if err := m.ResourceConditionNode.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceConditionNode")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("resourceConditionNode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourcePermissionPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourcePermissionPolicy) UnmarshalBinary(b []byte) error {
	var res ResourcePermissionPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
