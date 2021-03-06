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

// AuditLogMessage audit log message
//
// swagger:model AuditLogMessage
type AuditLogMessage struct {

	// Action that took place.
	// Enum: [Create View Update Delete Download Upload MemberAdd MemberUpdate MemberRemove Read ApplyProtection RevokeProtection UpdateRetention ReadAll Execute Publish Unpublish Activate Checkin Checkout Deactivate Debug Save Revert Transcode Enable Disable Authorize Deauthorize Authenticate ChangePassword Revoke Export Append Recycle]
	Action string `json:"action,omitempty"`

	// Client associated with this audit message.
	Client *AddressableEntityRef `json:"client,omitempty"`

	// Additional context for this message.
	Context map[string]string `json:"context,omitempty"`

	// Entity that was impacted.
	Entity *DomainEntityRef `json:"entity,omitempty"`

	// Type of the entity that was impacted.
	// Enum: [Document Queue Recording Role VoicemailUserPolicy UserPresence WrapupCode MaxOrgRoutingUtilizationCapacity AccessToken OAuthClient OAuthClientAuthorization AuthOrganization AuthUser OrganizationAuthorizationTrust OrganizationAuthorizationUserTrust BulkActions Feedback Topic Program Segment Outcome SessionType EventType ClickstreamSettings Schedule ScheduleGroup EmergencyGroup IVR Trigger Response DependencyTrackingBuild Flow Prompt PromptResource FlowOutcome FlowMilestone Team Edge EdgeGroup Trunk TrunkBase DID DIDPool Extension ExtensionPool Phone PhoneBase Line LineBase OutboundRoute NumberPlan Site AttemptLimits CallableTimeSet Campaign CampaignRule Sequence ContactList ContactListFilter DNCList CallAnalysisResponseSet RuleSet CampaignSchedule SequenceSchedule OrganizationProperties WrapUpCodeMapping MessagingCampaign TranscriptionSettings SpeechTextAnalyticsSettings Predictor]
	EntityType string `json:"entityType,omitempty"`

	// Date and time of when the audit message was logged. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EventDate strfmt.DateTime `json:"eventDate,omitempty"`

	// Id of the audit message.
	ID string `json:"id,omitempty"`

	// Message describing the event being audited.
	Message *MessageInfo `json:"message,omitempty"`

	// List of properties that were changed and changes made to those properties.
	PropertyChanges []*PropertyChange `json:"propertyChanges"`

	// List of IP addresses of systems that originated or handled the request.
	RemoteIP []string `json:"remoteIp"`

	// Name of the service that logged this audit message.
	// Enum: [Architect ContactCenter ContentManagement PeoplePermissions Presence Quality LanguageUnderstanding TopicsDefinitions PredictiveEngagement WorkforceManagement Triggers ResponseManagement Groups Telephony Outbound SpeechAndTextAnalytics Routing]
	ServiceName string `json:"serviceName,omitempty"`

	// User associated with this audit message.
	User *DomainEntityRef `json:"user,omitempty"`

	// Home Organization Id associated with this audit message.
	UserHomeOrgID string `json:"userHomeOrgId,omitempty"`
}

// Validate validates this audit log message
func (m *AuditLogMessage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePropertyChanges(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var auditLogMessageTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Create","View","Update","Delete","Download","Upload","MemberAdd","MemberUpdate","MemberRemove","Read","ApplyProtection","RevokeProtection","UpdateRetention","ReadAll","Execute","Publish","Unpublish","Activate","Checkin","Checkout","Deactivate","Debug","Save","Revert","Transcode","Enable","Disable","Authorize","Deauthorize","Authenticate","ChangePassword","Revoke","Export","Append","Recycle"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditLogMessageTypeActionPropEnum = append(auditLogMessageTypeActionPropEnum, v)
	}
}

const (

	// AuditLogMessageActionCreate captures enum value "Create"
	AuditLogMessageActionCreate string = "Create"

	// AuditLogMessageActionView captures enum value "View"
	AuditLogMessageActionView string = "View"

	// AuditLogMessageActionUpdate captures enum value "Update"
	AuditLogMessageActionUpdate string = "Update"

	// AuditLogMessageActionDelete captures enum value "Delete"
	AuditLogMessageActionDelete string = "Delete"

	// AuditLogMessageActionDownload captures enum value "Download"
	AuditLogMessageActionDownload string = "Download"

	// AuditLogMessageActionUpload captures enum value "Upload"
	AuditLogMessageActionUpload string = "Upload"

	// AuditLogMessageActionMemberAdd captures enum value "MemberAdd"
	AuditLogMessageActionMemberAdd string = "MemberAdd"

	// AuditLogMessageActionMemberUpdate captures enum value "MemberUpdate"
	AuditLogMessageActionMemberUpdate string = "MemberUpdate"

	// AuditLogMessageActionMemberRemove captures enum value "MemberRemove"
	AuditLogMessageActionMemberRemove string = "MemberRemove"

	// AuditLogMessageActionRead captures enum value "Read"
	AuditLogMessageActionRead string = "Read"

	// AuditLogMessageActionApplyProtection captures enum value "ApplyProtection"
	AuditLogMessageActionApplyProtection string = "ApplyProtection"

	// AuditLogMessageActionRevokeProtection captures enum value "RevokeProtection"
	AuditLogMessageActionRevokeProtection string = "RevokeProtection"

	// AuditLogMessageActionUpdateRetention captures enum value "UpdateRetention"
	AuditLogMessageActionUpdateRetention string = "UpdateRetention"

	// AuditLogMessageActionReadAll captures enum value "ReadAll"
	AuditLogMessageActionReadAll string = "ReadAll"

	// AuditLogMessageActionExecute captures enum value "Execute"
	AuditLogMessageActionExecute string = "Execute"

	// AuditLogMessageActionPublish captures enum value "Publish"
	AuditLogMessageActionPublish string = "Publish"

	// AuditLogMessageActionUnpublish captures enum value "Unpublish"
	AuditLogMessageActionUnpublish string = "Unpublish"

	// AuditLogMessageActionActivate captures enum value "Activate"
	AuditLogMessageActionActivate string = "Activate"

	// AuditLogMessageActionCheckin captures enum value "Checkin"
	AuditLogMessageActionCheckin string = "Checkin"

	// AuditLogMessageActionCheckout captures enum value "Checkout"
	AuditLogMessageActionCheckout string = "Checkout"

	// AuditLogMessageActionDeactivate captures enum value "Deactivate"
	AuditLogMessageActionDeactivate string = "Deactivate"

	// AuditLogMessageActionDebug captures enum value "Debug"
	AuditLogMessageActionDebug string = "Debug"

	// AuditLogMessageActionSave captures enum value "Save"
	AuditLogMessageActionSave string = "Save"

	// AuditLogMessageActionRevert captures enum value "Revert"
	AuditLogMessageActionRevert string = "Revert"

	// AuditLogMessageActionTranscode captures enum value "Transcode"
	AuditLogMessageActionTranscode string = "Transcode"

	// AuditLogMessageActionEnable captures enum value "Enable"
	AuditLogMessageActionEnable string = "Enable"

	// AuditLogMessageActionDisable captures enum value "Disable"
	AuditLogMessageActionDisable string = "Disable"

	// AuditLogMessageActionAuthorize captures enum value "Authorize"
	AuditLogMessageActionAuthorize string = "Authorize"

	// AuditLogMessageActionDeauthorize captures enum value "Deauthorize"
	AuditLogMessageActionDeauthorize string = "Deauthorize"

	// AuditLogMessageActionAuthenticate captures enum value "Authenticate"
	AuditLogMessageActionAuthenticate string = "Authenticate"

	// AuditLogMessageActionChangePassword captures enum value "ChangePassword"
	AuditLogMessageActionChangePassword string = "ChangePassword"

	// AuditLogMessageActionRevoke captures enum value "Revoke"
	AuditLogMessageActionRevoke string = "Revoke"

	// AuditLogMessageActionExport captures enum value "Export"
	AuditLogMessageActionExport string = "Export"

	// AuditLogMessageActionAppend captures enum value "Append"
	AuditLogMessageActionAppend string = "Append"

	// AuditLogMessageActionRecycle captures enum value "Recycle"
	AuditLogMessageActionRecycle string = "Recycle"
)

// prop value enum
func (m *AuditLogMessage) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditLogMessageTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditLogMessage) validateAction(formats strfmt.Registry) error {

	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *AuditLogMessage) validateClient(formats strfmt.Registry) error {

	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

func (m *AuditLogMessage) validateEntity(formats strfmt.Registry) error {

	if swag.IsZero(m.Entity) { // not required
		return nil
	}

	if m.Entity != nil {
		if err := m.Entity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("entity")
			}
			return err
		}
	}

	return nil
}

var auditLogMessageTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Document","Queue","Recording","Role","VoicemailUserPolicy","UserPresence","WrapupCode","MaxOrgRoutingUtilizationCapacity","AccessToken","OAuthClient","OAuthClientAuthorization","AuthOrganization","AuthUser","OrganizationAuthorizationTrust","OrganizationAuthorizationUserTrust","BulkActions","Feedback","Topic","Program","Segment","Outcome","SessionType","EventType","ClickstreamSettings","Schedule","ScheduleGroup","EmergencyGroup","IVR","Trigger","Response","DependencyTrackingBuild","Flow","Prompt","PromptResource","FlowOutcome","FlowMilestone","Team","Edge","EdgeGroup","Trunk","TrunkBase","DID","DIDPool","Extension","ExtensionPool","Phone","PhoneBase","Line","LineBase","OutboundRoute","NumberPlan","Site","AttemptLimits","CallableTimeSet","Campaign","CampaignRule","Sequence","ContactList","ContactListFilter","DNCList","CallAnalysisResponseSet","RuleSet","CampaignSchedule","SequenceSchedule","OrganizationProperties","WrapUpCodeMapping","MessagingCampaign","TranscriptionSettings","SpeechTextAnalyticsSettings","Predictor"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditLogMessageTypeEntityTypePropEnum = append(auditLogMessageTypeEntityTypePropEnum, v)
	}
}

const (

	// AuditLogMessageEntityTypeDocument captures enum value "Document"
	AuditLogMessageEntityTypeDocument string = "Document"

	// AuditLogMessageEntityTypeQueue captures enum value "Queue"
	AuditLogMessageEntityTypeQueue string = "Queue"

	// AuditLogMessageEntityTypeRecording captures enum value "Recording"
	AuditLogMessageEntityTypeRecording string = "Recording"

	// AuditLogMessageEntityTypeRole captures enum value "Role"
	AuditLogMessageEntityTypeRole string = "Role"

	// AuditLogMessageEntityTypeVoicemailUserPolicy captures enum value "VoicemailUserPolicy"
	AuditLogMessageEntityTypeVoicemailUserPolicy string = "VoicemailUserPolicy"

	// AuditLogMessageEntityTypeUserPresence captures enum value "UserPresence"
	AuditLogMessageEntityTypeUserPresence string = "UserPresence"

	// AuditLogMessageEntityTypeWrapupCode captures enum value "WrapupCode"
	AuditLogMessageEntityTypeWrapupCode string = "WrapupCode"

	// AuditLogMessageEntityTypeMaxOrgRoutingUtilizationCapacity captures enum value "MaxOrgRoutingUtilizationCapacity"
	AuditLogMessageEntityTypeMaxOrgRoutingUtilizationCapacity string = "MaxOrgRoutingUtilizationCapacity"

	// AuditLogMessageEntityTypeAccessToken captures enum value "AccessToken"
	AuditLogMessageEntityTypeAccessToken string = "AccessToken"

	// AuditLogMessageEntityTypeOAuthClient captures enum value "OAuthClient"
	AuditLogMessageEntityTypeOAuthClient string = "OAuthClient"

	// AuditLogMessageEntityTypeOAuthClientAuthorization captures enum value "OAuthClientAuthorization"
	AuditLogMessageEntityTypeOAuthClientAuthorization string = "OAuthClientAuthorization"

	// AuditLogMessageEntityTypeAuthOrganization captures enum value "AuthOrganization"
	AuditLogMessageEntityTypeAuthOrganization string = "AuthOrganization"

	// AuditLogMessageEntityTypeAuthUser captures enum value "AuthUser"
	AuditLogMessageEntityTypeAuthUser string = "AuthUser"

	// AuditLogMessageEntityTypeOrganizationAuthorizationTrust captures enum value "OrganizationAuthorizationTrust"
	AuditLogMessageEntityTypeOrganizationAuthorizationTrust string = "OrganizationAuthorizationTrust"

	// AuditLogMessageEntityTypeOrganizationAuthorizationUserTrust captures enum value "OrganizationAuthorizationUserTrust"
	AuditLogMessageEntityTypeOrganizationAuthorizationUserTrust string = "OrganizationAuthorizationUserTrust"

	// AuditLogMessageEntityTypeBulkActions captures enum value "BulkActions"
	AuditLogMessageEntityTypeBulkActions string = "BulkActions"

	// AuditLogMessageEntityTypeFeedback captures enum value "Feedback"
	AuditLogMessageEntityTypeFeedback string = "Feedback"

	// AuditLogMessageEntityTypeTopic captures enum value "Topic"
	AuditLogMessageEntityTypeTopic string = "Topic"

	// AuditLogMessageEntityTypeProgram captures enum value "Program"
	AuditLogMessageEntityTypeProgram string = "Program"

	// AuditLogMessageEntityTypeSegment captures enum value "Segment"
	AuditLogMessageEntityTypeSegment string = "Segment"

	// AuditLogMessageEntityTypeOutcome captures enum value "Outcome"
	AuditLogMessageEntityTypeOutcome string = "Outcome"

	// AuditLogMessageEntityTypeSessionType captures enum value "SessionType"
	AuditLogMessageEntityTypeSessionType string = "SessionType"

	// AuditLogMessageEntityTypeEventType captures enum value "EventType"
	AuditLogMessageEntityTypeEventType string = "EventType"

	// AuditLogMessageEntityTypeClickstreamSettings captures enum value "ClickstreamSettings"
	AuditLogMessageEntityTypeClickstreamSettings string = "ClickstreamSettings"

	// AuditLogMessageEntityTypeSchedule captures enum value "Schedule"
	AuditLogMessageEntityTypeSchedule string = "Schedule"

	// AuditLogMessageEntityTypeScheduleGroup captures enum value "ScheduleGroup"
	AuditLogMessageEntityTypeScheduleGroup string = "ScheduleGroup"

	// AuditLogMessageEntityTypeEmergencyGroup captures enum value "EmergencyGroup"
	AuditLogMessageEntityTypeEmergencyGroup string = "EmergencyGroup"

	// AuditLogMessageEntityTypeIVR captures enum value "IVR"
	AuditLogMessageEntityTypeIVR string = "IVR"

	// AuditLogMessageEntityTypeTrigger captures enum value "Trigger"
	AuditLogMessageEntityTypeTrigger string = "Trigger"

	// AuditLogMessageEntityTypeResponse captures enum value "Response"
	AuditLogMessageEntityTypeResponse string = "Response"

	// AuditLogMessageEntityTypeDependencyTrackingBuild captures enum value "DependencyTrackingBuild"
	AuditLogMessageEntityTypeDependencyTrackingBuild string = "DependencyTrackingBuild"

	// AuditLogMessageEntityTypeFlow captures enum value "Flow"
	AuditLogMessageEntityTypeFlow string = "Flow"

	// AuditLogMessageEntityTypePrompt captures enum value "Prompt"
	AuditLogMessageEntityTypePrompt string = "Prompt"

	// AuditLogMessageEntityTypePromptResource captures enum value "PromptResource"
	AuditLogMessageEntityTypePromptResource string = "PromptResource"

	// AuditLogMessageEntityTypeFlowOutcome captures enum value "FlowOutcome"
	AuditLogMessageEntityTypeFlowOutcome string = "FlowOutcome"

	// AuditLogMessageEntityTypeFlowMilestone captures enum value "FlowMilestone"
	AuditLogMessageEntityTypeFlowMilestone string = "FlowMilestone"

	// AuditLogMessageEntityTypeTeam captures enum value "Team"
	AuditLogMessageEntityTypeTeam string = "Team"

	// AuditLogMessageEntityTypeEdge captures enum value "Edge"
	AuditLogMessageEntityTypeEdge string = "Edge"

	// AuditLogMessageEntityTypeEdgeGroup captures enum value "EdgeGroup"
	AuditLogMessageEntityTypeEdgeGroup string = "EdgeGroup"

	// AuditLogMessageEntityTypeTrunk captures enum value "Trunk"
	AuditLogMessageEntityTypeTrunk string = "Trunk"

	// AuditLogMessageEntityTypeTrunkBase captures enum value "TrunkBase"
	AuditLogMessageEntityTypeTrunkBase string = "TrunkBase"

	// AuditLogMessageEntityTypeDID captures enum value "DID"
	AuditLogMessageEntityTypeDID string = "DID"

	// AuditLogMessageEntityTypeDIDPool captures enum value "DIDPool"
	AuditLogMessageEntityTypeDIDPool string = "DIDPool"

	// AuditLogMessageEntityTypeExtension captures enum value "Extension"
	AuditLogMessageEntityTypeExtension string = "Extension"

	// AuditLogMessageEntityTypeExtensionPool captures enum value "ExtensionPool"
	AuditLogMessageEntityTypeExtensionPool string = "ExtensionPool"

	// AuditLogMessageEntityTypePhone captures enum value "Phone"
	AuditLogMessageEntityTypePhone string = "Phone"

	// AuditLogMessageEntityTypePhoneBase captures enum value "PhoneBase"
	AuditLogMessageEntityTypePhoneBase string = "PhoneBase"

	// AuditLogMessageEntityTypeLine captures enum value "Line"
	AuditLogMessageEntityTypeLine string = "Line"

	// AuditLogMessageEntityTypeLineBase captures enum value "LineBase"
	AuditLogMessageEntityTypeLineBase string = "LineBase"

	// AuditLogMessageEntityTypeOutboundRoute captures enum value "OutboundRoute"
	AuditLogMessageEntityTypeOutboundRoute string = "OutboundRoute"

	// AuditLogMessageEntityTypeNumberPlan captures enum value "NumberPlan"
	AuditLogMessageEntityTypeNumberPlan string = "NumberPlan"

	// AuditLogMessageEntityTypeSite captures enum value "Site"
	AuditLogMessageEntityTypeSite string = "Site"

	// AuditLogMessageEntityTypeAttemptLimits captures enum value "AttemptLimits"
	AuditLogMessageEntityTypeAttemptLimits string = "AttemptLimits"

	// AuditLogMessageEntityTypeCallableTimeSet captures enum value "CallableTimeSet"
	AuditLogMessageEntityTypeCallableTimeSet string = "CallableTimeSet"

	// AuditLogMessageEntityTypeCampaign captures enum value "Campaign"
	AuditLogMessageEntityTypeCampaign string = "Campaign"

	// AuditLogMessageEntityTypeCampaignRule captures enum value "CampaignRule"
	AuditLogMessageEntityTypeCampaignRule string = "CampaignRule"

	// AuditLogMessageEntityTypeSequence captures enum value "Sequence"
	AuditLogMessageEntityTypeSequence string = "Sequence"

	// AuditLogMessageEntityTypeContactList captures enum value "ContactList"
	AuditLogMessageEntityTypeContactList string = "ContactList"

	// AuditLogMessageEntityTypeContactListFilter captures enum value "ContactListFilter"
	AuditLogMessageEntityTypeContactListFilter string = "ContactListFilter"

	// AuditLogMessageEntityTypeDNCList captures enum value "DNCList"
	AuditLogMessageEntityTypeDNCList string = "DNCList"

	// AuditLogMessageEntityTypeCallAnalysisResponseSet captures enum value "CallAnalysisResponseSet"
	AuditLogMessageEntityTypeCallAnalysisResponseSet string = "CallAnalysisResponseSet"

	// AuditLogMessageEntityTypeRuleSet captures enum value "RuleSet"
	AuditLogMessageEntityTypeRuleSet string = "RuleSet"

	// AuditLogMessageEntityTypeCampaignSchedule captures enum value "CampaignSchedule"
	AuditLogMessageEntityTypeCampaignSchedule string = "CampaignSchedule"

	// AuditLogMessageEntityTypeSequenceSchedule captures enum value "SequenceSchedule"
	AuditLogMessageEntityTypeSequenceSchedule string = "SequenceSchedule"

	// AuditLogMessageEntityTypeOrganizationProperties captures enum value "OrganizationProperties"
	AuditLogMessageEntityTypeOrganizationProperties string = "OrganizationProperties"

	// AuditLogMessageEntityTypeWrapUpCodeMapping captures enum value "WrapUpCodeMapping"
	AuditLogMessageEntityTypeWrapUpCodeMapping string = "WrapUpCodeMapping"

	// AuditLogMessageEntityTypeMessagingCampaign captures enum value "MessagingCampaign"
	AuditLogMessageEntityTypeMessagingCampaign string = "MessagingCampaign"

	// AuditLogMessageEntityTypeTranscriptionSettings captures enum value "TranscriptionSettings"
	AuditLogMessageEntityTypeTranscriptionSettings string = "TranscriptionSettings"

	// AuditLogMessageEntityTypeSpeechTextAnalyticsSettings captures enum value "SpeechTextAnalyticsSettings"
	AuditLogMessageEntityTypeSpeechTextAnalyticsSettings string = "SpeechTextAnalyticsSettings"

	// AuditLogMessageEntityTypePredictor captures enum value "Predictor"
	AuditLogMessageEntityTypePredictor string = "Predictor"
)

// prop value enum
func (m *AuditLogMessage) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditLogMessageTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditLogMessage) validateEntityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *AuditLogMessage) validateEventDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EventDate) { // not required
		return nil
	}

	if err := validate.FormatOf("eventDate", "body", "date-time", m.EventDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AuditLogMessage) validateMessage(formats strfmt.Registry) error {

	if swag.IsZero(m.Message) { // not required
		return nil
	}

	if m.Message != nil {
		if err := m.Message.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("message")
			}
			return err
		}
	}

	return nil
}

func (m *AuditLogMessage) validatePropertyChanges(formats strfmt.Registry) error {

	if swag.IsZero(m.PropertyChanges) { // not required
		return nil
	}

	for i := 0; i < len(m.PropertyChanges); i++ {
		if swag.IsZero(m.PropertyChanges[i]) { // not required
			continue
		}

		if m.PropertyChanges[i] != nil {
			if err := m.PropertyChanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("propertyChanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var auditLogMessageTypeServiceNamePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Architect","ContactCenter","ContentManagement","PeoplePermissions","Presence","Quality","LanguageUnderstanding","TopicsDefinitions","PredictiveEngagement","WorkforceManagement","Triggers","ResponseManagement","Groups","Telephony","Outbound","SpeechAndTextAnalytics","Routing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditLogMessageTypeServiceNamePropEnum = append(auditLogMessageTypeServiceNamePropEnum, v)
	}
}

const (

	// AuditLogMessageServiceNameArchitect captures enum value "Architect"
	AuditLogMessageServiceNameArchitect string = "Architect"

	// AuditLogMessageServiceNameContactCenter captures enum value "ContactCenter"
	AuditLogMessageServiceNameContactCenter string = "ContactCenter"

	// AuditLogMessageServiceNameContentManagement captures enum value "ContentManagement"
	AuditLogMessageServiceNameContentManagement string = "ContentManagement"

	// AuditLogMessageServiceNamePeoplePermissions captures enum value "PeoplePermissions"
	AuditLogMessageServiceNamePeoplePermissions string = "PeoplePermissions"

	// AuditLogMessageServiceNamePresence captures enum value "Presence"
	AuditLogMessageServiceNamePresence string = "Presence"

	// AuditLogMessageServiceNameQuality captures enum value "Quality"
	AuditLogMessageServiceNameQuality string = "Quality"

	// AuditLogMessageServiceNameLanguageUnderstanding captures enum value "LanguageUnderstanding"
	AuditLogMessageServiceNameLanguageUnderstanding string = "LanguageUnderstanding"

	// AuditLogMessageServiceNameTopicsDefinitions captures enum value "TopicsDefinitions"
	AuditLogMessageServiceNameTopicsDefinitions string = "TopicsDefinitions"

	// AuditLogMessageServiceNamePredictiveEngagement captures enum value "PredictiveEngagement"
	AuditLogMessageServiceNamePredictiveEngagement string = "PredictiveEngagement"

	// AuditLogMessageServiceNameWorkforceManagement captures enum value "WorkforceManagement"
	AuditLogMessageServiceNameWorkforceManagement string = "WorkforceManagement"

	// AuditLogMessageServiceNameTriggers captures enum value "Triggers"
	AuditLogMessageServiceNameTriggers string = "Triggers"

	// AuditLogMessageServiceNameResponseManagement captures enum value "ResponseManagement"
	AuditLogMessageServiceNameResponseManagement string = "ResponseManagement"

	// AuditLogMessageServiceNameGroups captures enum value "Groups"
	AuditLogMessageServiceNameGroups string = "Groups"

	// AuditLogMessageServiceNameTelephony captures enum value "Telephony"
	AuditLogMessageServiceNameTelephony string = "Telephony"

	// AuditLogMessageServiceNameOutbound captures enum value "Outbound"
	AuditLogMessageServiceNameOutbound string = "Outbound"

	// AuditLogMessageServiceNameSpeechAndTextAnalytics captures enum value "SpeechAndTextAnalytics"
	AuditLogMessageServiceNameSpeechAndTextAnalytics string = "SpeechAndTextAnalytics"

	// AuditLogMessageServiceNameRouting captures enum value "Routing"
	AuditLogMessageServiceNameRouting string = "Routing"
)

// prop value enum
func (m *AuditLogMessage) validateServiceNameEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, auditLogMessageTypeServiceNamePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AuditLogMessage) validateServiceName(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceName) { // not required
		return nil
	}

	// value enum
	if err := m.validateServiceNameEnum("serviceName", "body", m.ServiceName); err != nil {
		return err
	}

	return nil
}

func (m *AuditLogMessage) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditLogMessage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditLogMessage) UnmarshalBinary(b []byte) error {
	var res AuditLogMessage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
