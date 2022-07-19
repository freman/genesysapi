// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ValueWrapperDate value wrapper date
//
// swagger:model ValueWrapperDate
type ValueWrapperDate struct {

	// The value for the associated field. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	Value strfmt.DateTime `json:"value,omitempty"`
}

// Validate validates this value wrapper date
func (m *ValueWrapperDate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ValueWrapperDate) validateValue(formats strfmt.Registry) error {

	if swag.IsZero(m.Value) { // not required
		return nil
	}

	if err := validate.FormatOf("value", "body", "date-time", m.Value.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ValueWrapperDate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ValueWrapperDate) UnmarshalBinary(b []byte) error {
	var res ValueWrapperDate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
