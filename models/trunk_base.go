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

// TrunkBase trunk base
//
// swagger:model TrunkBase
type TrunkBase struct {

	// The ID of the user that created the resource.
	CreatedBy string `json:"createdBy,omitempty"`

	// The application that created the resource.
	CreatedByApp string `json:"createdByApp,omitempty"`

	// The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The resource's description.
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Is this trunk being managed remotely. This property is synchronized with the managed property of the Edge Group to which it is assigned.
	Managed bool `json:"managed"`

	// The ID of the user that last modified the resource.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
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

// Validate validates this trunk base
func (m *TrunkBase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
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

func (m *TrunkBase) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrunkBase) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrunkBase) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *TrunkBase) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var trunkBaseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trunkBaseTypeStatePropEnum = append(trunkBaseTypeStatePropEnum, v)
	}
}

const (

	// TrunkBaseStateActive captures enum value "active"
	TrunkBaseStateActive string = "active"

	// TrunkBaseStateInactive captures enum value "inactive"
	TrunkBaseStateInactive string = "inactive"

	// TrunkBaseStateDeleted captures enum value "deleted"
	TrunkBaseStateDeleted string = "deleted"
)

// prop value enum
func (m *TrunkBase) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trunkBaseTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TrunkBase) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

func (m *TrunkBase) validateTrunkMetabase(formats strfmt.Registry) error {

	if err := validate.Required("trunkMetabase", "body", m.TrunkMetabase); err != nil {
		return err
	}

	if m.TrunkMetabase != nil {
		if err := m.TrunkMetabase.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trunkMetabase")
			}
			return err
		}
	}

	return nil
}

var trunkBaseTypeTrunkTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXTERNAL","PHONE","EDGE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		trunkBaseTypeTrunkTypePropEnum = append(trunkBaseTypeTrunkTypePropEnum, v)
	}
}

const (

	// TrunkBaseTrunkTypeEXTERNAL captures enum value "EXTERNAL"
	TrunkBaseTrunkTypeEXTERNAL string = "EXTERNAL"

	// TrunkBaseTrunkTypePHONE captures enum value "PHONE"
	TrunkBaseTrunkTypePHONE string = "PHONE"

	// TrunkBaseTrunkTypeEDGE captures enum value "EDGE"
	TrunkBaseTrunkTypeEDGE string = "EDGE"
)

// prop value enum
func (m *TrunkBase) validateTrunkTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, trunkBaseTypeTrunkTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *TrunkBase) validateTrunkType(formats strfmt.Registry) error {

	if err := validate.Required("trunkType", "body", m.TrunkType); err != nil {
		return err
	}

	// value enum
	if err := m.validateTrunkTypeEnum("trunkType", "body", *m.TrunkType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrunkBase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrunkBase) UnmarshalBinary(b []byte) error {
	var res TrunkBase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
