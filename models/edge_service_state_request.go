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

// EdgeServiceStateRequest edge service state request
//
// swagger:model EdgeServiceStateRequest
type EdgeServiceStateRequest struct {

	// The number of seconds to wait for call draining to complete before initiating the reboot. A value of 0 will prevent call draining and all calls will disconnect immediately.
	CallDrainingWaitTimeSeconds int32 `json:"callDrainingWaitTimeSeconds,omitempty"`

	// A boolean that sets the Edge in-service or out-of-service.
	// Required: true
	InService *bool `json:"inService"`
}

// Validate validates this edge service state request
func (m *EdgeServiceStateRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInService(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeServiceStateRequest) validateInService(formats strfmt.Registry) error {

	if err := validate.Required("inService", "body", m.InService); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeServiceStateRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeServiceStateRequest) UnmarshalBinary(b []byte) error {
	var res EdgeServiceStateRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
