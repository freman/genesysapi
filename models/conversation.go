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

// Conversation conversation
//
// swagger:model Conversation
type Conversation struct {

	// The address of the conversation as seen from an external participant. For phone calls this will be the DNIS for inbound calls and the ANI for outbound calls. For other media types this will be the address of the destination participant for inbound and the address of the initiating participant for outbound.
	Address string `json:"address,omitempty"`

	// A list of conversations to merge into this conversation to create a conference. This field is null except when being used to create a conference.
	ConversationIds []string `json:"conversationIds"`

	// Identifiers of divisions associated with this conversation
	Divisions []*ConversationDivisionMembership `json:"divisions"`

	// The time when the conversation ended. This will be the time when the last participant left the conversation, or null when the conversation is still active. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	EndTime strfmt.DateTime `json:"endTime,omitempty"`

	// The external tag associated with the conversation.
	ExternalTag string `json:"externalTag,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
	MaxParticipants int32 `json:"maxParticipants,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The list of all participants in the conversation.
	// Required: true
	Participants []*Participant `json:"participants"`

	// The list of the most recent 20 transfer commands applied to this conversation.
	RecentTransfers []*TransferResponse `json:"recentTransfers"`

	// On update, 'paused' initiates a secure pause, 'active' resumes any paused recordings; otherwise indicates state of conversation recording.
	// Enum: [ACTIVE PAUSED NONE]
	RecordingState string `json:"recordingState,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The time when the conversation started. This will be the time when the first participant joined the conversation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Required: true
	// Format: date-time
	StartTime *strfmt.DateTime `json:"startTime"`

	// The conversation's state
	// Enum: [alerting dialing contacting offering connected disconnected terminated converting uploading transmitting none]
	State string `json:"state,omitempty"`
}

// Validate validates this conversation
func (m *Conversation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecentTransfers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordingState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Conversation) validateDivisions(formats strfmt.Registry) error {
	if swag.IsZero(m.Divisions) { // not required
		return nil
	}

	for i := 0; i < len(m.Divisions); i++ {
		if swag.IsZero(m.Divisions[i]) { // not required
			continue
		}

		if m.Divisions[i] != nil {
			if err := m.Divisions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("divisions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("divisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Conversation) validateEndTime(formats strfmt.Registry) error {
	if swag.IsZero(m.EndTime) { // not required
		return nil
	}

	if err := validate.FormatOf("endTime", "body", "date-time", m.EndTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Conversation) validateParticipants(formats strfmt.Registry) error {

	if err := validate.Required("participants", "body", m.Participants); err != nil {
		return err
	}

	for i := 0; i < len(m.Participants); i++ {
		if swag.IsZero(m.Participants[i]) { // not required
			continue
		}

		if m.Participants[i] != nil {
			if err := m.Participants[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Conversation) validateRecentTransfers(formats strfmt.Registry) error {
	if swag.IsZero(m.RecentTransfers) { // not required
		return nil
	}

	for i := 0; i < len(m.RecentTransfers); i++ {
		if swag.IsZero(m.RecentTransfers[i]) { // not required
			continue
		}

		if m.RecentTransfers[i] != nil {
			if err := m.RecentTransfers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recentTransfers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recentTransfers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var conversationTypeRecordingStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","PAUSED","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationTypeRecordingStatePropEnum = append(conversationTypeRecordingStatePropEnum, v)
	}
}

const (

	// ConversationRecordingStateACTIVE captures enum value "ACTIVE"
	ConversationRecordingStateACTIVE string = "ACTIVE"

	// ConversationRecordingStatePAUSED captures enum value "PAUSED"
	ConversationRecordingStatePAUSED string = "PAUSED"

	// ConversationRecordingStateNONE captures enum value "NONE"
	ConversationRecordingStateNONE string = "NONE"
)

// prop value enum
func (m *Conversation) validateRecordingStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationTypeRecordingStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Conversation) validateRecordingState(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordingState) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecordingStateEnum("recordingState", "body", m.RecordingState); err != nil {
		return err
	}

	return nil
}

func (m *Conversation) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Conversation) validateStartTime(formats strfmt.Registry) error {

	if err := validate.Required("startTime", "body", m.StartTime); err != nil {
		return err
	}

	if err := validate.FormatOf("startTime", "body", "date-time", m.StartTime.String(), formats); err != nil {
		return err
	}

	return nil
}

var conversationTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alerting","dialing","contacting","offering","connected","disconnected","terminated","converting","uploading","transmitting","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conversationTypeStatePropEnum = append(conversationTypeStatePropEnum, v)
	}
}

const (

	// ConversationStateAlerting captures enum value "alerting"
	ConversationStateAlerting string = "alerting"

	// ConversationStateDialing captures enum value "dialing"
	ConversationStateDialing string = "dialing"

	// ConversationStateContacting captures enum value "contacting"
	ConversationStateContacting string = "contacting"

	// ConversationStateOffering captures enum value "offering"
	ConversationStateOffering string = "offering"

	// ConversationStateConnected captures enum value "connected"
	ConversationStateConnected string = "connected"

	// ConversationStateDisconnected captures enum value "disconnected"
	ConversationStateDisconnected string = "disconnected"

	// ConversationStateTerminated captures enum value "terminated"
	ConversationStateTerminated string = "terminated"

	// ConversationStateConverting captures enum value "converting"
	ConversationStateConverting string = "converting"

	// ConversationStateUploading captures enum value "uploading"
	ConversationStateUploading string = "uploading"

	// ConversationStateTransmitting captures enum value "transmitting"
	ConversationStateTransmitting string = "transmitting"

	// ConversationStateNone captures enum value "none"
	ConversationStateNone string = "none"
)

// prop value enum
func (m *Conversation) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, conversationTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Conversation) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this conversation based on the context it is used
func (m *Conversation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDivisions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateParticipants(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecentTransfers(ctx, formats); err != nil {
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

func (m *Conversation) contextValidateDivisions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Divisions); i++ {

		if m.Divisions[i] != nil {
			if err := m.Divisions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("divisions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("divisions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Conversation) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Conversation) contextValidateParticipants(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Participants); i++ {

		if m.Participants[i] != nil {
			if err := m.Participants[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("participants" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("participants" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Conversation) contextValidateRecentTransfers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.RecentTransfers); i++ {

		if m.RecentTransfers[i] != nil {
			if err := m.RecentTransfers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recentTransfers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("recentTransfers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Conversation) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Conversation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Conversation) UnmarshalBinary(b []byte) error {
	var res Conversation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
