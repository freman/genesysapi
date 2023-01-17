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

// UserDevice user device
//
// swagger:model UserDevice
type UserDevice struct {

	// if the device accepts notifications
	// Required: true
	AcceptNotifications *bool `json:"acceptNotifications"`

	// device token sent by mobile clients.
	// Required: true
	DeviceToken *string `json:"deviceToken"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// make of the device.
	// Required: true
	Make *string `json:"make"`

	// Device model
	// Required: true
	Model *string `json:"model"`

	// name
	Name string `json:"name,omitempty"`

	// notification id of the device.
	// Required: true
	NotificationID *string `json:"notificationId"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// session hash
	SessionHash string `json:"sessionHash,omitempty"`

	// type of the device; ios or android
	// Required: true
	// Enum: [android ios]
	Type *string `json:"type"`
}

// Validate validates this user device
func (m *UserDevice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcceptNotifications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeviceToken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotificationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserDevice) validateAcceptNotifications(formats strfmt.Registry) error {

	if err := validate.Required("acceptNotifications", "body", m.AcceptNotifications); err != nil {
		return err
	}

	return nil
}

func (m *UserDevice) validateDeviceToken(formats strfmt.Registry) error {

	if err := validate.Required("deviceToken", "body", m.DeviceToken); err != nil {
		return err
	}

	return nil
}

func (m *UserDevice) validateMake(formats strfmt.Registry) error {

	if err := validate.Required("make", "body", m.Make); err != nil {
		return err
	}

	return nil
}

func (m *UserDevice) validateModel(formats strfmt.Registry) error {

	if err := validate.Required("model", "body", m.Model); err != nil {
		return err
	}

	return nil
}

func (m *UserDevice) validateNotificationID(formats strfmt.Registry) error {

	if err := validate.Required("notificationId", "body", m.NotificationID); err != nil {
		return err
	}

	return nil
}

func (m *UserDevice) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var userDeviceTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["android","ios"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userDeviceTypeTypePropEnum = append(userDeviceTypeTypePropEnum, v)
	}
}

const (

	// UserDeviceTypeAndroid captures enum value "android"
	UserDeviceTypeAndroid string = "android"

	// UserDeviceTypeIos captures enum value "ios"
	UserDeviceTypeIos string = "ios"
)

// prop value enum
func (m *UserDevice) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, userDeviceTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UserDevice) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this user device based on the context it is used
func (m *UserDevice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

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

func (m *UserDevice) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *UserDevice) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserDevice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDevice) UnmarshalBinary(b []byte) error {
	var res UserDevice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
