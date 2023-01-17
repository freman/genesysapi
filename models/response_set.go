// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResponseSet response set
//
// swagger:model ResponseSet
type ResponseSet struct {

	// Whether to enable answering machine detection
	AmdSpeechDistinguishEnabled bool `json:"amdSpeechDistinguishEnabled"`

	// Whether to enable answering machine beep detection
	BeepDetectionEnabled bool `json:"beepDetectionEnabled"`

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The name of the ResponseSet.
	// Required: true
	Name *string `json:"name"`

	// Map of disposition identifiers to reactions. For example: {"disposition.classification.callable.person": {"reactionType": "transfer"}}.
	// Required: true
	Responses map[string]Reaction `json:"responses"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this response set
func (m *ResponseSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponses(formats); err != nil {
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

func (m *ResponseSet) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponseSet) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ResponseSet) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ResponseSet) validateResponses(formats strfmt.Registry) error {

	if err := validate.Required("responses", "body", m.Responses); err != nil {
		return err
	}

	for k := range m.Responses {

		if err := validate.Required("responses"+"."+k, "body", m.Responses[k]); err != nil {
			return err
		}
		if val, ok := m.Responses[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("responses" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("responses" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResponseSet) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this response set based on the context it is used
func (m *ResponseSet) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateResponses(ctx, formats); err != nil {
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

func (m *ResponseSet) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *ResponseSet) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *ResponseSet) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ResponseSet) contextValidateResponses(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.Required("responses", "body", m.Responses); err != nil {
		return err
	}

	for k := range m.Responses {

		if val, ok := m.Responses[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *ResponseSet) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResponseSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseSet) UnmarshalBinary(b []byte) error {
	var res ResponseSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
