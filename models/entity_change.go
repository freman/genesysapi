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

// EntityChange entity change
//
// swagger:model EntityChange
type EntityChange struct {

	// Id of the entity that was changed
	EntityID string `json:"entityId,omitempty"`

	// Name of the entity that was changed
	EntityName string `json:"entityName,omitempty"`

	// Type of the entity that was changed
	// Enum: [AccessToken Action ActionDraft ActionMap ActionTemplate ActivityCode AdherenceExplanation AgentRoutingInfo AnalyticsReportingSettings Annotation Appointment Assignment AttemptLimits AuthOrganization AuthUser Bulk BulkActions BusinessUnit Calibration CallableTimeSet CallAnalysisResponseSet Campaign CampaignRule CampaignSchedule ChangeRequest ClickstreamSettings ComposerPage ComposerScript ComposerPublishedScript ComposerTemplate Configuration ConfigurationVersion ContactList ContactListFilter ContactSchema ConversationAttributes ConversationAccount ConversationDefaultSupportedContent ConversationPhoneNumber ConversationRecipient ConversationThreadingWindow Credential DashboardSettings DefaultPanelSettings DependencyTrackingBuild Deployment DID DIDPool DigitalRuleSet DirectoryGroup DNCList Document DynamicGroup DynamicSchema Edge EdgeGroup EdgeLog EdgeLogZip EdgePcaps EdgePreferences EdgeTraceLevel EmailCampaignSchedule EmergencyGroup EnterpriseAgreement Evaluation EvaluationForm EventType Exports Extension ExtensionPool ExternalMetricsData ExternalMetricsDefinition ExternalOrganizationSchema Feedback Flow FlowMilestone FlowOutcome Forecast GdprRequest Grammar GrammarLanguage Group HistoricalData InboundDomain InboundRoute InsightSettings Integration IntentMiner IVR KnowledgeBase KnowledgeCategory KnowledgeDocument KnowledgeDocumentVariation KnowledgeLabel KnowledgeSearchFeedback KnowledgeTraining Line LineBase Location ManagementUnit MaxOrgRoutingUtilizationCapacity MediaDiagnosticsTraceFile MessagingCampaign MessagingCampaignSchedule Metric Module NumberOrder NumberPlan OAuthClient OAuthClientAuthorization Organization OrganizationAuthorizationTrust OrganizationAuthorizationUserTrust OrganizationFeature OrganizationIntegrationsAccess OrganizationLimits OrganizationSettings OrphanedRecording OutboundDomain OutboundRoute Outcome Pcaps Phone PhoneBase PlanningGroup Policy Predictor Product Profile ProfileMembers Program Prompt PromptResource Public Queue Recording RecordingAnnotation RecordingKey RecordingKeyConfig RecordingSettings Response ResponseAsset Role RoleSettings Row RoutingTranscriptionSettings RoutingUtilizationTag Rule RuleSet Schedule ScheduledExports ScheduleGroup Schema ScreenRecording Segment SentimentFeedback Sequence SequenceSchedule ServiceGoalTemplate SessionType ShiftTrade Site SkillsGroup SkillGroupDefinition SpeechTextAnalyticsSettings StaffingGroup Status SupportedContent SupportFile Survey SurveyForm Tag Team TimeOffLimit TimeOffPlan TimeOffRequest Topic TopicMiner TranscriptionSettings Trigger Trunk TrunkBase User UserLanguage UserPresence UserSkill VoicemailPolicy VoicemailUserPolicy Webhook Workbin Workitem WorkPlan WorkPlanRotation Workspace Worktype WrapupCode WrapUpCodeMapping Participant]
	EntityType string `json:"entityType,omitempty"`

	// New values for the entity.
	NewValues []string `json:"newValues"`

	// Previous values for the entity.
	OldValues []string `json:"oldValues"`
}

// Validate validates this entity change
func (m *EntityChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var entityChangeTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AccessToken","Action","ActionDraft","ActionMap","ActionTemplate","ActivityCode","AdherenceExplanation","AgentRoutingInfo","AnalyticsReportingSettings","Annotation","Appointment","Assignment","AttemptLimits","AuthOrganization","AuthUser","Bulk","BulkActions","BusinessUnit","Calibration","CallableTimeSet","CallAnalysisResponseSet","Campaign","CampaignRule","CampaignSchedule","ChangeRequest","ClickstreamSettings","ComposerPage","ComposerScript","ComposerPublishedScript","ComposerTemplate","Configuration","ConfigurationVersion","ContactList","ContactListFilter","ContactSchema","ConversationAttributes","ConversationAccount","ConversationDefaultSupportedContent","ConversationPhoneNumber","ConversationRecipient","ConversationThreadingWindow","Credential","DashboardSettings","DefaultPanelSettings","DependencyTrackingBuild","Deployment","DID","DIDPool","DigitalRuleSet","DirectoryGroup","DNCList","Document","DynamicGroup","DynamicSchema","Edge","EdgeGroup","EdgeLog","EdgeLogZip","EdgePcaps","EdgePreferences","EdgeTraceLevel","EmailCampaignSchedule","EmergencyGroup","EnterpriseAgreement","Evaluation","EvaluationForm","EventType","Exports","Extension","ExtensionPool","ExternalMetricsData","ExternalMetricsDefinition","ExternalOrganizationSchema","Feedback","Flow","FlowMilestone","FlowOutcome","Forecast","GdprRequest","Grammar","GrammarLanguage","Group","HistoricalData","InboundDomain","InboundRoute","InsightSettings","Integration","IntentMiner","IVR","KnowledgeBase","KnowledgeCategory","KnowledgeDocument","KnowledgeDocumentVariation","KnowledgeLabel","KnowledgeSearchFeedback","KnowledgeTraining","Line","LineBase","Location","ManagementUnit","MaxOrgRoutingUtilizationCapacity","MediaDiagnosticsTraceFile","MessagingCampaign","MessagingCampaignSchedule","Metric","Module","NumberOrder","NumberPlan","OAuthClient","OAuthClientAuthorization","Organization","OrganizationAuthorizationTrust","OrganizationAuthorizationUserTrust","OrganizationFeature","OrganizationIntegrationsAccess","OrganizationLimits","OrganizationSettings","OrphanedRecording","OutboundDomain","OutboundRoute","Outcome","Pcaps","Phone","PhoneBase","PlanningGroup","Policy","Predictor","Product","Profile","ProfileMembers","Program","Prompt","PromptResource","Public","Queue","Recording","RecordingAnnotation","RecordingKey","RecordingKeyConfig","RecordingSettings","Response","ResponseAsset","Role","RoleSettings","Row","RoutingTranscriptionSettings","RoutingUtilizationTag","Rule","RuleSet","Schedule","ScheduledExports","ScheduleGroup","Schema","ScreenRecording","Segment","SentimentFeedback","Sequence","SequenceSchedule","ServiceGoalTemplate","SessionType","ShiftTrade","Site","SkillsGroup","SkillGroupDefinition","SpeechTextAnalyticsSettings","StaffingGroup","Status","SupportedContent","SupportFile","Survey","SurveyForm","Tag","Team","TimeOffLimit","TimeOffPlan","TimeOffRequest","Topic","TopicMiner","TranscriptionSettings","Trigger","Trunk","TrunkBase","User","UserLanguage","UserPresence","UserSkill","VoicemailPolicy","VoicemailUserPolicy","Webhook","Workbin","Workitem","WorkPlan","WorkPlanRotation","Workspace","Worktype","WrapupCode","WrapUpCodeMapping","Participant"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		entityChangeTypeEntityTypePropEnum = append(entityChangeTypeEntityTypePropEnum, v)
	}
}

const (

	// EntityChangeEntityTypeAccessToken captures enum value "AccessToken"
	EntityChangeEntityTypeAccessToken string = "AccessToken"

	// EntityChangeEntityTypeAction captures enum value "Action"
	EntityChangeEntityTypeAction string = "Action"

	// EntityChangeEntityTypeActionDraft captures enum value "ActionDraft"
	EntityChangeEntityTypeActionDraft string = "ActionDraft"

	// EntityChangeEntityTypeActionMap captures enum value "ActionMap"
	EntityChangeEntityTypeActionMap string = "ActionMap"

	// EntityChangeEntityTypeActionTemplate captures enum value "ActionTemplate"
	EntityChangeEntityTypeActionTemplate string = "ActionTemplate"

	// EntityChangeEntityTypeActivityCode captures enum value "ActivityCode"
	EntityChangeEntityTypeActivityCode string = "ActivityCode"

	// EntityChangeEntityTypeAdherenceExplanation captures enum value "AdherenceExplanation"
	EntityChangeEntityTypeAdherenceExplanation string = "AdherenceExplanation"

	// EntityChangeEntityTypeAgentRoutingInfo captures enum value "AgentRoutingInfo"
	EntityChangeEntityTypeAgentRoutingInfo string = "AgentRoutingInfo"

	// EntityChangeEntityTypeAnalyticsReportingSettings captures enum value "AnalyticsReportingSettings"
	EntityChangeEntityTypeAnalyticsReportingSettings string = "AnalyticsReportingSettings"

	// EntityChangeEntityTypeAnnotation captures enum value "Annotation"
	EntityChangeEntityTypeAnnotation string = "Annotation"

	// EntityChangeEntityTypeAppointment captures enum value "Appointment"
	EntityChangeEntityTypeAppointment string = "Appointment"

	// EntityChangeEntityTypeAssignment captures enum value "Assignment"
	EntityChangeEntityTypeAssignment string = "Assignment"

	// EntityChangeEntityTypeAttemptLimits captures enum value "AttemptLimits"
	EntityChangeEntityTypeAttemptLimits string = "AttemptLimits"

	// EntityChangeEntityTypeAuthOrganization captures enum value "AuthOrganization"
	EntityChangeEntityTypeAuthOrganization string = "AuthOrganization"

	// EntityChangeEntityTypeAuthUser captures enum value "AuthUser"
	EntityChangeEntityTypeAuthUser string = "AuthUser"

	// EntityChangeEntityTypeBulk captures enum value "Bulk"
	EntityChangeEntityTypeBulk string = "Bulk"

	// EntityChangeEntityTypeBulkActions captures enum value "BulkActions"
	EntityChangeEntityTypeBulkActions string = "BulkActions"

	// EntityChangeEntityTypeBusinessUnit captures enum value "BusinessUnit"
	EntityChangeEntityTypeBusinessUnit string = "BusinessUnit"

	// EntityChangeEntityTypeCalibration captures enum value "Calibration"
	EntityChangeEntityTypeCalibration string = "Calibration"

	// EntityChangeEntityTypeCallableTimeSet captures enum value "CallableTimeSet"
	EntityChangeEntityTypeCallableTimeSet string = "CallableTimeSet"

	// EntityChangeEntityTypeCallAnalysisResponseSet captures enum value "CallAnalysisResponseSet"
	EntityChangeEntityTypeCallAnalysisResponseSet string = "CallAnalysisResponseSet"

	// EntityChangeEntityTypeCampaign captures enum value "Campaign"
	EntityChangeEntityTypeCampaign string = "Campaign"

	// EntityChangeEntityTypeCampaignRule captures enum value "CampaignRule"
	EntityChangeEntityTypeCampaignRule string = "CampaignRule"

	// EntityChangeEntityTypeCampaignSchedule captures enum value "CampaignSchedule"
	EntityChangeEntityTypeCampaignSchedule string = "CampaignSchedule"

	// EntityChangeEntityTypeChangeRequest captures enum value "ChangeRequest"
	EntityChangeEntityTypeChangeRequest string = "ChangeRequest"

	// EntityChangeEntityTypeClickstreamSettings captures enum value "ClickstreamSettings"
	EntityChangeEntityTypeClickstreamSettings string = "ClickstreamSettings"

	// EntityChangeEntityTypeComposerPage captures enum value "ComposerPage"
	EntityChangeEntityTypeComposerPage string = "ComposerPage"

	// EntityChangeEntityTypeComposerScript captures enum value "ComposerScript"
	EntityChangeEntityTypeComposerScript string = "ComposerScript"

	// EntityChangeEntityTypeComposerPublishedScript captures enum value "ComposerPublishedScript"
	EntityChangeEntityTypeComposerPublishedScript string = "ComposerPublishedScript"

	// EntityChangeEntityTypeComposerTemplate captures enum value "ComposerTemplate"
	EntityChangeEntityTypeComposerTemplate string = "ComposerTemplate"

	// EntityChangeEntityTypeConfiguration captures enum value "Configuration"
	EntityChangeEntityTypeConfiguration string = "Configuration"

	// EntityChangeEntityTypeConfigurationVersion captures enum value "ConfigurationVersion"
	EntityChangeEntityTypeConfigurationVersion string = "ConfigurationVersion"

	// EntityChangeEntityTypeContactList captures enum value "ContactList"
	EntityChangeEntityTypeContactList string = "ContactList"

	// EntityChangeEntityTypeContactListFilter captures enum value "ContactListFilter"
	EntityChangeEntityTypeContactListFilter string = "ContactListFilter"

	// EntityChangeEntityTypeContactSchema captures enum value "ContactSchema"
	EntityChangeEntityTypeContactSchema string = "ContactSchema"

	// EntityChangeEntityTypeConversationAttributes captures enum value "ConversationAttributes"
	EntityChangeEntityTypeConversationAttributes string = "ConversationAttributes"

	// EntityChangeEntityTypeConversationAccount captures enum value "ConversationAccount"
	EntityChangeEntityTypeConversationAccount string = "ConversationAccount"

	// EntityChangeEntityTypeConversationDefaultSupportedContent captures enum value "ConversationDefaultSupportedContent"
	EntityChangeEntityTypeConversationDefaultSupportedContent string = "ConversationDefaultSupportedContent"

	// EntityChangeEntityTypeConversationPhoneNumber captures enum value "ConversationPhoneNumber"
	EntityChangeEntityTypeConversationPhoneNumber string = "ConversationPhoneNumber"

	// EntityChangeEntityTypeConversationRecipient captures enum value "ConversationRecipient"
	EntityChangeEntityTypeConversationRecipient string = "ConversationRecipient"

	// EntityChangeEntityTypeConversationThreadingWindow captures enum value "ConversationThreadingWindow"
	EntityChangeEntityTypeConversationThreadingWindow string = "ConversationThreadingWindow"

	// EntityChangeEntityTypeCredential captures enum value "Credential"
	EntityChangeEntityTypeCredential string = "Credential"

	// EntityChangeEntityTypeDashboardSettings captures enum value "DashboardSettings"
	EntityChangeEntityTypeDashboardSettings string = "DashboardSettings"

	// EntityChangeEntityTypeDefaultPanelSettings captures enum value "DefaultPanelSettings"
	EntityChangeEntityTypeDefaultPanelSettings string = "DefaultPanelSettings"

	// EntityChangeEntityTypeDependencyTrackingBuild captures enum value "DependencyTrackingBuild"
	EntityChangeEntityTypeDependencyTrackingBuild string = "DependencyTrackingBuild"

	// EntityChangeEntityTypeDeployment captures enum value "Deployment"
	EntityChangeEntityTypeDeployment string = "Deployment"

	// EntityChangeEntityTypeDID captures enum value "DID"
	EntityChangeEntityTypeDID string = "DID"

	// EntityChangeEntityTypeDIDPool captures enum value "DIDPool"
	EntityChangeEntityTypeDIDPool string = "DIDPool"

	// EntityChangeEntityTypeDigitalRuleSet captures enum value "DigitalRuleSet"
	EntityChangeEntityTypeDigitalRuleSet string = "DigitalRuleSet"

	// EntityChangeEntityTypeDirectoryGroup captures enum value "DirectoryGroup"
	EntityChangeEntityTypeDirectoryGroup string = "DirectoryGroup"

	// EntityChangeEntityTypeDNCList captures enum value "DNCList"
	EntityChangeEntityTypeDNCList string = "DNCList"

	// EntityChangeEntityTypeDocument captures enum value "Document"
	EntityChangeEntityTypeDocument string = "Document"

	// EntityChangeEntityTypeDynamicGroup captures enum value "DynamicGroup"
	EntityChangeEntityTypeDynamicGroup string = "DynamicGroup"

	// EntityChangeEntityTypeDynamicSchema captures enum value "DynamicSchema"
	EntityChangeEntityTypeDynamicSchema string = "DynamicSchema"

	// EntityChangeEntityTypeEdge captures enum value "Edge"
	EntityChangeEntityTypeEdge string = "Edge"

	// EntityChangeEntityTypeEdgeGroup captures enum value "EdgeGroup"
	EntityChangeEntityTypeEdgeGroup string = "EdgeGroup"

	// EntityChangeEntityTypeEdgeLog captures enum value "EdgeLog"
	EntityChangeEntityTypeEdgeLog string = "EdgeLog"

	// EntityChangeEntityTypeEdgeLogZip captures enum value "EdgeLogZip"
	EntityChangeEntityTypeEdgeLogZip string = "EdgeLogZip"

	// EntityChangeEntityTypeEdgePcaps captures enum value "EdgePcaps"
	EntityChangeEntityTypeEdgePcaps string = "EdgePcaps"

	// EntityChangeEntityTypeEdgePreferences captures enum value "EdgePreferences"
	EntityChangeEntityTypeEdgePreferences string = "EdgePreferences"

	// EntityChangeEntityTypeEdgeTraceLevel captures enum value "EdgeTraceLevel"
	EntityChangeEntityTypeEdgeTraceLevel string = "EdgeTraceLevel"

	// EntityChangeEntityTypeEmailCampaignSchedule captures enum value "EmailCampaignSchedule"
	EntityChangeEntityTypeEmailCampaignSchedule string = "EmailCampaignSchedule"

	// EntityChangeEntityTypeEmergencyGroup captures enum value "EmergencyGroup"
	EntityChangeEntityTypeEmergencyGroup string = "EmergencyGroup"

	// EntityChangeEntityTypeEnterpriseAgreement captures enum value "EnterpriseAgreement"
	EntityChangeEntityTypeEnterpriseAgreement string = "EnterpriseAgreement"

	// EntityChangeEntityTypeEvaluation captures enum value "Evaluation"
	EntityChangeEntityTypeEvaluation string = "Evaluation"

	// EntityChangeEntityTypeEvaluationForm captures enum value "EvaluationForm"
	EntityChangeEntityTypeEvaluationForm string = "EvaluationForm"

	// EntityChangeEntityTypeEventType captures enum value "EventType"
	EntityChangeEntityTypeEventType string = "EventType"

	// EntityChangeEntityTypeExports captures enum value "Exports"
	EntityChangeEntityTypeExports string = "Exports"

	// EntityChangeEntityTypeExtension captures enum value "Extension"
	EntityChangeEntityTypeExtension string = "Extension"

	// EntityChangeEntityTypeExtensionPool captures enum value "ExtensionPool"
	EntityChangeEntityTypeExtensionPool string = "ExtensionPool"

	// EntityChangeEntityTypeExternalMetricsData captures enum value "ExternalMetricsData"
	EntityChangeEntityTypeExternalMetricsData string = "ExternalMetricsData"

	// EntityChangeEntityTypeExternalMetricsDefinition captures enum value "ExternalMetricsDefinition"
	EntityChangeEntityTypeExternalMetricsDefinition string = "ExternalMetricsDefinition"

	// EntityChangeEntityTypeExternalOrganizationSchema captures enum value "ExternalOrganizationSchema"
	EntityChangeEntityTypeExternalOrganizationSchema string = "ExternalOrganizationSchema"

	// EntityChangeEntityTypeFeedback captures enum value "Feedback"
	EntityChangeEntityTypeFeedback string = "Feedback"

	// EntityChangeEntityTypeFlow captures enum value "Flow"
	EntityChangeEntityTypeFlow string = "Flow"

	// EntityChangeEntityTypeFlowMilestone captures enum value "FlowMilestone"
	EntityChangeEntityTypeFlowMilestone string = "FlowMilestone"

	// EntityChangeEntityTypeFlowOutcome captures enum value "FlowOutcome"
	EntityChangeEntityTypeFlowOutcome string = "FlowOutcome"

	// EntityChangeEntityTypeForecast captures enum value "Forecast"
	EntityChangeEntityTypeForecast string = "Forecast"

	// EntityChangeEntityTypeGdprRequest captures enum value "GdprRequest"
	EntityChangeEntityTypeGdprRequest string = "GdprRequest"

	// EntityChangeEntityTypeGrammar captures enum value "Grammar"
	EntityChangeEntityTypeGrammar string = "Grammar"

	// EntityChangeEntityTypeGrammarLanguage captures enum value "GrammarLanguage"
	EntityChangeEntityTypeGrammarLanguage string = "GrammarLanguage"

	// EntityChangeEntityTypeGroup captures enum value "Group"
	EntityChangeEntityTypeGroup string = "Group"

	// EntityChangeEntityTypeHistoricalData captures enum value "HistoricalData"
	EntityChangeEntityTypeHistoricalData string = "HistoricalData"

	// EntityChangeEntityTypeInboundDomain captures enum value "InboundDomain"
	EntityChangeEntityTypeInboundDomain string = "InboundDomain"

	// EntityChangeEntityTypeInboundRoute captures enum value "InboundRoute"
	EntityChangeEntityTypeInboundRoute string = "InboundRoute"

	// EntityChangeEntityTypeInsightSettings captures enum value "InsightSettings"
	EntityChangeEntityTypeInsightSettings string = "InsightSettings"

	// EntityChangeEntityTypeIntegration captures enum value "Integration"
	EntityChangeEntityTypeIntegration string = "Integration"

	// EntityChangeEntityTypeIntentMiner captures enum value "IntentMiner"
	EntityChangeEntityTypeIntentMiner string = "IntentMiner"

	// EntityChangeEntityTypeIVR captures enum value "IVR"
	EntityChangeEntityTypeIVR string = "IVR"

	// EntityChangeEntityTypeKnowledgeBase captures enum value "KnowledgeBase"
	EntityChangeEntityTypeKnowledgeBase string = "KnowledgeBase"

	// EntityChangeEntityTypeKnowledgeCategory captures enum value "KnowledgeCategory"
	EntityChangeEntityTypeKnowledgeCategory string = "KnowledgeCategory"

	// EntityChangeEntityTypeKnowledgeDocument captures enum value "KnowledgeDocument"
	EntityChangeEntityTypeKnowledgeDocument string = "KnowledgeDocument"

	// EntityChangeEntityTypeKnowledgeDocumentVariation captures enum value "KnowledgeDocumentVariation"
	EntityChangeEntityTypeKnowledgeDocumentVariation string = "KnowledgeDocumentVariation"

	// EntityChangeEntityTypeKnowledgeLabel captures enum value "KnowledgeLabel"
	EntityChangeEntityTypeKnowledgeLabel string = "KnowledgeLabel"

	// EntityChangeEntityTypeKnowledgeSearchFeedback captures enum value "KnowledgeSearchFeedback"
	EntityChangeEntityTypeKnowledgeSearchFeedback string = "KnowledgeSearchFeedback"

	// EntityChangeEntityTypeKnowledgeTraining captures enum value "KnowledgeTraining"
	EntityChangeEntityTypeKnowledgeTraining string = "KnowledgeTraining"

	// EntityChangeEntityTypeLine captures enum value "Line"
	EntityChangeEntityTypeLine string = "Line"

	// EntityChangeEntityTypeLineBase captures enum value "LineBase"
	EntityChangeEntityTypeLineBase string = "LineBase"

	// EntityChangeEntityTypeLocation captures enum value "Location"
	EntityChangeEntityTypeLocation string = "Location"

	// EntityChangeEntityTypeManagementUnit captures enum value "ManagementUnit"
	EntityChangeEntityTypeManagementUnit string = "ManagementUnit"

	// EntityChangeEntityTypeMaxOrgRoutingUtilizationCapacity captures enum value "MaxOrgRoutingUtilizationCapacity"
	EntityChangeEntityTypeMaxOrgRoutingUtilizationCapacity string = "MaxOrgRoutingUtilizationCapacity"

	// EntityChangeEntityTypeMediaDiagnosticsTraceFile captures enum value "MediaDiagnosticsTraceFile"
	EntityChangeEntityTypeMediaDiagnosticsTraceFile string = "MediaDiagnosticsTraceFile"

	// EntityChangeEntityTypeMessagingCampaign captures enum value "MessagingCampaign"
	EntityChangeEntityTypeMessagingCampaign string = "MessagingCampaign"

	// EntityChangeEntityTypeMessagingCampaignSchedule captures enum value "MessagingCampaignSchedule"
	EntityChangeEntityTypeMessagingCampaignSchedule string = "MessagingCampaignSchedule"

	// EntityChangeEntityTypeMetric captures enum value "Metric"
	EntityChangeEntityTypeMetric string = "Metric"

	// EntityChangeEntityTypeModule captures enum value "Module"
	EntityChangeEntityTypeModule string = "Module"

	// EntityChangeEntityTypeNumberOrder captures enum value "NumberOrder"
	EntityChangeEntityTypeNumberOrder string = "NumberOrder"

	// EntityChangeEntityTypeNumberPlan captures enum value "NumberPlan"
	EntityChangeEntityTypeNumberPlan string = "NumberPlan"

	// EntityChangeEntityTypeOAuthClient captures enum value "OAuthClient"
	EntityChangeEntityTypeOAuthClient string = "OAuthClient"

	// EntityChangeEntityTypeOAuthClientAuthorization captures enum value "OAuthClientAuthorization"
	EntityChangeEntityTypeOAuthClientAuthorization string = "OAuthClientAuthorization"

	// EntityChangeEntityTypeOrganization captures enum value "Organization"
	EntityChangeEntityTypeOrganization string = "Organization"

	// EntityChangeEntityTypeOrganizationAuthorizationTrust captures enum value "OrganizationAuthorizationTrust"
	EntityChangeEntityTypeOrganizationAuthorizationTrust string = "OrganizationAuthorizationTrust"

	// EntityChangeEntityTypeOrganizationAuthorizationUserTrust captures enum value "OrganizationAuthorizationUserTrust"
	EntityChangeEntityTypeOrganizationAuthorizationUserTrust string = "OrganizationAuthorizationUserTrust"

	// EntityChangeEntityTypeOrganizationFeature captures enum value "OrganizationFeature"
	EntityChangeEntityTypeOrganizationFeature string = "OrganizationFeature"

	// EntityChangeEntityTypeOrganizationIntegrationsAccess captures enum value "OrganizationIntegrationsAccess"
	EntityChangeEntityTypeOrganizationIntegrationsAccess string = "OrganizationIntegrationsAccess"

	// EntityChangeEntityTypeOrganizationLimits captures enum value "OrganizationLimits"
	EntityChangeEntityTypeOrganizationLimits string = "OrganizationLimits"

	// EntityChangeEntityTypeOrganizationSettings captures enum value "OrganizationSettings"
	EntityChangeEntityTypeOrganizationSettings string = "OrganizationSettings"

	// EntityChangeEntityTypeOrphanedRecording captures enum value "OrphanedRecording"
	EntityChangeEntityTypeOrphanedRecording string = "OrphanedRecording"

	// EntityChangeEntityTypeOutboundDomain captures enum value "OutboundDomain"
	EntityChangeEntityTypeOutboundDomain string = "OutboundDomain"

	// EntityChangeEntityTypeOutboundRoute captures enum value "OutboundRoute"
	EntityChangeEntityTypeOutboundRoute string = "OutboundRoute"

	// EntityChangeEntityTypeOutcome captures enum value "Outcome"
	EntityChangeEntityTypeOutcome string = "Outcome"

	// EntityChangeEntityTypePcaps captures enum value "Pcaps"
	EntityChangeEntityTypePcaps string = "Pcaps"

	// EntityChangeEntityTypePhone captures enum value "Phone"
	EntityChangeEntityTypePhone string = "Phone"

	// EntityChangeEntityTypePhoneBase captures enum value "PhoneBase"
	EntityChangeEntityTypePhoneBase string = "PhoneBase"

	// EntityChangeEntityTypePlanningGroup captures enum value "PlanningGroup"
	EntityChangeEntityTypePlanningGroup string = "PlanningGroup"

	// EntityChangeEntityTypePolicy captures enum value "Policy"
	EntityChangeEntityTypePolicy string = "Policy"

	// EntityChangeEntityTypePredictor captures enum value "Predictor"
	EntityChangeEntityTypePredictor string = "Predictor"

	// EntityChangeEntityTypeProduct captures enum value "Product"
	EntityChangeEntityTypeProduct string = "Product"

	// EntityChangeEntityTypeProfile captures enum value "Profile"
	EntityChangeEntityTypeProfile string = "Profile"

	// EntityChangeEntityTypeProfileMembers captures enum value "ProfileMembers"
	EntityChangeEntityTypeProfileMembers string = "ProfileMembers"

	// EntityChangeEntityTypeProgram captures enum value "Program"
	EntityChangeEntityTypeProgram string = "Program"

	// EntityChangeEntityTypePrompt captures enum value "Prompt"
	EntityChangeEntityTypePrompt string = "Prompt"

	// EntityChangeEntityTypePromptResource captures enum value "PromptResource"
	EntityChangeEntityTypePromptResource string = "PromptResource"

	// EntityChangeEntityTypePublic captures enum value "Public"
	EntityChangeEntityTypePublic string = "Public"

	// EntityChangeEntityTypeQueue captures enum value "Queue"
	EntityChangeEntityTypeQueue string = "Queue"

	// EntityChangeEntityTypeRecording captures enum value "Recording"
	EntityChangeEntityTypeRecording string = "Recording"

	// EntityChangeEntityTypeRecordingAnnotation captures enum value "RecordingAnnotation"
	EntityChangeEntityTypeRecordingAnnotation string = "RecordingAnnotation"

	// EntityChangeEntityTypeRecordingKey captures enum value "RecordingKey"
	EntityChangeEntityTypeRecordingKey string = "RecordingKey"

	// EntityChangeEntityTypeRecordingKeyConfig captures enum value "RecordingKeyConfig"
	EntityChangeEntityTypeRecordingKeyConfig string = "RecordingKeyConfig"

	// EntityChangeEntityTypeRecordingSettings captures enum value "RecordingSettings"
	EntityChangeEntityTypeRecordingSettings string = "RecordingSettings"

	// EntityChangeEntityTypeResponse captures enum value "Response"
	EntityChangeEntityTypeResponse string = "Response"

	// EntityChangeEntityTypeResponseAsset captures enum value "ResponseAsset"
	EntityChangeEntityTypeResponseAsset string = "ResponseAsset"

	// EntityChangeEntityTypeRole captures enum value "Role"
	EntityChangeEntityTypeRole string = "Role"

	// EntityChangeEntityTypeRoleSettings captures enum value "RoleSettings"
	EntityChangeEntityTypeRoleSettings string = "RoleSettings"

	// EntityChangeEntityTypeRow captures enum value "Row"
	EntityChangeEntityTypeRow string = "Row"

	// EntityChangeEntityTypeRoutingTranscriptionSettings captures enum value "RoutingTranscriptionSettings"
	EntityChangeEntityTypeRoutingTranscriptionSettings string = "RoutingTranscriptionSettings"

	// EntityChangeEntityTypeRoutingUtilizationTag captures enum value "RoutingUtilizationTag"
	EntityChangeEntityTypeRoutingUtilizationTag string = "RoutingUtilizationTag"

	// EntityChangeEntityTypeRule captures enum value "Rule"
	EntityChangeEntityTypeRule string = "Rule"

	// EntityChangeEntityTypeRuleSet captures enum value "RuleSet"
	EntityChangeEntityTypeRuleSet string = "RuleSet"

	// EntityChangeEntityTypeSchedule captures enum value "Schedule"
	EntityChangeEntityTypeSchedule string = "Schedule"

	// EntityChangeEntityTypeScheduledExports captures enum value "ScheduledExports"
	EntityChangeEntityTypeScheduledExports string = "ScheduledExports"

	// EntityChangeEntityTypeScheduleGroup captures enum value "ScheduleGroup"
	EntityChangeEntityTypeScheduleGroup string = "ScheduleGroup"

	// EntityChangeEntityTypeSchema captures enum value "Schema"
	EntityChangeEntityTypeSchema string = "Schema"

	// EntityChangeEntityTypeScreenRecording captures enum value "ScreenRecording"
	EntityChangeEntityTypeScreenRecording string = "ScreenRecording"

	// EntityChangeEntityTypeSegment captures enum value "Segment"
	EntityChangeEntityTypeSegment string = "Segment"

	// EntityChangeEntityTypeSentimentFeedback captures enum value "SentimentFeedback"
	EntityChangeEntityTypeSentimentFeedback string = "SentimentFeedback"

	// EntityChangeEntityTypeSequence captures enum value "Sequence"
	EntityChangeEntityTypeSequence string = "Sequence"

	// EntityChangeEntityTypeSequenceSchedule captures enum value "SequenceSchedule"
	EntityChangeEntityTypeSequenceSchedule string = "SequenceSchedule"

	// EntityChangeEntityTypeServiceGoalTemplate captures enum value "ServiceGoalTemplate"
	EntityChangeEntityTypeServiceGoalTemplate string = "ServiceGoalTemplate"

	// EntityChangeEntityTypeSessionType captures enum value "SessionType"
	EntityChangeEntityTypeSessionType string = "SessionType"

	// EntityChangeEntityTypeShiftTrade captures enum value "ShiftTrade"
	EntityChangeEntityTypeShiftTrade string = "ShiftTrade"

	// EntityChangeEntityTypeSite captures enum value "Site"
	EntityChangeEntityTypeSite string = "Site"

	// EntityChangeEntityTypeSkillsGroup captures enum value "SkillsGroup"
	EntityChangeEntityTypeSkillsGroup string = "SkillsGroup"

	// EntityChangeEntityTypeSkillGroupDefinition captures enum value "SkillGroupDefinition"
	EntityChangeEntityTypeSkillGroupDefinition string = "SkillGroupDefinition"

	// EntityChangeEntityTypeSpeechTextAnalyticsSettings captures enum value "SpeechTextAnalyticsSettings"
	EntityChangeEntityTypeSpeechTextAnalyticsSettings string = "SpeechTextAnalyticsSettings"

	// EntityChangeEntityTypeStaffingGroup captures enum value "StaffingGroup"
	EntityChangeEntityTypeStaffingGroup string = "StaffingGroup"

	// EntityChangeEntityTypeStatus captures enum value "Status"
	EntityChangeEntityTypeStatus string = "Status"

	// EntityChangeEntityTypeSupportedContent captures enum value "SupportedContent"
	EntityChangeEntityTypeSupportedContent string = "SupportedContent"

	// EntityChangeEntityTypeSupportFile captures enum value "SupportFile"
	EntityChangeEntityTypeSupportFile string = "SupportFile"

	// EntityChangeEntityTypeSurvey captures enum value "Survey"
	EntityChangeEntityTypeSurvey string = "Survey"

	// EntityChangeEntityTypeSurveyForm captures enum value "SurveyForm"
	EntityChangeEntityTypeSurveyForm string = "SurveyForm"

	// EntityChangeEntityTypeTag captures enum value "Tag"
	EntityChangeEntityTypeTag string = "Tag"

	// EntityChangeEntityTypeTeam captures enum value "Team"
	EntityChangeEntityTypeTeam string = "Team"

	// EntityChangeEntityTypeTimeOffLimit captures enum value "TimeOffLimit"
	EntityChangeEntityTypeTimeOffLimit string = "TimeOffLimit"

	// EntityChangeEntityTypeTimeOffPlan captures enum value "TimeOffPlan"
	EntityChangeEntityTypeTimeOffPlan string = "TimeOffPlan"

	// EntityChangeEntityTypeTimeOffRequest captures enum value "TimeOffRequest"
	EntityChangeEntityTypeTimeOffRequest string = "TimeOffRequest"

	// EntityChangeEntityTypeTopic captures enum value "Topic"
	EntityChangeEntityTypeTopic string = "Topic"

	// EntityChangeEntityTypeTopicMiner captures enum value "TopicMiner"
	EntityChangeEntityTypeTopicMiner string = "TopicMiner"

	// EntityChangeEntityTypeTranscriptionSettings captures enum value "TranscriptionSettings"
	EntityChangeEntityTypeTranscriptionSettings string = "TranscriptionSettings"

	// EntityChangeEntityTypeTrigger captures enum value "Trigger"
	EntityChangeEntityTypeTrigger string = "Trigger"

	// EntityChangeEntityTypeTrunk captures enum value "Trunk"
	EntityChangeEntityTypeTrunk string = "Trunk"

	// EntityChangeEntityTypeTrunkBase captures enum value "TrunkBase"
	EntityChangeEntityTypeTrunkBase string = "TrunkBase"

	// EntityChangeEntityTypeUser captures enum value "User"
	EntityChangeEntityTypeUser string = "User"

	// EntityChangeEntityTypeUserLanguage captures enum value "UserLanguage"
	EntityChangeEntityTypeUserLanguage string = "UserLanguage"

	// EntityChangeEntityTypeUserPresence captures enum value "UserPresence"
	EntityChangeEntityTypeUserPresence string = "UserPresence"

	// EntityChangeEntityTypeUserSkill captures enum value "UserSkill"
	EntityChangeEntityTypeUserSkill string = "UserSkill"

	// EntityChangeEntityTypeVoicemailPolicy captures enum value "VoicemailPolicy"
	EntityChangeEntityTypeVoicemailPolicy string = "VoicemailPolicy"

	// EntityChangeEntityTypeVoicemailUserPolicy captures enum value "VoicemailUserPolicy"
	EntityChangeEntityTypeVoicemailUserPolicy string = "VoicemailUserPolicy"

	// EntityChangeEntityTypeWebhook captures enum value "Webhook"
	EntityChangeEntityTypeWebhook string = "Webhook"

	// EntityChangeEntityTypeWorkbin captures enum value "Workbin"
	EntityChangeEntityTypeWorkbin string = "Workbin"

	// EntityChangeEntityTypeWorkitem captures enum value "Workitem"
	EntityChangeEntityTypeWorkitem string = "Workitem"

	// EntityChangeEntityTypeWorkPlan captures enum value "WorkPlan"
	EntityChangeEntityTypeWorkPlan string = "WorkPlan"

	// EntityChangeEntityTypeWorkPlanRotation captures enum value "WorkPlanRotation"
	EntityChangeEntityTypeWorkPlanRotation string = "WorkPlanRotation"

	// EntityChangeEntityTypeWorkspace captures enum value "Workspace"
	EntityChangeEntityTypeWorkspace string = "Workspace"

	// EntityChangeEntityTypeWorktype captures enum value "Worktype"
	EntityChangeEntityTypeWorktype string = "Worktype"

	// EntityChangeEntityTypeWrapupCode captures enum value "WrapupCode"
	EntityChangeEntityTypeWrapupCode string = "WrapupCode"

	// EntityChangeEntityTypeWrapUpCodeMapping captures enum value "WrapUpCodeMapping"
	EntityChangeEntityTypeWrapUpCodeMapping string = "WrapUpCodeMapping"

	// EntityChangeEntityTypeParticipant captures enum value "Participant"
	EntityChangeEntityTypeParticipant string = "Participant"
)

// prop value enum
func (m *EntityChange) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, entityChangeTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EntityChange) validateEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this entity change based on context it is used
func (m *EntityChange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EntityChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EntityChange) UnmarshalBinary(b []byte) error {
	var res EntityChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
