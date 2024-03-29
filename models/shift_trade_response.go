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

// ShiftTradeResponse shift trade response
//
// swagger:model ShiftTradeResponse
type ShiftTradeResponse struct {

	// acceptable intervals
	AcceptableIntervals []string `json:"acceptableIntervals"`

	// When this shift trade offer will expire if not matched or approved. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	Expiration strfmt.DateTime `json:"expiration,omitempty"`

	// The ID of this shift trade
	ID string `json:"id,omitempty"`

	// The end date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	InitiatingShiftEnd strfmt.DateTime `json:"initiatingShiftEnd,omitempty"`

	// The ID of the shift offered for trade by the initiating user
	InitiatingShiftID string `json:"initiatingShiftId,omitempty"`

	// The start date/time of the shift being offered for trade. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	InitiatingShiftStart strfmt.DateTime `json:"initiatingShiftStart,omitempty"`

	// The user who initiated this trade
	InitiatingUser *UserReference `json:"initiatingUser,omitempty"`

	// Version data for this trade
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// Whether this is a one-sided shift trade (e.g. the initiating user is not asking for a shift in return)
	OneSided bool `json:"oneSided"`

	// The end date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ReceivingShiftEnd strfmt.DateTime `json:"receivingShiftEnd,omitempty"`

	// The ID of the shift being exchanged for the initiating shift, null if the receiving user is picking up a shift
	ReceivingShiftID string `json:"receivingShiftId,omitempty"`

	// The start date/time of the receiving shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ReceivingShiftStart strfmt.DateTime `json:"receivingShiftStart,omitempty"`

	// The user matching the trade, or if the state is not Matched, the user to whom the trade request was sent
	ReceivingUser *UserReference `json:"receivingUser,omitempty"`

	// The user who reviewed this shift trade
	ReviewedBy *UserReference `json:"reviewedBy,omitempty"`

	// The timestamp when this shift trade was reviewed. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ReviewedDate strfmt.DateTime `json:"reviewedDate,omitempty"`

	// A reference to the associated schedule
	Schedule *BuScheduleReferenceForMuRoute `json:"schedule,omitempty"`

	// The state of this shift trade
	// Enum: [Unmatched Matched Approved Denied Expired Canceled]
	State string `json:"state,omitempty"`

	// The start week date of the associated schedule in yyyy-MM-dd format. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	WeekDate strfmt.Date `json:"weekDate,omitempty"`
}

// Validate validates this shift trade response
func (m *ShiftTradeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExpiration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatingShiftEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatingShiftStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiatingUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivingShiftEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivingShiftStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReceivingUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchedule(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShiftTradeResponse) validateExpiration(formats strfmt.Registry) error {
	if swag.IsZero(m.Expiration) { // not required
		return nil
	}

	if err := validate.FormatOf("expiration", "body", "date-time", m.Expiration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShiftTradeResponse) validateInitiatingShiftEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatingShiftEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("initiatingShiftEnd", "body", "date-time", m.InitiatingShiftEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShiftTradeResponse) validateInitiatingShiftStart(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatingShiftStart) { // not required
		return nil
	}

	if err := validate.FormatOf("initiatingShiftStart", "body", "date-time", m.InitiatingShiftStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShiftTradeResponse) validateInitiatingUser(formats strfmt.Registry) error {
	if swag.IsZero(m.InitiatingUser) { // not required
		return nil
	}

	if m.InitiatingUser != nil {
		if err := m.InitiatingUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiatingUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("initiatingUser")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) validateReceivingShiftEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.ReceivingShiftEnd) { // not required
		return nil
	}

	if err := validate.FormatOf("receivingShiftEnd", "body", "date-time", m.ReceivingShiftEnd.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShiftTradeResponse) validateReceivingShiftStart(formats strfmt.Registry) error {
	if swag.IsZero(m.ReceivingShiftStart) { // not required
		return nil
	}

	if err := validate.FormatOf("receivingShiftStart", "body", "date-time", m.ReceivingShiftStart.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShiftTradeResponse) validateReceivingUser(formats strfmt.Registry) error {
	if swag.IsZero(m.ReceivingUser) { // not required
		return nil
	}

	if m.ReceivingUser != nil {
		if err := m.ReceivingUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receivingUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("receivingUser")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) validateReviewedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ReviewedBy) { // not required
		return nil
	}

	if m.ReviewedBy != nil {
		if err := m.ReviewedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reviewedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reviewedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) validateReviewedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ReviewedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("reviewedDate", "body", "date-time", m.ReviewedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ShiftTradeResponse) validateSchedule(formats strfmt.Registry) error {
	if swag.IsZero(m.Schedule) { // not required
		return nil
	}

	if m.Schedule != nil {
		if err := m.Schedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

var shiftTradeResponseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Unmatched","Matched","Approved","Denied","Expired","Canceled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shiftTradeResponseTypeStatePropEnum = append(shiftTradeResponseTypeStatePropEnum, v)
	}
}

const (

	// ShiftTradeResponseStateUnmatched captures enum value "Unmatched"
	ShiftTradeResponseStateUnmatched string = "Unmatched"

	// ShiftTradeResponseStateMatched captures enum value "Matched"
	ShiftTradeResponseStateMatched string = "Matched"

	// ShiftTradeResponseStateApproved captures enum value "Approved"
	ShiftTradeResponseStateApproved string = "Approved"

	// ShiftTradeResponseStateDenied captures enum value "Denied"
	ShiftTradeResponseStateDenied string = "Denied"

	// ShiftTradeResponseStateExpired captures enum value "Expired"
	ShiftTradeResponseStateExpired string = "Expired"

	// ShiftTradeResponseStateCanceled captures enum value "Canceled"
	ShiftTradeResponseStateCanceled string = "Canceled"
)

// prop value enum
func (m *ShiftTradeResponse) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shiftTradeResponseTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ShiftTradeResponse) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *ShiftTradeResponse) validateWeekDate(formats strfmt.Registry) error {
	if swag.IsZero(m.WeekDate) { // not required
		return nil
	}

	if err := validate.FormatOf("weekDate", "body", "date", m.WeekDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this shift trade response based on the context it is used
func (m *ShiftTradeResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateInitiatingUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReceivingUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReviewedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ShiftTradeResponse) contextValidateInitiatingUser(ctx context.Context, formats strfmt.Registry) error {

	if m.InitiatingUser != nil {
		if err := m.InitiatingUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiatingUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("initiatingUser")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) contextValidateReceivingUser(ctx context.Context, formats strfmt.Registry) error {

	if m.ReceivingUser != nil {
		if err := m.ReceivingUser.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("receivingUser")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("receivingUser")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) contextValidateReviewedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ReviewedBy != nil {
		if err := m.ReviewedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("reviewedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("reviewedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ShiftTradeResponse) contextValidateSchedule(ctx context.Context, formats strfmt.Registry) error {

	if m.Schedule != nil {
		if err := m.Schedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("schedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ShiftTradeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ShiftTradeResponse) UnmarshalBinary(b []byte) error {
	var res ShiftTradeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
