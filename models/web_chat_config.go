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

// WebChatConfig web chat config
//
// swagger:model WebChatConfig
type WebChatConfig struct {

	// css class to be applied to the web chat widget.
	// Enum: [basic modern-caret-skin]
	WebChatSkin string `json:"webChatSkin,omitempty"`
}

// Validate validates this web chat config
func (m *WebChatConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebChatSkin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var webChatConfigTypeWebChatSkinPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["basic","modern-caret-skin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		webChatConfigTypeWebChatSkinPropEnum = append(webChatConfigTypeWebChatSkinPropEnum, v)
	}
}

const (

	// WebChatConfigWebChatSkinBasic captures enum value "basic"
	WebChatConfigWebChatSkinBasic string = "basic"

	// WebChatConfigWebChatSkinModernDashCaretDashSkin captures enum value "modern-caret-skin"
	WebChatConfigWebChatSkinModernDashCaretDashSkin string = "modern-caret-skin"
)

// prop value enum
func (m *WebChatConfig) validateWebChatSkinEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, webChatConfigTypeWebChatSkinPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WebChatConfig) validateWebChatSkin(formats strfmt.Registry) error {
	if swag.IsZero(m.WebChatSkin) { // not required
		return nil
	}

	// value enum
	if err := m.validateWebChatSkinEnum("webChatSkin", "body", m.WebChatSkin); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this web chat config based on context it is used
func (m *WebChatConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WebChatConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebChatConfig) UnmarshalBinary(b []byte) error {
	var res WebChatConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
