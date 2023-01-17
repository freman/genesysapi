// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateTimeOffLimitRequest create time off limit request
//
// swagger:model CreateTimeOffLimitRequest
type CreateTimeOffLimitRequest struct {

	// The default limit value in minutes per granularity. If not specified, then 0 is assumed, which means there are no time off minutes available
	DefaultLimitMinutes int32 `json:"defaultLimitMinutes,omitempty"`

	// Granularity choice for time off limit. If not specified, 'Daily' is assumed
	// Enum: [Daily]
	Granularity string `json:"granularity,omitempty"`
}

// Validate validates this create time off limit request
func (m *CreateTimeOffLimitRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGranularity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createTimeOffLimitRequestTypeGranularityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Daily"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createTimeOffLimitRequestTypeGranularityPropEnum = append(createTimeOffLimitRequestTypeGranularityPropEnum, v)
	}
}

const (

	// CreateTimeOffLimitRequestGranularityDaily captures enum value "Daily"
	CreateTimeOffLimitRequestGranularityDaily string = "Daily"
)

// prop value enum
func (m *CreateTimeOffLimitRequest) validateGranularityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createTimeOffLimitRequestTypeGranularityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateTimeOffLimitRequest) validateGranularity(formats strfmt.Registry) error {
	if swag.IsZero(m.Granularity) { // not required
		return nil
	}

	// value enum
	if err := m.validateGranularityEnum("granularity", "body", m.Granularity); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create time off limit request based on context it is used
func (m *CreateTimeOffLimitRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreateTimeOffLimitRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateTimeOffLimitRequest) UnmarshalBinary(b []byte) error {
	var res CreateTimeOffLimitRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
