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

// ContentOffer content offer
//
// swagger:model ContentOffer
type ContentOffer struct {

	// Body text of the content offer.
	Body string `json:"body,omitempty"`

	// Properties customizing the call to action button on the content offer.
	CallToAction *CallToAction `json:"callToAction,omitempty"`

	// The display mode of Genesys Widgets when displaying content offer.
	// Required: true
	// Enum: [Modal Overlay Toast]
	DisplayMode *string `json:"displayMode"`

	// Headline displayed above the body text of the content offer.
	Headline string `json:"headline,omitempty"`

	// URL for image displayed to the customer when displaying content offer.
	ImageURL string `json:"imageUrl,omitempty"`

	// The layout mode of the text shown to the user when displaying content offer.
	// Required: true
	// Enum: [TextOnly ImageOnly LeftText RightText TopText BottomText]
	LayoutMode *string `json:"layoutMode"`

	// Properties customizing the styling of the content offer.
	Style *ContentOfferStylingConfiguration `json:"style,omitempty"`

	// Title used in the header of the content offer.
	Title string `json:"title,omitempty"`
}

// Validate validates this content offer
func (m *ContentOffer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCallToAction(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLayoutMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStyle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentOffer) validateCallToAction(formats strfmt.Registry) error {
	if swag.IsZero(m.CallToAction) { // not required
		return nil
	}

	if m.CallToAction != nil {
		if err := m.CallToAction.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callToAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callToAction")
			}
			return err
		}
	}

	return nil
}

var contentOfferTypeDisplayModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Modal","Overlay","Toast"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentOfferTypeDisplayModePropEnum = append(contentOfferTypeDisplayModePropEnum, v)
	}
}

const (

	// ContentOfferDisplayModeModal captures enum value "Modal"
	ContentOfferDisplayModeModal string = "Modal"

	// ContentOfferDisplayModeOverlay captures enum value "Overlay"
	ContentOfferDisplayModeOverlay string = "Overlay"

	// ContentOfferDisplayModeToast captures enum value "Toast"
	ContentOfferDisplayModeToast string = "Toast"
)

// prop value enum
func (m *ContentOffer) validateDisplayModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contentOfferTypeDisplayModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContentOffer) validateDisplayMode(formats strfmt.Registry) error {

	if err := validate.Required("displayMode", "body", m.DisplayMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateDisplayModeEnum("displayMode", "body", *m.DisplayMode); err != nil {
		return err
	}

	return nil
}

var contentOfferTypeLayoutModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TextOnly","ImageOnly","LeftText","RightText","TopText","BottomText"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contentOfferTypeLayoutModePropEnum = append(contentOfferTypeLayoutModePropEnum, v)
	}
}

const (

	// ContentOfferLayoutModeTextOnly captures enum value "TextOnly"
	ContentOfferLayoutModeTextOnly string = "TextOnly"

	// ContentOfferLayoutModeImageOnly captures enum value "ImageOnly"
	ContentOfferLayoutModeImageOnly string = "ImageOnly"

	// ContentOfferLayoutModeLeftText captures enum value "LeftText"
	ContentOfferLayoutModeLeftText string = "LeftText"

	// ContentOfferLayoutModeRightText captures enum value "RightText"
	ContentOfferLayoutModeRightText string = "RightText"

	// ContentOfferLayoutModeTopText captures enum value "TopText"
	ContentOfferLayoutModeTopText string = "TopText"

	// ContentOfferLayoutModeBottomText captures enum value "BottomText"
	ContentOfferLayoutModeBottomText string = "BottomText"
)

// prop value enum
func (m *ContentOffer) validateLayoutModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, contentOfferTypeLayoutModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ContentOffer) validateLayoutMode(formats strfmt.Registry) error {

	if err := validate.Required("layoutMode", "body", m.LayoutMode); err != nil {
		return err
	}

	// value enum
	if err := m.validateLayoutModeEnum("layoutMode", "body", *m.LayoutMode); err != nil {
		return err
	}

	return nil
}

func (m *ContentOffer) validateStyle(formats strfmt.Registry) error {
	if swag.IsZero(m.Style) { // not required
		return nil
	}

	if m.Style != nil {
		if err := m.Style.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("style")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("style")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this content offer based on the context it is used
func (m *ContentOffer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCallToAction(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStyle(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContentOffer) contextValidateCallToAction(ctx context.Context, formats strfmt.Registry) error {

	if m.CallToAction != nil {
		if err := m.CallToAction.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("callToAction")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("callToAction")
			}
			return err
		}
	}

	return nil
}

func (m *ContentOffer) contextValidateStyle(ctx context.Context, formats strfmt.Registry) error {

	if m.Style != nil {
		if err := m.Style.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("style")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("style")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContentOffer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContentOffer) UnmarshalBinary(b []byte) error {
	var res ContentOffer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
