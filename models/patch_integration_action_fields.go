// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PatchIntegrationActionFields patch integration action fields
//
// swagger:model PatchIntegrationActionFields
type PatchIntegrationActionFields struct {

	// Reference to the Integration Action to be used when integrationAction type is qualified
	IntegrationAction *PatchIntegrationAction `json:"integrationAction,omitempty"`

	// Collection of Request Mappings to use
	RequestMappings []*RequestMapping `json:"requestMappings"`
}

// Validate validates this patch integration action fields
func (m *PatchIntegrationActionFields) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIntegrationAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestMappings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchIntegrationActionFields) validateIntegrationAction(formats strfmt.Registry) error {

	if swag.IsZero(m.IntegrationAction) { // not required
		return nil
	}

	if m.IntegrationAction != nil {
		if err := m.IntegrationAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrationAction")
			}
			return err
		}
	}

	return nil
}

func (m *PatchIntegrationActionFields) validateRequestMappings(formats strfmt.Registry) error {

	if swag.IsZero(m.RequestMappings) { // not required
		return nil
	}

	for i := 0; i < len(m.RequestMappings); i++ {
		if swag.IsZero(m.RequestMappings[i]) { // not required
			continue
		}

		if m.RequestMappings[i] != nil {
			if err := m.RequestMappings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("requestMappings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchIntegrationActionFields) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchIntegrationActionFields) UnmarshalBinary(b []byte) error {
	var res PatchIntegrationActionFields
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
