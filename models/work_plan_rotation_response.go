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

// WorkPlanRotationResponse work plan rotation response
//
// swagger:model WorkPlanRotationResponse
type WorkPlanRotationResponse struct {

	// Number of agents in this work plan rotation
	AgentCount int32 `json:"agentCount,omitempty"`

	// Agents in this work plan rotation. Populate with expand=agents for GET WorkPlanRotationsList (defaults to empty list)
	Agents []*WorkPlanRotationAgentResponse `json:"agents"`

	// The date range to which this work plan rotation applies
	DateRange *DateRangeWithOptionalEnd `json:"dateRange,omitempty"`

	// Whether the work plan rotation is enabled for scheduling
	Enabled bool `json:"enabled"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Version metadata for this work plan rotation
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// Pattern with ordered list of work plans that rotate on a weekly basis
	Pattern *WorkPlanPatternResponse `json:"pattern,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this work plan rotation response
func (m *WorkPlanRotationResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgents(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePattern(formats); err != nil {
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

func (m *WorkPlanRotationResponse) validateAgents(formats strfmt.Registry) error {
	if swag.IsZero(m.Agents) { // not required
		return nil
	}

	for i := 0; i < len(m.Agents); i++ {
		if swag.IsZero(m.Agents[i]) { // not required
			continue
		}

		if m.Agents[i] != nil {
			if err := m.Agents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkPlanRotationResponse) validateDateRange(formats strfmt.Registry) error {
	if swag.IsZero(m.DateRange) { // not required
		return nil
	}

	if m.DateRange != nil {
		if err := m.DateRange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dateRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dateRange")
			}
			return err
		}
	}

	return nil
}

func (m *WorkPlanRotationResponse) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *WorkPlanRotationResponse) validatePattern(formats strfmt.Registry) error {
	if swag.IsZero(m.Pattern) { // not required
		return nil
	}

	if m.Pattern != nil {
		if err := m.Pattern.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pattern")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pattern")
			}
			return err
		}
	}

	return nil
}

func (m *WorkPlanRotationResponse) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this work plan rotation response based on the context it is used
func (m *WorkPlanRotationResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateRange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePattern(ctx, formats); err != nil {
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

func (m *WorkPlanRotationResponse) contextValidateAgents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Agents); i++ {

		if m.Agents[i] != nil {
			if err := m.Agents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("agents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("agents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WorkPlanRotationResponse) contextValidateDateRange(ctx context.Context, formats strfmt.Registry) error {

	if m.DateRange != nil {
		if err := m.DateRange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("dateRange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("dateRange")
			}
			return err
		}
	}

	return nil
}

func (m *WorkPlanRotationResponse) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WorkPlanRotationResponse) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *WorkPlanRotationResponse) contextValidatePattern(ctx context.Context, formats strfmt.Registry) error {

	if m.Pattern != nil {
		if err := m.Pattern.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("pattern")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("pattern")
			}
			return err
		}
	}

	return nil
}

func (m *WorkPlanRotationResponse) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkPlanRotationResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkPlanRotationResponse) UnmarshalBinary(b []byte) error {
	var res WorkPlanRotationResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
