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

// RecordingContentStory Story object.
//
// swagger:model RecordingContentStory
type RecordingContentStory struct {

	// ID of the ephemeral story being replied to.
	ReplyToID string `json:"replyToId,omitempty"`

	// Type of ephemeral story attachment.
	// Enum: [Mention Reply]
	Type string `json:"type,omitempty"`

	// URL to the ephemeral story.
	URL string `json:"url,omitempty"`
}

// Validate validates this recording content story
func (m *RecordingContentStory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var recordingContentStoryTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Mention","Reply"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recordingContentStoryTypeTypePropEnum = append(recordingContentStoryTypeTypePropEnum, v)
	}
}

const (

	// RecordingContentStoryTypeMention captures enum value "Mention"
	RecordingContentStoryTypeMention string = "Mention"

	// RecordingContentStoryTypeReply captures enum value "Reply"
	RecordingContentStoryTypeReply string = "Reply"
)

// prop value enum
func (m *RecordingContentStory) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, recordingContentStoryTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *RecordingContentStory) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this recording content story based on context it is used
func (m *RecordingContentStory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *RecordingContentStory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecordingContentStory) UnmarshalBinary(b []byte) error {
	var res RecordingContentStory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
