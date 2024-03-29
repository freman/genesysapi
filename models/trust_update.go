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

// TrustUpdate trust update
//
// swagger:model TrustUpdate
type TrustUpdate struct {

	// The expiration date of the trust. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateExpired strfmt.DateTime `json:"dateExpired,omitempty"`

	// If disabled no trustee user will have access, even if they were previously added.
	// Required: true
	Enabled *bool `json:"enabled"`
}

// Validate validates this trust update
func (m *TrustUpdate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateExpired(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrustUpdate) validateDateExpired(formats strfmt.Registry) error {
	if swag.IsZero(m.DateExpired) { // not required
		return nil
	}

	if err := validate.FormatOf("dateExpired", "body", "date-time", m.DateExpired.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrustUpdate) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this trust update based on context it is used
func (m *TrustUpdate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TrustUpdate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustUpdate) UnmarshalBinary(b []byte) error {
	var res TrustUpdate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
