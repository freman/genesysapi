// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ActionProperties action properties
//
// swagger:model ActionProperties
type ActionProperties struct {

	// Accept button text shown to user, used for webchat type action.
	WebchatAcceptText string `json:"webchatAcceptText,omitempty"`

	// Decline button text shown to user, used for webchat type action.
	WebchatDeclineText string `json:"webchatDeclineText,omitempty"`

	// Prompt message shown to user, used for webchat type action.
	WebchatPrompt string `json:"webchatPrompt,omitempty"`

	// Survey provided to the user, used for webchat type action.
	WebchatSurvey *ActionSurvey `json:"webchatSurvey,omitempty"`

	// Title shown to the user, used for webchat type action.
	WebchatTitleText string `json:"webchatTitleText,omitempty"`
}

// Validate validates this action properties
func (m *ActionProperties) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWebchatSurvey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionProperties) validateWebchatSurvey(formats strfmt.Registry) error {
	if swag.IsZero(m.WebchatSurvey) { // not required
		return nil
	}

	if m.WebchatSurvey != nil {
		if err := m.WebchatSurvey.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webchatSurvey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webchatSurvey")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this action properties based on the context it is used
func (m *ActionProperties) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateWebchatSurvey(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActionProperties) contextValidateWebchatSurvey(ctx context.Context, formats strfmt.Registry) error {

	if m.WebchatSurvey != nil {
		if err := m.WebchatSurvey.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("webchatSurvey")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("webchatSurvey")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActionProperties) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActionProperties) UnmarshalBinary(b []byte) error {
	var res ActionProperties
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
