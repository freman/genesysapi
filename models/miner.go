// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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
	// Read Only: true
	// Format: date
	ConversationsDateRangeEnd strfmt.Date `json:"conversationsDateRangeEnd,omitempty"`

	// Date from which the conversations need to be taken for mining. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Read Only: true
	// Format: date
	ConversationsDateRangeStart strfmt.Date `json:"conversationsDateRangeStart,omitempty"`

	// Date when the mining process was completed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCompleted strfmt.DateTime `json:"dateCompleted,omitempty"`

	// Date when the miner was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date when the miner was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Date when the miner started execution. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateTriggered strfmt.DateTime `json:"dateTriggered,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Language Localization code.
	// Enum: [en-us en-gb en-au en-in en-za]
	Language string `json:"language,omitempty"`

	// Latest draft details of the miner.
	// Read Only: true
	LatestDraftVersion *Draft `json:"latestDraftVersion,omitempty"`

	// Media type for filtering conversations.
	// Read Only: true
	// Enum: [Chat Call]
	MediaType string `json:"mediaType,omitempty"`

	// Mining message if present.
	// Read Only: true
	Message string `json:"message,omitempty"`

	// Chat Corpus Name.
	// Required: true
	Name *string `json:"name"`

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

	if err := m.validateLanguage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestDraftVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

var minerTypeLanguagePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["en-us","en-gb","en-au","en-in","en-za"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		minerTypeLanguagePropEnum = append(minerTypeLanguagePropEnum, v)
	}
}

const (

	// MinerLanguageEnUs captures enum value "en-us"
	MinerLanguageEnUs string = "en-us"

	// MinerLanguageEnGb captures enum value "en-gb"
	MinerLanguageEnGb string = "en-gb"

	// MinerLanguageEnAu captures enum value "en-au"
	MinerLanguageEnAu string = "en-au"

	// MinerLanguageEnIn captures enum value "en-in"
	MinerLanguageEnIn string = "en-in"

	// MinerLanguageEnZa captures enum value "en-za"
	MinerLanguageEnZa string = "en-za"
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
			}
			return err
		}
	}

	return nil
}

var minerTypeMediaTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Chat","Call"]`), &res); err != nil {
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

func (m *Miner) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
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
