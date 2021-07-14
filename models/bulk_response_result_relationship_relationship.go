// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BulkResponseResultRelationshipRelationship bulk response result relationship relationship
//
// swagger:model BulkResponseResultRelationshipRelationship
type BulkResponseResultRelationshipRelationship struct {

	// entity
	Entity *Relationship `json:"entity,omitempty"`

	// error
	Error *BulkErrorRelationship `json:"error,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// success
	Success bool `json:"success"`
}

// Validate validates this bulk response result relationship relationship
func (m *BulkResponseResultRelationshipRelationship) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkResponseResultRelationshipRelationship) validateEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *BulkResponseResultRelationshipRelationship) validateError(formats strfmt.Registry) error {

	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkResponseResultRelationshipRelationship) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkResponseResultRelationshipRelationship) UnmarshalBinary(b []byte) error {
	var res BulkResponseResultRelationshipRelationship
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
