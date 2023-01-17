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

// BatchUserRoutingStatusEventRequest A maximum of 100 events are allowed per request
//
// swagger:model BatchUserRoutingStatusEventRequest
type BatchUserRoutingStatusEventRequest struct {

	// UserRoutingStatus events for this batch
	UserRoutingStatusEvents []*UserRoutingStatusEvent `json:"userRoutingStatusEvents"`
}

// Validate validates this batch user routing status event request
func (m *BatchUserRoutingStatusEventRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserRoutingStatusEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchUserRoutingStatusEventRequest) validateUserRoutingStatusEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.UserRoutingStatusEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.UserRoutingStatusEvents); i++ {
		if swag.IsZero(m.UserRoutingStatusEvents[i]) { // not required
			continue
		}

		if m.UserRoutingStatusEvents[i] != nil {
			if err := m.UserRoutingStatusEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userRoutingStatusEvents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userRoutingStatusEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this batch user routing status event request based on the context it is used
func (m *BatchUserRoutingStatusEventRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUserRoutingStatusEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BatchUserRoutingStatusEventRequest) contextValidateUserRoutingStatusEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.UserRoutingStatusEvents); i++ {

		if m.UserRoutingStatusEvents[i] != nil {
			if err := m.UserRoutingStatusEvents[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userRoutingStatusEvents" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("userRoutingStatusEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BatchUserRoutingStatusEventRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BatchUserRoutingStatusEventRequest) UnmarshalBinary(b []byte) error {
	var res BatchUserRoutingStatusEventRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
