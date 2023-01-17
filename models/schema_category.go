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

// SchemaCategory schema category
//
// swagger:model SchemaCategory
type SchemaCategory struct {

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

// Validate validates this schema category
func (m *SchemaCategory) Validate(formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemaCategory) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) validateDivision(formats strfmt.Registry) error {
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

func (m *SchemaCategory) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var schemaCategoryTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		schemaCategoryTypeStatePropEnum = append(schemaCategoryTypeStatePropEnum, v)
	}
}

const (

	// SchemaCategoryStateActive captures enum value "active"
	SchemaCategoryStateActive string = "active"

	// SchemaCategoryStateInactive captures enum value "inactive"
	SchemaCategoryStateInactive string = "inactive"

	// SchemaCategoryStateDeleted captures enum value "deleted"
	SchemaCategoryStateDeleted string = "deleted"
)

// prop value enum
func (m *SchemaCategory) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, schemaCategoryTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SchemaCategory) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this schema category based on the context it is used
func (m *SchemaCategory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SchemaCategory) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdBy", "body", string(m.CreatedBy)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateCreatedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createdByApp", "body", string(m.CreatedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateDateModified(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateModified", "body", strfmt.DateTime(m.DateModified)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateDivision(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SchemaCategory) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedBy", "body", string(m.ModifiedBy)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateModifiedByApp(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "modifiedByApp", "body", string(m.ModifiedByApp)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *SchemaCategory) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "state", "body", string(m.State)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SchemaCategory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SchemaCategory) UnmarshalBinary(b []byte) error {
	var res SchemaCategory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
