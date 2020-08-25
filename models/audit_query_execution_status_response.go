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

// AuditQueryExecutionStatusResponse audit query execution status response
//
// swagger:model AuditQueryExecutionStatusResponse
type AuditQueryExecutionStatusResponse struct {

	// Filters for the audit query.
	Filters []*AuditQueryFilter `json:"filters"`

	// Id of the audit query execution request.
	ID string `json:"id,omitempty"`

	// Interval for the audit query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	Interval string `json:"interval,omitempty"`

	// Service name for the audit query.
	// Enum: [ContactCenter ContentManagement PeoplePermissions Quality LanguageUnderstanding TopicsDefinitions PredictiveEngagement WorkforceManagement]
	ServiceName string `json:"serviceName,omitempty"`

	// Sort parameter for the audit query.
	Sort []*AuditQuerySort `json:"sort"`

	// Start date and time of the audit query execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// Status of the audit query execution request.
	// Enum: [Queued Running Succeeded Failed Cancelled]
	State string `json:"state,omitempty"`
}

// Validate validates this audit query execution status response
func (m *AuditQueryExecutionStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditQueryExecutionStatusResponse) validateFilters(formats strfmt.Registry) error {

	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var auditQueryExecutionStatusResponseTypeServiceNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ContactCenter","ContentManagement","PeoplePermissions","Quality","LanguageUnderstanding","TopicsDefinitions","PredictiveEngagement","WorkforceManagement"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditQueryExecutionStatusResponseTypeServiceNamePropEnum = append(auditQueryExecutionStatusResponseTypeServiceNamePropEnum, v)
	}
}

const (

	// AuditQueryExecutionStatusResponseServiceNameContactCenter captures enum value "ContactCenter"
	AuditQueryExecutionStatusResponseServiceNameContactCenter string = "ContactCenter"

	// AuditQueryExecutionStatusResponseServiceNameContentManagement captures enum value "ContentManagement"
	AuditQueryExecutionStatusResponseServiceNameContentManagement string = "ContentManagement"

	// AuditQueryExecutionStatusResponseServiceNamePeoplePermissions captures enum value "PeoplePermissions"
	AuditQueryExecutionStatusResponseServiceNamePeoplePermissions string = "PeoplePermissions"

	// AuditQueryExecutionStatusResponseServiceNameQuality captures enum value "Quality"
	AuditQueryExecutionStatusResponseServiceNameQuality string = "Quality"

	// AuditQueryExecutionStatusResponseServiceNameLanguageUnderstanding captures enum value "LanguageUnderstanding"
	AuditQueryExecutionStatusResponseServiceNameLanguageUnderstanding string = "LanguageUnderstanding"

	// AuditQueryExecutionStatusResponseServiceNameTopicsDefinitions captures enum value "TopicsDefinitions"
	AuditQueryExecutionStatusResponseServiceNameTopicsDefinitions string = "TopicsDefinitions"

	// AuditQueryExecutionStatusResponseServiceNamePredictiveEngagement captures enum value "PredictiveEngagement"
	AuditQueryExecutionStatusResponseServiceNamePredictiveEngagement string = "PredictiveEngagement"

	// AuditQueryExecutionStatusResponseServiceNameWorkforceManagement captures enum value "WorkforceManagement"
	AuditQueryExecutionStatusResponseServiceNameWorkforceManagement string = "WorkforceManagement"
)

// prop value enum
func (m *AuditQueryExecutionStatusResponse) validateServiceNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditQueryExecutionStatusResponseTypeServiceNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditQueryExecutionStatusResponse) validateServiceName(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceName) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceNameEnum("serviceName", "body", m.ServiceName); err != nil {
		return err
	}

	return nil
}

func (m *AuditQueryExecutionStatusResponse) validateSort(formats strfmt.Registry) error {

	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	for i := 0; i < len(m.Sort); i++ {
		if swag.IsZero(m.Sort[i]) { // not required
			continue
		}

		if m.Sort[i] != nil {
			if err := m.Sort[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditQueryExecutionStatusResponse) validateStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var auditQueryExecutionStatusResponseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Queued","Running","Succeeded","Failed","Cancelled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditQueryExecutionStatusResponseTypeStatePropEnum = append(auditQueryExecutionStatusResponseTypeStatePropEnum, v)
	}
}

const (

	// AuditQueryExecutionStatusResponseStateQueued captures enum value "Queued"
	AuditQueryExecutionStatusResponseStateQueued string = "Queued"

	// AuditQueryExecutionStatusResponseStateRunning captures enum value "Running"
	AuditQueryExecutionStatusResponseStateRunning string = "Running"

	// AuditQueryExecutionStatusResponseStateSucceeded captures enum value "Succeeded"
	AuditQueryExecutionStatusResponseStateSucceeded string = "Succeeded"

	// AuditQueryExecutionStatusResponseStateFailed captures enum value "Failed"
	AuditQueryExecutionStatusResponseStateFailed string = "Failed"

	// AuditQueryExecutionStatusResponseStateCancelled captures enum value "Cancelled"
	AuditQueryExecutionStatusResponseStateCancelled string = "Cancelled"
)

// prop value enum
func (m *AuditQueryExecutionStatusResponse) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditQueryExecutionStatusResponseTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditQueryExecutionStatusResponse) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditQueryExecutionStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditQueryExecutionStatusResponse) UnmarshalBinary(b []byte) error {
	var res AuditQueryExecutionStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}