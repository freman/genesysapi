// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StatusChange status change
//
// swagger:model StatusChange
type StatusChange struct {

	// If applicable, the user who updated the change request to this status
	// Read Only: true
	ChangedBy string `json:"changedBy,omitempty"`

	// The date of this status change. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateStatusChanged strfmt.DateTime `json:"dateStatusChanged,omitempty"`

	// A short message describing the status change
	// Read Only: true
	Message string `json:"message,omitempty"`

	// The status the change request transitioned from
	// Read Only: true
	// Enum: [Open Approved ImplementingChange ChangeImplemented Rejected Rollback ImplementingRollback RollbackImplemented]
	PreviousStatus string `json:"previousStatus,omitempty"`

	// The reason for rejecting the limit override request
	// Read Only: true
	// Enum: [AlternativeExists IncreaseNotRequired PlatformMisuse PlatformStability OtherReason]
	RejectReason string `json:"rejectReason,omitempty"`

	// The status the change request transitioned to
	// Read Only: true
	// Enum: [Open Approved ImplementingChange ChangeImplemented Rejected Rollback ImplementingRollback RollbackImplemented]
	Status string `json:"status,omitempty"`
}

// Validate validates this status change
func (m *StatusChange) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateStatusChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePreviousStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRejectReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StatusChange) validateDateStatusChanged(formats strfmt.Registry) error {

	if swag.IsZero(m.DateStatusChanged) { // not required
		return nil
	}

	if err := validate.FormatOf("dateStatusChanged", "body", "date-time", m.DateStatusChanged.String(), formats); err != nil {
		return err
	}

	return nil
}

var statusChangeTypePreviousStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Open","Approved","ImplementingChange","ChangeImplemented","Rejected","Rollback","ImplementingRollback","RollbackImplemented"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusChangeTypePreviousStatusPropEnum = append(statusChangeTypePreviousStatusPropEnum, v)
	}
}

const (

	// StatusChangePreviousStatusOpen captures enum value "Open"
	StatusChangePreviousStatusOpen string = "Open"

	// StatusChangePreviousStatusApproved captures enum value "Approved"
	StatusChangePreviousStatusApproved string = "Approved"

	// StatusChangePreviousStatusImplementingChange captures enum value "ImplementingChange"
	StatusChangePreviousStatusImplementingChange string = "ImplementingChange"

	// StatusChangePreviousStatusChangeImplemented captures enum value "ChangeImplemented"
	StatusChangePreviousStatusChangeImplemented string = "ChangeImplemented"

	// StatusChangePreviousStatusRejected captures enum value "Rejected"
	StatusChangePreviousStatusRejected string = "Rejected"

	// StatusChangePreviousStatusRollback captures enum value "Rollback"
	StatusChangePreviousStatusRollback string = "Rollback"

	// StatusChangePreviousStatusImplementingRollback captures enum value "ImplementingRollback"
	StatusChangePreviousStatusImplementingRollback string = "ImplementingRollback"

	// StatusChangePreviousStatusRollbackImplemented captures enum value "RollbackImplemented"
	StatusChangePreviousStatusRollbackImplemented string = "RollbackImplemented"
)

// prop value enum
func (m *StatusChange) validatePreviousStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusChangeTypePreviousStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusChange) validatePreviousStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PreviousStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePreviousStatusEnum("previousStatus", "body", m.PreviousStatus); err != nil {
		return err
	}

	return nil
}

var statusChangeTypeRejectReasonPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AlternativeExists","IncreaseNotRequired","PlatformMisuse","PlatformStability","OtherReason"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusChangeTypeRejectReasonPropEnum = append(statusChangeTypeRejectReasonPropEnum, v)
	}
}

const (

	// StatusChangeRejectReasonAlternativeExists captures enum value "AlternativeExists"
	StatusChangeRejectReasonAlternativeExists string = "AlternativeExists"

	// StatusChangeRejectReasonIncreaseNotRequired captures enum value "IncreaseNotRequired"
	StatusChangeRejectReasonIncreaseNotRequired string = "IncreaseNotRequired"

	// StatusChangeRejectReasonPlatformMisuse captures enum value "PlatformMisuse"
	StatusChangeRejectReasonPlatformMisuse string = "PlatformMisuse"

	// StatusChangeRejectReasonPlatformStability captures enum value "PlatformStability"
	StatusChangeRejectReasonPlatformStability string = "PlatformStability"

	// StatusChangeRejectReasonOtherReason captures enum value "OtherReason"
	StatusChangeRejectReasonOtherReason string = "OtherReason"
)

// prop value enum
func (m *StatusChange) validateRejectReasonEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusChangeTypeRejectReasonPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusChange) validateRejectReason(formats strfmt.Registry) error {

	if swag.IsZero(m.RejectReason) { // not required
		return nil
	}

	// value enum
	if err := m.validateRejectReasonEnum("rejectReason", "body", m.RejectReason); err != nil {
		return err
	}

	return nil
}

var statusChangeTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Open","Approved","ImplementingChange","ChangeImplemented","Rejected","Rollback","ImplementingRollback","RollbackImplemented"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusChangeTypeStatusPropEnum = append(statusChangeTypeStatusPropEnum, v)
	}
}

const (

	// StatusChangeStatusOpen captures enum value "Open"
	StatusChangeStatusOpen string = "Open"

	// StatusChangeStatusApproved captures enum value "Approved"
	StatusChangeStatusApproved string = "Approved"

	// StatusChangeStatusImplementingChange captures enum value "ImplementingChange"
	StatusChangeStatusImplementingChange string = "ImplementingChange"

	// StatusChangeStatusChangeImplemented captures enum value "ChangeImplemented"
	StatusChangeStatusChangeImplemented string = "ChangeImplemented"

	// StatusChangeStatusRejected captures enum value "Rejected"
	StatusChangeStatusRejected string = "Rejected"

	// StatusChangeStatusRollback captures enum value "Rollback"
	StatusChangeStatusRollback string = "Rollback"

	// StatusChangeStatusImplementingRollback captures enum value "ImplementingRollback"
	StatusChangeStatusImplementingRollback string = "ImplementingRollback"

	// StatusChangeStatusRollbackImplemented captures enum value "RollbackImplemented"
	StatusChangeStatusRollbackImplemented string = "RollbackImplemented"
)

// prop value enum
func (m *StatusChange) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statusChangeTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *StatusChange) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StatusChange) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StatusChange) UnmarshalBinary(b []byte) error {
	var res StatusChange
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}