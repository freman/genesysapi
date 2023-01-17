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

// ValidationServiceRequest validation service request
//
// swagger:model ValidationServiceRequest
type ValidationServiceRequest struct {

	// The last day of the data you are importing. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	DateImportEnded *strfmt.DateTime `json:"dateImportEnded"`

	// S3 key for the uploaded file
	// Required: true
	UploadKey *string `json:"uploadKey"`
}

// Validate validates this validation service request
func (m *ValidationServiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateImportEnded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValidationServiceRequest) validateDateImportEnded(formats strfmt.Registry) error {

	if err := validate.Required("dateImportEnded", "body", m.DateImportEnded); err != nil {
		return err
	}

	if err := validate.FormatOf("dateImportEnded", "body", "date-time", m.DateImportEnded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ValidationServiceRequest) validateUploadKey(formats strfmt.Registry) error {

	if err := validate.Required("uploadKey", "body", m.UploadKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this validation service request based on context it is used
func (m *ValidationServiceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ValidationServiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValidationServiceRequest) UnmarshalBinary(b []byte) error {
	var res ValidationServiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
