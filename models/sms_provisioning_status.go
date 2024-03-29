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

// SmsProvisioningStatus sms provisioning status
//
// swagger:model SmsProvisioningStatus
type SmsProvisioningStatus struct {

	// Provisioning action
	// Enum: [Unknown Create Update Delete]
	Action string `json:"action,omitempty"`

	// Any error associated with a Failed state
	Error *ErrorBody `json:"error,omitempty"`

	// Provisioning state
	// Enum: [Running Completed Failed]
	State string `json:"state,omitempty"`

	// The phone number version associated with the provisioning action
	Version int64 `json:"version,omitempty"`
}

// Validate validates this sms provisioning status
func (m *SmsProvisioningStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateError(formats); err != nil {
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

var smsProvisioningStatusTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unknown","Create","Update","Delete"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsProvisioningStatusTypeActionPropEnum = append(smsProvisioningStatusTypeActionPropEnum, v)
	}
}

const (

	// SmsProvisioningStatusActionUnknown captures enum value "Unknown"
	SmsProvisioningStatusActionUnknown string = "Unknown"

	// SmsProvisioningStatusActionCreate captures enum value "Create"
	SmsProvisioningStatusActionCreate string = "Create"

	// SmsProvisioningStatusActionUpdate captures enum value "Update"
	SmsProvisioningStatusActionUpdate string = "Update"

	// SmsProvisioningStatusActionDelete captures enum value "Delete"
	SmsProvisioningStatusActionDelete string = "Delete"
)

// prop value enum
func (m *SmsProvisioningStatus) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smsProvisioningStatusTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmsProvisioningStatus) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

func (m *SmsProvisioningStatus) validateError(formats strfmt.Registry) error {
	if swag.IsZero(m.Error) { // not required
		return nil
	}

	if m.Error != nil {
		if err := m.Error.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

var smsProvisioningStatusTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Running","Completed","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsProvisioningStatusTypeStatePropEnum = append(smsProvisioningStatusTypeStatePropEnum, v)
	}
}

const (

	// SmsProvisioningStatusStateRunning captures enum value "Running"
	SmsProvisioningStatusStateRunning string = "Running"

	// SmsProvisioningStatusStateCompleted captures enum value "Completed"
	SmsProvisioningStatusStateCompleted string = "Completed"

	// SmsProvisioningStatusStateFailed captures enum value "Failed"
	SmsProvisioningStatusStateFailed string = "Failed"
)

// prop value enum
func (m *SmsProvisioningStatus) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smsProvisioningStatusTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmsProvisioningStatus) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this sms provisioning status based on the context it is used
func (m *SmsProvisioningStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmsProvisioningStatus) contextValidateError(ctx context.Context, formats strfmt.Registry) error {

	if m.Error != nil {
		if err := m.Error.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("error")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("error")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmsProvisioningStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmsProvisioningStatus) UnmarshalBinary(b []byte) error {
	var res SmsProvisioningStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
