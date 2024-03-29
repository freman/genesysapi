// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportingExportMetadataJobResponse reporting export metadata job response
//
// swagger:model ReportingExportMetadataJobResponse
type ReportingExportMetadataJobResponse struct {

	// The list of available column ids for the export metadata
	AvailableColumnIds []string `json:"availableColumnIds"`

	// The date limitations of the export metadata
	DateLimitations string `json:"dateLimitations,omitempty"`

	// The list of dependent column ids for the export metadata
	DependentColumnIds map[string][]string `json:"dependentColumnIds,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The list of required column ids for the export metadata
	RequiredColumnIds []string `json:"requiredColumnIds"`

	// The list of required filters for the export metadata
	RequiredFilters []string `json:"requiredFilters"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The list of supported filters for the export metadata
	SupportedFilters []string `json:"supportedFilters"`

	// The view type of the export metadata
	// Enum: [QUEUE_PERFORMANCE_SUMMARY_VIEW QUEUE_PERFORMANCE_DETAIL_VIEW INTERACTION_SEARCH_VIEW AGENT_PERFORMANCE_SUMMARY_VIEW AGENT_PERFORMANCE_DETAIL_VIEW AGENT_STATUS_SUMMARY_VIEW AGENT_STATUS_DETAIL_VIEW AGENT_EVALUATION_SUMMARY_VIEW AGENT_EVALUATION_DETAIL_VIEW AGENT_QUEUE_DETAIL_VIEW AGENT_INTERACTION_DETAIL_VIEW ABANDON_INSIGHTS_VIEW SKILLS_PERFORMANCE_VIEW SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW SURVEY_FORM_PERFORMANCE_DETAIL_VIEW DNIS_PERFORMANCE_SUMMARY_VIEW DNIS_PERFORMANCE_DETAIL_VIEW WRAP_UP_PERFORMANCE_SUMMARY_VIEW AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW QUEUE_ACTIVITY_SUMMARY_VIEW QUEUE_ACTIVITY_DETAIL_VIEW AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW QUEUE_AGENT_DETAIL_VIEW QUEUE_INTERACTION_DETAIL_VIEW AGENT_SCHEDULE_DETAIL_VIEW IVR_PERFORMANCE_SUMMARY_VIEW IVR_PERFORMANCE_DETAIL_VIEW ANSWER_INSIGHTS_VIEW HANDLE_INSIGHTS_VIEW TALK_INSIGHTS_VIEW HOLD_INSIGHTS_VIEW ACW_INSIGHTS_VIEW WAIT_INSIGHTS_VIEW AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW FLOW_OUTCOME_SUMMARY_VIEW FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW FLOW_DESTINATION_SUMMARY_VIEW FLOW_DESTINATION_DETAIL_VIEW API_USAGE_VIEW SCHEDULED_CALLBACKS_VIEW CONTENT_SEARCH_VIEW LANDING_PAGE DASHBOARD_SUMMARY DASHBOARD_DETAIL JOURNEY_ACTION_MAP_SUMMARY_VIEW JOURNEY_OUTCOME_SUMMARY_VIEW JOURNEY_SEGMENT_SUMMARY_VIEW AGENT_DEVELOPMENT_DETAIL_VIEW AGENT_DEVELOPMENT_DETAIL_ME_VIEW AGENT_DEVELOPMENT_SUMMARY_VIEW AGENT_PERFORMANCE_ME_VIEW AGENT_STATUS_ME_VIEW AGENT_EVALUATION_ME_VIEW AGENT_SCORECARD_VIEW AGENT_SCORECARD_ME_VIEW AGENT_GAMIFICATION_LEADERSHIP_VIEW AGENT_SCHEDULE_ME_VIEW BOT_PERFORMANCE_SUMMARY_VIEW BOT_PERFORMANCE_DETAIL_VIEW SCHEDULED_EXPORTS_VIEW TOPIC_TREND_SUMMARY_VIEW TOPIC_TREND_DETAIL_VIEW ACTION_MAP_BLOCKED_CONSTRAINTS_DETAIL_VIEW ACTION_MAP_BLOCKED_CONSTRAINTS_INTERVAL_DETAIL_VIEW FLOW_MILESTONE_PERFORMANCE_DETAIL_VIEW FLOW_MILESTONE_PERFORMANCE_INTERVAL_DETAIL_VIEW AGENT_TOPIC_SUMMARY_VIEW AGENT_TOPIC_DETAIL_VIEW QUEUE_TOPIC_SUMMARY_VIEW QUEUE_TOPIC_DETAIL_VIEW FLOW_TOPIC_SUMMARY_VIEW FLOW_TOPIC_DETAIL_VIEW AGENT_INTERACTIONS_ME_VIEW ALERT_RULES_VIEW CONFIGURE_ALERT_RULE_VIEW PREDICTIVE_ROUTING_VIEW PREDICTIVE_ROUTING_QUEUE_OVERVIEW PREDICTIVE_ROUTING_MODEL_VIEW PREDICTIVE_ROUTING_IMPACT_VIEW DATA_ACTIONS_PERFORMANCE_SUMMARY_VIEW DATA_ACTIONS_PERFORMANCE_DETAIL_VIEW AGENT_TIMELINE_SUMMARY_VIEW AGENT_TIMELINE_DETAIL_VIEW AGENT_LOGIN_LOGOUT_SUMMARY_VIEW AGENT_LOGIN_LOGOUT_DETAIL_VIEW CAMPAIGN_PERFORMANCE_SUMMARY_VIEW CAMPAIGN_PERFORMANCE_DETAIL_VIEW KNOWLEDGE_PERFORMANCE_VIEW AGENT_SCORECARD_INSIGHTS_SUMMARY_VIEW AGENT_SCORECARD_INSIGHTS_DETAIL_VIEW QUEUE_WRAPUP_DETAIL_VIEW INTERACTION_DETAIL_VIEW CAMPAIGN_INTERACTION_DETAIL_VIEW CAMPAIGN_ATTEMPT_DETAIL_VIEW WORKITEM_PERFORMANCE_SUMMARY_VIEW]
	ViewType string `json:"viewType,omitempty"`
}

// Validate validates this reporting export metadata job response
func (m *ReportingExportMetadataJobResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *ReportingExportMetadataJobResponse) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var reportingExportMetadataJobResponseTypeViewTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["QUEUE_PERFORMANCE_SUMMARY_VIEW","QUEUE_PERFORMANCE_DETAIL_VIEW","INTERACTION_SEARCH_VIEW","AGENT_PERFORMANCE_SUMMARY_VIEW","AGENT_PERFORMANCE_DETAIL_VIEW","AGENT_STATUS_SUMMARY_VIEW","AGENT_STATUS_DETAIL_VIEW","AGENT_EVALUATION_SUMMARY_VIEW","AGENT_EVALUATION_DETAIL_VIEW","AGENT_QUEUE_DETAIL_VIEW","AGENT_INTERACTION_DETAIL_VIEW","ABANDON_INSIGHTS_VIEW","SKILLS_PERFORMANCE_VIEW","SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW","SURVEY_FORM_PERFORMANCE_DETAIL_VIEW","DNIS_PERFORMANCE_SUMMARY_VIEW","DNIS_PERFORMANCE_DETAIL_VIEW","WRAP_UP_PERFORMANCE_SUMMARY_VIEW","AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW","QUEUE_ACTIVITY_SUMMARY_VIEW","QUEUE_ACTIVITY_DETAIL_VIEW","AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW","QUEUE_AGENT_DETAIL_VIEW","QUEUE_INTERACTION_DETAIL_VIEW","AGENT_SCHEDULE_DETAIL_VIEW","IVR_PERFORMANCE_SUMMARY_VIEW","IVR_PERFORMANCE_DETAIL_VIEW","ANSWER_INSIGHTS_VIEW","HANDLE_INSIGHTS_VIEW","TALK_INSIGHTS_VIEW","HOLD_INSIGHTS_VIEW","ACW_INSIGHTS_VIEW","WAIT_INSIGHTS_VIEW","AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW","FLOW_OUTCOME_SUMMARY_VIEW","FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW","FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW","FLOW_DESTINATION_SUMMARY_VIEW","FLOW_DESTINATION_DETAIL_VIEW","API_USAGE_VIEW","SCHEDULED_CALLBACKS_VIEW","CONTENT_SEARCH_VIEW","LANDING_PAGE","DASHBOARD_SUMMARY","DASHBOARD_DETAIL","JOURNEY_ACTION_MAP_SUMMARY_VIEW","JOURNEY_OUTCOME_SUMMARY_VIEW","JOURNEY_SEGMENT_SUMMARY_VIEW","AGENT_DEVELOPMENT_DETAIL_VIEW","AGENT_DEVELOPMENT_DETAIL_ME_VIEW","AGENT_DEVELOPMENT_SUMMARY_VIEW","AGENT_PERFORMANCE_ME_VIEW","AGENT_STATUS_ME_VIEW","AGENT_EVALUATION_ME_VIEW","AGENT_SCORECARD_VIEW","AGENT_SCORECARD_ME_VIEW","AGENT_GAMIFICATION_LEADERSHIP_VIEW","AGENT_SCHEDULE_ME_VIEW","BOT_PERFORMANCE_SUMMARY_VIEW","BOT_PERFORMANCE_DETAIL_VIEW","SCHEDULED_EXPORTS_VIEW","TOPIC_TREND_SUMMARY_VIEW","TOPIC_TREND_DETAIL_VIEW","ACTION_MAP_BLOCKED_CONSTRAINTS_DETAIL_VIEW","ACTION_MAP_BLOCKED_CONSTRAINTS_INTERVAL_DETAIL_VIEW","FLOW_MILESTONE_PERFORMANCE_DETAIL_VIEW","FLOW_MILESTONE_PERFORMANCE_INTERVAL_DETAIL_VIEW","AGENT_TOPIC_SUMMARY_VIEW","AGENT_TOPIC_DETAIL_VIEW","QUEUE_TOPIC_SUMMARY_VIEW","QUEUE_TOPIC_DETAIL_VIEW","FLOW_TOPIC_SUMMARY_VIEW","FLOW_TOPIC_DETAIL_VIEW","AGENT_INTERACTIONS_ME_VIEW","ALERT_RULES_VIEW","CONFIGURE_ALERT_RULE_VIEW","PREDICTIVE_ROUTING_VIEW","PREDICTIVE_ROUTING_QUEUE_OVERVIEW","PREDICTIVE_ROUTING_MODEL_VIEW","PREDICTIVE_ROUTING_IMPACT_VIEW","DATA_ACTIONS_PERFORMANCE_SUMMARY_VIEW","DATA_ACTIONS_PERFORMANCE_DETAIL_VIEW","AGENT_TIMELINE_SUMMARY_VIEW","AGENT_TIMELINE_DETAIL_VIEW","AGENT_LOGIN_LOGOUT_SUMMARY_VIEW","AGENT_LOGIN_LOGOUT_DETAIL_VIEW","CAMPAIGN_PERFORMANCE_SUMMARY_VIEW","CAMPAIGN_PERFORMANCE_DETAIL_VIEW","KNOWLEDGE_PERFORMANCE_VIEW","AGENT_SCORECARD_INSIGHTS_SUMMARY_VIEW","AGENT_SCORECARD_INSIGHTS_DETAIL_VIEW","QUEUE_WRAPUP_DETAIL_VIEW","INTERACTION_DETAIL_VIEW","CAMPAIGN_INTERACTION_DETAIL_VIEW","CAMPAIGN_ATTEMPT_DETAIL_VIEW","WORKITEM_PERFORMANCE_SUMMARY_VIEW"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportingExportMetadataJobResponseTypeViewTypePropEnum = append(reportingExportMetadataJobResponseTypeViewTypePropEnum, v)
	}
}

const (

	// ReportingExportMetadataJobResponseViewTypeQUEUEPERFORMANCESUMMARYVIEW captures enum value "QUEUE_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUEPERFORMANCESUMMARYVIEW string = "QUEUE_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUEPERFORMANCEDETAILVIEW captures enum value "QUEUE_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUEPERFORMANCEDETAILVIEW string = "QUEUE_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeINTERACTIONSEARCHVIEW captures enum value "INTERACTION_SEARCH_VIEW"
	ReportingExportMetadataJobResponseViewTypeINTERACTIONSEARCHVIEW string = "INTERACTION_SEARCH_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTPERFORMANCESUMMARYVIEW captures enum value "AGENT_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTPERFORMANCESUMMARYVIEW string = "AGENT_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTPERFORMANCEDETAILVIEW captures enum value "AGENT_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTPERFORMANCEDETAILVIEW string = "AGENT_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSTATUSSUMMARYVIEW captures enum value "AGENT_STATUS_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSTATUSSUMMARYVIEW string = "AGENT_STATUS_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSTATUSDETAILVIEW captures enum value "AGENT_STATUS_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSTATUSDETAILVIEW string = "AGENT_STATUS_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTEVALUATIONSUMMARYVIEW captures enum value "AGENT_EVALUATION_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTEVALUATIONSUMMARYVIEW string = "AGENT_EVALUATION_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTEVALUATIONDETAILVIEW captures enum value "AGENT_EVALUATION_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTEVALUATIONDETAILVIEW string = "AGENT_EVALUATION_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTQUEUEDETAILVIEW captures enum value "AGENT_QUEUE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTQUEUEDETAILVIEW string = "AGENT_QUEUE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTINTERACTIONDETAILVIEW captures enum value "AGENT_INTERACTION_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTINTERACTIONDETAILVIEW string = "AGENT_INTERACTION_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeABANDONINSIGHTSVIEW captures enum value "ABANDON_INSIGHTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeABANDONINSIGHTSVIEW string = "ABANDON_INSIGHTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeSKILLSPERFORMANCEVIEW captures enum value "SKILLS_PERFORMANCE_VIEW"
	ReportingExportMetadataJobResponseViewTypeSKILLSPERFORMANCEVIEW string = "SKILLS_PERFORMANCE_VIEW"

	// ReportingExportMetadataJobResponseViewTypeSURVEYFORMPERFORMANCESUMMARYVIEW captures enum value "SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeSURVEYFORMPERFORMANCESUMMARYVIEW string = "SURVEY_FORM_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeSURVEYFORMPERFORMANCEDETAILVIEW captures enum value "SURVEY_FORM_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeSURVEYFORMPERFORMANCEDETAILVIEW string = "SURVEY_FORM_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeDNISPERFORMANCESUMMARYVIEW captures enum value "DNIS_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeDNISPERFORMANCESUMMARYVIEW string = "DNIS_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeDNISPERFORMANCEDETAILVIEW captures enum value "DNIS_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeDNISPERFORMANCEDETAILVIEW string = "DNIS_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeWRAPUPPERFORMANCESUMMARYVIEW captures enum value "WRAP_UP_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeWRAPUPPERFORMANCESUMMARYVIEW string = "WRAP_UP_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTWRAPUPPERFORMANCEDETAILVIEW captures enum value "AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTWRAPUPPERFORMANCEDETAILVIEW string = "AGENT_WRAP_UP_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUEACTIVITYSUMMARYVIEW captures enum value "QUEUE_ACTIVITY_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUEACTIVITYSUMMARYVIEW string = "QUEUE_ACTIVITY_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUEACTIVITYDETAILVIEW captures enum value "QUEUE_ACTIVITY_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUEACTIVITYDETAILVIEW string = "QUEUE_ACTIVITY_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTQUEUEACTIVITYSUMMARYVIEW captures enum value "AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTQUEUEACTIVITYSUMMARYVIEW string = "AGENT_QUEUE_ACTIVITY_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUEAGENTDETAILVIEW captures enum value "QUEUE_AGENT_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUEAGENTDETAILVIEW string = "QUEUE_AGENT_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUEINTERACTIONDETAILVIEW captures enum value "QUEUE_INTERACTION_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUEINTERACTIONDETAILVIEW string = "QUEUE_INTERACTION_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSCHEDULEDETAILVIEW captures enum value "AGENT_SCHEDULE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSCHEDULEDETAILVIEW string = "AGENT_SCHEDULE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeIVRPERFORMANCESUMMARYVIEW captures enum value "IVR_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeIVRPERFORMANCESUMMARYVIEW string = "IVR_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeIVRPERFORMANCEDETAILVIEW captures enum value "IVR_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeIVRPERFORMANCEDETAILVIEW string = "IVR_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeANSWERINSIGHTSVIEW captures enum value "ANSWER_INSIGHTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeANSWERINSIGHTSVIEW string = "ANSWER_INSIGHTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeHANDLEINSIGHTSVIEW captures enum value "HANDLE_INSIGHTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeHANDLEINSIGHTSVIEW string = "HANDLE_INSIGHTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeTALKINSIGHTSVIEW captures enum value "TALK_INSIGHTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeTALKINSIGHTSVIEW string = "TALK_INSIGHTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeHOLDINSIGHTSVIEW captures enum value "HOLD_INSIGHTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeHOLDINSIGHTSVIEW string = "HOLD_INSIGHTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeACWINSIGHTSVIEW captures enum value "ACW_INSIGHTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeACWINSIGHTSVIEW string = "ACW_INSIGHTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeWAITINSIGHTSVIEW captures enum value "WAIT_INSIGHTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeWAITINSIGHTSVIEW string = "WAIT_INSIGHTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTWRAPUPPERFORMANCEINTERVALDETAILVIEW captures enum value "AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTWRAPUPPERFORMANCEINTERVALDETAILVIEW string = "AGENT_WRAP_UP_PERFORMANCE_INTERVAL_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWOUTCOMESUMMARYVIEW captures enum value "FLOW_OUTCOME_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWOUTCOMESUMMARYVIEW string = "FLOW_OUTCOME_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWOUTCOMEPERFORMANCEDETAILVIEW captures enum value "FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWOUTCOMEPERFORMANCEDETAILVIEW string = "FLOW_OUTCOME_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWOUTCOMEPERFORMANCEINTERVALDETAILVIEW captures enum value "FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWOUTCOMEPERFORMANCEINTERVALDETAILVIEW string = "FLOW_OUTCOME_PERFORMANCE_INTERVAL_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWDESTINATIONSUMMARYVIEW captures enum value "FLOW_DESTINATION_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWDESTINATIONSUMMARYVIEW string = "FLOW_DESTINATION_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWDESTINATIONDETAILVIEW captures enum value "FLOW_DESTINATION_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWDESTINATIONDETAILVIEW string = "FLOW_DESTINATION_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAPIUSAGEVIEW captures enum value "API_USAGE_VIEW"
	ReportingExportMetadataJobResponseViewTypeAPIUSAGEVIEW string = "API_USAGE_VIEW"

	// ReportingExportMetadataJobResponseViewTypeSCHEDULEDCALLBACKSVIEW captures enum value "SCHEDULED_CALLBACKS_VIEW"
	ReportingExportMetadataJobResponseViewTypeSCHEDULEDCALLBACKSVIEW string = "SCHEDULED_CALLBACKS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeCONTENTSEARCHVIEW captures enum value "CONTENT_SEARCH_VIEW"
	ReportingExportMetadataJobResponseViewTypeCONTENTSEARCHVIEW string = "CONTENT_SEARCH_VIEW"

	// ReportingExportMetadataJobResponseViewTypeLANDINGPAGE captures enum value "LANDING_PAGE"
	ReportingExportMetadataJobResponseViewTypeLANDINGPAGE string = "LANDING_PAGE"

	// ReportingExportMetadataJobResponseViewTypeDASHBOARDSUMMARY captures enum value "DASHBOARD_SUMMARY"
	ReportingExportMetadataJobResponseViewTypeDASHBOARDSUMMARY string = "DASHBOARD_SUMMARY"

	// ReportingExportMetadataJobResponseViewTypeDASHBOARDDETAIL captures enum value "DASHBOARD_DETAIL"
	ReportingExportMetadataJobResponseViewTypeDASHBOARDDETAIL string = "DASHBOARD_DETAIL"

	// ReportingExportMetadataJobResponseViewTypeJOURNEYACTIONMAPSUMMARYVIEW captures enum value "JOURNEY_ACTION_MAP_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeJOURNEYACTIONMAPSUMMARYVIEW string = "JOURNEY_ACTION_MAP_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeJOURNEYOUTCOMESUMMARYVIEW captures enum value "JOURNEY_OUTCOME_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeJOURNEYOUTCOMESUMMARYVIEW string = "JOURNEY_OUTCOME_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeJOURNEYSEGMENTSUMMARYVIEW captures enum value "JOURNEY_SEGMENT_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeJOURNEYSEGMENTSUMMARYVIEW string = "JOURNEY_SEGMENT_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTDEVELOPMENTDETAILVIEW captures enum value "AGENT_DEVELOPMENT_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTDEVELOPMENTDETAILVIEW string = "AGENT_DEVELOPMENT_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTDEVELOPMENTDETAILMEVIEW captures enum value "AGENT_DEVELOPMENT_DETAIL_ME_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTDEVELOPMENTDETAILMEVIEW string = "AGENT_DEVELOPMENT_DETAIL_ME_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTDEVELOPMENTSUMMARYVIEW captures enum value "AGENT_DEVELOPMENT_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTDEVELOPMENTSUMMARYVIEW string = "AGENT_DEVELOPMENT_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTPERFORMANCEMEVIEW captures enum value "AGENT_PERFORMANCE_ME_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTPERFORMANCEMEVIEW string = "AGENT_PERFORMANCE_ME_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSTATUSMEVIEW captures enum value "AGENT_STATUS_ME_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSTATUSMEVIEW string = "AGENT_STATUS_ME_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTEVALUATIONMEVIEW captures enum value "AGENT_EVALUATION_ME_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTEVALUATIONMEVIEW string = "AGENT_EVALUATION_ME_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDVIEW captures enum value "AGENT_SCORECARD_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDVIEW string = "AGENT_SCORECARD_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDMEVIEW captures enum value "AGENT_SCORECARD_ME_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDMEVIEW string = "AGENT_SCORECARD_ME_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTGAMIFICATIONLEADERSHIPVIEW captures enum value "AGENT_GAMIFICATION_LEADERSHIP_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTGAMIFICATIONLEADERSHIPVIEW string = "AGENT_GAMIFICATION_LEADERSHIP_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSCHEDULEMEVIEW captures enum value "AGENT_SCHEDULE_ME_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSCHEDULEMEVIEW string = "AGENT_SCHEDULE_ME_VIEW"

	// ReportingExportMetadataJobResponseViewTypeBOTPERFORMANCESUMMARYVIEW captures enum value "BOT_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeBOTPERFORMANCESUMMARYVIEW string = "BOT_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeBOTPERFORMANCEDETAILVIEW captures enum value "BOT_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeBOTPERFORMANCEDETAILVIEW string = "BOT_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeSCHEDULEDEXPORTSVIEW captures enum value "SCHEDULED_EXPORTS_VIEW"
	ReportingExportMetadataJobResponseViewTypeSCHEDULEDEXPORTSVIEW string = "SCHEDULED_EXPORTS_VIEW"

	// ReportingExportMetadataJobResponseViewTypeTOPICTRENDSUMMARYVIEW captures enum value "TOPIC_TREND_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeTOPICTRENDSUMMARYVIEW string = "TOPIC_TREND_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeTOPICTRENDDETAILVIEW captures enum value "TOPIC_TREND_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeTOPICTRENDDETAILVIEW string = "TOPIC_TREND_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeACTIONMAPBLOCKEDCONSTRAINTSDETAILVIEW captures enum value "ACTION_MAP_BLOCKED_CONSTRAINTS_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeACTIONMAPBLOCKEDCONSTRAINTSDETAILVIEW string = "ACTION_MAP_BLOCKED_CONSTRAINTS_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeACTIONMAPBLOCKEDCONSTRAINTSINTERVALDETAILVIEW captures enum value "ACTION_MAP_BLOCKED_CONSTRAINTS_INTERVAL_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeACTIONMAPBLOCKEDCONSTRAINTSINTERVALDETAILVIEW string = "ACTION_MAP_BLOCKED_CONSTRAINTS_INTERVAL_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWMILESTONEPERFORMANCEDETAILVIEW captures enum value "FLOW_MILESTONE_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWMILESTONEPERFORMANCEDETAILVIEW string = "FLOW_MILESTONE_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWMILESTONEPERFORMANCEINTERVALDETAILVIEW captures enum value "FLOW_MILESTONE_PERFORMANCE_INTERVAL_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWMILESTONEPERFORMANCEINTERVALDETAILVIEW string = "FLOW_MILESTONE_PERFORMANCE_INTERVAL_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTTOPICSUMMARYVIEW captures enum value "AGENT_TOPIC_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTTOPICSUMMARYVIEW string = "AGENT_TOPIC_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTTOPICDETAILVIEW captures enum value "AGENT_TOPIC_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTTOPICDETAILVIEW string = "AGENT_TOPIC_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUETOPICSUMMARYVIEW captures enum value "QUEUE_TOPIC_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUETOPICSUMMARYVIEW string = "QUEUE_TOPIC_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUETOPICDETAILVIEW captures enum value "QUEUE_TOPIC_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUETOPICDETAILVIEW string = "QUEUE_TOPIC_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWTOPICSUMMARYVIEW captures enum value "FLOW_TOPIC_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWTOPICSUMMARYVIEW string = "FLOW_TOPIC_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeFLOWTOPICDETAILVIEW captures enum value "FLOW_TOPIC_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeFLOWTOPICDETAILVIEW string = "FLOW_TOPIC_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTINTERACTIONSMEVIEW captures enum value "AGENT_INTERACTIONS_ME_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTINTERACTIONSMEVIEW string = "AGENT_INTERACTIONS_ME_VIEW"

	// ReportingExportMetadataJobResponseViewTypeALERTRULESVIEW captures enum value "ALERT_RULES_VIEW"
	ReportingExportMetadataJobResponseViewTypeALERTRULESVIEW string = "ALERT_RULES_VIEW"

	// ReportingExportMetadataJobResponseViewTypeCONFIGUREALERTRULEVIEW captures enum value "CONFIGURE_ALERT_RULE_VIEW"
	ReportingExportMetadataJobResponseViewTypeCONFIGUREALERTRULEVIEW string = "CONFIGURE_ALERT_RULE_VIEW"

	// ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGVIEW captures enum value "PREDICTIVE_ROUTING_VIEW"
	ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGVIEW string = "PREDICTIVE_ROUTING_VIEW"

	// ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGQUEUEOVERVIEW captures enum value "PREDICTIVE_ROUTING_QUEUE_OVERVIEW"
	ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGQUEUEOVERVIEW string = "PREDICTIVE_ROUTING_QUEUE_OVERVIEW"

	// ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGMODELVIEW captures enum value "PREDICTIVE_ROUTING_MODEL_VIEW"
	ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGMODELVIEW string = "PREDICTIVE_ROUTING_MODEL_VIEW"

	// ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGIMPACTVIEW captures enum value "PREDICTIVE_ROUTING_IMPACT_VIEW"
	ReportingExportMetadataJobResponseViewTypePREDICTIVEROUTINGIMPACTVIEW string = "PREDICTIVE_ROUTING_IMPACT_VIEW"

	// ReportingExportMetadataJobResponseViewTypeDATAACTIONSPERFORMANCESUMMARYVIEW captures enum value "DATA_ACTIONS_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeDATAACTIONSPERFORMANCESUMMARYVIEW string = "DATA_ACTIONS_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeDATAACTIONSPERFORMANCEDETAILVIEW captures enum value "DATA_ACTIONS_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeDATAACTIONSPERFORMANCEDETAILVIEW string = "DATA_ACTIONS_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTTIMELINESUMMARYVIEW captures enum value "AGENT_TIMELINE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTTIMELINESUMMARYVIEW string = "AGENT_TIMELINE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTTIMELINEDETAILVIEW captures enum value "AGENT_TIMELINE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTTIMELINEDETAILVIEW string = "AGENT_TIMELINE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTLOGINLOGOUTSUMMARYVIEW captures enum value "AGENT_LOGIN_LOGOUT_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTLOGINLOGOUTSUMMARYVIEW string = "AGENT_LOGIN_LOGOUT_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTLOGINLOGOUTDETAILVIEW captures enum value "AGENT_LOGIN_LOGOUT_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTLOGINLOGOUTDETAILVIEW string = "AGENT_LOGIN_LOGOUT_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeCAMPAIGNPERFORMANCESUMMARYVIEW captures enum value "CAMPAIGN_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeCAMPAIGNPERFORMANCESUMMARYVIEW string = "CAMPAIGN_PERFORMANCE_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeCAMPAIGNPERFORMANCEDETAILVIEW captures enum value "CAMPAIGN_PERFORMANCE_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeCAMPAIGNPERFORMANCEDETAILVIEW string = "CAMPAIGN_PERFORMANCE_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeKNOWLEDGEPERFORMANCEVIEW captures enum value "KNOWLEDGE_PERFORMANCE_VIEW"
	ReportingExportMetadataJobResponseViewTypeKNOWLEDGEPERFORMANCEVIEW string = "KNOWLEDGE_PERFORMANCE_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDINSIGHTSSUMMARYVIEW captures enum value "AGENT_SCORECARD_INSIGHTS_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDINSIGHTSSUMMARYVIEW string = "AGENT_SCORECARD_INSIGHTS_SUMMARY_VIEW"

	// ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDINSIGHTSDETAILVIEW captures enum value "AGENT_SCORECARD_INSIGHTS_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeAGENTSCORECARDINSIGHTSDETAILVIEW string = "AGENT_SCORECARD_INSIGHTS_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeQUEUEWRAPUPDETAILVIEW captures enum value "QUEUE_WRAPUP_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeQUEUEWRAPUPDETAILVIEW string = "QUEUE_WRAPUP_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeINTERACTIONDETAILVIEW captures enum value "INTERACTION_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeINTERACTIONDETAILVIEW string = "INTERACTION_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeCAMPAIGNINTERACTIONDETAILVIEW captures enum value "CAMPAIGN_INTERACTION_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeCAMPAIGNINTERACTIONDETAILVIEW string = "CAMPAIGN_INTERACTION_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeCAMPAIGNATTEMPTDETAILVIEW captures enum value "CAMPAIGN_ATTEMPT_DETAIL_VIEW"
	ReportingExportMetadataJobResponseViewTypeCAMPAIGNATTEMPTDETAILVIEW string = "CAMPAIGN_ATTEMPT_DETAIL_VIEW"

	// ReportingExportMetadataJobResponseViewTypeWORKITEMPERFORMANCESUMMARYVIEW captures enum value "WORKITEM_PERFORMANCE_SUMMARY_VIEW"
	ReportingExportMetadataJobResponseViewTypeWORKITEMPERFORMANCESUMMARYVIEW string = "WORKITEM_PERFORMANCE_SUMMARY_VIEW"
)

// prop value enum
func (m *ReportingExportMetadataJobResponse) validateViewTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, reportingExportMetadataJobResponseTypeViewTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ReportingExportMetadataJobResponse) validateViewType(formats strfmt.Registry) error {
	if swag.IsZero(m.ViewType) { // not required
		return nil
	}

	// value enum
	if err := m.validateViewTypeEnum("viewType", "body", m.ViewType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this reporting export metadata job response based on the context it is used
func (m *ReportingExportMetadataJobResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportingExportMetadataJobResponse) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ReportingExportMetadataJobResponse) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportingExportMetadataJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportingExportMetadataJobResponse) UnmarshalBinary(b []byte) error {
	var res ReportingExportMetadataJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
