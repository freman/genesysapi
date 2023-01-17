// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PhonesReboot phones reboot
//
// swagger:model PhonesReboot
type PhonesReboot struct {

	// The list of phone Ids to reboot.
	PhoneIds []string `json:"phoneIds"`

	// ID of the site for which to reboot all phones at that site.
	// no.active.edge and phone.cannot.resolve errors are ignored.
	SiteID string `json:"siteId,omitempty"`
}

// Validate validates this phones reboot
func (m *PhonesReboot) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this phones reboot based on context it is used
func (m *PhonesReboot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhonesReboot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhonesReboot) UnmarshalBinary(b []byte) error {
	var res PhonesReboot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
