// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PhoneCapabilities phone capabilities
//
// swagger:model PhoneCapabilities
type PhoneCapabilities struct {

	// allow reboot
	AllowReboot bool `json:"allowReboot"`

	// cdm
	Cdm bool `json:"cdm"`

	// dual registers
	DualRegisters bool `json:"dualRegisters"`

	// hardware Id type
	HardwareIDType string `json:"hardwareIdType,omitempty"`

	// media codecs
	MediaCodecs []string `json:"mediaCodecs"`

	// no cloud provisioning
	NoCloudProvisioning bool `json:"noCloudProvisioning"`

	// no rebalance
	NoRebalance bool `json:"noRebalance"`

	// provisions
	Provisions bool `json:"provisions"`

	// registers
	Registers bool `json:"registers"`
}

// Validate validates this phone capabilities
func (m *PhoneCapabilities) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMediaCodecs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var phoneCapabilitiesMediaCodecsItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["audio/opus","audio/pcmu","audio/pcma","audio/g729","audio/g722"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		phoneCapabilitiesMediaCodecsItemsEnum = append(phoneCapabilitiesMediaCodecsItemsEnum, v)
	}
}

func (m *PhoneCapabilities) validateMediaCodecsItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, phoneCapabilitiesMediaCodecsItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PhoneCapabilities) validateMediaCodecs(formats strfmt.Registry) error {
	if swag.IsZero(m.MediaCodecs) { // not required
		return nil
	}

	for i := 0; i < len(m.MediaCodecs); i++ {

		// value enum
		if err := m.validateMediaCodecsItemsEnum("mediaCodecs"+"."+strconv.Itoa(i), "body", m.MediaCodecs[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this phone capabilities based on context it is used
func (m *PhoneCapabilities) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PhoneCapabilities) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PhoneCapabilities) UnmarshalBinary(b []byte) error {
	var res PhoneCapabilities
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
