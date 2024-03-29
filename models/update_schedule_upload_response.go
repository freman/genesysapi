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

// UpdateScheduleUploadResponse update schedule upload response
//
// swagger:model UpdateScheduleUploadResponse
type UpdateScheduleUploadResponse struct {

	// Required headers for the PUT request to the url
	Headers map[string]string `json:"headers,omitempty"`

	// Always null. Defines the schema of the json body to be PUT to the url. The json body should be gzip encoded before uploading
	UploadBodySchema *UpdateScheduleUploadSchema `json:"uploadBodySchema,omitempty"`

	// The key to pass to the secondary request to start processing of the upload
	UploadKey string `json:"uploadKey,omitempty"`

	// The url to which to PUT the upload body
	URL string `json:"url,omitempty"`
}

// Validate validates this update schedule upload response
func (m *UpdateScheduleUploadResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUploadBodySchema(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateScheduleUploadResponse) validateUploadBodySchema(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadBodySchema) { // not required
		return nil
	}

	if m.UploadBodySchema != nil {
		if err := m.UploadBodySchema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadBodySchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadBodySchema")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update schedule upload response based on the context it is used
func (m *UpdateScheduleUploadResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUploadBodySchema(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UpdateScheduleUploadResponse) contextValidateUploadBodySchema(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadBodySchema != nil {
		if err := m.UploadBodySchema.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadBodySchema")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadBodySchema")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UpdateScheduleUploadResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateScheduleUploadResponse) UnmarshalBinary(b []byte) error {
	var res UpdateScheduleUploadResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
