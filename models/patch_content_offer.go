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

// PatchContentOffer patch content offer
//
// swagger:model PatchContentOffer
type PatchContentOffer struct {

	// Body text of the content offer.
	Body string `json:"body,omitempty"`

	// Properties customizing the call to action button on the content offer.
	CallToAction *PatchCallToAction `json:"callToAction,omitempty"`

	// The display mode of Genesys Widgets when displaying content offer.
	// Enum: [Modal Overlay Toast]
	DisplayMode string `json:"displayMode,omitempty"`

	// Headline displayed above the body text of the content offer.
	Headline string `json:"headline,omitempty"`

	// URL for image displayed to the customer when displaying content offer.
	ImageURL string `json:"imageUrl,omitempty"`

	// The layout mode of the text shown to the user when displaying content offer.
	// Enum: [TextOnly ImageOnly LeftText RightText TopText BottomText]
	LayoutMode string `json:"layoutMode,omitempty"`

	// Properties customizing the styling of the content offer.
	Style *PatchContentOfferStylingConfiguration `json:"style,omitempty"`

	// Title used in the header of the content offer.
	Title string `json:"title,omitempty"`
}

// Validate validates this patch content offer
func (m *PatchContentOffer) Validate(formats strfmt.Registry) error {
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

func (m *PatchContentOffer) validateCallToAction(formats strfmt.Registry) error {
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

var patchContentOfferTypeDisplayModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Modal","Overlay","Toast"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchContentOfferTypeDisplayModePropEnum = append(patchContentOfferTypeDisplayModePropEnum, v)
	}
}

const (

	// PatchContentOfferDisplayModeModal captures enum value "Modal"
	PatchContentOfferDisplayModeModal string = "Modal"

	// PatchContentOfferDisplayModeOverlay captures enum value "Overlay"
	PatchContentOfferDisplayModeOverlay string = "Overlay"

	// PatchContentOfferDisplayModeToast captures enum value "Toast"
	PatchContentOfferDisplayModeToast string = "Toast"
)

// prop value enum
func (m *PatchContentOffer) validateDisplayModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchContentOfferTypeDisplayModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchContentOffer) validateDisplayMode(formats strfmt.Registry) error {
	if swag.IsZero(m.DisplayMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateDisplayModeEnum("displayMode", "body", m.DisplayMode); err != nil {
		return err
	}

	return nil
}

var patchContentOfferTypeLayoutModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["TextOnly","ImageOnly","LeftText","RightText","TopText","BottomText"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchContentOfferTypeLayoutModePropEnum = append(patchContentOfferTypeLayoutModePropEnum, v)
	}
}

const (

	// PatchContentOfferLayoutModeTextOnly captures enum value "TextOnly"
	PatchContentOfferLayoutModeTextOnly string = "TextOnly"

	// PatchContentOfferLayoutModeImageOnly captures enum value "ImageOnly"
	PatchContentOfferLayoutModeImageOnly string = "ImageOnly"

	// PatchContentOfferLayoutModeLeftText captures enum value "LeftText"
	PatchContentOfferLayoutModeLeftText string = "LeftText"

	// PatchContentOfferLayoutModeRightText captures enum value "RightText"
	PatchContentOfferLayoutModeRightText string = "RightText"

	// PatchContentOfferLayoutModeTopText captures enum value "TopText"
	PatchContentOfferLayoutModeTopText string = "TopText"

	// PatchContentOfferLayoutModeBottomText captures enum value "BottomText"
	PatchContentOfferLayoutModeBottomText string = "BottomText"
)

// prop value enum
func (m *PatchContentOffer) validateLayoutModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchContentOfferTypeLayoutModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PatchContentOffer) validateLayoutMode(formats strfmt.Registry) error {
	if swag.IsZero(m.LayoutMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateLayoutModeEnum("layoutMode", "body", m.LayoutMode); err != nil {
		return err
	}

	return nil
}

func (m *PatchContentOffer) validateStyle(formats strfmt.Registry) error {
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

// ContextValidate validate this patch content offer based on the context it is used
func (m *PatchContentOffer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (m *PatchContentOffer) contextValidateCallToAction(ctx context.Context, formats strfmt.Registry) error {

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

func (m *PatchContentOffer) contextValidateStyle(ctx context.Context, formats strfmt.Registry) error {

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
func (m *PatchContentOffer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PatchContentOffer) UnmarshalBinary(b []byte) error {
	var res PatchContentOffer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
