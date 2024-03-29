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

// UploadURLRequestBody upload Url request body
//
// swagger:model UploadUrlRequestBody
type UploadURLRequestBody struct {

	// The expected content length (in bytes) of the gzip-encoded data that will be PUT to the returned signed URL
	// Required: true
	ContentLengthBytes *int64 `json:"contentLengthBytes"`
}

// Validate validates this upload Url request body
func (m *UploadURLRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContentLengthBytes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UploadURLRequestBody) validateContentLengthBytes(formats strfmt.Registry) error {

	if err := validate.Required("contentLengthBytes", "body", m.ContentLengthBytes); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this upload Url request body based on context it is used
func (m *UploadURLRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UploadURLRequestBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UploadURLRequestBody) UnmarshalBinary(b []byte) error {
	var res UploadURLRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
