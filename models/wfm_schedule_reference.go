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

// WfmScheduleReference wfm schedule reference
//
// swagger:model WfmScheduleReference
type WfmScheduleReference struct {

	// A reference to a Workforce Management Business Unit
	// Required: true
	BusinessUnit *WfmBusinessUnitReference `json:"businessUnit"`

	// The ID of the WFM schedule
	// Required: true
	ID *string `json:"id"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Required: true
	// Format: date
	WeekDate *strfmt.Date `json:"weekDate"`
}

// Validate validates this wfm schedule reference
func (m *WfmScheduleReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmScheduleReference) validateBusinessUnit(formats strfmt.Registry) error {

	if err := validate.Required("businessUnit", "body", m.BusinessUnit); err != nil {
		return err
	}

	if m.BusinessUnit != nil {
		if err := m.BusinessUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessUnit")
			}
			return err
		}
	}

	return nil
}

func (m *WfmScheduleReference) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *WfmScheduleReference) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WfmScheduleReference) validateWeekDate(formats strfmt.Registry) error {

	if err := validate.Required("weekDate", "body", m.WeekDate); err != nil {
		return err
	}

	if err := validate.FormatOf("weekDate", "body", "date", m.WeekDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WfmScheduleReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmScheduleReference) UnmarshalBinary(b []byte) error {
	var res WfmScheduleReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
