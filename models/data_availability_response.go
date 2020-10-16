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

// DataAvailabilityResponse data availability response
//
// swagger:model DataAvailabilityResponse
type DataAvailabilityResponse struct {

	// Date and time before which data is guaranteed to be available in the datalake. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DataAvailabilityDate strfmt.DateTime `json:"dataAvailabilityDate,omitempty"`
}

// Validate validates this data availability response
func (m *DataAvailabilityResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataAvailabilityDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataAvailabilityResponse) validateDataAvailabilityDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DataAvailabilityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dataAvailabilityDate", "body", "date-time", m.DataAvailabilityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataAvailabilityResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataAvailabilityResponse) UnmarshalBinary(b []byte) error {
	var res DataAvailabilityResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}