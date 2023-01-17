// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReplaceResponse replace response
//
// swagger:model ReplaceResponse
type ReplaceResponse struct {

	// change number
	ChangeNumber int32 `json:"changeNumber,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// upload destination Uri
	// Format: uri
	UploadDestinationURI strfmt.URI `json:"uploadDestinationUri,omitempty"`

	// upload method
	// Enum: [SINGLE_PUT MULTIPART_POST]
	UploadMethod string `json:"uploadMethod,omitempty"`

	// upload status
	UploadStatus *DomainEntityRef `json:"uploadStatus,omitempty"`
}

// Validate validates this replace response
func (m *ReplaceResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUploadDestinationURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUploadStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceResponse) validateUploadDestinationURI(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadDestinationURI) { // not required
		return nil
	}

	if err := validate.FormatOf("uploadDestinationUri", "body", "uri", m.UploadDestinationURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var replaceResponseTypeUploadMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SINGLE_PUT","MULTIPART_POST"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		replaceResponseTypeUploadMethodPropEnum = append(replaceResponseTypeUploadMethodPropEnum, v)
	}
}

const (

	// ReplaceResponseUploadMethodSINGLEPUT captures enum value "SINGLE_PUT"
	ReplaceResponseUploadMethodSINGLEPUT string = "SINGLE_PUT"

	// ReplaceResponseUploadMethodMULTIPARTPOST captures enum value "MULTIPART_POST"
	ReplaceResponseUploadMethodMULTIPARTPOST string = "MULTIPART_POST"
)

// prop value enum
func (m *ReplaceResponse) validateUploadMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, replaceResponseTypeUploadMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReplaceResponse) validateUploadMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateUploadMethodEnum("uploadMethod", "body", m.UploadMethod); err != nil {
		return err
	}

	return nil
}

func (m *ReplaceResponse) validateUploadStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.UploadStatus) { // not required
		return nil
	}

	if m.UploadStatus != nil {
		if err := m.UploadStatus.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadStatus")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this replace response based on the context it is used
func (m *ReplaceResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUploadStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReplaceResponse) contextValidateUploadStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.UploadStatus != nil {
		if err := m.UploadStatus.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uploadStatus")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("uploadStatus")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReplaceResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReplaceResponse) UnmarshalBinary(b []byte) error {
	var res ReplaceResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
