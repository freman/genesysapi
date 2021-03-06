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

// JSONNode Json node
//
// swagger:model JsonNode
type JSONNode struct {

	// array
	Array bool `json:"array"`

	// big decimal
	BigDecimal bool `json:"bigDecimal"`

	// big integer
	BigInteger bool `json:"bigInteger"`

	// binary
	Binary bool `json:"binary"`

	// boolean
	Boolean bool `json:"boolean"`

	// container node
	ContainerNode bool `json:"containerNode"`

	// double
	Double bool `json:"double"`

	// float
	Float bool `json:"float"`

	// floating point number
	FloatingPointNumber bool `json:"floatingPointNumber"`

	// int
	Int bool `json:"int"`

	// integral number
	IntegralNumber bool `json:"integralNumber"`

	// long
	Long bool `json:"long"`

	// missing node
	MissingNode bool `json:"missingNode"`

	// node type
	// Enum: [ARRAY BINARY BOOLEAN MISSING NULL NUMBER OBJECT POJO STRING]
	NodeType string `json:"nodeType,omitempty"`

	// null
	Null bool `json:"null"`

	// number
	Number bool `json:"number"`

	// object
	Object bool `json:"object"`

	// pojo
	Pojo bool `json:"pojo"`

	// short
	Short bool `json:"short"`

	// textual
	Textual bool `json:"textual"`

	// value node
	ValueNode bool `json:"valueNode"`
}

// Validate validates this Json node
func (m *JSONNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var jsonNodeTypeNodeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ARRAY","BINARY","BOOLEAN","MISSING","NULL","NUMBER","OBJECT","POJO","STRING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jsonNodeTypeNodeTypePropEnum = append(jsonNodeTypeNodeTypePropEnum, v)
	}
}

const (

	// JSONNodeNodeTypeARRAY captures enum value "ARRAY"
	JSONNodeNodeTypeARRAY string = "ARRAY"

	// JSONNodeNodeTypeBINARY captures enum value "BINARY"
	JSONNodeNodeTypeBINARY string = "BINARY"

	// JSONNodeNodeTypeBOOLEAN captures enum value "BOOLEAN"
	JSONNodeNodeTypeBOOLEAN string = "BOOLEAN"

	// JSONNodeNodeTypeMISSING captures enum value "MISSING"
	JSONNodeNodeTypeMISSING string = "MISSING"

	// JSONNodeNodeTypeNULL captures enum value "NULL"
	JSONNodeNodeTypeNULL string = "NULL"

	// JSONNodeNodeTypeNUMBER captures enum value "NUMBER"
	JSONNodeNodeTypeNUMBER string = "NUMBER"

	// JSONNodeNodeTypeOBJECT captures enum value "OBJECT"
	JSONNodeNodeTypeOBJECT string = "OBJECT"

	// JSONNodeNodeTypePOJO captures enum value "POJO"
	JSONNodeNodeTypePOJO string = "POJO"

	// JSONNodeNodeTypeSTRING captures enum value "STRING"
	JSONNodeNodeTypeSTRING string = "STRING"
)

// prop value enum
func (m *JSONNode) validateNodeTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jsonNodeTypeNodeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JSONNode) validateNodeType(formats strfmt.Registry) error {

	if swag.IsZero(m.NodeType) { // not required
		return nil
	}

	// value enum
	if err := m.validateNodeTypeEnum("nodeType", "body", m.NodeType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONNode) UnmarshalBinary(b []byte) error {
	var res JSONNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
