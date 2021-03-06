// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AnalyticsResolution analytics resolution
//
// swagger:model AnalyticsResolution
type AnalyticsResolution struct {

	// The number of interactions for which next contact was avoided.
	GetnNextContactAvoided int64 `json:"getnNextContactAvoided,omitempty"`

	// The ID of the last queue on which the conversation was handled.
	QueueID string `json:"queueId,omitempty"`

	// The ID of the last user who handled the conversation.
	UserID string `json:"userId,omitempty"`
}

// Validate validates this analytics resolution
func (m *AnalyticsResolution) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AnalyticsResolution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AnalyticsResolution) UnmarshalBinary(b []byte) error {
	var res AnalyticsResolution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
