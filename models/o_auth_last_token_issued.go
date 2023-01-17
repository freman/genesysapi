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

// OAuthLastTokenIssued o auth last token issued
//
// swagger:model OAuthLastTokenIssued
type OAuthLastTokenIssued struct {

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateIssued strfmt.DateTime `json:"dateIssued,omitempty"`
}

// Validate validates this o auth last token issued
func (m *OAuthLastTokenIssued) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateIssued(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OAuthLastTokenIssued) validateDateIssued(formats strfmt.Registry) error {
	if swag.IsZero(m.DateIssued) { // not required
		return nil
	}

	if err := validate.FormatOf("dateIssued", "body", "date-time", m.DateIssued.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this o auth last token issued based on context it is used
func (m *OAuthLastTokenIssued) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OAuthLastTokenIssued) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OAuthLastTokenIssued) UnmarshalBinary(b []byte) error {
	var res OAuthLastTokenIssued
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
