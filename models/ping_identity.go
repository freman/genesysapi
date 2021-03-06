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

// PingIdentity ping identity
//
// swagger:model PingIdentity
type PingIdentity struct {

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// certificates
	Certificates []string `json:"certificates"`

	// disabled
	Disabled bool `json:"disabled"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// issuer URI
	IssuerURI string `json:"issuerURI,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// relying party identifier
	RelyingPartyIdentifier string `json:"relyingPartyIdentifier,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// sso target URI
	SsoTargetURI string `json:"ssoTargetURI,omitempty"`
}

// Validate validates this ping identity
func (m *PingIdentity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PingIdentity) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PingIdentity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PingIdentity) UnmarshalBinary(b []byte) error {
	var res PingIdentity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
