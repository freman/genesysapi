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

// KnowledgeImportJobReport knowledge import job report
//
// swagger:model KnowledgeImportJobReport
type KnowledgeImportJobReport struct {

	// List of errors occurred during processing import.
	Errors []*KnowledgeImportJobError `json:"errors"`

	// Statistics related to the import job.
	Statistics *KnowledgeImportJobStatistics `json:"statistics,omitempty"`
}

// Validate validates this knowledge import job report
func (m *KnowledgeImportJobReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeImportJobReport) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KnowledgeImportJobReport) validateStatistics(formats strfmt.Registry) error {
	if swag.IsZero(m.Statistics) { // not required
		return nil
	}

	if m.Statistics != nil {
		if err := m.Statistics.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this knowledge import job report based on the context it is used
func (m *KnowledgeImportJobReport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatistics(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeImportJobReport) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *KnowledgeImportJobReport) contextValidateStatistics(ctx context.Context, formats strfmt.Registry) error {

	if m.Statistics != nil {
		if err := m.Statistics.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("statistics")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("statistics")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeImportJobReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeImportJobReport) UnmarshalBinary(b []byte) error {
	var res KnowledgeImportJobReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
