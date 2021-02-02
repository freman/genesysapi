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

// ScimPhoneNumber Defines a SCIM phone number.
//
// swagger:model ScimPhoneNumber
type ScimPhoneNumber struct {

	// Indicates whether the phone number is the primary phone number.
	Primary bool `json:"primary"`

	// The type of phone number.
	// Enum: [work work2 work3 work4 home mobile other microsoftteams zoomphone ringcentral]
	Type string `json:"type,omitempty"`

	// The phone number in E.164 or tel URI format, for example, tel:+nnnnnnnn; ext=xxxxx.
	Value string `json:"value,omitempty"`
}

// Validate validates this scim phone number
func (m *ScimPhoneNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scimPhoneNumberTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["work","work2","work3","work4","home","mobile","other","microsoftteams","zoomphone","ringcentral"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimPhoneNumberTypeTypePropEnum = append(scimPhoneNumberTypeTypePropEnum, v)
	}
}

const (

	// ScimPhoneNumberTypeWork captures enum value "work"
	ScimPhoneNumberTypeWork string = "work"

	// ScimPhoneNumberTypeWork2 captures enum value "work2"
	ScimPhoneNumberTypeWork2 string = "work2"

	// ScimPhoneNumberTypeWork3 captures enum value "work3"
	ScimPhoneNumberTypeWork3 string = "work3"

	// ScimPhoneNumberTypeWork4 captures enum value "work4"
	ScimPhoneNumberTypeWork4 string = "work4"

	// ScimPhoneNumberTypeHome captures enum value "home"
	ScimPhoneNumberTypeHome string = "home"

	// ScimPhoneNumberTypeMobile captures enum value "mobile"
	ScimPhoneNumberTypeMobile string = "mobile"

	// ScimPhoneNumberTypeOther captures enum value "other"
	ScimPhoneNumberTypeOther string = "other"

	// ScimPhoneNumberTypeMicrosoftteams captures enum value "microsoftteams"
	ScimPhoneNumberTypeMicrosoftteams string = "microsoftteams"

	// ScimPhoneNumberTypeZoomphone captures enum value "zoomphone"
	ScimPhoneNumberTypeZoomphone string = "zoomphone"

	// ScimPhoneNumberTypeRingcentral captures enum value "ringcentral"
	ScimPhoneNumberTypeRingcentral string = "ringcentral"
)

// prop value enum
func (m *ScimPhoneNumber) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimPhoneNumberTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimPhoneNumber) validateType(formats strfmt.Registry) error {

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
func (m *ScimPhoneNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimPhoneNumber) UnmarshalBinary(b []byte) error {
	var res ScimPhoneNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
