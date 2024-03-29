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

// ImportScriptStatusResponse import script status response
//
// swagger:model ImportScriptStatusResponse
type ImportScriptStatusResponse struct {

	// message
	Message string `json:"message,omitempty"`

	// succeeded
	Succeeded bool `json:"succeeded"`

	// url
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this import script status response
func (m *ImportScriptStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ImportScriptStatusResponse) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this import script status response based on context it is used
func (m *ImportScriptStatusResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ImportScriptStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ImportScriptStatusResponse) UnmarshalBinary(b []byte) error {
	var res ImportScriptStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
