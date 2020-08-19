// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScimV2Group Defines a SCIM group.
//
// swagger:model ScimV2Group
type ScimV2Group struct {

	// The display name of the group.
	// Required: true
	// Read Only: true
	DisplayName string `json:"displayName"`

	// The external ID of the group. Set by the provisioning client. "caseExact" is set to "true". "mutability" is set to "readWrite".
	ExternalID string `json:"externalId,omitempty"`

	// The ID of the SCIM resource. Set by the service provider. "caseExact" is set to "true". "mutability" is set to "readOnly". "returned" is set to "always".
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The list of members in the group.
	Members []*ScimV2MemberReference `json:"members"`

	// The metadata of the SCIM resource.
	// Read Only: true
	Meta *ScimMetadata `json:"meta,omitempty"`

	// The list of supported schemas.
	// Read Only: true
	Schemas []string `json:"schemas"`
}

// Validate validates this scim v2 group
func (m *ScimV2Group) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMembers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScimV2Group) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.RequiredString("displayName", "body", string(m.DisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *ScimV2Group) validateMembers(formats strfmt.Registry) error {

	if swag.IsZero(m.Members) { // not required
		return nil
	}

	for i := 0; i < len(m.Members); i++ {
		if swag.IsZero(m.Members[i]) { // not required
			continue
		}

		if m.Members[i] != nil {
			if err := m.Members[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("members" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ScimV2Group) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimV2Group) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimV2Group) UnmarshalBinary(b []byte) error {
	var res ScimV2Group
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
