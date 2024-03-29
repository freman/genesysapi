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

// ManagementUnitSettingsResponse management unit settings response
//
// swagger:model ManagementUnitSettingsResponse
type ManagementUnitSettingsResponse struct {

	// Adherence settings for this management unit
	Adherence *AdherenceSettings `json:"adherence,omitempty"`

	// Version info metadata for the associated management unit
	// Required: true
	Metadata *WfmVersionedEntityMetadata `json:"metadata"`

	// Scheduling settings for this management unit. These settings are only available if you have the permission wfm:managementUnit:view
	Scheduling *SchedulingSettingsResponse `json:"scheduling,omitempty"`

	// Shift trade settings for this management unit
	ShiftTrading *ShiftTradeSettings `json:"shiftTrading,omitempty"`

	// Short term forecasting settings for this management unit
	ShortTermForecasting *ShortTermForecastingSettings `json:"shortTermForecasting,omitempty"`

	// Time off request settings for this management unit
	TimeOff *TimeOffRequestSettings `json:"timeOff,omitempty"`
}

// Validate validates this management unit settings response
func (m *ManagementUnitSettingsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdherence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduling(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShiftTrading(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortTermForecasting(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeOff(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementUnitSettingsResponse) validateAdherence(formats strfmt.Registry) error {
	if swag.IsZero(m.Adherence) { // not required
		return nil
	}

	if m.Adherence != nil {
		if err := m.Adherence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adherence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adherence")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) validateMetadata(formats strfmt.Registry) error {

	if err := validate.Required("metadata", "body", m.Metadata); err != nil {
		return err
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

func (m *ManagementUnitSettingsResponse) validateScheduling(formats strfmt.Registry) error {
	if swag.IsZero(m.Scheduling) { // not required
		return nil
	}

	if m.Scheduling != nil {
		if err := m.Scheduling.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduling")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) validateShiftTrading(formats strfmt.Registry) error {
	if swag.IsZero(m.ShiftTrading) { // not required
		return nil
	}

	if m.ShiftTrading != nil {
		if err := m.ShiftTrading.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shiftTrading")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shiftTrading")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) validateShortTermForecasting(formats strfmt.Registry) error {
	if swag.IsZero(m.ShortTermForecasting) { // not required
		return nil
	}

	if m.ShortTermForecasting != nil {
		if err := m.ShortTermForecasting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortTermForecasting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shortTermForecasting")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) validateTimeOff(formats strfmt.Registry) error {
	if swag.IsZero(m.TimeOff) { // not required
		return nil
	}

	if m.TimeOff != nil {
		if err := m.TimeOff.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeOff")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeOff")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this management unit settings response based on the context it is used
func (m *ManagementUnitSettingsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdherence(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScheduling(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShiftTrading(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShortTermForecasting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTimeOff(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementUnitSettingsResponse) contextValidateAdherence(ctx context.Context, formats strfmt.Registry) error {

	if m.Adherence != nil {
		if err := m.Adherence.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("adherence")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("adherence")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ManagementUnitSettingsResponse) contextValidateScheduling(ctx context.Context, formats strfmt.Registry) error {

	if m.Scheduling != nil {
		if err := m.Scheduling.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("scheduling")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("scheduling")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) contextValidateShiftTrading(ctx context.Context, formats strfmt.Registry) error {

	if m.ShiftTrading != nil {
		if err := m.ShiftTrading.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shiftTrading")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shiftTrading")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) contextValidateShortTermForecasting(ctx context.Context, formats strfmt.Registry) error {

	if m.ShortTermForecasting != nil {
		if err := m.ShortTermForecasting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortTermForecasting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shortTermForecasting")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnitSettingsResponse) contextValidateTimeOff(ctx context.Context, formats strfmt.Registry) error {

	if m.TimeOff != nil {
		if err := m.TimeOff.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("timeOff")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("timeOff")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementUnitSettingsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementUnitSettingsResponse) UnmarshalBinary(b []byte) error {
	var res ManagementUnitSettingsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
