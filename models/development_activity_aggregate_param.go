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

// DevelopmentActivityAggregateParam development activity aggregate param
//
// swagger:model DevelopmentActivityAggregateParam
type DevelopmentActivityAggregateParam struct {

	// The filter applied to the data. This is ANDed with the interval parameter.
	// Required: true
	Filter *DevelopmentActivityAggregateQueryRequestFilter `json:"filter"`

	// Specifies if the aggregated data is combined into a single set of metrics (groupBy is empty or not specified), or contains an element per attendeeId (groupBy is "attendeeId")
	GroupBy []string `json:"groupBy"`

	// Specifies the range of due dates to be used for filtering. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	// Required: true
	Interval *string `json:"interval"`

	// The list of metrics to be returned. If omitted, all metrics are returned.
	Metrics []string `json:"metrics"`
}

// Validate validates this development activity aggregate param
func (m *DevelopmentActivityAggregateParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetrics(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DevelopmentActivityAggregateParam) validateFilter(formats strfmt.Registry) error {

	if err := validate.Required("filter", "body", m.Filter); err != nil {
		return err
	}

	if m.Filter != nil {
		if err := m.Filter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("filter")
			}
			return err
		}
	}

	return nil
}

var developmentActivityAggregateParamGroupByItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attendeeId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		developmentActivityAggregateParamGroupByItemsEnum = append(developmentActivityAggregateParamGroupByItemsEnum, v)
	}
}

func (m *DevelopmentActivityAggregateParam) validateGroupByItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, developmentActivityAggregateParamGroupByItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DevelopmentActivityAggregateParam) validateGroupBy(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupBy) { // not required
		return nil
	}

	for i := 0; i < len(m.GroupBy); i++ {

		// value enum
		if err := m.validateGroupByItemsEnum("groupBy"+"."+strconv.Itoa(i), "body", m.GroupBy[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *DevelopmentActivityAggregateParam) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

var developmentActivityAggregateParamMetricsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nActivities","nPlannedActivities","nInProgressActivities","nCompleteActivities","nOverdueActivities","nInvalidScheduleActivities","nPassedActivities","nFailedActivities","oActivityScore","nNotCompletedActivities"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		developmentActivityAggregateParamMetricsItemsEnum = append(developmentActivityAggregateParamMetricsItemsEnum, v)
	}
}

func (m *DevelopmentActivityAggregateParam) validateMetricsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, developmentActivityAggregateParamMetricsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DevelopmentActivityAggregateParam) validateMetrics(formats strfmt.Registry) error {

	if swag.IsZero(m.Metrics) { // not required
		return nil
	}

	for i := 0; i < len(m.Metrics); i++ {

		// value enum
		if err := m.validateMetricsItemsEnum("metrics"+"."+strconv.Itoa(i), "body", m.Metrics[i]); err != nil {
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DevelopmentActivityAggregateParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DevelopmentActivityAggregateParam) UnmarshalBinary(b []byte) error {
	var res DevelopmentActivityAggregateParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
