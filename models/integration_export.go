// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IntegrationExport integration export
//
// swagger:model IntegrationExport
type IntegrationExport struct {

	// The aws-s3-recording-bulk-actions-integration that the policy uses for exports.
	// Required: true
	Integration *DomainEntityRef `json:"integration"`

	// True if the policy should export screen recordings in addition to the other conversation media. Default = true
	ShouldExportScreenRecordings bool `json:"shouldExportScreenRecordings"`
}

// Validate validates this integration export
func (m *IntegrationExport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationExport) validateIntegration(formats strfmt.Registry) error {

	if err := validate.Required("integration", "body", m.Integration); err != nil {
		return err
	}

	if m.Integration != nil {
		if err := m.Integration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this integration export based on the context it is used
func (m *IntegrationExport) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIntegration(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IntegrationExport) contextValidateIntegration(ctx context.Context, formats strfmt.Registry) error {

	if m.Integration != nil {
		if err := m.Integration.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integration")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("integration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IntegrationExport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IntegrationExport) UnmarshalBinary(b []byte) error {
	var res IntegrationExport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
