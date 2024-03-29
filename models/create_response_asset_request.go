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

// CreateResponseAssetRequest create response asset request
//
// swagger:model CreateResponseAssetRequest
type CreateResponseAssetRequest struct {

	// Content MD-5 of the file to upload
	ContentMd5 string `json:"contentMd5,omitempty"`

	// Division to associate to this asset. Can only be used with this division.
	DivisionID string `json:"divisionId,omitempty"`

	// Name of the file to upload. It must not start with a dot and not end with a forward slash. Whitespace and the following characters are not allowed: \{^}%`]">[~<#|
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this create response asset request
func (m *CreateResponseAssetRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateResponseAssetRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create response asset request based on context it is used
func (m *CreateResponseAssetRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateResponseAssetRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateResponseAssetRequest) UnmarshalBinary(b []byte) error {
	var res CreateResponseAssetRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
