// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TranscriptTopic transcript topic
//
// swagger:model TranscriptTopic
type TranscriptTopic struct {

	// The detection confidence of the topic.
	// Read Only: true
	Confidence int32 `json:"confidence,omitempty"`

	// duration
	Duration *TopicDuration `json:"duration,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the object.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The start time of the topic phrase.
	// Read Only: true
	StartTimeMilliseconds int64 `json:"startTimeMilliseconds,omitempty"`

	// The phrase which detected the topic.
	// Read Only: true
	TopicPhrase string `json:"topicPhrase,omitempty"`

	// The transcript phrase which detected the topic.
	// Read Only: true
	TranscriptPhrase string `json:"transcriptPhrase,omitempty"`
}

// Validate validates this transcript topic
func (m *TranscriptTopic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TranscriptTopic) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if m.Duration != nil {
		if err := m.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TranscriptTopic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TranscriptTopic) UnmarshalBinary(b []byte) error {
	var res TranscriptTopic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}