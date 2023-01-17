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

// FacebookIntegration facebook integration
//
// swagger:model FacebookIntegration
type FacebookIntegration struct {

	// The App Id from Facebook messenger
	// Required: true
	AppID *string `json:"appId"`

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

	// A unique Integration Id.
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// messaging setting
	MessagingSetting *MessagingSettingReference `json:"messagingSetting,omitempty"`

	// User reference that last modified this Integration
	ModifiedBy *DomainEntityRef `json:"modifiedBy,omitempty"`

	// The name of the Facebook Integration
	// Required: true
	Name *string `json:"name"`

	// The Page Id from Facebook messenger
	PageID string `json:"pageId,omitempty"`

	// The name of the Facebook page
	// Read Only: true
	PageName string `json:"pageName,omitempty"`

	// The url of the profile image of the Facebook page
	// Read Only: true
	PageProfileImageURL string `json:"pageProfileImageUrl,omitempty"`

	// The recipient reference associated to the Facebook Integration. This recipient is used to associate a flow to an integration
	// Read Only: true
	Recipient *DomainEntityRef `json:"recipient,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of the Facebook Integration
	Status string `json:"status,omitempty"`

	// Defines the SupportedContent profile configured for an integration
	SupportedContent *SupportedContentReference `json:"supportedContent,omitempty"`

	// Version number required for updates.
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this facebook integration
func (m *FacebookIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppID(formats); err != nil {
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

	if err := m.validateMessagingSetting(formats); err != nil {
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

	if err := m.validateSupportedContent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacebookIntegration) validateAppID(formats strfmt.Registry) error {

	if err := validate.Required("appId", "body", m.AppID); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) validateCreateError(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateError) { // not required
		return nil
	}

	if m.CreateError != nil {
		if err := m.CreateError.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createError")
			}
			return err
		}
	}

	return nil
}

var facebookIntegrationTypeCreateStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Initiated","Completed","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		facebookIntegrationTypeCreateStatusPropEnum = append(facebookIntegrationTypeCreateStatusPropEnum, v)
	}
}

const (

	// FacebookIntegrationCreateStatusInitiated captures enum value "Initiated"
	FacebookIntegrationCreateStatusInitiated string = "Initiated"

	// FacebookIntegrationCreateStatusCompleted captures enum value "Completed"
	FacebookIntegrationCreateStatusCompleted string = "Completed"

	// FacebookIntegrationCreateStatusError captures enum value "Error"
	FacebookIntegrationCreateStatusError string = "Error"
)

// prop value enum
func (m *FacebookIntegration) validateCreateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, facebookIntegrationTypeCreateStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FacebookIntegration) validateCreateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCreateStatusEnum("createStatus", "body", m.CreateStatus); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) validateCreatedBy(formats strfmt.Registry) error {
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

func (m *FacebookIntegration) validateDateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) validateDateModified(formats strfmt.Registry) error {
	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) validateMessagingSetting(formats strfmt.Registry) error {
	if swag.IsZero(m.MessagingSetting) { // not required
		return nil
	}

	if m.MessagingSetting != nil {
		if err := m.MessagingSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingSetting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("messagingSetting")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) validateModifiedBy(formats strfmt.Registry) error {
	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) validateRecipient(formats strfmt.Registry) error {
	if swag.IsZero(m.Recipient) { // not required
		return nil
	}

	if m.Recipient != nil {
		if err := m.Recipient.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) validateSupportedContent(formats strfmt.Registry) error {
	if swag.IsZero(m.SupportedContent) { // not required
		return nil
	}

	if m.SupportedContent != nil {
		if err := m.SupportedContent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supportedContent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supportedContent")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this facebook integration based on the context it is used
func (m *FacebookIntegration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreateError(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreatedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessagingSetting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModifiedBy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePageName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePageProfileImageURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRecipient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSupportedContent(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FacebookIntegration) contextValidateCreateError(ctx context.Context, formats strfmt.Registry) error {

	if m.CreateError != nil {
		if err := m.CreateError.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createError")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createError")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) contextValidateCreateStatus(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "createStatus", "body", string(m.CreateStatus)); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) contextValidateCreatedBy(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FacebookIntegration) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) contextValidateMessagingSetting(ctx context.Context, formats strfmt.Registry) error {

	if m.MessagingSetting != nil {
		if err := m.MessagingSetting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingSetting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("messagingSetting")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) contextValidateModifiedBy(ctx context.Context, formats strfmt.Registry) error {

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) contextValidatePageName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "pageName", "body", string(m.PageName)); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) contextValidatePageProfileImageURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "pageProfileImageUrl", "body", string(m.PageProfileImageURL)); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) contextValidateRecipient(ctx context.Context, formats strfmt.Registry) error {

	if m.Recipient != nil {
		if err := m.Recipient.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recipient")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("recipient")
			}
			return err
		}
	}

	return nil
}

func (m *FacebookIntegration) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

func (m *FacebookIntegration) contextValidateSupportedContent(ctx context.Context, formats strfmt.Registry) error {

	if m.SupportedContent != nil {
		if err := m.SupportedContent.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supportedContent")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("supportedContent")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FacebookIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FacebookIntegration) UnmarshalBinary(b []byte) error {
	var res FacebookIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
