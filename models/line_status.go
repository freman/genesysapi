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

// LineStatus line status
//
// swagger:model LineStatus
type LineStatus struct {

	// The line's address of record.
	AddressOfRecord string `json:"addressOfRecord,omitempty"`

	// The addresses used to contact the line.
	ContactAddresses []string `json:"contactAddresses"`

	// The id of this line
	ID string `json:"id,omitempty"`

	// Indicates whether the edge can reach the line.
	Reachable bool `json:"reachable"`

	// The time the line entered its current reachable state. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ReachableStateTime strfmt.DateTime `json:"reachableStateTime,omitempty"`
}

// Validate validates this line status
func (m *LineStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateReachableStateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineStatus) validateReachableStateTime(formats strfmt.Registry) error {

	if swag.IsZero(m.ReachableStateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("reachableStateTime", "body", "date-time", m.ReachableStateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LineStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LineStatus) UnmarshalBinary(b []byte) error {
	var res LineStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
