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

// PatchCtaButtonStyleProperties patch cta button style properties
//
// swagger:model PatchCtaButtonStyleProperties
type PatchCtaButtonStyleProperties struct {

	// Background color of the CTA button. (eg. #A04033)
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

// Validate validates this patch cta button style properties
func (m *PatchCtaButtonStyleProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTextAlign(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchCtaButtonStylePropertiesTypeTextAlignPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Left","Right","Center"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchCtaButtonStylePropertiesTypeTextAlignPropEnum = append(patchCtaButtonStylePropertiesTypeTextAlignPropEnum, v)
	}
}

const (

	// PatchCtaButtonStylePropertiesTextAlignLeft captures enum value "Left"
	PatchCtaButtonStylePropertiesTextAlignLeft string = "Left"

	// PatchCtaButtonStylePropertiesTextAlignRight captures enum value "Right"
	PatchCtaButtonStylePropertiesTextAlignRight string = "Right"

	// PatchCtaButtonStylePropertiesTextAlignCenter captures enum value "Center"
	PatchCtaButtonStylePropertiesTextAlignCenter string = "Center"
)

// prop value enum
func (m *PatchCtaButtonStyleProperties) validateTextAlignEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchCtaButtonStylePropertiesTypeTextAlignPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchCtaButtonStyleProperties) validateTextAlign(formats strfmt.Registry) error {
	if swag.IsZero(m.TextAlign) { // not required
		return nil
	}

	// value enum
	if err := m.validateTextAlignEnum("textAlign", "body", m.TextAlign); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch cta button style properties based on context it is used
func (m *PatchCtaButtonStyleProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PatchCtaButtonStyleProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchCtaButtonStyleProperties) UnmarshalBinary(b []byte) error {
	var res PatchCtaButtonStyleProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
