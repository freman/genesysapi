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

// Metabase metabase
//
// swagger:model Metabase
type Metabase struct {

	// The ID of the user that created the resource.
	CreatedBy string `json:"createdBy,omitempty"`

	// The application that created the resource.
	CreatedByApp string `json:"createdByApp,omitempty"`

	// The date the resource was created. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// The date of the last modification to the resource. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss.SSSZ
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// The ID of the user that last modified the resource.
	ModifiedBy string `json:"modifiedBy,omitempty"`

	// The application that last modified the resource.
	ModifiedByApp string `json:"modifiedByApp,omitempty"`

	// The name of the entity.
	// Required: true
	Name *string `json:"name"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// Indicates if the resource is active, inactive, or deleted.
	// Read Only: true
	// Enum: [active inactive deleted]
	State string `json:"state,omitempty"`

	// type
	// Enum: [EXTERNAL EXTERNAL_PCV EXTERNAL_PCV_AWS EXTERNAL_BYOC_CARRIER EXTERNAL_BYOC_PBX STATION_CDM STATION_CDM_WEBRTC STATION TIE TIE_DIRECT TIE_INDIRECT TIE_CLOUD_PROXY]
	Type string `json:"type,omitempty"`

	// The current version of the resource.
	Version int32 `json:"version,omitempty"`
}

// Validate validates this metabase
func (m *Metabase) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Metabase) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Metabase) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Metabase) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Metabase) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var metabaseTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","inactive","deleted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metabaseTypeStatePropEnum = append(metabaseTypeStatePropEnum, v)
	}
}

const (

	// MetabaseStateActive captures enum value "active"
	MetabaseStateActive string = "active"

	// MetabaseStateInactive captures enum value "inactive"
	MetabaseStateInactive string = "inactive"

	// MetabaseStateDeleted captures enum value "deleted"
	MetabaseStateDeleted string = "deleted"
)

// prop value enum
func (m *Metabase) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metabaseTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Metabase) validateState(formats strfmt.Registry) error {

	if swag.IsZero(m.State) { // not required
		return nil
	}

	// value enum
	if err := m.validateStateEnum("state", "body", m.State); err != nil {
		return err
	}

	return nil
}

var metabaseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["EXTERNAL","EXTERNAL_PCV","EXTERNAL_PCV_AWS","EXTERNAL_BYOC_CARRIER","EXTERNAL_BYOC_PBX","STATION_CDM","STATION_CDM_WEBRTC","STATION","TIE","TIE_DIRECT","TIE_INDIRECT","TIE_CLOUD_PROXY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		metabaseTypeTypePropEnum = append(metabaseTypeTypePropEnum, v)
	}
}

const (

	// MetabaseTypeEXTERNAL captures enum value "EXTERNAL"
	MetabaseTypeEXTERNAL string = "EXTERNAL"

	// MetabaseTypeEXTERNALPCV captures enum value "EXTERNAL_PCV"
	MetabaseTypeEXTERNALPCV string = "EXTERNAL_PCV"

	// MetabaseTypeEXTERNALPCVAWS captures enum value "EXTERNAL_PCV_AWS"
	MetabaseTypeEXTERNALPCVAWS string = "EXTERNAL_PCV_AWS"

	// MetabaseTypeEXTERNALBYOCCARRIER captures enum value "EXTERNAL_BYOC_CARRIER"
	MetabaseTypeEXTERNALBYOCCARRIER string = "EXTERNAL_BYOC_CARRIER"

	// MetabaseTypeEXTERNALBYOCPBX captures enum value "EXTERNAL_BYOC_PBX"
	MetabaseTypeEXTERNALBYOCPBX string = "EXTERNAL_BYOC_PBX"

	// MetabaseTypeSTATIONCDM captures enum value "STATION_CDM"
	MetabaseTypeSTATIONCDM string = "STATION_CDM"

	// MetabaseTypeSTATIONCDMWEBRTC captures enum value "STATION_CDM_WEBRTC"
	MetabaseTypeSTATIONCDMWEBRTC string = "STATION_CDM_WEBRTC"

	// MetabaseTypeSTATION captures enum value "STATION"
	MetabaseTypeSTATION string = "STATION"

	// MetabaseTypeTIE captures enum value "TIE"
	MetabaseTypeTIE string = "TIE"

	// MetabaseTypeTIEDIRECT captures enum value "TIE_DIRECT"
	MetabaseTypeTIEDIRECT string = "TIE_DIRECT"

	// MetabaseTypeTIEINDIRECT captures enum value "TIE_INDIRECT"
	MetabaseTypeTIEINDIRECT string = "TIE_INDIRECT"

	// MetabaseTypeTIECLOUDPROXY captures enum value "TIE_CLOUD_PROXY"
	MetabaseTypeTIECLOUDPROXY string = "TIE_CLOUD_PROXY"
)

// prop value enum
func (m *Metabase) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, metabaseTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Metabase) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Metabase) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metabase) UnmarshalBinary(b []byte) error {
	var res Metabase
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}