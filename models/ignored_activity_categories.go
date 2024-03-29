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

// IgnoredActivityCategories ignored activity categories
//
// swagger:model IgnoredActivityCategories
type IgnoredActivityCategories struct {

	// Activity categories list
	Values []string `json:"values"`
}

// Validate validates this ignored activity categories
func (m *IgnoredActivityCategories) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValues(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ignoredActivityCategoriesValuesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["OnQueueWork","Break","Meal","Meeting","OffQueueWork","TimeOff","Training","Unavailable","Unscheduled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ignoredActivityCategoriesValuesItemsEnum = append(ignoredActivityCategoriesValuesItemsEnum, v)
	}
}

func (m *IgnoredActivityCategories) validateValuesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ignoredActivityCategoriesValuesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IgnoredActivityCategories) validateValues(formats strfmt.Registry) error {
	if swag.IsZero(m.Values) { // not required
		return nil
	}

	for i := 0; i < len(m.Values); i++ {

		// value enum
		if err := m.validateValuesItemsEnum("values"+"."+strconv.Itoa(i), "body", m.Values[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this ignored activity categories based on context it is used
func (m *IgnoredActivityCategories) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IgnoredActivityCategories) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgnoredActivityCategories) UnmarshalBinary(b []byte) error {
	var res IgnoredActivityCategories
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
