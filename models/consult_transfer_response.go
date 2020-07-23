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

// ConsultTransferResponse consult transfer response
//
// swagger:model ConsultTransferResponse
type ConsultTransferResponse struct {

	// Participant ID to whom the call is being transferred.
	// Required: true
	DestinationParticipantID *string `json:"destinationParticipantId"`
}

// Validate validates this consult transfer response
func (m *ConsultTransferResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDestinationParticipantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsultTransferResponse) validateDestinationParticipantID(formats strfmt.Registry) error {

	if err := validate.Required("destinationParticipantId", "body", m.DestinationParticipantID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsultTransferResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsultTransferResponse) UnmarshalBinary(b []byte) error {
	var res ConsultTransferResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
