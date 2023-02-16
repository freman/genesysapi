// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AuditRealtimeQueryRequest audit realtime query request
//
// swagger:model AuditRealtimeQueryRequest
type AuditRealtimeQueryRequest struct {

	// Additional filters for the query.
	Filters []*AuditQueryFilter `json:"filters"`

	// Date and time range of data to query. Intervals are represented as an ISO-8601 string. For example: YYYY-MM-DDThh:mm:ssZ/YYYY-MM-DDThh:mm:ssZ
	// Required: true
	Interval *string `json:"interval"`

	// Page number
	PageNumber int32 `json:"pageNumber,omitempty"`

	// Page size
	PageSize int32 `json:"pageSize,omitempty"`

	// Name of the service to query audits for.
	// Required: true
	// Enum: [AgentConfig AnalyticsReporting Architect Callback Coaching ContactCenter ContentManagement Datatables Directory DynamicSchema Emails EmployeePerformance Gamification Groups Integrations Knowledge LanguageUnderstanding Learning Limits LogCapture Marketplace Messaging NumberPurchasing Outbound PeoplePermissions PredictiveEngagement Presence ProcessAutomation Quality ResponseManagement Routing SCIM Scripter SpeechAndTextAnalytics Supportability Telephony Triggers Voicemail WebDeployments Webhooks WorkforceManagement Workitems]
	ServiceName *string `json:"serviceName"`

	// Sort parameter for the query.
	Sort []*AuditQuerySort `json:"sort"`
}

// Validate validates this audit realtime query request
func (m *AuditRealtimeQueryRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditRealtimeQueryRequest) validateFilters(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditRealtimeQueryRequest) validateInterval(formats strfmt.Registry) error {

	if err := validate.Required("interval", "body", m.Interval); err != nil {
		return err
	}

	return nil
}

var auditRealtimeQueryRequestTypeServiceNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AgentConfig","AnalyticsReporting","Architect","Callback","Coaching","ContactCenter","ContentManagement","Datatables","Directory","DynamicSchema","Emails","EmployeePerformance","Gamification","Groups","Integrations","Knowledge","LanguageUnderstanding","Learning","Limits","LogCapture","Marketplace","Messaging","NumberPurchasing","Outbound","PeoplePermissions","PredictiveEngagement","Presence","ProcessAutomation","Quality","ResponseManagement","Routing","SCIM","Scripter","SpeechAndTextAnalytics","Supportability","Telephony","Triggers","Voicemail","WebDeployments","Webhooks","WorkforceManagement","Workitems"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditRealtimeQueryRequestTypeServiceNamePropEnum = append(auditRealtimeQueryRequestTypeServiceNamePropEnum, v)
	}
}

const (

	// AuditRealtimeQueryRequestServiceNameAgentConfig captures enum value "AgentConfig"
	AuditRealtimeQueryRequestServiceNameAgentConfig string = "AgentConfig"

	// AuditRealtimeQueryRequestServiceNameAnalyticsReporting captures enum value "AnalyticsReporting"
	AuditRealtimeQueryRequestServiceNameAnalyticsReporting string = "AnalyticsReporting"

	// AuditRealtimeQueryRequestServiceNameArchitect captures enum value "Architect"
	AuditRealtimeQueryRequestServiceNameArchitect string = "Architect"

	// AuditRealtimeQueryRequestServiceNameCallback captures enum value "Callback"
	AuditRealtimeQueryRequestServiceNameCallback string = "Callback"

	// AuditRealtimeQueryRequestServiceNameCoaching captures enum value "Coaching"
	AuditRealtimeQueryRequestServiceNameCoaching string = "Coaching"

	// AuditRealtimeQueryRequestServiceNameContactCenter captures enum value "ContactCenter"
	AuditRealtimeQueryRequestServiceNameContactCenter string = "ContactCenter"

	// AuditRealtimeQueryRequestServiceNameContentManagement captures enum value "ContentManagement"
	AuditRealtimeQueryRequestServiceNameContentManagement string = "ContentManagement"

	// AuditRealtimeQueryRequestServiceNameDatatables captures enum value "Datatables"
	AuditRealtimeQueryRequestServiceNameDatatables string = "Datatables"

	// AuditRealtimeQueryRequestServiceNameDirectory captures enum value "Directory"
	AuditRealtimeQueryRequestServiceNameDirectory string = "Directory"

	// AuditRealtimeQueryRequestServiceNameDynamicSchema captures enum value "DynamicSchema"
	AuditRealtimeQueryRequestServiceNameDynamicSchema string = "DynamicSchema"

	// AuditRealtimeQueryRequestServiceNameEmails captures enum value "Emails"
	AuditRealtimeQueryRequestServiceNameEmails string = "Emails"

	// AuditRealtimeQueryRequestServiceNameEmployeePerformance captures enum value "EmployeePerformance"
	AuditRealtimeQueryRequestServiceNameEmployeePerformance string = "EmployeePerformance"

	// AuditRealtimeQueryRequestServiceNameGamification captures enum value "Gamification"
	AuditRealtimeQueryRequestServiceNameGamification string = "Gamification"

	// AuditRealtimeQueryRequestServiceNameGroups captures enum value "Groups"
	AuditRealtimeQueryRequestServiceNameGroups string = "Groups"

	// AuditRealtimeQueryRequestServiceNameIntegrations captures enum value "Integrations"
	AuditRealtimeQueryRequestServiceNameIntegrations string = "Integrations"

	// AuditRealtimeQueryRequestServiceNameKnowledge captures enum value "Knowledge"
	AuditRealtimeQueryRequestServiceNameKnowledge string = "Knowledge"

	// AuditRealtimeQueryRequestServiceNameLanguageUnderstanding captures enum value "LanguageUnderstanding"
	AuditRealtimeQueryRequestServiceNameLanguageUnderstanding string = "LanguageUnderstanding"

	// AuditRealtimeQueryRequestServiceNameLearning captures enum value "Learning"
	AuditRealtimeQueryRequestServiceNameLearning string = "Learning"

	// AuditRealtimeQueryRequestServiceNameLimits captures enum value "Limits"
	AuditRealtimeQueryRequestServiceNameLimits string = "Limits"

	// AuditRealtimeQueryRequestServiceNameLogCapture captures enum value "LogCapture"
	AuditRealtimeQueryRequestServiceNameLogCapture string = "LogCapture"

	// AuditRealtimeQueryRequestServiceNameMarketplace captures enum value "Marketplace"
	AuditRealtimeQueryRequestServiceNameMarketplace string = "Marketplace"

	// AuditRealtimeQueryRequestServiceNameMessaging captures enum value "Messaging"
	AuditRealtimeQueryRequestServiceNameMessaging string = "Messaging"

	// AuditRealtimeQueryRequestServiceNameNumberPurchasing captures enum value "NumberPurchasing"
	AuditRealtimeQueryRequestServiceNameNumberPurchasing string = "NumberPurchasing"

	// AuditRealtimeQueryRequestServiceNameOutbound captures enum value "Outbound"
	AuditRealtimeQueryRequestServiceNameOutbound string = "Outbound"

	// AuditRealtimeQueryRequestServiceNamePeoplePermissions captures enum value "PeoplePermissions"
	AuditRealtimeQueryRequestServiceNamePeoplePermissions string = "PeoplePermissions"

	// AuditRealtimeQueryRequestServiceNamePredictiveEngagement captures enum value "PredictiveEngagement"
	AuditRealtimeQueryRequestServiceNamePredictiveEngagement string = "PredictiveEngagement"

	// AuditRealtimeQueryRequestServiceNamePresence captures enum value "Presence"
	AuditRealtimeQueryRequestServiceNamePresence string = "Presence"

	// AuditRealtimeQueryRequestServiceNameProcessAutomation captures enum value "ProcessAutomation"
	AuditRealtimeQueryRequestServiceNameProcessAutomation string = "ProcessAutomation"

	// AuditRealtimeQueryRequestServiceNameQuality captures enum value "Quality"
	AuditRealtimeQueryRequestServiceNameQuality string = "Quality"

	// AuditRealtimeQueryRequestServiceNameResponseManagement captures enum value "ResponseManagement"
	AuditRealtimeQueryRequestServiceNameResponseManagement string = "ResponseManagement"

	// AuditRealtimeQueryRequestServiceNameRouting captures enum value "Routing"
	AuditRealtimeQueryRequestServiceNameRouting string = "Routing"

	// AuditRealtimeQueryRequestServiceNameSCIM captures enum value "SCIM"
	AuditRealtimeQueryRequestServiceNameSCIM string = "SCIM"

	// AuditRealtimeQueryRequestServiceNameScripter captures enum value "Scripter"
	AuditRealtimeQueryRequestServiceNameScripter string = "Scripter"

	// AuditRealtimeQueryRequestServiceNameSpeechAndTextAnalytics captures enum value "SpeechAndTextAnalytics"
	AuditRealtimeQueryRequestServiceNameSpeechAndTextAnalytics string = "SpeechAndTextAnalytics"

	// AuditRealtimeQueryRequestServiceNameSupportability captures enum value "Supportability"
	AuditRealtimeQueryRequestServiceNameSupportability string = "Supportability"

	// AuditRealtimeQueryRequestServiceNameTelephony captures enum value "Telephony"
	AuditRealtimeQueryRequestServiceNameTelephony string = "Telephony"

	// AuditRealtimeQueryRequestServiceNameTriggers captures enum value "Triggers"
	AuditRealtimeQueryRequestServiceNameTriggers string = "Triggers"

	// AuditRealtimeQueryRequestServiceNameVoicemail captures enum value "Voicemail"
	AuditRealtimeQueryRequestServiceNameVoicemail string = "Voicemail"

	// AuditRealtimeQueryRequestServiceNameWebDeployments captures enum value "WebDeployments"
	AuditRealtimeQueryRequestServiceNameWebDeployments string = "WebDeployments"

	// AuditRealtimeQueryRequestServiceNameWebhooks captures enum value "Webhooks"
	AuditRealtimeQueryRequestServiceNameWebhooks string = "Webhooks"

	// AuditRealtimeQueryRequestServiceNameWorkforceManagement captures enum value "WorkforceManagement"
	AuditRealtimeQueryRequestServiceNameWorkforceManagement string = "WorkforceManagement"

	// AuditRealtimeQueryRequestServiceNameWorkitems captures enum value "Workitems"
	AuditRealtimeQueryRequestServiceNameWorkitems string = "Workitems"
)

// prop value enum
func (m *AuditRealtimeQueryRequest) validateServiceNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditRealtimeQueryRequestTypeServiceNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditRealtimeQueryRequest) validateServiceName(formats strfmt.Registry) error {

	if err := validate.Required("serviceName", "body", m.ServiceName); err != nil {
		return err
	}

	// value enum
	if err := m.validateServiceNameEnum("serviceName", "body", *m.ServiceName); err != nil {
		return err
	}

	return nil
}

func (m *AuditRealtimeQueryRequest) validateSort(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this audit realtime query request based on the context it is used
func (m *AuditRealtimeQueryRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuditRealtimeQueryRequest) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuditRealtimeQueryRequest) contextValidateSort(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Sort); i++ {

		if m.Sort[i] != nil {
			if err := m.Sort[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sort" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("sort" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditRealtimeQueryRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditRealtimeQueryRequest) UnmarshalBinary(b []byte) error {
	var res AuditRealtimeQueryRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
