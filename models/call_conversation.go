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

// CallConversation call conversation
//
// swagger:model CallConversation
type CallConversation struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// If this is a conference conversation, then this field indicates the maximum number of participants allowed to participant in the conference.
	MaxParticipants int32 `json:"maxParticipants,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The list of other media channels involved in the conversation.
	OtherMediaUris []strfmt.URI `json:"otherMediaUris"`

	// The list of participants involved in the conversation.
	Participants []*CallMediaParticipant `json:"participants"`

	// The list of the most recent 20 transfer commands applied to this conversation.
	RecentTransfers []*TransferResponse `json:"recentTransfers"`

	// recording state
	// Enum: [none active paused]
	RecordingState string `json:"recordingState,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this call conversation
func (m *CallConversation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOtherMediaUris(formats); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CallConversation) validateOtherMediaUris(formats strfmt.Registry) error {
	if swag.IsZero(m.OtherMediaUris) { // not required
		return nil
	}

	for i := 0; i < len(m.OtherMediaUris); i++ {

		if err := validate.FormatOf("otherMediaUris"+"."+strconv.Itoa(i), "body", "uri", m.OtherMediaUris[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

func (m *CallConversation) validateParticipants(formats strfmt.Registry) error {
	if swag.IsZero(m.Participants) { // not required
		return nil
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

func (m *CallConversation) validateRecentTransfers(formats strfmt.Registry) error {
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

var callConversationTypeRecordingStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","active","paused"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		callConversationTypeRecordingStatePropEnum = append(callConversationTypeRecordingStatePropEnum, v)
	}
}

const (

	// CallConversationRecordingStateNone captures enum value "none"
	CallConversationRecordingStateNone string = "none"

	// CallConversationRecordingStateActive captures enum value "active"
	CallConversationRecordingStateActive string = "active"

	// CallConversationRecordingStatePaused captures enum value "paused"
	CallConversationRecordingStatePaused string = "paused"
)

// prop value enum
func (m *CallConversation) validateRecordingStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, callConversationTypeRecordingStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CallConversation) validateRecordingState(formats strfmt.Registry) error {
	if swag.IsZero(m.RecordingState) { // not required
		return nil
	}

	// value enum
	if err := m.validateRecordingStateEnum("recordingState", "body", m.RecordingState); err != nil {
		return err
	}

	return nil
}

func (m *CallConversation) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this call conversation based on the context it is used
func (m *CallConversation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *CallConversation) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CallConversation) contextValidateParticipants(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CallConversation) contextValidateRecentTransfers(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CallConversation) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CallConversation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CallConversation) UnmarshalBinary(b []byte) error {
	var res CallConversation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
