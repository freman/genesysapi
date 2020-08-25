// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainPermissionPolicy domain permission policy
//
// swagger:model DomainPermissionPolicy
type DomainPermissionPolicy struct {

	// action set
	// Unique: true
	ActionSet []string `json:"actionSet"`

	// allow conditions
	AllowConditions bool `json:"allowConditions"`

	// domain
	Domain string `json:"domain,omitempty"`

	// entity name
	EntityName string `json:"entityName,omitempty"`

	// named resources
	// Unique: true
	NamedResources []string `json:"namedResources"`

	// policy description
	PolicyDescription string `json:"policyDescription,omitempty"`

	// policy name
	PolicyName string `json:"policyName,omitempty"`

	// resource condition node
	ResourceConditionNode *DomainResourceConditionNode `json:"resourceConditionNode,omitempty"`
}

// Validate validates this domain permission policy
func (m *DomainPermissionPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamedResources(formats); err != nil {
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

func (m *DomainPermissionPolicy) validateActionSet(formats strfmt.Registry) error {

	if swag.IsZero(m.ActionSet) { // not required
		return nil
	}

	if err := validate.UniqueItems("actionSet", "body", m.ActionSet); err != nil {
		return err
	}

	return nil
}

func (m *DomainPermissionPolicy) validateNamedResources(formats strfmt.Registry) error {

	if swag.IsZero(m.NamedResources) { // not required
		return nil
	}

	if err := validate.UniqueItems("namedResources", "body", m.NamedResources); err != nil {
		return err
	}

	return nil
}

func (m *DomainPermissionPolicy) validateResourceConditionNode(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceConditionNode) { // not required
		return nil
	}

	if m.ResourceConditionNode != nil {
		if err := m.ResourceConditionNode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resourceConditionNode")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainPermissionPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainPermissionPolicy) UnmarshalBinary(b []byte) error {
	var res DomainPermissionPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}