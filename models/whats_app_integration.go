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

// WhatsAppIntegration whats app integration
//
// swagger:model WhatsAppIntegration
type WhatsAppIntegration struct {

	// The error information of WhatsApp Integration activation process
	// Read Only: true
	ActivationErrorInfo *ErrorBody `json:"activationErrorInfo,omitempty"`

	// The status code of WhatsApp Integration activation process
	// Read Only: true
	// Enum: [CodeSent WaitRequired ActivationFailed CodeConfirmed ConfirmationFailed ResendCode]
	ActivationStatusCode string `json:"activationStatusCode,omitempty"`

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

	// Date this Integration was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
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

	// The name of the WhatsApp integration.
	// Required: true
	Name *string `json:"name"`

	// The phone number associated to the whatsApp integration.
	// Required: true
	PhoneNumber *string `json:"phoneNumber"`

	// The recipient associated to the WhatsApp Integration. This recipient is used to associate a flow to an integration
	// Read Only: true
	Recipient *DomainEntityRef `json:"recipient,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The status of the WhatsApp Integration
	// Enum: [Active Inactive Error Starting Incomplete Deleting DeletionFailed]
	Status string `json:"status,omitempty"`

	// Defines the SupportedContent profile configured for an integration
	SupportedContent *SupportedContentReference `json:"supportedContent,omitempty"`

	// Version number required for updates.
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this whats app integration
func (m *WhatsAppIntegration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActivationErrorInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateActivationStatusCode(formats); err != nil {
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

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipient(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *WhatsAppIntegration) validateActivationErrorInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivationErrorInfo) { // not required
		return nil
	}

	if m.ActivationErrorInfo != nil {
		if err := m.ActivationErrorInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("activationErrorInfo")
			}
			return err
		}
	}

	return nil
}

var whatsAppIntegrationTypeActivationStatusCodePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CodeSent","WaitRequired","ActivationFailed","CodeConfirmed","ConfirmationFailed","ResendCode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		whatsAppIntegrationTypeActivationStatusCodePropEnum = append(whatsAppIntegrationTypeActivationStatusCodePropEnum, v)
	}
}

const (

	// WhatsAppIntegrationActivationStatusCodeCodeSent captures enum value "CodeSent"
	WhatsAppIntegrationActivationStatusCodeCodeSent string = "CodeSent"

	// WhatsAppIntegrationActivationStatusCodeWaitRequired captures enum value "WaitRequired"
	WhatsAppIntegrationActivationStatusCodeWaitRequired string = "WaitRequired"

	// WhatsAppIntegrationActivationStatusCodeActivationFailed captures enum value "ActivationFailed"
	WhatsAppIntegrationActivationStatusCodeActivationFailed string = "ActivationFailed"

	// WhatsAppIntegrationActivationStatusCodeCodeConfirmed captures enum value "CodeConfirmed"
	WhatsAppIntegrationActivationStatusCodeCodeConfirmed string = "CodeConfirmed"

	// WhatsAppIntegrationActivationStatusCodeConfirmationFailed captures enum value "ConfirmationFailed"
	WhatsAppIntegrationActivationStatusCodeConfirmationFailed string = "ConfirmationFailed"

	// WhatsAppIntegrationActivationStatusCodeResendCode captures enum value "ResendCode"
	WhatsAppIntegrationActivationStatusCodeResendCode string = "ResendCode"
)

// prop value enum
func (m *WhatsAppIntegration) validateActivationStatusCodeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, whatsAppIntegrationTypeActivationStatusCodePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WhatsAppIntegration) validateActivationStatusCode(formats strfmt.Registry) error {

	if swag.IsZero(m.ActivationStatusCode) { // not required
		return nil
	}

	// value enum
	if err := m.validateActivationStatusCodeEnum("activationStatusCode", "body", m.ActivationStatusCode); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validateCreateError(formats strfmt.Registry) error {

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

var whatsAppIntegrationTypeCreateStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Initiated","Completed","Error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		whatsAppIntegrationTypeCreateStatusPropEnum = append(whatsAppIntegrationTypeCreateStatusPropEnum, v)
	}
}

const (

	// WhatsAppIntegrationCreateStatusInitiated captures enum value "Initiated"
	WhatsAppIntegrationCreateStatusInitiated string = "Initiated"

	// WhatsAppIntegrationCreateStatusCompleted captures enum value "Completed"
	WhatsAppIntegrationCreateStatusCompleted string = "Completed"

	// WhatsAppIntegrationCreateStatusError captures enum value "Error"
	WhatsAppIntegrationCreateStatusError string = "Error"
)

// prop value enum
func (m *WhatsAppIntegration) validateCreateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, whatsAppIntegrationTypeCreateStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WhatsAppIntegration) validateCreateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CreateStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCreateStatusEnum("createStatus", "body", m.CreateStatus); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validateCreatedBy(formats strfmt.Registry) error {

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

func (m *WhatsAppIntegration) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validateMessagingSetting(formats strfmt.Registry) error {

	if swag.IsZero(m.MessagingSetting) { // not required
		return nil
	}

	if m.MessagingSetting != nil {
		if err := m.MessagingSetting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("messagingSetting")
			}
			return err
		}
	}

	return nil
}

func (m *WhatsAppIntegration) validateModifiedBy(formats strfmt.Registry) error {

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

func (m *WhatsAppIntegration) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phoneNumber", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validateRecipient(formats strfmt.Registry) error {

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

func (m *WhatsAppIntegration) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var whatsAppIntegrationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Active","Inactive","Error","Starting","Incomplete","Deleting","DeletionFailed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		whatsAppIntegrationTypeStatusPropEnum = append(whatsAppIntegrationTypeStatusPropEnum, v)
	}
}

const (

	// WhatsAppIntegrationStatusActive captures enum value "Active"
	WhatsAppIntegrationStatusActive string = "Active"

	// WhatsAppIntegrationStatusInactive captures enum value "Inactive"
	WhatsAppIntegrationStatusInactive string = "Inactive"

	// WhatsAppIntegrationStatusError captures enum value "Error"
	WhatsAppIntegrationStatusError string = "Error"

	// WhatsAppIntegrationStatusStarting captures enum value "Starting"
	WhatsAppIntegrationStatusStarting string = "Starting"

	// WhatsAppIntegrationStatusIncomplete captures enum value "Incomplete"
	WhatsAppIntegrationStatusIncomplete string = "Incomplete"

	// WhatsAppIntegrationStatusDeleting captures enum value "Deleting"
	WhatsAppIntegrationStatusDeleting string = "Deleting"

	// WhatsAppIntegrationStatusDeletionFailed captures enum value "DeletionFailed"
	WhatsAppIntegrationStatusDeletionFailed string = "DeletionFailed"
)

// prop value enum
func (m *WhatsAppIntegration) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, whatsAppIntegrationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WhatsAppIntegration) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WhatsAppIntegration) validateSupportedContent(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedContent) { // not required
		return nil
	}

	if m.SupportedContent != nil {
		if err := m.SupportedContent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("supportedContent")
			}
			return err
		}
	}

	return nil
}

func (m *WhatsAppIntegration) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WhatsAppIntegration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WhatsAppIntegration) UnmarshalBinary(b []byte) error {
	var res WhatsAppIntegration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
