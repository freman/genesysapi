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

// EdgeTrunkBase edge trunk base
//
// swagger:model EdgeTrunkBase
type EdgeTrunkBase struct {

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

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The ID of the user that last modified the resource.
	// Read Only: true
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
	// Read Only: true
	ModifiedByApp string `json:"modifiedByApp,omitempty"`

	// The name of the entity.
	// Required: true
	Name *string `json:"name"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Indicates if the resource is active, inactive, or deleted.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// The meta-base this trunk is based on.
	// Required: true
	TrunkMetabase *DomainEntityRef `json:"trunkMetabase"`

	// The type of this trunk base.
	// Required: true
	// Enum: [EXTERNAL PHONE EDGE]
	TrunkType *string `json:"trunkType"`

	// The current version of the resource.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this edge trunk base
func (m *EdgeTrunkBase) Validate(formats strfmt.Registry) error {
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

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrunkMetabase(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrunkType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeTrunkBase) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) validateDivision(formats strfmt.Registry) error {
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

func (m *EdgeTrunkBase) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var edgeTrunkBaseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgeTrunkBaseTypeStatePropEnum = append(edgeTrunkBaseTypeStatePropEnum, v)
	}
}

const (

	// EdgeTrunkBaseStateActive captures enum value "active"
	EdgeTrunkBaseStateActive string = "active"

	// EdgeTrunkBaseStateInactive captures enum value "inactive"
	EdgeTrunkBaseStateInactive string = "inactive"

	// EdgeTrunkBaseStateDeleted captures enum value "deleted"
	EdgeTrunkBaseStateDeleted string = "deleted"
)

// prop value enum
func (m *EdgeTrunkBase) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgeTrunkBaseTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgeTrunkBase) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) validateTrunkMetabase(formats strfmt.Registry) error {

	if err := validate.Required("trunkMetabase", "body", m.TrunkMetabase); err != nil {
		return err
	}

	if m.TrunkMetabase != nil {
		if err := m.TrunkMetabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunkMetabase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trunkMetabase")
			}
			return err
		}
	}

	return nil
}

var edgeTrunkBaseTypeTrunkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXTERNAL","PHONE","EDGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgeTrunkBaseTypeTrunkTypePropEnum = append(edgeTrunkBaseTypeTrunkTypePropEnum, v)
	}
}

const (

	// EdgeTrunkBaseTrunkTypeEXTERNAL captures enum value "EXTERNAL"
	EdgeTrunkBaseTrunkTypeEXTERNAL string = "EXTERNAL"

	// EdgeTrunkBaseTrunkTypePHONE captures enum value "PHONE"
	EdgeTrunkBaseTrunkTypePHONE string = "PHONE"

	// EdgeTrunkBaseTrunkTypeEDGE captures enum value "EDGE"
	EdgeTrunkBaseTrunkTypeEDGE string = "EDGE"
)

// prop value enum
func (m *EdgeTrunkBase) validateTrunkTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgeTrunkBaseTypeTrunkTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgeTrunkBase) validateTrunkType(formats strfmt.Registry) error {

	if err := validate.Required("trunkType", "body", m.TrunkType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTrunkTypeEnum("trunkType", "body", *m.TrunkType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this edge trunk base based on the context it is used
func (m *EdgeTrunkBase) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedByApp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTrunkMetabase(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeTrunkBase) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateCreatedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdByApp", "body", string(m.CreatedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *EdgeTrunkBase) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedBy", "body", string(m.ModifiedBy)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateModifiedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedByApp", "body", string(m.ModifiedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

func (m *EdgeTrunkBase) contextValidateTrunkMetabase(ctx context.Context, formats strfmt.Registry) error {

	if m.TrunkMetabase != nil {
		if err := m.TrunkMetabase.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunkMetabase")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trunkMetabase")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeTrunkBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeTrunkBase) UnmarshalBinary(b []byte) error {
	var res EdgeTrunkBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
