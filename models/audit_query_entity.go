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

// AuditQueryEntity audit query entity
//
// swagger:model AuditQueryEntity
type AuditQueryEntity struct {

	// List of Actions
	Actions []string `json:"actions"`

	// Name of the Entity
	// Enum: [AccessToken ActionMap ActionTemplate Annotation Appointment AttemptLimits AuthOrganization AuthUser Bulk BulkActions Calibration CallableTimeSet CallAnalysisResponseSet Campaign CampaignRule CampaignSchedule ChangeRequest ClickstreamSettings Configuration ConfigurationVersion ContactList ContactListFilter ConversationAccount ConversationDefaultSupportedContent ConversationPhoneNumber ConversationRecipient ConversationThreadingWindow DashboardSettings DependencyTrackingBuild Deployment DID DIDPool DNCList Document Edge EdgeGroup EdgeLog EdgeLogZip EdgePcaps EdgePreferences EdgeTraceLevel EmergencyGroup Evaluation EvaluationForm EventType Exports Extension ExtensionPool ExternalMetricsData ExternalMetricsDefinition Feedback Flow FlowMilestone FlowOutcome Forecast HistoricalData InsightSettings Integration IVR Line LineBase MaxOrgRoutingUtilizationCapacity MediaDiagnosticsTraceFile MessagingCampaign Metric NumberPlan OAuthClient OAuthClientAuthorization OrganizationAuthorizationTrust OrganizationAuthorizationUserTrust OrganizationFeature OrganizationIntegrationsAccess OrganizationSettings OrphanedRecording OutboundRoute Outcome Pcaps Phone PhoneBase Policy Predictor Product Program Prompt PromptResource Queue Recording RecordingAnnotation RecordingSettings Response Role Row RoutingTranscriptionSettings RuleSet Schedule ScheduledExports ScheduleGroup Schema ScreenRecording Segment SentimentFeedback Sequence SequenceSchedule SessionType Site SpeechTextAnalyticsSettings Status SupportedContent SupportFile Survey SurveyForm Team Topic TranscriptionSettings Trigger Trunk TrunkBase User UserPresence VoicemailPolicy VoicemailUserPolicy Webhook WorkPlan Workspace WrapupCode WrapUpCodeMapping]
	Name string `json:"name,omitempty"`
}

// Validate validates this audit query entity
func (m *AuditQueryEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var auditQueryEntityActionsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Create","View","Update","Move","Delete","DeleteAll","Download","Upload","MemberAdd","MemberUpdate","MemberRemove","Read","ReadAll","Execute","ApplyProtection","RevokeProtection","UpdateRetention","Abandon","Archive","RestoreRequest","RestoreComplete","Publish","Unpublish","Activate","Checkin","Checkout","Deactivate","Debug","Save","Revert","Transcode","Enable","Disable","Authorize","Deauthorize","Authenticate","ChangePassword","Revoke","Export","Append","Recycle","Open","Approved","Rejected","Rollback","ImplementingChange","ChangeImplemented","ImplementingRollback","RollbackImplemented","Write","Purge","Processed","Replace","UpdateInService","UpdateOutOfService","Cycle","Scale","IpAllowlistClear","AddPairingRole","Add"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditQueryEntityActionsItemsEnum = append(auditQueryEntityActionsItemsEnum, v)
	}
}

func (m *AuditQueryEntity) validateActionsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditQueryEntityActionsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditQueryEntity) validateActions(formats strfmt.Registry) error {

	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	for i := 0; i < len(m.Actions); i++ {

		// value enum
		if err := m.validateActionsItemsEnum("actions"+"."+strconv.Itoa(i), "body", m.Actions[i]); err != nil {
			return err
		}

	}

	return nil
}

var auditQueryEntityTypeNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AccessToken","ActionMap","ActionTemplate","Annotation","Appointment","AttemptLimits","AuthOrganization","AuthUser","Bulk","BulkActions","Calibration","CallableTimeSet","CallAnalysisResponseSet","Campaign","CampaignRule","CampaignSchedule","ChangeRequest","ClickstreamSettings","Configuration","ConfigurationVersion","ContactList","ContactListFilter","ConversationAccount","ConversationDefaultSupportedContent","ConversationPhoneNumber","ConversationRecipient","ConversationThreadingWindow","DashboardSettings","DependencyTrackingBuild","Deployment","DID","DIDPool","DNCList","Document","Edge","EdgeGroup","EdgeLog","EdgeLogZip","EdgePcaps","EdgePreferences","EdgeTraceLevel","EmergencyGroup","Evaluation","EvaluationForm","EventType","Exports","Extension","ExtensionPool","ExternalMetricsData","ExternalMetricsDefinition","Feedback","Flow","FlowMilestone","FlowOutcome","Forecast","HistoricalData","InsightSettings","Integration","IVR","Line","LineBase","MaxOrgRoutingUtilizationCapacity","MediaDiagnosticsTraceFile","MessagingCampaign","Metric","NumberPlan","OAuthClient","OAuthClientAuthorization","OrganizationAuthorizationTrust","OrganizationAuthorizationUserTrust","OrganizationFeature","OrganizationIntegrationsAccess","OrganizationSettings","OrphanedRecording","OutboundRoute","Outcome","Pcaps","Phone","PhoneBase","Policy","Predictor","Product","Program","Prompt","PromptResource","Queue","Recording","RecordingAnnotation","RecordingSettings","Response","Role","Row","RoutingTranscriptionSettings","RuleSet","Schedule","ScheduledExports","ScheduleGroup","Schema","ScreenRecording","Segment","SentimentFeedback","Sequence","SequenceSchedule","SessionType","Site","SpeechTextAnalyticsSettings","Status","SupportedContent","SupportFile","Survey","SurveyForm","Team","Topic","TranscriptionSettings","Trigger","Trunk","TrunkBase","User","UserPresence","VoicemailPolicy","VoicemailUserPolicy","Webhook","WorkPlan","Workspace","WrapupCode","WrapUpCodeMapping"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditQueryEntityTypeNamePropEnum = append(auditQueryEntityTypeNamePropEnum, v)
	}
}

const (

	// AuditQueryEntityNameAccessToken captures enum value "AccessToken"
	AuditQueryEntityNameAccessToken string = "AccessToken"

	// AuditQueryEntityNameActionMap captures enum value "ActionMap"
	AuditQueryEntityNameActionMap string = "ActionMap"

	// AuditQueryEntityNameActionTemplate captures enum value "ActionTemplate"
	AuditQueryEntityNameActionTemplate string = "ActionTemplate"

	// AuditQueryEntityNameAnnotation captures enum value "Annotation"
	AuditQueryEntityNameAnnotation string = "Annotation"

	// AuditQueryEntityNameAppointment captures enum value "Appointment"
	AuditQueryEntityNameAppointment string = "Appointment"

	// AuditQueryEntityNameAttemptLimits captures enum value "AttemptLimits"
	AuditQueryEntityNameAttemptLimits string = "AttemptLimits"

	// AuditQueryEntityNameAuthOrganization captures enum value "AuthOrganization"
	AuditQueryEntityNameAuthOrganization string = "AuthOrganization"

	// AuditQueryEntityNameAuthUser captures enum value "AuthUser"
	AuditQueryEntityNameAuthUser string = "AuthUser"

	// AuditQueryEntityNameBulk captures enum value "Bulk"
	AuditQueryEntityNameBulk string = "Bulk"

	// AuditQueryEntityNameBulkActions captures enum value "BulkActions"
	AuditQueryEntityNameBulkActions string = "BulkActions"

	// AuditQueryEntityNameCalibration captures enum value "Calibration"
	AuditQueryEntityNameCalibration string = "Calibration"

	// AuditQueryEntityNameCallableTimeSet captures enum value "CallableTimeSet"
	AuditQueryEntityNameCallableTimeSet string = "CallableTimeSet"

	// AuditQueryEntityNameCallAnalysisResponseSet captures enum value "CallAnalysisResponseSet"
	AuditQueryEntityNameCallAnalysisResponseSet string = "CallAnalysisResponseSet"

	// AuditQueryEntityNameCampaign captures enum value "Campaign"
	AuditQueryEntityNameCampaign string = "Campaign"

	// AuditQueryEntityNameCampaignRule captures enum value "CampaignRule"
	AuditQueryEntityNameCampaignRule string = "CampaignRule"

	// AuditQueryEntityNameCampaignSchedule captures enum value "CampaignSchedule"
	AuditQueryEntityNameCampaignSchedule string = "CampaignSchedule"

	// AuditQueryEntityNameChangeRequest captures enum value "ChangeRequest"
	AuditQueryEntityNameChangeRequest string = "ChangeRequest"

	// AuditQueryEntityNameClickstreamSettings captures enum value "ClickstreamSettings"
	AuditQueryEntityNameClickstreamSettings string = "ClickstreamSettings"

	// AuditQueryEntityNameConfiguration captures enum value "Configuration"
	AuditQueryEntityNameConfiguration string = "Configuration"

	// AuditQueryEntityNameConfigurationVersion captures enum value "ConfigurationVersion"
	AuditQueryEntityNameConfigurationVersion string = "ConfigurationVersion"

	// AuditQueryEntityNameContactList captures enum value "ContactList"
	AuditQueryEntityNameContactList string = "ContactList"

	// AuditQueryEntityNameContactListFilter captures enum value "ContactListFilter"
	AuditQueryEntityNameContactListFilter string = "ContactListFilter"

	// AuditQueryEntityNameConversationAccount captures enum value "ConversationAccount"
	AuditQueryEntityNameConversationAccount string = "ConversationAccount"

	// AuditQueryEntityNameConversationDefaultSupportedContent captures enum value "ConversationDefaultSupportedContent"
	AuditQueryEntityNameConversationDefaultSupportedContent string = "ConversationDefaultSupportedContent"

	// AuditQueryEntityNameConversationPhoneNumber captures enum value "ConversationPhoneNumber"
	AuditQueryEntityNameConversationPhoneNumber string = "ConversationPhoneNumber"

	// AuditQueryEntityNameConversationRecipient captures enum value "ConversationRecipient"
	AuditQueryEntityNameConversationRecipient string = "ConversationRecipient"

	// AuditQueryEntityNameConversationThreadingWindow captures enum value "ConversationThreadingWindow"
	AuditQueryEntityNameConversationThreadingWindow string = "ConversationThreadingWindow"

	// AuditQueryEntityNameDashboardSettings captures enum value "DashboardSettings"
	AuditQueryEntityNameDashboardSettings string = "DashboardSettings"

	// AuditQueryEntityNameDependencyTrackingBuild captures enum value "DependencyTrackingBuild"
	AuditQueryEntityNameDependencyTrackingBuild string = "DependencyTrackingBuild"

	// AuditQueryEntityNameDeployment captures enum value "Deployment"
	AuditQueryEntityNameDeployment string = "Deployment"

	// AuditQueryEntityNameDID captures enum value "DID"
	AuditQueryEntityNameDID string = "DID"

	// AuditQueryEntityNameDIDPool captures enum value "DIDPool"
	AuditQueryEntityNameDIDPool string = "DIDPool"

	// AuditQueryEntityNameDNCList captures enum value "DNCList"
	AuditQueryEntityNameDNCList string = "DNCList"

	// AuditQueryEntityNameDocument captures enum value "Document"
	AuditQueryEntityNameDocument string = "Document"

	// AuditQueryEntityNameEdge captures enum value "Edge"
	AuditQueryEntityNameEdge string = "Edge"

	// AuditQueryEntityNameEdgeGroup captures enum value "EdgeGroup"
	AuditQueryEntityNameEdgeGroup string = "EdgeGroup"

	// AuditQueryEntityNameEdgeLog captures enum value "EdgeLog"
	AuditQueryEntityNameEdgeLog string = "EdgeLog"

	// AuditQueryEntityNameEdgeLogZip captures enum value "EdgeLogZip"
	AuditQueryEntityNameEdgeLogZip string = "EdgeLogZip"

	// AuditQueryEntityNameEdgePcaps captures enum value "EdgePcaps"
	AuditQueryEntityNameEdgePcaps string = "EdgePcaps"

	// AuditQueryEntityNameEdgePreferences captures enum value "EdgePreferences"
	AuditQueryEntityNameEdgePreferences string = "EdgePreferences"

	// AuditQueryEntityNameEdgeTraceLevel captures enum value "EdgeTraceLevel"
	AuditQueryEntityNameEdgeTraceLevel string = "EdgeTraceLevel"

	// AuditQueryEntityNameEmergencyGroup captures enum value "EmergencyGroup"
	AuditQueryEntityNameEmergencyGroup string = "EmergencyGroup"

	// AuditQueryEntityNameEvaluation captures enum value "Evaluation"
	AuditQueryEntityNameEvaluation string = "Evaluation"

	// AuditQueryEntityNameEvaluationForm captures enum value "EvaluationForm"
	AuditQueryEntityNameEvaluationForm string = "EvaluationForm"

	// AuditQueryEntityNameEventType captures enum value "EventType"
	AuditQueryEntityNameEventType string = "EventType"

	// AuditQueryEntityNameExports captures enum value "Exports"
	AuditQueryEntityNameExports string = "Exports"

	// AuditQueryEntityNameExtension captures enum value "Extension"
	AuditQueryEntityNameExtension string = "Extension"

	// AuditQueryEntityNameExtensionPool captures enum value "ExtensionPool"
	AuditQueryEntityNameExtensionPool string = "ExtensionPool"

	// AuditQueryEntityNameExternalMetricsData captures enum value "ExternalMetricsData"
	AuditQueryEntityNameExternalMetricsData string = "ExternalMetricsData"

	// AuditQueryEntityNameExternalMetricsDefinition captures enum value "ExternalMetricsDefinition"
	AuditQueryEntityNameExternalMetricsDefinition string = "ExternalMetricsDefinition"

	// AuditQueryEntityNameFeedback captures enum value "Feedback"
	AuditQueryEntityNameFeedback string = "Feedback"

	// AuditQueryEntityNameFlow captures enum value "Flow"
	AuditQueryEntityNameFlow string = "Flow"

	// AuditQueryEntityNameFlowMilestone captures enum value "FlowMilestone"
	AuditQueryEntityNameFlowMilestone string = "FlowMilestone"

	// AuditQueryEntityNameFlowOutcome captures enum value "FlowOutcome"
	AuditQueryEntityNameFlowOutcome string = "FlowOutcome"

	// AuditQueryEntityNameForecast captures enum value "Forecast"
	AuditQueryEntityNameForecast string = "Forecast"

	// AuditQueryEntityNameHistoricalData captures enum value "HistoricalData"
	AuditQueryEntityNameHistoricalData string = "HistoricalData"

	// AuditQueryEntityNameInsightSettings captures enum value "InsightSettings"
	AuditQueryEntityNameInsightSettings string = "InsightSettings"

	// AuditQueryEntityNameIntegration captures enum value "Integration"
	AuditQueryEntityNameIntegration string = "Integration"

	// AuditQueryEntityNameIVR captures enum value "IVR"
	AuditQueryEntityNameIVR string = "IVR"

	// AuditQueryEntityNameLine captures enum value "Line"
	AuditQueryEntityNameLine string = "Line"

	// AuditQueryEntityNameLineBase captures enum value "LineBase"
	AuditQueryEntityNameLineBase string = "LineBase"

	// AuditQueryEntityNameMaxOrgRoutingUtilizationCapacity captures enum value "MaxOrgRoutingUtilizationCapacity"
	AuditQueryEntityNameMaxOrgRoutingUtilizationCapacity string = "MaxOrgRoutingUtilizationCapacity"

	// AuditQueryEntityNameMediaDiagnosticsTraceFile captures enum value "MediaDiagnosticsTraceFile"
	AuditQueryEntityNameMediaDiagnosticsTraceFile string = "MediaDiagnosticsTraceFile"

	// AuditQueryEntityNameMessagingCampaign captures enum value "MessagingCampaign"
	AuditQueryEntityNameMessagingCampaign string = "MessagingCampaign"

	// AuditQueryEntityNameMetric captures enum value "Metric"
	AuditQueryEntityNameMetric string = "Metric"

	// AuditQueryEntityNameNumberPlan captures enum value "NumberPlan"
	AuditQueryEntityNameNumberPlan string = "NumberPlan"

	// AuditQueryEntityNameOAuthClient captures enum value "OAuthClient"
	AuditQueryEntityNameOAuthClient string = "OAuthClient"

	// AuditQueryEntityNameOAuthClientAuthorization captures enum value "OAuthClientAuthorization"
	AuditQueryEntityNameOAuthClientAuthorization string = "OAuthClientAuthorization"

	// AuditQueryEntityNameOrganizationAuthorizationTrust captures enum value "OrganizationAuthorizationTrust"
	AuditQueryEntityNameOrganizationAuthorizationTrust string = "OrganizationAuthorizationTrust"

	// AuditQueryEntityNameOrganizationAuthorizationUserTrust captures enum value "OrganizationAuthorizationUserTrust"
	AuditQueryEntityNameOrganizationAuthorizationUserTrust string = "OrganizationAuthorizationUserTrust"

	// AuditQueryEntityNameOrganizationFeature captures enum value "OrganizationFeature"
	AuditQueryEntityNameOrganizationFeature string = "OrganizationFeature"

	// AuditQueryEntityNameOrganizationIntegrationsAccess captures enum value "OrganizationIntegrationsAccess"
	AuditQueryEntityNameOrganizationIntegrationsAccess string = "OrganizationIntegrationsAccess"

	// AuditQueryEntityNameOrganizationSettings captures enum value "OrganizationSettings"
	AuditQueryEntityNameOrganizationSettings string = "OrganizationSettings"

	// AuditQueryEntityNameOrphanedRecording captures enum value "OrphanedRecording"
	AuditQueryEntityNameOrphanedRecording string = "OrphanedRecording"

	// AuditQueryEntityNameOutboundRoute captures enum value "OutboundRoute"
	AuditQueryEntityNameOutboundRoute string = "OutboundRoute"

	// AuditQueryEntityNameOutcome captures enum value "Outcome"
	AuditQueryEntityNameOutcome string = "Outcome"

	// AuditQueryEntityNamePcaps captures enum value "Pcaps"
	AuditQueryEntityNamePcaps string = "Pcaps"

	// AuditQueryEntityNamePhone captures enum value "Phone"
	AuditQueryEntityNamePhone string = "Phone"

	// AuditQueryEntityNamePhoneBase captures enum value "PhoneBase"
	AuditQueryEntityNamePhoneBase string = "PhoneBase"

	// AuditQueryEntityNamePolicy captures enum value "Policy"
	AuditQueryEntityNamePolicy string = "Policy"

	// AuditQueryEntityNamePredictor captures enum value "Predictor"
	AuditQueryEntityNamePredictor string = "Predictor"

	// AuditQueryEntityNameProduct captures enum value "Product"
	AuditQueryEntityNameProduct string = "Product"

	// AuditQueryEntityNameProgram captures enum value "Program"
	AuditQueryEntityNameProgram string = "Program"

	// AuditQueryEntityNamePrompt captures enum value "Prompt"
	AuditQueryEntityNamePrompt string = "Prompt"

	// AuditQueryEntityNamePromptResource captures enum value "PromptResource"
	AuditQueryEntityNamePromptResource string = "PromptResource"

	// AuditQueryEntityNameQueue captures enum value "Queue"
	AuditQueryEntityNameQueue string = "Queue"

	// AuditQueryEntityNameRecording captures enum value "Recording"
	AuditQueryEntityNameRecording string = "Recording"

	// AuditQueryEntityNameRecordingAnnotation captures enum value "RecordingAnnotation"
	AuditQueryEntityNameRecordingAnnotation string = "RecordingAnnotation"

	// AuditQueryEntityNameRecordingSettings captures enum value "RecordingSettings"
	AuditQueryEntityNameRecordingSettings string = "RecordingSettings"

	// AuditQueryEntityNameResponse captures enum value "Response"
	AuditQueryEntityNameResponse string = "Response"

	// AuditQueryEntityNameRole captures enum value "Role"
	AuditQueryEntityNameRole string = "Role"

	// AuditQueryEntityNameRow captures enum value "Row"
	AuditQueryEntityNameRow string = "Row"

	// AuditQueryEntityNameRoutingTranscriptionSettings captures enum value "RoutingTranscriptionSettings"
	AuditQueryEntityNameRoutingTranscriptionSettings string = "RoutingTranscriptionSettings"

	// AuditQueryEntityNameRuleSet captures enum value "RuleSet"
	AuditQueryEntityNameRuleSet string = "RuleSet"

	// AuditQueryEntityNameSchedule captures enum value "Schedule"
	AuditQueryEntityNameSchedule string = "Schedule"

	// AuditQueryEntityNameScheduledExports captures enum value "ScheduledExports"
	AuditQueryEntityNameScheduledExports string = "ScheduledExports"

	// AuditQueryEntityNameScheduleGroup captures enum value "ScheduleGroup"
	AuditQueryEntityNameScheduleGroup string = "ScheduleGroup"

	// AuditQueryEntityNameSchema captures enum value "Schema"
	AuditQueryEntityNameSchema string = "Schema"

	// AuditQueryEntityNameScreenRecording captures enum value "ScreenRecording"
	AuditQueryEntityNameScreenRecording string = "ScreenRecording"

	// AuditQueryEntityNameSegment captures enum value "Segment"
	AuditQueryEntityNameSegment string = "Segment"

	// AuditQueryEntityNameSentimentFeedback captures enum value "SentimentFeedback"
	AuditQueryEntityNameSentimentFeedback string = "SentimentFeedback"

	// AuditQueryEntityNameSequence captures enum value "Sequence"
	AuditQueryEntityNameSequence string = "Sequence"

	// AuditQueryEntityNameSequenceSchedule captures enum value "SequenceSchedule"
	AuditQueryEntityNameSequenceSchedule string = "SequenceSchedule"

	// AuditQueryEntityNameSessionType captures enum value "SessionType"
	AuditQueryEntityNameSessionType string = "SessionType"

	// AuditQueryEntityNameSite captures enum value "Site"
	AuditQueryEntityNameSite string = "Site"

	// AuditQueryEntityNameSpeechTextAnalyticsSettings captures enum value "SpeechTextAnalyticsSettings"
	AuditQueryEntityNameSpeechTextAnalyticsSettings string = "SpeechTextAnalyticsSettings"

	// AuditQueryEntityNameStatus captures enum value "Status"
	AuditQueryEntityNameStatus string = "Status"

	// AuditQueryEntityNameSupportedContent captures enum value "SupportedContent"
	AuditQueryEntityNameSupportedContent string = "SupportedContent"

	// AuditQueryEntityNameSupportFile captures enum value "SupportFile"
	AuditQueryEntityNameSupportFile string = "SupportFile"

	// AuditQueryEntityNameSurvey captures enum value "Survey"
	AuditQueryEntityNameSurvey string = "Survey"

	// AuditQueryEntityNameSurveyForm captures enum value "SurveyForm"
	AuditQueryEntityNameSurveyForm string = "SurveyForm"

	// AuditQueryEntityNameTeam captures enum value "Team"
	AuditQueryEntityNameTeam string = "Team"

	// AuditQueryEntityNameTopic captures enum value "Topic"
	AuditQueryEntityNameTopic string = "Topic"

	// AuditQueryEntityNameTranscriptionSettings captures enum value "TranscriptionSettings"
	AuditQueryEntityNameTranscriptionSettings string = "TranscriptionSettings"

	// AuditQueryEntityNameTrigger captures enum value "Trigger"
	AuditQueryEntityNameTrigger string = "Trigger"

	// AuditQueryEntityNameTrunk captures enum value "Trunk"
	AuditQueryEntityNameTrunk string = "Trunk"

	// AuditQueryEntityNameTrunkBase captures enum value "TrunkBase"
	AuditQueryEntityNameTrunkBase string = "TrunkBase"

	// AuditQueryEntityNameUser captures enum value "User"
	AuditQueryEntityNameUser string = "User"

	// AuditQueryEntityNameUserPresence captures enum value "UserPresence"
	AuditQueryEntityNameUserPresence string = "UserPresence"

	// AuditQueryEntityNameVoicemailPolicy captures enum value "VoicemailPolicy"
	AuditQueryEntityNameVoicemailPolicy string = "VoicemailPolicy"

	// AuditQueryEntityNameVoicemailUserPolicy captures enum value "VoicemailUserPolicy"
	AuditQueryEntityNameVoicemailUserPolicy string = "VoicemailUserPolicy"

	// AuditQueryEntityNameWebhook captures enum value "Webhook"
	AuditQueryEntityNameWebhook string = "Webhook"

	// AuditQueryEntityNameWorkPlan captures enum value "WorkPlan"
	AuditQueryEntityNameWorkPlan string = "WorkPlan"

	// AuditQueryEntityNameWorkspace captures enum value "Workspace"
	AuditQueryEntityNameWorkspace string = "Workspace"

	// AuditQueryEntityNameWrapupCode captures enum value "WrapupCode"
	AuditQueryEntityNameWrapupCode string = "WrapupCode"

	// AuditQueryEntityNameWrapUpCodeMapping captures enum value "WrapUpCodeMapping"
	AuditQueryEntityNameWrapUpCodeMapping string = "WrapUpCodeMapping"
)

// prop value enum
func (m *AuditQueryEntity) validateNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditQueryEntityTypeNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditQueryEntity) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameEnum("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditQueryEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditQueryEntity) UnmarshalBinary(b []byte) error {
	var res AuditQueryEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
