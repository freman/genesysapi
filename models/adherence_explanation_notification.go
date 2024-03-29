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

// AdherenceExplanationNotification adherence explanation notification
//
// swagger:model AdherenceExplanationNotification
type AdherenceExplanationNotification struct {

	// The agent for whom the adherence explanation applies
	Agent *UserReference `json:"agent,omitempty"`

	// The business unit to which the agent belonged at the time the adherence explanation was submitted
	BusinessUnit *BusinessUnitReference `json:"businessUnit,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The length of the adherence explanation in minutes
	LengthMinutes int32 `json:"lengthMinutes,omitempty"`

	// The management unit to which the agent belonged at the time the adherence explanation was submitted
	ManagementUnit *ManagementUnitReference `json:"managementUnit,omitempty"`

	// Notes about the adherence explanation
	Notes string `json:"notes,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The start date of the adherence explanation. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	StartDate strfmt.DateTime `json:"startDate,omitempty"`

	// The status of the adherence explanation
	// Enum: [Pending Approved Denied]
	Status string `json:"status,omitempty"`

	// The type of the adherence explanation
	// Enum: [Late]
	Type string `json:"type,omitempty"`
}

// Validate validates this adherence explanation notification
func (m *AdherenceExplanationNotification) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBusinessUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateManagementUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *AdherenceExplanationNotification) validateAgent(formats strfmt.Registry) error {
	if swag.IsZero(m.Agent) { // not required
		return nil
	}

	if m.Agent != nil {
		if err := m.Agent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *AdherenceExplanationNotification) validateBusinessUnit(formats strfmt.Registry) error {
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

func (m *AdherenceExplanationNotification) validateManagementUnit(formats strfmt.Registry) error {
	if swag.IsZero(m.ManagementUnit) { // not required
		return nil
	}

	if m.ManagementUnit != nil {
		if err := m.ManagementUnit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managementUnit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managementUnit")
			}
			return err
		}
	}

	return nil
}

func (m *AdherenceExplanationNotification) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *AdherenceExplanationNotification) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var adherenceExplanationNotificationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Approved","Denied"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		adherenceExplanationNotificationTypeStatusPropEnum = append(adherenceExplanationNotificationTypeStatusPropEnum, v)
	}
}

const (

	// AdherenceExplanationNotificationStatusPending captures enum value "Pending"
	AdherenceExplanationNotificationStatusPending string = "Pending"

	// AdherenceExplanationNotificationStatusApproved captures enum value "Approved"
	AdherenceExplanationNotificationStatusApproved string = "Approved"

	// AdherenceExplanationNotificationStatusDenied captures enum value "Denied"
	AdherenceExplanationNotificationStatusDenied string = "Denied"
)

// prop value enum
func (m *AdherenceExplanationNotification) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, adherenceExplanationNotificationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdherenceExplanationNotification) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

var adherenceExplanationNotificationTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Late"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		adherenceExplanationNotificationTypeTypePropEnum = append(adherenceExplanationNotificationTypeTypePropEnum, v)
	}
}

const (

	// AdherenceExplanationNotificationTypeLate captures enum value "Late"
	AdherenceExplanationNotificationTypeLate string = "Late"
)

// prop value enum
func (m *AdherenceExplanationNotification) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, adherenceExplanationNotificationTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *AdherenceExplanationNotification) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this adherence explanation notification based on the context it is used
func (m *AdherenceExplanationNotification) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAgent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateBusinessUnit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateManagementUnit(ctx, formats); err != nil {
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

func (m *AdherenceExplanationNotification) contextValidateAgent(ctx context.Context, formats strfmt.Registry) error {

	if m.Agent != nil {
		if err := m.Agent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("agent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("agent")
			}
			return err
		}
	}

	return nil
}

func (m *AdherenceExplanationNotification) contextValidateBusinessUnit(ctx context.Context, formats strfmt.Registry) error {

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

func (m *AdherenceExplanationNotification) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *AdherenceExplanationNotification) contextValidateManagementUnit(ctx context.Context, formats strfmt.Registry) error {

	if m.ManagementUnit != nil {
		if err := m.ManagementUnit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("managementUnit")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("managementUnit")
			}
			return err
		}
	}

	return nil
}

func (m *AdherenceExplanationNotification) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdherenceExplanationNotification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdherenceExplanationNotification) UnmarshalBinary(b []byte) error {
	var res AdherenceExplanationNotification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
