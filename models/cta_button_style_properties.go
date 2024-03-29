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

// CtaButtonStyleProperties cta button style properties
//
// swagger:model CtaButtonStyleProperties
type CtaButtonStyleProperties struct {

	// Background color of the CTA button. (eg. #FF0000)
	BackgroundColor string `json:"backgroundColor,omitempty"`

	// Color of the text. (eg. #FFFFFF)
	Color string `json:"color,omitempty"`

	// Font of the text. (eg. Helvetica)
	Font string `json:"font,omitempty"`

	// Font size of the text. (eg. '12')
	FontSize string `json:"fontSize,omitempty"`

	// Text alignment.
	// Enum: [Left Right Center]
	TextAlign string `json:"textAlign,omitempty"`
}

// Validate validates this cta button style properties
func (m *CtaButtonStyleProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTextAlign(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var ctaButtonStylePropertiesTypeTextAlignPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Left","Right","Center"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ctaButtonStylePropertiesTypeTextAlignPropEnum = append(ctaButtonStylePropertiesTypeTextAlignPropEnum, v)
	}
}

const (

	// CtaButtonStylePropertiesTextAlignLeft captures enum value "Left"
	CtaButtonStylePropertiesTextAlignLeft string = "Left"

	// CtaButtonStylePropertiesTextAlignRight captures enum value "Right"
	CtaButtonStylePropertiesTextAlignRight string = "Right"

	// CtaButtonStylePropertiesTextAlignCenter captures enum value "Center"
	CtaButtonStylePropertiesTextAlignCenter string = "Center"
)

// prop value enum
func (m *CtaButtonStyleProperties) validateTextAlignEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, ctaButtonStylePropertiesTypeTextAlignPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *CtaButtonStyleProperties) validateTextAlign(formats strfmt.Registry) error {
	if swag.IsZero(m.TextAlign) { // not required
		return nil
	}

	// value enum
	if err := m.validateTextAlignEnum("textAlign", "body", m.TextAlign); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this cta button style properties based on context it is used
func (m *CtaButtonStyleProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CtaButtonStyleProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CtaButtonStyleProperties) UnmarshalBinary(b []byte) error {
	var res CtaButtonStyleProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
