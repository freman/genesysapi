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

// EmailMediaPolicyConditions email media policy conditions
//
// swagger:model EmailMediaPolicyConditions
type EmailMediaPolicyConditions struct {

	// customer participation
	// Enum: [YES NO]
	CustomerParticipation string `json:"customerParticipation,omitempty"`

	// date ranges
	DateRanges []string `json:"dateRanges"`

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

// Validate validates this email media policy conditions
func (m *EmailMediaPolicyConditions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomerParticipation(formats); err != nil {
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

var emailMediaPolicyConditionsTypeCustomerParticipationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["YES","NO"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emailMediaPolicyConditionsTypeCustomerParticipationPropEnum = append(emailMediaPolicyConditionsTypeCustomerParticipationPropEnum, v)
	}
}

const (

	// EmailMediaPolicyConditionsCustomerParticipationYES captures enum value "YES"
	EmailMediaPolicyConditionsCustomerParticipationYES string = "YES"

	// EmailMediaPolicyConditionsCustomerParticipationNO captures enum value "NO"
	EmailMediaPolicyConditionsCustomerParticipationNO string = "NO"
)

// prop value enum
func (m *EmailMediaPolicyConditions) validateCustomerParticipationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, emailMediaPolicyConditionsTypeCustomerParticipationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EmailMediaPolicyConditions) validateCustomerParticipation(formats strfmt.Registry) error {
	if swag.IsZero(m.CustomerParticipation) { // not required
		return nil
	}

	// value enum
	if err := m.validateCustomerParticipationEnum("customerParticipation", "body", m.CustomerParticipation); err != nil {
		return err
	}

	return nil
}

func (m *EmailMediaPolicyConditions) validateForQueues(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forQueues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMediaPolicyConditions) validateForUsers(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMediaPolicyConditions) validateLanguages(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMediaPolicyConditions) validateTimeAllowed(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeAllowed) { // not required
		return nil
	}

	if m.TimeAllowed != nil {
		if err := m.TimeAllowed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeAllowed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeAllowed")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMediaPolicyConditions) validateWrapupCodes(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wrapupCodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this email media policy conditions based on the context it is used
func (m *EmailMediaPolicyConditions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateForQueues(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateForUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLanguages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeAllowed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWrapupCodes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailMediaPolicyConditions) contextValidateForQueues(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ForQueues); i++ {

		if m.ForQueues[i] != nil {
			if err := m.ForQueues[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forQueues" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forQueues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMediaPolicyConditions) contextValidateForUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ForUsers); i++ {

		if m.ForUsers[i] != nil {
			if err := m.ForUsers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("forUsers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("forUsers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMediaPolicyConditions) contextValidateLanguages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Languages); i++ {

		if m.Languages[i] != nil {
			if err := m.Languages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("languages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("languages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EmailMediaPolicyConditions) contextValidateTimeAllowed(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeAllowed != nil {
		if err := m.TimeAllowed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeAllowed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeAllowed")
			}
			return err
		}
	}

	return nil
}

func (m *EmailMediaPolicyConditions) contextValidateWrapupCodes(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.WrapupCodes); i++ {

		if m.WrapupCodes[i] != nil {
			if err := m.WrapupCodes[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("wrapupCodes" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("wrapupCodes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailMediaPolicyConditions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailMediaPolicyConditions) UnmarshalBinary(b []byte) error {
	var res EmailMediaPolicyConditions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
