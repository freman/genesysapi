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
	// Enum: [Document Queue Recording Role VoicemailUserPolicy UserPresence WrapupCode MaxOrgRoutingUtilizationCapacity AccessToken OAuthClient OAuthClientAuthorization AuthOrganization AuthUser OrganizationAuthorizationTrust OrganizationAuthorizationUserTrust BulkActions Feedback Topic Program Segment Outcome SessionType EventType ClickstreamSettings Schedule ScheduleGroup EmergencyGroup IVR Trigger Response DependencyTrackingBuild Flow Prompt PromptResource FlowOutcome FlowMilestone Team Edge EdgeGroup Trunk TrunkBase DID DIDPool Extension ExtensionPool Phone PhoneBase Line LineBase OutboundRoute NumberPlan Site AttemptLimits CallableTimeSet Campaign CampaignRule Sequence ContactList ContactListFilter DNCList CallAnalysisResponseSet RuleSet CampaignSchedule SequenceSchedule OrganizationProperties WrapUpCodeMapping MessagingCampaign TranscriptionSettings SpeechTextAnalyticsSettings Predictor]
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
	if err := json.Unmarshal([]byte(`["Create","View","Update","Delete","Download","Upload","MemberAdd","MemberUpdate","MemberRemove","Read","ApplyProtection","RevokeProtection","UpdateRetention","ReadAll","Execute","Publish","Unpublish","Activate","Checkin","Checkout","Deactivate","Debug","Save","Revert","Transcode","Enable","Disable","Authorize","Deauthorize","Authenticate","ChangePassword","Revoke","Export","Append","Recycle"]`), &res); err != nil {
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
	if err := json.Unmarshal([]byte(`["Document","Queue","Recording","Role","VoicemailUserPolicy","UserPresence","WrapupCode","MaxOrgRoutingUtilizationCapacity","AccessToken","OAuthClient","OAuthClientAuthorization","AuthOrganization","AuthUser","OrganizationAuthorizationTrust","OrganizationAuthorizationUserTrust","BulkActions","Feedback","Topic","Program","Segment","Outcome","SessionType","EventType","ClickstreamSettings","Schedule","ScheduleGroup","EmergencyGroup","IVR","Trigger","Response","DependencyTrackingBuild","Flow","Prompt","PromptResource","FlowOutcome","FlowMilestone","Team","Edge","EdgeGroup","Trunk","TrunkBase","DID","DIDPool","Extension","ExtensionPool","Phone","PhoneBase","Line","LineBase","OutboundRoute","NumberPlan","Site","AttemptLimits","CallableTimeSet","Campaign","CampaignRule","Sequence","ContactList","ContactListFilter","DNCList","CallAnalysisResponseSet","RuleSet","CampaignSchedule","SequenceSchedule","OrganizationProperties","WrapUpCodeMapping","MessagingCampaign","TranscriptionSettings","SpeechTextAnalyticsSettings","Predictor"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditQueryEntityTypeNamePropEnum = append(auditQueryEntityTypeNamePropEnum, v)
	}
}

const (

	// AuditQueryEntityNameDocument captures enum value "Document"
	AuditQueryEntityNameDocument string = "Document"

	// AuditQueryEntityNameQueue captures enum value "Queue"
	AuditQueryEntityNameQueue string = "Queue"

	// AuditQueryEntityNameRecording captures enum value "Recording"
	AuditQueryEntityNameRecording string = "Recording"

	// AuditQueryEntityNameRole captures enum value "Role"
	AuditQueryEntityNameRole string = "Role"

	// AuditQueryEntityNameVoicemailUserPolicy captures enum value "VoicemailUserPolicy"
	AuditQueryEntityNameVoicemailUserPolicy string = "VoicemailUserPolicy"

	// AuditQueryEntityNameUserPresence captures enum value "UserPresence"
	AuditQueryEntityNameUserPresence string = "UserPresence"

	// AuditQueryEntityNameWrapupCode captures enum value "WrapupCode"
	AuditQueryEntityNameWrapupCode string = "WrapupCode"

	// AuditQueryEntityNameMaxOrgRoutingUtilizationCapacity captures enum value "MaxOrgRoutingUtilizationCapacity"
	AuditQueryEntityNameMaxOrgRoutingUtilizationCapacity string = "MaxOrgRoutingUtilizationCapacity"

	// AuditQueryEntityNameAccessToken captures enum value "AccessToken"
	AuditQueryEntityNameAccessToken string = "AccessToken"

	// AuditQueryEntityNameOAuthClient captures enum value "OAuthClient"
	AuditQueryEntityNameOAuthClient string = "OAuthClient"

	// AuditQueryEntityNameOAuthClientAuthorization captures enum value "OAuthClientAuthorization"
	AuditQueryEntityNameOAuthClientAuthorization string = "OAuthClientAuthorization"

	// AuditQueryEntityNameAuthOrganization captures enum value "AuthOrganization"
	AuditQueryEntityNameAuthOrganization string = "AuthOrganization"

	// AuditQueryEntityNameAuthUser captures enum value "AuthUser"
	AuditQueryEntityNameAuthUser string = "AuthUser"

	// AuditQueryEntityNameOrganizationAuthorizationTrust captures enum value "OrganizationAuthorizationTrust"
	AuditQueryEntityNameOrganizationAuthorizationTrust string = "OrganizationAuthorizationTrust"

	// AuditQueryEntityNameOrganizationAuthorizationUserTrust captures enum value "OrganizationAuthorizationUserTrust"
	AuditQueryEntityNameOrganizationAuthorizationUserTrust string = "OrganizationAuthorizationUserTrust"

	// AuditQueryEntityNameBulkActions captures enum value "BulkActions"
	AuditQueryEntityNameBulkActions string = "BulkActions"

	// AuditQueryEntityNameFeedback captures enum value "Feedback"
	AuditQueryEntityNameFeedback string = "Feedback"

	// AuditQueryEntityNameTopic captures enum value "Topic"
	AuditQueryEntityNameTopic string = "Topic"

	// AuditQueryEntityNameProgram captures enum value "Program"
	AuditQueryEntityNameProgram string = "Program"

	// AuditQueryEntityNameSegment captures enum value "Segment"
	AuditQueryEntityNameSegment string = "Segment"

	// AuditQueryEntityNameOutcome captures enum value "Outcome"
	AuditQueryEntityNameOutcome string = "Outcome"

	// AuditQueryEntityNameSessionType captures enum value "SessionType"
	AuditQueryEntityNameSessionType string = "SessionType"

	// AuditQueryEntityNameEventType captures enum value "EventType"
	AuditQueryEntityNameEventType string = "EventType"

	// AuditQueryEntityNameClickstreamSettings captures enum value "ClickstreamSettings"
	AuditQueryEntityNameClickstreamSettings string = "ClickstreamSettings"

	// AuditQueryEntityNameSchedule captures enum value "Schedule"
	AuditQueryEntityNameSchedule string = "Schedule"

	// AuditQueryEntityNameScheduleGroup captures enum value "ScheduleGroup"
	AuditQueryEntityNameScheduleGroup string = "ScheduleGroup"

	// AuditQueryEntityNameEmergencyGroup captures enum value "EmergencyGroup"
	AuditQueryEntityNameEmergencyGroup string = "EmergencyGroup"

	// AuditQueryEntityNameIVR captures enum value "IVR"
	AuditQueryEntityNameIVR string = "IVR"

	// AuditQueryEntityNameTrigger captures enum value "Trigger"
	AuditQueryEntityNameTrigger string = "Trigger"

	// AuditQueryEntityNameResponse captures enum value "Response"
	AuditQueryEntityNameResponse string = "Response"

	// AuditQueryEntityNameDependencyTrackingBuild captures enum value "DependencyTrackingBuild"
	AuditQueryEntityNameDependencyTrackingBuild string = "DependencyTrackingBuild"

	// AuditQueryEntityNameFlow captures enum value "Flow"
	AuditQueryEntityNameFlow string = "Flow"

	// AuditQueryEntityNamePrompt captures enum value "Prompt"
	AuditQueryEntityNamePrompt string = "Prompt"

	// AuditQueryEntityNamePromptResource captures enum value "PromptResource"
	AuditQueryEntityNamePromptResource string = "PromptResource"

	// AuditQueryEntityNameFlowOutcome captures enum value "FlowOutcome"
	AuditQueryEntityNameFlowOutcome string = "FlowOutcome"

	// AuditQueryEntityNameFlowMilestone captures enum value "FlowMilestone"
	AuditQueryEntityNameFlowMilestone string = "FlowMilestone"

	// AuditQueryEntityNameTeam captures enum value "Team"
	AuditQueryEntityNameTeam string = "Team"

	// AuditQueryEntityNameEdge captures enum value "Edge"
	AuditQueryEntityNameEdge string = "Edge"

	// AuditQueryEntityNameEdgeGroup captures enum value "EdgeGroup"
	AuditQueryEntityNameEdgeGroup string = "EdgeGroup"

	// AuditQueryEntityNameTrunk captures enum value "Trunk"
	AuditQueryEntityNameTrunk string = "Trunk"

	// AuditQueryEntityNameTrunkBase captures enum value "TrunkBase"
	AuditQueryEntityNameTrunkBase string = "TrunkBase"

	// AuditQueryEntityNameDID captures enum value "DID"
	AuditQueryEntityNameDID string = "DID"

	// AuditQueryEntityNameDIDPool captures enum value "DIDPool"
	AuditQueryEntityNameDIDPool string = "DIDPool"

	// AuditQueryEntityNameExtension captures enum value "Extension"
	AuditQueryEntityNameExtension string = "Extension"

	// AuditQueryEntityNameExtensionPool captures enum value "ExtensionPool"
	AuditQueryEntityNameExtensionPool string = "ExtensionPool"

	// AuditQueryEntityNamePhone captures enum value "Phone"
	AuditQueryEntityNamePhone string = "Phone"

	// AuditQueryEntityNamePhoneBase captures enum value "PhoneBase"
	AuditQueryEntityNamePhoneBase string = "PhoneBase"

	// AuditQueryEntityNameLine captures enum value "Line"
	AuditQueryEntityNameLine string = "Line"

	// AuditQueryEntityNameLineBase captures enum value "LineBase"
	AuditQueryEntityNameLineBase string = "LineBase"

	// AuditQueryEntityNameOutboundRoute captures enum value "OutboundRoute"
	AuditQueryEntityNameOutboundRoute string = "OutboundRoute"

	// AuditQueryEntityNameNumberPlan captures enum value "NumberPlan"
	AuditQueryEntityNameNumberPlan string = "NumberPlan"

	// AuditQueryEntityNameSite captures enum value "Site"
	AuditQueryEntityNameSite string = "Site"

	// AuditQueryEntityNameAttemptLimits captures enum value "AttemptLimits"
	AuditQueryEntityNameAttemptLimits string = "AttemptLimits"

	// AuditQueryEntityNameCallableTimeSet captures enum value "CallableTimeSet"
	AuditQueryEntityNameCallableTimeSet string = "CallableTimeSet"

	// AuditQueryEntityNameCampaign captures enum value "Campaign"
	AuditQueryEntityNameCampaign string = "Campaign"

	// AuditQueryEntityNameCampaignRule captures enum value "CampaignRule"
	AuditQueryEntityNameCampaignRule string = "CampaignRule"

	// AuditQueryEntityNameSequence captures enum value "Sequence"
	AuditQueryEntityNameSequence string = "Sequence"

	// AuditQueryEntityNameContactList captures enum value "ContactList"
	AuditQueryEntityNameContactList string = "ContactList"

	// AuditQueryEntityNameContactListFilter captures enum value "ContactListFilter"
	AuditQueryEntityNameContactListFilter string = "ContactListFilter"

	// AuditQueryEntityNameDNCList captures enum value "DNCList"
	AuditQueryEntityNameDNCList string = "DNCList"

	// AuditQueryEntityNameCallAnalysisResponseSet captures enum value "CallAnalysisResponseSet"
	AuditQueryEntityNameCallAnalysisResponseSet string = "CallAnalysisResponseSet"

	// AuditQueryEntityNameRuleSet captures enum value "RuleSet"
	AuditQueryEntityNameRuleSet string = "RuleSet"

	// AuditQueryEntityNameCampaignSchedule captures enum value "CampaignSchedule"
	AuditQueryEntityNameCampaignSchedule string = "CampaignSchedule"

	// AuditQueryEntityNameSequenceSchedule captures enum value "SequenceSchedule"
	AuditQueryEntityNameSequenceSchedule string = "SequenceSchedule"

	// AuditQueryEntityNameOrganizationProperties captures enum value "OrganizationProperties"
	AuditQueryEntityNameOrganizationProperties string = "OrganizationProperties"

	// AuditQueryEntityNameWrapUpCodeMapping captures enum value "WrapUpCodeMapping"
	AuditQueryEntityNameWrapUpCodeMapping string = "WrapUpCodeMapping"

	// AuditQueryEntityNameMessagingCampaign captures enum value "MessagingCampaign"
	AuditQueryEntityNameMessagingCampaign string = "MessagingCampaign"

	// AuditQueryEntityNameTranscriptionSettings captures enum value "TranscriptionSettings"
	AuditQueryEntityNameTranscriptionSettings string = "TranscriptionSettings"

	// AuditQueryEntityNameSpeechTextAnalyticsSettings captures enum value "SpeechTextAnalyticsSettings"
	AuditQueryEntityNameSpeechTextAnalyticsSettings string = "SpeechTextAnalyticsSettings"

	// AuditQueryEntityNamePredictor captures enum value "Predictor"
	AuditQueryEntityNamePredictor string = "Predictor"
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
