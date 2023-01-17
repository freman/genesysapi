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

// Referrer referrer
//
// swagger:model Referrer
type Referrer struct {

	// Referrer URL domain.
	// Required: true
	Domain *string `json:"domain"`

	// Referrer URL fragment.
	Fragment string `json:"fragment,omitempty"`

	// Referrer URL hostname.
	// Required: true
	Hostname *string `json:"hostname"`

	// Referrer keywords.
	Keywords string `json:"keywords,omitempty"`

	// Type of referrer (e.g. search, social).
	// Enum: [internal search social email unknown paid]
	Medium string `json:"medium,omitempty"`

	// Name of referrer (e.g. Yahoo!, Google, InfoSpace).
	Name string `json:"name,omitempty"`

	// Referrer URL pathname.
	// Required: true
	Pathname *string `json:"pathname"`

	// Referrer URL querystring.
	QueryString string `json:"queryString,omitempty"`

	// Referrer URL.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this referrer
func (m *Referrer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePathname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Referrer) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *Referrer) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	return nil
}

var referrerTypeMediumPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internal","search","social","email","unknown","paid"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		referrerTypeMediumPropEnum = append(referrerTypeMediumPropEnum, v)
	}
}

const (

	// ReferrerMediumInternal captures enum value "internal"
	ReferrerMediumInternal string = "internal"

	// ReferrerMediumSearch captures enum value "search"
	ReferrerMediumSearch string = "search"

	// ReferrerMediumSocial captures enum value "social"
	ReferrerMediumSocial string = "social"

	// ReferrerMediumEmail captures enum value "email"
	ReferrerMediumEmail string = "email"

	// ReferrerMediumUnknown captures enum value "unknown"
	ReferrerMediumUnknown string = "unknown"

	// ReferrerMediumPaid captures enum value "paid"
	ReferrerMediumPaid string = "paid"
)

// prop value enum
func (m *Referrer) validateMediumEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, referrerTypeMediumPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Referrer) validateMedium(formats strfmt.Registry) error {
	if swag.IsZero(m.Medium) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediumEnum("medium", "body", m.Medium); err != nil {
		return err
	}

	return nil
}

func (m *Referrer) validatePathname(formats strfmt.Registry) error {

	if err := validate.Required("pathname", "body", m.Pathname); err != nil {
		return err
	}

	return nil
}

func (m *Referrer) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("url", "body", m.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this referrer based on context it is used
func (m *Referrer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Referrer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Referrer) UnmarshalBinary(b []byte) error {
	var res Referrer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
