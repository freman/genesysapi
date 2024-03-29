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
	"github.com/go-openapi/validate"
)

// MessagingCampaignSchedule messaging campaign schedule
//
// swagger:model MessagingCampaignSchedule
type MessagingCampaignSchedule struct {

	// Creation time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Last modified time of the entity. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// A list of intervals during which to run the associated Campaign.
	// Required: true
	Intervals []*ScheduleInterval `json:"intervals"`

	// The Campaign that this messaging campaign schedule is for.
	// Required: true
	MessagingCampaign *DomainEntityRef `json:"messagingCampaign"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The time zone for this messaging campaign schedule.
	// Example: Africa/Abidjan
	TimeZone string `json:"timeZone,omitempty"`

	// Required for updates, must match the version number of the most recent update
	Version int32 `json:"version,omitempty"`
}

// Validate validates this messaging campaign schedule
func (m *MessagingCampaignSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntervals(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessagingCampaign(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MessagingCampaignSchedule) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaignSchedule) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaignSchedule) validateIntervals(formats strfmt.Registry) error {

	if err := validate.Required("intervals", "body", m.Intervals); err != nil {
		return err
	}

	for i := 0; i < len(m.Intervals); i++ {
		if swag.IsZero(m.Intervals[i]) { // not required
			continue
		}

		if m.Intervals[i] != nil {
			if err := m.Intervals[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagingCampaignSchedule) validateMessagingCampaign(formats strfmt.Registry) error {

	if err := validate.Required("messagingCampaign", "body", m.MessagingCampaign); err != nil {
		return err
	}

	if m.MessagingCampaign != nil {
		if err := m.MessagingCampaign.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingCampaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("messagingCampaign")
			}
			return err
		}
	}

	return nil
}

func (m *MessagingCampaignSchedule) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this messaging campaign schedule based on the context it is used
func (m *MessagingCampaignSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIntervals(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessagingCampaign(ctx, formats); err != nil {
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

func (m *MessagingCampaignSchedule) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaignSchedule) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaignSchedule) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *MessagingCampaignSchedule) contextValidateIntervals(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Intervals); i++ {

		if m.Intervals[i] != nil {
			if err := m.Intervals[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("intervals" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("intervals" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MessagingCampaignSchedule) contextValidateMessagingCampaign(ctx context.Context, formats strfmt.Registry) error {

	if m.MessagingCampaign != nil {
		if err := m.MessagingCampaign.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingCampaign")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("messagingCampaign")
			}
			return err
		}
	}

	return nil
}

func (m *MessagingCampaignSchedule) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MessagingCampaignSchedule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MessagingCampaignSchedule) UnmarshalBinary(b []byte) error {
	var res MessagingCampaignSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
