// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BatchDownloadRequest batch download request
//
// swagger:model BatchDownloadRequest
type BatchDownloadRequest struct {

	// Conversation id requested
	ConversationID string `json:"conversationId,omitempty"`

	// Recording id requested, optional.  Leave null for all recordings on the conversation
	RecordingID string `json:"recordingId,omitempty"`
}

// Validate validates this batch download request
func (m *BatchDownloadRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this batch download request based on context it is used
func (m *BatchDownloadRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BatchDownloadRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchDownloadRequest) UnmarshalBinary(b []byte) error {
	var res BatchDownloadRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
