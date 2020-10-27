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

// Wrapup wrapup
//
// swagger:model Wrapup
type Wrapup struct {

	// The user configured wrap up code id.
	Code string `json:"code,omitempty"`

	// The length of time in seconds that the agent spent doing after call work.
	DurationSeconds int32 `json:"durationSeconds,omitempty"`

	// The timestamp when the wrapup was finished. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// The user configured wrap up code name.
	Name string `json:"name,omitempty"`

	// Text entered by the agent to describe the call or disposition.
	Notes string `json:"notes,omitempty"`

	// Indicates if this is a pending save and should not require a code to be specified.  This allows someone to save some temporary wrapup that will be used later.
	Provisional bool `json:"provisional"`

	// List of tags selected by the agent to describe the call or disposition.
	Tags []string `json:"tags"`
}

// Validate validates this wrapup
func (m *Wrapup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Wrapup) validateEndTime(formats strfmt.Registry) error {

	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Wrapup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Wrapup) UnmarshalBinary(b []byte) error {
	var res Wrapup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
