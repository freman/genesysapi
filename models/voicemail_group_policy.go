// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VoicemailGroupPolicy voicemail group policy
//
// swagger:model VoicemailGroupPolicy
type VoicemailGroupPolicy struct {

	// Whether voicemail is enabled for the group
	Enabled bool `json:"enabled"`

	// The group associated with the policy
	// Read Only: true
	Group *Group `json:"group,omitempty"`

	// Specifies if the members in this group should be contacted randomly, in a specific order, or by round-robin.
	// Enum: [RANDOM ROUND_ROBIN SEQUENTIAL]
	GroupAlertType string `json:"groupAlertType,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	//  A fallback group to contact when all of the members in this group did not answer the call.
	OverflowGroupID string `json:"overflowGroupId,omitempty"`

	// How many seconds to ring before rotating to the next member in the group
	RotateCallsSecs int32 `json:"rotateCallsSecs,omitempty"`

	// Whether email notifications are sent to group members when a new voicemail is received
	SendEmailNotifications bool `json:"sendEmailNotifications"`

	// How many rotations to go through
	StopRingingAfterRotations int32 `json:"stopRingingAfterRotations,omitempty"`
}

// Validate validates this voicemail group policy
func (m *VoicemailGroupPolicy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroupAlertType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VoicemailGroupPolicy) validateGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.Group) { // not required
		return nil
	}

	if m.Group != nil {
		if err := m.Group.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("group")
			}
			return err
		}
	}

	return nil
}

var voicemailGroupPolicyTypeGroupAlertTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RANDOM","ROUND_ROBIN","SEQUENTIAL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		voicemailGroupPolicyTypeGroupAlertTypePropEnum = append(voicemailGroupPolicyTypeGroupAlertTypePropEnum, v)
	}
}

const (

	// VoicemailGroupPolicyGroupAlertTypeRANDOM captures enum value "RANDOM"
	VoicemailGroupPolicyGroupAlertTypeRANDOM string = "RANDOM"

	// VoicemailGroupPolicyGroupAlertTypeROUNDROBIN captures enum value "ROUND_ROBIN"
	VoicemailGroupPolicyGroupAlertTypeROUNDROBIN string = "ROUND_ROBIN"

	// VoicemailGroupPolicyGroupAlertTypeSEQUENTIAL captures enum value "SEQUENTIAL"
	VoicemailGroupPolicyGroupAlertTypeSEQUENTIAL string = "SEQUENTIAL"
)

// prop value enum
func (m *VoicemailGroupPolicy) validateGroupAlertTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, voicemailGroupPolicyTypeGroupAlertTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *VoicemailGroupPolicy) validateGroupAlertType(formats strfmt.Registry) error {

	if swag.IsZero(m.GroupAlertType) { // not required
		return nil
	}

	// value enum
	if err := m.validateGroupAlertTypeEnum("groupAlertType", "body", m.GroupAlertType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VoicemailGroupPolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VoicemailGroupPolicy) UnmarshalBinary(b []byte) error {
	var res VoicemailGroupPolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}