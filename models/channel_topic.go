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

// ChannelTopic channel topic
//
// swagger:model ChannelTopic
type ChannelTopic struct {

	// id
	ID string `json:"id,omitempty"`

	// rejection reason
	RejectionReason string `json:"rejectionReason,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// state
	// Enum: [Permitted Rejected]
	State string `json:"state,omitempty"`
}

// Validate validates this channel topic
func (m *ChannelTopic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *ChannelTopic) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var channelTopicTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Permitted","Rejected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		channelTopicTypeStatePropEnum = append(channelTopicTypeStatePropEnum, v)
	}
}

const (

	// ChannelTopicStatePermitted captures enum value "Permitted"
	ChannelTopicStatePermitted string = "Permitted"

	// ChannelTopicStateRejected captures enum value "Rejected"
	ChannelTopicStateRejected string = "Rejected"
)

// prop value enum
func (m *ChannelTopic) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, channelTopicTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ChannelTopic) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this channel topic based on the context it is used
func (m *ChannelTopic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChannelTopic) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChannelTopic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChannelTopic) UnmarshalBinary(b []byte) error {
	var res ChannelTopic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
