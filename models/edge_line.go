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

// EdgeLine edge line
//
// swagger:model EdgeLine
type EdgeLine struct {

	// The ID of the user that created the resource.
	CreatedBy string `json:"createdBy,omitempty"`

	// The application that created the resource.
	CreatedByApp string `json:"createdByApp,omitempty"`

	// The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The resource's description.
	Description string `json:"description,omitempty"`

	// The division to which this entity belongs.
	Division *Division `json:"division,omitempty"`

	// edge
	Edge *Edge `json:"edge,omitempty"`

	// edge group
	EdgeGroup *EdgeGroup `json:"edgeGroup,omitempty"`

	// endpoint
	Endpoint *Endpoint `json:"endpoint,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// ip address
	IPAddress string `json:"ipAddress,omitempty"`

	// line type
	// Enum: [TIE NETWORK TRUNK STATION]
	LineType string `json:"lineType,omitempty"`

	// logical interface Id
	LogicalInterfaceID string `json:"logicalInterfaceId,omitempty"`

	// The ID of the user that last modified the resource.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
	ModifiedByApp string `json:"modifiedByApp,omitempty"`

	// The name of the entity.
	// Required: true
	Name *string `json:"name"`

	// properties
	Properties map[string]interface{} `json:"properties,omitempty"`

	// schema
	Schema *DomainEntityRef `json:"schema,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Indicates if the resource is active, inactive, or deleted.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// The current version of the resource.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this edge line
func (m *EdgeLine) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDivision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEdgeGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEndpoint(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSchema(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EdgeLine) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLine) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLine) validateDivision(formats strfmt.Registry) error {

	if swag.IsZero(m.Division) { // not required
		return nil
	}

	if m.Division != nil {
		if err := m.Division.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("division")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeLine) validateEdge(formats strfmt.Registry) error {

	if swag.IsZero(m.Edge) { // not required
		return nil
	}

	if m.Edge != nil {
		if err := m.Edge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edge")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeLine) validateEdgeGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.EdgeGroup) { // not required
		return nil
	}

	if m.EdgeGroup != nil {
		if err := m.EdgeGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("edgeGroup")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeLine) validateEndpoint(formats strfmt.Registry) error {

	if swag.IsZero(m.Endpoint) { // not required
		return nil
	}

	if m.Endpoint != nil {
		if err := m.Endpoint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("endpoint")
			}
			return err
		}
	}

	return nil
}

var edgeLineTypeLineTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TIE","NETWORK","TRUNK","STATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgeLineTypeLineTypePropEnum = append(edgeLineTypeLineTypePropEnum, v)
	}
}

const (

	// EdgeLineLineTypeTIE captures enum value "TIE"
	EdgeLineLineTypeTIE string = "TIE"

	// EdgeLineLineTypeNETWORK captures enum value "NETWORK"
	EdgeLineLineTypeNETWORK string = "NETWORK"

	// EdgeLineLineTypeTRUNK captures enum value "TRUNK"
	EdgeLineLineTypeTRUNK string = "TRUNK"

	// EdgeLineLineTypeSTATION captures enum value "STATION"
	EdgeLineLineTypeSTATION string = "STATION"
)

// prop value enum
func (m *EdgeLine) validateLineTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgeLineTypeLineTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgeLine) validateLineType(formats strfmt.Registry) error {

	if swag.IsZero(m.LineType) { // not required
		return nil
	}

	// value enum
	if err := m.validateLineTypeEnum("lineType", "body", m.LineType); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLine) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EdgeLine) validateSchema(formats strfmt.Registry) error {

	if swag.IsZero(m.Schema) { // not required
		return nil
	}

	if m.Schema != nil {
		if err := m.Schema.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("schema")
			}
			return err
		}
	}

	return nil
}

func (m *EdgeLine) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var edgeLineTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		edgeLineTypeStatePropEnum = append(edgeLineTypeStatePropEnum, v)
	}
}

const (

	// EdgeLineStateActive captures enum value "active"
	EdgeLineStateActive string = "active"

	// EdgeLineStateInactive captures enum value "inactive"
	EdgeLineStateInactive string = "inactive"

	// EdgeLineStateDeleted captures enum value "deleted"
	EdgeLineStateDeleted string = "deleted"
)

// prop value enum
func (m *EdgeLine) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, edgeLineTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *EdgeLine) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EdgeLine) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EdgeLine) UnmarshalBinary(b []byte) error {
	var res EdgeLine
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
