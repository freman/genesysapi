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

// WfmHistoricalShrinkageRequest wfm historical shrinkage request
//
// swagger:model WfmHistoricalShrinkageRequest
type WfmHistoricalShrinkageRequest struct {

	// End of the date range to query in ISO-8601 format. If it is not set, end date will be set to current time
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// Shrinkage aggregation interval granularity.
	// Enum: [Daily Weekly]
	Granularity string `json:"granularity,omitempty"`

	// Beginning of the date range to query in ISO-8601 format
	// Required: true
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate"`

	// The time zone, in olson format, to use in defining days when computing shrinkage for requested granularity. If it is not set, the business unit time zone will be used. The results will be returned as UTC timestamps regardless of the time zone input.
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this wfm historical shrinkage request
func (m *WfmHistoricalShrinkageRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGranularity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmHistoricalShrinkageRequest) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var wfmHistoricalShrinkageRequestTypeGranularityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Daily","Weekly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wfmHistoricalShrinkageRequestTypeGranularityPropEnum = append(wfmHistoricalShrinkageRequestTypeGranularityPropEnum, v)
	}
}

const (

	// WfmHistoricalShrinkageRequestGranularityDaily captures enum value "Daily"
	WfmHistoricalShrinkageRequestGranularityDaily string = "Daily"

	// WfmHistoricalShrinkageRequestGranularityWeekly captures enum value "Weekly"
	WfmHistoricalShrinkageRequestGranularityWeekly string = "Weekly"
)

// prop value enum
func (m *WfmHistoricalShrinkageRequest) validateGranularityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wfmHistoricalShrinkageRequestTypeGranularityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WfmHistoricalShrinkageRequest) validateGranularity(formats strfmt.Registry) error {
	if swag.IsZero(m.Granularity) { // not required
		return nil
	}

	// value enum
	if err := m.validateGranularityEnum("granularity", "body", m.Granularity); err != nil {
		return err
	}

	return nil
}

func (m *WfmHistoricalShrinkageRequest) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", m.StartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this wfm historical shrinkage request based on context it is used
func (m *WfmHistoricalShrinkageRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WfmHistoricalShrinkageRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmHistoricalShrinkageRequest) UnmarshalBinary(b []byte) error {
	var res WfmHistoricalShrinkageRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
