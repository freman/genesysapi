// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SubscriptionOverviewUsage subscription overview usage
//
// swagger:model SubscriptionOverviewUsage
type SubscriptionOverviewUsage struct {

	// Quantity multiplier for this charge
	BundleQuantity string `json:"bundleQuantity,omitempty"`

	// UI grouping key
	// Required: true
	Grouping *string `json:"grouping"`

	// Indicates whether the item is cancellable
	IsCancellable bool `json:"isCancellable"`

	// A charge from a third party entity
	IsThirdParty bool `json:"isThirdParty"`

	// Product charge name
	// Required: true
	Name *string `json:"name"`

	// Price for usage / overage charge
	OveragePrice string `json:"overagePrice,omitempty"`

	// Product part number
	// Required: true
	PartNumber *string `json:"partNumber"`

	// Price for prepay charge
	PrepayPrice string `json:"prepayPrice,omitempty"`

	// Items prepaid for specified period
	// Required: true
	PrepayQuantity *string `json:"prepayQuantity"`

	// UI unit of measure
	// Required: true
	UnitOfMeasureType *string `json:"unitOfMeasureType"`

	// Notes about the usage/charge item
	UsageNotes string `json:"usageNotes,omitempty"`

	// Usage count for specified period
	// Required: true
	UsageQuantity *string `json:"usageQuantity"`
}

// Validate validates this subscription overview usage
func (m *SubscriptionOverviewUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGrouping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePartNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrepayQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitOfMeasureType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsageQuantity(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SubscriptionOverviewUsage) validateGrouping(formats strfmt.Registry) error {

	if err := validate.Required("grouping", "body", m.Grouping); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionOverviewUsage) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionOverviewUsage) validatePartNumber(formats strfmt.Registry) error {

	if err := validate.Required("partNumber", "body", m.PartNumber); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionOverviewUsage) validatePrepayQuantity(formats strfmt.Registry) error {

	if err := validate.Required("prepayQuantity", "body", m.PrepayQuantity); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionOverviewUsage) validateUnitOfMeasureType(formats strfmt.Registry) error {

	if err := validate.Required("unitOfMeasureType", "body", m.UnitOfMeasureType); err != nil {
		return err
	}

	return nil
}

func (m *SubscriptionOverviewUsage) validateUsageQuantity(formats strfmt.Registry) error {

	if err := validate.Required("usageQuantity", "body", m.UsageQuantity); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SubscriptionOverviewUsage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubscriptionOverviewUsage) UnmarshalBinary(b []byte) error {
	var res SubscriptionOverviewUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
