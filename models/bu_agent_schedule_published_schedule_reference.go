// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BuAgentSchedulePublishedScheduleReference bu agent schedule published schedule reference
//
// swagger:model BuAgentSchedulePublishedScheduleReference
type BuAgentSchedulePublishedScheduleReference struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The number of weeks encompassed by the schedule
	WeekCount int32 `json:"weekCount,omitempty"`

	// The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	WeekDate strfmt.Date `json:"weekDate,omitempty"`
}

// Validate validates this bu agent schedule published schedule reference
func (m *BuAgentSchedulePublishedScheduleReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *BuAgentSchedulePublishedScheduleReference) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuAgentSchedulePublishedScheduleReference) validateWeekDate(formats strfmt.Registry) error {
	if swag.IsZero(m.WeekDate) { // not required
		return nil
	}

	if err := validate.FormatOf("weekDate", "body", "date", m.WeekDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bu agent schedule published schedule reference based on the context it is used
func (m *BuAgentSchedulePublishedScheduleReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuAgentSchedulePublishedScheduleReference) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *BuAgentSchedulePublishedScheduleReference) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuAgentSchedulePublishedScheduleReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuAgentSchedulePublishedScheduleReference) UnmarshalBinary(b []byte) error {
	var res BuAgentSchedulePublishedScheduleReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
