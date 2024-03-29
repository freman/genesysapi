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

// GenerateBuForecastRequest generate bu forecast request
//
// swagger:model GenerateBuForecastRequest
type GenerateBuForecastRequest struct {

	// Whether this forecast can be used for scheduling
	CanUseForScheduling bool `json:"canUseForScheduling"`

	// The description for the forecast
	// Required: true
	Description *string `json:"description"`

	// The number of weeks this forecast covers
	WeekCount int32 `json:"weekCount,omitempty"`
}

// Validate validates this generate bu forecast request
func (m *GenerateBuForecastRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GenerateBuForecastRequest) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this generate bu forecast request based on context it is used
func (m *GenerateBuForecastRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenerateBuForecastRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenerateBuForecastRequest) UnmarshalBinary(b []byte) error {
	var res GenerateBuForecastRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
