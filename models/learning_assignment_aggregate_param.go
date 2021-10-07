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

// LearningAssignmentAggregateParam learning assignment aggregate param
//
// swagger:model LearningAssignmentAggregateParam
type LearningAssignmentAggregateParam struct {

	// The filter applied to the data.  This is ANDed with the interval parameter.
	// Required: true
	Filter *LearningAssignmentAggregateQueryRequestFilter `json:"filter"`

	// Specifies if the aggregated data is combined into a single set of metrics (groupBy is empty or not specified), or contains an element per attendeeId (groupBy is "attendeeId")
	GroupBy []string `json:"groupBy"`

	// Specifies the range of due dates to be used for filtering. Milliseconds will be truncated. A maximum of 1 year can be specified in the range. End date is not inclusive. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	// Required: true
	Interval *string `json:"interval"`

	// The list of metrics to be returned. If omitted, all metrics are returned.
	Metrics []string `json:"metrics"`
}

// Validate validates this learning assignment aggregate param
func (m *LearningAssignmentAggregateParam) Validate(formats strfmt.Registry) error {
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

func (m *LearningAssignmentAggregateParam) validateFilter(formats strfmt.Registry) error {

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

var learningAssignmentAggregateParamGroupByItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["attendeeId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningAssignmentAggregateParamGroupByItemsEnum = append(learningAssignmentAggregateParamGroupByItemsEnum, v)
	}
}

func (m *LearningAssignmentAggregateParam) validateGroupByItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningAssignmentAggregateParamGroupByItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningAssignmentAggregateParam) validateGroupBy(formats strfmt.Registry) error {

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

func (m *LearningAssignmentAggregateParam) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

var learningAssignmentAggregateParamMetricsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["nActivities","nPlannedActivities","nInProgressActivities","nCompleteActivities","nOverdueActivities","nPassedActivities","nFailedActivities","oActivityScore"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		learningAssignmentAggregateParamMetricsItemsEnum = append(learningAssignmentAggregateParamMetricsItemsEnum, v)
	}
}

func (m *LearningAssignmentAggregateParam) validateMetricsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, learningAssignmentAggregateParamMetricsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LearningAssignmentAggregateParam) validateMetrics(formats strfmt.Registry) error {

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
func (m *LearningAssignmentAggregateParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LearningAssignmentAggregateParam) UnmarshalBinary(b []byte) error {
	var res LearningAssignmentAggregateParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
