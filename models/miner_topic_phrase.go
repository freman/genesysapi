// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MinerTopicPhrase miner topic phrase
//
// swagger:model MinerTopicPhrase
type MinerTopicPhrase struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Phrase name.
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Topic associated with a phrase.
	Topic *MinerTopic `json:"topic,omitempty"`

	// Number of utterances belonging to a phrase
	UtteranceCount int32 `json:"utteranceCount,omitempty"`

	// List of utterances related to a phrase.
	Utterances []*Utterance `json:"utterances"`
}

// Validate validates this miner topic phrase
func (m *MinerTopicPhrase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTopic(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUtterances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MinerTopicPhrase) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MinerTopicPhrase) validateTopic(formats strfmt.Registry) error {
	if swag.IsZero(m.Topic) { // not required
		return nil
	}

	if m.Topic != nil {
		if err := m.Topic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topic")
			}
			return err
		}
	}

	return nil
}

func (m *MinerTopicPhrase) validateUtterances(formats strfmt.Registry) error {
	if swag.IsZero(m.Utterances) { // not required
		return nil
	}

	for i := 0; i < len(m.Utterances); i++ {
		if swag.IsZero(m.Utterances[i]) { // not required
			continue
		}

		if m.Utterances[i] != nil {
			if err := m.Utterances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("utterances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("utterances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this miner topic phrase based on the context it is used
func (m *MinerTopicPhrase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTopic(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUtterances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MinerTopicPhrase) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *MinerTopicPhrase) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *MinerTopicPhrase) contextValidateTopic(ctx context.Context, formats strfmt.Registry) error {

	if m.Topic != nil {
		if err := m.Topic.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("topic")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("topic")
			}
			return err
		}
	}

	return nil
}

func (m *MinerTopicPhrase) contextValidateUtterances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Utterances); i++ {

		if m.Utterances[i] != nil {
			if err := m.Utterances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("utterances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("utterances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *MinerTopicPhrase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MinerTopicPhrase) UnmarshalBinary(b []byte) error {
	var res MinerTopicPhrase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
