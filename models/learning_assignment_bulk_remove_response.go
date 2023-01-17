// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// LearningAssignmentBulkRemoveResponse learning assignment bulk remove response
//
// swagger:model LearningAssignmentBulkRemoveResponse
type LearningAssignmentBulkRemoveResponse struct {

	// The learning assignments that were not removed due to missing permissions
	DisallowedEntities []*DisallowedEntityLearningAssignmentReference `json:"disallowedEntities"`

	// The learning assignments that were removed successfully
	Entities []*LearningAssignmentEntity `json:"entities"`
}

// Validate validates this learning assignment bulk remove response
func (m *LearningAssignmentBulkRemoveResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisallowedEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningAssignmentBulkRemoveResponse) validateDisallowedEntities(formats strfmt.Registry) error {
	if swag.IsZero(m.DisallowedEntities) { // not required
		return nil
	}

	for i := 0; i < len(m.DisallowedEntities); i++ {
		if swag.IsZero(m.DisallowedEntities[i]) { // not required
			continue
		}

		if m.DisallowedEntities[i] != nil {
			if err := m.DisallowedEntities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disallowedEntities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disallowedEntities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningAssignmentBulkRemoveResponse) validateEntities(formats strfmt.Registry) error {
	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this learning assignment bulk remove response based on the context it is used
func (m *LearningAssignmentBulkRemoveResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDisallowedEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEntities(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LearningAssignmentBulkRemoveResponse) contextValidateDisallowedEntities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DisallowedEntities); i++ {

		if m.DisallowedEntities[i] != nil {
			if err := m.DisallowedEntities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("disallowedEntities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("disallowedEntities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LearningAssignmentBulkRemoveResponse) contextValidateEntities(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Entities); i++ {

		if m.Entities[i] != nil {
			if err := m.Entities[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LearningAssignmentBulkRemoveResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningAssignmentBulkRemoveResponse) UnmarshalBinary(b []byte) error {
	var res LearningAssignmentBulkRemoveResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
