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

// ChatMediaPolicyConditions chat media policy conditions
//
// swagger:model ChatMediaPolicyConditions
type ChatMediaPolicyConditions struct {

	// date ranges
	DateRanges []string `json:"dateRanges"`

	// duration
	Duration *DurationCondition `json:"duration,omitempty"`

	// for queues
	ForQueues []*Queue `json:"forQueues"`

	// for users
	ForUsers []*User `json:"forUsers"`

	// languages
	Languages []*Language `json:"languages"`

	// time allowed
	TimeAllowed *TimeAllowed `json:"timeAllowed,omitempty"`

	// wrapup codes
	WrapupCodes []*WrapupCode `json:"wrapupCodes"`
}

// Validate validates this chat media policy conditions
func (m *ChatMediaPolicyConditions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForQueues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateForUsers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLanguages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeAllowed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWrapupCodes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ChatMediaPolicyConditions) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if m.Duration != nil {
		if err := m.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duration")
			}
			return err
		}
	}

	return nil
}

func (m *ChatMediaPolicyConditions) validateForQueues(formats strfmt.Registry) error {

	if swag.IsZero(m.ForQueues) { // not required
		return nil
	}

	for i := 0; i < len(m.ForQueues); i++ {
		if swag.IsZero(m.ForQueues[i]) { // not required
			continue
		}

		if m.ForQueues[i] != nil {
			if err := m.ForQueues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forQueues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChatMediaPolicyConditions) validateForUsers(formats strfmt.Registry) error {

	if swag.IsZero(m.ForUsers) { // not required
		return nil
	}

	for i := 0; i < len(m.ForUsers); i++ {
		if swag.IsZero(m.ForUsers[i]) { // not required
			continue
		}

		if m.ForUsers[i] != nil {
			if err := m.ForUsers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ChatMediaPolicyConditions) validateLanguages(formats strfmt.Registry) error {

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

func (m *ChatMediaPolicyConditions) validateTimeAllowed(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeAllowed) { // not required
		return nil
	}

	if m.TimeAllowed != nil {
		if err := m.TimeAllowed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeAllowed")
			}
			return err
		}
	}

	return nil
}

func (m *ChatMediaPolicyConditions) validateWrapupCodes(formats strfmt.Registry) error {

	if swag.IsZero(m.WrapupCodes) { // not required
		return nil
	}

	for i := 0; i < len(m.WrapupCodes); i++ {
		if swag.IsZero(m.WrapupCodes[i]) { // not required
			continue
		}

		if m.WrapupCodes[i] != nil {
			if err := m.WrapupCodes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wrapupCodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ChatMediaPolicyConditions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ChatMediaPolicyConditions) UnmarshalBinary(b []byte) error {
	var res ChatMediaPolicyConditions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}