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

// ScimV2SchemaAttribute A complex type that defines service provider attributes or subattributes and their qualities.
//
// swagger:model ScimV2SchemaAttribute
type ScimV2SchemaAttribute struct {

	// The list of standard values that service providers may use. Service providers may ignore unsupported values.
	// Read Only: true
	CanonicalValues []string `json:"canonicalValues"`

	// Indicates whether a string attribute is case-sensitive. If set to "true", the server preserves case sensitivity. If set to "false", the server may change the case. The server also uses case sensitivity when evaluating filters. See section 3.4.2.2 "Filtering" in RFC 7644 for details.
	// Read Only: true
	CaseExact *bool `json:"caseExact"`

	// The description of the attribute.
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Indicates whether an attribute contains multiple values.
	// Read Only: true
	MultiValued *bool `json:"multiValued"`

	// The circumstances under which an attribute can be defined or redefined. The default is "readWrite".
	// Read Only: true
	// Enum: [readWrite readOnly immutable writeOnly]
	Mutability string `json:"mutability,omitempty"`

	// The name of the attribute.
	// Read Only: true
	Name string `json:"name,omitempty"`

	// The list of SCIM resource types that may be referenced. Only applies when "type" is set to "reference".
	// Read Only: true
	ReferenceTypes []string `json:"referenceTypes"`

	// Indicates whether an attribute is required.
	// Read Only: true
	Required *bool `json:"required"`

	// The circumstances under which an attribute and its values are returned in response to a GET, PUT, POST, or PATCH request.
	// Read Only: true
	// Enum: [always never default request]
	Returned string `json:"returned,omitempty"`

	// The list of subattributes for an attribute of the type "complex". Uses the same schema as "attributes".
	// Read Only: true
	SubAttributes []*ScimV2SchemaAttribute `json:"subAttributes"`

	// The data type of the attribute.
	// Read Only: true
	// Enum: [string boolean decimal integer dateTime reference complex]
	Type string `json:"type,omitempty"`

	// The method by which the service provider enforces the uniqueness of an attribute value. A server can reject a value by returning the HTTP response code 400 (Bad Request). A client can enforce uniqueness to a greater degree than the server provider enforces. For example, a client could make a value unique even though the server has "uniqueness" set to "none".
	// Read Only: true
	// Enum: [none server global]
	Uniqueness string `json:"uniqueness,omitempty"`
}

// Validate validates this scim v2 schema attribute
func (m *ScimV2SchemaAttribute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMutability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferenceTypes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReturned(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubAttributes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUniqueness(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var scimV2SchemaAttributeTypeMutabilityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["readWrite","readOnly","immutable","writeOnly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimV2SchemaAttributeTypeMutabilityPropEnum = append(scimV2SchemaAttributeTypeMutabilityPropEnum, v)
	}
}

const (

	// ScimV2SchemaAttributeMutabilityReadWrite captures enum value "readWrite"
	ScimV2SchemaAttributeMutabilityReadWrite string = "readWrite"

	// ScimV2SchemaAttributeMutabilityReadOnly captures enum value "readOnly"
	ScimV2SchemaAttributeMutabilityReadOnly string = "readOnly"

	// ScimV2SchemaAttributeMutabilityImmutable captures enum value "immutable"
	ScimV2SchemaAttributeMutabilityImmutable string = "immutable"

	// ScimV2SchemaAttributeMutabilityWriteOnly captures enum value "writeOnly"
	ScimV2SchemaAttributeMutabilityWriteOnly string = "writeOnly"
)

// prop value enum
func (m *ScimV2SchemaAttribute) validateMutabilityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimV2SchemaAttributeTypeMutabilityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimV2SchemaAttribute) validateMutability(formats strfmt.Registry) error {

	if swag.IsZero(m.Mutability) { // not required
		return nil
	}

	// value enum
	if err := m.validateMutabilityEnum("mutability", "body", m.Mutability); err != nil {
		return err
	}

	return nil
}

var scimV2SchemaAttributeReferenceTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["User","Group","external","uri"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimV2SchemaAttributeReferenceTypesItemsEnum = append(scimV2SchemaAttributeReferenceTypesItemsEnum, v)
	}
}

func (m *ScimV2SchemaAttribute) validateReferenceTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimV2SchemaAttributeReferenceTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimV2SchemaAttribute) validateReferenceTypes(formats strfmt.Registry) error {

	if swag.IsZero(m.ReferenceTypes) { // not required
		return nil
	}

	for i := 0; i < len(m.ReferenceTypes); i++ {

		// value enum
		if err := m.validateReferenceTypesItemsEnum("referenceTypes"+"."+strconv.Itoa(i), "body", m.ReferenceTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

var scimV2SchemaAttributeTypeReturnedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["always","never","default","request"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimV2SchemaAttributeTypeReturnedPropEnum = append(scimV2SchemaAttributeTypeReturnedPropEnum, v)
	}
}

const (

	// ScimV2SchemaAttributeReturnedAlways captures enum value "always"
	ScimV2SchemaAttributeReturnedAlways string = "always"

	// ScimV2SchemaAttributeReturnedNever captures enum value "never"
	ScimV2SchemaAttributeReturnedNever string = "never"

	// ScimV2SchemaAttributeReturnedDefault captures enum value "default"
	ScimV2SchemaAttributeReturnedDefault string = "default"

	// ScimV2SchemaAttributeReturnedRequest captures enum value "request"
	ScimV2SchemaAttributeReturnedRequest string = "request"
)

// prop value enum
func (m *ScimV2SchemaAttribute) validateReturnedEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimV2SchemaAttributeTypeReturnedPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimV2SchemaAttribute) validateReturned(formats strfmt.Registry) error {

	if swag.IsZero(m.Returned) { // not required
		return nil
	}

	// value enum
	if err := m.validateReturnedEnum("returned", "body", m.Returned); err != nil {
		return err
	}

	return nil
}

func (m *ScimV2SchemaAttribute) validateSubAttributes(formats strfmt.Registry) error {

	if swag.IsZero(m.SubAttributes) { // not required
		return nil
	}

	for i := 0; i < len(m.SubAttributes); i++ {
		if swag.IsZero(m.SubAttributes[i]) { // not required
			continue
		}

		if m.SubAttributes[i] != nil {
			if err := m.SubAttributes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subAttributes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var scimV2SchemaAttributeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["string","boolean","decimal","integer","dateTime","reference","complex"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimV2SchemaAttributeTypeTypePropEnum = append(scimV2SchemaAttributeTypeTypePropEnum, v)
	}
}

const (

	// ScimV2SchemaAttributeTypeString captures enum value "string"
	ScimV2SchemaAttributeTypeString string = "string"

	// ScimV2SchemaAttributeTypeBoolean captures enum value "boolean"
	ScimV2SchemaAttributeTypeBoolean string = "boolean"

	// ScimV2SchemaAttributeTypeDecimal captures enum value "decimal"
	ScimV2SchemaAttributeTypeDecimal string = "decimal"

	// ScimV2SchemaAttributeTypeInteger captures enum value "integer"
	ScimV2SchemaAttributeTypeInteger string = "integer"

	// ScimV2SchemaAttributeTypeDateTime captures enum value "dateTime"
	ScimV2SchemaAttributeTypeDateTime string = "dateTime"

	// ScimV2SchemaAttributeTypeReference captures enum value "reference"
	ScimV2SchemaAttributeTypeReference string = "reference"

	// ScimV2SchemaAttributeTypeComplex captures enum value "complex"
	ScimV2SchemaAttributeTypeComplex string = "complex"
)

// prop value enum
func (m *ScimV2SchemaAttribute) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimV2SchemaAttributeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimV2SchemaAttribute) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

var scimV2SchemaAttributeTypeUniquenessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["none","server","global"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		scimV2SchemaAttributeTypeUniquenessPropEnum = append(scimV2SchemaAttributeTypeUniquenessPropEnum, v)
	}
}

const (

	// ScimV2SchemaAttributeUniquenessNone captures enum value "none"
	ScimV2SchemaAttributeUniquenessNone string = "none"

	// ScimV2SchemaAttributeUniquenessServer captures enum value "server"
	ScimV2SchemaAttributeUniquenessServer string = "server"

	// ScimV2SchemaAttributeUniquenessGlobal captures enum value "global"
	ScimV2SchemaAttributeUniquenessGlobal string = "global"
)

// prop value enum
func (m *ScimV2SchemaAttribute) validateUniquenessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, scimV2SchemaAttributeTypeUniquenessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ScimV2SchemaAttribute) validateUniqueness(formats strfmt.Registry) error {

	if swag.IsZero(m.Uniqueness) { // not required
		return nil
	}

	// value enum
	if err := m.validateUniquenessEnum("uniqueness", "body", m.Uniqueness); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ScimV2SchemaAttribute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScimV2SchemaAttribute) UnmarshalBinary(b []byte) error {
	var res ScimV2SchemaAttribute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
