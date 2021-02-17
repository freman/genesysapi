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

// GenericSAML generic s a m l
//
// swagger:model GenericSAML
type GenericSAML struct {

	// certificate
	Certificate string `json:"certificate,omitempty"`

	// certificates
	Certificates []string `json:"certificates"`

	// disabled
	Disabled bool `json:"disabled"`

	// endpoint compression
	EndpointCompression bool `json:"endpointCompression"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// issuer URI
	IssuerURI string `json:"issuerURI,omitempty"`

	// logo image data
	LogoImageData string `json:"logoImageData,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// name identifier format
	// Enum: [urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName urn:oasis:names:tc:SAML:1.1:nameid-format:WindowsDomainQualifiedName urn:oasis:names:tc:SAML:2.0:nameid-format:kerberos urn:oasis:names:tc:SAML:2.0:nameid-format:entity urn:oasis:names:tc:SAML:2.0:nameid-format:persistent urn:oasis:names:tc:SAML:2.0:nameid-format:transient]
	NameIdentifierFormat string `json:"nameIdentifierFormat,omitempty"`

	// relying party identifier
	RelyingPartyIdentifier string `json:"relyingPartyIdentifier,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// sso target URI
	SsoTargetURI string `json:"ssoTargetURI,omitempty"`
}

// Validate validates this generic s a m l
func (m *GenericSAML) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNameIdentifierFormat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var genericSAMLTypeNameIdentifierFormatPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified","urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress","urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName","urn:oasis:names:tc:SAML:1.1:nameid-format:WindowsDomainQualifiedName","urn:oasis:names:tc:SAML:2.0:nameid-format:kerberos","urn:oasis:names:tc:SAML:2.0:nameid-format:entity","urn:oasis:names:tc:SAML:2.0:nameid-format:persistent","urn:oasis:names:tc:SAML:2.0:nameid-format:transient"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		genericSAMLTypeNameIdentifierFormatPropEnum = append(genericSAMLTypeNameIdentifierFormatPropEnum, v)
	}
}

const (

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatUnspecified captures enum value "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatUnspecified string = "urn:oasis:names:tc:SAML:1.1:nameid-format:unspecified"

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatEmailAddress captures enum value "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatEmailAddress string = "urn:oasis:names:tc:SAML:1.1:nameid-format:emailAddress"

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatX509SubjectName captures enum value "urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatX509SubjectName string = "urn:oasis:names:tc:SAML:1.1:nameid-format:X509SubjectName"

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatWindowsDomainQualifiedName captures enum value "urn:oasis:names:tc:SAML:1.1:nameid-format:WindowsDomainQualifiedName"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML11NameidFormatWindowsDomainQualifiedName string = "urn:oasis:names:tc:SAML:1.1:nameid-format:WindowsDomainQualifiedName"

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatKerberos captures enum value "urn:oasis:names:tc:SAML:2.0:nameid-format:kerberos"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatKerberos string = "urn:oasis:names:tc:SAML:2.0:nameid-format:kerberos"

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatEntity captures enum value "urn:oasis:names:tc:SAML:2.0:nameid-format:entity"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatEntity string = "urn:oasis:names:tc:SAML:2.0:nameid-format:entity"

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatPersistent captures enum value "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatPersistent string = "urn:oasis:names:tc:SAML:2.0:nameid-format:persistent"

	// GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatTransient captures enum value "urn:oasis:names:tc:SAML:2.0:nameid-format:transient"
	GenericSAMLNameIdentifierFormatUrnOasisNamesTcSAML20NameidFormatTransient string = "urn:oasis:names:tc:SAML:2.0:nameid-format:transient"
)

// prop value enum
func (m *GenericSAML) validateNameIdentifierFormatEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, genericSAMLTypeNameIdentifierFormatPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GenericSAML) validateNameIdentifierFormat(formats strfmt.Registry) error {

	if swag.IsZero(m.NameIdentifierFormat) { // not required
		return nil
	}

	// value enum
	if err := m.validateNameIdentifierFormatEnum("nameIdentifierFormat", "body", m.NameIdentifierFormat); err != nil {
		return err
	}

	return nil
}

func (m *GenericSAML) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GenericSAML) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericSAML) UnmarshalBinary(b []byte) error {
	var res GenericSAML
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
