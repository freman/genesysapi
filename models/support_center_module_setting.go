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

// SupportCenterModuleSetting support center module setting
//
// swagger:model SupportCenterModuleSetting
type SupportCenterModuleSetting struct {

	// Whether or not support center screen module is enabled
	Enabled bool `json:"enabled"`

	// Screen module type
	// Enum: [Search Categories FAQ Contact Results Article]
	Type string `json:"type,omitempty"`
}

// Validate validates this support center module setting
func (m *SupportCenterModuleSetting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var supportCenterModuleSettingTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Search","Categories","FAQ","Contact","Results","Article"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		supportCenterModuleSettingTypeTypePropEnum = append(supportCenterModuleSettingTypeTypePropEnum, v)
	}
}

const (

	// SupportCenterModuleSettingTypeSearch captures enum value "Search"
	SupportCenterModuleSettingTypeSearch string = "Search"

	// SupportCenterModuleSettingTypeCategories captures enum value "Categories"
	SupportCenterModuleSettingTypeCategories string = "Categories"

	// SupportCenterModuleSettingTypeFAQ captures enum value "FAQ"
	SupportCenterModuleSettingTypeFAQ string = "FAQ"

	// SupportCenterModuleSettingTypeContact captures enum value "Contact"
	SupportCenterModuleSettingTypeContact string = "Contact"

	// SupportCenterModuleSettingTypeResults captures enum value "Results"
	SupportCenterModuleSettingTypeResults string = "Results"

	// SupportCenterModuleSettingTypeArticle captures enum value "Article"
	SupportCenterModuleSettingTypeArticle string = "Article"
)

// prop value enum
func (m *SupportCenterModuleSetting) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, supportCenterModuleSettingTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SupportCenterModuleSetting) validateType(formats strfmt.Registry) error {

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
func (m *SupportCenterModuleSetting) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportCenterModuleSetting) UnmarshalBinary(b []byte) error {
	var res SupportCenterModuleSetting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
