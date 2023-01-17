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

// DNSRecordEntry Dns record entry
//
// swagger:model DnsRecordEntry
type DNSRecordEntry struct {

	// the hostname of the DNS entry
	Host string `json:"host,omitempty"`

	// the payload of the DNS entry
	RecordContents string `json:"recordContents,omitempty"`

	// the current status of the related verification process
	// Enum: [Pending Verified Failed Unknown]
	VerificationStatus string `json:"verificationStatus,omitempty"`
}

// Validate validates this Dns record entry
func (m *DNSRecordEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateVerificationStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var dnsRecordEntryTypeVerificationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Verified","Failed","Unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		dnsRecordEntryTypeVerificationStatusPropEnum = append(dnsRecordEntryTypeVerificationStatusPropEnum, v)
	}
}

const (

	// DNSRecordEntryVerificationStatusPending captures enum value "Pending"
	DNSRecordEntryVerificationStatusPending string = "Pending"

	// DNSRecordEntryVerificationStatusVerified captures enum value "Verified"
	DNSRecordEntryVerificationStatusVerified string = "Verified"

	// DNSRecordEntryVerificationStatusFailed captures enum value "Failed"
	DNSRecordEntryVerificationStatusFailed string = "Failed"

	// DNSRecordEntryVerificationStatusUnknown captures enum value "Unknown"
	DNSRecordEntryVerificationStatusUnknown string = "Unknown"
)

// prop value enum
func (m *DNSRecordEntry) validateVerificationStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, dnsRecordEntryTypeVerificationStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *DNSRecordEntry) validateVerificationStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.VerificationStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateVerificationStatusEnum("verificationStatus", "body", m.VerificationStatus); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this Dns record entry based on context it is used
func (m *DNSRecordEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DNSRecordEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DNSRecordEntry) UnmarshalBinary(b []byte) error {
	var res DNSRecordEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
