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

// ReportingExportJobResponse reporting export job response
//
// swagger:model ReportingExportJobResponse
type ReportingExportJobResponse struct {

	// The created date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	CreatedDateTime *strfmt.DateTime `json:"createdDateTime"`

	// The user supplied csv delimiter string value either of type 'comma' or 'semicolon' permitted for the export request
	// Enum: [SEMICOLON COMMA]
	CsvDelimiter string `json:"csvDelimiter,omitempty"`

	// The url to download the request if it's status is completed
	DownloadURL string `json:"downloadUrl,omitempty"`

	// The optional error message in case the export fail to email
	EmailErrorDescription string `json:"emailErrorDescription,omitempty"`

	// The status of individual email addresses as a map
	EmailStatuses map[string]string `json:"emailStatuses,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// Excludes empty rows from the exports
	ExcludeEmptyRows bool `json:"excludeEmptyRows"`

	// The error message in case the export request failed
	// Enum: [FAILED_CONVERTING_EXPORT_JOB FAILED_NO_DATA_EXPORT_JOB_FOUND FAILED_GETTING_DATA_FROM_SERVICE FAILED_GENERATING_TEMP_FILE FAILED_SAVING_FILE_TO_S3 FAILED_NOTIFYING_SKYWALKER_OF_DOWNLOAD FAILED_BUILDING_DOWNLOAD_URL_FROM_SKYWALKER_RESPONSE FAILED_CONVERTING_EXPORT_JOB_TO_QUEUE_PERFORMANCE_JOB EXPORT_TYPE_NOT_IMPLEMENTED REACHED_MAXIMUM_ATTEMPT_OF_RETRY FAILED_LONG_RUNNING_EXPORT TOO_MANY_REQUESTS_FROM_AN_ORGANIZATION FAILED_AS_EXPORT_FILE_SIZE_IS_GREATER_THAN_10MB NOT_AUTHORIZED_TO_VIEW_EXPORT]
	ExportErrorMessagesType string `json:"exportErrorMessagesType,omitempty"`

	// The requested format of the exported data
	// Required: true
	// Enum: [CSV PDF]
	ExportFormat *string `json:"exportFormat"`

	// Filters to apply to create the view
	// Required: true
	Filter *ViewFilter `json:"filter"`

	// Indicates if custom participant attributes will be exported
	HasCustomParticipantAttributes bool `json:"hasCustomParticipantAttributes"`

	// Indicates if durations are formatted in hh:mm:ss format instead of ms
	HasFormatDurations bool `json:"hasFormatDurations"`

	// Indicates if media type will be split in aggregate detail exports
	HasSplitByMedia bool `json:"hasSplitByMedia"`

	// Indicates if filters will be split in aggregate detail exports
	HasSplitFilters bool `json:"hasSplitFilters"`

	// Indicates if summary row needs to be present in exports
	HasSummaryRow bool `json:"hasSummaryRow"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The time period used to limit the the exported data. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ss/YYYY-MM-DDThh:mm:ss
	// Required: true
	Interval *string `json:"interval"`

	// The locale use for localization of the exported data, i.e. en-us, es-mx
	// Required: true
	Locale *string `json:"locale"`

	// The last modified date/time of the request. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	ModifiedDateTime *strfmt.DateTime `json:"modifiedDateTime"`

	// name
	Name string `json:"name,omitempty"`

	// The percentage of the job that has completed processing
	// Required: true
	PercentageComplete *float64 `json:"percentageComplete"`

	// The Period of the request in which to break down the intervals. Periods are represented as an ISO-8601 string. For example: P1D or P1DT12H
	// Required: true
	Period *string `json:"period"`

	// Indicates if the request has been marked as read
	// Required: true
	Read *bool `json:"read"`

	// The list of email recipients for the exports
	RecipientEmails []string `json:"recipientEmails"`

	// The unique run id of the export schedule execute
	// Required: true
	RunID *string `json:"runId"`

	// The list of ordered selected columns from the export view by the user
	SelectedColumns []*SelectedColumns `json:"selectedColumns"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The current status of the export request
	// Required: true
	// Enum: [SUBMITTED RUNNING CANCELLING CANCELLED COMPLETED COMPLETED_WITH_PARTIAL_RESULTS FAILED]
	Status *string `json:"status"`

	// The requested timezone of the exported data. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
	// Required: true
	TimeZone *string `json:"timeZone"`

	// The type of view export job to be created
	// Required: true
	// Enum: [QUEUE_PERFORMANCE_SUMMARY_VIEW QUEUE_PERFORMANCE_DETAIL_VIEW INTERACTION_SEARCH_VIEW AGENT_PERFORMANCE_SUMMARY_VIEW AGENT_PERFORMANCE_DETAIL_VIEW AGENT_STATUS_SUMMARY_VIEW AGENT_STATUS_DETAIL_VIEW AGENT_EVALUATION_SUMMARY_VIEW AGENT_EVALUATION_DETAIL_VIEW AGENT_QUEUE_DETAIL_VIEW AGENT_INTERACTION_DETAIL_VIEW ABANDON_INSIGHTS_VIEW SKILLS_PERFORMANCE_VIEW SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW SURVEY_FORM_PERFORMANCE_DETAIL_VIEW DNIS_PERFORMANCE_SUMMARY_VIEW DNIS_PERFORMANCE_DETAIL_VIEW WRAP_UP_PERFORMANCE_SUMMARY_VIEW AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW QUEUE_ACTIVITY_SUMMARY_VIEW QUEUE_ACTIVITY_DETAIL_VIEW AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW QUEUE_AGENT_DETAIL_VIEW QUEUE_INTERACTION_DETAIL_VIEW AGENT_SCHEDULE_DETAIL_VIEW IVR_PERFORMANCE_SUMMARY_VIEW IVR_PERFORMANCE_DETAIL_VIEW ANSWER_INSIGHTS_VIEW HANDLE_INSIGHTS_VIEW TALK_INSIGHTS_VIEW HOLD_INSIGHTS_VIEW ACW_INSIGHTS_VIEW WAIT_INSIGHTS_VIEW AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW FLOW_OUTCOME_SUMMARY_VIEW FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW FLOW_DESTINATION_SUMMARY_VIEW FLOW_DESTINATION_DETAIL_VIEW API_USAGE_VIEW SCHEDULED_CALLBACKS_VIEW CONTENT_SEARCH_VIEW LANDING_PAGE DASHBOARD_SUMMARY DASHBOARD_DETAIL JOURNEY_ACTION_MAP_SUMMARY_VIEW JOURNEY_OUTCOME_SUMMARY_VIEW JOURNEY_SEGMENT_SUMMARY_VIEW AGENT_DEVELOPMENT_DETAIL_VIEW AGENT_DEVELOPMENT_DETAIL_ME_VIEW AGENT_DEVELOPMENT_SUMMARY_VIEW AGENT_PERFORMANCE_ME_VIEW AGENT_STATUS_ME_VIEW AGENT_EVALUATION_ME_VIEW AGENT_SCORECARD_VIEW AGENT_SCORECARD_ME_VIEW AGENT_GAMIFICATION_LEADERSHIP_VIEW]
	ViewType *string `json:"viewType"`
}

// Validate validates this reporting export job response
func (m *ReportingExportJobResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsvDelimiter(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmailStatuses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportErrorMessagesType(formats); err != nil {
		res = append(res, err)
	}

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

	if err := m.validateModifiedDateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePercentageComplete(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRunID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelectedColumns(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *ReportingExportJobResponse) validateCreatedDateTime(formats strfmt.Registry) error {

	if err := validate.Required("createdDateTime", "body", m.CreatedDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("createdDateTime", "body", "date-time", m.CreatedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var reportingExportJobResponseTypeCsvDelimiterPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SEMICOLON","COMMA"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobResponseTypeCsvDelimiterPropEnum = append(reportingExportJobResponseTypeCsvDelimiterPropEnum, v)
	}
}

const (

	// ReportingExportJobResponseCsvDelimiterSEMICOLON captures enum value "SEMICOLON"
	ReportingExportJobResponseCsvDelimiterSEMICOLON string = "SEMICOLON"

	// ReportingExportJobResponseCsvDelimiterCOMMA captures enum value "COMMA"
	ReportingExportJobResponseCsvDelimiterCOMMA string = "COMMA"
)

// prop value enum
func (m *ReportingExportJobResponse) validateCsvDelimiterEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobResponseTypeCsvDelimiterPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobResponse) validateCsvDelimiter(formats strfmt.Registry) error {

	if swag.IsZero(m.CsvDelimiter) { // not required
		return nil
	}

	// value enum
	if err := m.validateCsvDelimiterEnum("csvDelimiter", "body", m.CsvDelimiter); err != nil {
		return err
	}

	return nil
}

// additional properties value enum
var reportingExportJobResponseEmailStatusesValueEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sent","Pending","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobResponseEmailStatusesValueEnum = append(reportingExportJobResponseEmailStatusesValueEnum, v)
	}
}

func (m *ReportingExportJobResponse) validateEmailStatusesValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobResponseEmailStatusesValueEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobResponse) validateEmailStatuses(formats strfmt.Registry) error {

	if swag.IsZero(m.EmailStatuses) { // not required
		return nil
	}

	for k := range m.EmailStatuses {

		// value enum
		if err := m.validateEmailStatusesValueEnum("emailStatuses"+"."+k, "body", m.EmailStatuses[k]); err != nil {
			return err
		}

	}

	return nil
}

var reportingExportJobResponseTypeExportErrorMessagesTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["FAILED_CONVERTING_EXPORT_JOB","FAILED_NO_DATA_EXPORT_JOB_FOUND","FAILED_GETTING_DATA_FROM_SERVICE","FAILED_GENERATING_TEMP_FILE","FAILED_SAVING_FILE_TO_S3","FAILED_NOTIFYING_SKYWALKER_OF_DOWNLOAD","FAILED_BUILDING_DOWNLOAD_URL_FROM_SKYWALKER_RESPONSE","FAILED_CONVERTING_EXPORT_JOB_TO_QUEUE_PERFORMANCE_JOB","EXPORT_TYPE_NOT_IMPLEMENTED","REACHED_MAXIMUM_ATTEMPT_OF_RETRY","FAILED_LONG_RUNNING_EXPORT","TOO_MANY_REQUESTS_FROM_AN_ORGANIZATION","FAILED_AS_EXPORT_FILE_SIZE_IS_GREATER_THAN_10MB","NOT_AUTHORIZED_TO_VIEW_EXPORT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobResponseTypeExportErrorMessagesTypePropEnum = append(reportingExportJobResponseTypeExportErrorMessagesTypePropEnum, v)
	}
}

const (

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDCONVERTINGEXPORTJOB captures enum value "FAILED_CONVERTING_EXPORT_JOB"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDCONVERTINGEXPORTJOB string = "FAILED_CONVERTING_EXPORT_JOB"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDNODATAEXPORTJOBFOUND captures enum value "FAILED_NO_DATA_EXPORT_JOB_FOUND"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDNODATAEXPORTJOBFOUND string = "FAILED_NO_DATA_EXPORT_JOB_FOUND"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDGETTINGDATAFROMSERVICE captures enum value "FAILED_GETTING_DATA_FROM_SERVICE"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDGETTINGDATAFROMSERVICE string = "FAILED_GETTING_DATA_FROM_SERVICE"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDGENERATINGTEMPFILE captures enum value "FAILED_GENERATING_TEMP_FILE"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDGENERATINGTEMPFILE string = "FAILED_GENERATING_TEMP_FILE"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDSAVINGFILETOS3 captures enum value "FAILED_SAVING_FILE_TO_S3"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDSAVINGFILETOS3 string = "FAILED_SAVING_FILE_TO_S3"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDNOTIFYINGSKYWALKEROFDOWNLOAD captures enum value "FAILED_NOTIFYING_SKYWALKER_OF_DOWNLOAD"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDNOTIFYINGSKYWALKEROFDOWNLOAD string = "FAILED_NOTIFYING_SKYWALKER_OF_DOWNLOAD"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDBUILDINGDOWNLOADURLFROMSKYWALKERRESPONSE captures enum value "FAILED_BUILDING_DOWNLOAD_URL_FROM_SKYWALKER_RESPONSE"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDBUILDINGDOWNLOADURLFROMSKYWALKERRESPONSE string = "FAILED_BUILDING_DOWNLOAD_URL_FROM_SKYWALKER_RESPONSE"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDCONVERTINGEXPORTJOBTOQUEUEPERFORMANCEJOB captures enum value "FAILED_CONVERTING_EXPORT_JOB_TO_QUEUE_PERFORMANCE_JOB"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDCONVERTINGEXPORTJOBTOQUEUEPERFORMANCEJOB string = "FAILED_CONVERTING_EXPORT_JOB_TO_QUEUE_PERFORMANCE_JOB"

	// ReportingExportJobResponseExportErrorMessagesTypeEXPORTTYPENOTIMPLEMENTED captures enum value "EXPORT_TYPE_NOT_IMPLEMENTED"
	ReportingExportJobResponseExportErrorMessagesTypeEXPORTTYPENOTIMPLEMENTED string = "EXPORT_TYPE_NOT_IMPLEMENTED"

	// ReportingExportJobResponseExportErrorMessagesTypeREACHEDMAXIMUMATTEMPTOFRETRY captures enum value "REACHED_MAXIMUM_ATTEMPT_OF_RETRY"
	ReportingExportJobResponseExportErrorMessagesTypeREACHEDMAXIMUMATTEMPTOFRETRY string = "REACHED_MAXIMUM_ATTEMPT_OF_RETRY"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDLONGRUNNINGEXPORT captures enum value "FAILED_LONG_RUNNING_EXPORT"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDLONGRUNNINGEXPORT string = "FAILED_LONG_RUNNING_EXPORT"

	// ReportingExportJobResponseExportErrorMessagesTypeTOOMANYREQUESTSFROMANORGANIZATION captures enum value "TOO_MANY_REQUESTS_FROM_AN_ORGANIZATION"
	ReportingExportJobResponseExportErrorMessagesTypeTOOMANYREQUESTSFROMANORGANIZATION string = "TOO_MANY_REQUESTS_FROM_AN_ORGANIZATION"

	// ReportingExportJobResponseExportErrorMessagesTypeFAILEDASEXPORTFILESIZEISGREATERTHAN10MB captures enum value "FAILED_AS_EXPORT_FILE_SIZE_IS_GREATER_THAN_10MB"
	ReportingExportJobResponseExportErrorMessagesTypeFAILEDASEXPORTFILESIZEISGREATERTHAN10MB string = "FAILED_AS_EXPORT_FILE_SIZE_IS_GREATER_THAN_10MB"

	// ReportingExportJobResponseExportErrorMessagesTypeNOTAUTHORIZEDTOVIEWEXPORT captures enum value "NOT_AUTHORIZED_TO_VIEW_EXPORT"
	ReportingExportJobResponseExportErrorMessagesTypeNOTAUTHORIZEDTOVIEWEXPORT string = "NOT_AUTHORIZED_TO_VIEW_EXPORT"
)

// prop value enum
func (m *ReportingExportJobResponse) validateExportErrorMessagesTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobResponseTypeExportErrorMessagesTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobResponse) validateExportErrorMessagesType(formats strfmt.Registry) error {

	if swag.IsZero(m.ExportErrorMessagesType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExportErrorMessagesTypeEnum("exportErrorMessagesType", "body", m.ExportErrorMessagesType); err != nil {
		return err
	}

	return nil
}

var reportingExportJobResponseTypeExportFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CSV","PDF"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobResponseTypeExportFormatPropEnum = append(reportingExportJobResponseTypeExportFormatPropEnum, v)
	}
}

const (

	// ReportingExportJobResponseExportFormatCSV captures enum value "CSV"
	ReportingExportJobResponseExportFormatCSV string = "CSV"

	// ReportingExportJobResponseExportFormatPDF captures enum value "PDF"
	ReportingExportJobResponseExportFormatPDF string = "PDF"
)

// prop value enum
func (m *ReportingExportJobResponse) validateExportFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobResponseTypeExportFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobResponse) validateExportFormat(formats strfmt.Registry) error {

	if err := validate.Required("exportFormat", "body", m.ExportFormat); err != nil {
		return err
	}

	// value enum
	if err := m.validateExportFormatEnum("exportFormat", "body", *m.ExportFormat); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validateFilter(formats strfmt.Registry) error {

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

func (m *ReportingExportJobResponse) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validateLocale(formats strfmt.Registry) error {

	if err := validate.Required("locale", "body", m.Locale); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validateModifiedDateTime(formats strfmt.Registry) error {

	if err := validate.Required("modifiedDateTime", "body", m.ModifiedDateTime); err != nil {
		return err
	}

	if err := validate.FormatOf("modifiedDateTime", "body", "date-time", m.ModifiedDateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validatePercentageComplete(formats strfmt.Registry) error {

	if err := validate.Required("percentageComplete", "body", m.PercentageComplete); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validatePeriod(formats strfmt.Registry) error {

	if err := validate.Required("period", "body", m.Period); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validateRead(formats strfmt.Registry) error {

	if err := validate.Required("read", "body", m.Read); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validateRunID(formats strfmt.Registry) error {

	if err := validate.Required("runId", "body", m.RunID); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validateSelectedColumns(formats strfmt.Registry) error {

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

func (m *ReportingExportJobResponse) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var reportingExportJobResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUBMITTED","RUNNING","CANCELLING","CANCELLED","COMPLETED","COMPLETED_WITH_PARTIAL_RESULTS","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobResponseTypeStatusPropEnum = append(reportingExportJobResponseTypeStatusPropEnum, v)
	}
}

const (

	// ReportingExportJobResponseStatusSUBMITTED captures enum value "SUBMITTED"
	ReportingExportJobResponseStatusSUBMITTED string = "SUBMITTED"

	// ReportingExportJobResponseStatusRUNNING captures enum value "RUNNING"
	ReportingExportJobResponseStatusRUNNING string = "RUNNING"

	// ReportingExportJobResponseStatusCANCELLING captures enum value "CANCELLING"
	ReportingExportJobResponseStatusCANCELLING string = "CANCELLING"

	// ReportingExportJobResponseStatusCANCELLED captures enum value "CANCELLED"
	ReportingExportJobResponseStatusCANCELLED string = "CANCELLED"

	// ReportingExportJobResponseStatusCOMPLETED captures enum value "COMPLETED"
	ReportingExportJobResponseStatusCOMPLETED string = "COMPLETED"

	// ReportingExportJobResponseStatusCOMPLETEDWITHPARTIALRESULTS captures enum value "COMPLETED_WITH_PARTIAL_RESULTS"
	ReportingExportJobResponseStatusCOMPLETEDWITHPARTIALRESULTS string = "COMPLETED_WITH_PARTIAL_RESULTS"

	// ReportingExportJobResponseStatusFAILED captures enum value "FAILED"
	ReportingExportJobResponseStatusFAILED string = "FAILED"
)

// prop value enum
func (m *ReportingExportJobResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobResponse) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportJobResponse) validateTimeZone(formats strfmt.Registry) error {

	if err := validate.Required("timeZone", "body", m.TimeZone); err != nil {
		return err
	}

	return nil
}

var reportingExportJobResponseTypeViewTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QUEUE_PERFORMANCE_SUMMARY_VIEW","QUEUE_PERFORMANCE_DETAIL_VIEW","INTERACTION_SEARCH_VIEW","AGENT_PERFORMANCE_SUMMARY_VIEW","AGENT_PERFORMANCE_DETAIL_VIEW","AGENT_STATUS_SUMMARY_VIEW","AGENT_STATUS_DETAIL_VIEW","AGENT_EVALUATION_SUMMARY_VIEW","AGENT_EVALUATION_DETAIL_VIEW","AGENT_QUEUE_DETAIL_VIEW","AGENT_INTERACTION_DETAIL_VIEW","ABANDON_INSIGHTS_VIEW","SKILLS_PERFORMANCE_VIEW","SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW","SURVEY_FORM_PERFORMANCE_DETAIL_VIEW","DNIS_PERFORMANCE_SUMMARY_VIEW","DNIS_PERFORMANCE_DETAIL_VIEW","WRAP_UP_PERFORMANCE_SUMMARY_VIEW","AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW","QUEUE_ACTIVITY_SUMMARY_VIEW","QUEUE_ACTIVITY_DETAIL_VIEW","AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW","QUEUE_AGENT_DETAIL_VIEW","QUEUE_INTERACTION_DETAIL_VIEW","AGENT_SCHEDULE_DETAIL_VIEW","IVR_PERFORMANCE_SUMMARY_VIEW","IVR_PERFORMANCE_DETAIL_VIEW","ANSWER_INSIGHTS_VIEW","HANDLE_INSIGHTS_VIEW","TALK_INSIGHTS_VIEW","HOLD_INSIGHTS_VIEW","ACW_INSIGHTS_VIEW","WAIT_INSIGHTS_VIEW","AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW","FLOW_OUTCOME_SUMMARY_VIEW","FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW","FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW","FLOW_DESTINATION_SUMMARY_VIEW","FLOW_DESTINATION_DETAIL_VIEW","API_USAGE_VIEW","SCHEDULED_CALLBACKS_VIEW","CONTENT_SEARCH_VIEW","LANDING_PAGE","DASHBOARD_SUMMARY","DASHBOARD_DETAIL","JOURNEY_ACTION_MAP_SUMMARY_VIEW","JOURNEY_OUTCOME_SUMMARY_VIEW","JOURNEY_SEGMENT_SUMMARY_VIEW","AGENT_DEVELOPMENT_DETAIL_VIEW","AGENT_DEVELOPMENT_DETAIL_ME_VIEW","AGENT_DEVELOPMENT_SUMMARY_VIEW","AGENT_PERFORMANCE_ME_VIEW","AGENT_STATUS_ME_VIEW","AGENT_EVALUATION_ME_VIEW","AGENT_SCORECARD_VIEW","AGENT_SCORECARD_ME_VIEW","AGENT_GAMIFICATION_LEADERSHIP_VIEW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportJobResponseTypeViewTypePropEnum = append(reportingExportJobResponseTypeViewTypePropEnum, v)
	}
}

const (

	// ReportingExportJobResponseViewTypeQUEUEPERFORMANCESUMMARYVIEW captures enum value "QUEUE_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeQUEUEPERFORMANCESUMMARYVIEW string = "QUEUE_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeQUEUEPERFORMANCEDETAILVIEW captures enum value "QUEUE_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeQUEUEPERFORMANCEDETAILVIEW string = "QUEUE_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeINTERACTIONSEARCHVIEW captures enum value "INTERACTION_SEARCH_VIEW"
	ReportingExportJobResponseViewTypeINTERACTIONSEARCHVIEW string = "INTERACTION_SEARCH_VIEW"

	// ReportingExportJobResponseViewTypeAGENTPERFORMANCESUMMARYVIEW captures enum value "AGENT_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeAGENTPERFORMANCESUMMARYVIEW string = "AGENT_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeAGENTPERFORMANCEDETAILVIEW captures enum value "AGENT_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTPERFORMANCEDETAILVIEW string = "AGENT_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAGENTSTATUSSUMMARYVIEW captures enum value "AGENT_STATUS_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeAGENTSTATUSSUMMARYVIEW string = "AGENT_STATUS_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeAGENTSTATUSDETAILVIEW captures enum value "AGENT_STATUS_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTSTATUSDETAILVIEW string = "AGENT_STATUS_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAGENTEVALUATIONSUMMARYVIEW captures enum value "AGENT_EVALUATION_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeAGENTEVALUATIONSUMMARYVIEW string = "AGENT_EVALUATION_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeAGENTEVALUATIONDETAILVIEW captures enum value "AGENT_EVALUATION_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTEVALUATIONDETAILVIEW string = "AGENT_EVALUATION_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAGENTQUEUEDETAILVIEW captures enum value "AGENT_QUEUE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTQUEUEDETAILVIEW string = "AGENT_QUEUE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAGENTINTERACTIONDETAILVIEW captures enum value "AGENT_INTERACTION_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTINTERACTIONDETAILVIEW string = "AGENT_INTERACTION_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeABANDONINSIGHTSVIEW captures enum value "ABANDON_INSIGHTS_VIEW"
	ReportingExportJobResponseViewTypeABANDONINSIGHTSVIEW string = "ABANDON_INSIGHTS_VIEW"

	// ReportingExportJobResponseViewTypeSKILLSPERFORMANCEVIEW captures enum value "SKILLS_PERFORMANCE_VIEW"
	ReportingExportJobResponseViewTypeSKILLSPERFORMANCEVIEW string = "SKILLS_PERFORMANCE_VIEW"

	// ReportingExportJobResponseViewTypeSURVEYFORMPERFORMANCESUMMARYVIEW captures enum value "SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeSURVEYFORMPERFORMANCESUMMARYVIEW string = "SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeSURVEYFORMPERFORMANCEDETAILVIEW captures enum value "SURVEY_FORM_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeSURVEYFORMPERFORMANCEDETAILVIEW string = "SURVEY_FORM_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeDNISPERFORMANCESUMMARYVIEW captures enum value "DNIS_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeDNISPERFORMANCESUMMARYVIEW string = "DNIS_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeDNISPERFORMANCEDETAILVIEW captures enum value "DNIS_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeDNISPERFORMANCEDETAILVIEW string = "DNIS_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeWRAPUPPERFORMANCESUMMARYVIEW captures enum value "WRAP_UP_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeWRAPUPPERFORMANCESUMMARYVIEW string = "WRAP_UP_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeAGENTWRAPUPPERFORMANCEDETAILVIEW captures enum value "AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTWRAPUPPERFORMANCEDETAILVIEW string = "AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeQUEUEACTIVITYSUMMARYVIEW captures enum value "QUEUE_ACTIVITY_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeQUEUEACTIVITYSUMMARYVIEW string = "QUEUE_ACTIVITY_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeQUEUEACTIVITYDETAILVIEW captures enum value "QUEUE_ACTIVITY_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeQUEUEACTIVITYDETAILVIEW string = "QUEUE_ACTIVITY_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAGENTQUEUEACTIVITYSUMMARYVIEW captures enum value "AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeAGENTQUEUEACTIVITYSUMMARYVIEW string = "AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeQUEUEAGENTDETAILVIEW captures enum value "QUEUE_AGENT_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeQUEUEAGENTDETAILVIEW string = "QUEUE_AGENT_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeQUEUEINTERACTIONDETAILVIEW captures enum value "QUEUE_INTERACTION_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeQUEUEINTERACTIONDETAILVIEW string = "QUEUE_INTERACTION_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAGENTSCHEDULEDETAILVIEW captures enum value "AGENT_SCHEDULE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTSCHEDULEDETAILVIEW string = "AGENT_SCHEDULE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeIVRPERFORMANCESUMMARYVIEW captures enum value "IVR_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeIVRPERFORMANCESUMMARYVIEW string = "IVR_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeIVRPERFORMANCEDETAILVIEW captures enum value "IVR_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeIVRPERFORMANCEDETAILVIEW string = "IVR_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeANSWERINSIGHTSVIEW captures enum value "ANSWER_INSIGHTS_VIEW"
	ReportingExportJobResponseViewTypeANSWERINSIGHTSVIEW string = "ANSWER_INSIGHTS_VIEW"

	// ReportingExportJobResponseViewTypeHANDLEINSIGHTSVIEW captures enum value "HANDLE_INSIGHTS_VIEW"
	ReportingExportJobResponseViewTypeHANDLEINSIGHTSVIEW string = "HANDLE_INSIGHTS_VIEW"

	// ReportingExportJobResponseViewTypeTALKINSIGHTSVIEW captures enum value "TALK_INSIGHTS_VIEW"
	ReportingExportJobResponseViewTypeTALKINSIGHTSVIEW string = "TALK_INSIGHTS_VIEW"

	// ReportingExportJobResponseViewTypeHOLDINSIGHTSVIEW captures enum value "HOLD_INSIGHTS_VIEW"
	ReportingExportJobResponseViewTypeHOLDINSIGHTSVIEW string = "HOLD_INSIGHTS_VIEW"

	// ReportingExportJobResponseViewTypeACWINSIGHTSVIEW captures enum value "ACW_INSIGHTS_VIEW"
	ReportingExportJobResponseViewTypeACWINSIGHTSVIEW string = "ACW_INSIGHTS_VIEW"

	// ReportingExportJobResponseViewTypeWAITINSIGHTSVIEW captures enum value "WAIT_INSIGHTS_VIEW"
	ReportingExportJobResponseViewTypeWAITINSIGHTSVIEW string = "WAIT_INSIGHTS_VIEW"

	// ReportingExportJobResponseViewTypeAGENTWRAPUPPERFORMANCEINTERVALDETAILVIEW captures enum value "AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTWRAPUPPERFORMANCEINTERVALDETAILVIEW string = "AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeFLOWOUTCOMESUMMARYVIEW captures enum value "FLOW_OUTCOME_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeFLOWOUTCOMESUMMARYVIEW string = "FLOW_OUTCOME_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeFLOWOUTCOMEPERFORMANCEDETAILVIEW captures enum value "FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeFLOWOUTCOMEPERFORMANCEDETAILVIEW string = "FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeFLOWOUTCOMEPERFORMANCEINTERVALDETAILVIEW captures enum value "FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeFLOWOUTCOMEPERFORMANCEINTERVALDETAILVIEW string = "FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeFLOWDESTINATIONSUMMARYVIEW captures enum value "FLOW_DESTINATION_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeFLOWDESTINATIONSUMMARYVIEW string = "FLOW_DESTINATION_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeFLOWDESTINATIONDETAILVIEW captures enum value "FLOW_DESTINATION_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeFLOWDESTINATIONDETAILVIEW string = "FLOW_DESTINATION_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAPIUSAGEVIEW captures enum value "API_USAGE_VIEW"
	ReportingExportJobResponseViewTypeAPIUSAGEVIEW string = "API_USAGE_VIEW"

	// ReportingExportJobResponseViewTypeSCHEDULEDCALLBACKSVIEW captures enum value "SCHEDULED_CALLBACKS_VIEW"
	ReportingExportJobResponseViewTypeSCHEDULEDCALLBACKSVIEW string = "SCHEDULED_CALLBACKS_VIEW"

	// ReportingExportJobResponseViewTypeCONTENTSEARCHVIEW captures enum value "CONTENT_SEARCH_VIEW"
	ReportingExportJobResponseViewTypeCONTENTSEARCHVIEW string = "CONTENT_SEARCH_VIEW"

	// ReportingExportJobResponseViewTypeLANDINGPAGE captures enum value "LANDING_PAGE"
	ReportingExportJobResponseViewTypeLANDINGPAGE string = "LANDING_PAGE"

	// ReportingExportJobResponseViewTypeDASHBOARDSUMMARY captures enum value "DASHBOARD_SUMMARY"
	ReportingExportJobResponseViewTypeDASHBOARDSUMMARY string = "DASHBOARD_SUMMARY"

	// ReportingExportJobResponseViewTypeDASHBOARDDETAIL captures enum value "DASHBOARD_DETAIL"
	ReportingExportJobResponseViewTypeDASHBOARDDETAIL string = "DASHBOARD_DETAIL"

	// ReportingExportJobResponseViewTypeJOURNEYACTIONMAPSUMMARYVIEW captures enum value "JOURNEY_ACTION_MAP_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeJOURNEYACTIONMAPSUMMARYVIEW string = "JOURNEY_ACTION_MAP_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeJOURNEYOUTCOMESUMMARYVIEW captures enum value "JOURNEY_OUTCOME_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeJOURNEYOUTCOMESUMMARYVIEW string = "JOURNEY_OUTCOME_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeJOURNEYSEGMENTSUMMARYVIEW captures enum value "JOURNEY_SEGMENT_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeJOURNEYSEGMENTSUMMARYVIEW string = "JOURNEY_SEGMENT_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeAGENTDEVELOPMENTDETAILVIEW captures enum value "AGENT_DEVELOPMENT_DETAIL_VIEW"
	ReportingExportJobResponseViewTypeAGENTDEVELOPMENTDETAILVIEW string = "AGENT_DEVELOPMENT_DETAIL_VIEW"

	// ReportingExportJobResponseViewTypeAGENTDEVELOPMENTDETAILMEVIEW captures enum value "AGENT_DEVELOPMENT_DETAIL_ME_VIEW"
	ReportingExportJobResponseViewTypeAGENTDEVELOPMENTDETAILMEVIEW string = "AGENT_DEVELOPMENT_DETAIL_ME_VIEW"

	// ReportingExportJobResponseViewTypeAGENTDEVELOPMENTSUMMARYVIEW captures enum value "AGENT_DEVELOPMENT_SUMMARY_VIEW"
	ReportingExportJobResponseViewTypeAGENTDEVELOPMENTSUMMARYVIEW string = "AGENT_DEVELOPMENT_SUMMARY_VIEW"

	// ReportingExportJobResponseViewTypeAGENTPERFORMANCEMEVIEW captures enum value "AGENT_PERFORMANCE_ME_VIEW"
	ReportingExportJobResponseViewTypeAGENTPERFORMANCEMEVIEW string = "AGENT_PERFORMANCE_ME_VIEW"

	// ReportingExportJobResponseViewTypeAGENTSTATUSMEVIEW captures enum value "AGENT_STATUS_ME_VIEW"
	ReportingExportJobResponseViewTypeAGENTSTATUSMEVIEW string = "AGENT_STATUS_ME_VIEW"

	// ReportingExportJobResponseViewTypeAGENTEVALUATIONMEVIEW captures enum value "AGENT_EVALUATION_ME_VIEW"
	ReportingExportJobResponseViewTypeAGENTEVALUATIONMEVIEW string = "AGENT_EVALUATION_ME_VIEW"

	// ReportingExportJobResponseViewTypeAGENTSCORECARDVIEW captures enum value "AGENT_SCORECARD_VIEW"
	ReportingExportJobResponseViewTypeAGENTSCORECARDVIEW string = "AGENT_SCORECARD_VIEW"

	// ReportingExportJobResponseViewTypeAGENTSCORECARDMEVIEW captures enum value "AGENT_SCORECARD_ME_VIEW"
	ReportingExportJobResponseViewTypeAGENTSCORECARDMEVIEW string = "AGENT_SCORECARD_ME_VIEW"

	// ReportingExportJobResponseViewTypeAGENTGAMIFICATIONLEADERSHIPVIEW captures enum value "AGENT_GAMIFICATION_LEADERSHIP_VIEW"
	ReportingExportJobResponseViewTypeAGENTGAMIFICATIONLEADERSHIPVIEW string = "AGENT_GAMIFICATION_LEADERSHIP_VIEW"
)

// prop value enum
func (m *ReportingExportJobResponse) validateViewTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportJobResponseTypeViewTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportJobResponse) validateViewType(formats strfmt.Registry) error {

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
func (m *ReportingExportJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingExportJobResponse) UnmarshalBinary(b []byte) error {
	var res ReportingExportJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
