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

// Topic topic
//
// swagger:model Topic
type Topic struct {

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DatePublished strfmt.DateTime `json:"datePublished,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// dialect
	Dialect string `json:"dialect,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// modified by
	ModifiedBy *AddressableEntityRef `json:"modifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// participants
	// Enum: [External Internal All]
	Participants string `json:"participants,omitempty"`

	// phrases
	Phrases []*Phrase `json:"phrases"`

	// programs
	Programs []*BaseProgramEntity `json:"programs"`

	// published
	Published bool `json:"published"`

	// published by
	PublishedBy *AddressableEntityRef `json:"publishedBy,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// strictness
	// Enum: [1 55 65 72 85 90]
	Strictness string `json:"strictness,omitempty"`

	// tags
	Tags []string `json:"tags"`
}

// Validate validates this topic
func (m *Topic) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatePublished(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipants(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhrases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrograms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublishedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStrictness(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Topic) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Topic) validateDatePublished(formats strfmt.Registry) error {
	if swag.IsZero(m.DatePublished) { // not required
		return nil
	}

	if err := validate.FormatOf("datePublished", "body", "date-time", m.DatePublished.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Topic) validateModifiedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

var topicTypeParticipantsPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["External","Internal","All"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		topicTypeParticipantsPropEnum = append(topicTypeParticipantsPropEnum, v)
	}
}

const (

	// TopicParticipantsExternal captures enum value "External"
	TopicParticipantsExternal string = "External"

	// TopicParticipantsInternal captures enum value "Internal"
	TopicParticipantsInternal string = "Internal"

	// TopicParticipantsAll captures enum value "All"
	TopicParticipantsAll string = "All"
)

// prop value enum
func (m *Topic) validateParticipantsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, topicTypeParticipantsPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Topic) validateParticipants(formats strfmt.Registry) error {
	if swag.IsZero(m.Participants) { // not required
		return nil
	}

	// value enum
	if err := m.validateParticipantsEnum("participants", "body", m.Participants); err != nil {
		return err
	}

	return nil
}

func (m *Topic) validatePhrases(formats strfmt.Registry) error {
	if swag.IsZero(m.Phrases) { // not required
		return nil
	}

	for i := 0; i < len(m.Phrases); i++ {
		if swag.IsZero(m.Phrases[i]) { // not required
			continue
		}

		if m.Phrases[i] != nil {
			if err := m.Phrases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phrases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phrases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Topic) validatePrograms(formats strfmt.Registry) error {
	if swag.IsZero(m.Programs) { // not required
		return nil
	}

	for i := 0; i < len(m.Programs); i++ {
		if swag.IsZero(m.Programs[i]) { // not required
			continue
		}

		if m.Programs[i] != nil {
			if err := m.Programs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("programs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("programs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Topic) validatePublishedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.PublishedBy) { // not required
		return nil
	}

	if m.PublishedBy != nil {
		if err := m.PublishedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publishedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Topic) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var topicTypeStrictnessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","55","65","72","85","90"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		topicTypeStrictnessPropEnum = append(topicTypeStrictnessPropEnum, v)
	}
}

const (

	// TopicStrictnessNr1 captures enum value "1"
	TopicStrictnessNr1 string = "1"

	// TopicStrictnessNr55 captures enum value "55"
	TopicStrictnessNr55 string = "55"

	// TopicStrictnessNr65 captures enum value "65"
	TopicStrictnessNr65 string = "65"

	// TopicStrictnessNr72 captures enum value "72"
	TopicStrictnessNr72 string = "72"

	// TopicStrictnessNr85 captures enum value "85"
	TopicStrictnessNr85 string = "85"

	// TopicStrictnessNr90 captures enum value "90"
	TopicStrictnessNr90 string = "90"
)

// prop value enum
func (m *Topic) validateStrictnessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, topicTypeStrictnessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Topic) validateStrictness(formats strfmt.Registry) error {
	if swag.IsZero(m.Strictness) { // not required
		return nil
	}

	// value enum
	if err := m.validateStrictnessEnum("strictness", "body", m.Strictness); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this topic based on the context it is used
func (m *Topic) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhrases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrograms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePublishedBy(ctx, formats); err != nil {
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

func (m *Topic) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Topic) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Topic) contextValidatePhrases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Phrases); i++ {

		if m.Phrases[i] != nil {
			if err := m.Phrases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phrases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phrases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Topic) contextValidatePrograms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Programs); i++ {

		if m.Programs[i] != nil {
			if err := m.Programs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("programs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("programs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Topic) contextValidatePublishedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.PublishedBy != nil {
		if err := m.PublishedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("publishedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("publishedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Topic) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Topic) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Topic) UnmarshalBinary(b []byte) error {
	var res Topic
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
