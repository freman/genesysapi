// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateAdminTimeOffRequest create admin time off request
//
// swagger:model CreateAdminTimeOffRequest
type CreateAdminTimeOffRequest struct {

	// The ID of the activity code associated with this time off request. Activity code must be of the TimeOff category
	// Required: true
	ActivityCodeID *string `json:"activityCodeId"`

	// The daily duration of this time off request in minutes
	// Required: true
	DailyDurationMinutes *int32 `json:"dailyDurationMinutes"`

	// A set of dates in yyyy-MM-dd format.  Should be interpreted in the management unit's configured time zone.
	FullDayManagementUnitDates []string `json:"fullDayManagementUnitDates"`

	// Notes about the time off request
	Notes string `json:"notes,omitempty"`

	// Whether this is a paid time off request
	Paid bool `json:"paid"`

	// A set of start date-times in ISO-8601 format for partial day requests.
	PartialDayStartDateTimes []strfmt.DateTime `json:"partialDayStartDateTimes"`

	// The status of this time off request
	// Required: true
	// Enum: [PENDING APPROVED]
	Status *string `json:"status"`

	// A set of IDs for users to associate with this time off request
	// Required: true
	// Unique: true
	Users []*UserReference `json:"users"`
}

// Validate validates this create admin time off request
func (m *CreateAdminTimeOffRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivityCodeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDailyDurationMinutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartialDayStartDateTimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAdminTimeOffRequest) validateActivityCodeID(formats strfmt.Registry) error {

	if err := validate.Required("activityCodeId", "body", m.ActivityCodeID); err != nil {
		return err
	}

	return nil
}

func (m *CreateAdminTimeOffRequest) validateDailyDurationMinutes(formats strfmt.Registry) error {

	if err := validate.Required("dailyDurationMinutes", "body", m.DailyDurationMinutes); err != nil {
		return err
	}

	return nil
}

func (m *CreateAdminTimeOffRequest) validatePartialDayStartDateTimes(formats strfmt.Registry) error {
	if swag.IsZero(m.PartialDayStartDateTimes) { // not required
		return nil
	}

	for i := 0; i < len(m.PartialDayStartDateTimes); i++ {

		if err := validate.FormatOf("partialDayStartDateTimes"+"."+strconv.Itoa(i), "body", "date-time", m.PartialDayStartDateTimes[i].String(), formats); err != nil {
			return err
		}

	}

	return nil
}

var createAdminTimeOffRequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PENDING","APPROVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createAdminTimeOffRequestTypeStatusPropEnum = append(createAdminTimeOffRequestTypeStatusPropEnum, v)
	}
}

const (

	// CreateAdminTimeOffRequestStatusPENDING captures enum value "PENDING"
	CreateAdminTimeOffRequestStatusPENDING string = "PENDING"

	// CreateAdminTimeOffRequestStatusAPPROVED captures enum value "APPROVED"
	CreateAdminTimeOffRequestStatusAPPROVED string = "APPROVED"
)

// prop value enum
func (m *CreateAdminTimeOffRequest) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createAdminTimeOffRequestTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateAdminTimeOffRequest) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *CreateAdminTimeOffRequest) validateUsers(formats strfmt.Registry) error {

	if err := validate.Required("users", "body", m.Users); err != nil {
		return err
	}

	if err := validate.UniqueItems("users", "body", m.Users); err != nil {
		return err
	}

	for i := 0; i < len(m.Users); i++ {
		if swag.IsZero(m.Users[i]) { // not required
			continue
		}

		if m.Users[i] != nil {
			if err := m.Users[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create admin time off request based on the context it is used
func (m *CreateAdminTimeOffRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateUsers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateAdminTimeOffRequest) contextValidateUsers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Users); i++ {

		if m.Users[i] != nil {
			if err := m.Users[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("users" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("users" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateAdminTimeOffRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateAdminTimeOffRequest) UnmarshalBinary(b []byte) error {
	var res CreateAdminTimeOffRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
