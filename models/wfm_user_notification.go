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

// WfmUserNotification wfm user notification
//
// swagger:model WfmUserNotification
type WfmUserNotification struct {

	// An adherence explanation notification.  Only set if type == AdherenceExplanation
	// Read Only: true
	AdherenceExplanation *AdherenceExplanationNotification `json:"adherenceExplanation,omitempty"`

	// Whether this notification is for an agent
	// Read Only: true
	AgentNotification *bool `json:"agentNotification"`

	// The immutable globally unique identifier for the object.
	// Required: true
	ID *string `json:"id"`

	// Whether this notification has been marked "read"
	// Required: true
	MarkedAsRead *bool `json:"markedAsRead"`

	// The group ID of the notification (mutable, may change  on update)
	// Required: true
	MutableGroupID *string `json:"mutableGroupId"`

	// Other notification IDs in group.  This field is only populated in real-time notifications
	// Read Only: true
	// Unique: true
	OtherNotificationIdsInGroup []string `json:"otherNotificationIdsInGroup"`

	// A shift trade notification.  Only set if type == ShiftTrade
	// Read Only: true
	ShiftTrade *ShiftTradeNotification `json:"shiftTrade,omitempty"`

	// A time off request notification.  Only set if type == TimeOffRequest
	// Read Only: true
	TimeOffRequest *TimeOffRequestNotification `json:"timeOffRequest,omitempty"`

	// The timestamp for this notification. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`

	// The type of this notification
	// Read Only: true
	// Enum: [AdherenceExplanation ShiftTrade TimeOffRequest]
	Type string `json:"type,omitempty"`
}

// Validate validates this wfm user notification
func (m *WfmUserNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdherenceExplanation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMarkedAsRead(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMutableGroupID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOtherNotificationIdsInGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShiftTrade(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOffRequest(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmUserNotification) validateAdherenceExplanation(formats strfmt.Registry) error {
	if swag.IsZero(m.AdherenceExplanation) { // not required
		return nil
	}

	if m.AdherenceExplanation != nil {
		if err := m.AdherenceExplanation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adherenceExplanation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adherenceExplanation")
			}
			return err
		}
	}

	return nil
}

func (m *WfmUserNotification) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *WfmUserNotification) validateMarkedAsRead(formats strfmt.Registry) error {

	if err := validate.Required("markedAsRead", "body", m.MarkedAsRead); err != nil {
		return err
	}

	return nil
}

func (m *WfmUserNotification) validateMutableGroupID(formats strfmt.Registry) error {

	if err := validate.Required("mutableGroupId", "body", m.MutableGroupID); err != nil {
		return err
	}

	return nil
}

func (m *WfmUserNotification) validateOtherNotificationIdsInGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.OtherNotificationIdsInGroup) { // not required
		return nil
	}

	if err := validate.UniqueItems("otherNotificationIdsInGroup", "body", m.OtherNotificationIdsInGroup); err != nil {
		return err
	}

	return nil
}

func (m *WfmUserNotification) validateShiftTrade(formats strfmt.Registry) error {
	if swag.IsZero(m.ShiftTrade) { // not required
		return nil
	}

	if m.ShiftTrade != nil {
		if err := m.ShiftTrade.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shiftTrade")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shiftTrade")
			}
			return err
		}
	}

	return nil
}

func (m *WfmUserNotification) validateTimeOffRequest(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeOffRequest) { // not required
		return nil
	}

	if m.TimeOffRequest != nil {
		if err := m.TimeOffRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeOffRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeOffRequest")
			}
			return err
		}
	}

	return nil
}

func (m *WfmUserNotification) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

var wfmUserNotificationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AdherenceExplanation","ShiftTrade","TimeOffRequest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		wfmUserNotificationTypeTypePropEnum = append(wfmUserNotificationTypeTypePropEnum, v)
	}
}

const (

	// WfmUserNotificationTypeAdherenceExplanation captures enum value "AdherenceExplanation"
	WfmUserNotificationTypeAdherenceExplanation string = "AdherenceExplanation"

	// WfmUserNotificationTypeShiftTrade captures enum value "ShiftTrade"
	WfmUserNotificationTypeShiftTrade string = "ShiftTrade"

	// WfmUserNotificationTypeTimeOffRequest captures enum value "TimeOffRequest"
	WfmUserNotificationTypeTimeOffRequest string = "TimeOffRequest"
)

// prop value enum
func (m *WfmUserNotification) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, wfmUserNotificationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WfmUserNotification) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this wfm user notification based on the context it is used
func (m *WfmUserNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdherenceExplanation(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateAgentNotification(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOtherNotificationIdsInGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShiftTrade(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeOffRequest(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimestamp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WfmUserNotification) contextValidateAdherenceExplanation(ctx context.Context, formats strfmt.Registry) error {

	if m.AdherenceExplanation != nil {
		if err := m.AdherenceExplanation.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adherenceExplanation")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adherenceExplanation")
			}
			return err
		}
	}

	return nil
}

func (m *WfmUserNotification) contextValidateAgentNotification(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "agentNotification", "body", m.AgentNotification); err != nil {
		return err
	}

	return nil
}

func (m *WfmUserNotification) contextValidateOtherNotificationIdsInGroup(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "otherNotificationIdsInGroup", "body", []string(m.OtherNotificationIdsInGroup)); err != nil {
		return err
	}

	return nil
}

func (m *WfmUserNotification) contextValidateShiftTrade(ctx context.Context, formats strfmt.Registry) error {

	if m.ShiftTrade != nil {
		if err := m.ShiftTrade.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shiftTrade")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shiftTrade")
			}
			return err
		}
	}

	return nil
}

func (m *WfmUserNotification) contextValidateTimeOffRequest(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeOffRequest != nil {
		if err := m.TimeOffRequest.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeOffRequest")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeOffRequest")
			}
			return err
		}
	}

	return nil
}

func (m *WfmUserNotification) contextValidateTimestamp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "timestamp", "body", strfmt.DateTime(m.Timestamp)); err != nil {
		return err
	}

	return nil
}

func (m *WfmUserNotification) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "type", "body", string(m.Type)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WfmUserNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WfmUserNotification) UnmarshalBinary(b []byte) error {
	var res WfmUserNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
