// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BuForecastResultResponse bu forecast result response
//
// swagger:model BuForecastResultResponse
type BuForecastResultResponse struct {

	// The download url to fetch the result.  Only populated if the result is too large to pass through the api directly
	DownloadURL string `json:"downloadUrl,omitempty"`

	// The result of the operation.  Populated whenever the result is small enough to pass through the api directly
	Result *BuForecastResult `json:"result,omitempty"`
}

// Validate validates this bu forecast result response
func (m *BuForecastResultResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResult(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuForecastResultResponse) validateResult(formats strfmt.Registry) error {

	if swag.IsZero(m.Result) { // not required
		return nil
	}

	if m.Result != nil {
		if err := m.Result.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("result")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuForecastResultResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuForecastResultResponse) UnmarshalBinary(b []byte) error {
	var res BuForecastResultResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}