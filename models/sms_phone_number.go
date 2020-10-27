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

// SmsPhoneNumber sms phone number
//
// swagger:model SmsPhoneNumber
type SmsPhoneNumber struct {

	// The id of an address attached to this phone number.
	AddressID *SmsAddress `json:"addressId,omitempty"`

	// Renewal time period of this phone number, if the phoneNumberType is shortcode.
	// Enum: [Quarterly]
	AutoRenewable string `json:"autoRenewable,omitempty"`

	// Contract end date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	CancellationDate strfmt.DateTime `json:"cancellationDate,omitempty"`

	// The capabilities of the phone number available for provisioning.
	// Read Only: true
	Capabilities []string `json:"capabilities"`

	// The ISO 3166-1 alpha-2 country code of the country this phone number is associated with.
	CountryCode string `json:"countryCode,omitempty"`

	// User that provisioned this phone number
	CreatedBy *User `json:"createdBy,omitempty"`

	// Date this phone number was provisioned. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateCreated strfmt.DateTime `json:"dateCreated,omitempty"`

	// Date this phone number was modified. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	DateModified strfmt.DateTime `json:"dateModified,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// User that last modified this phone number
	ModifiedBy *User `json:"modifiedBy,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// A phone number provisioned for SMS communications in E.164 format. E.g. +13175555555 or +34234234234
	// Required: true
	PhoneNumber *string `json:"phoneNumber"`

	// Status of the provisioned phone number.
	// Enum: [INVALID ACTIVE PORTING PENDING PENDING_CANCELLATION]
	PhoneNumberStatus string `json:"phoneNumberStatus,omitempty"`

	// Type of the phone number provisioned.
	// Read Only: true
	// Enum: [local mobile tollfree shortcode]
	PhoneNumberType string `json:"phoneNumberType,omitempty"`

	// Is set to false, if the phone number is provisioned through a SMS provider, outside of PureCloud
	ProvisionedThroughPureCloud bool `json:"provisionedThroughPureCloud"`

	// Date this phone number was purchased, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	PurchaseDate strfmt.DateTime `json:"purchaseDate,omitempty"`

	// Contract renewal date of this phone number, if the phoneNumberType is shortcode. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
	// Format: date-time
	RenewalDate strfmt.DateTime `json:"renewalDate,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// BillingType of this phone number, if the phoneNumberType is shortcode.
	// Enum: [Basic Vanity]
	ShortCodeBillingType string `json:"shortCodeBillingType,omitempty"`

	// Version number required for updates.
	// Required: true
	Version *int32 `json:"version"`
}

// Validate validates this sms phone number
func (m *SmsPhoneNumber) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddressID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoRenewable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCancellationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCapabilities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateModified(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModifiedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumberStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhoneNumberType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePurchaseDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenewalDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShortCodeBillingType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SmsPhoneNumber) validateAddressID(formats strfmt.Registry) error {

	if swag.IsZero(m.AddressID) { // not required
		return nil
	}

	if m.AddressID != nil {
		if err := m.AddressID.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addressId")
			}
			return err
		}
	}

	return nil
}

var smsPhoneNumberTypeAutoRenewablePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Quarterly"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsPhoneNumberTypeAutoRenewablePropEnum = append(smsPhoneNumberTypeAutoRenewablePropEnum, v)
	}
}

const (

	// SmsPhoneNumberAutoRenewableQuarterly captures enum value "Quarterly"
	SmsPhoneNumberAutoRenewableQuarterly string = "Quarterly"
)

// prop value enum
func (m *SmsPhoneNumber) validateAutoRenewableEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smsPhoneNumberTypeAutoRenewablePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmsPhoneNumber) validateAutoRenewable(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoRenewable) { // not required
		return nil
	}

	// value enum
	if err := m.validateAutoRenewableEnum("autoRenewable", "body", m.AutoRenewable); err != nil {
		return err
	}

	return nil
}

func (m *SmsPhoneNumber) validateCancellationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.CancellationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("cancellationDate", "body", "date-time", m.CancellationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var smsPhoneNumberCapabilitiesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["sms","mms","voice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsPhoneNumberCapabilitiesItemsEnum = append(smsPhoneNumberCapabilitiesItemsEnum, v)
	}
}

func (m *SmsPhoneNumber) validateCapabilitiesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smsPhoneNumberCapabilitiesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmsPhoneNumber) validateCapabilities(formats strfmt.Registry) error {

	if swag.IsZero(m.Capabilities) { // not required
		return nil
	}

	for i := 0; i < len(m.Capabilities); i++ {

		// value enum
		if err := m.validateCapabilitiesItemsEnum("capabilities"+"."+strconv.Itoa(i), "body", m.Capabilities[i]); err != nil {
			return err
		}

	}

	return nil
}

func (m *SmsPhoneNumber) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *SmsPhoneNumber) validateDateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.DateCreated) { // not required
		return nil
	}

	if err := validate.FormatOf("dateCreated", "body", "date-time", m.DateCreated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmsPhoneNumber) validateDateModified(formats strfmt.Registry) error {

	if swag.IsZero(m.DateModified) { // not required
		return nil
	}

	if err := validate.FormatOf("dateModified", "body", "date-time", m.DateModified.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmsPhoneNumber) validateModifiedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.ModifiedBy) { // not required
		return nil
	}

	if m.ModifiedBy != nil {
		if err := m.ModifiedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("modifiedBy")
			}
			return err
		}
	}

	return nil
}

func (m *SmsPhoneNumber) validatePhoneNumber(formats strfmt.Registry) error {

	if err := validate.Required("phoneNumber", "body", m.PhoneNumber); err != nil {
		return err
	}

	return nil
}

var smsPhoneNumberTypePhoneNumberStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["INVALID","ACTIVE","PORTING","PENDING","PENDING_CANCELLATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsPhoneNumberTypePhoneNumberStatusPropEnum = append(smsPhoneNumberTypePhoneNumberStatusPropEnum, v)
	}
}

const (

	// SmsPhoneNumberPhoneNumberStatusINVALID captures enum value "INVALID"
	SmsPhoneNumberPhoneNumberStatusINVALID string = "INVALID"

	// SmsPhoneNumberPhoneNumberStatusACTIVE captures enum value "ACTIVE"
	SmsPhoneNumberPhoneNumberStatusACTIVE string = "ACTIVE"

	// SmsPhoneNumberPhoneNumberStatusPORTING captures enum value "PORTING"
	SmsPhoneNumberPhoneNumberStatusPORTING string = "PORTING"

	// SmsPhoneNumberPhoneNumberStatusPENDING captures enum value "PENDING"
	SmsPhoneNumberPhoneNumberStatusPENDING string = "PENDING"

	// SmsPhoneNumberPhoneNumberStatusPENDINGCANCELLATION captures enum value "PENDING_CANCELLATION"
	SmsPhoneNumberPhoneNumberStatusPENDINGCANCELLATION string = "PENDING_CANCELLATION"
)

// prop value enum
func (m *SmsPhoneNumber) validatePhoneNumberStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smsPhoneNumberTypePhoneNumberStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmsPhoneNumber) validatePhoneNumberStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumberStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhoneNumberStatusEnum("phoneNumberStatus", "body", m.PhoneNumberStatus); err != nil {
		return err
	}

	return nil
}

var smsPhoneNumberTypePhoneNumberTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["local","mobile","tollfree","shortcode"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsPhoneNumberTypePhoneNumberTypePropEnum = append(smsPhoneNumberTypePhoneNumberTypePropEnum, v)
	}
}

const (

	// SmsPhoneNumberPhoneNumberTypeLocal captures enum value "local"
	SmsPhoneNumberPhoneNumberTypeLocal string = "local"

	// SmsPhoneNumberPhoneNumberTypeMobile captures enum value "mobile"
	SmsPhoneNumberPhoneNumberTypeMobile string = "mobile"

	// SmsPhoneNumberPhoneNumberTypeTollfree captures enum value "tollfree"
	SmsPhoneNumberPhoneNumberTypeTollfree string = "tollfree"

	// SmsPhoneNumberPhoneNumberTypeShortcode captures enum value "shortcode"
	SmsPhoneNumberPhoneNumberTypeShortcode string = "shortcode"
)

// prop value enum
func (m *SmsPhoneNumber) validatePhoneNumberTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smsPhoneNumberTypePhoneNumberTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmsPhoneNumber) validatePhoneNumberType(formats strfmt.Registry) error {

	if swag.IsZero(m.PhoneNumberType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePhoneNumberTypeEnum("phoneNumberType", "body", m.PhoneNumberType); err != nil {
		return err
	}

	return nil
}

func (m *SmsPhoneNumber) validatePurchaseDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PurchaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("purchaseDate", "body", "date-time", m.PurchaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmsPhoneNumber) validateRenewalDate(formats strfmt.Registry) error {

	if swag.IsZero(m.RenewalDate) { // not required
		return nil
	}

	if err := validate.FormatOf("renewalDate", "body", "date-time", m.RenewalDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SmsPhoneNumber) validateSelfURI(formats strfmt.Registry) error {

	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

var smsPhoneNumberTypeShortCodeBillingTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Basic","Vanity"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		smsPhoneNumberTypeShortCodeBillingTypePropEnum = append(smsPhoneNumberTypeShortCodeBillingTypePropEnum, v)
	}
}

const (

	// SmsPhoneNumberShortCodeBillingTypeBasic captures enum value "Basic"
	SmsPhoneNumberShortCodeBillingTypeBasic string = "Basic"

	// SmsPhoneNumberShortCodeBillingTypeVanity captures enum value "Vanity"
	SmsPhoneNumberShortCodeBillingTypeVanity string = "Vanity"
)

// prop value enum
func (m *SmsPhoneNumber) validateShortCodeBillingTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, smsPhoneNumberTypeShortCodeBillingTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SmsPhoneNumber) validateShortCodeBillingType(formats strfmt.Registry) error {

	if swag.IsZero(m.ShortCodeBillingType) { // not required
		return nil
	}

	// value enum
	if err := m.validateShortCodeBillingTypeEnum("shortCodeBillingType", "body", m.ShortCodeBillingType); err != nil {
		return err
	}

	return nil
}

func (m *SmsPhoneNumber) validateVersion(formats strfmt.Registry) error {

	if err := validate.Required("version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SmsPhoneNumber) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SmsPhoneNumber) UnmarshalBinary(b []byte) error {
	var res SmsPhoneNumber
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
