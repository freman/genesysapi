// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateShareResponse create share response
//
// swagger:model CreateShareResponse
type CreateShareResponse struct {

	// failed
	Failed []*Share `json:"failed"`

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

	// succeeded
	Succeeded []*Share `json:"succeeded"`

	// workspace
	Workspace *DomainEntityRef `json:"workspace,omitempty"`
}

// Validate validates this create share response
func (m *CreateShareResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFailed(formats); err != nil {
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

	if err := m.validateSucceeded(formats); err != nil {
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

func (m *CreateShareResponse) validateFailed(formats strfmt.Registry) error {
	if swag.IsZero(m.Failed) { // not required
		return nil
	}

	for i := 0; i < len(m.Failed); i++ {
		if swag.IsZero(m.Failed[i]) { // not required
			continue
		}

		if m.Failed[i] != nil {
			if err := m.Failed[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateShareResponse) validateMember(formats strfmt.Registry) error {
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

var createShareResponseTypeMemberTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USER","GROUP","PUBLIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createShareResponseTypeMemberTypePropEnum = append(createShareResponseTypeMemberTypePropEnum, v)
	}
}

const (

	// CreateShareResponseMemberTypeUSER captures enum value "USER"
	CreateShareResponseMemberTypeUSER string = "USER"

	// CreateShareResponseMemberTypeGROUP captures enum value "GROUP"
	CreateShareResponseMemberTypeGROUP string = "GROUP"

	// CreateShareResponseMemberTypePUBLIC captures enum value "PUBLIC"
	CreateShareResponseMemberTypePUBLIC string = "PUBLIC"
)

// prop value enum
func (m *CreateShareResponse) validateMemberTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createShareResponseTypeMemberTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateShareResponse) validateMemberType(formats strfmt.Registry) error {
	if swag.IsZero(m.MemberType) { // not required
		return nil
	}

	// value enum
	if err := m.validateMemberTypeEnum("memberType", "body", m.MemberType); err != nil {
		return err
	}

	return nil
}

func (m *CreateShareResponse) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CreateShareResponse) validateSharedBy(formats strfmt.Registry) error {
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

func (m *CreateShareResponse) validateSharedEntity(formats strfmt.Registry) error {
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

var createShareResponseTypeSharedEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DOCUMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createShareResponseTypeSharedEntityTypePropEnum = append(createShareResponseTypeSharedEntityTypePropEnum, v)
	}
}

const (

	// CreateShareResponseSharedEntityTypeDOCUMENT captures enum value "DOCUMENT"
	CreateShareResponseSharedEntityTypeDOCUMENT string = "DOCUMENT"
)

// prop value enum
func (m *CreateShareResponse) validateSharedEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createShareResponseTypeSharedEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CreateShareResponse) validateSharedEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.SharedEntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateSharedEntityTypeEnum("sharedEntityType", "body", m.SharedEntityType); err != nil {
		return err
	}

	return nil
}

func (m *CreateShareResponse) validateSucceeded(formats strfmt.Registry) error {
	if swag.IsZero(m.Succeeded) { // not required
		return nil
	}

	for i := 0; i < len(m.Succeeded); i++ {
		if swag.IsZero(m.Succeeded[i]) { // not required
			continue
		}

		if m.Succeeded[i] != nil {
			if err := m.Succeeded[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("succeeded" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("succeeded" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateShareResponse) validateWorkspace(formats strfmt.Registry) error {
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

// ContextValidate validate this create share response based on the context it is used
func (m *CreateShareResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFailed(ctx, formats); err != nil {
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

	if err := m.contextValidateSucceeded(ctx, formats); err != nil {
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

func (m *CreateShareResponse) contextValidateFailed(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Failed); i++ {

		if m.Failed[i] != nil {
			if err := m.Failed[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("failed" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("failed" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateShareResponse) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *CreateShareResponse) contextValidateMember(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateShareResponse) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *CreateShareResponse) contextValidateSharedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateShareResponse) contextValidateSharedEntity(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CreateShareResponse) contextValidateSucceeded(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Succeeded); i++ {

		if m.Succeeded[i] != nil {
			if err := m.Succeeded[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("succeeded" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("succeeded" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CreateShareResponse) contextValidateWorkspace(ctx context.Context, formats strfmt.Registry) error {

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
func (m *CreateShareResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateShareResponse) UnmarshalBinary(b []byte) error {
	var res CreateShareResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
