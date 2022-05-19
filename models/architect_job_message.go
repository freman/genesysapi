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

// ArchitectJobMessage architect job message
//
// swagger:model ArchitectJobMessage
type ArchitectJobMessage struct {

	// The DateTime when the message was generated. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateTime strfmt.DateTime `json:"dateTime,omitempty"`

	// The text of the message.
	Text string `json:"text,omitempty"`

	// The message type.
	// Enum: [Error Warning]
	Type string `json:"type,omitempty"`
}

// Validate validates this architect job message
func (m *ArchitectJobMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateTime(formats); err != nil {
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

func (m *ArchitectJobMessage) validateDateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.DateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTime", "body", "date-time", m.DateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var architectJobMessageTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Error","Warning"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		architectJobMessageTypeTypePropEnum = append(architectJobMessageTypeTypePropEnum, v)
	}
}

const (

	// ArchitectJobMessageTypeError captures enum value "Error"
	ArchitectJobMessageTypeError string = "Error"

	// ArchitectJobMessageTypeWarning captures enum value "Warning"
	ArchitectJobMessageTypeWarning string = "Warning"
)

// prop value enum
func (m *ArchitectJobMessage) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, architectJobMessageTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchitectJobMessage) validateType(formats strfmt.Registry) error {

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
func (m *ArchitectJobMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchitectJobMessage) UnmarshalBinary(b []byte) error {
	var res ArchitectJobMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
