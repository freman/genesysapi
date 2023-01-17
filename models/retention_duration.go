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

// RetentionDuration retention duration
//
// swagger:model RetentionDuration
type RetentionDuration struct {

	// archive retention
	ArchiveRetention *ArchiveRetention `json:"archiveRetention,omitempty"`

	// delete retention
	DeleteRetention *DeleteRetention `json:"deleteRetention,omitempty"`
}

// Validate validates this retention duration
func (m *RetentionDuration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArchiveRetention(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleteRetention(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetentionDuration) validateArchiveRetention(formats strfmt.Registry) error {
	if swag.IsZero(m.ArchiveRetention) { // not required
		return nil
	}

	if m.ArchiveRetention != nil {
		if err := m.ArchiveRetention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archiveRetention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archiveRetention")
			}
			return err
		}
	}

	return nil
}

func (m *RetentionDuration) validateDeleteRetention(formats strfmt.Registry) error {
	if swag.IsZero(m.DeleteRetention) { // not required
		return nil
	}

	if m.DeleteRetention != nil {
		if err := m.DeleteRetention.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteRetention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteRetention")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this retention duration based on the context it is used
func (m *RetentionDuration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateArchiveRetention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDeleteRetention(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RetentionDuration) contextValidateArchiveRetention(ctx context.Context, formats strfmt.Registry) error {

	if m.ArchiveRetention != nil {
		if err := m.ArchiveRetention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("archiveRetention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("archiveRetention")
			}
			return err
		}
	}

	return nil
}

func (m *RetentionDuration) contextValidateDeleteRetention(ctx context.Context, formats strfmt.Registry) error {

	if m.DeleteRetention != nil {
		if err := m.DeleteRetention.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("deleteRetention")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("deleteRetention")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RetentionDuration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RetentionDuration) UnmarshalBinary(b []byte) error {
	var res RetentionDuration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
