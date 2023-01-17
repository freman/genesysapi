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

// EdgeGroup edge group
//
// swagger:model EdgeGroup
type EdgeGroup struct {

	// The ID of the user that created the resource.
	// Read Only: true
	CreatedBy string `json:"createdBy,omitempty"`

	// The application that created the resource.
	// Read Only: true
	CreatedByApp string `json:"createdByApp,omitempty"`

	// The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The resource's description.
	Description string `json:"description,omitempty"`

	// The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// A trunk base settings assignment of trunkType "EDGE" to use for edge-to-edge communication.
	// Required: true
	EdgeTrunkBaseAssignment *TrunkBaseAssignment `json:"edgeTrunkBaseAssignment"`

	// Is this edge group hybrid.
	Hybrid bool `json:"hybrid"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Is this edge group being managed remotely.
	Managed bool `json:"managed"`

	// The ID of the user that last modified the resource.
	// Read Only: true
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
	// Read Only: true
	ModifiedByApp string `json:"modifiedByApp,omitempty"`

	// The name of the entity.
	// Required: true
	Name *string `json:"name"`

	// Trunk base settings of trunkType "PHONE" to inherit to edge logical interface for phone communication.
	// Required: true
	PhoneTrunkBases []*TrunkBase `json:"phoneTrunkBases"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Indicates if the resource is active, inactive, or deleted.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// The current version of the resource.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this edge group
func (m *EdgeGroup) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeTrunkBaseAssignment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneTrunkBases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeGroup) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) validateDivision(formats strfmt.Registry) error {
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

func (m *EdgeGroup) validateEdgeTrunkBaseAssignment(formats strfmt.Registry) error {

	if err := validate.Required("edgeTrunkBaseAssignment", "body", m.EdgeTrunkBaseAssignment); err != nil {
		return err
	}

	if m.EdgeTrunkBaseAssignment != nil {
		if err := m.EdgeTrunkBaseAssignment.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeTrunkBaseAssignment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeTrunkBaseAssignment")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeGroup) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) validatePhoneTrunkBases(formats strfmt.Registry) error {

	if err := validate.Required("phoneTrunkBases", "body", m.PhoneTrunkBases); err != nil {
		return err
	}

	for i := 0; i < len(m.PhoneTrunkBases); i++ {
		if swag.IsZero(m.PhoneTrunkBases[i]) { // not required
			continue
		}

		if m.PhoneTrunkBases[i] != nil {
			if err := m.PhoneTrunkBases[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneTrunkBases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phoneTrunkBases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeGroup) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var edgeGroupTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgeGroupTypeStatePropEnum = append(edgeGroupTypeStatePropEnum, v)
	}
}

const (

	// EdgeGroupStateActive captures enum value "active"
	EdgeGroupStateActive string = "active"

	// EdgeGroupStateInactive captures enum value "inactive"
	EdgeGroupStateInactive string = "inactive"

	// EdgeGroupStateDeleted captures enum value "deleted"
	EdgeGroupStateDeleted string = "deleted"
)

// prop value enum
func (m *EdgeGroup) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgeGroupTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgeGroup) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this edge group based on the context it is used
func (m *EdgeGroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedByApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateModified(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDivision(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEdgeTrunkBaseAssignment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedByApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePhoneTrunkBases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeGroup) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidateCreatedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdByApp", "body", string(m.CreatedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EdgeGroup) contextValidateEdgeTrunkBaseAssignment(ctx context.Context, formats strfmt.Registry) error {

	if m.EdgeTrunkBaseAssignment != nil {
		if err := m.EdgeTrunkBaseAssignment.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeTrunkBaseAssignment")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("edgeTrunkBaseAssignment")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeGroup) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedBy", "body", string(m.ModifiedBy)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidateModifiedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedByApp", "body", string(m.ModifiedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidatePhoneTrunkBases(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PhoneTrunkBases); i++ {

		if m.PhoneTrunkBases[i] != nil {
			if err := m.PhoneTrunkBases[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("phoneTrunkBases" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("phoneTrunkBases" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *EdgeGroup) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeGroup) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeGroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeGroup) UnmarshalBinary(b []byte) error {
	var res EdgeGroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
