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

// KnowledgeExportJobRequest knowledge export job request
//
// swagger:model KnowledgeExportJobRequest
type KnowledgeExportJobRequest struct {

	// What to export.
	// Required: true
	ExportFilter *KnowledgeExportJobFilter `json:"exportFilter"`
}

// Validate validates this knowledge export job request
func (m *KnowledgeExportJobRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExportFilter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KnowledgeExportJobRequest) validateExportFilter(formats strfmt.Registry) error {

	if err := validate.Required("exportFilter", "body", m.ExportFilter); err != nil {
		return err
	}

	if m.ExportFilter != nil {
		if err := m.ExportFilter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exportFilter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *KnowledgeExportJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KnowledgeExportJobRequest) UnmarshalBinary(b []byte) error {
	var res KnowledgeExportJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}