// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WfmAgent Workforce management agent data
//
// swagger:model WfmAgent
type WfmAgent struct {

	// Whether the agent accepts direct shift trade requests
	AcceptDirectShiftTrades bool `json:"acceptDirectShiftTrades"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The list of languages
	Languages []*LanguageReference `json:"languages"`

	// Metadata for this agent
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// List of queues to which the agent belongs and which are defined in the service goal groups in this management unit
	Queues []*QueueReference `json:"queues"`

	// Whether the agent has the permission to be included in schedule generation
	Schedulable bool `json:"schedulable"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The list of skills
	Skills []*RoutingSkillReference `json:"skills"`

	// The time zone for this agent. Defaults to the time zone of the management unit to which the agent belongs
	TimeZone *WfmTimeZone `json:"timeZone,omitempty"`

	// The user associated with this data
	User *UserReference `json:"user,omitempty"`

	// The work plan associated with this agent
	WorkPlan *WorkPlanReference `json:"workPlan,omitempty"`
}

// Validate validates this wfm agent
func (m *WfmAgent) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQueues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkills(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkPlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmAgent) validateLanguages(formats strfmt.Registry) error {

	if swag.IsZero(m.Languages) { // not required
		return nil
	}

	for i := 0; i < len(m.Languages); i++ {
		if swag.IsZero(m.Languages[i]) { // not required
			continue
		}

		if m.Languages[i] != nil {
			if err := m.Languages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WfmAgent) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *WfmAgent) validateQueues(formats strfmt.Registry) error {

	if swag.IsZero(m.Queues) { // not required
		return nil
	}

	for i := 0; i < len(m.Queues); i++ {
		if swag.IsZero(m.Queues[i]) { // not required
			continue
		}

		if m.Queues[i] != nil {
			if err := m.Queues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("queues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WfmAgent) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WfmAgent) validateSkills(formats strfmt.Registry) error {

	if swag.IsZero(m.Skills) { // not required
		return nil
	}

	for i := 0; i < len(m.Skills); i++ {
		if swag.IsZero(m.Skills[i]) { // not required
			continue
		}

		if m.Skills[i] != nil {
			if err := m.Skills[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("skills" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WfmAgent) validateTimeZone(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeZone) { // not required
		return nil
	}

	if m.TimeZone != nil {
		if err := m.TimeZone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeZone")
			}
			return err
		}
	}

	return nil
}

func (m *WfmAgent) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *WfmAgent) validateWorkPlan(formats strfmt.Registry) error {

	if swag.IsZero(m.WorkPlan) { // not required
		return nil
	}

	if m.WorkPlan != nil {
		if err := m.WorkPlan.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workPlan")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WfmAgent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmAgent) UnmarshalBinary(b []byte) error {
	var res WfmAgent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}