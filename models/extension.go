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

// Extension extension
//
// swagger:model Extension
type Extension struct {

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

	// extension pool
	ExtensionPool *DomainEntityRef `json:"extensionPool,omitempty"`

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

	// number
	Number string `json:"number,omitempty"`

	// A Uri reference to the owner of this extension, which is either a User or an IVR
	Owner *DomainEntityRef `json:"owner,omitempty"`

	// owner type
	// Enum: [USER PHONE IVR_CONFIG GROUP]
	OwnerType string `json:"ownerType,omitempty"`

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

// Validate validates this extension
func (m *Extension) Validate(formats strfmt.Registry) error {
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

	if err := m.validateExtensionPool(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerType(formats); err != nil {
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

func (m *Extension) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Extension) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Extension) validateDivision(formats strfmt.Registry) error {
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

func (m *Extension) validateExtensionPool(formats strfmt.Registry) error {
	if swag.IsZero(m.ExtensionPool) { // not required
		return nil
	}

	if m.ExtensionPool != nil {
		if err := m.ExtensionPool.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensionPool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensionPool")
			}
			return err
		}
	}

	return nil
}

func (m *Extension) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Extension) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

var extensionTypeOwnerTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["USER","PHONE","IVR_CONFIG","GROUP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extensionTypeOwnerTypePropEnum = append(extensionTypeOwnerTypePropEnum, v)
	}
}

const (

	// ExtensionOwnerTypeUSER captures enum value "USER"
	ExtensionOwnerTypeUSER string = "USER"

	// ExtensionOwnerTypePHONE captures enum value "PHONE"
	ExtensionOwnerTypePHONE string = "PHONE"

	// ExtensionOwnerTypeIVRCONFIG captures enum value "IVR_CONFIG"
	ExtensionOwnerTypeIVRCONFIG string = "IVR_CONFIG"

	// ExtensionOwnerTypeGROUP captures enum value "GROUP"
	ExtensionOwnerTypeGROUP string = "GROUP"
)

// prop value enum
func (m *Extension) validateOwnerTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, extensionTypeOwnerTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Extension) validateOwnerType(formats strfmt.Registry) error {
	if swag.IsZero(m.OwnerType) { // not required
		return nil
	}

	// value enum
	if err := m.validateOwnerTypeEnum("ownerType", "body", m.OwnerType); err != nil {
		return err
	}

	return nil
}

func (m *Extension) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var extensionTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		extensionTypeStatePropEnum = append(extensionTypeStatePropEnum, v)
	}
}

const (

	// ExtensionStateActive captures enum value "active"
	ExtensionStateActive string = "active"

	// ExtensionStateInactive captures enum value "inactive"
	ExtensionStateInactive string = "inactive"

	// ExtensionStateDeleted captures enum value "deleted"
	ExtensionStateDeleted string = "deleted"
)

// prop value enum
func (m *Extension) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, extensionTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Extension) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this extension based on the context it is used
func (m *Extension) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateExtensionPool(ctx, formats); err != nil {
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

	if err := m.contextValidateOwner(ctx, formats); err != nil {
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

func (m *Extension) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateCreatedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdByApp", "body", string(m.CreatedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Extension) contextValidateExtensionPool(ctx context.Context, formats strfmt.Registry) error {

	if m.ExtensionPool != nil {
		if err := m.ExtensionPool.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensionPool")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensionPool")
			}
			return err
		}
	}

	return nil
}

func (m *Extension) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedBy", "body", string(m.ModifiedBy)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateModifiedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedByApp", "body", string(m.ModifiedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {
		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *Extension) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Extension) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Extension) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Extension) UnmarshalBinary(b []byte) error {
	var res Extension
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
