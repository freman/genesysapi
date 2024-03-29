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

// SecureSession secure session
//
// swagger:model SecureSession
type SecureSession struct {

	// If true, disconnect the agent after creating the session
	Disconnect bool `json:"disconnect"`

	// The flow to execute securely
	// Required: true
	Flow *DomainEntityRef `json:"flow"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Unique identifier for the participant initiating the secure session.
	SourceParticipantID string `json:"sourceParticipantId,omitempty"`

	// The current state of a secure session
	// Required: true
	// Enum: [PENDING COMPLETED FAILED]
	State *string `json:"state"`

	// Customer-provided data
	UserData string `json:"userData,omitempty"`
}

// Validate validates this secure session
func (m *SecureSession) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

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

func (m *SecureSession) validateFlow(formats strfmt.Registry) error {

	if err := validate.Required("flow", "body", m.Flow); err != nil {
		return err
	}

	if m.Flow != nil {
		if err := m.Flow.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *SecureSession) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var secureSessionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","COMPLETED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		secureSessionTypeStatePropEnum = append(secureSessionTypeStatePropEnum, v)
	}
}

const (

	// SecureSessionStatePENDING captures enum value "PENDING"
	SecureSessionStatePENDING string = "PENDING"

	// SecureSessionStateCOMPLETED captures enum value "COMPLETED"
	SecureSessionStateCOMPLETED string = "COMPLETED"

	// SecureSessionStateFAILED captures enum value "FAILED"
	SecureSessionStateFAILED string = "FAILED"
)

// prop value enum
func (m *SecureSession) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, secureSessionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SecureSession) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this secure session based on the context it is used
func (m *SecureSession) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
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

func (m *SecureSession) contextValidateFlow(ctx context.Context, formats strfmt.Registry) error {

	if m.Flow != nil {
		if err := m.Flow.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flow")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flow")
			}
			return err
		}
	}

	return nil
}

func (m *SecureSession) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *SecureSession) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SecureSession) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SecureSession) UnmarshalBinary(b []byte) error {
	var res SecureSession
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
