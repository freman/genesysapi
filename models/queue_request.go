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

// QueueRequest queue request
//
// swagger:model QueueRequest
type QueueRequest struct {

	// The ACW settings for the queue.
	AcwSettings *AcwSettings `json:"acwSettings,omitempty"`

	// Specifies whether the configured whisper should play for all ACD calls, or only for those which are auto-answered.
	AutoAnswerOnly bool `json:"autoAnswerOnly"`

	// The bulls-eye settings for the queue.
	Bullseye *Bullseye `json:"bullseye,omitempty"`

	// The name to use for caller identification for outbound calls from this queue.
	CallingPartyName string `json:"callingPartyName,omitempty"`

	// The phone number to use for caller identification for outbound calls from this queue.
	CallingPartyNumber string `json:"callingPartyNumber,omitempty"`

	// The ID of the user that created the queue.
	CreatedBy string `json:"createdBy,omitempty"`

	// The date the queue was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the queue. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The default script Ids for the communication types.
	DefaultScripts map[string]Script `json:"defaultScripts,omitempty"`

	// The queue description.
	Description string `json:"description,omitempty"`

	// The division to which this entity belongs.
	Division *WritableDivision `json:"division,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The media settings for the queue. Valid key values: CALL, CALLBACK, CHAT, EMAIL, MESSAGE, SOCIAL_EXPRESSION, VIDEO_COMM
	MediaSettings map[string]MediaSetting `json:"mediaSettings,omitempty"`

	// The number of users in the queue.
	// Read Only: true
	MemberCount int32 `json:"memberCount,omitempty"`

	// The ID of the user that last modified the queue.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The queue name
	// Required: true
	Name *string `json:"name"`

	// outbound email address
	OutboundEmailAddress *QueueEmailAddress `json:"outboundEmailAddress,omitempty"`

	// The messaging addresses for the queue.
	OutboundMessagingAddresses *QueueMessagingAddresses `json:"outboundMessagingAddresses,omitempty"`

	// The in-queue flow to use for conversations waiting in queue.
	QueueFlow *DomainEntityRef `json:"queueFlow,omitempty"`

	// The routing rules for the queue, used for routing to known or preferred agents.
	RoutingRules []*RoutingRule `json:"routingRules"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The skill evaluation method to use when routing conversations.
	// Enum: [NONE BEST ALL]
	SkillEvaluationMethod string `json:"skillEvaluationMethod,omitempty"`

	// The prompt used for whisper on the queue, if configured.
	WhisperPrompt *DomainEntityRef `json:"whisperPrompt,omitempty"`
}

// Validate validates this queue request
func (m *QueueRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcwSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBullseye(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDefaultScripts(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundEmailAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOutboundMessagingAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueueFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingRules(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkillEvaluationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhisperPrompt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QueueRequest) validateAcwSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.AcwSettings) { // not required
		return nil
	}

	if m.AcwSettings != nil {
		if err := m.AcwSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acwSettings")
			}
			return err
		}
	}

	return nil
}

func (m *QueueRequest) validateBullseye(formats strfmt.Registry) error {

	if swag.IsZero(m.Bullseye) { // not required
		return nil
	}

	if m.Bullseye != nil {
		if err := m.Bullseye.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bullseye")
			}
			return err
		}
	}

	return nil
}

func (m *QueueRequest) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueueRequest) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *QueueRequest) validateDefaultScripts(formats strfmt.Registry) error {

	if swag.IsZero(m.DefaultScripts) { // not required
		return nil
	}

	for k := range m.DefaultScripts {

		if err := validate.Required("defaultScripts"+"."+k, "body", m.DefaultScripts[k]); err != nil {
			return err
		}
		if val, ok := m.DefaultScripts[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *QueueRequest) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *QueueRequest) validateMediaSettings(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaSettings) { // not required
		return nil
	}

	for k := range m.MediaSettings {

		if err := validate.Required("mediaSettings"+"."+k, "body", m.MediaSettings[k]); err != nil {
			return err
		}
		if val, ok := m.MediaSettings[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *QueueRequest) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *QueueRequest) validateOutboundEmailAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.OutboundEmailAddress) { // not required
		return nil
	}

	if m.OutboundEmailAddress != nil {
		if err := m.OutboundEmailAddress.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundEmailAddress")
			}
			return err
		}
	}

	return nil
}

func (m *QueueRequest) validateOutboundMessagingAddresses(formats strfmt.Registry) error {

	if swag.IsZero(m.OutboundMessagingAddresses) { // not required
		return nil
	}

	if m.OutboundMessagingAddresses != nil {
		if err := m.OutboundMessagingAddresses.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("outboundMessagingAddresses")
			}
			return err
		}
	}

	return nil
}

func (m *QueueRequest) validateQueueFlow(formats strfmt.Registry) error {

	if swag.IsZero(m.QueueFlow) { // not required
		return nil
	}

	if m.QueueFlow != nil {
		if err := m.QueueFlow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("queueFlow")
			}
			return err
		}
	}

	return nil
}

func (m *QueueRequest) validateRoutingRules(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingRules) { // not required
		return nil
	}

	for i := 0; i < len(m.RoutingRules); i++ {
		if swag.IsZero(m.RoutingRules[i]) { // not required
			continue
		}

		if m.RoutingRules[i] != nil {
			if err := m.RoutingRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("routingRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *QueueRequest) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var queueRequestTypeSkillEvaluationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","BEST","ALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		queueRequestTypeSkillEvaluationMethodPropEnum = append(queueRequestTypeSkillEvaluationMethodPropEnum, v)
	}
}

const (

	// QueueRequestSkillEvaluationMethodNONE captures enum value "NONE"
	QueueRequestSkillEvaluationMethodNONE string = "NONE"

	// QueueRequestSkillEvaluationMethodBEST captures enum value "BEST"
	QueueRequestSkillEvaluationMethodBEST string = "BEST"

	// QueueRequestSkillEvaluationMethodALL captures enum value "ALL"
	QueueRequestSkillEvaluationMethodALL string = "ALL"
)

// prop value enum
func (m *QueueRequest) validateSkillEvaluationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, queueRequestTypeSkillEvaluationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *QueueRequest) validateSkillEvaluationMethod(formats strfmt.Registry) error {

	if swag.IsZero(m.SkillEvaluationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateSkillEvaluationMethodEnum("skillEvaluationMethod", "body", m.SkillEvaluationMethod); err != nil {
		return err
	}

	return nil
}

func (m *QueueRequest) validateWhisperPrompt(formats strfmt.Registry) error {

	if swag.IsZero(m.WhisperPrompt) { // not required
		return nil
	}

	if m.WhisperPrompt != nil {
		if err := m.WhisperPrompt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("whisperPrompt")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QueueRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QueueRequest) UnmarshalBinary(b []byte) error {
	var res QueueRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}