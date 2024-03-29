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

// Share share
//
// swagger:model Share
type Share struct {

	// group
	Group *Group `json:"group,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// member
	Member *DomainEntityRef `json:"member,omitempty"`

	// member type
	// Enum: [USER GROUP PUBLIC]
	MemberType string `json:"memberType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// shared by
	SharedBy *DomainEntityRef `json:"sharedBy,omitempty"`

	// shared entity
	SharedEntity *DomainEntityRef `json:"sharedEntity,omitempty"`

	// shared entity type
	// Enum: [DOCUMENT]
	SharedEntityType string `json:"sharedEntityType,omitempty"`

	// user
	User *User `json:"user,omitempty"`

	// workspace
	Workspace *DomainEntityRef `json:"workspace,omitempty"`
}

// Validate validates this share
func (m *Share) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMember(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedEntity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSharedEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Share) validateGroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Share) validateMember(formats strfmt.Registry) error {
	if swag.IsZero(m.Member) { // not required
		return nil
	}

	if m.Member != nil {
		if err := m.Member.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

var shareTypeMemberTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USER","GROUP","PUBLIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shareTypeMemberTypePropEnum = append(shareTypeMemberTypePropEnum, v)
	}
}

const (

	// ShareMemberTypeUSER captures enum value "USER"
	ShareMemberTypeUSER string = "USER"

	// ShareMemberTypeGROUP captures enum value "GROUP"
	ShareMemberTypeGROUP string = "GROUP"

	// ShareMemberTypePUBLIC captures enum value "PUBLIC"
	ShareMemberTypePUBLIC string = "PUBLIC"
)

// prop value enum
func (m *Share) validateMemberTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shareTypeMemberTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Share) validateMemberType(formats strfmt.Registry) error {
	if swag.IsZero(m.MemberType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMemberTypeEnum("memberType", "body", m.MemberType); err != nil {
		return err
	}

	return nil
}

func (m *Share) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Share) validateSharedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.SharedBy) { // not required
		return nil
	}

	if m.SharedBy != nil {
		if err := m.SharedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Share) validateSharedEntity(formats strfmt.Registry) error {
	if swag.IsZero(m.SharedEntity) { // not required
		return nil
	}

	if m.SharedEntity != nil {
		if err := m.SharedEntity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharedEntity")
			}
			return err
		}
	}

	return nil
}

var shareTypeSharedEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DOCUMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		shareTypeSharedEntityTypePropEnum = append(shareTypeSharedEntityTypePropEnum, v)
	}
}

const (

	// ShareSharedEntityTypeDOCUMENT captures enum value "DOCUMENT"
	ShareSharedEntityTypeDOCUMENT string = "DOCUMENT"
)

// prop value enum
func (m *Share) validateSharedEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, shareTypeSharedEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Share) validateSharedEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.SharedEntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSharedEntityTypeEnum("sharedEntityType", "body", m.SharedEntityType); err != nil {
		return err
	}

	return nil
}

func (m *Share) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *Share) validateWorkspace(formats strfmt.Registry) error {
	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {
		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this share based on the context it is used
func (m *Share) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMember(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSharedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSharedEntity(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWorkspace(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Share) contextValidateGroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Group != nil {
		if err := m.Group.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

func (m *Share) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Share) contextValidateMember(ctx context.Context, formats strfmt.Registry) error {

	if m.Member != nil {
		if err := m.Member.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("member")
			}
			return err
		}
	}

	return nil
}

func (m *Share) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Share) contextValidateSharedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.SharedBy != nil {
		if err := m.SharedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Share) contextValidateSharedEntity(ctx context.Context, formats strfmt.Registry) error {

	if m.SharedEntity != nil {
		if err := m.SharedEntity.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedEntity")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sharedEntity")
			}
			return err
		}
	}

	return nil
}

func (m *Share) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *Share) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

	if m.Workspace != nil {
		if err := m.Workspace.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Share) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Share) UnmarshalBinary(b []byte) error {
	var res Share
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
