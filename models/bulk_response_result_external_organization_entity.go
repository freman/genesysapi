// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BulkResponseResultExternalOrganizationEntity bulk response result external organization entity
//
// swagger:model BulkResponseResultExternalOrganizationEntity
type BulkResponseResultExternalOrganizationEntity struct {

	// entity
	Entity *ExternalOrganization `json:"entity,omitempty"`

	// error
	Error *BulkErrorEntity `json:"error,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// success
	Success bool `json:"success"`
}

// Validate validates this bulk response result external organization entity
func (m *BulkResponseResultExternalOrganizationEntity) Validate(formats strfmt.Registry) error {
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

func (m *BulkResponseResultExternalOrganizationEntity) validateEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *BulkResponseResultExternalOrganizationEntity) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bulk response result external organization entity based on the context it is used
func (m *BulkResponseResultExternalOrganizationEntity) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkResponseResultExternalOrganizationEntity) contextValidateEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.Entity != nil {
		if err := m.Entity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

func (m *BulkResponseResultExternalOrganizationEntity) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkResponseResultExternalOrganizationEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkResponseResultExternalOrganizationEntity) UnmarshalBinary(b []byte) error {
	var res BulkResponseResultExternalOrganizationEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
