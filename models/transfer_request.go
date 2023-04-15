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

// TransferRequest transfer request
//
// swagger:model TransferRequest
type TransferRequest struct {

	// The user ID or queue ID of the transfer target. Address like a phone number can not be used for callbacks, but they can be used for other forms of communication.
	Address string `json:"address,omitempty"`

	// The queue ID of the transfer target.
	QueueID string `json:"queueId,omitempty"`

	// The type of transfer to perform.
	// Enum: [Attended Unattended]
	TransferType string `json:"transferType,omitempty"`

	// The user ID of the transfer target.
	UserID string `json:"userId,omitempty"`

	// The user name of the transfer target.
	UserName string `json:"userName,omitempty"`

	// If true, transfer to the voicemail inbox of the participant that is being replaced.
	Voicemail bool `json:"voicemail"`
}

// Validate validates this transfer request
func (m *TransferRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransferType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var transferRequestTypeTransferTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Attended","Unattended"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transferRequestTypeTransferTypePropEnum = append(transferRequestTypeTransferTypePropEnum, v)
	}
}

const (

	// TransferRequestTransferTypeAttended captures enum value "Attended"
	TransferRequestTransferTypeAttended string = "Attended"

	// TransferRequestTransferTypeUnattended captures enum value "Unattended"
	TransferRequestTransferTypeUnattended string = "Unattended"
)

// prop value enum
func (m *TransferRequest) validateTransferTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, transferRequestTypeTransferTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TransferRequest) validateTransferType(formats strfmt.Registry) error {
	if swag.IsZero(m.TransferType) { // not required
		return nil
	}

	// value enum
	if err := m.validateTransferTypeEnum("transferType", "body", m.TransferType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this transfer request based on context it is used
func (m *TransferRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TransferRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransferRequest) UnmarshalBinary(b []byte) error {
	var res TransferRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
