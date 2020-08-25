// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PureEngage pure engage
//
// swagger:model PureEngage
type PureEngage struct {

	// auto provision users
	AutoProvisionUsers bool `json:"autoProvisionUsers"`

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// disabled
	Disabled bool `json:"disabled"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// issuer URI
	IssuerURI string `json:"issuerURI,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// sso target URI
	SsoTargetURI string `json:"ssoTargetURI,omitempty"`
}

// Validate validates this pure engage
func (m *PureEngage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PureEngage) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PureEngage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PureEngage) UnmarshalBinary(b []byte) error {
	var res PureEngage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}