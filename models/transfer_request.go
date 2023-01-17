// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransferRequest transfer request
//
// swagger:model TransferRequest
type TransferRequest struct {

	// The phone number or address of the transfer target.
	Address string `json:"address,omitempty"`

	// The queue ID of the transfer target.
	QueueID string `json:"queueId,omitempty"`

	// The user ID of the transfer target.
	UserID string `json:"userId,omitempty"`

	// The user name of the transfer target.
	UserName string `json:"userName,omitempty"`

	// If true, transfer to the voicemail inbox of the participant that is being replaced.
	Voicemail bool `json:"voicemail"`
}

// Validate validates this transfer request
func (m *TransferRequest) Validate(formats strfmt.Registry) error {
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
