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

// ExternalDataSource Describes a link to a record in an external system that contributed data to a Relate record
//
// swagger:model ExternalDataSource
type ExternalDataSource struct {

	// The platform that was the source of the data.  Example: a CRM like SALESFORCE.
	// Enum: [SALESFORCE]
	Platform string `json:"platform,omitempty"`

	// An URL that links to the source record that contributed data to the associated entity.
	URL string `json:"url,omitempty"`
}

// Validate validates this external data source
func (m *ExternalDataSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePlatform(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var externalDataSourceTypePlatformPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SALESFORCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		externalDataSourceTypePlatformPropEnum = append(externalDataSourceTypePlatformPropEnum, v)
	}
}

const (

	// ExternalDataSourcePlatformSALESFORCE captures enum value "SALESFORCE"
	ExternalDataSourcePlatformSALESFORCE string = "SALESFORCE"
)

// prop value enum
func (m *ExternalDataSource) validatePlatformEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, externalDataSourceTypePlatformPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ExternalDataSource) validatePlatform(formats strfmt.Registry) error {
	if swag.IsZero(m.Platform) { // not required
		return nil
	}

	// value enum
	if err := m.validatePlatformEnum("platform", "body", m.Platform); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this external data source based on context it is used
func (m *ExternalDataSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ExternalDataSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExternalDataSource) UnmarshalBinary(b []byte) error {
	var res ExternalDataSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
