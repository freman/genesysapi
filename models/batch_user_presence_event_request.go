// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BatchUserPresenceEventRequest A maximum of 100 events are allowed per request
//
// swagger:model BatchUserPresenceEventRequest
type BatchUserPresenceEventRequest struct {

	// UserPresence events for this batch
	UserPresenceEvents []*UserPresenceEvent `json:"userPresenceEvents"`
}

// Validate validates this batch user presence event request
func (m *BatchUserPresenceEventRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserPresenceEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchUserPresenceEventRequest) validateUserPresenceEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.UserPresenceEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.UserPresenceEvents); i++ {
		if swag.IsZero(m.UserPresenceEvents[i]) { // not required
			continue
		}

		if m.UserPresenceEvents[i] != nil {
			if err := m.UserPresenceEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userPresenceEvents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userPresenceEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this batch user presence event request based on the context it is used
func (m *BatchUserPresenceEventRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserPresenceEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchUserPresenceEventRequest) contextValidateUserPresenceEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserPresenceEvents); i++ {

		if m.UserPresenceEvents[i] != nil {
			if err := m.UserPresenceEvents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userPresenceEvents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userPresenceEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchUserPresenceEventRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchUserPresenceEventRequest) UnmarshalBinary(b []byte) error {
	var res BatchUserPresenceEventRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
