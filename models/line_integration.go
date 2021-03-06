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

// LineIntegration line integration
//
// swagger:model LineIntegration
type LineIntegration struct {

	// The Channel Id from LINE messenger
	// Required: true
	ChannelID *string `json:"channelId"`

	// Error information returned, if createStatus is set to Error
	// Read Only: true
	CreateError *ErrorBody `json:"createError,omitempty"`

	// Status of asynchronous create operation
	// Read Only: true
	// Enum: [Initiated Completed Error]
	CreateStatus string `json:"createStatus,omitempty"`

	// User reference that created this Integration
	CreatedBy *DomainEntityRef `json:"createdBy,omitempty"`

	// Date this Integration was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date this Integration was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// A unique Integration Id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// User reference that last modified this Integration
	ModifiedBy *DomainEntityRef `json:"modifiedBy,omitempty"`

	// The name of the LINE Integration
	// Required: true
	Name *string `json:"name"`

	// The recipient associated to the Line Integration. This recipient is used to associate a flow to an integration
	// Read Only: true
	Recipient *DomainEntityRef `json:"recipient,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of the LINE Integration
	Status string `json:"status,omitempty"`

	// Version number required for updates.
	// Required: true
	Version *int32 `json:"version"`

	// The Webhook URI to be updated in LINE platform
	// Required: true
	// Format: uri
	WebhookURI *strfmt.URI `json:"webhookUri"`
}

// Validate validates this line integration
func (m *LineIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChannelID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateError(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWebhookURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LineIntegration) validateChannelID(formats strfmt.Registry) error {

	if err := validate.Required("channelId", "body", m.ChannelID); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateCreateError(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateError) { // not required
		return nil
	}

	if m.CreateError != nil {
		if err := m.CreateError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createError")
			}
			return err
		}
	}

	return nil
}

var lineIntegrationTypeCreateStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Initiated","Completed","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lineIntegrationTypeCreateStatusPropEnum = append(lineIntegrationTypeCreateStatusPropEnum, v)
	}
}

const (

	// LineIntegrationCreateStatusInitiated captures enum value "Initiated"
	LineIntegrationCreateStatusInitiated string = "Initiated"

	// LineIntegrationCreateStatusCompleted captures enum value "Completed"
	LineIntegrationCreateStatusCompleted string = "Completed"

	// LineIntegrationCreateStatusError captures enum value "Error"
	LineIntegrationCreateStatusError string = "Error"
)

// prop value enum
func (m *LineIntegration) validateCreateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, lineIntegrationTypeCreateStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *LineIntegration) validateCreateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCreateStatusEnum("createStatus", "body", m.CreateStatus); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *LineIntegration) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *LineIntegration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateRecipient(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipient) { // not required
		return nil
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *LineIntegration) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

func (m *LineIntegration) validateWebhookURI(formats strfmt.Registry) error {

	if err := validate.Required("webhookUri", "body", m.WebhookURI); err != nil {
		return err
	}

	if err := validate.FormatOf("webhookUri", "body", "uri", m.WebhookURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LineIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LineIntegration) UnmarshalBinary(b []byte) error {
	var res LineIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
