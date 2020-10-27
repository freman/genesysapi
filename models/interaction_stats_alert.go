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

// InteractionStatsAlert interaction stats alert
//
// swagger:model InteractionStatsAlert
type InteractionStatsAlert struct {

	// A collection of notification methods.
	// Required: true
	// Read Only: true
	// Unique: true
	AlertTypes []string `json:"alertTypes"`

	// The dimension of concern.
	// Required: true
	// Read Only: true
	// Enum: [queueId userId]
	Dimension string `json:"dimension"`

	// The value of the dimension.
	// Required: true
	// Read Only: true
	DimensionValue string `json:"dimensionValue"`

	// The date/time the owning rule exiting in alarm status. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	EndDate strfmt.DateTime `json:"endDate,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The media type.
	// Required: true
	// Read Only: true
	// Enum: [voice chat email callback message]
	MediaType string `json:"mediaType"`

	// The metric to be assessed.
	// Required: true
	// Read Only: true
	// Enum: [tAbandon tAnswered tTalk nOffered tHandle nTransferred oServiceLevel tWait tHeld tAcw]
	Metric string `json:"metric"`

	// Name of the rule that generated the alert
	// Required: true
	// Read Only: true
	Name string `json:"name"`

	// The ids of users who were notified of alarm state change.
	// Required: true
	// Read Only: true
	// Unique: true
	NotificationUsers []*User `json:"notificationUsers"`

	// The comparison descriptor used against the metric's value.
	// Required: true
	// Read Only: true
	// Enum: [gt gte lt lte eq ne]
	NumericRange string `json:"numericRange"`

	// The id of the rule.
	// Required: true
	// Read Only: true
	RuleID string `json:"ruleId"`

	// rule Uri
	// Format: uri
	RuleURI strfmt.URI `json:"ruleUri,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The date/time the alert was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Read Only: true
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate"`

	// The statistic of concern for the metric.
	// Required: true
	// Read Only: true
	// Enum: [count min ratio max]
	Statistic string `json:"statistic"`

	// Indicates if the alert has been read.
	// Required: true
	Unread *bool `json:"unread"`

	// The threshold value.
	// Required: true
	// Read Only: true
	Value float64 `json:"value"`
}

// Validate validates this interaction stats alert
func (m *InteractionStatsAlert) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimension(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDimensionValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetric(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumericRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuleURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatistic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnread(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var interactionStatsAlertAlertTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SMS","DEVICE","EMAIL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionStatsAlertAlertTypesItemsEnum = append(interactionStatsAlertAlertTypesItemsEnum, v)
	}
}

func (m *InteractionStatsAlert) validateAlertTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interactionStatsAlertAlertTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InteractionStatsAlert) validateAlertTypes(formats strfmt.Registry) error {

	if err := validate.Required("alertTypes", "body", m.AlertTypes); err != nil {
		return err
	}

	if err := validate.UniqueItems("alertTypes", "body", m.AlertTypes); err != nil {
		return err
	}

	for i := 0; i < len(m.AlertTypes); i++ {

		// value enum
		if err := m.validateAlertTypesItemsEnum("alertTypes"+"."+strconv.Itoa(i), "body", m.AlertTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var interactionStatsAlertTypeDimensionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["queueId","userId"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionStatsAlertTypeDimensionPropEnum = append(interactionStatsAlertTypeDimensionPropEnum, v)
	}
}

const (

	// InteractionStatsAlertDimensionQueueID captures enum value "queueId"
	InteractionStatsAlertDimensionQueueID string = "queueId"

	// InteractionStatsAlertDimensionUserID captures enum value "userId"
	InteractionStatsAlertDimensionUserID string = "userId"
)

// prop value enum
func (m *InteractionStatsAlert) validateDimensionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interactionStatsAlertTypeDimensionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InteractionStatsAlert) validateDimension(formats strfmt.Registry) error {

	if err := validate.RequiredString("dimension", "body", string(m.Dimension)); err != nil {
		return err
	}

	// value enum
	if err := m.validateDimensionEnum("dimension", "body", m.Dimension); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateDimensionValue(formats strfmt.Registry) error {

	if err := validate.RequiredString("dimensionValue", "body", string(m.DimensionValue)); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var interactionStatsAlertTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["voice","chat","email","callback","message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionStatsAlertTypeMediaTypePropEnum = append(interactionStatsAlertTypeMediaTypePropEnum, v)
	}
}

const (

	// InteractionStatsAlertMediaTypeVoice captures enum value "voice"
	InteractionStatsAlertMediaTypeVoice string = "voice"

	// InteractionStatsAlertMediaTypeChat captures enum value "chat"
	InteractionStatsAlertMediaTypeChat string = "chat"

	// InteractionStatsAlertMediaTypeEmail captures enum value "email"
	InteractionStatsAlertMediaTypeEmail string = "email"

	// InteractionStatsAlertMediaTypeCallback captures enum value "callback"
	InteractionStatsAlertMediaTypeCallback string = "callback"

	// InteractionStatsAlertMediaTypeMessage captures enum value "message"
	InteractionStatsAlertMediaTypeMessage string = "message"
)

// prop value enum
func (m *InteractionStatsAlert) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interactionStatsAlertTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InteractionStatsAlert) validateMediaType(formats strfmt.Registry) error {

	if err := validate.RequiredString("mediaType", "body", string(m.MediaType)); err != nil {
		return err
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

var interactionStatsAlertTypeMetricPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tAbandon","tAnswered","tTalk","nOffered","tHandle","nTransferred","oServiceLevel","tWait","tHeld","tAcw"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionStatsAlertTypeMetricPropEnum = append(interactionStatsAlertTypeMetricPropEnum, v)
	}
}

const (

	// InteractionStatsAlertMetricTAbandon captures enum value "tAbandon"
	InteractionStatsAlertMetricTAbandon string = "tAbandon"

	// InteractionStatsAlertMetricTAnswered captures enum value "tAnswered"
	InteractionStatsAlertMetricTAnswered string = "tAnswered"

	// InteractionStatsAlertMetricTTalk captures enum value "tTalk"
	InteractionStatsAlertMetricTTalk string = "tTalk"

	// InteractionStatsAlertMetricNOffered captures enum value "nOffered"
	InteractionStatsAlertMetricNOffered string = "nOffered"

	// InteractionStatsAlertMetricTHandle captures enum value "tHandle"
	InteractionStatsAlertMetricTHandle string = "tHandle"

	// InteractionStatsAlertMetricNTransferred captures enum value "nTransferred"
	InteractionStatsAlertMetricNTransferred string = "nTransferred"

	// InteractionStatsAlertMetricOServiceLevel captures enum value "oServiceLevel"
	InteractionStatsAlertMetricOServiceLevel string = "oServiceLevel"

	// InteractionStatsAlertMetricTWait captures enum value "tWait"
	InteractionStatsAlertMetricTWait string = "tWait"

	// InteractionStatsAlertMetricTHeld captures enum value "tHeld"
	InteractionStatsAlertMetricTHeld string = "tHeld"

	// InteractionStatsAlertMetricTAcw captures enum value "tAcw"
	InteractionStatsAlertMetricTAcw string = "tAcw"
)

// prop value enum
func (m *InteractionStatsAlert) validateMetricEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interactionStatsAlertTypeMetricPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InteractionStatsAlert) validateMetric(formats strfmt.Registry) error {

	if err := validate.RequiredString("metric", "body", string(m.Metric)); err != nil {
		return err
	}

	// value enum
	if err := m.validateMetricEnum("metric", "body", m.Metric); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateNotificationUsers(formats strfmt.Registry) error {

	if err := validate.Required("notificationUsers", "body", m.NotificationUsers); err != nil {
		return err
	}

	if err := validate.UniqueItems("notificationUsers", "body", m.NotificationUsers); err != nil {
		return err
	}

	for i := 0; i < len(m.NotificationUsers); i++ {
		if swag.IsZero(m.NotificationUsers[i]) { // not required
			continue
		}

		if m.NotificationUsers[i] != nil {
			if err := m.NotificationUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("notificationUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var interactionStatsAlertTypeNumericRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["gt","gte","lt","lte","eq","ne"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionStatsAlertTypeNumericRangePropEnum = append(interactionStatsAlertTypeNumericRangePropEnum, v)
	}
}

const (

	// InteractionStatsAlertNumericRangeGt captures enum value "gt"
	InteractionStatsAlertNumericRangeGt string = "gt"

	// InteractionStatsAlertNumericRangeGte captures enum value "gte"
	InteractionStatsAlertNumericRangeGte string = "gte"

	// InteractionStatsAlertNumericRangeLt captures enum value "lt"
	InteractionStatsAlertNumericRangeLt string = "lt"

	// InteractionStatsAlertNumericRangeLte captures enum value "lte"
	InteractionStatsAlertNumericRangeLte string = "lte"

	// InteractionStatsAlertNumericRangeEq captures enum value "eq"
	InteractionStatsAlertNumericRangeEq string = "eq"

	// InteractionStatsAlertNumericRangeNe captures enum value "ne"
	InteractionStatsAlertNumericRangeNe string = "ne"
)

// prop value enum
func (m *InteractionStatsAlert) validateNumericRangeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interactionStatsAlertTypeNumericRangePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InteractionStatsAlert) validateNumericRange(formats strfmt.Registry) error {

	if err := validate.RequiredString("numericRange", "body", string(m.NumericRange)); err != nil {
		return err
	}

	// value enum
	if err := m.validateNumericRangeEnum("numericRange", "body", m.NumericRange); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateRuleID(formats strfmt.Registry) error {

	if err := validate.RequiredString("ruleId", "body", string(m.RuleID)); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateRuleURI(formats strfmt.Registry) error {

	if swag.IsZero(m.RuleURI) { // not required
		return nil
	}

	if err := validate.FormatOf("ruleUri", "body", "uri", m.RuleURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateStartDate(formats strfmt.Registry) error {

	if err := validate.Required("startDate", "body", strfmt.DateTime(m.StartDate)); err != nil {
		return err
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var interactionStatsAlertTypeStatisticPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["count","min","ratio","max"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		interactionStatsAlertTypeStatisticPropEnum = append(interactionStatsAlertTypeStatisticPropEnum, v)
	}
}

const (

	// InteractionStatsAlertStatisticCount captures enum value "count"
	InteractionStatsAlertStatisticCount string = "count"

	// InteractionStatsAlertStatisticMin captures enum value "min"
	InteractionStatsAlertStatisticMin string = "min"

	// InteractionStatsAlertStatisticRatio captures enum value "ratio"
	InteractionStatsAlertStatisticRatio string = "ratio"

	// InteractionStatsAlertStatisticMax captures enum value "max"
	InteractionStatsAlertStatisticMax string = "max"
)

// prop value enum
func (m *InteractionStatsAlert) validateStatisticEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, interactionStatsAlertTypeStatisticPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InteractionStatsAlert) validateStatistic(formats strfmt.Registry) error {

	if err := validate.RequiredString("statistic", "body", string(m.Statistic)); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatisticEnum("statistic", "body", m.Statistic); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateUnread(formats strfmt.Registry) error {

	if err := validate.Required("unread", "body", m.Unread); err != nil {
		return err
	}

	return nil
}

func (m *InteractionStatsAlert) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", float64(m.Value)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InteractionStatsAlert) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InteractionStatsAlert) UnmarshalBinary(b []byte) error {
	var res InteractionStatsAlert
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
