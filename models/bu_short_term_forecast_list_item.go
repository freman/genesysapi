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

// BuShortTermForecastListItem bu short term forecast list item
//
// swagger:model BuShortTermForecastListItem
type BuShortTermForecastListItem struct {

	// Whether this forecast can be used for scheduling
	CanUseForScheduling bool `json:"canUseForScheduling"`

	// The method by which this forecast was created
	// Enum: [Import ImportedHistoricalWeightedAverage HistoricalWeightedAverage Advanced]
	CreationMethod string `json:"creationMethod,omitempty"`

	// The description of this forecast
	Description string `json:"description,omitempty"`

	// The globally unique identifier for the object.
	// Read Only: true
	ID string `json:"id,omitempty"`

	// Whether this forecast contains modifications on legacy metrics
	// Read Only: true
	Legacy *bool `json:"legacy"`

	// Metadata for this forecast
	Metadata *WfmVersionedEntityMetadata `json:"metadata,omitempty"`

	// The URI for this object
	// Read Only: true
	// Format: uri
	SelfURI strfmt.URI `json:"selfUri,omitempty"`

	// The number of weeks this forecast covers
	WeekCount int32 `json:"weekCount,omitempty"`

	// The start week date of this forecast in yyyy-MM-dd.  Must fall on the start day of week for the associated business unit. Dates are represented as an ISO-8601 string. For example: yyyy-MM-dd
	// Format: date
	WeekDate strfmt.Date `json:"weekDate,omitempty"`
}

// Validate validates this bu short term forecast list item
func (m *BuShortTermForecastListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelfURI(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWeekDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var buShortTermForecastListItemTypeCreationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Import","ImportedHistoricalWeightedAverage","HistoricalWeightedAverage","Advanced"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		buShortTermForecastListItemTypeCreationMethodPropEnum = append(buShortTermForecastListItemTypeCreationMethodPropEnum, v)
	}
}

const (

	// BuShortTermForecastListItemCreationMethodImport captures enum value "Import"
	BuShortTermForecastListItemCreationMethodImport string = "Import"

	// BuShortTermForecastListItemCreationMethodImportedHistoricalWeightedAverage captures enum value "ImportedHistoricalWeightedAverage"
	BuShortTermForecastListItemCreationMethodImportedHistoricalWeightedAverage string = "ImportedHistoricalWeightedAverage"

	// BuShortTermForecastListItemCreationMethodHistoricalWeightedAverage captures enum value "HistoricalWeightedAverage"
	BuShortTermForecastListItemCreationMethodHistoricalWeightedAverage string = "HistoricalWeightedAverage"

	// BuShortTermForecastListItemCreationMethodAdvanced captures enum value "Advanced"
	BuShortTermForecastListItemCreationMethodAdvanced string = "Advanced"
)

// prop value enum
func (m *BuShortTermForecastListItem) validateCreationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, buShortTermForecastListItemTypeCreationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BuShortTermForecastListItem) validateCreationMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.CreationMethod) { // not required
		return nil
	}

	// value enum
	if err := m.validateCreationMethodEnum("creationMethod", "body", m.CreationMethod); err != nil {
		return err
	}

	return nil
}

func (m *BuShortTermForecastListItem) validateMetadata(formats strfmt.Registry) error {
	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if m.Metadata != nil {
		if err := m.Metadata.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *BuShortTermForecastListItem) validateSelfURI(formats strfmt.Registry) error {
	if swag.IsZero(m.SelfURI) { // not required
		return nil
	}

	if err := validate.FormatOf("selfUri", "body", "uri", m.SelfURI.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BuShortTermForecastListItem) validateWeekDate(formats strfmt.Registry) error {
	if swag.IsZero(m.WeekDate) { // not required
		return nil
	}

	if err := validate.FormatOf("weekDate", "body", "date", m.WeekDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this bu short term forecast list item based on the context it is used
func (m *BuShortTermForecastListItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLegacy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetadata(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelfURI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BuShortTermForecastListItem) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *BuShortTermForecastListItem) contextValidateLegacy(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "legacy", "body", m.Legacy); err != nil {
		return err
	}

	return nil
}

func (m *BuShortTermForecastListItem) contextValidateMetadata(ctx context.Context, formats strfmt.Registry) error {

	if m.Metadata != nil {
		if err := m.Metadata.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("metadata")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("metadata")
			}
			return err
		}
	}

	return nil
}

func (m *BuShortTermForecastListItem) contextValidateSelfURI(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "selfUri", "body", strfmt.URI(m.SelfURI)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BuShortTermForecastListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BuShortTermForecastListItem) UnmarshalBinary(b []byte) error {
	var res BuShortTermForecastListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
