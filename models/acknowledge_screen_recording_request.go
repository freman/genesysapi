// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AcknowledgeScreenRecordingRequest acknowledge screen recording request
//
// swagger:model AcknowledgeScreenRecordingRequest
type AcknowledgeScreenRecordingRequest struct {

	// conversation Id
	ConversationID string `json:"conversationId,omitempty"`

	// participant jid
	ParticipantJid string `json:"participantJid,omitempty"`

	// room Id
	RoomID string `json:"roomId,omitempty"`
}

// Validate validates this acknowledge screen recording request
func (m *AcknowledgeScreenRecordingRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this acknowledge screen recording request based on context it is used
func (m *AcknowledgeScreenRecordingRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AcknowledgeScreenRecordingRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AcknowledgeScreenRecordingRequest) UnmarshalBinary(b []byte) error {
	var res AcknowledgeScreenRecordingRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
