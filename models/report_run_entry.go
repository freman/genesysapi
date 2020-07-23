// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportRunEntry report run entry
//
// swagger:model ReportRunEntry
type ReportRunEntry struct {

	// error message
	ErrorMessage string `json:"errorMessage,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// report format
	ReportFormat string `json:"reportFormat,omitempty"`

	// report Id
	ReportID string `json:"reportId,omitempty"`

	// report Url
	ReportURL string `json:"reportUrl,omitempty"`

	// run duration msec
	RunDurationMsec int64 `json:"runDurationMsec,omitempty"`

	// run status
	// Enum: [RUNNING COMPLETED COMPLETED_WITH_ERRORS FAILED FAILED_TIMEOUT FAILED_DATALIMIT UNABLE_TO_COMPLETE]
	RunStatus string `json:"runStatus,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	RunTime strfmt.DateTime `json:"runTime,omitempty"`

	// schedule Uri
	// Format: uri
	ScheduleURI strfmt.URI `json:"scheduleUri,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this report run entry
func (m *ReportRunEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRunStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduleURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportRunEntryTypeRunStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","COMPLETED","COMPLETED_WITH_ERRORS","FAILED","FAILED_TIMEOUT","FAILED_DATALIMIT","UNABLE_TO_COMPLETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportRunEntryTypeRunStatusPropEnum = append(reportRunEntryTypeRunStatusPropEnum, v)
	}
}

const (

	// ReportRunEntryRunStatusRUNNING captures enum value "RUNNING"
	ReportRunEntryRunStatusRUNNING string = "RUNNING"

	// ReportRunEntryRunStatusCOMPLETED captures enum value "COMPLETED"
	ReportRunEntryRunStatusCOMPLETED string = "COMPLETED"

	// ReportRunEntryRunStatusCOMPLETEDWITHERRORS captures enum value "COMPLETED_WITH_ERRORS"
	ReportRunEntryRunStatusCOMPLETEDWITHERRORS string = "COMPLETED_WITH_ERRORS"

	// ReportRunEntryRunStatusFAILED captures enum value "FAILED"
	ReportRunEntryRunStatusFAILED string = "FAILED"

	// ReportRunEntryRunStatusFAILEDTIMEOUT captures enum value "FAILED_TIMEOUT"
	ReportRunEntryRunStatusFAILEDTIMEOUT string = "FAILED_TIMEOUT"

	// ReportRunEntryRunStatusFAILEDDATALIMIT captures enum value "FAILED_DATALIMIT"
	ReportRunEntryRunStatusFAILEDDATALIMIT string = "FAILED_DATALIMIT"

	// ReportRunEntryRunStatusUNABLETOCOMPLETE captures enum value "UNABLE_TO_COMPLETE"
	ReportRunEntryRunStatusUNABLETOCOMPLETE string = "UNABLE_TO_COMPLETE"
)

// prop value enum
func (m *ReportRunEntry) validateRunStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportRunEntryTypeRunStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportRunEntry) validateRunStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RunStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateRunStatusEnum("runStatus", "body", m.RunStatus); err != nil {
		return err
	}

	return nil
}

func (m *ReportRunEntry) validateRunTime(formats strfmt.Registry) error {

	if swag.IsZero(m.RunTime) { // not required
		return nil
	}

	if err := validate.FormatOf("runTime", "body", "date-time", m.RunTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRunEntry) validateScheduleURI(formats strfmt.Registry) error {

	if swag.IsZero(m.ScheduleURI) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduleUri", "body", "uri", m.ScheduleURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportRunEntry) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportRunEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportRunEntry) UnmarshalBinary(b []byte) error {
	var res ReportRunEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
