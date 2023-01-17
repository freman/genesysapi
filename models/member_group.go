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

// MemberGroup member group
//
// swagger:model MemberGroup
type MemberGroup struct {

	// The division to which this entity belongs.
	Division *WritableDivision `json:"division,omitempty"`

	// The globally unique identifier for the object.
	ID string `json:"id,omitempty"`

	// The number of members in this group
	// Read Only: true
	MemberCount int32 `json:"memberCount,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The group type
	// Enum: [TEAM GROUP SKILLGROUP]
	Type string `json:"type,omitempty"`
}

// Validate validates this member group
func (m *MemberGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
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

func (m *MemberGroup) validateDivision(formats strfmt.Registry) error {
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

func (m *MemberGroup) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var memberGroupTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TEAM","GROUP","SKILLGROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		memberGroupTypeTypePropEnum = append(memberGroupTypeTypePropEnum, v)
	}
}

const (

	// MemberGroupTypeTEAM captures enum value "TEAM"
	MemberGroupTypeTEAM string = "TEAM"

	// MemberGroupTypeGROUP captures enum value "GROUP"
	MemberGroupTypeGROUP string = "GROUP"

	// MemberGroupTypeSKILLGROUP captures enum value "SKILLGROUP"
	MemberGroupTypeSKILLGROUP string = "SKILLGROUP"
)

// prop value enum
func (m *MemberGroup) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, memberGroupTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MemberGroup) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this member group based on the context it is used
func (m *MemberGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDivision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMemberCount(ctx, formats); err != nil {
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

func (m *MemberGroup) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MemberGroup) contextValidateMemberCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "memberCount", "body", int32(m.MemberCount)); err != nil {
		return err
	}

	return nil
}

func (m *MemberGroup) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MemberGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MemberGroup) UnmarshalBinary(b []byte) error {
	var res MemberGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
