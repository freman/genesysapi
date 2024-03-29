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

// DataActionContactColumnFieldMapping data action contact column field mapping
//
// swagger:model DataActionContactColumnFieldMapping
type DataActionContactColumnFieldMapping struct {

	// The name of a contact column whose data will be passed to the data action
	// Required: true
	ContactColumnName *string `json:"contactColumnName"`

	// The name of an input field from the data action that the contact column data will be passed to
	// Required: true
	DataActionField *string `json:"dataActionField"`
}

// Validate validates this data action contact column field mapping
func (m *DataActionContactColumnFieldMapping) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContactColumnName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDataActionField(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataActionContactColumnFieldMapping) validateContactColumnName(formats strfmt.Registry) error {

	if err := validate.Required("contactColumnName", "body", m.ContactColumnName); err != nil {
		return err
	}

	return nil
}

func (m *DataActionContactColumnFieldMapping) validateDataActionField(formats strfmt.Registry) error {

	if err := validate.Required("dataActionField", "body", m.DataActionField); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data action contact column field mapping based on context it is used
func (m *DataActionContactColumnFieldMapping) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataActionContactColumnFieldMapping) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataActionContactColumnFieldMapping) UnmarshalBinary(b []byte) error {
	var res DataActionContactColumnFieldMapping
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
