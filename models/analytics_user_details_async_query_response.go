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

// AnalyticsUserDetailsAsyncQueryResponse analytics user details async query response
//
// swagger:model AnalyticsUserDetailsAsyncQueryResponse
type AnalyticsUserDetailsAsyncQueryResponse struct {

	// Optional cursor to indicate where to resume the results
	Cursor string `json:"cursor,omitempty"`

	// Data available up to at least this datetime. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DataAvailabilityDate strfmt.DateTime `json:"dataAvailabilityDate,omitempty"`

	// user details
	UserDetails []*AnalyticsUserDetail `json:"userDetails"`
}

// Validate validates this analytics user details async query response
func (m *AnalyticsUserDetailsAsyncQueryResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataAvailabilityDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AnalyticsUserDetailsAsyncQueryResponse) validateDataAvailabilityDate(formats strfmt.Registry) error {

	if swag.IsZero(m.DataAvailabilityDate) { // not required
		return nil
	}

	if err := validate.FormatOf("dataAvailabilityDate", "body", "date-time", m.DataAvailabilityDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AnalyticsUserDetailsAsyncQueryResponse) validateUserDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.UserDetails) { // not required
		return nil
	}

	for i := 0; i < len(m.UserDetails); i++ {
		if swag.IsZero(m.UserDetails[i]) { // not required
			continue
		}

		if m.UserDetails[i] != nil {
			if err := m.UserDetails[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userDetails" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsUserDetailsAsyncQueryResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsUserDetailsAsyncQueryResponse) UnmarshalBinary(b []byte) error {
	var res AnalyticsUserDetailsAsyncQueryResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
