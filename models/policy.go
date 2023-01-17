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

// Policy policy
//
// swagger:model Policy
type Policy struct {

	// Actions
	Actions *PolicyActions `json:"actions,omitempty"`

	// Conditions
	Conditions *PolicyConditions `json:"conditions,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enabled
	Enabled bool `json:"enabled"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Conditions and actions per media type
	MediaPolicies *MediaPolicies `json:"mediaPolicies,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// order
	Order int32 `json:"order,omitempty"`

	// policy errors
	PolicyErrors *PolicyErrors `json:"policyErrors,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this policy
func (m *Policy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMediaPolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePolicyErrors(formats); err != nil {
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

func (m *Policy) validateActions(formats strfmt.Registry) error {
	if swag.IsZero(m.Actions) { // not required
		return nil
	}

	if m.Actions != nil {
		if err := m.Actions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.Conditions) { // not required
		return nil
	}

	if m.Conditions != nil {
		if err := m.Conditions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateCreatedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validateMediaPolicies(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaPolicies) { // not required
		return nil
	}

	if m.MediaPolicies != nil {
		if err := m.MediaPolicies.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaPolicies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaPolicies")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateModifiedDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Policy) validatePolicyErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyErrors) { // not required
		return nil
	}

	if m.PolicyErrors != nil {
		if err := m.PolicyErrors.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyErrors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policyErrors")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this policy based on the context it is used
func (m *Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMediaPolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePolicyErrors(ctx, formats); err != nil {
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

func (m *Policy) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	if m.Actions != nil {
		if err := m.Actions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actions")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateConditions(ctx context.Context, formats strfmt.Registry) error {

	if m.Conditions != nil {
		if err := m.Conditions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("conditions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("conditions")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Policy) contextValidateMediaPolicies(ctx context.Context, formats strfmt.Registry) error {

	if m.MediaPolicies != nil {
		if err := m.MediaPolicies.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mediaPolicies")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mediaPolicies")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidatePolicyErrors(ctx context.Context, formats strfmt.Registry) error {

	if m.PolicyErrors != nil {
		if err := m.PolicyErrors.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("policyErrors")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("policyErrors")
			}
			return err
		}
	}

	return nil
}

func (m *Policy) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
