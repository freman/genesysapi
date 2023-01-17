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

// BuScheduleMetadata bu schedule metadata
//
// swagger:model BuScheduleMetadata
type BuScheduleMetadata struct {

	// The description of this schedule
	Description string `json:"description,omitempty"`

	// Generation result summary for this schedule, if applicable
	GenerationResults *ScheduleGenerationResultSummary `json:"generationResults,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// High level per-management unit schedule metadata
	ManagementUnits []*BuManagementUnitScheduleSummary `json:"managementUnits"`

	// Version metadata for this schedule
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// Whether this schedule is published
	Published bool `json:"published"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The forecast used for this schedule, if applicable
	ShortTermForecast *BuShortTermForecastReference `json:"shortTermForecast,omitempty"`

	// The number of weeks spanned by this schedule
	WeekCount int32 `json:"weekCount,omitempty"`

	// The start week date for this schedule. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	WeekDate strfmt.Date `json:"weekDate,omitempty"`
}

// Validate validates this bu schedule metadata
func (m *BuScheduleMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGenerationResults(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementUnits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortTermForecast(formats); err != nil {
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

func (m *BuScheduleMetadata) validateGenerationResults(formats strfmt.Registry) error {
	if swag.IsZero(m.GenerationResults) { // not required
		return nil
	}

	if m.GenerationResults != nil {
		if err := m.GenerationResults.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generationResults")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("generationResults")
			}
			return err
		}
	}

	return nil
}

func (m *BuScheduleMetadata) validateManagementUnits(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagementUnits) { // not required
		return nil
	}

	for i := 0; i < len(m.ManagementUnits); i++ {
		if swag.IsZero(m.ManagementUnits[i]) { // not required
			continue
		}

		if m.ManagementUnits[i] != nil {
			if err := m.ManagementUnits[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("managementUnits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("managementUnits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuScheduleMetadata) validateMetadata(formats strfmt.Registry) error {
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

func (m *BuScheduleMetadata) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuScheduleMetadata) validateShortTermForecast(formats strfmt.Registry) error {
	if swag.IsZero(m.ShortTermForecast) { // not required
		return nil
	}

	if m.ShortTermForecast != nil {
		if err := m.ShortTermForecast.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortTermForecast")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shortTermForecast")
			}
			return err
		}
	}

	return nil
}

func (m *BuScheduleMetadata) validateWeekDate(formats strfmt.Registry) error {
	if swag.IsZero(m.WeekDate) { // not required
		return nil
	}

	if err := validate.FormatOf("weekDate", "body", "date", m.WeekDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bu schedule metadata based on the context it is used
func (m *BuScheduleMetadata) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGenerationResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManagementUnits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateShortTermForecast(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuScheduleMetadata) contextValidateGenerationResults(ctx context.Context, formats strfmt.Registry) error {

	if m.GenerationResults != nil {
		if err := m.GenerationResults.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generationResults")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("generationResults")
			}
			return err
		}
	}

	return nil
}

func (m *BuScheduleMetadata) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *BuScheduleMetadata) contextValidateManagementUnits(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ManagementUnits); i++ {

		if m.ManagementUnits[i] != nil {
			if err := m.ManagementUnits[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("managementUnits" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("managementUnits" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BuScheduleMetadata) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *BuScheduleMetadata) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *BuScheduleMetadata) contextValidateShortTermForecast(ctx context.Context, formats strfmt.Registry) error {

	if m.ShortTermForecast != nil {
		if err := m.ShortTermForecast.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shortTermForecast")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("shortTermForecast")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuScheduleMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuScheduleMetadata) UnmarshalBinary(b []byte) error {
	var res BuScheduleMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
