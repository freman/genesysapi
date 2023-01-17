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

// Library library
//
// swagger:model Library
type Library struct {

	// User that created the library.
	// Read Only: true
	CreatedBy *User `json:"createdBy,omitempty"`

	// The date and time the response was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Read Only: true
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The library name.
	// Required: true
	Name *string `json:"name"`

	// This value is deprecated. Responses representing message templates may be added to any library.
	// Enum: [MessagingTemplate CampaignSmsTemplate CampaignEmailTemplate Footer]
	ResponseType string `json:"responseType,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Current version for this resource.
	// Read Only: true
	Version int32 `json:"version,omitempty"`
}

// Validate validates this library
func (m *Library) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResponseType(formats); err != nil {
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

func (m *Library) validateCreatedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.CreatedBy) { // not required
		return nil
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

func (m *Library) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Library) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var libraryTypeResponseTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["MessagingTemplate","CampaignSmsTemplate","CampaignEmailTemplate","Footer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		libraryTypeResponseTypePropEnum = append(libraryTypeResponseTypePropEnum, v)
	}
}

const (

	// LibraryResponseTypeMessagingTemplate captures enum value "MessagingTemplate"
	LibraryResponseTypeMessagingTemplate string = "MessagingTemplate"

	// LibraryResponseTypeCampaignSmsTemplate captures enum value "CampaignSmsTemplate"
	LibraryResponseTypeCampaignSmsTemplate string = "CampaignSmsTemplate"

	// LibraryResponseTypeCampaignEmailTemplate captures enum value "CampaignEmailTemplate"
	LibraryResponseTypeCampaignEmailTemplate string = "CampaignEmailTemplate"

	// LibraryResponseTypeFooter captures enum value "Footer"
	LibraryResponseTypeFooter string = "Footer"
)

// prop value enum
func (m *Library) validateResponseTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, libraryTypeResponseTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Library) validateResponseType(formats strfmt.Registry) error {
	if swag.IsZero(m.ResponseType) { // not required
		return nil
	}

	// value enum
	if err := m.validateResponseTypeEnum("responseType", "body", m.ResponseType); err != nil {
		return err
	}

	return nil
}

func (m *Library) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this library based on the context it is used
func (m *Library) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Library) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Library) contextValidateDateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "dateCreated", "body", strfmt.DateTime(m.DateCreated)); err != nil {
		return err
	}

	return nil
}

func (m *Library) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *Library) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *Library) contextValidateVersion(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "version", "body", int32(m.Version)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Library) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Library) UnmarshalBinary(b []byte) error {
	var res Library
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
