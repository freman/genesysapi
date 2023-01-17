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

// ManagementUnit management unit
//
// swagger:model ManagementUnit
type ManagementUnit struct {

	// The business unit to which this management unit belongs
	BusinessUnit *BusinessUnitReference `json:"businessUnit,omitempty"`

	// The date and time at which this entity was last modified.  Deprecated, use field from settings.metadata instead. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The division to which this entity belongs.
	Division *DivisionReference `json:"division,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Version info metadata for this management unit. Deprecated, use settings.metadata
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// The user who last modified this entity.  Deprecated, use field from settings.metadata instead
	// Read Only: true
	ModifiedBy *UserReference `json:"modifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The configuration settings for this management unit
	Settings *ManagementUnitSettingsResponse `json:"settings,omitempty"`

	// Start day of week for scheduling and forecasting purposes. Moving to Business Unit
	// Enum: [Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
	StartDayOfWeek string `json:"startDayOfWeek,omitempty"`

	// The time zone for the management unit in standard Olson format.  Moving to Business Unit
	TimeZone string `json:"timeZone,omitempty"`

	// The version of the underlying entity.  Deprecated, use field from settings.metadata instead
	// Read Only: true
	Version int32 `json:"version,omitempty"`
}

// Validate validates this management unit
func (m *ManagementUnit) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBusinessUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDayOfWeek(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementUnit) validateBusinessUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.BusinessUnit) { // not required
		return nil
	}

	if m.BusinessUnit != nil {
		if err := m.BusinessUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessUnit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessUnit")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnit) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ManagementUnit) validateDivision(formats strfmt.Registry) error {
	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnit) validateMetadata(formats strfmt.Registry) error {
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

func (m *ManagementUnit) validateModifiedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnit) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ManagementUnit) validateSettings(formats strfmt.Registry) error {
	if swag.IsZero(m.Settings) { // not required
		return nil
	}

	if m.Settings != nil {
		if err := m.Settings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

var managementUnitTypeStartDayOfWeekPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Sunday","Monday","Tuesday","Wednesday","Thursday","Friday","Saturday"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		managementUnitTypeStartDayOfWeekPropEnum = append(managementUnitTypeStartDayOfWeekPropEnum, v)
	}
}

const (

	// ManagementUnitStartDayOfWeekSunday captures enum value "Sunday"
	ManagementUnitStartDayOfWeekSunday string = "Sunday"

	// ManagementUnitStartDayOfWeekMonday captures enum value "Monday"
	ManagementUnitStartDayOfWeekMonday string = "Monday"

	// ManagementUnitStartDayOfWeekTuesday captures enum value "Tuesday"
	ManagementUnitStartDayOfWeekTuesday string = "Tuesday"

	// ManagementUnitStartDayOfWeekWednesday captures enum value "Wednesday"
	ManagementUnitStartDayOfWeekWednesday string = "Wednesday"

	// ManagementUnitStartDayOfWeekThursday captures enum value "Thursday"
	ManagementUnitStartDayOfWeekThursday string = "Thursday"

	// ManagementUnitStartDayOfWeekFriday captures enum value "Friday"
	ManagementUnitStartDayOfWeekFriday string = "Friday"

	// ManagementUnitStartDayOfWeekSaturday captures enum value "Saturday"
	ManagementUnitStartDayOfWeekSaturday string = "Saturday"
)

// prop value enum
func (m *ManagementUnit) validateStartDayOfWeekEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, managementUnitTypeStartDayOfWeekPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ManagementUnit) validateStartDayOfWeek(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDayOfWeek) { // not required
		return nil
	}

	// value enum
	if err := m.validateStartDayOfWeekEnum("startDayOfWeek", "body", m.StartDayOfWeek); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this management unit based on the context it is used
func (m *ManagementUnit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBusinessUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDivision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ManagementUnit) contextValidateBusinessUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.BusinessUnit != nil {
		if err := m.BusinessUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("businessUnit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("businessUnit")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnit) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *ManagementUnit) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

	if m.Division != nil {
		if err := m.Division.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnit) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ManagementUnit) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ManagementUnit) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnit) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *ManagementUnit) contextValidateSettings(ctx context.Context, formats strfmt.Registry) error {

	if m.Settings != nil {
		if err := m.Settings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("settings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("settings")
			}
			return err
		}
	}

	return nil
}

func (m *ManagementUnit) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", int32(m.Version)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ManagementUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ManagementUnit) UnmarshalBinary(b []byte) error {
	var res ManagementUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
