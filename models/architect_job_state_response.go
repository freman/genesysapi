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

// ArchitectJobStateResponse architect job state response
//
// swagger:model ArchitectJobStateResponse
type ArchitectJobStateResponse struct {

	// The command executed by the Architect Job
	// Enum: [Publish Create Update]
	Command string `json:"command,omitempty"`

	// Flow created from the Architect Job
	Flow *AddressableEntityRef `json:"flow,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Warnings and Errors messages of the Architect Job
	Messages []*ArchitectJobMessage `json:"messages"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Status of the Architect Job
	// Enum: [Registered Started Success Failure]
	Status string `json:"status,omitempty"`
}

// Validate validates this architect job state response
func (m *ArchitectJobStateResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommand(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFlow(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var architectJobStateResponseTypeCommandPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Publish","Create","Update"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		architectJobStateResponseTypeCommandPropEnum = append(architectJobStateResponseTypeCommandPropEnum, v)
	}
}

const (

	// ArchitectJobStateResponseCommandPublish captures enum value "Publish"
	ArchitectJobStateResponseCommandPublish string = "Publish"

	// ArchitectJobStateResponseCommandCreate captures enum value "Create"
	ArchitectJobStateResponseCommandCreate string = "Create"

	// ArchitectJobStateResponseCommandUpdate captures enum value "Update"
	ArchitectJobStateResponseCommandUpdate string = "Update"
)

// prop value enum
func (m *ArchitectJobStateResponse) validateCommandEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, architectJobStateResponseTypeCommandPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchitectJobStateResponse) validateCommand(formats strfmt.Registry) error {
	if swag.IsZero(m.Command) { // not required
		return nil
	}

	// value enum
	if err := m.validateCommandEnum("command", "body", m.Command); err != nil {
		return err
	}

	return nil
}

func (m *ArchitectJobStateResponse) validateFlow(formats strfmt.Registry) error {
	if swag.IsZero(m.Flow) { // not required
		return nil
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

func (m *ArchitectJobStateResponse) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ArchitectJobStateResponse) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var architectJobStateResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Registered","Started","Success","Failure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		architectJobStateResponseTypeStatusPropEnum = append(architectJobStateResponseTypeStatusPropEnum, v)
	}
}

const (

	// ArchitectJobStateResponseStatusRegistered captures enum value "Registered"
	ArchitectJobStateResponseStatusRegistered string = "Registered"

	// ArchitectJobStateResponseStatusStarted captures enum value "Started"
	ArchitectJobStateResponseStatusStarted string = "Started"

	// ArchitectJobStateResponseStatusSuccess captures enum value "Success"
	ArchitectJobStateResponseStatusSuccess string = "Success"

	// ArchitectJobStateResponseStatusFailure captures enum value "Failure"
	ArchitectJobStateResponseStatusFailure string = "Failure"
)

// prop value enum
func (m *ArchitectJobStateResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, architectJobStateResponseTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ArchitectJobStateResponse) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this architect job state response based on the context it is used
func (m *ArchitectJobStateResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFlow(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
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

func (m *ArchitectJobStateResponse) contextValidateFlow(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ArchitectJobStateResponse) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ArchitectJobStateResponse) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Messages); i++ {

		if m.Messages[i] != nil {
			if err := m.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ArchitectJobStateResponse) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArchitectJobStateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArchitectJobStateResponse) UnmarshalBinary(b []byte) error {
	var res ArchitectJobStateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
