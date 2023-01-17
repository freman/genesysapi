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

// TimeOffBalanceJobReference time off balance job reference
//
// swagger:model TimeOffBalanceJobReference
type TimeOffBalanceJobReference struct {

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of the job
	// Enum: [Processing Complete Error]
	Status string `json:"status,omitempty"`
}

// Validate validates this time off balance job reference
func (m *TimeOffBalanceJobReference) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *TimeOffBalanceJobReference) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var timeOffBalanceJobReferenceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Processing","Complete","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		timeOffBalanceJobReferenceTypeStatusPropEnum = append(timeOffBalanceJobReferenceTypeStatusPropEnum, v)
	}
}

const (

	// TimeOffBalanceJobReferenceStatusProcessing captures enum value "Processing"
	TimeOffBalanceJobReferenceStatusProcessing string = "Processing"

	// TimeOffBalanceJobReferenceStatusComplete captures enum value "Complete"
	TimeOffBalanceJobReferenceStatusComplete string = "Complete"

	// TimeOffBalanceJobReferenceStatusError captures enum value "Error"
	TimeOffBalanceJobReferenceStatusError string = "Error"
)

// prop value enum
func (m *TimeOffBalanceJobReference) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, timeOffBalanceJobReferenceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TimeOffBalanceJobReference) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this time off balance job reference based on the context it is used
func (m *TimeOffBalanceJobReference) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *TimeOffBalanceJobReference) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *TimeOffBalanceJobReference) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TimeOffBalanceJobReference) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TimeOffBalanceJobReference) UnmarshalBinary(b []byte) error {
	var res TimeOffBalanceJobReference
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
