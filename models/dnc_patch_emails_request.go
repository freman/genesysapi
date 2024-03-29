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

// DncPatchEmailsRequest dnc patch emails request
//
// swagger:model DncPatchEmailsRequest
type DncPatchEmailsRequest struct {

	// The action to perform
	// Enum: [Add Remove]
	Action string `json:"action,omitempty"`

	// The list of email addresses to Add to / Remove from the DNC list
	EmailAddresses []string `json:"emailAddresses"`

	// Expiration date for DNC email addresses in yyyy-MM-ddTHH:mmZ format
	ExpirationDateTime string `json:"expirationDateTime,omitempty"`
}

// Validate validates this dnc patch emails request
func (m *DncPatchEmailsRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAction(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dncPatchEmailsRequestTypeActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Add","Remove"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dncPatchEmailsRequestTypeActionPropEnum = append(dncPatchEmailsRequestTypeActionPropEnum, v)
	}
}

const (

	// DncPatchEmailsRequestActionAdd captures enum value "Add"
	DncPatchEmailsRequestActionAdd string = "Add"

	// DncPatchEmailsRequestActionRemove captures enum value "Remove"
	DncPatchEmailsRequestActionRemove string = "Remove"
)

// prop value enum
func (m *DncPatchEmailsRequest) validateActionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dncPatchEmailsRequestTypeActionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DncPatchEmailsRequest) validateAction(formats strfmt.Registry) error {
	if swag.IsZero(m.Action) { // not required
		return nil
	}

	// value enum
	if err := m.validateActionEnum("action", "body", m.Action); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this dnc patch emails request based on context it is used
func (m *DncPatchEmailsRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DncPatchEmailsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DncPatchEmailsRequest) UnmarshalBinary(b []byte) error {
	var res DncPatchEmailsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
