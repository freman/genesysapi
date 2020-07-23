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

// ReportingExportJobRequest reporting export job request
//
// swagger:model ReportingExportJobRequest
type ReportingExportJobRequest struct {

	// The requested format of the exported data
	// Required: true
	// Enum: [CSV PDF]
	ExportFormat *string `json:"exportFormat"`

	// Filters to apply to create the view
	// Required: true
	Filter *ViewFilter `json:"filter"`

	// Indicates if custom participant attributes will be exported
	HasCustomParticipantAttributes bool `json:"hasCustomParticipantAttributes,omitempty"`

	// Indicates if durations are formatted in hh:mm:ss format instead of ms
	HasFormatDurations bool `json:"hasFormatDurations,omitempty"`

	// Indicates if filters will be split in aggregate detail exports
	HasSplitFilters bool `json:"hasSplitFilters,omitempty"`

	// The time period used to limit the the exported data. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	// Required: true
	Interval *string `json:"interval"`

	// The locale use for localization of the exported data, i.e. en-us, es-mx
	// Required: true
	Locale *string `json:"locale"`

	// The user supplied name of the export request
	// Required: true
	Name *string `json:"name"`

	// The Period of the request in which to break down the intervals. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	// Required: true
	Period *string `json:"period"`

	// Indicates if the request has been marked as read
	Read bool `json:"read,omitempty"`

	// The list of email recipients for the exports
	RecipientEmails []string `json:"recipientEmails"`

	// The list of ordered selected columns from the export view by the user
	SelectedColumns []*SelectedColumns `json:"selectedColumns"`

	// The requested timezone of the exported data. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	// Required: true
	TimeZone *string `json:"timeZone"`

	// The type of view export job to be created
	// Required: true
	// Enum: [QUEUE_PERFORMANCE_SUMMARY_VIEW QUEUE_PERFORMANCE_DETAIL_VIEW INTERACTION_SEARCH_VIEW AGENT_PERFORMANCE_SUMMARY_VIEW AGENT_PERFORMANCE_DETAIL_VIEW AGENT_STATUS_SUMMARY_VIEW AGENT_STATUS_DETAIL_VIEW AGENT_EVALUATION_SUMMARY_VIEW AGENT_EVALUATION_DETAIL_VIEW AGENT_QUEUE_DETAIL_VIEW AGENT_INTERACTION_DETAIL_VIEW ABANDON_INSIGHTS_VIEW SKILLS_PERFORMANCE_VIEW SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW SURVEY_FORM_PERFORMANCE_DETAIL_VIEW DNIS_PERFORMANCE_SUMMARY_VIEW DNIS_PERFORMANCE_DETAIL_VIEW WRAP_UP_PERFORMANCE_SUMMARY_VIEW AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW QUEUE_ACTIVITY_SUMMARY_VIEW QUEUE_ACTIVITY_DETAIL_VIEW AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW QUEUE_AGENT_DETAIL_VIEW QUEUE_INTERACTION_DETAIL_VIEW AGENT_SCHEDULE_DETAIL_VIEW IVR_PERFORMANCE_SUMMARY_VIEW IVR_PERFORMANCE_DETAIL_VIEW ANSWER_INSIGHTS_VIEW HANDLE_INSIGHTS_VIEW TALK_INSIGHTS_VIEW HOLD_INSIGHTS_VIEW ACW_INSIGHTS_VIEW WAIT_INSIGHTS_VIEW AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW FLOW_OUTCOME_SUMMARY_VIEW FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW FLOW_DESTINATION_SUMMARY_VIEW FLOW_DESTINATION_DETAIL_VIEW API_USAGE_VIEW]
	ViewType *string `json:"viewType"`
}

// Validate validates this reporting export job request
func (m *ReportingExportJobRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExportFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocale(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectedColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateViewType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var reportingExportJobRequestTypeExportFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CSV","PDF"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobRequestTypeExportFormatPropEnum = append(reportingExportJobRequestTypeExportFormatPropEnum, v)
	}
}

const (

	// ReportingExportJobRequestExportFormatCSV captures enum value "CSV"
	ReportingExportJobRequestExportFormatCSV string = "CSV"

	// ReportingExportJobRequestExportFormatPDF captures enum value "PDF"
	ReportingExportJobRequestExportFormatPDF string = "PDF"
)

// prop value enum
func (m *ReportingExportJobRequest) validateExportFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobRequestTypeExportFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobRequest) validateExportFormat(formats strfmt.Registry) error {

	if err := validate.Required("exportFormat", "body", m.ExportFormat); err != nil {
		return err
	}

	// value enum
	if err := m.validateExportFormatEnum("exportFormat", "body", *m.ExportFormat); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobRequest) validateFilter(formats strfmt.Registry) error {

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

func (m *ReportingExportJobRequest) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobRequest) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobRequest) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobRequest) validateSelectedColumns(formats strfmt.Registry) error {

	if swag.IsZero(m.SelectedColumns) { // not required
		return nil
	}

	for i := 0; i < len(m.SelectedColumns); i++ {
		if swag.IsZero(m.SelectedColumns[i]) { // not required
			continue
		}

		if m.SelectedColumns[i] != nil {
			if err := m.SelectedColumns[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("selectedColumns" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ReportingExportJobRequest) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("timeZone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}

var reportingExportJobRequestTypeViewTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QUEUE_PERFORMANCE_SUMMARY_VIEW","QUEUE_PERFORMANCE_DETAIL_VIEW","INTERACTION_SEARCH_VIEW","AGENT_PERFORMANCE_SUMMARY_VIEW","AGENT_PERFORMANCE_DETAIL_VIEW","AGENT_STATUS_SUMMARY_VIEW","AGENT_STATUS_DETAIL_VIEW","AGENT_EVALUATION_SUMMARY_VIEW","AGENT_EVALUATION_DETAIL_VIEW","AGENT_QUEUE_DETAIL_VIEW","AGENT_INTERACTION_DETAIL_VIEW","ABANDON_INSIGHTS_VIEW","SKILLS_PERFORMANCE_VIEW","SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW","SURVEY_FORM_PERFORMANCE_DETAIL_VIEW","DNIS_PERFORMANCE_SUMMARY_VIEW","DNIS_PERFORMANCE_DETAIL_VIEW","WRAP_UP_PERFORMANCE_SUMMARY_VIEW","AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW","QUEUE_ACTIVITY_SUMMARY_VIEW","QUEUE_ACTIVITY_DETAIL_VIEW","AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW","QUEUE_AGENT_DETAIL_VIEW","QUEUE_INTERACTION_DETAIL_VIEW","AGENT_SCHEDULE_DETAIL_VIEW","IVR_PERFORMANCE_SUMMARY_VIEW","IVR_PERFORMANCE_DETAIL_VIEW","ANSWER_INSIGHTS_VIEW","HANDLE_INSIGHTS_VIEW","TALK_INSIGHTS_VIEW","HOLD_INSIGHTS_VIEW","ACW_INSIGHTS_VIEW","WAIT_INSIGHTS_VIEW","AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW","FLOW_OUTCOME_SUMMARY_VIEW","FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW","FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW","FLOW_DESTINATION_SUMMARY_VIEW","FLOW_DESTINATION_DETAIL_VIEW","API_USAGE_VIEW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobRequestTypeViewTypePropEnum = append(reportingExportJobRequestTypeViewTypePropEnum, v)
	}
}

const (

	// ReportingExportJobRequestViewTypeQUEUEPERFORMANCESUMMARYVIEW captures enum value "QUEUE_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeQUEUEPERFORMANCESUMMARYVIEW string = "QUEUE_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeQUEUEPERFORMANCEDETAILVIEW captures enum value "QUEUE_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeQUEUEPERFORMANCEDETAILVIEW string = "QUEUE_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeINTERACTIONSEARCHVIEW captures enum value "INTERACTION_SEARCH_VIEW"
	ReportingExportJobRequestViewTypeINTERACTIONSEARCHVIEW string = "INTERACTION_SEARCH_VIEW"

	// ReportingExportJobRequestViewTypeAGENTPERFORMANCESUMMARYVIEW captures enum value "AGENT_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeAGENTPERFORMANCESUMMARYVIEW string = "AGENT_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeAGENTPERFORMANCEDETAILVIEW captures enum value "AGENT_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTPERFORMANCEDETAILVIEW string = "AGENT_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeAGENTSTATUSSUMMARYVIEW captures enum value "AGENT_STATUS_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeAGENTSTATUSSUMMARYVIEW string = "AGENT_STATUS_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeAGENTSTATUSDETAILVIEW captures enum value "AGENT_STATUS_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTSTATUSDETAILVIEW string = "AGENT_STATUS_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeAGENTEVALUATIONSUMMARYVIEW captures enum value "AGENT_EVALUATION_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeAGENTEVALUATIONSUMMARYVIEW string = "AGENT_EVALUATION_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeAGENTEVALUATIONDETAILVIEW captures enum value "AGENT_EVALUATION_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTEVALUATIONDETAILVIEW string = "AGENT_EVALUATION_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeAGENTQUEUEDETAILVIEW captures enum value "AGENT_QUEUE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTQUEUEDETAILVIEW string = "AGENT_QUEUE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeAGENTINTERACTIONDETAILVIEW captures enum value "AGENT_INTERACTION_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTINTERACTIONDETAILVIEW string = "AGENT_INTERACTION_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeABANDONINSIGHTSVIEW captures enum value "ABANDON_INSIGHTS_VIEW"
	ReportingExportJobRequestViewTypeABANDONINSIGHTSVIEW string = "ABANDON_INSIGHTS_VIEW"

	// ReportingExportJobRequestViewTypeSKILLSPERFORMANCEVIEW captures enum value "SKILLS_PERFORMANCE_VIEW"
	ReportingExportJobRequestViewTypeSKILLSPERFORMANCEVIEW string = "SKILLS_PERFORMANCE_VIEW"

	// ReportingExportJobRequestViewTypeSURVEYFORMPERFORMANCESUMMARYVIEW captures enum value "SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeSURVEYFORMPERFORMANCESUMMARYVIEW string = "SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeSURVEYFORMPERFORMANCEDETAILVIEW captures enum value "SURVEY_FORM_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeSURVEYFORMPERFORMANCEDETAILVIEW string = "SURVEY_FORM_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeDNISPERFORMANCESUMMARYVIEW captures enum value "DNIS_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeDNISPERFORMANCESUMMARYVIEW string = "DNIS_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeDNISPERFORMANCEDETAILVIEW captures enum value "DNIS_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeDNISPERFORMANCEDETAILVIEW string = "DNIS_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeWRAPUPPERFORMANCESUMMARYVIEW captures enum value "WRAP_UP_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeWRAPUPPERFORMANCESUMMARYVIEW string = "WRAP_UP_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeAGENTWRAPUPPERFORMANCEDETAILVIEW captures enum value "AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTWRAPUPPERFORMANCEDETAILVIEW string = "AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeQUEUEACTIVITYSUMMARYVIEW captures enum value "QUEUE_ACTIVITY_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeQUEUEACTIVITYSUMMARYVIEW string = "QUEUE_ACTIVITY_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeQUEUEACTIVITYDETAILVIEW captures enum value "QUEUE_ACTIVITY_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeQUEUEACTIVITYDETAILVIEW string = "QUEUE_ACTIVITY_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeAGENTQUEUEACTIVITYSUMMARYVIEW captures enum value "AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeAGENTQUEUEACTIVITYSUMMARYVIEW string = "AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeQUEUEAGENTDETAILVIEW captures enum value "QUEUE_AGENT_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeQUEUEAGENTDETAILVIEW string = "QUEUE_AGENT_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeQUEUEINTERACTIONDETAILVIEW captures enum value "QUEUE_INTERACTION_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeQUEUEINTERACTIONDETAILVIEW string = "QUEUE_INTERACTION_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeAGENTSCHEDULEDETAILVIEW captures enum value "AGENT_SCHEDULE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTSCHEDULEDETAILVIEW string = "AGENT_SCHEDULE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeIVRPERFORMANCESUMMARYVIEW captures enum value "IVR_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeIVRPERFORMANCESUMMARYVIEW string = "IVR_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeIVRPERFORMANCEDETAILVIEW captures enum value "IVR_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeIVRPERFORMANCEDETAILVIEW string = "IVR_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeANSWERINSIGHTSVIEW captures enum value "ANSWER_INSIGHTS_VIEW"
	ReportingExportJobRequestViewTypeANSWERINSIGHTSVIEW string = "ANSWER_INSIGHTS_VIEW"

	// ReportingExportJobRequestViewTypeHANDLEINSIGHTSVIEW captures enum value "HANDLE_INSIGHTS_VIEW"
	ReportingExportJobRequestViewTypeHANDLEINSIGHTSVIEW string = "HANDLE_INSIGHTS_VIEW"

	// ReportingExportJobRequestViewTypeTALKINSIGHTSVIEW captures enum value "TALK_INSIGHTS_VIEW"
	ReportingExportJobRequestViewTypeTALKINSIGHTSVIEW string = "TALK_INSIGHTS_VIEW"

	// ReportingExportJobRequestViewTypeHOLDINSIGHTSVIEW captures enum value "HOLD_INSIGHTS_VIEW"
	ReportingExportJobRequestViewTypeHOLDINSIGHTSVIEW string = "HOLD_INSIGHTS_VIEW"

	// ReportingExportJobRequestViewTypeACWINSIGHTSVIEW captures enum value "ACW_INSIGHTS_VIEW"
	ReportingExportJobRequestViewTypeACWINSIGHTSVIEW string = "ACW_INSIGHTS_VIEW"

	// ReportingExportJobRequestViewTypeWAITINSIGHTSVIEW captures enum value "WAIT_INSIGHTS_VIEW"
	ReportingExportJobRequestViewTypeWAITINSIGHTSVIEW string = "WAIT_INSIGHTS_VIEW"

	// ReportingExportJobRequestViewTypeAGENTWRAPUPPERFORMANCEINTERVALDETAILVIEW captures enum value "AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeAGENTWRAPUPPERFORMANCEINTERVALDETAILVIEW string = "AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeFLOWOUTCOMESUMMARYVIEW captures enum value "FLOW_OUTCOME_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeFLOWOUTCOMESUMMARYVIEW string = "FLOW_OUTCOME_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeFLOWOUTCOMEPERFORMANCEDETAILVIEW captures enum value "FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeFLOWOUTCOMEPERFORMANCEDETAILVIEW string = "FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeFLOWOUTCOMEPERFORMANCEINTERVALDETAILVIEW captures enum value "FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeFLOWOUTCOMEPERFORMANCEINTERVALDETAILVIEW string = "FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeFLOWDESTINATIONSUMMARYVIEW captures enum value "FLOW_DESTINATION_SUMMARY_VIEW"
	ReportingExportJobRequestViewTypeFLOWDESTINATIONSUMMARYVIEW string = "FLOW_DESTINATION_SUMMARY_VIEW"

	// ReportingExportJobRequestViewTypeFLOWDESTINATIONDETAILVIEW captures enum value "FLOW_DESTINATION_DETAIL_VIEW"
	ReportingExportJobRequestViewTypeFLOWDESTINATIONDETAILVIEW string = "FLOW_DESTINATION_DETAIL_VIEW"

	// ReportingExportJobRequestViewTypeAPIUSAGEVIEW captures enum value "API_USAGE_VIEW"
	ReportingExportJobRequestViewTypeAPIUSAGEVIEW string = "API_USAGE_VIEW"
)

// prop value enum
func (m *ReportingExportJobRequest) validateViewTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobRequestTypeViewTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobRequest) validateViewType(formats strfmt.Registry) error {

	if err := validate.Required("viewType", "body", m.ViewType); err != nil {
		return err
	}

	// value enum
	if err := m.validateViewTypeEnum("viewType", "body", *m.ViewType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingExportJobRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingExportJobRequest) UnmarshalBinary(b []byte) error {
	var res ReportingExportJobRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
