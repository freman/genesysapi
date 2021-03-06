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

// PatchShiftTradeRequest patch shift trade request
//
// swagger:model PatchShiftTradeRequest
type PatchShiftTradeRequest struct {

	// Update the acceptable intervals the initiating user is willing to accept in trade. Setting the enclosed list to empty will make this a one sided trade request
	AcceptableIntervals *ListWrapperInterval `json:"acceptableIntervals,omitempty"`

	// Update the expiration time for this shift trade.
	Expiration *ValueWrapperDate `json:"expiration,omitempty"`

	// Version metadata
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`

	// Update the ID of the receiving user to direct the request at a specific user, or set the wrapped id to null to open up a trade to be matched by any user.
	ReceivingUserID *ValueWrapperString `json:"receivingUserId,omitempty"`
}

// Validate validates this patch shift trade request
func (m *PatchShiftTradeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptableIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivingUserID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PatchShiftTradeRequest) validateAcceptableIntervals(formats strfmt.Registry) error {

	if swag.IsZero(m.AcceptableIntervals) { // not required
		return nil
	}

	if m.AcceptableIntervals != nil {
		if err := m.AcceptableIntervals.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acceptableIntervals")
			}
			return err
		}
	}

	return nil
}

func (m *PatchShiftTradeRequest) validateExpiration(formats strfmt.Registry) error {

	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if m.Expiration != nil {
		if err := m.Expiration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("expiration")
			}
			return err
		}
	}

	return nil
}

func (m *PatchShiftTradeRequest) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
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

func (m *PatchShiftTradeRequest) validateReceivingUserID(formats strfmt.Registry) error {

	if swag.IsZero(m.ReceivingUserID) { // not required
		return nil
	}

	if m.ReceivingUserID != nil {
		if err := m.ReceivingUserID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receivingUserId")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PatchShiftTradeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchShiftTradeRequest) UnmarshalBinary(b []byte) error {
	var res PatchShiftTradeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
