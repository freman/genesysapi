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

// UpdatePlanningGroupRequest update planning group request
//
// swagger:model UpdatePlanningGroupRequest
type UpdatePlanningGroupRequest struct {

	// Version metadata for the planning group
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`

	// The name of the planning group
	Name string `json:"name,omitempty"`

	// Set of route paths to associate with the planning group
	RoutePaths *SetWrapperRoutePathRequest `json:"routePaths,omitempty"`

	// The ID of the service goal template to associate with this planning group
	ServiceGoalTemplateID string `json:"serviceGoalTemplateId,omitempty"`
}

// Validate validates this update planning group request
func (m *UpdatePlanningGroupRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutePaths(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdatePlanningGroupRequest) validateMetadata(formats strfmt.Registry) error {

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

func (m *UpdatePlanningGroupRequest) validateRoutePaths(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutePaths) { // not required
		return nil
	}

	if m.RoutePaths != nil {
		if err := m.RoutePaths.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("routePaths")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdatePlanningGroupRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdatePlanningGroupRequest) UnmarshalBinary(b []byte) error {
	var res UpdatePlanningGroupRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
