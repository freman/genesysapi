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

// DncPatchPhoneNumbersRequest dnc patch phone numbers request
//
// swagger:model DncPatchPhoneNumbersRequest
type DncPatchPhoneNumbersRequest struct {

	// The action to perform
	// Enum: [Add Remove]
	Action string `json:"action,omitempty"`

	// Expiration date for DNC phone numbers in yyyy-MM-ddTHH:mmZ format
	ExpirationDateTime string `json:"expirationDateTime,omitempty"`

	// The list of phone numbers to Add to / Remove from the DNC list
	PhoneNumbers []string `json:"phoneNumbers"`
}

// Validate validates this dnc patch phone numbers request
func (m *DncPatchPhoneNumbersRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dncPatchPhoneNumbersRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Add","Remove"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dncPatchPhoneNumbersRequestTypeActionPropEnum = append(dncPatchPhoneNumbersRequestTypeActionPropEnum, v)
	}
}

const (

	// DncPatchPhoneNumbersRequestActionAdd captures enum value "Add"
	DncPatchPhoneNumbersRequestActionAdd string = "Add"

	// DncPatchPhoneNumbersRequestActionRemove captures enum value "Remove"
	DncPatchPhoneNumbersRequestActionRemove string = "Remove"
)

// prop value enum
func (m *DncPatchPhoneNumbersRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dncPatchPhoneNumbersRequestTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DncPatchPhoneNumbersRequest) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dnc patch phone numbers request based on context it is used
func (m *DncPatchPhoneNumbersRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DncPatchPhoneNumbersRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DncPatchPhoneNumbersRequest) UnmarshalBinary(b []byte) error {
	var res DncPatchPhoneNumbersRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
