// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CrossPlatformPolicyActions cross platform policy actions
//
// swagger:model CrossPlatformPolicyActions
type CrossPlatformPolicyActions struct {

	// true to delete the recording associated with the conversation regardless of the values of retainRecording or deleteRecording. Default = false
	AlwaysDelete bool `json:"alwaysDelete"`

	// assign calibrations
	AssignCalibrations []*CalibrationAssignment `json:"assignCalibrations"`

	// assign evaluations
	AssignEvaluations []*EvaluationAssignment `json:"assignEvaluations"`

	// assign metered assignment by agent
	AssignMeteredAssignmentByAgent []*MeteredAssignmentByAgent `json:"assignMeteredAssignmentByAgent"`

	// assign metered evaluations
	AssignMeteredEvaluations []*MeteredEvaluationAssignment `json:"assignMeteredEvaluations"`

	// true to delete the recording associated with the conversation. If retainRecording = true, this will be ignored. Default = false
	DeleteRecording bool `json:"deleteRecording"`

	// Policy action for exporting recordings using an integration to 3rd party s3.
	IntegrationExport *IntegrationExport `json:"integrationExport,omitempty"`

	// media transcriptions
	MediaTranscriptions []*MediaTranscription `json:"mediaTranscriptions"`

	// true to retain the recording associated with the conversation. Default = true
	RetainRecording bool `json:"retainRecording"`

	// retention duration
	RetentionDuration *RetentionDuration `json:"retentionDuration,omitempty"`
}

// Validate validates this cross platform policy actions
func (m *CrossPlatformPolicyActions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignCalibrations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignEvaluations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignMeteredAssignmentByAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignMeteredEvaluations(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntegrationExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaTranscriptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRetentionDuration(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CrossPlatformPolicyActions) validateAssignCalibrations(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignCalibrations) { // not required
		return nil
	}

	for i := 0; i < len(m.AssignCalibrations); i++ {
		if swag.IsZero(m.AssignCalibrations[i]) { // not required
			continue
		}

		if m.AssignCalibrations[i] != nil {
			if err := m.AssignCalibrations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignCalibrations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CrossPlatformPolicyActions) validateAssignEvaluations(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignEvaluations) { // not required
		return nil
	}

	for i := 0; i < len(m.AssignEvaluations); i++ {
		if swag.IsZero(m.AssignEvaluations[i]) { // not required
			continue
		}

		if m.AssignEvaluations[i] != nil {
			if err := m.AssignEvaluations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignEvaluations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CrossPlatformPolicyActions) validateAssignMeteredAssignmentByAgent(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignMeteredAssignmentByAgent) { // not required
		return nil
	}

	for i := 0; i < len(m.AssignMeteredAssignmentByAgent); i++ {
		if swag.IsZero(m.AssignMeteredAssignmentByAgent[i]) { // not required
			continue
		}

		if m.AssignMeteredAssignmentByAgent[i] != nil {
			if err := m.AssignMeteredAssignmentByAgent[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignMeteredAssignmentByAgent" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CrossPlatformPolicyActions) validateAssignMeteredEvaluations(formats strfmt.Registry) error {

	if swag.IsZero(m.AssignMeteredEvaluations) { // not required
		return nil
	}

	for i := 0; i < len(m.AssignMeteredEvaluations); i++ {
		if swag.IsZero(m.AssignMeteredEvaluations[i]) { // not required
			continue
		}

		if m.AssignMeteredEvaluations[i] != nil {
			if err := m.AssignMeteredEvaluations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignMeteredEvaluations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CrossPlatformPolicyActions) validateIntegrationExport(formats strfmt.Registry) error {

	if swag.IsZero(m.IntegrationExport) { // not required
		return nil
	}

	if m.IntegrationExport != nil {
		if err := m.IntegrationExport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("integrationExport")
			}
			return err
		}
	}

	return nil
}

func (m *CrossPlatformPolicyActions) validateMediaTranscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.MediaTranscriptions) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaTranscriptions); i++ {
		if swag.IsZero(m.MediaTranscriptions[i]) { // not required
			continue
		}

		if m.MediaTranscriptions[i] != nil {
			if err := m.MediaTranscriptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("mediaTranscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CrossPlatformPolicyActions) validateRetentionDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.RetentionDuration) { // not required
		return nil
	}

	if m.RetentionDuration != nil {
		if err := m.RetentionDuration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("retentionDuration")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CrossPlatformPolicyActions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CrossPlatformPolicyActions) UnmarshalBinary(b []byte) error {
	var res CrossPlatformPolicyActions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
