// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainNetworkRoute domain network route
//
// swagger:model DomainNetworkRoute
type DomainNetworkRoute struct {

	// The address family for this route.
	// Enum: [2 23]
	Family int32 `json:"family,omitempty"`

	// The metric being used for route. Lower values will have a higher priority.
	Metric int32 `json:"metric,omitempty"`

	// The IPv4 or IPv6 nexthop IP address.
	Nexthop string `json:"nexthop,omitempty"`

	// True if this route will persist on Edge restart.  Routes assigned by DHCP will be returned as false.
	Persistent bool `json:"persistent"`

	// The IPv4 or IPv6 route prefix in CIDR notation.
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this domain network route
func (m *DomainNetworkRoute) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFamily(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var domainNetworkRouteTypeFamilyPropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`[2,23]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainNetworkRouteTypeFamilyPropEnum = append(domainNetworkRouteTypeFamilyPropEnum, v)
	}
}

// prop value enum
func (m *DomainNetworkRoute) validateFamilyEnum(path, location string, value int32) error {
	if err := validate.EnumCase(path, location, value, domainNetworkRouteTypeFamilyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DomainNetworkRoute) validateFamily(formats strfmt.Registry) error {
	if swag.IsZero(m.Family) { // not required
		return nil
	}

	// value enum
	if err := m.validateFamilyEnum("family", "body", m.Family); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this domain network route based on context it is used
func (m *DomainNetworkRoute) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainNetworkRoute) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainNetworkRoute) UnmarshalBinary(b []byte) error {
	var res DomainNetworkRoute
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
