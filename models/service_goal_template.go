// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ServiceGoalTemplate Service Goal Template
//
// swagger:model ServiceGoalTemplate
type ServiceGoalTemplate struct {

	// Abandon rate targets for this service goal template
	AbandonRate *BuAbandonRate `json:"abandonRate,omitempty"`

	// Average speed of answer targets for this service goal template
	AverageSpeedOfAnswer *BuAverageSpeedOfAnswer `json:"averageSpeedOfAnswer,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Version metadata for the service goal template
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Service level targets for this service goal template
	ServiceLevel *BuServiceLevel `json:"serviceLevel,omitempty"`
}

// Validate validates this service goal template
func (m *ServiceGoalTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAbandonRate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAverageSpeedOfAnswer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceGoalTemplate) validateAbandonRate(formats strfmt.Registry) error {

	if swag.IsZero(m.AbandonRate) { // not required
		return nil
	}

	if m.AbandonRate != nil {
		if err := m.AbandonRate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("abandonRate")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceGoalTemplate) validateAverageSpeedOfAnswer(formats strfmt.Registry) error {

	if swag.IsZero(m.AverageSpeedOfAnswer) { // not required
		return nil
	}

	if m.AverageSpeedOfAnswer != nil {
		if err := m.AverageSpeedOfAnswer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("averageSpeedOfAnswer")
			}
			return err
		}
	}

	return nil
}

func (m *ServiceGoalTemplate) validateMetadata(formats strfmt.Registry) error {

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

func (m *ServiceGoalTemplate) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ServiceGoalTemplate) validateServiceLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceLevel) { // not required
		return nil
	}

	if m.ServiceLevel != nil {
		if err := m.ServiceLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceLevel")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceGoalTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceGoalTemplate) UnmarshalBinary(b []byte) error {
	var res ServiceGoalTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
