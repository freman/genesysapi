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

// Miner miner
//
// swagger:model Miner
type Miner struct {

	// Flag to indicate whether data file to be mined was uploaded.
	// Read Only: true
	ConversationDataUploaded *bool `json:"conversationDataUploaded"`

	// Date till which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Example: 2019-12-20
	// Read Only: true
	// Format: date
	ConversationsDateRangeEnd strfmt.Date `json:"conversationsDateRangeEnd,omitempty"`

	// Date from which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Example: 2019-06-20
	// Read Only: true
	// Format: date
	ConversationsDateRangeStart strfmt.Date `json:"conversationsDateRangeStart,omitempty"`

	// Number of conversations/transcripts fetched.
	// Read Only: true
	ConversationsFetchedCount int32 `json:"conversationsFetchedCount,omitempty"`

	// Number of conversations/recordings/transcripts that were found valid for mining purposes.
	// Read Only: true
	ConversationsValidCount int32 `json:"conversationsValidCount,omitempty"`

	// Date when the mining process was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Example: 2020-05-20T23:56:07.268
	// Read Only: true
	// Format: date-time
	DateCompleted strfmt.DateTime `json:"dateCompleted,omitempty"`

	// Date when the miner was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Example: 2020-04-29T17:12:06.613
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date when the miner was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Example: 2020-04-30T23:56:07.268
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Date when the miner started execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Example: 2020-04-30T23:56:07.268
	// Read Only: true
	// Format: date-time
	DateTriggered strfmt.DateTime `json:"dateTriggered,omitempty"`

	// Error Information
	// Read Only: true
	ErrorInfo *ErrorInfo `json:"errorInfo,omitempty"`

	// Number of intents or topics based on the miner type.
	// Read Only: true
	GetminedItemCount int32 `json:"getminedItemCount,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Language Localization code.
	// Enum: [en-us en-gb en-au en-in en-za es-us es-es fr-fr fr-ca de-de]
	Language string `json:"language,omitempty"`

	// Latest draft details of the miner.
	// Read Only: true
	LatestDraftVersion *Draft `json:"latestDraftVersion,omitempty"`

	// Media type for filtering conversations.
	// Read Only: true
	// Enum: [Chat Call Message]
	MediaType string `json:"mediaType,omitempty"`

	// Mining message if present.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Type of the miner, intent or topic.
	// Enum: [Intent Topic]
	MinerType string `json:"minerType,omitempty"`

	// Chat Corpus Name.
	// Required: true
	Name *string `json:"name"`

	// Type of the participant, either agent, customer or both.
	// Read Only: true
	// Enum: [Customer Agent Both]
	ParticipantType string `json:"participantType,omitempty"`

	// List of queue IDs for filtering conversations.
	// Read Only: true
	QueueIds []string `json:"queueIds"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Status of the miner.
	// Read Only: true
	// Enum: [NotStarted FetchingConversationIds ConversationIdsFetched ConversationIdsFetchError FetchingConversations ConversationsFetched ConversationsFetchError Queued QueuingError MiningStarted MaskingUtterances MaskingError ComputingAnalytics ComputingAnalyticsError MiningCompleted MiningError ModelValidationError Deleted]
	Status string `json:"status,omitempty"`

	// Warning Information
	// Read Only: true
	WarningInfo *ErrorInfo `json:"warningInfo,omitempty"`
}

// Validate validates this miner
func (m *Miner) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConversationsDateRangeEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConversationsDateRangeStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateTriggered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateErrorInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestDraftVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMinerType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipantType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarningInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Miner) validateConversationsDateRangeEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationsDateRangeEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationsDateRangeEnd", "body", "date", m.ConversationsDateRangeEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateConversationsDateRangeStart(formats strfmt.Registry) error {
	if swag.IsZero(m.ConversationsDateRangeStart) { // not required
		return nil
	}

	if err := validate.FormatOf("conversationsDateRangeStart", "body", "date", m.ConversationsDateRangeStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateDateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCompleted) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCompleted", "body", "date-time", m.DateCompleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateDateTriggered(formats strfmt.Registry) error {
	if swag.IsZero(m.DateTriggered) { // not required
		return nil
	}

	if err := validate.FormatOf("dateTriggered", "body", "date-time", m.DateTriggered.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateErrorInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.ErrorInfo) { // not required
		return nil
	}

	if m.ErrorInfo != nil {
		if err := m.ErrorInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInfo")
			}
			return err
		}
	}

	return nil
}

var minerTypeLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-us","en-gb","en-au","en-in","en-za","es-us","es-es","fr-fr","fr-ca","de-de"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		minerTypeLanguagePropEnum = append(minerTypeLanguagePropEnum, v)
	}
}

const (

	// MinerLanguageEnDashUs captures enum value "en-us"
	MinerLanguageEnDashUs string = "en-us"

	// MinerLanguageEnDashGb captures enum value "en-gb"
	MinerLanguageEnDashGb string = "en-gb"

	// MinerLanguageEnDashAu captures enum value "en-au"
	MinerLanguageEnDashAu string = "en-au"

	// MinerLanguageEnDashIn captures enum value "en-in"
	MinerLanguageEnDashIn string = "en-in"

	// MinerLanguageEnDashZa captures enum value "en-za"
	MinerLanguageEnDashZa string = "en-za"

	// MinerLanguageEsDashUs captures enum value "es-us"
	MinerLanguageEsDashUs string = "es-us"

	// MinerLanguageEsDashEs captures enum value "es-es"
	MinerLanguageEsDashEs string = "es-es"

	// MinerLanguageFrDashFr captures enum value "fr-fr"
	MinerLanguageFrDashFr string = "fr-fr"

	// MinerLanguageFrDashCa captures enum value "fr-ca"
	MinerLanguageFrDashCa string = "fr-ca"

	// MinerLanguageDeDashDe captures enum value "de-de"
	MinerLanguageDeDashDe string = "de-de"
)

// prop value enum
func (m *Miner) validateLanguageEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, minerTypeLanguagePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Miner) validateLanguage(formats strfmt.Registry) error {
	if swag.IsZero(m.Language) { // not required
		return nil
	}

	// value enum
	if err := m.validateLanguageEnum("language", "body", m.Language); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateLatestDraftVersion(formats strfmt.Registry) error {
	if swag.IsZero(m.LatestDraftVersion) { // not required
		return nil
	}

	if m.LatestDraftVersion != nil {
		if err := m.LatestDraftVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestDraftVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestDraftVersion")
			}
			return err
		}
	}

	return nil
}

var minerTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Chat","Call","Message"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		minerTypeMediaTypePropEnum = append(minerTypeMediaTypePropEnum, v)
	}
}

const (

	// MinerMediaTypeChat captures enum value "Chat"
	MinerMediaTypeChat string = "Chat"

	// MinerMediaTypeCall captures enum value "Call"
	MinerMediaTypeCall string = "Call"

	// MinerMediaTypeMessage captures enum value "Message"
	MinerMediaTypeMessage string = "Message"
)

// prop value enum
func (m *Miner) validateMediaTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, minerTypeMediaTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Miner) validateMediaType(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMediaTypeEnum("mediaType", "body", m.MediaType); err != nil {
		return err
	}

	return nil
}

var minerTypeMinerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Intent","Topic"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		minerTypeMinerTypePropEnum = append(minerTypeMinerTypePropEnum, v)
	}
}

const (

	// MinerMinerTypeIntent captures enum value "Intent"
	MinerMinerTypeIntent string = "Intent"

	// MinerMinerTypeTopic captures enum value "Topic"
	MinerMinerTypeTopic string = "Topic"
)

// prop value enum
func (m *Miner) validateMinerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, minerTypeMinerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Miner) validateMinerType(formats strfmt.Registry) error {
	if swag.IsZero(m.MinerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMinerTypeEnum("minerType", "body", m.MinerType); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var minerTypeParticipantTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Customer","Agent","Both"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		minerTypeParticipantTypePropEnum = append(minerTypeParticipantTypePropEnum, v)
	}
}

const (

	// MinerParticipantTypeCustomer captures enum value "Customer"
	MinerParticipantTypeCustomer string = "Customer"

	// MinerParticipantTypeAgent captures enum value "Agent"
	MinerParticipantTypeAgent string = "Agent"

	// MinerParticipantTypeBoth captures enum value "Both"
	MinerParticipantTypeBoth string = "Both"
)

// prop value enum
func (m *Miner) validateParticipantTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, minerTypeParticipantTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Miner) validateParticipantType(formats strfmt.Registry) error {
	if swag.IsZero(m.ParticipantType) { // not required
		return nil
	}

	// value enum
	if err := m.validateParticipantTypeEnum("participantType", "body", m.ParticipantType); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var minerTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NotStarted","FetchingConversationIds","ConversationIdsFetched","ConversationIdsFetchError","FetchingConversations","ConversationsFetched","ConversationsFetchError","Queued","QueuingError","MiningStarted","MaskingUtterances","MaskingError","ComputingAnalytics","ComputingAnalyticsError","MiningCompleted","MiningError","ModelValidationError","Deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		minerTypeStatusPropEnum = append(minerTypeStatusPropEnum, v)
	}
}

const (

	// MinerStatusNotStarted captures enum value "NotStarted"
	MinerStatusNotStarted string = "NotStarted"

	// MinerStatusFetchingConversationIds captures enum value "FetchingConversationIds"
	MinerStatusFetchingConversationIds string = "FetchingConversationIds"

	// MinerStatusConversationIdsFetched captures enum value "ConversationIdsFetched"
	MinerStatusConversationIdsFetched string = "ConversationIdsFetched"

	// MinerStatusConversationIdsFetchError captures enum value "ConversationIdsFetchError"
	MinerStatusConversationIdsFetchError string = "ConversationIdsFetchError"

	// MinerStatusFetchingConversations captures enum value "FetchingConversations"
	MinerStatusFetchingConversations string = "FetchingConversations"

	// MinerStatusConversationsFetched captures enum value "ConversationsFetched"
	MinerStatusConversationsFetched string = "ConversationsFetched"

	// MinerStatusConversationsFetchError captures enum value "ConversationsFetchError"
	MinerStatusConversationsFetchError string = "ConversationsFetchError"

	// MinerStatusQueued captures enum value "Queued"
	MinerStatusQueued string = "Queued"

	// MinerStatusQueuingError captures enum value "QueuingError"
	MinerStatusQueuingError string = "QueuingError"

	// MinerStatusMiningStarted captures enum value "MiningStarted"
	MinerStatusMiningStarted string = "MiningStarted"

	// MinerStatusMaskingUtterances captures enum value "MaskingUtterances"
	MinerStatusMaskingUtterances string = "MaskingUtterances"

	// MinerStatusMaskingError captures enum value "MaskingError"
	MinerStatusMaskingError string = "MaskingError"

	// MinerStatusComputingAnalytics captures enum value "ComputingAnalytics"
	MinerStatusComputingAnalytics string = "ComputingAnalytics"

	// MinerStatusComputingAnalyticsError captures enum value "ComputingAnalyticsError"
	MinerStatusComputingAnalyticsError string = "ComputingAnalyticsError"

	// MinerStatusMiningCompleted captures enum value "MiningCompleted"
	MinerStatusMiningCompleted string = "MiningCompleted"

	// MinerStatusMiningError captures enum value "MiningError"
	MinerStatusMiningError string = "MiningError"

	// MinerStatusModelValidationError captures enum value "ModelValidationError"
	MinerStatusModelValidationError string = "ModelValidationError"

	// MinerStatusDeleted captures enum value "Deleted"
	MinerStatusDeleted string = "Deleted"
)

// prop value enum
func (m *Miner) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, minerTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Miner) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *Miner) validateWarningInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.WarningInfo) { // not required
		return nil
	}

	if m.WarningInfo != nil {
		if err := m.WarningInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("warningInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("warningInfo")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this miner based on the context it is used
func (m *Miner) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConversationDataUploaded(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversationsDateRangeEnd(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversationsDateRangeStart(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversationsFetchedCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConversationsValidCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCompleted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateTriggered(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateErrorInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGetminedItemCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLatestDraftVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediaType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParticipantType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQueueIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWarningInfo(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Miner) contextValidateConversationDataUploaded(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "conversationDataUploaded", "body", m.ConversationDataUploaded); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateConversationsDateRangeEnd(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "conversationsDateRangeEnd", "body", strfmt.Date(m.ConversationsDateRangeEnd)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateConversationsDateRangeStart(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "conversationsDateRangeStart", "body", strfmt.Date(m.ConversationsDateRangeStart)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateConversationsFetchedCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "conversationsFetchedCount", "body", int32(m.ConversationsFetchedCount)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateConversationsValidCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "conversationsValidCount", "body", int32(m.ConversationsValidCount)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateDateCompleted(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCompleted", "body", strfmt.DateTime(m.DateCompleted)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateDateTriggered(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateTriggered", "body", strfmt.DateTime(m.DateTriggered)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateErrorInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.ErrorInfo != nil {
		if err := m.ErrorInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("errorInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("errorInfo")
			}
			return err
		}
	}

	return nil
}

func (m *Miner) contextValidateGetminedItemCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "getminedItemCount", "body", int32(m.GetminedItemCount)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateLatestDraftVersion(ctx context.Context, formats strfmt.Registry) error {

	if m.LatestDraftVersion != nil {
		if err := m.LatestDraftVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("latestDraftVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("latestDraftVersion")
			}
			return err
		}
	}

	return nil
}

func (m *Miner) contextValidateMediaType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "mediaType", "body", string(m.MediaType)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateMessage(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "message", "body", string(m.Message)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateParticipantType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "participantType", "body", string(m.ParticipantType)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateQueueIds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "queueIds", "body", []string(m.QueueIds)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "status", "body", string(m.Status)); err != nil {
		return err
	}

	return nil
}

func (m *Miner) contextValidateWarningInfo(ctx context.Context, formats strfmt.Registry) error {

	if m.WarningInfo != nil {
		if err := m.WarningInfo.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("warningInfo")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("warningInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Miner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Miner) UnmarshalBinary(b []byte) error {
	var res Miner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
