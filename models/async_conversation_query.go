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

// AsyncConversationQuery async conversation query
//
// swagger:model AsyncConversationQuery
type AsyncConversationQuery struct {

	// Filters that target conversation-level data
	ConversationFilters []*ConversationDetailQueryFilter `json:"conversationFilters"`

	// Filters that target evaluations
	EvaluationFilters []*EvaluationDetailQueryFilter `json:"evaluationFilters"`

	// Specifies the date and time range of data being queried. Results will include all conversations that had activity during the interval. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	// Required: true
	Interval *string `json:"interval"`

	// Specify number of results to be returned
	Limit int32 `json:"limit,omitempty"`

	// Sort the result set in ascending/descending order. Default is ascending
	// Enum: [asc desc]
	Order string `json:"order,omitempty"`

	// Specify which data element within the result set to use for sorting. The options  to use as a basis for sorting the results: conversationStart, segmentStart, and segmentEnd. If not specified, the default is conversationStart
	// Enum: [conversationStart conversationEnd segmentStart segmentEnd]
	OrderBy string `json:"orderBy,omitempty"`

	// Filters that target resolutions
	ResolutionFilters []*ResolutionDetailQueryFilter `json:"resolutionFilters"`

	// Filters that target individual segments within a conversation
	SegmentFilters []*SegmentDetailQueryFilter `json:"segmentFilters"`

	// Add a filter to only include conversations that started after the beginning of the interval start date (UTC)
	StartOfDayIntervalMatching bool `json:"startOfDayIntervalMatching"`

	// Filters that target surveys
	SurveyFilters []*SurveyDetailQueryFilter `json:"surveyFilters"`
}

// Validate validates this async conversation query
func (m *AsyncConversationQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversationFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvaluationFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResolutionFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSegmentFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSurveyFilters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AsyncConversationQuery) validateConversationFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.ConversationFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.ConversationFilters); i++ {
		if swag.IsZero(m.ConversationFilters[i]) { // not required
			continue
		}

		if m.ConversationFilters[i] != nil {
			if err := m.ConversationFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("conversationFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AsyncConversationQuery) validateEvaluationFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.EvaluationFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.EvaluationFilters); i++ {
		if swag.IsZero(m.EvaluationFilters[i]) { // not required
			continue
		}

		if m.EvaluationFilters[i] != nil {
			if err := m.EvaluationFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evaluationFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AsyncConversationQuery) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

var asyncConversationQueryTypeOrderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		asyncConversationQueryTypeOrderPropEnum = append(asyncConversationQueryTypeOrderPropEnum, v)
	}
}

const (

	// AsyncConversationQueryOrderAsc captures enum value "asc"
	AsyncConversationQueryOrderAsc string = "asc"

	// AsyncConversationQueryOrderDesc captures enum value "desc"
	AsyncConversationQueryOrderDesc string = "desc"
)

// prop value enum
func (m *AsyncConversationQuery) validateOrderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, asyncConversationQueryTypeOrderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AsyncConversationQuery) validateOrder(formats strfmt.Registry) error {

	if swag.IsZero(m.Order) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderEnum("order", "body", m.Order); err != nil {
		return err
	}

	return nil
}

var asyncConversationQueryTypeOrderByPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["conversationStart","conversationEnd","segmentStart","segmentEnd"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		asyncConversationQueryTypeOrderByPropEnum = append(asyncConversationQueryTypeOrderByPropEnum, v)
	}
}

const (

	// AsyncConversationQueryOrderByConversationStart captures enum value "conversationStart"
	AsyncConversationQueryOrderByConversationStart string = "conversationStart"

	// AsyncConversationQueryOrderByConversationEnd captures enum value "conversationEnd"
	AsyncConversationQueryOrderByConversationEnd string = "conversationEnd"

	// AsyncConversationQueryOrderBySegmentStart captures enum value "segmentStart"
	AsyncConversationQueryOrderBySegmentStart string = "segmentStart"

	// AsyncConversationQueryOrderBySegmentEnd captures enum value "segmentEnd"
	AsyncConversationQueryOrderBySegmentEnd string = "segmentEnd"
)

// prop value enum
func (m *AsyncConversationQuery) validateOrderByEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, asyncConversationQueryTypeOrderByPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AsyncConversationQuery) validateOrderBy(formats strfmt.Registry) error {

	if swag.IsZero(m.OrderBy) { // not required
		return nil
	}

	// value enum
	if err := m.validateOrderByEnum("orderBy", "body", m.OrderBy); err != nil {
		return err
	}

	return nil
}

func (m *AsyncConversationQuery) validateResolutionFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.ResolutionFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.ResolutionFilters); i++ {
		if swag.IsZero(m.ResolutionFilters[i]) { // not required
			continue
		}

		if m.ResolutionFilters[i] != nil {
			if err := m.ResolutionFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("resolutionFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AsyncConversationQuery) validateSegmentFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.SegmentFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.SegmentFilters); i++ {
		if swag.IsZero(m.SegmentFilters[i]) { // not required
			continue
		}

		if m.SegmentFilters[i] != nil {
			if err := m.SegmentFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("segmentFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AsyncConversationQuery) validateSurveyFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.SurveyFilters) { // not required
		return nil
	}

	for i := 0; i < len(m.SurveyFilters); i++ {
		if swag.IsZero(m.SurveyFilters[i]) { // not required
			continue
		}

		if m.SurveyFilters[i] != nil {
			if err := m.SurveyFilters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("surveyFilters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AsyncConversationQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AsyncConversationQuery) UnmarshalBinary(b []byte) error {
	var res AsyncConversationQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
