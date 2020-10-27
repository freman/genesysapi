// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuIntradayResponse bu intraday response
//
// swagger:model BuIntradayResponse
type BuIntradayResponse struct {

	// The categories to which this data corresponds
	Categories []string `json:"categories"`

	// The end of the date range for which this data applies. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// The aggregation period in minutes, which determines the interval duration of the returned data
	IntervalLengthMinutes int32 `json:"intervalLengthMinutes,omitempty"`

	// Intraday data grouped by a single media type and set of planning group IDs
	IntradayDataGroupings []*BuIntradayDataGroup `json:"intradayDataGroupings"`

	// If not null, the reason there was no data for the request
	// Enum: [NoPublishedSchedule NoSourceForecast]
	NoDataReason string `json:"noDataReason,omitempty"`

	// Schedule reference
	Schedule *BuScheduleReference `json:"schedule,omitempty"`

	// Short term forecast reference
	ShortTermForecast *BuShortTermForecastReference `json:"shortTermForecast,omitempty"`

	// The start of the date range for which this data applies.  This is also the start reference point for the intervals represented in the various arrays. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`
}

// Validate validates this bu intraday response
func (m *BuIntradayResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategories(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntradayDataGroupings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNoDataReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortTermForecast(formats); err != nil {
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

var buIntradayResponseCategoriesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ForecastData","ScheduleData","PerformancePredictionData"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buIntradayResponseCategoriesItemsEnum = append(buIntradayResponseCategoriesItemsEnum, v)
	}
}

func (m *BuIntradayResponse) validateCategoriesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buIntradayResponseCategoriesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuIntradayResponse) validateCategories(formats strfmt.Registry) error {

	if swag.IsZero(m.Categories) { // not required
		return nil
	}

	for i := 0; i < len(m.Categories); i++ {

		// value enum
		if err := m.validateCategoriesItemsEnum("categories"+"."+strconv.Itoa(i), "body", m.Categories[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *BuIntradayResponse) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuIntradayResponse) validateIntradayDataGroupings(formats strfmt.Registry) error {

	if swag.IsZero(m.IntradayDataGroupings) { // not required
		return nil
	}

	for i := 0; i < len(m.IntradayDataGroupings); i++ {
		if swag.IsZero(m.IntradayDataGroupings[i]) { // not required
			continue
		}

		if m.IntradayDataGroupings[i] != nil {
			if err := m.IntradayDataGroupings[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intradayDataGroupings" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var buIntradayResponseTypeNoDataReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NoPublishedSchedule","NoSourceForecast"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buIntradayResponseTypeNoDataReasonPropEnum = append(buIntradayResponseTypeNoDataReasonPropEnum, v)
	}
}

const (

	// BuIntradayResponseNoDataReasonNoPublishedSchedule captures enum value "NoPublishedSchedule"
	BuIntradayResponseNoDataReasonNoPublishedSchedule string = "NoPublishedSchedule"

	// BuIntradayResponseNoDataReasonNoSourceForecast captures enum value "NoSourceForecast"
	BuIntradayResponseNoDataReasonNoSourceForecast string = "NoSourceForecast"
)

// prop value enum
func (m *BuIntradayResponse) validateNoDataReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buIntradayResponseTypeNoDataReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuIntradayResponse) validateNoDataReason(formats strfmt.Registry) error {

	if swag.IsZero(m.NoDataReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateNoDataReasonEnum("noDataReason", "body", m.NoDataReason); err != nil {
		return err
	}

	return nil
}

func (m *BuIntradayResponse) validateSchedule(formats strfmt.Registry) error {

	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

func (m *BuIntradayResponse) validateShortTermForecast(formats strfmt.Registry) error {

	if swag.IsZero(m.ShortTermForecast) { // not required
		return nil
	}

	if m.ShortTermForecast != nil {
		if err := m.ShortTermForecast.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortTermForecast")
			}
			return err
		}
	}

	return nil
}

func (m *BuIntradayResponse) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuIntradayResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuIntradayResponse) UnmarshalBinary(b []byte) error {
	var res BuIntradayResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
