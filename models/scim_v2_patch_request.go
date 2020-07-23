// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScimV2PatchRequest Defines a SCIM PATCH request. See section 3.5.2 "Modifying with PATCH" in RFC 7644 for details.
//
// swagger:model ScimV2PatchRequest
type ScimV2PatchRequest struct {

	// The list of operations to perform for the PATCH request.
	Operations []*ScimV2PatchOperation `json:"Operations"`

	// The list of schemas used in the PATCH request.
	// Required: true
	Schemas []string `json:"schemas"`
}

// Validate validates this scim v2 patch request
func (m *ScimV2PatchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOperations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchemas(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimV2PatchRequest) validateOperations(formats strfmt.Registry) error {

	if swag.IsZero(m.Operations) { // not required
		return nil
	}

	for i := 0; i < len(m.Operations); i++ {
		if swag.IsZero(m.Operations[i]) { // not required
			continue
		}

		if m.Operations[i] != nil {
			if err := m.Operations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Operations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScimV2PatchRequest) validateSchemas(formats strfmt.Registry) error {

	if err := validate.Required("schemas", "body", m.Schemas); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimV2PatchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimV2PatchRequest) UnmarshalBinary(b []byte) error {
	var res ScimV2PatchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
