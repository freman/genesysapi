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

// Note note
//
// swagger:model Note
type Note struct {

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreateDate strfmt.DateTime `json:"createDate,omitempty"`

	// When creating or updating a note, only User.id is required. User object is fully populated when expanding a note.
	// Required: true
	CreatedBy *User `json:"createdBy"`

	// The id of the contact or organization to which this note refers. This only needs to be set for input when using the Bulk APIs.
	EntityID string `json:"entityId,omitempty"`

	// This is only need to be set when using Bulk API. Using any other value than contact or organization will result in null being used.
	// Enum: [contact organization]
	EntityType string `json:"entityType,omitempty"`

	// Links to the sources of data (e.g. one source might be a CRM) that contributed data to this record.  Read-only, and only populated when requested via expand param.
	// Read Only: true
	ExternalDataSources []*ExternalDataSource `json:"externalDataSources"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ModifyDate strfmt.DateTime `json:"modifyDate,omitempty"`

	// note text
	NoteText string `json:"noteText,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`
}

// Validate validates this note
func (m *Note) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalDataSources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifyDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Note) validateCreateDate(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createDate", "body", "date-time", m.CreateDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Note) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

var noteTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["contact","organization"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		noteTypeEntityTypePropEnum = append(noteTypeEntityTypePropEnum, v)
	}
}

const (

	// NoteEntityTypeContact captures enum value "contact"
	NoteEntityTypeContact string = "contact"

	// NoteEntityTypeOrganization captures enum value "organization"
	NoteEntityTypeOrganization string = "organization"
)

// prop value enum
func (m *Note) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, noteTypeEntityTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Note) validateEntityType(formats strfmt.Registry) error {
	if swag.IsZero(m.EntityType) { // not required
		return nil
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *Note) validateExternalDataSources(formats strfmt.Registry) error {
	if swag.IsZero(m.ExternalDataSources) { // not required
		return nil
	}

	for i := 0; i < len(m.ExternalDataSources); i++ {
		if swag.IsZero(m.ExternalDataSources[i]) { // not required
			continue
		}

		if m.ExternalDataSources[i] != nil {
			if err := m.ExternalDataSources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalDataSources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("externalDataSources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Note) validateModifyDate(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifyDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifyDate", "body", "date-time", m.ModifyDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Note) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this note based on the context it is used
func (m *Note) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExternalDataSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
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

func (m *Note) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.CreatedBy != nil {
		if err := m.CreatedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *Note) contextValidateExternalDataSources(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "externalDataSources", "body", []*ExternalDataSource(m.ExternalDataSources)); err != nil {
		return err
	}

	for i := 0; i < len(m.ExternalDataSources); i++ {

		if m.ExternalDataSources[i] != nil {
			if err := m.ExternalDataSources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("externalDataSources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("externalDataSources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Note) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Note) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Note) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Note) UnmarshalBinary(b []byte) error {
	var res Note
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
