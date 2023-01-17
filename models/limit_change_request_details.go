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

// LimitChangeRequestDetails limit change request details
//
// swagger:model LimitChangeRequestDetails
type LimitChangeRequestDetails struct {

	// The approval breakdown for this override request.
	// Read Only: true
	ApprovalNamespaces []*ApprovalNamespace `json:"approvalNamespaces"`

	// Current limit value for a given key
	// Read Only: true
	CurrentValue float64 `json:"currentValue,omitempty"`

	// The date of the limit change request completion (ChangeImplemented, Rejected, or RollbackImplemented. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCompleted strfmt.DateTime `json:"dateCompleted,omitempty"`

	// The date of the limit change request creation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Description of the need for the limit change request
	// Required: true
	Description *string `json:"description"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Limit key to be overridden (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
	// Required: true
	Key *string `json:"key"`

	// Namespace the key belongs to (see https://developer.mypurecloud.com/api/rest/v2/organization/limits.html#available_limits)
	// Required: true
	// Enum: [contacts agent.assistant analytics.alerting analytics analytics.realtime analytics.reporting.settings architect audiohook audit auth.api authorization automation.testing bots bots.voice cobrowse content.management conversation dataactions datatables directory email event.orchestration external.contacts gcv gdpr groups historical.adherence infrastructureascode integrations intent.miner journey knowledge language.understanding limit.registry marketplace messaging notifications onboarding outbound platform.api predictive.routing quality recording response.management routing scim search secondary.automation.testing skills speech.and.text.analytics speech.integration supportability task.management telephony.configuration web.deployments web.messaging webchat webhooks workforce.management]
	Namespace *string `json:"namespace"`

	// The reason for rejecting the limit override request
	// Read Only: true
	// Enum: [AlternativeExists IncreaseNotRequired PlatformMisuse PlatformStability OtherReason]
	RejectReason string `json:"rejectReason,omitempty"`

	// Requested limit value for a given key
	// Required: true
	RequestedValue *float64 `json:"requestedValue"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Current status of the limit change request
	// Read Only: true
	// Enum: [Approved Rejected Rollback Pending Open SecondaryApprovalNamespacesAdded ReviewerApproved ReviewerRejected ReviewerRollback ImplementingChange ChangeImplemented ImplementingRollback RollbackImplemented]
	Status string `json:"status,omitempty"`

	// List of statuses that a limit change request has gone through
	// Read Only: true
	StatusHistory []*StatusChange `json:"statusHistory"`

	// The support case url created by Care
	// Required: true
	SupportCaseURL *string `json:"supportCaseUrl"`
}

// Validate validates this limit change request details
func (m *LimitChangeRequestDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApprovalNamespaces(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNamespace(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRequestedValue(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatusHistory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportCaseURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LimitChangeRequestDetails) validateApprovalNamespaces(formats strfmt.Registry) error {
	if swag.IsZero(m.ApprovalNamespaces) { // not required
		return nil
	}

	for i := 0; i < len(m.ApprovalNamespaces); i++ {
		if swag.IsZero(m.ApprovalNamespaces[i]) { // not required
			continue
		}

		if m.ApprovalNamespaces[i] != nil {
			if err := m.ApprovalNamespaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("approvalNamespaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("approvalNamespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LimitChangeRequestDetails) validateDateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCompleted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCompleted", "body", "date-time", m.DateCompleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	return nil
}

var limitChangeRequestDetailsTypeNamespacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["contacts","agent.assistant","analytics.alerting","analytics","analytics.realtime","analytics.reporting.settings","architect","audiohook","audit","auth.api","authorization","automation.testing","bots","bots.voice","cobrowse","content.management","conversation","dataactions","datatables","directory","email","event.orchestration","external.contacts","gcv","gdpr","groups","historical.adherence","infrastructureascode","integrations","intent.miner","journey","knowledge","language.understanding","limit.registry","marketplace","messaging","notifications","onboarding","outbound","platform.api","predictive.routing","quality","recording","response.management","routing","scim","search","secondary.automation.testing","skills","speech.and.text.analytics","speech.integration","supportability","task.management","telephony.configuration","web.deployments","web.messaging","webchat","webhooks","workforce.management"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitChangeRequestDetailsTypeNamespacePropEnum = append(limitChangeRequestDetailsTypeNamespacePropEnum, v)
	}
}

const (

	// LimitChangeRequestDetailsNamespaceContacts captures enum value "contacts"
	LimitChangeRequestDetailsNamespaceContacts string = "contacts"

	// LimitChangeRequestDetailsNamespaceAgentDotAssistant captures enum value "agent.assistant"
	LimitChangeRequestDetailsNamespaceAgentDotAssistant string = "agent.assistant"

	// LimitChangeRequestDetailsNamespaceAnalyticsDotAlerting captures enum value "analytics.alerting"
	LimitChangeRequestDetailsNamespaceAnalyticsDotAlerting string = "analytics.alerting"

	// LimitChangeRequestDetailsNamespaceAnalytics captures enum value "analytics"
	LimitChangeRequestDetailsNamespaceAnalytics string = "analytics"

	// LimitChangeRequestDetailsNamespaceAnalyticsDotRealtime captures enum value "analytics.realtime"
	LimitChangeRequestDetailsNamespaceAnalyticsDotRealtime string = "analytics.realtime"

	// LimitChangeRequestDetailsNamespaceAnalyticsDotReportingDotSettings captures enum value "analytics.reporting.settings"
	LimitChangeRequestDetailsNamespaceAnalyticsDotReportingDotSettings string = "analytics.reporting.settings"

	// LimitChangeRequestDetailsNamespaceArchitect captures enum value "architect"
	LimitChangeRequestDetailsNamespaceArchitect string = "architect"

	// LimitChangeRequestDetailsNamespaceAudiohook captures enum value "audiohook"
	LimitChangeRequestDetailsNamespaceAudiohook string = "audiohook"

	// LimitChangeRequestDetailsNamespaceAudit captures enum value "audit"
	LimitChangeRequestDetailsNamespaceAudit string = "audit"

	// LimitChangeRequestDetailsNamespaceAuthDotAPI captures enum value "auth.api"
	LimitChangeRequestDetailsNamespaceAuthDotAPI string = "auth.api"

	// LimitChangeRequestDetailsNamespaceAuthorization captures enum value "authorization"
	LimitChangeRequestDetailsNamespaceAuthorization string = "authorization"

	// LimitChangeRequestDetailsNamespaceAutomationDotTesting captures enum value "automation.testing"
	LimitChangeRequestDetailsNamespaceAutomationDotTesting string = "automation.testing"

	// LimitChangeRequestDetailsNamespaceBots captures enum value "bots"
	LimitChangeRequestDetailsNamespaceBots string = "bots"

	// LimitChangeRequestDetailsNamespaceBotsDotVoice captures enum value "bots.voice"
	LimitChangeRequestDetailsNamespaceBotsDotVoice string = "bots.voice"

	// LimitChangeRequestDetailsNamespaceCobrowse captures enum value "cobrowse"
	LimitChangeRequestDetailsNamespaceCobrowse string = "cobrowse"

	// LimitChangeRequestDetailsNamespaceContentDotManagement captures enum value "content.management"
	LimitChangeRequestDetailsNamespaceContentDotManagement string = "content.management"

	// LimitChangeRequestDetailsNamespaceConversation captures enum value "conversation"
	LimitChangeRequestDetailsNamespaceConversation string = "conversation"

	// LimitChangeRequestDetailsNamespaceDataactions captures enum value "dataactions"
	LimitChangeRequestDetailsNamespaceDataactions string = "dataactions"

	// LimitChangeRequestDetailsNamespaceDatatables captures enum value "datatables"
	LimitChangeRequestDetailsNamespaceDatatables string = "datatables"

	// LimitChangeRequestDetailsNamespaceDirectory captures enum value "directory"
	LimitChangeRequestDetailsNamespaceDirectory string = "directory"

	// LimitChangeRequestDetailsNamespaceEmail captures enum value "email"
	LimitChangeRequestDetailsNamespaceEmail string = "email"

	// LimitChangeRequestDetailsNamespaceEventDotOrchestration captures enum value "event.orchestration"
	LimitChangeRequestDetailsNamespaceEventDotOrchestration string = "event.orchestration"

	// LimitChangeRequestDetailsNamespaceExternalDotContacts captures enum value "external.contacts"
	LimitChangeRequestDetailsNamespaceExternalDotContacts string = "external.contacts"

	// LimitChangeRequestDetailsNamespaceGcv captures enum value "gcv"
	LimitChangeRequestDetailsNamespaceGcv string = "gcv"

	// LimitChangeRequestDetailsNamespaceGdpr captures enum value "gdpr"
	LimitChangeRequestDetailsNamespaceGdpr string = "gdpr"

	// LimitChangeRequestDetailsNamespaceGroups captures enum value "groups"
	LimitChangeRequestDetailsNamespaceGroups string = "groups"

	// LimitChangeRequestDetailsNamespaceHistoricalDotAdherence captures enum value "historical.adherence"
	LimitChangeRequestDetailsNamespaceHistoricalDotAdherence string = "historical.adherence"

	// LimitChangeRequestDetailsNamespaceInfrastructureascode captures enum value "infrastructureascode"
	LimitChangeRequestDetailsNamespaceInfrastructureascode string = "infrastructureascode"

	// LimitChangeRequestDetailsNamespaceIntegrations captures enum value "integrations"
	LimitChangeRequestDetailsNamespaceIntegrations string = "integrations"

	// LimitChangeRequestDetailsNamespaceIntentDotMiner captures enum value "intent.miner"
	LimitChangeRequestDetailsNamespaceIntentDotMiner string = "intent.miner"

	// LimitChangeRequestDetailsNamespaceJourney captures enum value "journey"
	LimitChangeRequestDetailsNamespaceJourney string = "journey"

	// LimitChangeRequestDetailsNamespaceKnowledge captures enum value "knowledge"
	LimitChangeRequestDetailsNamespaceKnowledge string = "knowledge"

	// LimitChangeRequestDetailsNamespaceLanguageDotUnderstanding captures enum value "language.understanding"
	LimitChangeRequestDetailsNamespaceLanguageDotUnderstanding string = "language.understanding"

	// LimitChangeRequestDetailsNamespaceLimitDotRegistry captures enum value "limit.registry"
	LimitChangeRequestDetailsNamespaceLimitDotRegistry string = "limit.registry"

	// LimitChangeRequestDetailsNamespaceMarketplace captures enum value "marketplace"
	LimitChangeRequestDetailsNamespaceMarketplace string = "marketplace"

	// LimitChangeRequestDetailsNamespaceMessaging captures enum value "messaging"
	LimitChangeRequestDetailsNamespaceMessaging string = "messaging"

	// LimitChangeRequestDetailsNamespaceNotifications captures enum value "notifications"
	LimitChangeRequestDetailsNamespaceNotifications string = "notifications"

	// LimitChangeRequestDetailsNamespaceOnboarding captures enum value "onboarding"
	LimitChangeRequestDetailsNamespaceOnboarding string = "onboarding"

	// LimitChangeRequestDetailsNamespaceOutbound captures enum value "outbound"
	LimitChangeRequestDetailsNamespaceOutbound string = "outbound"

	// LimitChangeRequestDetailsNamespacePlatformDotAPI captures enum value "platform.api"
	LimitChangeRequestDetailsNamespacePlatformDotAPI string = "platform.api"

	// LimitChangeRequestDetailsNamespacePredictiveDotRouting captures enum value "predictive.routing"
	LimitChangeRequestDetailsNamespacePredictiveDotRouting string = "predictive.routing"

	// LimitChangeRequestDetailsNamespaceQuality captures enum value "quality"
	LimitChangeRequestDetailsNamespaceQuality string = "quality"

	// LimitChangeRequestDetailsNamespaceRecording captures enum value "recording"
	LimitChangeRequestDetailsNamespaceRecording string = "recording"

	// LimitChangeRequestDetailsNamespaceResponseDotManagement captures enum value "response.management"
	LimitChangeRequestDetailsNamespaceResponseDotManagement string = "response.management"

	// LimitChangeRequestDetailsNamespaceRouting captures enum value "routing"
	LimitChangeRequestDetailsNamespaceRouting string = "routing"

	// LimitChangeRequestDetailsNamespaceScim captures enum value "scim"
	LimitChangeRequestDetailsNamespaceScim string = "scim"

	// LimitChangeRequestDetailsNamespaceSearch captures enum value "search"
	LimitChangeRequestDetailsNamespaceSearch string = "search"

	// LimitChangeRequestDetailsNamespaceSecondaryDotAutomationDotTesting captures enum value "secondary.automation.testing"
	LimitChangeRequestDetailsNamespaceSecondaryDotAutomationDotTesting string = "secondary.automation.testing"

	// LimitChangeRequestDetailsNamespaceSkills captures enum value "skills"
	LimitChangeRequestDetailsNamespaceSkills string = "skills"

	// LimitChangeRequestDetailsNamespaceSpeechDotAndDotTextDotAnalytics captures enum value "speech.and.text.analytics"
	LimitChangeRequestDetailsNamespaceSpeechDotAndDotTextDotAnalytics string = "speech.and.text.analytics"

	// LimitChangeRequestDetailsNamespaceSpeechDotIntegration captures enum value "speech.integration"
	LimitChangeRequestDetailsNamespaceSpeechDotIntegration string = "speech.integration"

	// LimitChangeRequestDetailsNamespaceSupportability captures enum value "supportability"
	LimitChangeRequestDetailsNamespaceSupportability string = "supportability"

	// LimitChangeRequestDetailsNamespaceTaskDotManagement captures enum value "task.management"
	LimitChangeRequestDetailsNamespaceTaskDotManagement string = "task.management"

	// LimitChangeRequestDetailsNamespaceTelephonyDotConfiguration captures enum value "telephony.configuration"
	LimitChangeRequestDetailsNamespaceTelephonyDotConfiguration string = "telephony.configuration"

	// LimitChangeRequestDetailsNamespaceWebDotDeployments captures enum value "web.deployments"
	LimitChangeRequestDetailsNamespaceWebDotDeployments string = "web.deployments"

	// LimitChangeRequestDetailsNamespaceWebDotMessaging captures enum value "web.messaging"
	LimitChangeRequestDetailsNamespaceWebDotMessaging string = "web.messaging"

	// LimitChangeRequestDetailsNamespaceWebchat captures enum value "webchat"
	LimitChangeRequestDetailsNamespaceWebchat string = "webchat"

	// LimitChangeRequestDetailsNamespaceWebhooks captures enum value "webhooks"
	LimitChangeRequestDetailsNamespaceWebhooks string = "webhooks"

	// LimitChangeRequestDetailsNamespaceWorkforceDotManagement captures enum value "workforce.management"
	LimitChangeRequestDetailsNamespaceWorkforceDotManagement string = "workforce.management"
)

// prop value enum
func (m *LimitChangeRequestDetails) validateNamespaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, limitChangeRequestDetailsTypeNamespacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LimitChangeRequestDetails) validateNamespace(formats strfmt.Registry) error {

	if err := validate.Required("namespace", "body", m.Namespace); err != nil {
		return err
	}

	// value enum
	if err := m.validateNamespaceEnum("namespace", "body", *m.Namespace); err != nil {
		return err
	}

	return nil
}

var limitChangeRequestDetailsTypeRejectReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AlternativeExists","IncreaseNotRequired","PlatformMisuse","PlatformStability","OtherReason"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitChangeRequestDetailsTypeRejectReasonPropEnum = append(limitChangeRequestDetailsTypeRejectReasonPropEnum, v)
	}
}

const (

	// LimitChangeRequestDetailsRejectReasonAlternativeExists captures enum value "AlternativeExists"
	LimitChangeRequestDetailsRejectReasonAlternativeExists string = "AlternativeExists"

	// LimitChangeRequestDetailsRejectReasonIncreaseNotRequired captures enum value "IncreaseNotRequired"
	LimitChangeRequestDetailsRejectReasonIncreaseNotRequired string = "IncreaseNotRequired"

	// LimitChangeRequestDetailsRejectReasonPlatformMisuse captures enum value "PlatformMisuse"
	LimitChangeRequestDetailsRejectReasonPlatformMisuse string = "PlatformMisuse"

	// LimitChangeRequestDetailsRejectReasonPlatformStability captures enum value "PlatformStability"
	LimitChangeRequestDetailsRejectReasonPlatformStability string = "PlatformStability"

	// LimitChangeRequestDetailsRejectReasonOtherReason captures enum value "OtherReason"
	LimitChangeRequestDetailsRejectReasonOtherReason string = "OtherReason"
)

// prop value enum
func (m *LimitChangeRequestDetails) validateRejectReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, limitChangeRequestDetailsTypeRejectReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LimitChangeRequestDetails) validateRejectReason(formats strfmt.Registry) error {
	if swag.IsZero(m.RejectReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateRejectReasonEnum("rejectReason", "body", m.RejectReason); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) validateRequestedValue(formats strfmt.Registry) error {

	if err := validate.Required("requestedValue", "body", m.RequestedValue); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var limitChangeRequestDetailsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Approved","Rejected","Rollback","Pending","Open","SecondaryApprovalNamespacesAdded","ReviewerApproved","ReviewerRejected","ReviewerRollback","ImplementingChange","ChangeImplemented","ImplementingRollback","RollbackImplemented"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		limitChangeRequestDetailsTypeStatusPropEnum = append(limitChangeRequestDetailsTypeStatusPropEnum, v)
	}
}

const (

	// LimitChangeRequestDetailsStatusApproved captures enum value "Approved"
	LimitChangeRequestDetailsStatusApproved string = "Approved"

	// LimitChangeRequestDetailsStatusRejected captures enum value "Rejected"
	LimitChangeRequestDetailsStatusRejected string = "Rejected"

	// LimitChangeRequestDetailsStatusRollback captures enum value "Rollback"
	LimitChangeRequestDetailsStatusRollback string = "Rollback"

	// LimitChangeRequestDetailsStatusPending captures enum value "Pending"
	LimitChangeRequestDetailsStatusPending string = "Pending"

	// LimitChangeRequestDetailsStatusOpen captures enum value "Open"
	LimitChangeRequestDetailsStatusOpen string = "Open"

	// LimitChangeRequestDetailsStatusSecondaryApprovalNamespacesAdded captures enum value "SecondaryApprovalNamespacesAdded"
	LimitChangeRequestDetailsStatusSecondaryApprovalNamespacesAdded string = "SecondaryApprovalNamespacesAdded"

	// LimitChangeRequestDetailsStatusReviewerApproved captures enum value "ReviewerApproved"
	LimitChangeRequestDetailsStatusReviewerApproved string = "ReviewerApproved"

	// LimitChangeRequestDetailsStatusReviewerRejected captures enum value "ReviewerRejected"
	LimitChangeRequestDetailsStatusReviewerRejected string = "ReviewerRejected"

	// LimitChangeRequestDetailsStatusReviewerRollback captures enum value "ReviewerRollback"
	LimitChangeRequestDetailsStatusReviewerRollback string = "ReviewerRollback"

	// LimitChangeRequestDetailsStatusImplementingChange captures enum value "ImplementingChange"
	LimitChangeRequestDetailsStatusImplementingChange string = "ImplementingChange"

	// LimitChangeRequestDetailsStatusChangeImplemented captures enum value "ChangeImplemented"
	LimitChangeRequestDetailsStatusChangeImplemented string = "ChangeImplemented"

	// LimitChangeRequestDetailsStatusImplementingRollback captures enum value "ImplementingRollback"
	LimitChangeRequestDetailsStatusImplementingRollback string = "ImplementingRollback"

	// LimitChangeRequestDetailsStatusRollbackImplemented captures enum value "RollbackImplemented"
	LimitChangeRequestDetailsStatusRollbackImplemented string = "RollbackImplemented"
)

// prop value enum
func (m *LimitChangeRequestDetails) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, limitChangeRequestDetailsTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LimitChangeRequestDetails) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) validateStatusHistory(formats strfmt.Registry) error {
	if swag.IsZero(m.StatusHistory) { // not required
		return nil
	}

	for i := 0; i < len(m.StatusHistory); i++ {
		if swag.IsZero(m.StatusHistory[i]) { // not required
			continue
		}

		if m.StatusHistory[i] != nil {
			if err := m.StatusHistory[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statusHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statusHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LimitChangeRequestDetails) validateSupportCaseURL(formats strfmt.Registry) error {

	if err := validate.Required("supportCaseUrl", "body", m.SupportCaseURL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this limit change request details based on the context it is used
func (m *LimitChangeRequestDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApprovalNamespaces(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCurrentValue(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCompleted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRejectReason(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatusHistory(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LimitChangeRequestDetails) contextValidateApprovalNamespaces(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "approvalNamespaces", "body", []*ApprovalNamespace(m.ApprovalNamespaces)); err != nil {
		return err
	}

	for i := 0; i < len(m.ApprovalNamespaces); i++ {

		if m.ApprovalNamespaces[i] != nil {
			if err := m.ApprovalNamespaces[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("approvalNamespaces" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("approvalNamespaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateCurrentValue(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "currentValue", "body", float64(m.CurrentValue)); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateDateCompleted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCompleted", "body", strfmt.DateTime(m.DateCompleted)); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateRejectReason(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "rejectReason", "body", string(m.RejectReason)); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *LimitChangeRequestDetails) contextValidateStatusHistory(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "statusHistory", "body", []*StatusChange(m.StatusHistory)); err != nil {
		return err
	}

	for i := 0; i < len(m.StatusHistory); i++ {

		if m.StatusHistory[i] != nil {
			if err := m.StatusHistory[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("statusHistory" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("statusHistory" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *LimitChangeRequestDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LimitChangeRequestDetails) UnmarshalBinary(b []byte) error {
	var res LimitChangeRequestDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
