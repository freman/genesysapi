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

// StatusChange status change
//
// swagger:model StatusChange
type StatusChange struct {

	// The date of this status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateStatusChanged strfmt.DateTime `json:"dateStatusChanged,omitempty"`

	// A short message describing the status change
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The namespace for the status change
	// Read Only: true
	// Enum: [agent.assistant analytics.alerting analytics analytics.realtime analytics.reporting.settings architect audiohook audit auth.api authorization automation.testing bots bots.voice callback cobrowse content.management conversation dataactions datatables directory email event.orchestration external.contacts gcv gdpr groups historical.adherence infrastructureascode integrations intent.miner journey knowledge language.understanding limit.registry marketplace messaging notifications onboarding outbound platform.api predictive.routing presence quality recording response.management routing scim search secondary.automation.testing skills speech.and.text.analytics speech.integration supportability task.management telephony.configuration users web.deployments web.messaging webchat webhooks workforce.management]
	Namespace string `json:"namespace,omitempty"`

	// The status the change request transitioned from
	// Read Only: true
	// Enum: [Approved Rejected Rollback Pending Open SecondaryApprovalNamespacesAdded ReviewerApproved ReviewerRejected ReviewerRollback ImplementingChange ChangeImplemented ImplementingRollback RollbackImplemented]
	PreviousStatus string `json:"previousStatus,omitempty"`

	// The reason for rejecting the limit override request
	// Read Only: true
	// Enum: [AlternativeExists IncreaseNotRequired PlatformMisuse PlatformStability OtherReason]
	RejectReason string `json:"rejectReason,omitempty"`

	// The status the change request transitioned to
	// Read Only: true
	// Enum: [Approved Rejected Rollback Pending Open SecondaryApprovalNamespacesAdded ReviewerApproved ReviewerRejected ReviewerRollback ImplementingChange ChangeImplemented ImplementingRollback RollbackImplemented]
	Status string `json:"status,omitempty"`
}

// Validate validates this status change
func (m *StatusChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateStatusChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusChange) validateDateStatusChanged(formats strfmt.Registry) error {
	if swag.IsZero(m.DateStatusChanged) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStatusChanged", "body", "date-time", m.DateStatusChanged.String(), formats); err != nil {
		return err
	}

	return nil
}

var statusChangeTypeNamespacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["agent.assistant","analytics.alerting","analytics","analytics.realtime","analytics.reporting.settings","architect","audiohook","audit","auth.api","authorization","automation.testing","bots","bots.voice","callback","cobrowse","content.management","conversation","dataactions","datatables","directory","email","event.orchestration","external.contacts","gcv","gdpr","groups","historical.adherence","infrastructureascode","integrations","intent.miner","journey","knowledge","language.understanding","limit.registry","marketplace","messaging","notifications","onboarding","outbound","platform.api","predictive.routing","presence","quality","recording","response.management","routing","scim","search","secondary.automation.testing","skills","speech.and.text.analytics","speech.integration","supportability","task.management","telephony.configuration","users","web.deployments","web.messaging","webchat","webhooks","workforce.management"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusChangeTypeNamespacePropEnum = append(statusChangeTypeNamespacePropEnum, v)
	}
}

const (

	// StatusChangeNamespaceAgentDotAssistant captures enum value "agent.assistant"
	StatusChangeNamespaceAgentDotAssistant string = "agent.assistant"

	// StatusChangeNamespaceAnalyticsDotAlerting captures enum value "analytics.alerting"
	StatusChangeNamespaceAnalyticsDotAlerting string = "analytics.alerting"

	// StatusChangeNamespaceAnalytics captures enum value "analytics"
	StatusChangeNamespaceAnalytics string = "analytics"

	// StatusChangeNamespaceAnalyticsDotRealtime captures enum value "analytics.realtime"
	StatusChangeNamespaceAnalyticsDotRealtime string = "analytics.realtime"

	// StatusChangeNamespaceAnalyticsDotReportingDotSettings captures enum value "analytics.reporting.settings"
	StatusChangeNamespaceAnalyticsDotReportingDotSettings string = "analytics.reporting.settings"

	// StatusChangeNamespaceArchitect captures enum value "architect"
	StatusChangeNamespaceArchitect string = "architect"

	// StatusChangeNamespaceAudiohook captures enum value "audiohook"
	StatusChangeNamespaceAudiohook string = "audiohook"

	// StatusChangeNamespaceAudit captures enum value "audit"
	StatusChangeNamespaceAudit string = "audit"

	// StatusChangeNamespaceAuthDotAPI captures enum value "auth.api"
	StatusChangeNamespaceAuthDotAPI string = "auth.api"

	// StatusChangeNamespaceAuthorization captures enum value "authorization"
	StatusChangeNamespaceAuthorization string = "authorization"

	// StatusChangeNamespaceAutomationDotTesting captures enum value "automation.testing"
	StatusChangeNamespaceAutomationDotTesting string = "automation.testing"

	// StatusChangeNamespaceBots captures enum value "bots"
	StatusChangeNamespaceBots string = "bots"

	// StatusChangeNamespaceBotsDotVoice captures enum value "bots.voice"
	StatusChangeNamespaceBotsDotVoice string = "bots.voice"

	// StatusChangeNamespaceCallback captures enum value "callback"
	StatusChangeNamespaceCallback string = "callback"

	// StatusChangeNamespaceCobrowse captures enum value "cobrowse"
	StatusChangeNamespaceCobrowse string = "cobrowse"

	// StatusChangeNamespaceContentDotManagement captures enum value "content.management"
	StatusChangeNamespaceContentDotManagement string = "content.management"

	// StatusChangeNamespaceConversation captures enum value "conversation"
	StatusChangeNamespaceConversation string = "conversation"

	// StatusChangeNamespaceDataactions captures enum value "dataactions"
	StatusChangeNamespaceDataactions string = "dataactions"

	// StatusChangeNamespaceDatatables captures enum value "datatables"
	StatusChangeNamespaceDatatables string = "datatables"

	// StatusChangeNamespaceDirectory captures enum value "directory"
	StatusChangeNamespaceDirectory string = "directory"

	// StatusChangeNamespaceEmail captures enum value "email"
	StatusChangeNamespaceEmail string = "email"

	// StatusChangeNamespaceEventDotOrchestration captures enum value "event.orchestration"
	StatusChangeNamespaceEventDotOrchestration string = "event.orchestration"

	// StatusChangeNamespaceExternalDotContacts captures enum value "external.contacts"
	StatusChangeNamespaceExternalDotContacts string = "external.contacts"

	// StatusChangeNamespaceGcv captures enum value "gcv"
	StatusChangeNamespaceGcv string = "gcv"

	// StatusChangeNamespaceGdpr captures enum value "gdpr"
	StatusChangeNamespaceGdpr string = "gdpr"

	// StatusChangeNamespaceGroups captures enum value "groups"
	StatusChangeNamespaceGroups string = "groups"

	// StatusChangeNamespaceHistoricalDotAdherence captures enum value "historical.adherence"
	StatusChangeNamespaceHistoricalDotAdherence string = "historical.adherence"

	// StatusChangeNamespaceInfrastructureascode captures enum value "infrastructureascode"
	StatusChangeNamespaceInfrastructureascode string = "infrastructureascode"

	// StatusChangeNamespaceIntegrations captures enum value "integrations"
	StatusChangeNamespaceIntegrations string = "integrations"

	// StatusChangeNamespaceIntentDotMiner captures enum value "intent.miner"
	StatusChangeNamespaceIntentDotMiner string = "intent.miner"

	// StatusChangeNamespaceJourney captures enum value "journey"
	StatusChangeNamespaceJourney string = "journey"

	// StatusChangeNamespaceKnowledge captures enum value "knowledge"
	StatusChangeNamespaceKnowledge string = "knowledge"

	// StatusChangeNamespaceLanguageDotUnderstanding captures enum value "language.understanding"
	StatusChangeNamespaceLanguageDotUnderstanding string = "language.understanding"

	// StatusChangeNamespaceLimitDotRegistry captures enum value "limit.registry"
	StatusChangeNamespaceLimitDotRegistry string = "limit.registry"

	// StatusChangeNamespaceMarketplace captures enum value "marketplace"
	StatusChangeNamespaceMarketplace string = "marketplace"

	// StatusChangeNamespaceMessaging captures enum value "messaging"
	StatusChangeNamespaceMessaging string = "messaging"

	// StatusChangeNamespaceNotifications captures enum value "notifications"
	StatusChangeNamespaceNotifications string = "notifications"

	// StatusChangeNamespaceOnboarding captures enum value "onboarding"
	StatusChangeNamespaceOnboarding string = "onboarding"

	// StatusChangeNamespaceOutbound captures enum value "outbound"
	StatusChangeNamespaceOutbound string = "outbound"

	// StatusChangeNamespacePlatformDotAPI captures enum value "platform.api"
	StatusChangeNamespacePlatformDotAPI string = "platform.api"

	// StatusChangeNamespacePredictiveDotRouting captures enum value "predictive.routing"
	StatusChangeNamespacePredictiveDotRouting string = "predictive.routing"

	// StatusChangeNamespacePresence captures enum value "presence"
	StatusChangeNamespacePresence string = "presence"

	// StatusChangeNamespaceQuality captures enum value "quality"
	StatusChangeNamespaceQuality string = "quality"

	// StatusChangeNamespaceRecording captures enum value "recording"
	StatusChangeNamespaceRecording string = "recording"

	// StatusChangeNamespaceResponseDotManagement captures enum value "response.management"
	StatusChangeNamespaceResponseDotManagement string = "response.management"

	// StatusChangeNamespaceRouting captures enum value "routing"
	StatusChangeNamespaceRouting string = "routing"

	// StatusChangeNamespaceScim captures enum value "scim"
	StatusChangeNamespaceScim string = "scim"

	// StatusChangeNamespaceSearch captures enum value "search"
	StatusChangeNamespaceSearch string = "search"

	// StatusChangeNamespaceSecondaryDotAutomationDotTesting captures enum value "secondary.automation.testing"
	StatusChangeNamespaceSecondaryDotAutomationDotTesting string = "secondary.automation.testing"

	// StatusChangeNamespaceSkills captures enum value "skills"
	StatusChangeNamespaceSkills string = "skills"

	// StatusChangeNamespaceSpeechDotAndDotTextDotAnalytics captures enum value "speech.and.text.analytics"
	StatusChangeNamespaceSpeechDotAndDotTextDotAnalytics string = "speech.and.text.analytics"

	// StatusChangeNamespaceSpeechDotIntegration captures enum value "speech.integration"
	StatusChangeNamespaceSpeechDotIntegration string = "speech.integration"

	// StatusChangeNamespaceSupportability captures enum value "supportability"
	StatusChangeNamespaceSupportability string = "supportability"

	// StatusChangeNamespaceTaskDotManagement captures enum value "task.management"
	StatusChangeNamespaceTaskDotManagement string = "task.management"

	// StatusChangeNamespaceTelephonyDotConfiguration captures enum value "telephony.configuration"
	StatusChangeNamespaceTelephonyDotConfiguration string = "telephony.configuration"

	// StatusChangeNamespaceUsers captures enum value "users"
	StatusChangeNamespaceUsers string = "users"

	// StatusChangeNamespaceWebDotDeployments captures enum value "web.deployments"
	StatusChangeNamespaceWebDotDeployments string = "web.deployments"

	// StatusChangeNamespaceWebDotMessaging captures enum value "web.messaging"
	StatusChangeNamespaceWebDotMessaging string = "web.messaging"

	// StatusChangeNamespaceWebchat captures enum value "webchat"
	StatusChangeNamespaceWebchat string = "webchat"

	// StatusChangeNamespaceWebhooks captures enum value "webhooks"
	StatusChangeNamespaceWebhooks string = "webhooks"

	// StatusChangeNamespaceWorkforceDotManagement captures enum value "workforce.management"
	StatusChangeNamespaceWorkforceDotManagement string = "workforce.management"
)

// prop value enum
func (m *StatusChange) validateNamespaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusChangeTypeNamespacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusChange) validateNamespace(formats strfmt.Registry) error {
	if swag.IsZero(m.Namespace) { // not required
		return nil
	}

	// value enum
	if err := m.validateNamespaceEnum("namespace", "body", m.Namespace); err != nil {
		return err
	}

	return nil
}

var statusChangeTypePreviousStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Approved","Rejected","Rollback","Pending","Open","SecondaryApprovalNamespacesAdded","ReviewerApproved","ReviewerRejected","ReviewerRollback","ImplementingChange","ChangeImplemented","ImplementingRollback","RollbackImplemented"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusChangeTypePreviousStatusPropEnum = append(statusChangeTypePreviousStatusPropEnum, v)
	}
}

const (

	// StatusChangePreviousStatusApproved captures enum value "Approved"
	StatusChangePreviousStatusApproved string = "Approved"

	// StatusChangePreviousStatusRejected captures enum value "Rejected"
	StatusChangePreviousStatusRejected string = "Rejected"

	// StatusChangePreviousStatusRollback captures enum value "Rollback"
	StatusChangePreviousStatusRollback string = "Rollback"

	// StatusChangePreviousStatusPending captures enum value "Pending"
	StatusChangePreviousStatusPending string = "Pending"

	// StatusChangePreviousStatusOpen captures enum value "Open"
	StatusChangePreviousStatusOpen string = "Open"

	// StatusChangePreviousStatusSecondaryApprovalNamespacesAdded captures enum value "SecondaryApprovalNamespacesAdded"
	StatusChangePreviousStatusSecondaryApprovalNamespacesAdded string = "SecondaryApprovalNamespacesAdded"

	// StatusChangePreviousStatusReviewerApproved captures enum value "ReviewerApproved"
	StatusChangePreviousStatusReviewerApproved string = "ReviewerApproved"

	// StatusChangePreviousStatusReviewerRejected captures enum value "ReviewerRejected"
	StatusChangePreviousStatusReviewerRejected string = "ReviewerRejected"

	// StatusChangePreviousStatusReviewerRollback captures enum value "ReviewerRollback"
	StatusChangePreviousStatusReviewerRollback string = "ReviewerRollback"

	// StatusChangePreviousStatusImplementingChange captures enum value "ImplementingChange"
	StatusChangePreviousStatusImplementingChange string = "ImplementingChange"

	// StatusChangePreviousStatusChangeImplemented captures enum value "ChangeImplemented"
	StatusChangePreviousStatusChangeImplemented string = "ChangeImplemented"

	// StatusChangePreviousStatusImplementingRollback captures enum value "ImplementingRollback"
	StatusChangePreviousStatusImplementingRollback string = "ImplementingRollback"

	// StatusChangePreviousStatusRollbackImplemented captures enum value "RollbackImplemented"
	StatusChangePreviousStatusRollbackImplemented string = "RollbackImplemented"
)

// prop value enum
func (m *StatusChange) validatePreviousStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusChangeTypePreviousStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusChange) validatePreviousStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.PreviousStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePreviousStatusEnum("previousStatus", "body", m.PreviousStatus); err != nil {
		return err
	}

	return nil
}

var statusChangeTypeRejectReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AlternativeExists","IncreaseNotRequired","PlatformMisuse","PlatformStability","OtherReason"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusChangeTypeRejectReasonPropEnum = append(statusChangeTypeRejectReasonPropEnum, v)
	}
}

const (

	// StatusChangeRejectReasonAlternativeExists captures enum value "AlternativeExists"
	StatusChangeRejectReasonAlternativeExists string = "AlternativeExists"

	// StatusChangeRejectReasonIncreaseNotRequired captures enum value "IncreaseNotRequired"
	StatusChangeRejectReasonIncreaseNotRequired string = "IncreaseNotRequired"

	// StatusChangeRejectReasonPlatformMisuse captures enum value "PlatformMisuse"
	StatusChangeRejectReasonPlatformMisuse string = "PlatformMisuse"

	// StatusChangeRejectReasonPlatformStability captures enum value "PlatformStability"
	StatusChangeRejectReasonPlatformStability string = "PlatformStability"

	// StatusChangeRejectReasonOtherReason captures enum value "OtherReason"
	StatusChangeRejectReasonOtherReason string = "OtherReason"
)

// prop value enum
func (m *StatusChange) validateRejectReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusChangeTypeRejectReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusChange) validateRejectReason(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateRejectReasonEnum("rejectReason", "body", m.RejectReason); err != nil {
		return err
	}

	return nil
}

var statusChangeTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Approved","Rejected","Rollback","Pending","Open","SecondaryApprovalNamespacesAdded","ReviewerApproved","ReviewerRejected","ReviewerRollback","ImplementingChange","ChangeImplemented","ImplementingRollback","RollbackImplemented"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusChangeTypeStatusPropEnum = append(statusChangeTypeStatusPropEnum, v)
	}
}

const (

	// StatusChangeStatusApproved captures enum value "Approved"
	StatusChangeStatusApproved string = "Approved"

	// StatusChangeStatusRejected captures enum value "Rejected"
	StatusChangeStatusRejected string = "Rejected"

	// StatusChangeStatusRollback captures enum value "Rollback"
	StatusChangeStatusRollback string = "Rollback"

	// StatusChangeStatusPending captures enum value "Pending"
	StatusChangeStatusPending string = "Pending"

	// StatusChangeStatusOpen captures enum value "Open"
	StatusChangeStatusOpen string = "Open"

	// StatusChangeStatusSecondaryApprovalNamespacesAdded captures enum value "SecondaryApprovalNamespacesAdded"
	StatusChangeStatusSecondaryApprovalNamespacesAdded string = "SecondaryApprovalNamespacesAdded"

	// StatusChangeStatusReviewerApproved captures enum value "ReviewerApproved"
	StatusChangeStatusReviewerApproved string = "ReviewerApproved"

	// StatusChangeStatusReviewerRejected captures enum value "ReviewerRejected"
	StatusChangeStatusReviewerRejected string = "ReviewerRejected"

	// StatusChangeStatusReviewerRollback captures enum value "ReviewerRollback"
	StatusChangeStatusReviewerRollback string = "ReviewerRollback"

	// StatusChangeStatusImplementingChange captures enum value "ImplementingChange"
	StatusChangeStatusImplementingChange string = "ImplementingChange"

	// StatusChangeStatusChangeImplemented captures enum value "ChangeImplemented"
	StatusChangeStatusChangeImplemented string = "ChangeImplemented"

	// StatusChangeStatusImplementingRollback captures enum value "ImplementingRollback"
	StatusChangeStatusImplementingRollback string = "ImplementingRollback"

	// StatusChangeStatusRollbackImplemented captures enum value "RollbackImplemented"
	StatusChangeStatusRollbackImplemented string = "RollbackImplemented"
)

// prop value enum
func (m *StatusChange) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusChangeTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusChange) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this status change based on the context it is used
func (m *StatusChange) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateStatusChanged(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNamespace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePreviousStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRejectReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusChange) contextValidateDateStatusChanged(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateStatusChanged", "body", strfmt.DateTime(m.DateStatusChanged)); err != nil {
		return err
	}

	return nil
}

func (m *StatusChange) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *StatusChange) contextValidateNamespace(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "namespace", "body", string(m.Namespace)); err != nil {
		return err
	}

	return nil
}

func (m *StatusChange) contextValidatePreviousStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "previousStatus", "body", string(m.PreviousStatus)); err != nil {
		return err
	}

	return nil
}

func (m *StatusChange) contextValidateRejectReason(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rejectReason", "body", string(m.RejectReason)); err != nil {
		return err
	}

	return nil
}

func (m *StatusChange) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusChange) UnmarshalBinary(b []byte) error {
	var res StatusChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
