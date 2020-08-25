// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Contact contact
//
// swagger:model Contact
type Contact struct {

	// Email address or phone number for this contact type
	Address string `json:"address,omitempty"`

	// country code
	CountryCode string `json:"countryCode,omitempty"`

	// Formatted version of the address property
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Use internal extension instead of address. Mutually exclusive with the address field.
	Extension string `json:"extension,omitempty"`

	// media type
	// Enum: [PHONE EMAIL SMS]
	MediaType string `json:"mediaType,omitempty"`

	// type
	// Enum: [PRIMARY WORK WORK2 WORK3 WORK4 HOME MOBILE MAIN]
	Type string `json:"type,omitempty"`
}

// Validate validates this contact
func (m *Contact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var contactTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PHONE","EMAIL","SMS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactTypeMediaTypePropEnum = append(contactTypeMediaTypePropEnum, v)
	}
}

const (

	// ContactMediaTypePHONE captures enum value "PHONE"
	ContactMediaTypePHONE string = "PHONE"

	// ContactMediaTypeEMAIL captures enum value "EMAIL"
	ContactMediaTypeEMAIL string = "EMAIL"

	// ContactMediaTypeSMS captures enum value "SMS"
	ContactMediaTypeSMS string = "SMS"
)

// prop value enum
func (m *Contact) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contactTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Contact) validateMediaType(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

var contactTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRIMARY","WORK","WORK2","WORK3","WORK4","HOME","MOBILE","MAIN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactTypeTypePropEnum = append(contactTypeTypePropEnum, v)
	}
}

const (

	// ContactTypePRIMARY captures enum value "PRIMARY"
	ContactTypePRIMARY string = "PRIMARY"

	// ContactTypeWORK captures enum value "WORK"
	ContactTypeWORK string = "WORK"

	// ContactTypeWORK2 captures enum value "WORK2"
	ContactTypeWORK2 string = "WORK2"

	// ContactTypeWORK3 captures enum value "WORK3"
	ContactTypeWORK3 string = "WORK3"

	// ContactTypeWORK4 captures enum value "WORK4"
	ContactTypeWORK4 string = "WORK4"

	// ContactTypeHOME captures enum value "HOME"
	ContactTypeHOME string = "HOME"

	// ContactTypeMOBILE captures enum value "MOBILE"
	ContactTypeMOBILE string = "MOBILE"

	// ContactTypeMAIN captures enum value "MAIN"
	ContactTypeMAIN string = "MAIN"
)

// prop value enum
func (m *Contact) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contactTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Contact) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Contact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Contact) UnmarshalBinary(b []byte) error {
	var res Contact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
