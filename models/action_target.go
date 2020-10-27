// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ActionTarget action target
//
// swagger:model ActionTarget
type ActionTarget struct {

	// The date the target was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CreatedDate strfmt.DateTime `json:"createdDate,omitempty"`

	// Description of the target.
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The date the target was last modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	ModifiedDate strfmt.DateTime `json:"modifiedDate,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Service Level of the action target. Chat offers for the target will be throttled with the aim of achieving this service level.
	ServiceLevel *ServiceLevel `json:"serviceLevel,omitempty"`

	// Indicates the state of the target.
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// Supported media types of the target.
	SupportedMediaTypes []string `json:"supportedMediaTypes"`

	// Additional user data associated with the target in key/value format.
	UserData []*KeyValue `json:"userData"`
}

// Validate validates this action target
func (m *ActionTarget) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedMediaTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionTarget) validateCreatedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("createdDate", "body", "date-time", m.CreatedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActionTarget) validateModifiedDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedDate) { // not required
		return nil
	}

	if err := validate.FormatOf("modifiedDate", "body", "date-time", m.ModifiedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActionTarget) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ActionTarget) validateServiceLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceLevel) { // not required
		return nil
	}

	if m.ServiceLevel != nil {
		if err := m.ServiceLevel.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceLevel")
			}
			return err
		}
	}

	return nil
}

var actionTargetTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionTargetTypeStatePropEnum = append(actionTargetTypeStatePropEnum, v)
	}
}

const (

	// ActionTargetStateActive captures enum value "active"
	ActionTargetStateActive string = "active"

	// ActionTargetStateInactive captures enum value "inactive"
	ActionTargetStateInactive string = "inactive"

	// ActionTargetStateDeleted captures enum value "deleted"
	ActionTargetStateDeleted string = "deleted"
)

// prop value enum
func (m *ActionTarget) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actionTargetTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActionTarget) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var actionTargetSupportedMediaTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["callback","call","email","generic","messaging","social","webchat"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		actionTargetSupportedMediaTypesItemsEnum = append(actionTargetSupportedMediaTypesItemsEnum, v)
	}
}

func (m *ActionTarget) validateSupportedMediaTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, actionTargetSupportedMediaTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ActionTarget) validateSupportedMediaTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedMediaTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.SupportedMediaTypes); i++ {

		// value enum
		if err := m.validateSupportedMediaTypesItemsEnum("supportedMediaTypes"+"."+strconv.Itoa(i), "body", m.SupportedMediaTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *ActionTarget) validateUserData(formats strfmt.Registry) error {

	if swag.IsZero(m.UserData) { // not required
		return nil
	}

	for i := 0; i < len(m.UserData); i++ {
		if swag.IsZero(m.UserData[i]) { // not required
			continue
		}

		if m.UserData[i] != nil {
			if err := m.UserData[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("userData" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionTarget) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionTarget) UnmarshalBinary(b []byte) error {
	var res ActionTarget
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
