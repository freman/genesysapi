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

// Station station
//
// swagger:model Station
type Station struct {

	// description
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// line appearance Id
	LineAppearanceID string `json:"lineAppearanceId,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// primary edge
	PrimaryEdge *DomainEntityRef `json:"primaryEdge,omitempty"`

	// secondary edge
	SecondaryEdge *DomainEntityRef `json:"secondaryEdge,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// status
	// Enum: [AVAILABLE ASSOCIATED]
	Status string `json:"status,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// The Id of the user currently logged in and associated with the station.
	UserID string `json:"userId,omitempty"`

	// The number of call appearances on the station.
	// Read Only: true
	WebRtcCallAppearances int32 `json:"webRtcCallAppearances,omitempty"`

	// Whether the station is configured to require TURN for routing WebRTC calls. Empty if station type is not inin_webrtc_softphone.
	// Read Only: true
	WebRtcForceTurn *bool `json:"webRtcForceTurn"`

	// The default or configured value of media dscp for the station. Empty if station type is not inin_webrtc_softphone.
	// Read Only: true
	WebRtcMediaDscp int32 `json:"webRtcMediaDscp,omitempty"`

	// The default or configured value of persistent connection setting for the station. Empty if station type is not inin_webrtc_softphone.
	// Read Only: true
	WebRtcPersistentEnabled *bool `json:"webRtcPersistentEnabled"`

	// The Id of the user configured for the station if it is of type inin_webrtc_softphone. Empty if station type is not inin_webrtc_softphone.
	WebRtcUserID string `json:"webRtcUserId,omitempty"`
}

// Validate validates this station
func (m *Station) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrimaryEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecondaryEdge(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Station) validatePrimaryEdge(formats strfmt.Registry) error {

	if swag.IsZero(m.PrimaryEdge) { // not required
		return nil
	}

	if m.PrimaryEdge != nil {
		if err := m.PrimaryEdge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("primaryEdge")
			}
			return err
		}
	}

	return nil
}

func (m *Station) validateSecondaryEdge(formats strfmt.Registry) error {

	if swag.IsZero(m.SecondaryEdge) { // not required
		return nil
	}

	if m.SecondaryEdge != nil {
		if err := m.SecondaryEdge.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("secondaryEdge")
			}
			return err
		}
	}

	return nil
}

func (m *Station) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var stationTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AVAILABLE","ASSOCIATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stationTypeStatusPropEnum = append(stationTypeStatusPropEnum, v)
	}
}

const (

	// StationStatusAVAILABLE captures enum value "AVAILABLE"
	StationStatusAVAILABLE string = "AVAILABLE"

	// StationStatusASSOCIATED captures enum value "ASSOCIATED"
	StationStatusASSOCIATED string = "ASSOCIATED"
)

// prop value enum
func (m *Station) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, stationTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Station) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Station) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Station) UnmarshalBinary(b []byte) error {
	var res Station
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
