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

// PatchActionMapScheduleGroups patch action map schedule groups
//
// swagger:model PatchActionMapScheduleGroups
type PatchActionMapScheduleGroups struct {

	// The actions map's associated schedule group.
	// Required: true
	ActionMapScheduleGroup *ActionMapScheduleGroup `json:"actionMapScheduleGroup"`

	// The action map's associated emergency schedule group.
	EmergencyActionMapScheduleGroup *ActionMapScheduleGroup `json:"emergencyActionMapScheduleGroup,omitempty"`
}

// Validate validates this patch action map schedule groups
func (m *PatchActionMapScheduleGroups) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActionMapScheduleGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmergencyActionMapScheduleGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchActionMapScheduleGroups) validateActionMapScheduleGroup(formats strfmt.Registry) error {

	if err := validate.Required("actionMapScheduleGroup", "body", m.ActionMapScheduleGroup); err != nil {
		return err
	}

	if m.ActionMapScheduleGroup != nil {
		if err := m.ActionMapScheduleGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actionMapScheduleGroup")
			}
			return err
		}
	}

	return nil
}

func (m *PatchActionMapScheduleGroups) validateEmergencyActionMapScheduleGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.EmergencyActionMapScheduleGroup) { // not required
		return nil
	}

	if m.EmergencyActionMapScheduleGroup != nil {
		if err := m.EmergencyActionMapScheduleGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("emergencyActionMapScheduleGroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchActionMapScheduleGroups) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchActionMapScheduleGroups) UnmarshalBinary(b []byte) error {
	var res PatchActionMapScheduleGroups
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
