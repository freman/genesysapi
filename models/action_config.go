// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActionConfig Defines components of the Action Config.
//
// swagger:model ActionConfig
type ActionConfig struct {

	// Configuration of outbound request.
	Request *RequestConfig `json:"request,omitempty"`

	// Configuration of response processing.
	Response *ResponseConfig `json:"response,omitempty"`
}

// Validate validates this action config
func (m *ActionConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionConfig) validateRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.Request) { // not required
		return nil
	}

	if m.Request != nil {
		if err := m.Request.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("request")
			}
			return err
		}
	}

	return nil
}

func (m *ActionConfig) validateResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.Response) { // not required
		return nil
	}

	if m.Response != nil {
		if err := m.Response.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("response")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionConfig) UnmarshalBinary(b []byte) error {
	var res ActionConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
