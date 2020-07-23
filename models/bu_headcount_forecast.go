// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuHeadcountForecast bu headcount forecast
//
// swagger:model BuHeadcountForecast
type BuHeadcountForecast struct {

	// entities
	Entities []*BuPlanningGroupHeadcountForecast `json:"entities"`

	// Reference start date for the interval values in each forecast entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	ReferenceStartDate strfmt.DateTime `json:"referenceStartDate,omitempty"`
}

// Validate validates this bu headcount forecast
func (m *BuHeadcountForecast) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuHeadcountForecast) validateEntities(formats strfmt.Registry) error {

	if swag.IsZero(m.Entities) { // not required
		return nil
	}

	for i := 0; i < len(m.Entities); i++ {
		if swag.IsZero(m.Entities[i]) { // not required
			continue
		}

		if m.Entities[i] != nil {
			if err := m.Entities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("entities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuHeadcountForecast) validateReferenceStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("referenceStartDate", "body", "date-time", m.ReferenceStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuHeadcountForecast) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuHeadcountForecast) UnmarshalBinary(b []byte) error {
	var res BuHeadcountForecast
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
