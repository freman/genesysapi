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

// LearningAssignmentItem learning assignment item
//
// swagger:model LearningAssignmentItem
type LearningAssignmentItem struct {

	// The Learning Module ID associated with this assignment
	// Required: true
	ModuleID *string `json:"moduleId"`

	// The User ID associated with this assignment
	// Required: true
	UserID *string `json:"userId"`
}

// Validate validates this learning assignment item
func (m *LearningAssignmentItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateModuleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningAssignmentItem) validateModuleID(formats strfmt.Registry) error {

	if err := validate.Required("moduleId", "body", m.ModuleID); err != nil {
		return err
	}

	return nil
}

func (m *LearningAssignmentItem) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("userId", "body", m.UserID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningAssignmentItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningAssignmentItem) UnmarshalBinary(b []byte) error {
	var res LearningAssignmentItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
